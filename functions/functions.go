package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}

func beSocialble(name string) {
	sayHello(name)
	fmt.Println("You are here!")
	sayGoodbye(name)
}

func addOne(x int) int {
	return x + 1
}

func sayHelloaBunch() {
	fmt.Println("Hello")
	sayHelloaBunch()
}

func main() {

	beSocialble("Benny")
	beSocialble("General Kenobi")

	x := 5
	x = addOne(x)
	fmt.Println(x)

	x = addOne(addOne(x))
	fmt.Println(x)

	//sayHelloaBunch()

}
