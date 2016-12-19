package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func pull(i int, a []flv) (t flv, o []flv) {
	if i < 0 || i >= len(a) {
		panic(errors.New("Index out of range"))
	}
	if i == 0 {
		return a[0], a[1:]
	}
	if i == len(a)-1 {
		return a[len(a)-1], a[:len(a)-1]
	}

	t = a[i]
	o = append(o, a[:i-1]...)
	o = append(o, a[i:]...)
	return
}

type flv struct {
	cst int
	id  int
}

func Search(t int, s []flv) []flv {
	if len(s) <= 1 {
		return s
	}

	n := len(s) / 2
	if s[n].cst == t {
		return []flv{s[n]}
	} else if s[n].cst > t {
		return Search(t, s[n:])
	}
	return Search(t, s[:n])
}

type cse struct {
	bgt int
	chk int
	dat []flv
}

func (c cse) Len() int {
	return len(c.dat) - 1
}

func (c cse) Less(a int, b int) bool {
	return c.dat[a].cst > c.dat[b].cst
}

func (c cse) Swap(a int, b int) {
	c.dat[a], c.dat[b] = c.dat[b], c.dat[a]
}

func (c cse) Process() string {
	sort.Sort(c)

	var t flv
	var s []flv
	var a []flv

	for i, _ := range c.dat {
		t, a = pull(i, c.dat)
		rdr := math.Abs(float64(t.cst - c.bgt))
		s = Search(int(rdr), a)
		if s[0].cst+t.cst == c.bgt {
			if t.id < s[0].id {
				return fmt.Sprintf("%v %v", t.id, s[0].id)
			} else {
				return fmt.Sprintf("%v %v", s[0].id, t.id)
			}
		}
	}
	return ""
}

func Solution(reader *bufio.Reader) string {
	var ln []byte
	var err error

	ln, _, _ = reader.ReadLine()
	chk, _ := strconv.Atoi(strings.TrimSpace(string(ln)))
	sets := []cse{}

	/*
		2 chk := v
		4 bgt := v
		5 chk := v
		1 4 5 3 2 cst, id := v, i
	*/

	out := []string{}

	for range make([]byte, chk) {
		set := cse{}
		for i, _ := range make([]byte, 3) {
			if ln, _, err = reader.ReadLine(); err == nil {
				ss := strings.TrimSpace(string(ln))
				switch i {
				case 0:
					set.bgt, _ = strconv.Atoi(ss)
				case 1:
					set.chk, _ = strconv.Atoi(ss)
				case 2:
					set.dat = []flv{}
					for j, s := range strings.Split(ss, " ") {
						n, _ := strconv.Atoi(s)
						set.dat = append(set.dat, flv{cst: n, id: j + 1})
					}
				}
			}
		}

		out = append(out, set.Process())
		sets = append(sets, set)
	}

	if len(sets) != chk {
		panic(errors.New("Error parsing data"))
	}

	return strings.Join(out, "\n")
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	file := os.Stdin
	nfo, _ := file.Stat()

	reader := bufio.NewReaderSize(file, int(nfo.Size()))
	out := Solution(reader)

	fmt.Println(out)
}
