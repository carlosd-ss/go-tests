package main

import "fmt"

const hi = "Hellow "

func Hellow(name string) string {
	return hi + name
}

func main() {
	fmt.Println(Hellow("Carlos"))
}
