package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dump_val(mtx [][]node) {
	fmt.Println("=======")
	for _, v := range mtx {
		fmt.Printf("%v %v %v %v\n", v[0].val, v[1].val, v[2].val, v[3].val)
	}
	fmt.Println("=======")
}

func dump_vst(mtx [][]node) {
	fmt.Println("=======")
	for _, v := range mtx {
		fmt.Printf("%v %v %v %v\n", v[0].vst, v[1].vst, v[2].vst, v[3].vst)
	}
	fmt.Println("=======")
}

type node struct {
	val int
	vst bool
}

type space struct {
	r   int
	c   int
	dat [][]node
	run int
}

func (s space) check(i int, j int) int {
	// if a node is zero OR visited true, next node
	if s.dat[i][j].val == 0 || s.dat[i][j].vst {
		return 0
	}

	// if a node is non-zero, find relations (compass), set to visited, recursive
	s.dat[i][j].vst = true

	cnt := 1
	if i-1 >= 0 {
		// n
		cnt += s.check(i-1, j)
	}
	if i-1 >= 0 && j+1 < s.c {
		// ne
		cnt += s.check(i-1, j+1)
	}
	if j+1 < s.c {
		// e
		cnt += s.check(i, j+1)
	}
	if j+1 < s.c && i+1 < s.r {
		// se
		cnt += s.check(i+1, j+1)
	}
	if i+1 < s.r {
		// s
		cnt += s.check(i+1, j)
	}
	if i+1 < s.r && j-1 >= 0 {
		// sw
		cnt += s.check(i+1, j-1)
	}
	if j-1 >= 0 {
		// w
		cnt += s.check(i, j-1)
	}
	if j-1 >= 0 && i-1 >= 0 {
		// nw
		cnt += s.check(i-1, j-1)
	}

	return cnt
}

func (s space) Search() string {

	hi, cnt := 0, 0
	for i, r := range s.dat {
	C:
		for j, v := range r {
			if v.val == 0 {
				s.dat[i][j].vst = true
				continue C
			}
			if v.vst {
				continue C
			}

			// once we have a value, we search the chain
			cnt = s.check(i, j)
			if cnt > hi {
				hi = cnt
			}
		}
	}
	return strconv.Itoa(hi)
}

func Solution(reader *bufio.Reader) string {
	var ln []byte
	var err error

	ln, _, _ = reader.ReadLine()
	rows, _ := strconv.Atoi(strings.TrimSpace(string(ln)))
	ln, _, _ = reader.ReadLine()
	cols, _ := strconv.Atoi(strings.TrimSpace(string(ln)))

	mtx := make([][]node, rows)
	for i := range make([]byte, rows) {
		mtx[i] = make([]node, cols)
		if ln, _, err = reader.ReadLine(); err == nil {
			ss := strings.TrimSpace(string(ln))
			for j, v := range strings.Split(ss, " ") {
				val, _ := strconv.Atoi(v)
				mtx[i][j] = node{val: val, vst: false}
			}
		}
	}

	spc := space{r: rows, c: cols, dat: mtx}
	out := spc.Search()

	return out
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	file := os.Stdin
	nfo, _ := file.Stat()

	reader := bufio.NewReaderSize(file, int(nfo.Size()))
	out := Solution(reader)

	fmt.Println(out)
}
