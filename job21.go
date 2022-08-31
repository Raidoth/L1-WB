package main

import "fmt"

/*
21.	Реализовать паттерн «адаптер» на любом примере
*/

//create interface with function get temperature in celsius
type Sensor interface {
	GetTemperatureCelsius() float64
}

//create struct sensor with celsius
type SensorCelsius struct {
}

//get temperature in celsius from sensor
func (s *SensorCelsius) GetTemperatureCelsius() float64 {
	return 20
}

//create struct sensor with fahrenheit
type SensorFah struct {
}

//get temperature in fahrenheit from sensor
func (s *SensorFah) GetTemperatureFar() float64 {
	return 55.0
}

//create struct sensor adapter fahrenheit with pointer on struct
type SernsorFarAdapter struct {
	sensFar *SensorFah
}

//translation fahrenheit temperature in celsius
func (s *SernsorFarAdapter) GetTemperatureCelsius() float64 {
	return ((s.sensFar.GetTemperatureFar() - 32) * 5 / 9)
}

func main() {
	//create sensor fahrenheit and celsius
	var fah Sensor = &SernsorFarAdapter{&SensorFah{}} //12,78
	var cel Sensor = &SensorCelsius{}                 //20
	//compute general temperature in celsius
	fmt.Println(fah.GetTemperatureCelsius() + cel.GetTemperatureCelsius()) //32,78

}
