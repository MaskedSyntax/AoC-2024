package part1

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
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
	for scanner.Scan() {
		s := scanner.Text()
		line := strings.Fields(s)
		fmt.Println(line)

		re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		for _, v := range line {
			matches := re.FindAllStringSubmatch(v, -1)
			fmt.Println(matches)
			for _, match := range matches {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])

				fmt.Println(x, y)
				ans += (x * y)

			}
		}

	}

	fmt.Println(ans)
}
