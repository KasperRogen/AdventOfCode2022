package day5

import (
	"advent_of_code/shared"
	"bufio"
	"fmt"
	"strconv"
	"unicode"
)

func Main() {
	stackingProcedureScanner := shared.NewFileScanner("day5/data.txt")
	stacks := createStacks(stackingProcedureScanner)
	part1Stacks := make(map[int][]rune)
	part2Stacks := make(map[int][]rune)

	for i, stack := range stacks {
		part1Stacks[i] = stack
		part2Stacks[i] = stack
	}

	for stackingProcedureScanner.Scan() {
		line := stackingProcedureScanner.Text()
		OperateCreateMover(part1Stacks, line, 9000)
		OperateCreateMover(part2Stacks, line, 9001)
	}

	part1Answer := ""
	for i := 0; i < len(part1Stacks); i++ {
		part1Answer += string(part1Stacks[i][len(part1Stacks[i])-1])
	}

	part2Answer := ""
	for i := 0; i < len(part2Stacks); i++ {
		part2Answer += string(part2Stacks[i][len(part2Stacks[i])-1])
	}

	fmt.Println("The answer to part one is", part1Answer)
	fmt.Println("The answer to part two is", part2Answer)

}

func OperateCreateMover(stacks map[int][]rune, operation string, craneModel int) {
	if operation == "" {
		return
	}

	stripCharsTillNumeral(&operation)
	moves, err := strconv.Atoi(readNumerals(&operation))
	if err != nil {
		panic(err)
	}

	stripCharsTillNumeral(&operation)
	from, err := strconv.Atoi(readNumerals(&operation))
	from -= 1
	if err != nil {
		panic(err)
	}

	stripCharsTillNumeral(&operation)
	to, err := strconv.Atoi(readNumerals(&operation))
	to -= 1
	if err != nil {
		panic(err)
	}

	if craneModel == 9000 {
		for i := 0; i < moves; i++ {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}

	if craneModel == 9001 {
		fromStackSize := len(stacks[from])
		stacks[to] = append(stacks[to], stacks[from][fromStackSize-moves:]...)
		stacks[from] = stacks[from][:fromStackSize-moves]
	}

}

func readNumerals(s *string) string {
	numString := ""
	for _, c := range *s {
		if !unicode.IsNumber(rune(c)) {
			return numString
		}
		*s = (*s)[1:]
		numString += string(rune(c))
	}

	return numString
}

func stripCharsTillNumeral(s *string) {
	for _, c := range *s {
		if unicode.IsNumber(rune(c)) {
			return
		}
		*s = (*s)[1:]
	}

	panic("no numeral reached")
}

func printStacks(stacks map[int][]rune) {
	for i := 0; i < len(stacks); i++ {
		fmt.Printf("%d, %s", i+1, string(stacks[i]))
		fmt.Println()
	}
	fmt.Println()
}

func containsBoxes(line string) bool {
	for _, c := range line {
		if c == '[' || c == ']' {
			return true
		}
	}
	return false
}

func createStacks(stackingProcedureScanner *bufio.Scanner) map[int][]rune {
	stacks := make(map[int][]rune)

	for stackingProcedureScanner.Scan() {

		line := stackingProcedureScanner.Text()
		if !containsBoxes(line) {
			return stacks
		}

		colCount := (len(line) / 4) + 1
		for i := 0; i < colCount; i++ {
			col := line[0:3]
			letter := rune(col[1])
			if unicode.IsLetter(letter) {
				stacks[i] = append([]rune{letter}, stacks[i]...)
			}

			if i < colCount-1 {
				line = line[4:]
			}

		}
	}

	panic("didn't catch end of stacks")
}
