# Основы языка Go на примере HTTP-сервера

## Базовые моменты:

- Программа на `Go` хранится в одном или нескольких файлах. Каждый файл с программным кодом должен принадлежать какому-нибудь пакету. И вначале каждого файла должно идти объявление пакета, к которому этот файл принадлежит. Пакет объявляется с помощью ключевого слова `package`.

- Функция определяется с помощью ключевого слова `func`, после которого идет имя функции. Затем в скобках идет список параметров. После списка параметров определяются типы возвращаемых из функции значений (если функция возвращает значения). И далее в фигурных скобках идут собственно те операторы, из которых состоит функция.

- Функция `main` — главная функция любой программы на `Go`. По сути всё, что выполняется в программе, выполняется именно функцией `main`.

- Компиляция и выполнение кода происходит через команду в терминале, которая начинается с `go`.
`Go` — это и есть компилятор, которому указывается параметр `run` с указанием файла разширения `.go` для выполнения кода. В случае использования параметра `build` компилятор скопилирует код в выполняемый файл.

Пример кода:

```
package main

func main() {
	println("Hello, Go!")
}
```

- Для хранения данных в программе применяются переменные. Переменная представляет именованный участок в памяти, который может хранить некотоое значение. Для определения переменной применяется ключевое слово `var`, после которого идет имя переменной, а затем указывается ее тип: `var имя_переменной тип_данных`

Пример кода:

```
package main

func main() {
	var str string
  str = "Hello, Go!"
  println(str)
}
```

- Также допустимо краткое определение переменной в формате: `имя_переменной := значение`

Пример кода:

```
package main

func main() {
	var str string
  str = "Hello, Go!"
  println(str)
}
```

- В файле может использоваться функционал из других пакетов. В этом случае используемые пакеты надо импортировать с помощью ключевого слова `import`. Импортируемые пакеты должны идти после объявления пакета для текущего файла

Пример кода:
```
package main

import (
	"log"
	"net/http"
)
```