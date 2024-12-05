// Day 5: Print Queue

package day05

import (
 "io/ioutil"
 "fmt"
 //"sort"
 "strings"
 "strconv"
 //"regexp"
)

func absDiff(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func getLines(path string) ([]string, []string){
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	isRules := true
	var rules []string
	var updates []string
	for _, line := range lines {
		if len(line) == 0 {
			isRules = false
			continue
		}

		if isRules {
			rules = append(rules, line)
		} else {
			updates = append(updates, line)
		}
	}
	return rules, updates
}

func IndexOf(arr []int, candidate int) int {
    for index, c := range arr {
        if c == candidate {
            return index
        }
    }
    return -1
}

func part1(rules [][]int, updates [][]int) [][]int {
	total := 0
        var badUpdates [][]int
	for _, update := range updates {
		correct := true
		for _, rule := range rules {
			X, Y := rule[0], rule[1]
			indexX, indexY := IndexOf(update, X), IndexOf(update, Y)
			if indexX == -1 || indexY == -1 {
				continue;
			}

			if indexX > indexY {
				correct = false
			}


		}

		if !correct {
			badUpdates = append(badUpdates, update)
			continue
		}

		total += update[(len(update) - 1) / 2]
	}
	fmt.Printf("Part 1 - <%d>\n", total)
	return badUpdates
}

func insertInt(array []int, value int, index int) []int {
    return append(array[:index], append([]int{value}, array[index:]...)...)
}

func removeInt(array []int, index int) []int {
    return append(array[:index], array[index+1:]...)
}

func moveInt(array []int, srcIndex int, dstIndex int) []int {
    value := array[srcIndex]
    return insertInt(removeInt(array, srcIndex), value, dstIndex)
}

func checkRulesAndFix(rules [][]int, update []int) bool {
                for _, rule := range rules {
                        X, Y := rule[0], rule[1]
                        indexX, indexY := IndexOf(update, X), IndexOf(update, Y)
                        if indexX == -1 || indexY == -1 {
                               continue;
                        }
	
       		        if indexX < indexY {
				continue
       		        }
	
			fmt.Println(rule, X, Y, indexX, indexY)
			update = moveInt(update, indexX, indexY)
			fmt.Println(update)
			return true
		}
		return false
}

func part2(rules [][]int, updates [][]int) {
	total := 0
        for _, update := range updates {
		fmt.Println("-----")
		wrong := true

		for wrong {
			wrong = checkRulesAndFix(rules, update)
		}

                total += update[(len(update) - 1) / 2]
        }
	fmt.Printf("Part 2 - <%d>\n", total)
}

func main() {
	rawRules, rawUpdates := getLines("inputs/day05")

	var rules [][]int
	for _, rule := range rawRules {
		parts := strings.Split(rule, "|")
		val, _ := strconv.Atoi(parts[0])
		val2, _ := strconv.Atoi(parts[1])
		rules = append(rules, []int{val, val2})
	}

	var updates [][]int
	for _, update := range rawUpdates {
		line := strings.Split(update, ",")
		var row []int
		for _, s := range line {
			val, _ := strconv.Atoi(s)
			row = append(row, val)
		}
		updates = append(updates, row)
	}

	badUpdates := part1(rules, updates)
	part2(rules, badUpdates)
}
