package arrays

import "fmt"

func Demo3() {
	numbers := [5]int{20, 30, 50, 10, 2}
	fmt.Println(numbers)

	biggest := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > biggest {
			biggest = numbers[i]
		}
	}

	fmt.Printf("Biggest number is: %d", biggest)
}
