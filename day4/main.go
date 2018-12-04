package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"regexp"
	"strconv"
)

func atoi(s string) (answer int) {
	answer, _ = strconv.Atoi(s)
	return
}

func find_match(s string) int {
	wakes := regexp.MustCompile(`wakes`)
	sleep := regexp.MustCompile(`falls`)
	if wakes.MatchString(s) {
		return 2
	}
	if sleep.MatchString(s) {
		return 1
	}
	return 0
}

func time_stamp(s string) int {
	re := regexp.MustCompile(":(\\d+)")
	return atoi((re.FindStringSubmatch(s)[1]))
}

func get_guard(s string) int {
	re := regexp.MustCompile("#(\\d+)")
	return atoi(re.FindStringSubmatch(s)[1])
}

func sum_arr(xs []int) (sum int) {
	for _, v := range xs {
		sum += v
	}
	return
}

func find_max(xs []int) (answer int) {
	max := -1
	for i, v := range xs {
		if v > max {
			max = v
			answer = i
		}
	}
	return
}

func find_max_row(xs map[int][]int) (answer int) {
	sum := 0
	max := -1
	for i, v := range xs {
		sum = sum_arr(v)
		if sum > max {
			answer = i
			max = sum
		}
	}
	return
}

func main() {
	data, _ := ioutil.ReadFile(os.Args[1])
	s := strings.Split(string(data), "\n")

	shifts := make(map[int][]int)
	temp := 0
	on_duty := 0
	for _, v := range s {
		time := time_stamp(v)
		switch find_match(v) {
		case 0:
			on_duty = get_guard(v)
			_, ok := shifts[on_duty]
			if !ok {
				shifts[on_duty] = make([]int, 60)
			}
		case 1:
			temp = time
		case 2:
			for i := temp; i < time; i++ {
				shifts[on_duty][i]++
			}
		}
	}

	// Part 1
	fmt.Println(find_max(shifts[find_max_row(shifts)]) * find_max_row(shifts)) // Correct! 85296

	// Part 2
	arr := make(map[int]int)
	for i, v := range shifts {
		arr[i] = find_max(v)
	}

	max := 0
	index := 0
	for i, v := range arr {
		if shifts[i][v] > max {
			max = shifts[i][v]
			index = i
		}
	}
	fmt.Println(index * arr[index]) // Correct! 58559
}
