package defer_statement

import "fmt"

func Demo1(number int) string {
	defer DeferFunc()
	if number%2 == 0 {
		return "even number"
	}
	if number%2 != 0 {
		return "odd number"
	}
	return "Wrong input!"
}

func Test() {
	result := Demo1(10)
	fmt.Println(result)
}

func DeferFunc() {
	fmt.Println("Finish")
}
