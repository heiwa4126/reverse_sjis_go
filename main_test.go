package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test_reverseString(t *testing.T) {
	wants := "helloこんにちは"
	rc := reverseString("はちにんこolleh")
	if wants != rc {
		t.Fatal("ERROR")
	}
}

func Test_reverseFile(t *testing.T) {

	r := strings.NewReader(
		`ああ、あなたが見ることができると言う、
夜明けの早い光によって
何を誇らしげに歓迎したか
夕暮れの最後のキラリと光る？
`)
	w := bytes.NewBuffer([]byte{})

	if err := reverseFile(r, w); err != nil {
		t.Fatal(err)
	}
	rc := w.String()

	wants := `、う言とるきでがとこる見がたなあ、ああ
てっよに光い早のけ明夜
かたし迎歓にげしら誇を何
？る光とリラキの後最のれ暮夕
`
	if wants != rc {
		t.Fatal("ERROR")
	}
}
