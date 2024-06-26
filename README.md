## Техническое задание

Программа читает из stdin строки, содержащие URL. На каждый URL нужно отправить HTTP-запрос методом GET и посчитать кол-во вхождений строки "Go" в теле ответа. В конце работы приложение выводит на экран общее кол-во найденных строк "Go" во всех переданных URL, например:

```
$ cat urls | go run src/cmd/wb-test/main.go
Count for https://golang.org/: 9
Count for https://golang.org/doc/: 64
Count for https://golang.org/compress/: 6
...
...
Total: 253
```

Каждый URL должен начать обрабатываться сразу после вычитывания и параллельно с вычитыванием следующего. URL должны обрабатываться параллельно, но не более k=5 одновременно. Обработчики URL не должны порождать лишних горутин, т.е. если k=5, а обрабатываемых URL-ов всего 2, не должно создаваться 5 горутин.

Нужно обойтись без глобальных переменных и использовать только стандартную библиотеку.

Для проверки работы используем следующий файл `urls`:

```
https://golang.org/
https://golang.org/doc/
https://golang.org/pkg/compress/
https://golang.org/pkg/compress/gzip/
https://golang.org/pkg/crypto/md5/
https://golang.org/pkg/debug/pe/
https://golang.org/pkg/log/syslog/
https://golang.org/pkg/sort/
https://golang.org/pkg/strconv/
https://golang.org/pkg/strings/
https://golang.org/pkg/sync/
https://golang.org/pkg/strings/
https://golang.org/pkg/time/
https://golang.org/pkg/unicode/
https://golang.org/pkg/unsafe/
https://godoc.org/golang.org/x/benchmarks
https://godoc.org/golang.org/x/net
https://godoc.org/golang.org/x/mobile
```

## Решение

| Файлы   | Для чего       |
| ------- | -------------- |
| main.go | Исходный код   |
| urls    | Входные данные |

Моя система для проверки

```
Linux username 6.5.0-28-generic #29~22.04.1-Ubuntu SMP PREEMPT_DYNAMIC Thu Apr  4 14:39:20 UTC 2 x86_64 x86_64 x86_64 GNU/Linux
```

Пример запуска

```
stepanoff@stepanoff-wpc:/mnt/c/Users/StepanOFF/Desktop/wildberries-test$ cat urls | go run main.go
// Count for https://golang.org/doc/: 116
// Count for https://golang.org/: 246
// Count for https://golang.org/pkg/crypto/md5/: 55
// Count for https://golang.org/pkg/compress/: 42
// Count for https://golang.org/pkg/compress/gzip/: 71
// Count for https://golang.org/pkg/debug/pe/: 80
// Count for https://golang.org/pkg/log/syslog/: 68
// Count for https://golang.org/pkg/sort/: 146
// Count for https://golang.org/pkg/strconv/: 147
// Count for https://golang.org/pkg/strings/: 232
// Count for https://golang.org/pkg/strings/: 232
// Count for https://golang.org/pkg/sync/: 104
// Count for https://golang.org/pkg/unicode/: 113
// Count for https://golang.org/pkg/time/: 239
// Count for https://godoc.org/golang.org/x/benchmarks: 47
// Count for https://godoc.org/golang.org/x/net: 48
// Count for https://golang.org/pkg/unsafe/: 69
// Count for https://godoc.org/golang.org/x/mobile: 77
// Total: 2132
```
