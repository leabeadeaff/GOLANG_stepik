// https://stepik.org/submissions/775601/922313836?unit=205068
package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	if a%400 == 0 || (a%4 == 0 && a%100 != 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
