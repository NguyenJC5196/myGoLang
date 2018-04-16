package main
import ("fmt"
"math/rand")

func main(){
	fmt.Println("Welcome to Go")
	foo()
}
func foo(){
	fmt.Println("The random number is, ",rand.Intn(100))
}
func variables(){
	fmt.Println("the add, ", 2 + 2)
	fmt.Println("the divide, ", 7.0 / 3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}