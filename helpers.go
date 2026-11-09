package main

import (
	"log"
	"os/exec"
	"runtime"
	"strings"
)

var GreenText = "\033[32m"
var RedText = "\033[31m"
var ResetText = "\033[0m"


// I thought to use gopsutil but it doesnt have default network detection
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

