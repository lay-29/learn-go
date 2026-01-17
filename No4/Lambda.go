package No4

import "fmt"

func LambdaTest() {

	add := func(a int, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2))

	add2 := func(a int, b int) int {
		return a + b
	}(1, 10)

	fmt.Println(add2)
}
