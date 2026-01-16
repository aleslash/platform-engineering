package main

import "fmt"

func main() {
	fmt.Println(sumup(1, 2, 3, 4, 5))

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(sumup(numbers...))

}

func sumup(n ...int) int {
	result := 0
	for _, v := range n {
		result += v
	}
	return result
}
