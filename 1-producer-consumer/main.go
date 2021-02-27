//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, tChan chan<- *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(tChan)
			return
		}

		tChan <- tweet
	}
}

func consumer(tChan <-chan *Tweet) {
	for t := range tChan {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	tChan := make(chan *Tweet)
	// Producer
	go producer(stream, tChan)

	// Consumer
	consumer(tChan)

	fmt.Printf("Process took %s\n", time.Since(start))
}
