package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type q struct {
	dat []string
	sp  int
	ep  int
}

func (q *q) Enqueue(s string) {
	q.dat[q.ep] = s
	q.ep++
}

func (q *q) Dequeue() string {
	q.sp++
	return q.dat[q.sp]
}

func (q *q) Peek() string {
	return q.dat[q.sp]
}

func NewQ() *q {
	nq := new(q)
	nq.dat = make([]string, 100000)
	nq.sp = 0
	nq.ep = 0
	return nq
}

func Queues(reader *bufio.Reader) string {

	var ln []byte
	var err error

	ln, _, _ = reader.ReadLine()
	chk := string(ln)

	lns := []string{}

	for err == nil {
		if ln, _, err = reader.ReadLine(); err == nil {
			lns = append(lns, string(ln))
		}
	}

	n, err := strconv.Atoi(chk)
	if err != nil {
		panic(err)
	}

	if n != len(lns) {
		panic(fmt.Errorf("Error reading input: %#v != %#v", n, len(lns)))
	}

	queue := NewQ()
	out := []string{}

	for _, ln := range lns {
		dat := strings.Split(ln, " ")
		switch dat[0] {
		case "1":
			queue.Enqueue(dat[1])
		case "2":
			queue.Dequeue()
		case "3":
			out = append(out, queue.Peek())
		}
	}

	return strings.Join(out, "\n")
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	file := os.Stdin
	nfo, _ := file.Stat()

	reader := bufio.NewReaderSize(file, int(nfo.Size()))
	out := Queues(reader)

	fmt.Println(out)
}
