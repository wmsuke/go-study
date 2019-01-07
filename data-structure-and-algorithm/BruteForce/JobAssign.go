package IntPerm

import (
	"fmt"
	"math"
)

// リスト2-2

//matrix := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
constMatrix := [][]int{
	{20,13,8,10,17},
	{10,5,7,12,8}
}
var assign []int
var min int

func JobAssign(data []int) {
	mdata = data
	assign = make([]int, 5)
	min = math.MaxInt64
}

func eval() {
	cost := 0
	for i:=0; i<2; i++{
		cost += constMatrix[i][mdata[i]]
	}

}

func ShowAnswer() {
	fmt.Println("最小の組み合わせ :")
	for i := 0; i < 2; i++ {

	}
	fmt.Println(" = ")
}

// var mdata []int
//
// func New(data []int) {
// 	mdata = data
// }
//
// func Perm() {
// 	_perm(0)
// }
//
// func _perm(i int) {
// 	if i == len(mdata) {
// 		eval()
// 	} else {
// 		for j := i; j < len(mdata); j++ {
// 			swap(i, j)
// 			_perm(i + 1)
// 			swap(i, j)
// 		}
// 	}
// }
//
// func swap(i, j int) {
// 	temp := mdata[i]
// 	mdata[i] = mdata[j]
// 	mdata[j] = temp
// }
//
// func eval() {
// 	var str string
// 	for _, v := range mdata {
// 		str = fmt.Sprintf("%s %d", str, v)
// 	}
// 	fmt.Println(str)
// }
