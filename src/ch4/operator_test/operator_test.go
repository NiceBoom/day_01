package operator_test

import "testing"

//按位清零操作符
//数组比较，当维数不同则编译错误，维数与元素以及元素的顺序全相同才会输出true
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	e := [...]int{1, 3, 2, 4}
	t.Log(a == b)
	//t.Log(a==c)
	t.Log(a == d)
	t.Log(a == e)
}

//位运算符&^ 按位置零
func TestOperator1(t *testing.T) {
	//右方操作数为1，无论左方操作数是0还是1，结果都是0 （二进制位）
	//右方操作数为0，左方操作数是什么，结果就是什么（二进制位）
	//可以用二进制标志位标识状态，用状态定义常量
	a := 69
	b := 23
	c := a &^ b
	d := a &^ a
	e := b &^ a
	f := b &^ b
	t.Log(c, d, e, f)
}
