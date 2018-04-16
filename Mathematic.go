package main

import ("math"
	"fmt")
const s = "constant"
func main(){
	x, y := 8.26, 5.15
	s := "change here"
	var w1, w2 string = "Hey", "there"
	var a int64 = 23
	var b float64 = float64(a)
	const n = 500000000
	const d = 3e20 / n

	fmt.Println(add(x, y))
	fmt.Println(s)
	fmt.Println(addText(w1, w2))
	fmt.Println(b)
	fmt.Println(math.Sin(n), d)
}
func add(x, y float64) float64 {
	return x + y
}
func addText(a, b string) string{
	return a + " " +b
}

