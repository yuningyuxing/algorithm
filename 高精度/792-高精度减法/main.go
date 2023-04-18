package main

//切片模拟减法运算 注意负数情况 以及保证大减小即可
import "fmt"

func sub(A []int, B []int) []int {
	var C []int
	t := 0
	la := len(A)
	lb := len(B)
	for i := 0; i < la; i++ {
		t = A[i] - t
		if i < lb {
			t -= B[i]
		}
		C = append(C, (t+10)%10)
		if t < 0 {
			t = 1
		} else {
			t = 0
		}
	}
	for {
		lc := len(C)
		if lc > 1 && C[lc-1] == 0 {
			C = C[:lc-1]
		} else {
			break
		}
	}
	return C
}
func main() {
	var a, b string
	var A, B, C []int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	la := len(a)
	lb := len(b)
	for i := la - 1; i >= 0; i-- {
		A = append(A, int(a[i]-'0'))
	}
	for i := lb - 1; i >= 0; i-- {
		B = append(B, int(b[i]-'0'))
	}
	if la == lb {
		falg := false
		for i := la - 1; i >= 0; i-- {
			if A[i] > B[i] {
				C = sub(A, B)
				for i := len(C) - 1; i >= 0; i-- {
					fmt.Printf("%d", C[i])
				}
				falg = true
				break
			} else if A[i] < B[i] {
				C = sub(B, A)
				fmt.Printf("-")
				for i := len(C) - 1; i >= 0; i-- {
					fmt.Printf("%d", C[i])
				}
				falg = true
				break
			}
		}
		if !falg {
			C = sub(A, B)
			for i := len(C) - 1; i >= 0; i-- {
				fmt.Printf("%d", C[i])
			}
		}
	} else if la > lb {
		C = sub(A, B)
		for i := len(C) - 1; i >= 0; i-- {
			fmt.Printf("%d", C[i])
		}
	} else {
		C = sub(B, A)
		fmt.Printf("-")
		for i := len(C) - 1; i >= 0; i-- {
			fmt.Printf("%d", C[i])
		}
	}
}
