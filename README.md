# hw3_bench

Решение задачи **Оптимизация программы**. Coursera, курс: [Разработка веб-сервисов на Go - основы языка.](https://www.coursera.org/learn/golang-webservices-1/home/welcome)

Задание выполнено. Результат бенчмарка:
```
BenchmarkSlow-4               10         112224361 ns/op       320462801 B/op     284142 allocs/op
BenchmarkSolution-8          500           2782432 ns/op          559910 B/op      10422 allocs/op
BenchmarkFast-4             1000           2518339 ns/op          298974 B/op       3889 allocs/op // with easyjson
BenchmarkFast-4              200           7868943 ns/op          494176 B/op       6790 allocs/op // without easyjson

```
