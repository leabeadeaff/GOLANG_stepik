// https://stepik.org/submissions/714699/923632118?unit=200796
package main

import (
	"fmt"
)

func main() {
	var a int
	for fmt.Scan(&a); a <= 100; fmt.Scan(&a) {
		if a < 10 {
			continue
		}
		fmt.Println(a)
	}
}
