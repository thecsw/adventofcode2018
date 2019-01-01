package main

import ("os"
	"io/ioutil"
	"fmt"
	"log"
	"regexp"
	"bytes"
	"strconv"
)

func parse_input(input []byte) (y []int, x []int) {
	arr := bytes.Split(input, []byte("\n"))
	re := regexp.MustCompile(`(\d+),\s+(\d+)`)
	y = make([]int, len(arr))
	x = make([]int, len(arr))
	for i, v := range arr {
		res := re.FindSubmatch(v)
		y[i], _ = strconv.Atoi(string(res[1]))
		x[i], _ = strconv.Atoi(string(res[2]))
	}
	return
}

func find_max(arr []int) (max int){
	max = 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return
}

// Don't want to import huge math package
func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func taxgeo(p1, p2, q1, q2 int) int {
	return abs(p1 - q1) + abs(p2 - q2)
}

func find(arr []int, target int) int {
	counter := 0
	for _, v := range arr {
		if v == target {
			counter++
		}
	}
	return counter
}

func find_minimum(arr []int) int {
	min := 999999
	ans := 0
	for i, v := range arr {
		if v < min {
			min = v
			ans = i
		}
	}
	return ans
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	y, x := parse_input(data)
	max_y, max_x := find_max(y), find_max(x)
	coords := len(y)
	ans := make([]int, coords)
	for i := 0; i <= max_y; i++ {
		for j := 0; j <= max_x; j++ {
			temp := make([]int, 0)
			// Calculate all distances
			for k := 0; k < coords; k++ {
				temp = append(temp, taxgeo(i, j, y[k], x[k]))
			}
			min_dist_index := find_minimum(temp)
			if find(temp, temp[min_dist_index]) <= 1 {
				ans[min_dist_index]++
			}
		}
	}
	fmt.Println(ans)
}
