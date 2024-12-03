// Day 2: Red-Nosed Reports 

package day02

import (
 "io/ioutil"
 "fmt"
 "strings"
 "strconv"
// "sort"
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

func createList() [][]int {
	lines := getLines("inputs/day02")

	var list [][]int

	for _, value := range lines {
		if len(strings.TrimSpace(value)) == 0 {
			continue
		}
		var subList []int
		parts := strings.Split(value, " ")
		for _, num := range parts {
			out, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			subList = append(subList, out)
		}
		list = append(list, subList)
	}
	

	return list
}


type Direction int8

const (
	Unchanged Direction = iota
	Ascending
	Descending
)

func getDirection(a int, b int) Direction {
	if a == b {
		return Unchanged
	}

	if a > b {
		return Descending
	}

	return Ascending
}

func part1(reports [][]int) {
	safe := 0
	for _, report := range reports {
		ok := isSafe(report)
		if ok {
			safe += 1
		}
	}

	fmt.Printf("Part 1 - <%d>\n", safe)
}

func isSafe(report []int) bool {
		direction := Unchanged
		isSafe := true
		for i := 0; i < len(report) - 1; i += 1 {
			a, b := report[i], report[i + 1]
			newDirection := getDirection(a, b)
			
			if direction == Unchanged {
				direction = newDirection
			} else if direction != newDirection {
				isSafe = false
				break
			}

			diff := absDiff(a, b)
			if diff < 1 || diff > 3 {
				isSafe = false
				break
			}
		}
		return isSafe 
}

func part2(reports [][]int) {
	safe := 0
	for _, report := range reports {
		var tempReport []int

		tempReport = append(tempReport, report...)
		index := 0
		for {
			ok := isSafe(tempReport)
			if ok {
				safe += 1	
				break;
			}

			if index + 1 > len(report) {
				break;
			}

			tempReport = nil
			tempReport = append(tempReport, report...)
			tempReport = append(tempReport[:index], tempReport[index + 1:]...)

			index += 1
		}
	}

	fmt.Printf("Part 2 - <%d>\n", safe)
}

func main() {
	reports := createList()
	part1(reports)
}
