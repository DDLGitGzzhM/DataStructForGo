package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_swap(t *testing.T) {
	// 测试整数交换
	x := 1
	y := 2
	Swap(&x, &y)
	assert.Equal(t, 2, x)
	assert.Equal(t, 1, y)

	// 测试字符串交换
	str1 := "Hello"
	str2 := "World"
	Swap(&str1, &str2)
	assert.Equal(t, "World", str1)
	assert.Equal(t, "Hello", str2)

	// 测试切片交换
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	Swap(&slice1, &slice2)

	assert.Equal(t, []int{4, 5, 6}, slice1)
	assert.Equal(t, []int{1, 2, 3}, slice2)

	// 测试映射交换
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"c": 3, "d": 4}
	Swap(&map1, &map2)

	assert.Equal(t, map[string]int{"c": 3, "d": 4}, map1)
	assert.Equal(t, map[string]int{"a": 1, "b": 2}, map2)
}
