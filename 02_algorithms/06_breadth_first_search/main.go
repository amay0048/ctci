package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	id   int
	rels []int
}

type graph struct {
	s     int // start node
	n     int // num nodes
	e     int // num edges
	spc   map[int]node
	queue []int
}

func (g *graph) path(t int) int {
	// track visited nodes
	vst := map[int]bool{}
	// set the start node to visited
	st := g.spc[g.s]
	(*g).spc[st.id] = st
	// get the children from the start node
	g.queue = st.rels
	// path length
	l := 0
	// Get the len of the queue for each iteration to dequeue post process
	for len(g.queue) > 0 {
		ql := len(g.queue)
		l += 6
		for _, v := range g.queue {
			// get candidate node
			cd := (*g).spc[v]
			// if visited continue
			if _, ok := vst[cd.id]; ok {
				continue
			}
			// add the candidate to visited
			vst[cd.id] = true

			// found the target
			if cd.id == t {
				return l
			}
			// queue children
			g.queue = append(g.queue, cd.rels...)
		}
		// dequeue
		g.queue = g.queue[ql:]
	}
	return -1
}

func (g *graph) Search() string {
	out := []string{}
	for i := range make([]byte, g.n) {
		// Search every node in the space to a link to the start
		if i+1 != g.s {
			out = append(out, strconv.Itoa(g.path(i+1)))
		}
	}
	return strings.Join(out, " ")
}

func NewG() graph {
	return graph{
		spc: map[int]node{},
	}
}

func Solution(reader *bufio.Reader) string {
	var ln []byte
	var err error

	ln, _, _ = reader.ReadLine()
	chk, _ := strconv.Atoi(strings.TrimSpace(string(ln)))

	out := []string{}

	for range make([]byte, chk) {
		g := NewG()

		if ln, _, err = reader.ReadLine(); err == nil {
			ss := strings.Split(strings.TrimSpace(string(ln)), " ")
			g.n, _ = strconv.Atoi(ss[0])
			g.e, _ = strconv.Atoi(ss[1])
		}

		for i := range make([]byte, g.n) {
			g.spc[i+1] = node{id: i + 1}
		}

		for range make([]byte, g.e) {
			if ln, _, err = reader.ReadLine(); err == nil {
				ss := strings.Split(strings.TrimSpace(string(ln)), " ")
				ida, _ := strconv.Atoi(strings.TrimSpace(ss[0]))
				idb, _ := strconv.Atoi(strings.TrimSpace(ss[1]))

				na, nb := g.spc[ida], g.spc[idb]
				na.rels = append(na.rels, idb)
				nb.rels = append(nb.rels, ida)
				g.spc[ida] = na
				g.spc[idb] = nb
			}
		}

		if ln, _, err = reader.ReadLine(); err == nil {
			ss := strings.TrimSpace(string(ln))
			g.s, _ = strconv.Atoi(ss)
		}

		// fmt.Printf("%v\n", g.spc)
		out = append(out, g.Search())
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
