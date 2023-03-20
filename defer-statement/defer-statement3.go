package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("Product saved: ", p.productName)
	defer Log()
}

func Log() {
	fmt.Println("Logged")
}

func Demo3() {
	var p = product{productName: "Laptop", unitPrice: 5000}
	defer p.save()
	p = product{productName: "Mouse", unitPrice: 45}

	fmt.Println("Success!")
	fmt.Println(p.productName)
	/* PRINTS:
	Success!
	Mouse
	Product saved:  Laptop
	Logged
	*/
}
