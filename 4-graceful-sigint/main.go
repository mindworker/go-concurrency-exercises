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
)

func main() {

    signal_chan := make(chan os.Signal, 1)
    signal.Notify(signal_chan,
        syscall.SIGHUP,
        syscall.SIGINT,
        syscall.SIGTERM,
        syscall.SIGQUIT)

    exit_chan := make(chan int)
    go func() {

	    // Create a process
	    proc := MockProcess{}

	    // Run the process (blocking)
	    go proc.Run()

        count := 0
        for {
            s := <-signal_chan
            switch s {
/*
            // kill -SIGTERM XXXX
            case syscall.SIGTERM:
                fmt.Println("sigterm")

            // kill -SIGQUIT XXXX
            case syscall.SIGQUIT:
                fmt.Println("sigquit")

            // kill -SIGHUP XXXX
            case syscall.SIGHUP:
                fmt.Println("hungup")
*/
            // kill -SIGINT XXXX or Ctrl+c
            case syscall.SIGINT:
                fmt.Println("CTRL+C")
                count++
                if count == 2 {
                    exit_chan <- 0
                }
                go proc.Stop()

            default:
                fmt.Println("Unknown signal.")
                exit_chan <- 1
            }
        }
    }()

    code := <-exit_chan
    os.Exit(code)
}
