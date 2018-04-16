package main
import "fmt"

const Use16Bitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal uint16 // (0, 65535)
	brake_pedal uint16 
	steering_wheel int16 // (-32k, +32k) 
	top_speed_kmh float64
}
func (c *car) KmPerHour() float64{
	c.top_speed_kmh = 500
	return float64(c.gas_pedal) * (c.top_speed_kmh/Use16Bitmax)
}
func (c *car) MilePerHour() float64{
	c.top_speed_kmh = 500
	return float64(c.gas_pedal) * (c.top_speed_kmh/kmh_multiple)
}
func (c *car) New_top_speed(newSpeed float64){
	c.top_speed_kmh = newSpeed
}
func newer_top_speed(c car, speed float64) car{
	c.top_speed_kmh = speed
	return c
}
func main(){
	a_car := car{
		gas_pedal: 22564,
		brake_pedal: 0,
		steering_wheel: 351,
		top_speed_kmh: 99999.0,
	}
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.KmPerHour())
	fmt.Println(a_car.MilePerHour())

	// a_car.New_top_speed(500)
	a_car = newer_top_speed(a_car, 500)
	fmt.Println(a_car.KmPerHour())
	fmt.Println(a_car.MilePerHour())
}