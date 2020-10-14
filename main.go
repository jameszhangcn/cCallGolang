package main
import (
	"fmt"
	"os"
	"os/signal"
	event "events/events"
)
func main() {
	fmt.Println("main enter...")
	exit := make(chan struct{})
	//for coverage test, gofer will support this later
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	done := make(chan struct{})
	go func() {
		<-c
		done <- struct{}{}
	}()
	fmt.Println("This is the test for go coverage")
	//end for coverage test

	event.MdpEventInit()

	FOR_LOOP:
	for {
		select {
		case <-exit:
			break
		case <-done:
			break FOR_LOOP
		default:

		}
	}

	fmt.Println("EBM main exited!!")
}