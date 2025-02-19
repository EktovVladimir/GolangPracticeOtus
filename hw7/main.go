package hw7

func DrawBoard(x int, y int) string {
	isHash := false
	str := ""
	for i := 0; i < y; i++ {
		isHash = i%2 == 0
		for j := 0; j < x; j++ {
			if isHash {
				str += "#"
			} else {
				str += " "
			}
			isHash = !isHash
		}
		str += "\n"
	}
	return str
}
