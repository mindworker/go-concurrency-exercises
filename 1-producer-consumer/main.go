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

func producer(stream Stream, finished chan<- string) (tweets []*Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
		  c <-  {"","done"}	
		}
		c <- tweet
	}
}

func consumer(c <-chan string, finished chan<- string) {
	for _, t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else if {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
	  } else {
      finished <- "msg"	
    }
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
  c := make(chan string)
  finshed := make(chan string)

	// Producer
	go producer(stream, c)

	// Consumer
	go consumer(tweets, c)
  <- finished
	fmt.Printf("Process took %s\n", time.Since(start))
}
