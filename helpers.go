package main

import ( 
	"runtime"
	"os/exec"
	"strings"
	"fmt"
)

var GreenText = "\033[32m"
var ResetText = "\033[0m"

func getRuntime() {
	
	// let's find the default interface with a little hack if on macOS
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println(GreenText + "Current runtime is MacOS" + ResetText)
			cmd := exec.Command("sh", "-c", "route -n get default")
			out, err := cmd.Output()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			parsedOut := string(out)
			parsedLines := strings.Split(parsedOut, "\n")
			var defaultInterface string;
			for _, line := range parsedLines {
				if strings.Contains(line, "interface:"){
					defaultInterface = strings.Trim(line, " ")
				}
			}
			fmt.Printf("%v The default interface is %v%v\n", GreenText, defaultInterface, ResetText)

		case "linux":
			fmt.Println(GreenText + "Current runtime is Linux" + ResetText)
			cmd := exec.Command("sh", "-c", "ip route show default")
			out, err := cmd.Output()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			fmt.Printf("%vThe default interface is %v%v\n", GreenText, string(out), ResetText)

		}
}
