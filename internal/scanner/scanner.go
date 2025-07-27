package scanner

import (
	// "fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

type PortScanner struct {
	Host  string
	Ports []int
}

func (ps *PortScanner) CheckPort(port int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	address := net.JoinHostPort(ps.Host, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err == nil {
		conn.Close()
		ch <- port
	}
	// Error logs
	// if err != nil {
	// 		fmt.Printf("Port %d closed or unreachable: %v\n", port, err)
	// }
}
