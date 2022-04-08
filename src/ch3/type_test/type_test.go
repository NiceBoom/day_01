package type_test

import (
	"math"
	"testing"
)

//golang不支持任何类型的隐式转换
//不支持别名到原类型的隐式转换
//go支持指针类型，但是不支持指针的运算
//go字符串是值类型，默认初始化为空字符串，不是nil
//go对于数据类型转换要求非常严苛
//转换前需要一个显式的转换声明
type MyInt int64

func TestName(t *testing.T) {
	var a int = 1
	var b int64
	b = int64(a)
	t.Log(a, b)
}

//类型的预定义值
func TestNumber(t *testing.T) {
	maxInt64 := math.MaxInt64
	maxFloat64 := math.MaxFloat64
	maxUint32 := math.MaxUint32
	t.Log(maxUint32, maxFloat64, maxInt64)
}

//指针
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)

	//%T 格式化符号，获得变量类型
	t.Logf("%T %T", a, aPtr)
	//输出内容为int，*int 意义为int类型与int类型指针
	//指针不支持运算。在其他语言容易用指针运算进行连续的空间获取
	//但是go不支持指针运算
}

//字符串
func TestString(t *testing.T) {
	//字符串在初始化的时候会赋值为空字符串，而不是空nil
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	//判断是否初始化
	if s == "" {

	}
}
