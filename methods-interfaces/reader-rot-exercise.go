package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	l, e := rot.r.Read(b)
	for i, c := range b {
		if c <= 'Z' && c >= 'A' {
			b[i] = (c-'A'+13)%26 + 'A'
		} else if c >= 'a' && c <= 'z' {
			b[i] = (c-'a'+13)%26 + 'a'
		}
	}
	return l, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

}
