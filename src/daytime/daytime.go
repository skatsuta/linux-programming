package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

const (
	network = "tcp"
	service = "daytime"
)

func main() {
	args := os.Args
	host := "localhost"
	if len(args) > 1 {
		host = args[1]
	}

	r, err := openConnection(host, service)
	if err != nil {
		perror(err)
		return
	}

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		fmt.Printf("%v", line)
	}
}

func openConnection(host, service string) (*bufio.Reader, error) {
	port, err := net.LookupPort(network, service)
	if err != nil {
		return nil, fmt.Errorf("net.LookupPort(%v, %v) failed: %v", network, service, err)
	}

	addr := fmt.Sprintf("%v:%v", host, port)
	conn, err := net.Dial(network, addr)
	if err != nil {
		return nil, fmt.Errorf("net.Dial(%v, %v) failed: %v", network, addr, err)
	}

	return bufio.NewReader(conn), nil
}

func perror(err error) {
	fmt.Fprintf(os.Stderr, "[error] %v\n", err)
}
