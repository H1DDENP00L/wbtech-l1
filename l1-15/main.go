package main

import (
	"bytes"
	"fmt"
	"strings"
)

/*
	К каким негативным последствиям может привести данный фрагмент кода,
	и как это исправить?
	Приведите корректный пример реализации.

Code:
```go
var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
```

Ответ:

Функция someFunc присваивает глобальной переменной justString срез (v[:100]) от строки,
созданной createHugeString. Срезы в Go не копируют данные, а предоставляют представление существующего массива.
justString продолжает ссылаться на массив байтов исходной, большой строки.
Это приводит к утечке памяти (сходный массив и вся большая строка остаются в памяти)

В Go строки иммутабельные, и при динамическом создании строк здесь, эффективно применять буфер.
*/

var regularString string

func createHugeString(stringLen int) string {
	var buf bytes.Buffer
	for i := 0; i < stringLen; i++ {
		buf.WriteString("z")
	}
	return buf.String()
}

func someFunc() {
	v := createHugeString(1 << 10)
	regularString = strings.Clone(v[:100])
	fmt.Printf("Длина строки = [%d];\nСама строка : <%s>\n", len(regularString), regularString)
}

func main() {
	someFunc()
}
