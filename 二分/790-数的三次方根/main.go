package main

import "fmt"

const eps = 1.0e-8

var n float64

func main() {
	fmt.Scan(&n)
	var l, r float64
	//注意>0 <1 数的立方根 也一定是小数 还有负数同上的情况
	if n >= 0 && n <= 1 {
		l = 0
		r = 1
	} else if n >= 0 {
		l = 0
		r = n
	} else if n < 0 && n >= -1 {
		l = -1
		r = 0
	} else {
		l = n
		r = 0
	}
	for r-l > eps {
		mid := (l + r) / 2
		if mid*mid*mid < n {
			l = mid
		} else {
			r = mid
		}
	}
	fmt.Printf("%.6f", l)
}
