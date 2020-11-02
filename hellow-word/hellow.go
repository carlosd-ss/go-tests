package main

import "fmt"

const hi = "Hellow "

func hellow(name string) string {
	return hi + name
}

func main() {
	fmt.Println(hellow("Carlos"))
}
