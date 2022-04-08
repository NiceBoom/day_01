package _map

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	t.Logf("len m1=%d", len(m1))
	t.Log(len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))
	t.Log(m2)
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[1] = 0
	t.Log(m1[2])
	//当访问的key不存在，仍然返回0，不能通过nil来判断元素是否存在
	if v, ok := m1[3]; ok {
		t.Logf("key 3 is %d.", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestTraveMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	for i, i2 := range m1 {
		t.Log(i, i2)
	}
}
