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

func getInitialPoint(matrix [][]rune) (i int, j int) {
	for ; i < len(matrix); i++ {
		for j = 0; j < len(matrix[0]); j++ {
			if string(matrix[i][j]) == "^" {
				fmt.Println(i, j)
				return
			}
		}
	}

	return
}

func printMat(matrix [][]rune) (cnt int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(string(matrix[i][j]), " ")
			if string(matrix[i][j]) == "X" {
				cnt++
			}
		}
		fmt.Println()
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
		// fmt.Println(line)

		var temp []rune
		for _, v := range line[0] {
			temp = append(temp, v)
		}

		matrix = append(matrix, temp)

	}

	x, y := getInitialPoint(matrix)
	matrix[x][y] = 'X'

	i := x
	j := y

	m := len(matrix)
	n := len(matrix[0])
	dir := 0 // up

	for i >= 0 && i < m && j >= 0 && j < n {
		matrix[i][j] = 'X'
		if dir == 0 { // up
			// check for the above element
			if i-1 >= 0 && string(matrix[i-1][j]) == "#" {
				dir = 1
				j++
			} else {
				i--
			}
		} else if dir == 1 { // right
			// check for the right element
			if j+1 < n && string(matrix[i][j+1]) == "#" {
				dir = 2
				i++
			} else {
				j++
			}
		} else if dir == 2 { // down
			// check for the element below
			if i+1 < m && string(matrix[i+1][j]) == "#" {
				dir = 3
				j--
			} else {
				i++
			}
		} else if dir == 3 { // left
			// check for the left element
			if j-1 >= 0 && string(matrix[i][j-1]) == "#" {
				dir = 0
				i--
			} else {
				j--
			}
		}

	}

	ans = printMat(matrix)

	fmt.Println(ans)
}
