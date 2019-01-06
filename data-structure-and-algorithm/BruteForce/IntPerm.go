// package main
package IntPerm

import "fmt"

var mdata []int

func New(data []int) {
	mdata = data
}

func Perm() {
	_perm(0)
}

func _perm(i int) {
	if i == len(mdata) {
		eval()
	} else {
		for j := i; j < len(mdata); j++ {
			swap(i, j)
			_perm(i + 1)
			swap(i, j)
		}
	}
}

func swap(i, j int) {
	temp := mdata[i]
	mdata[i] = mdata[j]
	mdata[j] = temp
}

func eval() {
	var str string
	for _, v := range mdata {
		str = fmt.Sprintf("%s %d", str, v)
	}
	fmt.Println(str)
}
