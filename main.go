package main

import (
	"log"
	"time"
	"os/signal"
	"syscall"
	"os"
	"fmt"
)

/*
Metrics we need to track:
1. Bytes received/sent 											| Got this part down
2. Packets received/sent and packet size    | Can use the same IOCounters struct.
3. Errors
4. Dropped packets
5. Latency																	| can use ping and then parse

Packet size can be (Bytes received / Packets received)
*/

func main(){

	start := time.Now()

	log.SetFlags(log.LstdFlags) // this crazy thing takes bitwise OR for arguments so add | for more args
	log.SetPrefix(GreenText)

	defaultInterface := getDefaultInterfaceName()

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
				bytesPerSecond := calculateBytesReceived(defaultInterface)
				mbps := (bytesPerSecond * 8) / 1_000_000
				fmt.Printf("\rRX: %.2f MB/s (%.2f Mbps)\033[K", bytesPerSecond/1_000_000, mbps)		}
		}
	// clear flags
	program_duration := time.Since(start)
	log.SetFlags(0)
	log.Printf("Program ended after %v\n", program_duration)
}
