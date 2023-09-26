package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var shapeScore = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var roundTie = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var roundWin = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}

var roundLoss = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}

func roundScore(op, mine string) int {
	if roundTie[op] == mine {
		return 3
	}

	if roundWin[op] == mine {
		return 6
	}

	return 0
}

func main() {
	var res int = 0

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		vals := strings.Split(sc.Text(), " ")
		res += roundScore(vals[0], vals[1]) + shapeScore[vals[1]]
	}

	fmt.Println(res)
}

func main2() {
	var res int = 0

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		vals := strings.Split(sc.Text(), " ")

		if vals[1] == "X" {
			res += shapeScore[roundLoss[vals[0]]]
		} else if vals[1] == "Y" {
			res += 3 + shapeScore[roundTie[vals[0]]]
		} else {
			res += 6 + shapeScore[roundWin[vals[0]]]
		}
	}

	fmt.Println(res)
}
