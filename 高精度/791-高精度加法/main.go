package main

import "fmt"

func add(A []int, B []int) []int {
	var C []int
	t := 0
	var la, lb = len(A), len(B)
	for i := 0; i < la || i < lb; i++ {
		if i < la {
			t += A[i]
		}
		if i < lb {
			t += B[i]
		}
		C = append(C, t%10)
		t /= 10
	}
	if t != 0 {
		C = append(C, 1)
	}
	return C
}
func main() {
	var a, b string
	var A, B []int
	fmt.Scan(&a, &b)
	var la, lb = len(a), len(b)
	for i := la - 1; i >= 0; i-- {
		A = append(A, int(a[i]-'0'))
	}
	for i := lb - 1; i >= 0; i-- {
		B = append(B, int(b[i]-'0'))
	}
	var C []int = add(A, B)
	for i := len(C) - 1; i >= 0; i-- {
		fmt.Printf("%d", C[i])
	}
}
