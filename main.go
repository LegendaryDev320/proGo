package main

import (
	"fmt"
	"time"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
func writeToChannel(channel chan<- string) {
	timer := time.NewTimer(time.Minute * 10)
	go func() {
		time.Sleep(time.Second * 2)
		Printfln("Resetting timer")
		timer.Reset(time.Second)
	}()
	Printfln("Waiting for initial duration...")
	<-timer.C
	Printfln("Initial duration elapsed.")
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		//time.Sleep(time.Second * 3)
	}
	close(channel)
}
func main() {
	nameChannel := make(chan string)
	go writeToChannel(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}
