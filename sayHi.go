package main

import "fmt"

func sayHi() {
	fmt.Println("Theze Nutz")
}

func main() {
	// declaration of a variable holding a function
	var myFunction func()
	myFunction = sayHi
	myFunction()
}
