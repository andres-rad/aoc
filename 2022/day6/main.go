package main

import (
	"bufio"
	"fmt"
	"os"
)

func Unique(bs []byte) bool {
	for i := 0; i < len(bs)-1; i++ {
		for j := i + 1; j < len(bs); j++ {
			if bs[i] == bs[j] {
				return false
			}
		}
	}

	return true
}

func main() {

	// const size = 4 // PART 1
	const size = 14 // PART 2

	window := make([]byte, size)
	sc := bufio.NewReader(os.Stdin)

	for i := 0; i < len(window); i++ {
		b, _ := sc.ReadByte()
		window[i] = b
	}

	if Unique(window) {
		fmt.Println(size)
	}

	count := len(window)
	for {
		b, err := sc.ReadByte()

		window[count%len(window)] = b

		if Unique(window) || err != nil {
			fmt.Println(count + 1)
			break
		}
		count += 1

	}
}
