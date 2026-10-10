package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
	"runtime"
)

// type Interface struct {
// 	Index        int          // positive integer that starts at one, zero is never used
// 	MTU          int          // maximum transmission unit
// 	Name         string       // e.g., "en0", "lo0", "eth0.100"
// 	HardwareAddr HardwareAddr // IEEE MAC-48, EUI-48 and EUI-64 form
// 	Flags        Flags        // e.g., FlagUp, FlagLoopback, FlagMulticast
// }

func main(){
	var GreenText = "\033[32m"
	var ResetText = "\033[0m"

	my_interfaces, err := net.Interfaces() // this seems to be a slice
	if err != nil {
		log.Fatal("could not gather interfaces")
	}

  // why can't I loop over each item?
	for _, iface := range my_interfaces {
		defaultAddress, _ := iface.Addrs()	
		fmt.Println("Name:", iface.Name)
		fmt.Println("Index:", iface.Index)
		fmt.Println("Hardware Address:", iface.HardwareAddr)
		fmt.Println("Network Address:", defaultAddress)
		fmt.Println("Flags:", iface.Flags)
		fmt.Println("MTU:", iface.MTU)
		fmt.Print("\n")
	}

	// let's find the default interface with a little hack if on macOS
	if runtime.GOOS == "darwin"{
		fmt.Println(GreenText + "Current runtime is MacOS" + ResetText)
		cmd := exec.Command("sh", "-c", "route -n get default")
		out, err := cmd.Output()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}else{
			parsedOut := string(out)
			parsedLines := strings.Split(parsedOut, "\n")
			var defaultInterface string;
			for _, line := range parsedLines {
				if strings.Contains(line, "interface:"){
					defaultInterface = strings.Trim(line, " ")
				}
			}
			fmt.Printf("%v The default interface is %v%v\n", GreenText, defaultInterface, ResetText)
		}
	}
}
