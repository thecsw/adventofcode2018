package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
	"strconv"
)

type elf_case struct {
	number int
	left int
	top int
	width int
	height int
}

func get_object(s string) (obj elf_case) {
	r, _ := regexp.Compile("#(\\d+)\\s@\\s(\\d+),(\\d+):\\s(\\d+)x(\\d+)")
	temp1, _ := strconv.Atoi(r.FindStringSubmatch(s)[1])
	temp2, _ := strconv.Atoi(r.FindStringSubmatch(s)[2])
	temp3, _ := strconv.Atoi(r.FindStringSubmatch(s)[3])
	temp4, _ := strconv.Atoi(r.FindStringSubmatch(s)[4])
	temp5, _ := strconv.Atoi(r.FindStringSubmatch(s)[5])
	obj = elf_case{number: temp1,
		left: temp2,
		top: temp3,
		width: temp4,
		height: temp5,
	}
	return
}

func max_dimensions(xs []elf_case) (width int, height int) {
	for _, v := range xs {
		if v.left + v.width > width {
			width = v.left + v.width
		}
		if v.top + v.height > height {
			height = v.top + v.height
		}
	}
	return
}

func find_case(board [][]int, cases []elf_case) int {
	case_number := -1
	for _, v := range cases {
		case_number = v.number
		hits := 0
		for i:= v.left; i < v.left + v.width; i++ {
			for j:= v.top; j < v.top + v.height; j++ {
				if board[i][j] != 1 {
					hits++
				}
			}
		}
		if hits == 0 {
			return case_number
		}
	}
	return -1
}

func main() {
	data, _ := ioutil.ReadFile(os.Args[1])
	s := strings.Split(string(data), "\n")

	arr := make([]elf_case, 0, 100)
	for _, v := range s {
		arr = append(arr, get_object(v))
	}

	y, x := max_dimensions(arr)

	board := make([][]int, y)
	for i, _ := range board {
		board[i] = make([]int, x)
	}

	for _, v := range arr {
		for i:= v.left; i < v.left + v.width; i++ {
			for j:= v.top; j < v.top + v.height; j++ {
				switch board[i][j] {
				case 0:
					board[i][j] = 1
				case 1:
					board[i][j] = 2
				}
			}
		} 
	}

	answer := 0
	for _, column := range board {
		for _, rownum := range column {
			if rownum == 2 {
				answer++
			}
		}
	}

	fmt.Println("Part one:", answer)
	fmt.Println("Part two:", find_case(board, arr))
}
