package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Anagrams(reader *bufio.Reader) string {

	var ln []byte

	ln, _, _ = reader.ReadLine()
	a := string(ln)
	// fmt.Printf("%#v", a)

	ln, _, _ = reader.ReadLine()
	b := string(ln)
	// fmt.Printf("%#v", b)

	var dltd string
	var ngrm string

	dels := 0

L:
	for i, rna := range a {
		for j, rnb := range b {
			if rna == rnb {
				ngrm += b[j : j+1]
				b = b[:j] + b[j+1:]
				continue L
			}
		}
		dltd += a[i : i+1]
	}

	dels += len(dltd)
	dels += len(b)

	return strconv.Itoa(dels)
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	file := os.Stdin
	nfo, _ := file.Stat()

	reader := bufio.NewReaderSize(file, int(nfo.Size()))
	out := Anagrams(reader)

	fmt.Println(out)
}
