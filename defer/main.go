package main

import (
	"fmt"
	"time"
)

// defer和panic的关系
func main() {
	defer func() {
		fmt.Println("main defer")
		if r := recover(); r != nil {
			fmt.Println("recover ", r)
		}
	}()
	a()
	panic("main panic")
}

func a() {
	defer func() {
		fmt.Printf("a defer\n")
	}()
	time.Sleep(time.Second * 1)
	panic("a panic")
}
