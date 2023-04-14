package domainProcess

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

// DNS search execution
func DomainProcess(input string, out string, flagSpec byte) ([]string, error) {
	file, err := os.Open(input)
	if err != nil {
		log.Println("Didn't manage to open the input file")
		return nil, err
	}

	inputLog := make([]string, 0)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("An error occured during reading the file: %#v\n", err)
			return nil, err
		}

		if line != "\n" {
			line = strings.Trim(line, "\n")
			inputLog = append(inputLog, line)
		}
	}

	dnsResult := make([]string, 0)
	switch flagSpec {
	case 'd':
		dnsResult, err = getDomainHosts(inputLog)
		return dnsResult, err
	case 'i':
		dnsResult, err = getDomainIPs(inputLog)
		return dnsResult, err
	default:
		return nil, errors.New("Not a correct flag passed")
	}
}

// Returns names mapping to the passed IP address
func getDomainHosts(inputLog []string) ([]string, error) {
	hostsSlice := make([]string, 0)

	for _, ip := range inputLog {

		ipAddr := net.ParseIP(ip)
		if ipAddr != nil {
			hosts, err := getHosts(ip)

			if err == nil {

				for _, hostName := range hosts {
					hostsSlice = append(hostsSlice, hostName)
					fmt.Printf("\tFor IP address %#v found host name: %#v\n", ip, hostName)
				}
			}
		}
	}
	fmt.Println()

	return hostsSlice, nil
}

func getDomainIPs(inputLog []string) ([]string, error) {
	return nil, nil
}

// Getting names mapped to an IP address helping function
func getHosts(ip string) ([]string, error) {
	hosts, err := net.LookupAddr(ip)
	if err != nil {
		return nil, err
	}

	return hosts, nil
}

// Getting IP addresses mapped to domain hosts names helping function
func getIPs() {

}
