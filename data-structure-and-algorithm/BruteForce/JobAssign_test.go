package IntPerm

import "testing"

func TestJobAssign(t *testing.T) {
	_data := []int{1, 2, 3}
	New(_data)
	JobAssign(_data)
}

func TestShowAnswer(t *testing.T) {
	_data := []int{1, 2, 3}
	// New(_data)
	// Perm()
	New(_data)
	ShowAnswer()
}
