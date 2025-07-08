package main

import "fmt"

func hello(done chan bool) {
	fmt.Println("eu sou uma nova goroutine")
	done <- true
}
func main() {
	done := make(chan bool)

	go hello(done)

	<-done

	fmt.Println("eu sou a goroutine principal")
}
