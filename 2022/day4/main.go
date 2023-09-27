package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseRow(s string, l1, r1, l2, r2 *int) {
	ranges := strings.Split(s, ",")

	range_ := strings.Split(ranges[0], "-")
	left, _ := strconv.Atoi(range_[0])
	right, _ := strconv.Atoi(range_[1])

	*l1, *r1 = left, right

	range_ = strings.Split(ranges[1], "-")
	left, _ = strconv.Atoi(range_[0])
	right, _ = strconv.Atoi(range_[1])

	*l2, *r2 = left, right

}

func main2() {
	sc := bufio.NewScanner(os.Stdin)

	var l1, r1, l2, r2 int
	count := 0
	for sc.Scan() {
		parseRow(sc.Text(), &l1, &r1, &l2, &r2)

		if l1 <= l2 && r1 >= r2 {
			count++
		} else if l2 <= l1 && r2 >= r1 {
			count++
		}
	}

	fmt.Println(count)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var l1, r1, l2, r2 int
	count := 0
	for sc.Scan() {
		parseRow(sc.Text(), &l1, &r1, &l2, &r2)

		if !(r1 < l2 || l1 > r2) {
			count++
		}
	}

	fmt.Println(count)
}
