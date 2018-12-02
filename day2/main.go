package main

import ("fmt"
	"io/ioutil"
	"os"
	"strings"
)

func convert_to_string(xs []uint8, end *int) (string) {
	start := *end
	for i:= start; xs[i] != 10; i++ {
		(*end)++
	}
	return string(xs[start:*end])
}

func make_string(xs []uint8) []string {
	s := make([]string, 0, 100)
	end := 0
	for end != len(xs) {
		s = append(s, convert_to_string(xs, &end))
		end++
	}
	return s
}

func find_two_three(s string) (two int, three int) {
	for i:=97; i < 123; i++ {
		count := strings.Count(s, string(i))
		if count == 2 {
			two = 1
		} else if count == 3 {
			three = 1
		}
		if two == 1 && three == 1 {
			return
		}
	}
	return
}

func part1(target []string) int {
	var two, three int
	for _, v := range target {
		temp1, temp2 := find_two_three(v)
		two += temp1
		three += temp2
	}
	answer := two * three
	return answer
}

func prototype_box(a string, b string) (bool, string) {
	diff := 0
	index := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if diff == 0 {
				diff = 1
				index = i
			} else {
				return false, ""
			}
		}
	}
	if diff == 1 {
		return true, strings.Replace(a, string(a[index]), "", -1)
	}
	return false, ""
}

func part2(target []string) string {
	for i, v := range target {
 		for j:=i; j < len(target); j++ {
			found, answer := prototype_box(v, target[j])
			if found {
				return answer
			}
 		}
	}
	return ""
}

func main() {
	dat, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Print(err)
	}

	target := make_string(dat)
	fmt.Println(part1(target))
	fmt.Println(part2(target))
	// Correct!!!
}
