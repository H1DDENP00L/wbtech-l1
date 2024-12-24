package main

import (
	"fmt"
	"reflect"
)

/*
	Разработать программу, которая в рантайме способна определить тип переменной:
	int, string, bool, channel из переменной типа interface{}.


*/

// defineDataType - функция для определения и вывода типа переменной interface{}
func defineDataType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("[%v] is int type\n", v)
	case string:
		fmt.Printf("[%v] is string type\n", v)
	case bool:
		fmt.Printf("[%v] is bool type\n", v)
	case nil:
		fmt.Printf("[%v] is nil\n", v)

	default:
		if reflect.ValueOf(i).Kind() == reflect.Chan {
			fmt.Printf("[%v] is channel type\n", v)
		} else {
			fmt.Printf("[%v] is unknown type\n", v)
		}
	}
}

func main() {

	var i interface{}
	i = "Hello"
	defineDataType(i) // string

	i = 54
	defineDataType(i) // int

	i = false
	defineDataType(i) // bool

	i = make(chan int)
	defineDataType(i) // channel

	i = nil
	defineDataType(i) // nil

	i = 55.555
	defineDataType(i) // unknown

}
