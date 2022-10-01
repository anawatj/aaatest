package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution2Success(t *testing.T) {
	var input string
	input = "We test coders. Give us a try?"

	ret := Solution2(input)
	assert.True(t, ret == 4)

}

func TestSolution2WithEmpty(t *testing.T) {
	var input string
	input = ""
	ret := Solution2(input)
	assert.True(t, ret == 0)
}

func TestSolution2With101String(t *testing.T) {
	var input string
	input = "We test coders. Give us a try?We test coders. Give us a try?We test coders. Give us a try?We test coders. Give us a try?We test coders. Give us a try?We test coders. Give us a try?We test coders. Give us a try?"
	ret := Solution2(input)
	assert.True(t, ret == 0)
}

func TestSolutionWithNumber(t *testing.T) {
	var input string
	input = "Hello Tao 1A"
	ret := Solution2(input)
	assert.True(t, ret == 2)
}
