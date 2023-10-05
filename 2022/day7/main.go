package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IFile interface {
	GetName() string
	GetSize() int
	ToString() string
}

type Directory struct {
	Name     string
	Children []IFile
	Parent   *Directory
}

func (d *Directory) GetName() string {
	return d.Name
}

func (d *Directory) GetSize() int {
	size := 0
	for i := range d.Children {
		size += (d.Children[i]).GetSize()
	}
	return size
}

func (d *Directory) ToString() string {
	return "dir " + d.Name
}

func (d *Directory) HasDir(name string) bool {
	if len(d.Children) == 0 {
		return false
	}

	for i := range d.Children {
		switch v := (d.Children[i]).(type) {
		case *Directory:
			if v.Name == name {
				return true
			}
		default:
		}
	}

	return false
}

func (d *Directory) HasFile(name string) bool {
	for i := range d.Children {
		switch v := (d.Children[i]).(type) {
		case *File:
			if v.Name == name {
				return true
			}
		default:
		}
	}

	return false
}

type File struct {
	Name   string
	Size   int
	Parent *Directory
}

func (f *File) GetName() string {
	return f.Name
}

func (f *File) GetSize() int {
	return f.Size
}

func (f *File) ToString() string {
	return fmt.Sprintf(`%d %s`, f.Size, f.Name)
}

type Command interface {
	Execute(cmd string, sc *bufio.Scanner, dir *Directory) (*Directory, string)
	Match(cmd string) bool
}

type ChangeDir struct{}

func (*ChangeDir) Match(s string) bool {
	cmd := strings.Split(s, " ")[1]
	return cmd == "cd"
}

func (*ChangeDir) Execute(cmd string, sc *bufio.Scanner, dir *Directory) (*Directory, string) {
	target := strings.Split(cmd, " ")[2]
	if target == ".." {
		if dir.Parent == nil {
			return dir, ""
		}
		return dir.Parent, ""
	}

	if len(dir.Children) == 0 {
		return dir, ""
	}

	for i := range dir.Children {
		switch v := (dir.Children[i]).(type) {
		case *File:
		case *Directory:
			if v.Name == target {
				return v, ""
			}
		}
	}

	return nil, ""
}

type ListDir struct{}

func (*ListDir) Match(s string) bool {
	//fmt.Println("MATCHING: ", s)
	cmd := strings.Split(s, " ")[1]
	return cmd == "ls"
}

func (*ListDir) Execute(cmd string, sc *bufio.Scanner, dir *Directory) (*Directory, string) {
	//fmt.Println("LIST DIR ", cmd, " FOR DIR ", dir)

	for sc.Scan() {
		line := sc.Text()
		if line[0] == '$' {
			break
		}

		parts := strings.Split(line, " ")

		if parts[0] == "dir" && !dir.HasDir(parts[0]) {
			newDir := Directory{
				Name:     parts[1],
				Children: []IFile{},
				Parent:   dir,
			}

			dir.Children = append(dir.Children, &newDir)
		} else if !dir.HasFile(parts[1]) {
			size, _ := strconv.Atoi(parts[0])
			file := File{
				Name:   parts[1],
				Size:   size,
				Parent: dir,
			}
			//fmt.Println("APPENDING ", file, " TO ", dir)
			dir.Children = append(dir.Children, &file)
		}
	}
	return dir, sc.Text()
}

func Solve(d *Directory) (int, []int) {
	size := 0
	sol := []int{}

	for i := range d.Children {
		switch f := d.Children[i].(type) {
		case *File:
			size += f.GetSize()
		case *Directory:
			childSize, rec := Solve(f)
			size += childSize
			sol = append(sol, rec...)
		}
	}

	if size < 100000 {
		sol = append(sol, size)
	}

	return size, sol
}

type DirSize struct {
	Name string
	Size int
}

func ResumeDirs(d *Directory) (int, []*DirSize) {
	size := 0
	sol := []*DirSize{}

	for i := range d.Children {
		switch f := d.Children[i].(type) {
		case *File:
			size += f.GetSize()
		case *Directory:
			childSize, rec := ResumeDirs(f)
			size += childSize
			sol = append(sol, rec...)
		}
	}

	sol = append(sol, &DirSize{
		Name: d.Name,
		Size: size,
	})

	return size, sol
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	rootDir := Directory{
		Name:     "/",
		Children: []IFile{},
		Parent:   nil,
	}
	cmds := []Command{
		&ChangeDir{},
		&ListDir{},
	}

	curDir := &rootDir
	sc.Scan()
	line := sc.Text()
	var test string = ""
	for {
		for i := range cmds {
			if cmds[i].Match(line) {
				curDir, test = cmds[i].Execute(line, sc, curDir)
			}
		}
		line = test

		if line == "" {
			sc.Scan()
			line = sc.Text()
		}

		if len(line) == 0 || line[0] != '$' {
			break
		}
	}

	_, solutions := Solve(&rootDir)

	solution := 0
	for i := range solutions {
		solution += solutions[i]
	}

	fmt.Println("PART 1: ", solution)

	totalSize, solution2 := ResumeDirs(&rootDir)
	remainingSpace := 70000000 - totalSize
	spaceToFree := 30000000 - remainingSpace
	solution = 70000000

	for i := range solution2 {
		candidate := solution2[i]
		if (candidate.Size >= spaceToFree) && candidate.Size < solution {
			solution = candidate.Size
		}
	}

	fmt.Println("PART 2: ", solution)
}
