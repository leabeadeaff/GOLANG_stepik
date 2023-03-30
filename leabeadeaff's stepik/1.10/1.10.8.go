// https://stepik.org/submissions/776050/924356391?unit=200796
package main

import (
	"fmt"
)

func main() {
	var x, y, c int
	fmt.Scan(&x, &y)
	c = y
	reversedX := 0
	for x > 0 {
		remainder := x % 10
		reversedX = (reversedX * 10) + remainder
		x /= 10
	}

	for r := 0; reversedX != 0; reversedX /= 10 {
		r = reversedX % 10
		for i := 0; y != 0; y /= 10 {
			i = y % 10
			if r == i {
				fmt.Print(r, " ")
				break
			}
		}
		y = c
	}
}
