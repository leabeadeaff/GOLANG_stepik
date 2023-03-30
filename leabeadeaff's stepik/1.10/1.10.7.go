// https://stepik.org/submissions/775725/924282628?unit=200796
package main

import (
	"fmt"
)

func main() {
	var x, p, y, c int
	fmt.Scan(&x, &p, &y)
	for ; x < y; c++ {
		x += x * p / 100
	}
	fmt.Println(c)
}
