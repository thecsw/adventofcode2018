package main

import ("fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func convert_number(xs []uint8, start int) int {
	end := start
	for i := start; xs[i] != 10; i++ {
		end++
	}
	number, _ := strconv.Atoi(string(xs[start:end]))
	return number
}

func find(xs []int, target int) bool {
	for _, v := range xs {
		if v == target {
			return true
		}
	}
	return false
}

func calc_freqs(xs []uint8, sums *[]int, start int) (sum int, found bool) {
	sum = start
	for i, v := range xs {
		mul := 1;
		if v == 45  || v == 43 {
			if v == 45 {
				mul *= -1
			}
			result := mul * convert_number(xs, i + 1)
			sum += result
			if found = find(*sums, sum); found {
				return sum, found
			}
			*sums = append(*sums, sum)
		}
	}
	return sum, false
}

func main() {
	dat, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Print(err)
	}

	sums := make([]int, 0, 100)
	target := 0
	found := false
	for ; !found; {
		target, found = calc_freqs(dat, &sums, target)
	}
	
	fmt.Println(target)
}
