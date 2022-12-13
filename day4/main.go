package day4

import (
	"advent_of_code/shared"
	"fmt"
	"strconv"
	"strings"
)

func Main() {
	sectionScanner := shared.NewFileScanner("day4/data.txt")
	overlappingSectionListCount := 0
	overlappingAnySectionCount := 0
	for sectionScanner.Scan() {
		sectionLists := strings.Split(sectionScanner.Text(), ",")
		sectionList1 := sectionLists[0]
		sectionList2 := sectionLists[1]

		intSectionList1 := parseSection(sectionList1)
		intSectionList2 := parseSection(sectionList2)

		if ListIsContained(intSectionList1, intSectionList2) {
			overlappingSectionListCount++
		} else if ListIsContained(intSectionList2, intSectionList1) {
			overlappingSectionListCount++
		}

		if isAnyInListContained(intSectionList1, intSectionList2) {
			overlappingAnySectionCount++
			continue
		} else if isAnyInListContained(intSectionList2, intSectionList1) {
			overlappingAnySectionCount++
		}
	}

	fmt.Printf("The amount of overlapping lists for part 1 is %v\n", overlappingSectionListCount)
	fmt.Printf("The amount of overlapping lists for part 2 is %v\n", overlappingAnySectionCount)

}

func ListIsContained(list1, list2 []int) bool {
	for _, list1Section := range list1 {
		list1SectionContained := false
		for _, list2Section := range list2 {
			if list1Section == list2Section {
				list1SectionContained = true
			}
		}
		if list1SectionContained == false {
			return false
		}
	}

	return true
}

func isAnyInListContained(list1, list2 []int) bool {
	for _, list1Section := range list1 {
		for _, list2Section := range list2 {
			if list1Section == list2Section {
				return true
			}
		}
	}

	return false
}

func parseSection(sections string) []int {
	minMaxSections := strings.Split(sections, "-")
	minSection, err := strconv.Atoi(minMaxSections[0])

	if err != nil {
		panic(err)
	}

	maxSection, err := strconv.Atoi(minMaxSections[1])

	if err != nil {
		panic(err)
	}

	intSectionList := make([]int, maxSection-minSection+1)
	for i := range intSectionList {
		intSectionList[i] = minSection + i
	}
	return intSectionList
}
