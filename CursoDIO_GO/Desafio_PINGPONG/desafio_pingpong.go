package main

import (
	"fmt"
	"time"
)


func main() {
	/*canal1 := make(chan string)
	canal2 := make(chan string)~
	

	go func() {
		for {
			time.Sleep(1 * time.Second)
			canal1 <- "ping"
		}
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			canal2 <- "pong"
		}
	}()

	for {
		select {
			case msg1 := <-canal1:
			fmt.Println("Recebido:", msg1)
			case msg2 := <-canal2:
			fmt.Println("Recebido:", msg2)
		}
	}*/

	canal1 := make(chan string)
	canal2 := make(chan string)
	canalParar := make(chan bool)
	

	go func() {
		for {
			select {
			case <-canalParar:
				return
			default:
				time.Sleep(1 * time.Second)
				canal1 <- "ping"
			}
		}
	}()

	go func() {
		for {
			select {
			case <-canalParar:
				return
			default:
				time.Sleep(1 * time.Second)
				canal2 <- "pong"
			}
		}
	}()

	for i := 0; i <= 10; i++ {
		select {
		case msg1 := <-canal1:
			fmt.Println("Recebido:", msg1)
		case msg2 := <-canal2:
			fmt.Println("Recebido:", msg2)
		}
	}

	close(canalParar)
	time.Sleep(100 * time.Millisecond)


}