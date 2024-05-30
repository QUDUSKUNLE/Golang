package main

import (
	"fmt"
	"bytes"
	"unicode/utf8"
)

func Runes2Bytes(rs []rune) []byte {
	n := 0
	for _, r := range rs {
		n += utf8.RuneLen(r)
	}
	n, bs := 0, make([]byte, n)
	for _, r := range rs {
		n += utf8.EncodeRune(bs[n:], r)
	}
	return bs
}

func main() {
	s := "Color Infection is a fun game."
	sb := []byte(s) // string <-> []byte
	bs := string(sb) // []byte <-> string
	sr := []rune(bs) // string <-> []rune
	_ = bytes.Runes(sb) // []byte <-> []rune

	s = string(sr)
	fmt.Println(sb, bs, sr, s)
	fmt.Println("Done")
}
