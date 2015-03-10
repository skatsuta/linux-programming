package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	l := len(os.Args)

	if l < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s n [file1 file2 ...]\n", os.Args[0])
		os.Exit(1)
	}

	nlines, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if l == 2 {
		doHead(os.Stdin, nlines)
		os.Exit(0)
	}

	for i := 2; i < l; i++ {
		fp, err := os.Open(os.Args[i])
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		defer func() {
			if e := fp.Close(); e != nil {
				fmt.Fprintln(os.Stderr, e.Error())
				os.Exit(1)
			}
		}()

		doHead(fp, nlines)
	}
}

func doHead(fp *os.File, nlines int) {
	if nlines <= 0 {
		return
	}

	r := bufio.NewReader(fp)
	w := bufio.NewWriter(os.Stdout)

	for {
		c, err := r.ReadByte()
		if err == io.EOF {
			break
		}

		if e := w.WriteByte(c); e != nil {
			fmt.Fprintln(os.Stderr, e.Error())
			os.Exit(1)
		}

		if c == '\n' {
			if e := w.Flush(); e != nil {
				fmt.Fprintln(os.Stderr, e.Error())
				os.Exit(1)
			}

			nlines--
			if nlines == 0 {
				return
			}
		}
	}

	if e := w.Flush(); e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
		os.Exit(1)
	}
}
