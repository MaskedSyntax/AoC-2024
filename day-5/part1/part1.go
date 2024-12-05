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

func getFileData() (*os.File, *bufio.Scanner) {
	wd, err := os.Getwd()
	check(err)

	inputPath := filepath.Join(wd, "input.txt")

	file, err := os.Open(inputPath)
	check(err)

	return file, bufio.NewScanner(file)
}

func checkUpdatesOrder(rules []string, updates []string) bool {
	for i := range updates {
		ele := updates[i]
		for j := i + 1; j < len(updates); j++ {
			check := updates[j] + "|" + string(ele)
			if slices.Contains(rules, check) {
				return false
			}
		}
	}

	return true
}

func Run() {

	file, scanner := getFileData()
	defer file.Close()

	ans := 0
	rules := []string{}
	updates := []string{}
	flag := false
	for scanner.Scan() {
		s := scanner.Text()
		line := strings.Fields(s)
		// fmt.Println(line)

		if len(line) == 0 {
			flag = true
			continue
		}

		if flag {
			updates = append(updates, line[0])
		} else {
			rules = append(rules, line[0])
		}
	}

	for _, v := range updates {
		updates := strings.Split(v, ",")
		if checkUpdatesOrder(rules, updates) {
			n := len(updates) / 2
			x, _ := strconv.Atoi(updates[n])
			ans += x
		}
	}

	fmt.Println(ans)
}
