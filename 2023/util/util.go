package util

import (
	"bufio"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetScanner(inputPath string) *bufio.Scanner {
	fd, err := os.Open(inputPath)
	Check(err)

	return bufio.NewScanner(fd)
}
