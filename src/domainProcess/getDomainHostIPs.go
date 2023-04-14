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

// Writing to an output file
func writeBytesToOutput(data1 string, data2 string, out *os.File) int {
	strToWrite := data1 + ": " + data2 + "\n"

	n, err := out.Write([]byte(strToWrite))
	if err != nil {
		log.Fatalf("Didn't manage to write to a file: %v\n", err)
	}

	return n
}

// DNS search execution
func DomainProcess(input string, out string, flagSpec byte) ([]string, error) {
	file, err := os.Open(input)
	if err != nil {
		log.Println("Didn't manage to open the input file")
		return nil, err
	}
	defer file.Close()

	outFile, err := os.OpenFile(out, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Printf("Didn't manage to open/create the output file: %v\n", out)
		return nil, err
	}
	defer outFile.Close()

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
		dnsResult, err = getDomainHosts(inputLog, outFile)
		return dnsResult, err
	case 'i':
		dnsResult, err = getDomainIPs(inputLog, outFile)
		return dnsResult, err
	case 'n':
		dnsResult, err = getNameServers(inputLog, outFile)
		return dnsResult, err
	default:
		return nil, errors.New("Incorrect flag passed")
	}
}

// Returns names mapping to the passed IP address
func getDomainHosts(inputLog []string, out *os.File) ([]string, error) {
	hostsSlice := make([]string, 0)

	for _, ip := range inputLog {

		ipAddr := net.ParseIP(ip)
		if ipAddr != nil {
			hosts, err := getHosts(ip)

			if err == nil {

				for _, hostName := range hosts {
					hostsSlice = append(hostsSlice, hostName)
					writeBytesToOutput(ip, hostName, out)
					fmt.Printf("\tFor IP address %#v found host name: %#v\n", ip, hostName)
				}
			}
		}
	}

	return hostsSlice, nil
}

// Returns IP addresses mapping to the domain host names
func getDomainIPs(inputLog []string, out *os.File) ([]string, error) {
	ipsSlice := make([]string, 0)

	for _, host := range inputLog {

		hostCheck := net.ParseIP(host)
		if hostCheck == nil {
			hosts, err := getIPs(host)

			if err == nil {

				for _, ip := range hosts {
					ipsSlice = append(ipsSlice, ip)
					writeBytesToOutput(host, ip, out)
					fmt.Printf("\tFor domain host name %#v found IP: %#v\n", host, ip)
				}
			}
		}
	}

	return ipsSlice, nil
}

// Getting a domain name server using NS recods of a domain
func getNameServers(inputLog []string, out *os.File) ([]string, error) {
	nameServersSlice := make([]string, 0)

	for _, domain := range inputLog {
		NSs, err := net.LookupNS(domain)
		if err != nil {
			log.Println("Didn't manage to get a domain name server occured from NS records")
			continue
		}

		for _, ns := range NSs {
			nameServersSlice = append(nameServersSlice, ns.Host)
			writeBytesToOutput(domain, ns.Host, out)
			fmt.Printf("\tFor domain host name %#v found domain name server: %#v\n", domain, ns.Host)
		}
		fmt.Println()
	}

	return nameServersSlice, nil
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
func getIPs(host string) ([]string, error) {
	IPs, err := net.LookupHost(host)
	if err != nil {
		return nil, err
	}

	return IPs, nil
}
