package main

import (
	"log"
	"net"
	"os/exec"
	"runtime"
	"strings"
)

var GreenText = "\033[32m"
var RedText = "\033[31m"
var ResetText = "\033[0m"


// I should have used gopsutil before doing all this but at least it was good exercise lol
func getDefaultInterfaceName() string {
	
	// let's find the default interface with a little hack depending on runtime
	switch os := runtime.GOOS; os {
		case "darwin":
			log.Println(GreenText + "Current runtime is MacOS" + ResetText)
			cmd := exec.Command("sh", "-c", "route -n get default")
			out, err := cmd.Output()
			if err != nil {
				log.Printf("Error: %v\n", err)
			}
			parsedOut := string(out)
			parsedLines := strings.Split(parsedOut, "\n")
			var defaultInterface string;
			for _, line := range parsedLines {
				if strings.Contains(line, "interface:"){
					defaultInterface = strings.Trim(line, " ")
				}
			}
			defaultInterfaceName := strings.TrimSpace(strings.Trim(defaultInterface, "interface:"))
			log.Printf("%vThe default interface is %v%v\n", GreenText, defaultInterfaceName, ResetText)
			return defaultInterfaceName

		case "linux":
			log.Println(GreenText + "Current runtime is Linux" + ResetText)
			cmd := exec.Command("sh", "-c", "ip route show default")
			out, err := cmd.Output()
			if err != nil {
				log.Printf("Error: %v\n", err)
			}
			defaultInterface := string(out)	
			log.Printf("%vThe default interface is %v%v\n", GreenText, string(out), ResetText)
			return defaultInterface
		}
		_:
			return ""
}

func getDefaultInterfaceStruct() (net.Interface, error) {

	my_interfaces, err := net.Interfaces() // this seems to be a slice
	if err != nil {
		log.Fatal("could not gather interfaces")
		return net.Interface{}, err
	}

	numInterfaces := 0

	defaultInterfaceName := getDefaultInterfaceName()	
	var defaultInterface net.Interface;

	for _, iface := range my_interfaces {
		defaultAddress, _ := iface.Addrs()	
		if iface.Name == defaultInterfaceName {
			log.Println("Name:", iface.Name)
			log.Println("Index:", iface.Index)
			log.Println("Hardware Address:", iface.HardwareAddr)
			log.Println("Network Address:", defaultAddress)
			log.Println("Flags:", iface.Flags)
			log.Println("MTU:", iface.MTU)
			defaultInterface = iface
		}
		numInterfaces += 1
	}
	log.Printf(GreenText + "%v interfaces found\n" + ResetText, numInterfaces)
	// let's get the addresses
	addresses, err := defaultInterface.Addrs()
	if err != nil {
		log.Printf("%vAn error happened: %v%v", RedText, err, ResetText)
	}
	for _, address := range addresses {
		log.Printf("%vDefault interface address: %v%v\n", GreenText, address, ResetText)
	}

	return defaultInterface, nil
}
