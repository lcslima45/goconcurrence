package main

import "fmt"

func hello(done chan bool) {
	fmt.Println("eu sou uma nova goroutine")
	done <- true
}

func hello2(done chan<- bool) {
	fmt.Println("eu sou uma goroutine com chan direcional")
	done <- true
}

func numeros(v chan<- int) {
	for i := 0; i < 10; i++ {
		v <- i
	}
	close(v)
}

func main() {
	done := make(chan bool)
	c := make(chan int)
	go numeros(c)
	go hello(done)

	<-done

	done2 := make(chan bool)
	go hello2(done2)
	<-done2

	for v := range c {
		fmt.Println("valor:", v)
	}

	fmt.Println("eu sou a goro utine principal")
}
