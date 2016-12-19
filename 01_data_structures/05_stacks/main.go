package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scanLines(scanner *bufio.Scanner) (out []string, err error) {
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	err = scanner.Err()
	return
}

func shift(a []string) (string, []string) {
	return a[0], a[1:len(a)]
}

func unshift(s string, a []string) []string {
	return append([]string{s}, a...)
}

func bal(s []string) bool {
	var tgt string
	stack := []string{}

	// Loop through char at a time
	for _, v := range s {
		switch v {
		// if it's a open brace then add it to the pile
		case "{", "[", "(":
			stack = unshift(v, stack)
		// if it's a closing brace, expect the first of pile to be the balance
		case "}":
			if len(stack) == 0 {
				return false
			}

			tgt, stack = shift(stack)
			if !strings.ContainsRune(tgt, '{') {
				return false
			}
		case "]":
			if len(stack) == 0 {
				return false
			}

			tgt, stack = shift(stack)
			if !strings.ContainsRune(tgt, '[') {
				return false
			}
		case ")":
			if len(stack) == 0 {
				return false
			}

			tgt, stack = shift(stack)
			if !strings.ContainsRune(tgt, '(') {
				return false
			}
		}
	}

	// Finally, if the stack has a length, then it contains unmatched braces, so return false
	if len(stack) > 0 {
		return false
	}

	return true
}

func Balance(reader *bufio.Reader) string {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var ln []byte
	var err error

	ln, _, _ = reader.ReadLine()
	chk, _ := strconv.Atoi(string(ln))

	var lns []string
	for err == nil {
		if ln, _, err = reader.ReadLine(); err == nil {
			lns = append(lns, string(ln))
		}
	}

	if chk != len(lns) {
		panic(errors.New("Error reading input"))
	}

	out := []string{}
	// Range over each line and check the balance
	for _, ln := range lns {
		if bal(strings.Split(ln, "")) == false {
			out = append(out, "NO")
		} else {
			out = append(out, "YES")
		}
	}

	return strings.Join(out, "\n")
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	file := os.Stdin
	nfo, _ := file.Stat()

	reader := bufio.NewReaderSize(file, int(nfo.Size()))
	out := Balance(reader)

	fmt.Println(out)
}
