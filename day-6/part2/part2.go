package part2

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Pos struct {
	X int
	Y int
}

type State struct {
	P   Pos
	Dir int
}

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
			if (matrix[i][j]) == '^' {
				// fmt.Println(i, j)
				return
			}
		}
	}

	return
}

func getGuardPos(matrix [][]rune, stx int, sty int) map[Pos]struct{} {
	positions := make(map[Pos]struct{})
	m, n := len(matrix), len(matrix[0])

	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	d := 0

	// i, j are the starting points
	i, j := stx, sty
	for {
		if matrix[i][j] != '^' {
			s := Pos{X: i, Y: j}
			positions[s] = struct{}{}
		}

		nexti, nextj := i+dir[d][0], j+dir[d][1]
		if nexti < 0 || nexti >= m || nextj < 0 || nextj >= n {
			return positions
		}

		if matrix[nexti][nextj] == '#' {
			d = (d + 1) % 4
		} else {
			i, j = nexti, nextj
		}

	}

}

func isValid(matrix [][]rune, stx, sty int) bool {
	positions := make(map[State]struct{})
	m, n := len(matrix), len(matrix[0])

	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	d := 0

	// i, j are the starting points
	i, j := stx, sty
	for {
		s := State{P: Pos{X: i, Y: j}, Dir: d}
		if _, exists := positions[s]; exists {
			return true
		}

		nexti, nextj := i+dir[d][0], j+dir[d][1]
		if nexti < 0 || nexti >= m || nextj < 0 || nextj >= n {
			return false
		}

		if matrix[nexti][nextj] == '#' {
			d = (d + 1) % 4
		} else {
			i, j = nexti, nextj
		}

		positions[s] = struct{}{}
	}

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

	pos := getGuardPos(matrix, x, y)
	// fmt.Println(pos, len(pos))

	// print char present at each of the pos
	for i := range pos {
		x_, y_ := i.X, i.Y
		// fmt.Println(x_, y_, " : ", string(matrix[x_][y_]))

		matrix[x_][y_] = '#'
		if isValid(matrix, x, y) {
			ans++
		}
		matrix[x_][y_] = '.'
	}

	fmt.Println(ans)
}
