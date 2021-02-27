//////////////////////////////////////////////////////////////////////
//
// Given is a mock process which runs indefinitely and blocks the
// program. Right now the only way to stop the program is to send a
// SIGINT (Ctrl-C). Killing a process like that is not graceful, so we
// want to try to gracefully stop the process first.
//
// Change the program to do the following:
//   1. On SIGINT try to gracefully stop the process using
//          `proc.Stop()`
//   2. If SIGINT is called again, just kill the program (last resort)
//

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	proc := MockProcess{}

	// cleanup outer
	exitFunc := func() {
		fmt.Println("exit outer gracefully")
		os.Exit(0)
	}

	go func(ef func()) {
		// cleanup inner
		defer func() {
			fmt.Println("\nexit inner")
			ef()
		}()

		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		<-c
		go proc.Stop()

		time.Sleep(time.Second)
	}(exitFunc)

	proc.Run()
}
