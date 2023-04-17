// go版本
package main

import "fmt"

const N = 100005

var n int
var s [N]int

func quicksort(l int, r int) {
	if l >= r {
		return
	}
	i := l - 1
	j := r + 1
	// 注意go中的l+r>>1优先级和c++不同  我们要括号
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
	quicksort(l, j)
	quicksort(j+1, r)
}
func main() {
	// 注意我们尽量用Scanf输入
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &s[i])
	}
	quicksort(0, n-1)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", s[i])
	}
}
