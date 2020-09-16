package main

import "fmt"

func sayHi() {
	fmt.Println("Theze Nutz 2")
}

func sayBye() {
	fmt.Println("Bye Bye Theze Nutz 2")
}

func twice(theFunction func()) {
	theFunction()
	theFunction()
}

func main() {

	twice(sayHi)
	twice(sayBye)
}
