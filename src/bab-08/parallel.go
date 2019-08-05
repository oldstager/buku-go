// https://hub.packtpub.com/concurrency-and-parallelism-in-golang-tutorial/
package main

import (
	"fmt"
	"sync"
	"time"
)

func printTime(msg string) {
	fmt.Println(msg, time.Now().Format("15:04:05"))
}

// Task that will be done over time
func writeMail1(wg *sync.WaitGroup) {
	printTime("Done writing mail #1.")
	wg.Done()
}
func writeMail2(wg *sync.WaitGroup) {
	printTime("Done writing mail #2.")
	wg.Done()
}
func writeMail3(wg *sync.WaitGroup) {
	printTime("Done writing mail #3.")
	wg.Done()
}
func writeMail4(wg *sync.WaitGroup) {
	time.Sleep(time.Nanosecond * 10)
	printTime("Done writing mail #4.")
	wg.Done()
}
func writeMail5(wg *sync.WaitGroup) {
	printTime("Done writing mail #5.")
	wg.Done()
}
func writeMail6(wg *sync.WaitGroup) {
	printTime("Done writing mail #6.")
	wg.Done()
}
func writeMail7(wg *sync.WaitGroup) {
	printTime("Done writing mail #7.")
	wg.Done()
}
func writeMail8(wg *sync.WaitGroup) {
	time.Sleep(time.Nanosecond * 10)
	printTime("Done writing mail #8.")
	wg.Done()
}
func writeMail9(wg *sync.WaitGroup) {
	time.Sleep(time.Nanosecond * 10)
	printTime("Done writing mail #9.")
	wg.Done()
}

// Task done in parallel
func listenForever() {
	for {
		printTime("Listening...")
	}
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(20)

	go listenForever()

	// Give some time for listenForever to start
	time.Sleep(time.Nanosecond * 10)

	// Let's start writing the mails
	go writeMail1(&waitGroup)
	go writeMail2(&waitGroup)
	go writeMail3(&waitGroup)
	go writeMail4(&waitGroup)
	go writeMail5(&waitGroup)
	go writeMail6(&waitGroup)
	go writeMail7(&waitGroup)
	go writeMail8(&waitGroup)
	go writeMail9(&waitGroup)

	waitGroup.Wait()
}
