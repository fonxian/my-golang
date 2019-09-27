package main

import "fmt"

type inner struct {
	ss4 string
}

type oo struct {
	inner
	ss1 string
	ss2 int
	ss3 bool
}

func main() {

	oo1 := new(oo)
	oo1.ss2 = 2
	fmt.Println(oo1.ss1)
	fmt.Println(oo1.ss2)
}
