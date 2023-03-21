package interfaces

import "fmt"

func Demo2() {
	// type assertion (type casting)
	var number1 interface{} = "Gorkem"
	validate(number1)

	var number2 interface{} = 5
	validate(number2)
}

func validate(i interface{}) {
	val, ok := i.(int)
	fmt.Println(val, ok)
}
