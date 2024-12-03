// Day 3: Mull It Over

package day03

import (
 "io/ioutil"
 "fmt"
 "strings"
 "strconv"
 "regexp"
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

func mul(mem []string) int {
	a, _ := strconv.Atoi(mem[1])
	b, _ := strconv.Atoi(mem[2])
	return a * b

}

func part1(memory []string) {
	total := 0
	r, _ := regexp.Compile(`mul\((?P<A>\d{1,3})\,(?P<B>\d{1,3})\)`)
	for _, mem := range memory {
		matches := r.FindAllStringSubmatch(mem, -1)
		for _, match := range matches {
			total += mul(match)
		}
	}
	fmt.Printf("Part 1 - <%d>\n", total)
}

func part2(memory []string) {
	total := 0
	r, _ := regexp.Compile(`do\(\)|don't\(\)|mul\((?P<A>\d{1,3})\,(?P<B>\d{1,3})\)`)
	enabled := true
	for _, mem := range memory {
		matches := r.FindAllStringSubmatch(mem, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				enabled = true
			} else if match[0] == "don't()" {
				enabled = false
			} else if enabled {
				total += mul(match)
			}
		}
	}
	fmt.Printf("Part 2 - <%d>\n", total)
}

func main() {
	memory := getLines("inputs/day03")
	part1(memory)
	part2(memory)
}
