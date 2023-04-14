package main

import (
	dns "dns/main/domainProcess"
	"fmt"
	"log"
	"os"
	"strings"
)

type flagRequest struct {
	inputFileFlag  string
	outFileFlag    string
	getDomainNames bool
}

var flagRequestStructure flagRequest

func main() {
	flags := os.Args
	if len(flags) == 1 {
		log.Println("Not enough arguments passed!")
	}

	for _, flag := range flags[1:] {

		flagSlice := strings.Split(flag, "=")
		switch flagSlice[0] {
		case "-input":
			flagRequestStructure.inputFileFlag = flagSlice[1]
		case "-out":
			flagRequestStructure.outFileFlag = flagSlice[1]
		case "-d":
			flagRequestStructure.getDomainNames = true
		default:
			log.Println("Incorrect argument was passed")
		}
	}

	if flagRequestStructure.getDomainNames {
		fmt.Println("Mapping IP addresses with its domain host names:")
		hosts, err := dns.DomainProcess(flagRequestStructure.inputFileFlag, flagRequestStructure.outFileFlag, 'd')
		if err == nil {
			fmt.Printf("*recived hosts: %#v\n", hosts)
		}
	}
}
