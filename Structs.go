package main
import "fmt"

type car struct {
	gas_pedal uint16 // (0, 65535)
	brake_pedal uint16 
	steering_wheel int16 // (-32k, +32k) 
	top_speed_kmh float64
}
func main(){
	a_car := car{
		gas_pedal: 22564,
		brake_pedal: 0,
		steering_wheel: 351,
		top_speed_kmh: 99999.0,
	}
	fmt.Println(a_car.gas_pedal)
}