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
type games [][]map[string]int

func main() {
	startTime := time.Now()
	games := getGames()
	fmt.Println(getSumOfIds(games))
	fmt.Println(getSumOfPowers(games))
	fmt.Println(time.Since(startTime))
}

func getSumOfIds(games games)(sum int) {
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
	for _, id := range valid {
		sum += id
	}
	return sum
}

func getGames() (games games) {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		game := parseGamesFromLine(fileScanner.Text())
		games = append(games, game)

	}
	return games
}

func parseGamesFromLine(line string) ([]map[string]int) {
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

func getMinCubes(game []map[string]int)(int) {
	var minBlue, minRed, minGreen int
	for _, round := range game {
		for k, v := range round {
			switch {
			case k == "blue":
				if v > minBlue || minBlue == 0 {
					minBlue = v
				}
			case k == "red":
				if v > minRed || minRed == 0 {
					minRed = v
				}	
			case k == "green":
				if v > minGreen || minGreen == 0 {
					minGreen = v
				}
			}
		}
	}
	power := minBlue * minGreen * minRed
	return power
}

func getSumOfPowers(games games) (sum int) {
	for _, game := range games {
		sum += getMinCubes(game)
	}
	return sum
}