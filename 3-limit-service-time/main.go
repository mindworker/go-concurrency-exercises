//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import "time"
import "fmt"
//import "sync"

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {

    fmt.Println("handling request for ", u.ID)

    quit := time.NewTimer(time.Second * 10).C
    tickChan := time.NewTicker(time.Second * 1).C

    doneChan := make(chan bool)
    go func() {
        process()
        doneChan <- true
    }()

    for {
        select {
            case <-quit:
                if !u.IsPremium {
                    return false
                }
            case <-tickChan:
                u.TimeUsed += 1
                if u.TimeUsed >= 10 {
                    fmt.Println("time up")
                    return false
                }
            case <-doneChan:
                return true
        }
    }
}

func main() {
	RunMockServer()
}
