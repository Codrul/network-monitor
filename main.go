package main

import (
	"log"
	"time"
	"os/signal"
	"syscall"
	"os"
)

/*
Metrics we need to track:
1. Bytes received/sent
2. Packets received/sent
3. Errors
4. Dropped packets
*/

func main(){

	start := time.Now()

	log.SetFlags(log.LstdFlags) // this crazy thing takes bitwise OR for arguments so add | for more args
	log.SetPrefix(GreenText)

	defaultInterface, err := getDefaultInterfaceStruct()
	if err != nil {
		log.Panic("Error I guess")
	}
	// listen to ctrl + c to exit the loop while dev testing
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	monitor: // yes apparently we can name fucking loops in go lang
		for {
			select {
			case <-sigChan:
				log.Println("Shutdown signal received...")
				break monitor

			default:
				bytesReceived := calculateBytesReceived(defaultInterface)
				log.Printf("%v bytes received\n", bytesReceived)
			}
		}
	// clear flags
	program_duration := time.Since(start)
	log.SetFlags(0)
	log.Printf("Program ended after %v\n", program_duration)
}
