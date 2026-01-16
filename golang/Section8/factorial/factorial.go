package factorial

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}

func factorial(n int) int {
	// result := 1
	// for i := 1; i <= n; i++ {
	// 	result *= i
	// }
	// return result
	if n == 1 || n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
