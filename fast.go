package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"regexp"
)

// глобальные переменные запрещены
// cgo запрещен

type Info struct {
	Browsers []string `json:"browsers"`
	Hits     []string `json:"hits"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Indx     int
}

func Fast(in io.Reader, w io.Writer, networks []string) {
	regex := regexp.MustCompile(`Chrome/(57.0.2987.133|60.0.3112.90|52.0.2743.116)`)
	dec := json.NewDecoder(in)
	users := make([]*Info, 5000)
	res := make([]*Info, 1000)
	var i, j int
	for dec.More() {
		dec.Decode(&users[i])
		if users[i].check(networks, regex) {
			users[i].Indx = i + 1
			res[j] = users[i]
			j++
		}
		i++
	}
	output(w, j, res[:j])
}

func output(w io.Writer, total int, out []*Info) {
	fmt.Fprintf(w, "Total: %d\n", total)
	for i := range out {

		//email := strings.Split(out[i].Email, "@")
		//fmt.Fprintf(w, "[%d] %s <%s [at] %s>\n", out[i].Indx, out[i].Name, email[0], email[1])

		//email := strings.Replace(out[i].Email, "@", " [at] ", 1)
		fmt.Fprint(w, "[", out[i].Indx, "] ", out[i].Name, " <", out[i].Email, ">\n")
		//fmt.Fprintf(w, "[%d] %s <%s>\n", out[i].Indx, out[i].Name, out[i].Email)
	}
}

func (inf *Info) check(networks []string, regex *regexp.Regexp) bool {
	return inf.checkBrowser(regex) && inf.checkNetwork(networks)
}

func (inf *Info) checkNetwork(networks []string) bool {
	matchCount := 0
	for j := range inf.Hits {
		for i := range networks {
			if networks[i][:5] != inf.Hits[j][:5] {
				continue
			}
			_, n, _ := net.ParseCIDR(networks[i])
			ip := net.ParseIP(inf.Hits[j])
			if n.Contains(ip) {
				matchCount++
				if matchCount == 3 {
					return true
				}
			}
		}
	}
	return false
}

func (inf *Info) checkBrowser(regex *regexp.Regexp) bool {
	matchCount := 0
	for i := range inf.Browsers {
		part := inf.Browsers[i][len(inf.Browsers[i])-34:]
		//log.Println(part)
		if part[0] != 'C' && part[1] != 'C' {
			continue
		}
		if regex.MatchString(part) {
			matchCount++
			if matchCount == 3 {
				return true
			}
		}
	}
	return false
}
