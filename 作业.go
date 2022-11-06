package main

import "fmt"

func main() {
	var (
		n, i, r, j, x, b, l, k, m, t, p int16
	)
	fmt.Scanln(&n, &i, &r)
	//var a []int16
	fmt.Println()
	fmt.Scanf("%d", &p)

	a := make([]int16, 1)
	a[0] = p
	for j = 1; j <= n-1; j++ {
		fmt.Scanf("%d", &x)
		a = append(a, x)
	}
	b = r - i
	m = i
	for l = 1; l <= b; l++ {
		for k = 1; k <= b; k++ {
			if a[i] > a[i+1] {
				t = a[i]
				a[i] = a[i+1]
				a[i+1] = t
			}
			i++
		}
		i = m
	}
	fmt.Println(a)
}

//https://zhuanlan.zhihu.com/p/193140870
