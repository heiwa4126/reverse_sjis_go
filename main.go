package main

import (
	"bufio"
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

func reverseFile(in io.Reader, out io.Writer) (err error) {
	s := bufio.NewScanner(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	for s.Scan() {
		_, err = w.WriteString(reverseString(s.Text()) + "\n")
		if err != nil {
			return
		}
	}
	err = s.Err()
	return
}

func reverseSJIS(in io.Reader, out io.Writer) (err error) {
	return reverseFile(
		transform.NewReader(in, japanese.ShiftJIS.NewDecoder()),
		transform.NewWriter(out, japanese.ShiftJIS.NewEncoder()),
	)
}

func main() {
	if err := reverseSJIS(os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
