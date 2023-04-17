package main

import "fmt"

var q []int
var sum int

func mergesort(s []int, l int, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	mergesort(s, l, mid)
	mergesort(s, mid+1, r)
	i := l
	j := mid + 1
	k := 0
	for i <= mid && j <= r {
		if s[i] <= s[j] {
			q[k] = s[i]
			k++
			i++
		} else {
			q[k] = s[j]
			k++
			j++
			//对两个排好序进行合并
			sum += mid - i + 1
		}
	}
	for i <= mid {
		q[k] = s[i]
		k++
		i++
	}
	for j <= r {
		q[k] = s[j]
		k++
		j++
	}
	i = 0
	j = l
	for j <= r {
		s[j] = q[i]
		i++
		j++
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n+10)
	q = make([]int, n+10)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	mergesort(s, 0, n-1)
	fmt.Println(sum)
}
