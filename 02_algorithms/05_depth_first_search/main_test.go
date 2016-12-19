package main_test

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	m "github.com/amay0048/ctci/02_algorithms/04_ice_cream"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileContents(filename string) ([]byte, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	return bytes, err
}

func testCase(t *testing.T, c string) {

	in, _ := fileContents("./test/" + c + "/in.txt")
	sin := string(in)

	out := m.Solution(bufio.NewReaderSize(strings.NewReader(sin), len(in)))
	chk, _ := fileContents("./test/" + c + "/expect.txt")

	ioutil.WriteFile("./test/"+c+"/out.txt", []byte(out), 0777)

	if string(chk) != out {
		t.FailNow()
	}
}

func Test000(t *testing.T) {
	testCase(t, "case000")
}

func Test001(t *testing.T) {
	testCase(t, "case001")
}
