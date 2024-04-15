package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const red, green, blue = 12, 13, 14

func main() {
	startTime := time.Now()
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	games := make([][]map[string]int, 0)
	for fileScanner.Scan() {
		game := parseInput(fileScanner.Text())
		games = append(games, game)

	}
	var valid []int
	for i, game := range games {
		isValid := true
		for _, round := range game {
			if round["blue"] > blue || round["red"] > red || round["green"] > green {
				isValid = false
				break
			} else {
				continue
			}
		}
		if isValid {
			valid = append(valid, i + 1)
		}
	}
	var sum int
	fmt.Println(valid)
	for _, id := range valid {
		sum += id
	}
	fmt.Println(sum)
	fmt.Println(time.Since(startTime))

}

func parseInput(line string) ([]map[string]int) {
	split := strings.Split(line, "Game ")
	after := strings.Split(split[1], ":")
	roundsSlice := strings.Split(after[1], ";")
	rounds := make([]map[string]int, 0)
	for _, r := range(roundsSlice) {
		colorCounts := make(map[string]int)
		counts := strings.Split(r, ",")
		for _, count := range counts {
			countsSlice := strings.Split(strings.Trim(count, " "), " ")
			countInt, _ := strconv.Atoi(countsSlice[0])
			colorCounts[countsSlice[1]] = countInt
		}
		rounds = append(rounds, colorCounts)
	}
	return rounds
}