package main

import (
	"bufio"
	"fmt"
	"os"
)

func convertRune(r rune) int {

	i := int(r)

	if i < int('a') {
		return i - int('A') + 27
	}

	return i - int('a') + 1
}

func main2() {
	sc := bufio.NewScanner(os.Stdin)

	total := 0

	for sc.Scan() {
		m := make(map[rune]bool)
		l := sc.Text()
		var i = 0
		for ; i < len(l)/2; i++ {
			r := []rune(l)[i]
			_, ok := m[r]

			if !ok {
				m[r] = true
			}
		}

		for ; i < len(l); i++ {
			r := []rune(l)[i]
			_, ok := m[r]

			if ok {
				total += convertRune(r)
				break
			}
		}
	}

	fmt.Println(total)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	total := 0

	for {
		if !sc.Scan() {
			break
		}
		m := make(map[rune]int)

		l1 := sc.Text()
		sc.Scan()
		l2 := sc.Text()
		sc.Scan()
		l3 := sc.Text()

		for i := 0; i < len(l1); i++ {
			r := []rune(l1)[i]
			_, ok := m[r]

			if !ok {
				m[r] = 1
			}
		}

		for i := 0; i < len(l2); i++ {
			r := []rune(l2)[i]
			_, ok := m[r]

			if ok {
				m[r] = 2
			}
		}

		for i := 0; i < len(l3); i++ {
			r := []rune(l3)[i]
			_, ok := m[r]

			if ok && m[r] == 2 {
				total += convertRune(r)
				break
			}
		}
	}

	fmt.Println(total)
}
