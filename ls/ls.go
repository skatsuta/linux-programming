package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		perror(fmt.Errorf("%s: no arguments\n", os.Args[0]))
	}

	for _, path := range os.Args {
		doLs(path)
	}
}

func doLs(path string) {
	fis, err := ioutil.ReadDir(path)
	if err != nil {
		perror(err)
		return
	}

	for _, fi := range fis {
		fmt.Printf("%s\n", fi.Name())
	}
}

func perror(err error) {
	fmt.Fprintf(os.Stderr, "[error] %v\n", err)
}
