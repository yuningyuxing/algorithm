package main

import "fmt"

var n, q int
var s []int

func main() {
	fmt.Scan(&n, &q)
	s = make([]int, n+10)
	for i := 1; i <= n; i++ {
		fmt.Scan(&s[i])
	}
	for q != 0 {
		var k int
		fmt.Scan(&k)
		l := 1
		r := n
		for l < r {
			mid := (l + r) >> 1
			if s[mid] >= k {
				r = mid
			} else {
				l = mid + 1
			}
		}
		if s[l] == k {
			fmt.Printf("%d ", l-1)
		} else {
			fmt.Printf("-1 ")
		}
		l = 1
		r = n
		for l < r {
			mid := (l + r + 1) >> 1
			if s[mid] <= k {
				l = mid
			} else {
				r = mid - 1
			}
		}
		if s[l] == k {
			fmt.Println(l - 1)
		} else {
			fmt.Println("-1")
		}
		q--
	}
}
