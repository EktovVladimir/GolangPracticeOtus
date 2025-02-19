package main

import (
	"GolangPracticeOtus/hw7"
	"fmt"
)

func main() {
	var x, y int

	fmt.Print("Ширина = ")
	_, _ = fmt.Scanln(&x)
	fmt.Print("Высота = ")
	_, _ = fmt.Scanln(&y)

	board := hw7.DrawBoard(x, y)

	fmt.Println(board)
}
