package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeView struct {
	L, R, U, D int
}

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

	viewD := make([][]TreeView, len(m))
	for i := range viewD {
		viewD[i] = make([]TreeView, len(m[0]))
	}

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

			if j < len(m[0])-1 {
				curT := m[i][j][0]
				v := 1
				for nextT := j + v; nextT < len(m[0])-1; nextT = j + v {
					if m[i][nextT][0] < curT {
						v += viewD[i][nextT].R
					} else {
						break
					}
				}
				viewD[i][j].R = v
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

			if j > 0 {
				curT := m[i][j][0]
				v := 1
				for nextT := j - v; nextT > 0; nextT = j - v {
					if m[i][nextT][0] < curT {
						v += viewD[i][nextT].L
					} else {
						break
					}
				}
				viewD[i][j].L = v
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

			if j > 0 {
				curT := m[j][i][0]
				v := 1
				for nextT := j - v; nextT > 0; nextT = j - v {
					if m[nextT][i][0] < curT {
						v += viewD[nextT][i].U
					} else {
						break
					}
				}
				viewD[j][i].U = v
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

			if j < len(m[0])-1 {
				curT := m[j][i][0]
				v := 1
				for nextT := j + v; nextT < len(m[0])-1; nextT = j + v {
					if m[nextT][i][0] < curT {
						v += viewD[nextT][i].D
					} else {
						break
					}
				}
				viewD[j][i].D = v
			}
		}
	}

	maxView := 0
	for i := range viewD {
		for j := range viewD[i] {
			curView := viewD[i][j]
			val := curView.D * curView.R * curView.U * curView.L
			if val > maxView {
				maxView = val
			}
		}
	}
	fmt.Println("PART 1: ", count)
	fmt.Println("PART 2: ", maxView)
}
