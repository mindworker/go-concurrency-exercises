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

import (
	"log"
	"sync"
	"time"
)

const timeAllowance = 10

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
	mu        sync.Mutex
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User, pid int) bool {
	u.mu.Lock()
	defer u.mu.Unlock()

	done := make(chan bool)
	tm := time.NewTicker(time.Second)

	go func() {
		log.Println("start: job id", pid, "user id:", u.ID, "time used:", u.TimeUsed)
		process()
		log.Println("complete: job id", pid, "user id:", u.ID, "time used:", u.TimeUsed)
		done <- true
	}()

	for {
		select {
		case <-tm.C:
			u.TimeUsed++
			if !u.IsPremium && u.TimeUsed > timeAllowance {
				return false
			}
		case <-done:
			return true
		}
	}
}

func main() {
	RunMockServer()
}
