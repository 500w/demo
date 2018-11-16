package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{}
	var b interface{}
	a = A()
	b = B()
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.ValueOf(b))
	fmt.Println(a == b)
}

// A return nil
func A() interface{} {
	var a []int
	a = nil
	return a
}

// B return nil
func B() interface{} {
	var b []string
	b = nil
	return b
}
