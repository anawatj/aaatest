package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func TestSolution41(t *testing.T) {
	var a []int = []int{3, 8, 9, 7, 6}
	var k int = 3
	ret := solution4(a, k)
	fmt.Println(ret)
	var result []int = []int{9, 7, 6, 3, 8}
	assert.True(t, testEq(ret, result))
}

func TestSolution42(t *testing.T) {
	var a []int = []int{0, 0, 0}
	var k int = 1
	ret := solution4(a, k)
	fmt.Println(ret)
	var result []int = []int{0, 0, 0}
	assert.True(t, testEq(ret, result))
}

func TestSolution43(t *testing.T) {
	var a []int = []int{1, 2, 3, 4}
	var k int = 4
	ret := solution4(a, k)
	var result []int = []int{1, 2, 3, 4}
	assert.True(t, testEq(ret, result))
}

func TestSolution4k101(t *testing.T) {
	var a []int
	for index := 1; index <= 101; index++ {
		a = append(a, index)
	}
	var k int = 4
	ret := solution4(a, k)
	var result []int = []int{}
	assert.True(t, testEq(ret, result))
}
