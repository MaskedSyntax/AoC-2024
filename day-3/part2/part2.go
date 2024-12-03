package part2

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFileData() (*os.File, *bufio.Scanner) {
	wd, err := os.Getwd()
	check(err)

	inputPath := filepath.Join(wd, "input.txt")

	file, err := os.Open(inputPath)
	check(err)

	return file, bufio.NewScanner(file)
}

func Run() {

	file, scanner := getFileData()
	defer file.Close()

	ans := 0
	enabled := true
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|(do\(\))|don't\(\)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if match[1] != "" && match[2] != "" {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				if enabled {
					ans += (x * y)
				}
			} else if match[0] == "do()" {
				enabled = true
			} else if match[0] == "don't()" {
				enabled = false
			}
		}

	}

	fmt.Println(ans)
}
