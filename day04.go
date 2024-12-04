// Day 4: Ceres Search

package day04

import (
 "io/ioutil"
 "fmt"
 "sort"
 "strings"
 //"regexp"
)

func absDiff(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func getLines(path string) [][]rune {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	var wordSearch [][]rune
	for _, line := range lines {
		wordSearch = append(wordSearch, []rune(line))
	}
	return wordSearch
}

func walk(wordSearch [][]rune, x int, y int, dirX int, dirY int, word []rune) bool {
	if len(word) == 0 {
		return true
	}

	checkChar := word[0]
	height := len(wordSearch)
	width := len(wordSearch[0])
	if x < 0 || y < 0  || y >= height || x >= width || len(wordSearch[y])  == 0 || checkChar != wordSearch[y][x] {
		return false
	}

	if len(word) == 1 {
		return true
	}
        
	return walk(wordSearch, x + dirX, y + dirY, dirX, dirY, word[1:])
}

func part1(wordSearch [][]rune) {
	total := 0

	for row := 0; row < len(wordSearch); row += 1 {
		for col := 0; col < len(wordSearch[row]); col += 1 {
			if walk(wordSearch, col, row, 0, -1, []rune{'X', 'M', 'A', 'S'}) {
				total += 1
			}
			if walk(wordSearch, col, row, 1, 0, []rune{'X', 'M', 'A', 'S'}) {
				total += 1
			}
			if walk(wordSearch, col, row, 0, 1, []rune{'X', 'M', 'A', 'S'}) {
				total += 1
			}
			if walk(wordSearch, col, row, -1, 0, []rune{'X', 'M', 'A', 'S'}) {
				total += 1
			}
			if walk(wordSearch, col, row, 1, -1, []rune{'X', 'M', 'A', 'S'}) {
				total += 1
			}
			if walk(wordSearch, col, row, -1, -1, []rune{'X', 'M', 'A', 'S'}) {
				total += 1
			}
			if walk(wordSearch, col, row, 1, 1, []rune{'X', 'M', 'A', 'S'}) {
				total += 1
			}
			if walk(wordSearch, col, row, -1, 1, []rune{'X', 'M', 'A', 'S'}) {
				total += 1
			}
		}
	}

	fmt.Printf("Part 1 - <%d>\n", total)
}

func part2(wordSearch [][]rune) {
	total := 0
	for row := 0; row < len(wordSearch); row += 1 {
		for col := 0; col < len(wordSearch[row]); col += 1 {
			if wordSearch[row][col] != 'A' || len(wordSearch[row]) == 0 {
				continue
			}

			counter:= 0

			// CHECK top left
			if row - 1 >= 0 && col -1 >= 0 {
				char1 := wordSearch[row-1][col-1]
				if row + 1 < len(wordSearch) && col + 1 < len(wordSearch[row + 1]) {
					char2 := wordSearch[row+1][col+1]
					s := string([]rune{char1, char2})
					splitS := strings.Split(s, "")

					sort.Strings(splitS)
					if strings.Join(splitS, "") == "MS" {
						counter += 1
					}
				}
			}

			// CHECK top left
			if row - 1 >= 0 && col + 1 < len(wordSearch[row - 1]) {
				char1 := wordSearch[row-1][col + 1]
				fmt.Println(wordSearch[row -1 ])
				if row + 1 < len(wordSearch) && col - 1 >= 0 && col - 1 < len(wordSearch[row + 1]) {
					char2 := wordSearch[row + 1][col - 1]
					s := string([]rune{char1, char2})
					splitS := strings.Split(s, "")

					sort.Strings(splitS)
					if strings.Join(splitS, "") == "MS" {
						counter += 1
					}
				}
			}

			if counter == 2 {
				total += 1
			}

		}
	}
	fmt.Printf("Part 2 - <%d>\n", total)
}

func main() {
	wordSearch := getLines("inputs/day04")
	part1(wordSearch)
	part2(wordSearch)
}
