package conditionals

import "fmt"

func Demo1() {
	num1, num2, num3 := 1, 2, 3

	biggest := num1
	if num2 > biggest {
		biggest = num2
	} else if num3 > biggest {
		biggest = num3
	}

	fmt.Print("Biggest number is: " + fmt.Sprintf("%d", biggest))
}
