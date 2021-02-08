package main

import "fmt"

func addOne(x *int) {
	*x = *x + 1
}

func main() {

	x := 5
	xPtr := &x
	fmt.Println(x)
	fmt.Println(xPtr)

	addOne(xPtr)
	fmt.Println(x)

}
