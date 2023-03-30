// https://stepik.org/submissions/775597/922303166?unit=205068
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a)
	b = a
	count := 0
	for a > 0 {
		a = a / 10
		count++
	}
	fmt.Println((b / int(math.Pow(10, float64(count-1)))) % 10)
}
