package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RansomNote(reader *bufio.Reader) string {

	var err error
	var ln []byte

	for err != nil {
		ln, _, _ = reader.ReadLine()
	}
	chk := strings.Split(string(ln), " ")
	// fmt.Printf("%#v", chk)

	ln, _, _ = reader.ReadLine()
	spc := strings.Split(string(ln), " ")
	// fmt.Printf("%#v", spc)

	ln, _, _ = reader.ReadLine()
	qry := strings.Split(string(ln), " ")
	// fmt.Printf("%#v", qry)

	// Check the format of the input matches expectations
	var chklen int
	if chklen, err = strconv.Atoi(chk[0]); err != nil {
		panic(err)
	}

	if len(spc) != chklen {
		panic(errors.New("Search space parsed incorrectly"))
	}

	if chklen, err = strconv.Atoi(chk[1]); err != nil {
		panic(err)
	}

	if len(qry) != chklen {
		panic(errors.New("Query parsed incorrectly"))
	}

	// Map each value in hash table => return boolean
	var mp = map[string]bool{}
L:
	for len(qry) > 0 {
		var str string
		str, qry = qry[len(qry)-1], qry[:len(qry)-1]
		mp[str] = false
		for i, val := range spc {
			if val == str {
				spc = append(spc[:i], spc[i+1:]...)
				mp[str] = true
				continue L
			}
		}
		if mp[str] == false {
			return "No"
		}
	}

	return "Yes"
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	file := os.Stdin
	nfo, _ := file.Stat()

	reader := bufio.NewReaderSize(file, int(nfo.Size()))
	out := RansomNote(reader)

	fmt.Println(out)
}
