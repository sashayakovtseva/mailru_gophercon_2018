Конкурс "У кого быстрее" от Mail.Ru Group

Проходит в рамках конференции GopherCon Russia 2018 17 марта, Москва

Давайте представим, что у нас есть список пользователей.

Мы хотим найти пользователей со следующими условиями:
* >= 3-х IP из указанного списка сетей
* >= 3-х браузеров, попадающих под регулярку `Chrome/(60.0.3112.90|52.0.2743.116|57.0.2987.133)`

И вывести их в нужном формате. Смотрите формат и параметры запуска в тестах.

Всё просто, но написать надо максимально производительно. Меряться будем параметром ns/op.

Условия:
* Весь код должен быть помещен в 1 файле, его мы поместим в докер и запустим там тесты и бенчмарк.
* Только stdlib. Ничего другого на сервере нету. Но вы можете добавить в свой файл любой код
* Нельзя использовать cgo
* Нельзя использовать глобальные переменные
* Нельзя хардкодить под задачу, нельзя кешировать результат между банчмарками. Вычисления всегда должны производиться в полном объёме
* У вас есть 4 cpu на сервере
* Все входящие данные гарантированно корректные. Т.е. все сети корректные, json всегда валиден

Работы, которые пытаются подхачить бенчмарк, будут исключены из рейтинга
В случае если ns/op окажется одинаковы у нескольких работ - первенство отдадим тому, кто отправил решение раньше.

https://go2018.highloadcup.ru