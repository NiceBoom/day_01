package test

import (
	"fmt"
	"testing"
)

func main(t *testing.T) {
	var a int = 1
	var b int = 1
	/**

	var(

		)
	*/
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	//变量互换
	//tmp := a
	//a = b
	//b = tmp
	a, b = b, a
	t.Log(a, b)

}
