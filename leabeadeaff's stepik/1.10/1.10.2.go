// https://stepik.org/submissions/720577/922330799?unit=200796
package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	count := 0
	for ; a <= b; a++ {
		count = count + a
	}
	fmt.Println(count)
}
