package main

func Rep(name string) string {
	var rep string
	for i := 0; i < 5; i++ {
		rep += name
	}
	return rep
}
