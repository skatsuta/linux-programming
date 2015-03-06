package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s: file name not given\n", os.Args[0])
	}

	for i := 1; i < len(os.Args); i++ {
		doCat(os.Args[i])
	}

	os.Exit(0)
}

const bufferSize = 2048

func doCat(path string) {
	fd, err := syscall.Open(path, os.O_RDONLY, 0644)
	if err != nil {
		die(path)
	}

	buf := make([]byte, bufferSize)
	for {
		n, err := syscall.Read(fd, buf)
		if err != nil {
			die(path)
		}
		if n == 0 {
			break
		}
		if _, e := syscall.Write(syscall.Stdout, buf); e != nil {
			die(path)
		}
	}

	if e := syscall.Close(fd); e != nil {
		die(path)
	}
}

func die(msg string) {
	fmt.Fprint(os.Stderr, msg)
	os.Exit(1)
}
