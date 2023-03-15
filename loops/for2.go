package loops

import "fmt"

func Demo2() {
	actualNumber := 80
	var estimatedNumber int

	for {
		fmt.Scanln(&estimatedNumber)

		if estimatedNumber > actualNumber {
			fmt.Print("Biraz aşşa")
		} else if estimatedNumber < actualNumber {
			fmt.Print("Biraz yukarı")
		} else {
			fmt.Print("Bildin!")
			break
		}
	}
}
