package main

import (
	"flag"
	"fmt"
	"go-port-scanner/internal/scanner"
	"math"
	"sync"
)

func main() {
	host := flag.String("host", "localhost", "Target host to scan")
	endPort := math.MaxUint16
	startPort := 1
	flag.Parse()

	ps := scanner.PortScanner{
		Host: *host,
	}

	var wg sync.WaitGroup
	openPorts := make(chan int)

	go func() {
		wg.Wait()
		close(openPorts)
	}()

	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go func(p int) {
			ps.CheckPort(p, openPorts, &wg)
		}(port)
	}

	for port := range openPorts {
		fmt.Printf("Port %d is open\n", port)
	}
}
