package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	Val  string
	Next *Stack
}

func (s *Stack) Pop() (*Stack, *Stack) {
	return s, s.Next
}

func (s *Stack) Print() {
	str := ""

	for n := s; n != nil; n = n.Next {
		str = str + " " + n.Val
	}

	fmt.Println(str)
}

func (s *Stack) AddLast(e *Stack) {
	cur := s
	for cur = s; cur.Next != nil; cur = cur.Next {
	}
	cur.Next = e
}

func swap(amt int, from, to *Stack) (*Stack, *Stack) { //Part 1
	for i := 0; i < amt; i++ {
		from, from.Next, to = from.Next, to, from
	}

	return from, to
}

func swapN(amt int, from, to *Stack) (*Stack, *Stack) { //Part 2
	last := from
	for i := 0; i < amt-1; i++ {
		last = last.Next
	}

	to, from, last.Next = from, last.Next, to

	return from, to
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()

	line := sc.Text()
	nStacks := (len(line) + 1) / 4

	stacks := make([]*Stack, nStacks)

	for i := 0; i < nStacks; i++ {
		crateVal := string(line[4*i+1])
		if crateVal != " " {
			tail := Stack{
				Val:  crateVal,
				Next: nil,
			}

			if stacks[i] != nil {
				stacks[i].AddLast(&tail)
			} else {
				stacks[i] = &tail
			}
		}
	}

	keepGoing := true
	for keepGoing {
		sc.Scan()
		line := sc.Text()
		for i := 0; i < nStacks; i++ {
			crateVal := string(line[4*i+1])
			_, err := strconv.Atoi(crateVal)

			if err == nil {
				keepGoing = false
				break
			}

			if crateVal != " " {
				tail := &Stack{
					Val:  crateVal,
					Next: nil,
				}

				if stacks[i] != nil {
					stacks[i].AddLast(tail)
				} else {
					stacks[i] = tail
				}
			}
		}
	}

	sc.Scan()

	for sc.Scan() {
		instruction := strings.Split(sc.Text(), " ")
		amtS, fromS, toS := instruction[1], instruction[3], instruction[5]

		amt, _ := strconv.Atoi(amtS)
		from, _ := strconv.Atoi(fromS)
		to, _ := strconv.Atoi(toS)

		stacks[from-1], stacks[to-1] = swapN(amt, stacks[from-1], stacks[to-1])
	}

	res := ""
	for i := range stacks {
		res = res + stacks[i].Val
	}

	fmt.Println(res)
}
