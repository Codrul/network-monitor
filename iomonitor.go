package main

import (
	"log"
	"time"
	"net"

	gnet "github.com/shirou/gopsutil/v4/net" // to avoid conflicts with std package net
)

func calculateBytesReceived(iface net.Interface) float64 {

	stats, err := gnet.IOCounters(true)
	if err != nil {
			log.Fatal(err)
	}

	bytesReceivedStart := 0.0
	for _, stat := range stats {
			if stat.Name == iface.Name {
					log.Printf("Received: %v", stat.BytesRecv)
					bytesReceivedStart += float64(stat.BytesRecv)
			}
	}

	bytesReceivedEnd := 0.0
	time.Sleep(time.Second)
	stats, err = gnet.IOCounters(true)
	for _, stat := range stats {
			if stat.Name == iface.Name {
					log.Printf("Received: %v\n", stat.BytesRecv)
					bytesReceivedEnd += float64(stat.BytesRecv)
			}
	}

	return bytesReceivedEnd - bytesReceivedStart
}
