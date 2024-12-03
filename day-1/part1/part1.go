package part1

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
	var firstList []int
	var secondList []int
	for scanner.Scan() {
		s := scanner.Text()
		line := strings.Fields(s)
		firstEl, _ := strconv.Atoi(line[0])
		secEl, _ := strconv.Atoi(line[1])
		firstList = append(firstList, firstEl)
		secondList = append(secondList, secEl)
	}

	slices.Sort(firstList)
	slices.Sort(secondList)

	for i := 0; i < len(firstList); i++ {
		ans += abs(firstList[i] - secondList[i])
	}

	fmt.Println(ans)
}
