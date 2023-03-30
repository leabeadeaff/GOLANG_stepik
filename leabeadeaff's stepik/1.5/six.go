// https://stepik.org/submissions/776062/922144032?unit=200794
package main

import (
	"fmt"
)

func main() {

	var a int
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	// 1 deg - 86400/720 sec => 120
	// 1 hour - 3600 sec
	// 1 min - 60 sec
	fmt.Println("It is", a*120/3600, "hours", (a*120/60)%60, "minutes.")
}
