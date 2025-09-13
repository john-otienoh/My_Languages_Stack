package main

import (
	"fmt"
	"time"
)

func slowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("Sleeper() finished")
}

func slowFuncChannel(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "SlowFuncChannel() finished"
}
func receiver(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func pinger(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "ping"
		<-t.C
	}
}

// Using Channels as Function Arguments
func channelReader(messages <-chan string) {
	msg := <-messages
	fmt.Println(msg)
}

func channelWriter(messages chan<- string) {
	messages <- "Hello World"
}
func channelReaderAndWriter(messages chan string) {
	msg := <-messages
	fmt.Println(msg)
	messages <- "Hello World"
}

func ping1(c chan string) {
	time.Sleep(time.Second * 1)
	c <- "ping on channel1"
}
func ping2(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "ping on channel2"
}
func sender(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "I'm sending a message"
		<-t.C
	}
}
func main() {
	// Simple Goroutine Example
	go slowFunc()
	fmt.Println("I am not shown until slowFunc() completes")
	time.Sleep(time.Second * 3)

	// Initializing a channel
	c := make(chan string)
	go slowFuncChannel(c)

	// Receiving a message on a channel
	msg := <-c
	fmt.Println(msg)

	// Using Buffered Channel
	messages := make(chan string, 2)
	messages <- "Hello"
	messages <- "World"

	// Receiving two messages in Buffered Channel
	close(messages)
	fmt.Println("Pushed 2 messages into Channel with no receivers")
	time.Sleep(time.Second * 1)
	receiver(messages)

	// Channels and Flow Control
	msg1 := make(chan string)
	go pinger(msg1)
	msg2 := <-msg1
	fmt.Println(msg2)

	// Running a Process Indefinitely
	for i := 0; i < 5; i++ {
		msg3 := <-msg1
		fmt.Println(msg3)
	}

	// Employing the select Statement
	channel1 := make(chan string)
	channel2 := make(chan string)

	go ping1(channel1)
	go ping2(channel2)

	select {
	case msg3 := <-channel1:
		fmt.Println("received", msg3)

	case msg4 := <-channel2:
		fmt.Println("received", msg4)

	case <-time.After(500 * time.Millisecond):
		fmt.Println("no messages received. giving up.")
	}

	msg5 := make(chan string)
	stop := make(chan bool)
	go sender(msg5)

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Time's Up")
		stop <- true
	}()
	for {
		select {
		case <-stop:
			return
		case msg6 := <-msg5:
			fmt.Println(msg6)
		}
	}

}
