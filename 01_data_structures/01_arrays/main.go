package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RotateLeft(reader *bufio.Reader) string {

	var ln []byte

	ln, _, _ = reader.ReadLine()
	npt := strings.Split(string(ln), " ")

	ln, _, _ = reader.ReadLine()
	a := strings.Split(string(ln), " ")

	c, _ := strconv.Atoi(npt[0])
	if c != len(a) {
		panic(fmt.Errorf("Error Parsing Input: %v != %v", c, len(a)))
	}

	d, _ := strconv.Atoi(npt[1])
	for d > 0 {
		a = append(a[1:], a[0])
		d--
	}

	return strings.Join(a[:], " ")
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	file := os.Stdin
	nfo, _ := file.Stat()

	reader := bufio.NewReaderSize(file, int(nfo.Size()))
	out := RotateLeft(reader)

	fmt.Println(out)
}
