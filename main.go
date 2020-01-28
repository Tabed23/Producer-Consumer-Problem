package main

import (
	"fmt"
	"time"
)

var (
	buffer     [10]int
	counter    = 0
	in         = 0
	out        = 0
	bufferSize = 10
	done       = make(chan bool)
)

func main() {
	go producer()
	go consumer()
	<-done
	fmt.Println("Program teminated")
}

func producer() {
	for i := 0; i <= bufferSize; i++ {
		if counter == bufferSize {
			time.Sleep(time.Second * 1)
		} else {
			buffer[in] = i
			in = (in + 1) % bufferSize
		}
		counter++
	}
	done <- true
}

func consumer() {
	for i := 0; i <= bufferSize; i++ {
		if counter <= 0 {
			time.Sleep(time.Second * 1)
		} else {
			fmt.Printf("Consumer value %d \n", buffer[out])
			out = (out + 1) % bufferSize
		}
		counter--
	}
}
