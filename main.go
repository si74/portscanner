package main

import (
	"time"
	"flag"
	"fmt"
	"net"
	"sync"
)
//IsOpen connect to ports
func IsOpen(ipstr string,port int,timeout time.Duration) bool{
	target:= fmt.Sprintf("%s:%d",ipstr,port)
	conn, err := net.DialTimeout("tcp", target, timeout)
	
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func main() {

	// input user -ip "1.1.1.1"
	var ipstr string
    flag.StringVar(&ipstr, "ip", "localhost", "ip target")
	flag.Parse()

	// worker pool channel

	timeout:=500*time.Millisecond
	var wg sync.WaitGroup
	conSema:= make(chan bool,10)

	for i := 1; i < 65535; i++ {
		wg.Add(1)
		conSema <- true
		go func(port int) {
			
			if IsOpen(ipstr,port,timeout){
				fmt.Printf("Open port: %d \n",port)
			}

			<-conSema
			wg.Done()
		}(i)
	}
	wg.Wait()
}
