package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	// "sort"
	// "strconv"
	// "strings"
	"time"
)

// --- Day 1: Trebuchet?! ---
// Something is wrong with global snow production, and you've been selected to take a look. The Elves have even given you a map; on it,
// they've used stars to mark the top fifty locations that are likely to be having problems.

// You've been doing this long enough to know that to restore snow operations, you need to check all fifty stars by December 25th.

// Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked
// when you complete the first. Each puzzle grants one star. Good luck!

// You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky") and why
// your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky ("of course, where do you think
// snow comes from") when you realize that the Elves are already loading you into a trebuchet ("please hold still, we need to strap you in").

// As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young
//  Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.

// The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves
// now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

// For example:

// 1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet
// In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

// Consider your entire calibration document. What is the sum of all of the calibration values?

func main() {
	startTime := time.Now()
	var sum int
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		sum += getCode(fileScanner.Text())
	}
	fmt.Println(sum)
	fmt.Println(time.Since(startTime))
}

func getCode(line string) (int){
	matches := make(map[int]string)
	substrings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	substrings2 := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	findSubstrings(line, substrings, matches)
	findSubstrings(line, substrings2, matches)
	var keys []int
	for k := range matches {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	code, _ := strconv.Atoi(matches[keys[0]] + matches[keys[len(keys)-1]])
	return code
}

func findSubstrings (line string, substrings []string, matches map[int]string) {
	for i, str := range substrings {
        start := 0
        for {
            index := strings.Index(line[start:], str)
            if index == -1 {
                break
            }
            matches[index+start] = fmt.Sprint(i + 1)
            index += start
            start = index + len(str)
        }
    }
}

// 	for i := 0; i < len(str); i++ {
// 		if str[i:i+len(substring)] == substring {
// 		   return i
// 		}
// 	 }
// 	 return -1
// }