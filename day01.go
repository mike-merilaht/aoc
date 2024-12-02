// Day 1: Historian Hysteria

package day01

import (
 "io/ioutil"
 "fmt"
 "strings"
 "strconv"
 "sort"
)

func absDiff(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func getLines(path string) []string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(content), "\n")
}

func createLists() ([]int, []int) {
	lines := getLines("inputs/day01")

	var leftList []int
	var rightList []int

	for _, value := range lines {
		if len(strings.TrimSpace(value)) == 0 {
			continue
		}
		parts := strings.Split(value, "   ")
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, left)

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		rightList = append(rightList, right)
	}
	
	sort.Sort(sort.IntSlice(leftList))
	sort.Sort(sort.IntSlice(rightList))

	return leftList, rightList
}

func part1(leftList []int, rightList []int ) map[int]int {
	nums := make(map[int]int)
	numCount := len(leftList)
	total := 0
	for i := 0; i < numCount; i++ {
		a, b := leftList[i], rightList[i]
		nums[b] = nums[b] + 1;
		total += absDiff(a, b)
	}

	fmt.Printf("Part 1 - <%d>\n", total)
	return nums
}

func part2(leftList []int, nums map[int]int) {
	total := 0
	for _, value := range leftList {
		total += nums[value] * value
	}
	fmt.Printf("Part 2 - <%d>\n", total)
}

func main() {
	leftList, rightList := createLists()
	nums := part1(leftList, rightList)
	part2(leftList, nums)
}
