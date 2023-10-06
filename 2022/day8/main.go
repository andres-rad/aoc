package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	m := [][][2]int{}

	for sc.Scan() {
		line := sc.Text()
		ps := strings.Split(line, "")
		row := make([][2]int, len(ps))

		for i := range ps {
			n, _ := strconv.Atoi(ps[i])
			row[i] = [2]int{n, 0}
		}

		m = append(m, row)
	}

	count := 0

	for i := range m {
		maxH := -1
		for j := len(m[i]) - 1; j >= 0; j-- {
			if m[i][j][0] > maxH {
				if m[i][j][1] == 0 {
					count++
					m[i][j][1] += 1
				}
				maxH = m[i][j][0]
			}
		}
	}

	for i := range m {
		maxH := -1
		for j := range m[i] {
			if m[i][j][0] > maxH {
				if m[i][j][1] == 0 {
					count++
					m[i][j][1] += 1
				}
				maxH = m[i][j][0]
			}
		}
	}

	for i := range m[0] {
		maxH := -1
		for j := range m {
			if m[j][i][0] > maxH {
				if m[j][i][1] == 0 {
					count++
					m[j][i][1] += 1
				}
				maxH = m[j][i][0]
			}
		}
	}

	for i := range m[0] {
		maxH := -1
		for j := len(m) - 1; j >= 0; j-- {
			if m[j][i][0] > maxH {
				if m[j][i][1] == 0 {
					m[j][i][1] += 1
					count++
				}

				maxH = m[j][i][0]
			}
		}
	}

	fmt.Println(count)
}
