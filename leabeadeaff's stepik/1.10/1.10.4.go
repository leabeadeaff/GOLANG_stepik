// https://stepik.org/submissions/729482/923465300?unit=200796
package main

import (
	"fmt"
)

func main() {
	var a, c, maxInt int
	for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
		if a > maxInt {
			maxInt = a
			c = 0
		}
		if maxInt == a {
			c++
		}
	}
	fmt.Println(c)
}
