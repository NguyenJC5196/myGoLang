package main
import ("fmt"
"math/rand")
import "math"

const s string = "constant"

func main(){
	fmt.Println("Welcome to Go")
	// num1,num2 :=5.5,9.6
	w1,w2:="Hey","there"
	var a int=86
	var b float64=float64(a)
	x := a
	const n = 500000000
	const d = 3e20 / n

	// fmt.Println(add(num1,num2))
	fmt.Println(multiple(w1,w2))
	fmt.Println( a, b, x)
	fmt.Println(s)
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
func foo(){
	fmt.Println("The random number is, ",rand.Intn(100))
}
func variables(){
	fmt.Println("the add, ",2+2)
	fmt.Println("the divide, ",7.0/3.0)
	fmt.Println(true&&false)
	fmt.Println(true||false)
	fmt.Println(!true)
}
func add(x, y float32) float32{
	return x+y
}
func multiple(a,b string) (string,string){
	return a,b
}