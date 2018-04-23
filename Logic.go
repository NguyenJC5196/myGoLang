package main

import "fmt"

func main(){
	mapFunction()
}
func different(){
	if 7 % 2 == 0{
		fmt.Println("number is even")
	} else{
		fmt.Println("number is odd")
	}

	if 8 %4 == 0 {
		fmt.Println("can be divided")
	}

	if num := 9; num < 0{
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "bigger than 10")
	}
}
func loop(){
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	for{
		fmt.Println("loop for once")
		break
	}
	for n := 0; n <=5; n++ {
		if n % 2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
func mapFunction(){
	grandes := make(map[string]float32)

	grandes["Timmy"] = 42.7
	grandes["Jess"] = 92.1
	grandes["Sam"] = 67.3

	fmt.Println(grandes)
	TimsGrade := grandes["Timmy"]
	fmt.Println(TimsGrade)

	delete(grandes, "Timmy")
	fmt.Println(grandes)

	for k, v := range grandes {
		fmt.Println(k,":", v)
	}	
}