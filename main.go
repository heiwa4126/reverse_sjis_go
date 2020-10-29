package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func reverseString(s string) string {
	r := []rune(s)
	i := 0
	j := len(r) - 1
	for i < j {
		r[i], r[j] = r[j], r[i]
		i += 1
		j -= 1
	}
	return string(r)
}

func reverseSjis(in io.Reader, out io.Writer) (err error) {
	s := bufio.NewScanner(transform.NewReader(in, japanese.ShiftJIS.NewDecoder()))
	w := bufio.NewWriter(transform.NewWriter(out, japanese.ShiftJIS.NewEncoder()))
	defer w.Flush()

	for s.Scan() {
		_, err = fmt.Fprintln(w, reverseString(s.Text()))
		if err != nil {
			return
		}
	}
	err = s.Err()
	return
}

func main() {
	if err := reverseSjis(os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
