package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dump(nds map[rune]TrieNode, pre string) {
	for k, v := range nds {
		if len(v.c) > 0 {
			pre = "-" + pre
			dump(v.c, pre)
		}
		fmt.Printf("%s Key: %s, Rune: %#v Kids: %v\n", pre, string(k), v.r, v.n)
	}
}

type TrieNode struct {
	r rune
	c map[rune]TrieNode
	n int
}

type Trie struct {
	dat TrieNode
	loc *TrieNode
	ndx map[string]int
}

func (t *Trie) Add(s string) {
	if v, ok := t.ndx[s]; ok {
		t.ndx[s] = v + 1
		t.Find(s)
		return
	}

	ndx := t.ndx
	ndx[s] = 1
	t.ndx = ndx

	t.loc = &t.dat
	for _, rn := range s {
		t.upstep(rn)
	}
	t.upstep('*')
}

func (t *Trie) Find(s string) int {
	t.loc = &t.dat
	for _, rn := range s {
		if nd, ok := t.loc.c[rn]; ok {
			t.loc = &nd
		} else {
			return 0
		}
	}
	return t.Count()
}

func (t *Trie) Count() int {
	if t.loc == nil {
		t.loc = &t.dat
	}
	return t.loc.n
}

func (t *Trie) upstep(r rune) {
	if v, ok := t.loc.c[r]; ok {
		v.n++
		t.loc.c[r] = v
		t.loc = &v
	} else {
		nd := TrieNode{r: r, c: map[rune]TrieNode{}, n: 1}
		t.loc.c[r] = nd
		t.loc = &nd
	}
}

func NewTrie() *Trie {
	t := new(Trie)
	t.dat = TrieNode{c: map[rune]TrieNode{}}
	t.ndx = map[string]int{}
	return t
}

func Contacts(reader *bufio.Reader) string {
	var ln []byte
	var err error

	ln, _, _ = reader.ReadLine()
	chk, _ := strconv.Atoi(string(ln))

	var ops [][]string
	for err == nil {
		if ln, _, err = reader.ReadLine(); err == nil {
			s := string(ln)
			ops = append(ops, strings.Split(s, " "))
		}
	}
	if len(ops) != chk {
		panic(errors.New("Error parsing data"))
	}

	trie := NewTrie()
	out := []string{}
	for _, op := range ops {
		switch op[0] {
		case "add":
			trie.Add(op[1])
		case "find":
			out = append(out, strconv.Itoa(trie.Find(op[1])))
		}
	}
	return strings.Join(out, "\n")
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	file := os.Stdin
	nfo, _ := file.Stat()

	reader := bufio.NewReaderSize(file, int(nfo.Size()))
	out := Contacts(reader)

	fmt.Println(out)
}
