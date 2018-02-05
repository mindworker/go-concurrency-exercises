/////////////////////////////////////////////////////////////////////
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
	"sync"
)

func producer(stream Stream, tweetChan chan<- *Tweet, wg *sync.WaitGroup) {
	defer wg.Done()
    fmt.Println("producer")
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
            close(tweetChan)
            return
			//return tweets
		}

        tweetChan <- tweet
		//tweets = append(tweets, tweet)
	}
    //close(tweetChan)
}

func consumer(tweetChan <-chan *Tweet, wg *sync.WaitGroup) {
///tweets []*Tweet) {

	defer wg.Done()
    fmt.Println("consumer")
	for t := range tweetChan {
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

    tweetChan := make(chan *Tweet)

	var wg sync.WaitGroup

    wg.Add(2)
	// Producer
    go producer(stream, tweetChan, &wg)
	//tweets := producer(stream)

	// Consumer
	go consumer(tweetChan, &wg)

	wg.Wait()
	fmt.Printf("Process took %s\n", time.Since(start))
}
