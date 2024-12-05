package part2

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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

func checkMASCross(matrix [][]rune, r int, c int) (result int) {
	m := len(matrix)
	n := len(matrix[0])

	var diag1 string
	var diag2 string

	if r-1 >= 0 && c-1 >= 0 {
		diag1 += string(matrix[r-1][c-1])
	}
	diag1 += "A"
	if r+1 < m && c+1 < n {
		diag1 += string(matrix[r+1][c+1])
	}

	if diag1 != "MAS" && diag1 != "SAM" {
		result = 0
		return
	}

	if r-1 >= 0 && c+1 < n {
		diag2 += string(matrix[r-1][c+1])
	}
	diag2 += "A"
	if r+1 < m && c-1 >= 0 {
		diag2 += string(matrix[r+1][c-1])
	}

	if diag2 == "MAS" || diag2 == "SAM" {
		result = 1
		return
	}

	result = 0
	return
}

func Run() {

	file, scanner := getFileData()
	defer file.Close()

	ans := 0
	var matrix [][]rune
	for scanner.Scan() {
		s := scanner.Text()
		line := strings.Fields(s)
		var temp []rune
		for _, v := range line[0] {
			temp = append(temp, v)
		}
		matrix = append(matrix, temp)
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if string(matrix[row][col]) == "A" {
				ans += checkMASCross(matrix, row, col)
			}
		}
	}

	fmt.Println(ans)
}
