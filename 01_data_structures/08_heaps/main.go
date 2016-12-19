package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type MinIntHeap []int

func (h MinIntHeap) Len() int           { return len(h) }
func (h MinIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinIntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxIntHeap []int

func (h MaxIntHeap) Len() int           { return len(h) }
func (h MaxIntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func RunningMedian(reader *bufio.Reader) string {

	var ln []byte
	var err error

	ln, _, _ = reader.ReadLine()
	chk, _ := strconv.Atoi(string(ln))

	hi := &MinIntHeap{}
	heap.Init(hi)
	lo := &MaxIntHeap{}
	heap.Init(lo)

	out := []string{}
	tot := 0

	for err == nil {
		if ln, _, err = reader.ReadLine(); err == nil {
			i, _ := strconv.Atoi(string(ln))

			tot = lo.Len() + hi.Len()

			// Step 1: Add next item to one of the heaps
			// if next item is smaller than maxHeap root add it to maxHeap,
			// else add it to minHeap
			if tot == 0 {
				heap.Push(lo, i)
			} else if i < (*lo)[0] {
				heap.Push(lo, i)
			} else {
				heap.Push(hi, i)
			}

			// Step 2: Balance the heaps (after this step heaps will be either balanced or
			// one of them will contain 1 more item)
			// if number of elements in one of the heaps is greater than the other by
			// more than 1, remove the root element from the one containing more elements and
			// add to the other one
			dh := math.Abs(float64(hi.Len() - lo.Len()))
			if dh > float64(1) {
				// Balance the heaps
				if hi.Len() > lo.Len() {
					heap.Push(lo, heap.Pop(hi))
				} else {
					heap.Push(hi, heap.Pop(lo))
				}
			}

			if hi.Len() == lo.Len() {
				// even
				m := float64((*lo)[0]+(*hi)[0]) / 2
				out = append(out, fmt.Sprintf("%0.1f", m))
			} else {
				// odd
				if hi.Len() > lo.Len() {
					out = append(out, fmt.Sprintf("%0.1f", float64((*hi)[0])))
				} else {
					out = append(out, fmt.Sprintf("%0.1f", float64((*lo)[0])))
				}
			}
		}
	}

	if chk != tot+1 {
		panic(errors.New("Error reading input"))
	}

	return strings.Join(out, "\n")
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	file := os.Stdin
	nfo, _ := file.Stat()

	reader := bufio.NewReaderSize(file, int(nfo.Size()))
	out := RunningMedian(reader)

	fmt.Println(out)
}
