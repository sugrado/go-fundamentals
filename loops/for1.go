package loops

import "fmt"

func Demo1() {
	i := 0
	for i <= 5 {
		fmt.Print(i)
		i++
	}
	fmt.Print("\n")

	j := 0
	for ; j <= 5; j++ {
		fmt.Print(j)
	}
	fmt.Print("\n")

	var k int = 0
	for {
		if k <= 5 {
			fmt.Print(k)
		} else {
			break
		}
		k++
	}
}
