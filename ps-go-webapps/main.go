package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	w.Write([]byte("oh hi"))
}

func searchInsert(nums []int, target int) int {
	start, mid, end := 0, 0, len(nums)-1
	for start <= end {
		mid = ((end - start) / 2) + start
		c := nums[mid]
		fmt.Println(mid, c, start, end)
		switch {
		case c == target:
			return mid
		case c < target:
			if mid == end {
				return mid + 1
			} else {
				start = mid + 1
			}
		case c > target:
			if mid == start {
				if mid == 0 {
					return 0
				} else {
					return mid - 1
				}
			} else {
				end = mid - 1
			}
		default:
			panic("oh noes")

		}
	}
	return mid
}

func main() {
	fmt.Println(searchInsert([]int{1, 3}, 2))
}
