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
	var firstList []string
	secondMap := make(map[string]int)
	for scanner.Scan() {
		s := scanner.Text()
		line := strings.Fields(s)
		firstList = append(firstList, line[0])
		secondMap[line[1]]++
	}

	for _, v := range firstList {
		firstEl, _ := strconv.Atoi(v)
		ans += (firstEl * secondMap[v])
	}

	fmt.Println(ans)
}
