package shared

import (
	"bufio"
	"fmt"
	"os"
)

func NewFileScanner(filePath string) *bufio.Scanner {
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner //TODO: CLOSE FILE AFTER READING
}
