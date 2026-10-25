package day01

import (
	"bufio"
	"errors"
	"log"
	"os"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day01/input.txt"
	if len(inputFile) > 0 {
		file = inputFile[0]
	}

	switch part {
	case 1:
		return SolvePart1(file, logger), nil
	case 2:
		return SolvePart2(file, logger), nil
	default:
		return -1, errors.New("incorrect part number recieved")
	}
}

func SolvePart1(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	answer := 0

	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()

		logger.Printf("%d : %s", i, line)
		var left, right int
		foundL, foundR := false, false
		idxL := 0
		idxR := len(line) - 1
		for idxL <= idxR {
			if foundL && foundR {
				break
			}

			if !foundL {
				if line[idxL] >= '0' && line[idxL] <= '9' {
					left = int(line[idxL] - '0')
					foundL = true
				} else {
					idxL++
				}
			}

			if !foundR {
				if line[idxR] >= '0' && line[idxR] <= '9' {
					right = int(line[idxR] - '0')
					foundR = true
				} else {
					idxR--
				}
			}
		}

		if !(foundL && foundR) {
			logger.Fatalf("Left and right digits not found in %s", line)
		}

		logger.Print(10*left + right)
		answer += 10*left + right
	}
	return answer
}

var digitWords = map[int]string{
	0: "zero",
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

func isValidDigit(idx int, line string) (int, bool) {
	// Return if number
	if line[idx] >= '0' && line[idx] <= '9' {
		return int(line[idx] - '0'), true
	}

	// Return if spells out a number
	for i, word := range digitWords {
		if (len(word) + idx) > len(line) {
			continue
		}

		if word == line[idx:len(word)+idx] {
			return i, true
		}
	}

	return 0, false
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	answer := 0

	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()

		logger.Printf("%d : %s", i, line)
		var left, right int
		foundL, foundR := false, false
		idxL := 0
		idxR := len(line) - 1
		for idxL <= idxR {
			if foundL && foundR {
				break
			}

			if !foundL {
				if l, ok := isValidDigit(idxL, line); ok {
					left = l
					foundL = true
				} else {
					idxL++
				}
			}

			if !foundR {
				if r, ok := isValidDigit(idxR, line); ok {
					right = r
					foundR = true
				} else {
					idxR--
				}
			}
		}

		if !(foundL && foundR) {
			logger.Fatalf("Left and right digits not found in %s", line)
		}

		logger.Print(10*left + right)
		answer += 10*left + right
	}
	return answer
}
