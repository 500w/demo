package main

import (
	"fmt"
)

// 利用goroutien,求1000以内的素数
func main() {
	origin, wait := make(chan int), make(chan struct{})
	Processor(origin, wait)
	for index := 2; index < 1000; index++ {
		origin <- index
	}
	close(origin)
	<-wait
}

// Processor p
func Processor(seq chan int, wait chan struct{}) {
	go func() {
		// 素数
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		fmt.Println(prime)
		out := make(chan int)
		Processor(out, wait)
		for num := range seq {
			if num%prime != 0 {
				out <- num
			}
		}
		close(out)
	}()
}
