package main

import "fmt"

func main(){
	x := 15
	a := &x
	
	// *a = 59

	fmt.Println(x)

	*a = *a * *a
	fmt.Println(x)
}