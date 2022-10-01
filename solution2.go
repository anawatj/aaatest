package main

import (
	"regexp"
	"strings"
)

func RegSplit(text string, delimeter string) []string {
	reg := regexp.MustCompile(delimeter)
	indexes := reg.FindAllStringIndex(text, -1)
	laststart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = text[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = text[laststart:len(text)]
	return result
}
func Solution2(s string) int {
	if len(s) == 0 || len(s) > 100 {
		return 0
	}
	str := RegSplit(s, "[$&+,:;=?@#|'<>.-^*()%!0-9]+")
	var max int
	for index, item := range str {
		if index == 0 {

			max = len(strings.Split(item, " "))

		} else {
			current := len(strings.Split(item, " "))
			if current > max {
				max = current
			}
		}
	}
	return max

}
func main() {
	print(Solution2("Hello Tao 1A"))
}
