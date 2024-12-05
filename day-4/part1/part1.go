package part1

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

func reverse(word string) (result string) {
	for _, v := range word {
		result = string(v) + result
	}
	return
}

func checkVertical(matrix [][]rune, r int, c int) (result int) {
	m := len(matrix)

	word := ""
	// going UP
	for i := 0; i < 4; i++ {
		if r-i >= 0 {
			word += string(matrix[r-i][c])
		}
	}
	if word == "XMAS" {
		result++
	}
	if reverse(word) == "XMAS" {
		result++
	}
	// going DOWN
	word = ""
	for i := 0; i < 4; i++ {
		if r+i < m {
			word += string(matrix[r+i][c])
		}
	}
	if word == "XMAS" {
		result++
	}
	if reverse(word) == "XMAS" {
		result++
	}
	return
}

func checkHorizontal(matrix [][]rune, r int, c int) (result int) {
	n := len(matrix[0])

	var word string

	// going LEFT
	for i := 0; i < 4; i++ {
		if c-i >= 0 {
			word += string(matrix[r][c-i])
		}
	}
	if word == "XMAS" {
		result++
	}
	if reverse(word) == "XMAS" {
		result++
	}
	// going RIGHT
	word = ""
	for i := 0; i < 4; i++ {
		if c+i < n {
			word += string(matrix[r][c+i])
		}
	}
	if word == "XMAS" {
		result++
	}
	if reverse(word) == "XMAS" {
		result++
	}
	return
}

func checkLeftDiagonal(matrix [][]rune, r int, c int) (result int) {
	m := len(matrix)
	n := len(matrix[0])

	var word string

	// Upper Half
	for i := 0; i < 4; i++ {
		if r-i >= 0 && c-i >= 0 {
			word += string(matrix[r-i][c-i])
		}
	}
	if word == "XMAS" {
		result++
	}
	if reverse(word) == "XMAS" {
		result++
	}
	// Lower Half
	word = ""
	for i := 0; i < 4; i++ {
		if r+i < m && c+i < n {
			word += string(matrix[r+i][c+i])
		}
	}

	if word == "XMAS" {
		result++
	}
	if reverse(word) == "XMAS" {
		result++
	}
	return
}

func checkRightDiagonal(matrix [][]rune, r int, c int) (result int) {
	m := len(matrix)
	n := len(matrix[0])

	var word string

	// Upper Half
	for i := 0; i < 4; i++ {
		if r-i >= 0 && c+i < n {
			word += string(matrix[r-i][c+i])
		}
	}
	if word == "XMAS" {
		result++
	}
	if reverse(word) == "XMAS" {
		result++
	}
	// Lower Half
	word = ""
	for i := 0; i < 4; i++ {
		if r+i < m && c-i >= 0 {
			word += string(matrix[r+i][c-i])
		}
	}
	if word == "XMAS" {
		result++
	}
	if reverse(word) == "XMAS" {
		result++
	}
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
			ans += checkVertical(matrix, row, col)
			ans += checkHorizontal(matrix, row, col)
			ans += checkLeftDiagonal(matrix, row, col)
			ans += checkRightDiagonal(matrix, row, col)
		}
	}

	fmt.Println(ans / 2)
}
