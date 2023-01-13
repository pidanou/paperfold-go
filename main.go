package main

import (
	"fmt"

	"github.com/pidanou/go-matrix-lib"
)

func main() {

	// createPaper()

	test := matrix.RotateMatrixClockwise(createPaper())

	fmt.Println(test)

}

func createPaper() [][]string {

	paper := make([][]string, 0)
	var row = 0
	var col = 0

	fmt.Println("Input rows then cols lengths:")
	fmt.Scan(&row, &col)

	for i := 0; i < row; i++ {
		tmp := make([]string, 0)
		for j := 0; j < col; j++ {
			fmt.Printf("Input %dx%v: ", i+1, j+1)
			var input string
			fmt.Scan(&input)
			tmp = append(tmp, input)
		}
		paper = append(paper, tmp)
	}

	fmt.Println(paper)

	return paper
}
