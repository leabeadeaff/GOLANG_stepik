// https://stepik.org/submissions/729495/922288922?unit=205068
package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a%10 != (a/10)%10 && (a/10)%10 != ((a/100)%10) && a%10 != ((a/100)%10) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
