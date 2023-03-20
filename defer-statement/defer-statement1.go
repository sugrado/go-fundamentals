package defer_statement

import "fmt"

func A() {
	fmt.Println("func A running")
}

func B() {
	defer A()
	defer C()
	fmt.Println("func B running")
}

func C() {
	fmt.Println("func C running")
}
