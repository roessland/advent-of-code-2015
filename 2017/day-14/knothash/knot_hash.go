package knothash

import "fmt"
import "strings"

const N int = 256

type Circ []byte

func NewCirc() Circ {
	c := make(Circ, N)
	for i := 0; i < N; i++ {
		c[i] = byte(i)
	}
	return c
}

func (c Circ) Get(i int) byte {
	return c[(i+N)%N]
}

func (c Circ) Put(i int, val byte) {
	c[(i+N)%N] = val
}

func (c Circ) Reverse(start, length int) {
	for i, j := start, start+length-1; i < j; i, j = i+1, j-1 {
		tmp := c.Get(i)
		c.Put(i, c.Get(j))
		c.Put(j, tmp)
	}
}

func ToDense(s []byte) []byte {
	dense := make([]byte, 16)
	for i, _ := range dense {
		x := s[i*16+0]
		for j := 1; j < 16; j++ {
			x ^= s[i*16+j]
		}
		dense[i] = x
	}
	return dense
}

func Sum(data []byte) []byte {
	input := append(data, 17, 31, 73, 47, 23)
	c := NewCirc()
	curr := 0
	i := 0
	for round := 0; round < 64; round++ {
		for _, length := range input {
			c.Reverse(curr, int(length))
			curr += int(length) + i
			i++
		}
	}
	dense := ToDense([]byte(c))
	return dense
}

func HexDigest(data []byte) string {
	dense := Sum(data)
	strParts := []string{}
	for _, d := range dense {
		strParts = append(strParts, fmt.Sprintf("%02x", d))
	}
	return strings.Join(strParts, "")
}
