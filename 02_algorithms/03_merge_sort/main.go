package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var flips = 0

func Merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			// flips++
			// Inversions is not a simple count, if you are moving from the right
			// you need to count the wole length of left as the inversion
			flips += len(l)
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}

func Sort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	n := len(s) / 2
	l := Sort(s[:n])
	r := Sort(s[n:])
	return Merge(l, r)
}

func MergeSort(reader *bufio.Reader) string {
	var ln []byte
	var err error

	ln, _, _ = reader.ReadLine()
	chk, _ := strconv.Atoi(strings.TrimSpace(string(ln)))
	sets := [][]int{}

	for err == nil {
		if ln, _, err = reader.ReadLine(); err == nil {
			// lchk, _ := strconv.Atoi(string(ln))
			ln, _, err = reader.ReadLine()
			ss := strings.TrimSpace(string(ln))

			set := []int{}
			for _, s := range strings.Split(ss, " ") {
				n, _ := strconv.Atoi(s)
				set = append(set, n)
			}
			sets = append(sets, set)
		}
	}

	if len(sets) != chk {
		panic(errors.New("Error parsing data"))
	}

	out := []string{}
	for _, v := range sets {
		flips = 0
		Sort(v)
		out = append(out, strconv.Itoa(flips))
	}
	return strings.Join(out, "\n")
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	file := os.Stdin
	nfo, _ := file.Stat()

	reader := bufio.NewReaderSize(file, int(nfo.Size()))
	out := MergeSort(reader)

	fmt.Println(out)
}
