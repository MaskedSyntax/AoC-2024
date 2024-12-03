package part1

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isValid(word []string, tone string) bool {
	for i, v := range word {
		if i == 0 {
			continue
		}
		x, _ := strconv.Atoi(v)
		prev, _ := strconv.Atoi(word[i-1])

		if tone == "increasing" {
			if x <= prev || x > prev+3 {
				return false
			}
		} else {
			if prev <= x || prev > x+3 {
				return false
			}
		}
	}
	return true
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
		words := strings.Fields(s)
		a, _ := strconv.Atoi(words[0])
		b, _ := strconv.Atoi(words[1])
		tone := "increasing"
		if a > b {
			tone = "decreasing"
		}

		if isValid(words, tone) {
			ans++
		}

	}

	fmt.Println(ans)
}
