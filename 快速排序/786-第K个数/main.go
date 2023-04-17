package main

// go版本
import "fmt"

// 递归左右
func quicksort(s []int, l int, r int, k int) int {
	if l >= r {
		return s[l]
	}
	i := l - 1
	j := r + 1
	x := s[(l+r)>>1]
	for i < j {
		for {
			i++
			if s[i] >= x {
				break
			}
		}
		for {
			j--
			if s[j] <= x {
				break
			}
		}
		if i < j {
			s[i], s[j] = s[j], s[i]
		}
	}
	if j-l+1 >= k {
		return quicksort(s, l, j, k)
	} else {
		return quicksort(s, j+1, r, k-(j-l+1))
	}
	return 0
}
func main() {
	var n int
	var k int
	fmt.Scan(&n, &k)
	s := make([]int, n+10)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	t := quicksort(s, 0, n-1, k)
	fmt.Println(t)
}
