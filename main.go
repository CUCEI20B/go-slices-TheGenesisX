package main

import "fmt"

func main() {
	var n, i uint64 = 0, 0
	var value, suma int64 = 0, 0
	var s []int64

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&value)
		s = append(s, value)
		suma += value
	}

	fmt.Println(suma)
}
