// https://stepik.org/submissions/775599/922310608?unit=205068
package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)
	if a[0]+a[1]+a[2] == a[3]+a[4]+a[5] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
