package arrays

import "fmt"

func Demo2() {
	var cities [5]string = [5]string{"Ankara", "İstanbul", "İzmir", "Adana", "Diyarbakır"}
	cities[4] = "Eskişehir"
	fmt.Println(cities)

	for _, v := range cities {
		fmt.Println(v)
	}

	for i, v := range cities {
		fmt.Println(i, v)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(cities[i])
	}
}
