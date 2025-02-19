package main

import (
	"fmt"
	"github.com/EktovVladimir/GolangPracticeOtus/hw7"
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
