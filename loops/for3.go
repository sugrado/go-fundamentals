package loops

import "fmt"

func PrimeNumber() {
	var inputNumber int

	fmt.Scanln(&inputNumber)
	if inputNumber < 2 {
		fmt.Printf("%d : Asal Değildir", inputNumber)
		return
	}

	for i := 2; i <= inputNumber/2; i++ {
		if inputNumber%i == 0 {
			fmt.Printf("%d : Asal Değildir", inputNumber)
			return
		}
	}
	fmt.Printf("%d : Asaldır", inputNumber)
}
