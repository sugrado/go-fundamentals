package slices

import "fmt"

func Demo1() {
	var cities []string = make([]string, 2)
	cities[0] = "Ankara"
	cities[1] = "İstanbul"
	cities = append(cities, "Adana")
	fmt.Println(cities)
}
