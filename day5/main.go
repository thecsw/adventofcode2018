package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func make_run(xs []uint8) ([]uint8, bool) {
	for i := 0; i < len(xs) - 1; i++ {
		if (xs[i] < 97 && xs[i + 1] == xs[i] + 32) || (xs[i] > 96 && xs[i] - 32 == xs[i + 1]) {
			xs = append(xs[:i], xs[i+2:]...)
			return xs, true
		}
	}
	return xs, false
}

func part1(xs []uint8) int {
	arr := xs
	cont := true
	for cont {
		arr, cont = make_run(arr)
	}
	return len(arr)
}

func delete_element(xs []uint8, letter1 uint8, letter2 uint8) []uint8 {
	arr := make([]uint8, 0)
	for _, v := range xs {
		if v != letter1 && v != letter2 {
			arr = append(arr, v)
		}
	}
	return arr
}

func part2(xs []uint8) int {
	answers := make([]int, 0, 100)
	for i:= uint8(65); i < 91; i++ {
		answers = append(answers, part1(delete_element(xs, i, i + 32)))
	}

	min := answers[0]
	for _, v := range answers {
		if v < min {
			min = v
		} 
	}
	return min
}

func main() {
	data, _ := ioutil.ReadFile(os.Args[1])
	//	fmt.Println(part1(data))
	fmt.Println(part2(data))
}
