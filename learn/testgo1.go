package main

import "fmt"
import "time"

func swap(x string, y string) (string, string) {
	return y, x
}
func printc() {
	fmt.Println("golang1")
}
func printd() {
	fmt.Println("golang2")
}

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	x := "test"
	y := "xyz"
	x, y = swap(x, y)
	go printc()
	go printd()
	time.Sleep(time.Duration(2) * time.Second)
	fmt.Println(x)
	fmt.Println(y)
}
