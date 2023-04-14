package main

import (
	dns "dns/main/domainProcess"
	"fmt"
	"log"
	"os"
	"strings"
)

type flagRequest struct {
	inputFileFlag   string
	outFileFlag     string
	getDomainNames  bool
	getDomainIPs    bool
	getNamesServers bool
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
		case "-i":
			flagRequestStructure.getDomainIPs = true
		case "-ns":
			flagRequestStructure.getNamesServers = true
		default:
			log.Println("Incorrect argument was passed")
		}
	}

	if flagRequestStructure.getDomainNames {
		fmt.Println("Mapping IP addresses with its domain host names:")
		hosts, err := dns.DomainProcess(flagRequestStructure.inputFileFlag, flagRequestStructure.outFileFlag, 'd')
		if err == nil {
			fmt.Printf("\n*recived hosts: %#v\n\n", hosts)
		}

	} else if flagRequestStructure.getDomainIPs {
		fmt.Println("Mapping domain hosts names with its IP addresses")
		IPs, err := dns.DomainProcess(flagRequestStructure.inputFileFlag, flagRequestStructure.outFileFlag, 'i')
		if err == nil {
			fmt.Printf("\n*recived IPs: %#v\n\n", IPs)
		}

	} else if flagRequestStructure.getNamesServers {
		fmt.Println("List of domain name servers found:")
		nameServers, err := dns.DomainProcess(flagRequestStructure.inputFileFlag, flagRequestStructure.outFileFlag, 'n')
		if err == nil {
			fmt.Printf("*domain name servers: %#v\n\n", nameServers)
		}
	}
}
