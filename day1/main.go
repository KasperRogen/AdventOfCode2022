package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type elf struct {
	kcals int
}

func Main() {
	kcalReader := readCals("day1/data.txt")
	elfs := createElves(kcalReader)
	sort.Slice(elfs, func(i, j int) bool {
		return elfs[i].kcals > elfs[j].kcals
	})

	fmt.Printf("The largest amount of kcals carried by an elf is %d", elfs[0].kcals)
	fmt.Printf("The total amount of kcals carried by the top three elfs is %d", elfs[0].kcals+elfs[1].kcals+elfs[2].kcals)

}

func readCals(filePath string) *bufio.Scanner {
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner //TODO: CLOSE FILE AFTER READING
}

func createElves(kcalReader *bufio.Scanner) []elf {
	var elfs []elf

	var currentCals int
	for kcalReader.Scan() {
		if kcalReader.Text() == "" {
			newElf := elf{kcals: currentCals}
			elfs = append(elfs, newElf)
			currentCals = 0
			continue
		}

		intKcals, err := strconv.Atoi(kcalReader.Text())
		if err != nil {
			fmt.Println(err)
		}

		currentCals += intKcals
	}

	return elfs
}
