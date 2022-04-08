package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	t.Log(len(arr))
	t.Log(arr[1], arr[2])
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4}
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	//数组传统遍历
	arr := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}
	//goland遍历写法
	for i, i2 := range arr {
		t.Log(i, i2)
	}
	//goland不需要角标，用_占位
	for _, i := range arr {
		t.Log(i)
	}
}
func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	arr_sec := arr[3:len(arr)] //(不支持负数截取，也就是反顺序截取)
	t.Log(arr_sec)
}
