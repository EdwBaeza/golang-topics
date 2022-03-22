package main

import (
	"log"
	"time"
)

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	time1 := time.Second * 4
	time2 := time.Second * 2

	go runMultiplexing(channel1, time1, 1)
	go runMultiplexing(channel2, time2, 2)

	for i := 0; i < 2; i++ {
		select {
		case channelMsg1 := <-channel1:
			log.Println(channelMsg1)
		case channelMsg2 := <-channel2:
			log.Println(channelMsg2)
		}
	}
}

func runMultiplexing(channel chan<- int, duration time.Duration, id int) {
	time.Sleep(duration)
	channel <- id
}
