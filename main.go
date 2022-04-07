package main

import "fmt"

func main() {
	// d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d := []int{84, 1, 95, 44, 2, 8, 101, 300, 50, 40, 20, 78}

	fmt.Print(quickSort(d))
}

// 遞迴作法
func quickSort(d []int) []int {
	fmt.Printf("data: %v\n", d)
	dataLen := len(d)
	// 結束遞迴的條件
	if len(d) <= 1 {
		return d
	}

	// 基準點
	pivot := d[0]

	// 小於基準點的資料
	small := make([]int, 0, dataLen)
	// 大於基準點的資料
	big := make([]int, 0, dataLen)
	// 等於基準點的資料
	equal := make([]int, 0, dataLen)

	// num := 0
	for _, v := range d {
		if v < pivot {
			small = append(small, v)
		} else if v > pivot {
			big = append(big, v)
		} else if v == pivot {
			equal = append(equal, v)
		}
		// num++
	}
	// fmt.Printf("num: %v\n", num)

	fmt.Printf("small: %v\n", small)
	fmt.Printf("big: %v\n", big)

	// 遞迴
	small = quickSort(small)
	fmt.Printf("small Recursion over\n")

	// 遞迴
	big = quickSort(big)
	fmt.Printf("big Recursion over\n")

	r := make([]int, 0, dataLen)
	r = append(r, small...)
	r = append(r, equal...)
	r = append(r, big...)

	fmt.Printf("r: %v\n", r)

	return r
}
