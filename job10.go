package main

import (
	"fmt"
	"strconv"
)

/*
10.	Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
	Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
	Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/
func main() {
	//create array with temperatures
	temps := [...]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 3, -4, -10}
	//create map for union temperatures
	maps := make(map[string][]float64)
	//create variable for range
	var rng int
	//cycle bypass array and puts temperatures by union in map
	for _, val := range temps {
		rng = (int(val / 10)) * 10
		if val < 0 && val > -10 { //handling case with temperatures from <10 and > -10, create union 0 and -0
			maps["-0"] = append(maps["-0"], val)
			continue
		}
		maps[strconv.Itoa(rng)] = append(maps[strconv.Itoa(rng)], val)

	}
	//cycle output union temperatures
	for i, val := range maps {
		fmt.Println(i, val)
	}
}
