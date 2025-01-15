package main

import (
	"fmt"
	"os"
)

func main() {
	var name string

	fmt.Println("Enter your name:")

	fmt.Fscan(os.Stdin, &name)

	fmt.Println("Hello,", name, "!")

}
