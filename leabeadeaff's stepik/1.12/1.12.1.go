// https://stepik.org/submissions/1257942/924791398?unit=200798
package main

import (
	"fmt"
)

func main() {
	var workArray [10]uint8
	var idxArray [6]uint8
	var a uint8
	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&a)
		workArray[i] = a
	}
	for k := 0; k < len(idxArray); k++ {
		fmt.Scan(&a)
		idxArray[k] = a
	}

	workArray[idxArray[0]], workArray[idxArray[1]] = workArray[idxArray[1]], workArray[idxArray[0]]
	workArray[idxArray[2]], workArray[idxArray[3]] = workArray[idxArray[3]], workArray[idxArray[2]]
	workArray[idxArray[4]], workArray[idxArray[5]] = workArray[idxArray[5]], workArray[idxArray[4]]

	for _, value := range workArray {
		fmt.Printf("%v ", value)
	}
}
