package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var maxCals int = 0
	var elfCals int = 0

	for scanner.Scan() {
		val := scanner.Text()
		if len(val) == 0 {
			if maxCals < elfCals {
				maxCals = elfCals
			}
			elfCals = 0
		} else {
			v, _ := strconv.Atoi(val)
			elfCals = elfCals + v
		}
	}

	fmt.Println(maxCals)
}

func main2() {
	scanner := bufio.NewScanner(os.Stdin)

	maxCals := [3]int{0, 0, 0}
	var elfCals int = 0

	for scanner.Scan() {
		val := scanner.Text()
		if len(val) == 0 {
			if maxCals[0] < elfCals {
				maxCals[0], maxCals[1], maxCals[2] = elfCals, maxCals[0], maxCals[1]
			} else if maxCals[1] < elfCals {
				maxCals[1], maxCals[2] = elfCals, maxCals[1]
			} else if maxCals[2] < elfCals {
				maxCals[2] = elfCals
			}
			elfCals = 0
		} else {
			v, _ := strconv.Atoi(val)
			elfCals = elfCals + v
		}
	}
	fmt.Println(maxCals)
	fmt.Println(maxCals[0] + maxCals[1] + maxCals[2])
}
