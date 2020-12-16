package main

func Sums(arr []int) int {
	cont := 0
	for _, v := range arr {
		cont += v
	}
	return cont
}
