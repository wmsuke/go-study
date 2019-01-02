package main

import (
	"fmt"
	"strconv"
)

//リスト1-1
// 鶴亀算

//力ずく方による鶴の数の求め方
func tsuru1(n, m int) int {
	for i := 0; i <= n; i++ {
		f := 2*i + 4*(n-i)
		if f == m {
			return i
		}
	}
	return -1
}

// 連立方程式によるつるの数の求め方
func tsuru2(n, m int) int {
	return (4*n - m) / 2
}

func main() {
	var n int = 2
	var m int = 4
	fmt.Println("鶴の数は" + strconv.Itoa(tsuru1(n, m)) + "です")
	fmt.Println("鶴の数は" + strconv.Itoa(tsuru2(n, m)) + "です")
}
