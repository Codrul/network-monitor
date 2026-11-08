package main

import (
	"log"
	"time"
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

	log.SetFlags(log.LstdFlags | log.Ltime, )
	log.SetPrefix(GreenText)

	defaultInterface, err := getDefaultInterfaceStruct()
	if err != nil {
		log.Panic("Error I guess")
	}
	log.Printf("%v", defaultInterface.Name)

	program_duration := time.Since(start)
	// clear flags
	log.SetFlags(0)
	log.Printf("Program ended after %v", program_duration)
}
