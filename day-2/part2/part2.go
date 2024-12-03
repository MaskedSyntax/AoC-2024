package part2

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

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func getFileData() (*os.File, *bufio.Scanner) {
	wd, err := os.Getwd()
	check(err)

	inputPath := filepath.Join(wd, "input.txt")

	file, err := os.Open(inputPath)
	check(err)

	return file, bufio.NewScanner(file)
}

func isValid(word []string) bool {
	neg := 0
	pos := 0
	for i := 1; i < len(word); i++ {
		curr, _ := strconv.Atoi(word[i])
		prev, _ := strconv.Atoi(word[i-1])
		diff := curr - prev
		if diff < 0 {
			if pos > 0 {
				return false
			}
			neg++
		} else {
			if neg > 0 {
				return false
			}
			pos++
		}
		diff = abs(diff)

		if diff > 3 {
			return false
		}

		if diff == 0 {
			return false
		}

	}

	return true
}

func isValidAfterRemoval(words []string) bool {
	if isValid(words) {
		return true
	}

	var temp []string
	for i := 0; i < len(words); i++ {
		temp = append(temp[:0], words[:i]...)
		temp = append(temp, words[i+1:]...)
		if isValid(temp) {
			return true
		}
	}

	return false
}

func Run() {

	file, scanner := getFileData()
	defer file.Close()

	ans := 0
	for scanner.Scan() {
		s := scanner.Text()
		words := strings.Fields(s)

		if isValidAfterRemoval(words) {
			ans++
		}

	}

	fmt.Println(ans)
}
