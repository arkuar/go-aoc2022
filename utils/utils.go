package utils

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
)

func SortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func ParseInt(lines []string, base int) (result []int) {
	for _, l := range lines {
		if i, err := strconv.ParseInt(l, base, 0); err != nil {
			log.Fatal(err)
		} else {
			result = append(result, int(i))
		}
	}
	return
}

// ReadLines reads a file and returns the lines
func ReadLines(filename string) []string {
	f, err := os.Open(filename)
	Check(err)
	defer f.Close()

	sc := bufio.NewScanner(f)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

// ReadIntLines reads a file and returns the lines as integers
func ReadIntLines(filename string) []int {
	f, err := os.Open(filename)
	Check(err)
	defer f.Close()

	sc := bufio.NewScanner(f)
	var lines []int
	for sc.Scan() {
		lines = append(lines, ConvertToInt(sc.Text()))
	}
	return lines
}

// Read lines and split them at the separator
func ReadLinesSplit(filename string, separator rune) []string {
	splitFunc := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := bytes.IndexByte(data, byte(separator)); i >= 0 {
			return i + 1, data[0:i], nil
		}

		if atEOF {
			return len(data), data, nil
		}
		return
	}
	return ReadLinesSplitFunc(filename, splitFunc)
}

// ReadLinesSplit reads a file using a user defined split function and returns the lines
func ReadLinesSplitFunc(filename string, split bufio.SplitFunc) []string {
	f, err := os.Open(filename)
	Check(err)
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Split(split)

	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

// ReadFile reads a file and returns it as a single string
func ReadFile(filename string) string {
	buf, err := ioutil.ReadFile(filename)
	Check(err)

	return string(buf)

}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
