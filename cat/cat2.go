package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	for i := 1; i < len(os.Args); i++ {
		fp, err := os.Open(os.Args[i])
		if err != nil {
			die(err.Error())
		}
		defer func() {
			if e := fp.Close(); e != nil {
				die(e.Error())
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
				die(e.Error())
			}
		}

		if e := w.Flush(); e != nil {
			die(e.Error())
		}
	}

}

func die(msg string) {
	fmt.Fprintf(os.Stderr, msg)
	os.Exit(1)
}
