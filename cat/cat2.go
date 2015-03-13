package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Usage = usage

	// process flags
	flag.Bool("e", false, "display a dollar sign (`$`) at the end of each line")
	flag.Bool("t", false, "display tab characters as `	`")
	flag.Parse()

	for i := 1; i < len(os.Args); i++ {
		fp, err := os.Open(os.Args[i])
		if err != nil {
			die(err)
		}
		defer func() {
			if e := fp.Close(); e != nil {
				die(e)
			}
		}()

		r := bufio.NewReader(fp)
		w := bufio.NewWriter(os.Stdout)

		for {
			c, err := r.ReadByte()
			if err == io.EOF {
				break
			}

			if e := w.WriteByte(c); e != nil {
				die(e)
			}
		}

		if e := w.Flush(); e != nil {
			die(e)
		}
	}

}

func die(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s [-et] [file ...]:\n", os.Args[0])
	flag.PrintDefaults()
}
