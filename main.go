package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {

	hostname := "localhost"

	// worker pool channel
	conSema := make(chan struct{}, 10)

	var wg sync.WaitGroup
	for i := 1; i < 65535; i++ {
		wg.Add(1)
		conSema <- struct{}{}
		go func(port int) {
			addr := fmt.Sprintf("%s:%d", hostname, port)
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				// handle error
				fmt.Printf("port %d closed: %v\n", port, err)
			} else {
				fmt.Printf("port %d open\n", port)
				conn.Close()
			}
			<-conSema
			wg.Done()
		}(i)
	}
	wg.Wait()

}
