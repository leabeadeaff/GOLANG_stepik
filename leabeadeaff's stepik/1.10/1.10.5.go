// https://stepik.org/submissions/716358/923625548?unit=200796
package main

import (
	"fmt"
)

func main() {
	var n, c, d int
	fmt.Scan(&n)
	fmt.Scan(&c)
	fmt.Scan(&d)
	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
	}
}
