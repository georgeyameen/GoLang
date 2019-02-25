package main

import "fmt"

func main() {
	fmt.Println("hello everyone")
	foo()
	fmt.Println("something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm in foo")
}

//control flow"
// 1 sequence
// 2 loop
// 3 conditional
