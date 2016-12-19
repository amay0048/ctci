package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bubble struct {
	a []int
	s int
	p int
}

func (b *Bubble) swap() (err error) {
	i, j := b.p, b.p+1
	if b.a[i] > b.a[j] {
		b.a[i], b.a[j] = b.a[j], b.a[i]
		b.s++
	}
	return
}

func (b *Bubble) sort() {
	b.p = 0
	swp := b.s
	for b.p < len(b.a)-1 {
		b.swap()
		b.p++
	}
	if b.s != swp {
		b.sort()
	}
}

func NewBubble(a []int) *Bubble {
	b := &Bubble{a: a, s: 0, p: 0}
	b.sort()
	return b
}

func BubbleSort(reader *bufio.Reader) string {
	var ln []byte
	var err error

	ln, _, _ = reader.ReadLine()
	chk, _ := strconv.Atoi(string(ln))

	var dat []int
	if ln, _, err = reader.ReadLine(); err == nil {
		nts := strings.Split(string(ln), " ")
		for _, v := range nts {
			n, _ := strconv.Atoi(v)
			dat = append(dat, n)
		}
	}

	if len(dat) != chk {
		panic(errors.New("Error parsing data"))
	}

	bbl := NewBubble(dat)
	out := []string{}
	out = append(out, fmt.Sprintf("Array is sorted in %v swaps.", bbl.s))
	out = append(out, fmt.Sprintf("First Element: %v", bbl.a[0]))
	out = append(out, fmt.Sprintf("Last Element: %v", bbl.a[len(bbl.a)-1]))

	return strings.Join(out, "\n")
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	file := os.Stdin
	nfo, _ := file.Stat()

	reader := bufio.NewReaderSize(file, int(nfo.Size()))
	out := BubbleSort(reader)

	fmt.Println(out)
}
