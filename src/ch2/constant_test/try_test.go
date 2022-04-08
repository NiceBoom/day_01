package constant_test

import "testing"

//可以做递增递减
const (
	monday = iota + 1
	tuesday
	wednesday
)

//可以做状态位标识
const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(monday, tuesday, wednesday)
}

func TestConstantTry1(t *testing.T) {
	a := 1
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
