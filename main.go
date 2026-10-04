package main

import (
	"fmt"
	"log"
	"net"
)

// type Interface struct {
// 	Index        int          // positive integer that starts at one, zero is never used
// 	MTU          int          // maximum transmission unit
// 	Name         string       // e.g., "en0", "lo0", "eth0.100"
// 	HardwareAddr HardwareAddr // IEEE MAC-48, EUI-48 and EUI-64 form
// 	Flags        Flags        // e.g., FlagUp, FlagLoopback, FlagMulticast
// }

func main(){

	my_interfaces, err := net.Interfaces() // this seems to be a slice
	if err != nil {
		log.Fatal("could not gather interfaces")
	}

  // why can't I loop over each item?
	for _, iface := range my_interfaces {
		fmt.Println("Name:", iface.Name)
		fmt.Println("Index:", iface.Index)
		fmt.Println("Address:", iface.HardwareAddr)
		fmt.Println("Flags:", iface.Flags)
		fmt.Println("MTU:", iface.MTU)
		fmt.Print("\n")
	}
}
