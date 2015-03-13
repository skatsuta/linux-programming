package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "no pattern\n")
		return
	}

	re, err := regexp.Compile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	if len(os.Args) == 2 {
		doGrep(re, os.Stdin)
		return
	}

	for _, f := range os.Args[2:] {
		fp, err := os.Open(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}
		defer fp.Close()

		if e := doGrep(re, fp); e != nil {
			fmt.Fprintf(os.Stderr, "%v\n", e)
			return
		}
	}
}

func doGrep(re *regexp.Regexp, fp *os.File) error {
	r := bufio.NewReader(fp)
	w := bufio.NewWriter(os.Stdout)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		if re.MatchString(line) {
			if _, e := w.WriteString(line); e != nil {
				return e
			}

			if e := w.Flush(); e != nil {
				return e
			}
		}
	}

	return nil
}
