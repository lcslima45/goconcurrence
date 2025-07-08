package main

import (
	"fmt"
	"time"
)

func numeros(v chan<- int) {
	for i := 0; i < 10; i++ {
		v <- i
		fmt.Printf("número %d escrito no channel\n", i)
	}
	close(v)
}

func main() {
	c := make(chan int, 3)
	go numeros(c)
	time.Sleep(time.Second * 2)
	for v := range c {
		fmt.Println(v)
		time.Sleep(time.Second * 2)
	}
}
