// https://stepik.org/submissions/720578/923365379?unit=200796
package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		if b > 9 && b < 100 && b%8 == 0 {
			c = c + b
		}
	}
	fmt.Println(c)
}
