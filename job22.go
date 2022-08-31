package main

import (
	"fmt"
	"math/big"
)

/*
22.	Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
*/

//check type
func getType(in interface{}) string {

	switch in.(type) {
	case int:
		return "int"
	case float64:
		return "float"
	default:
		return "Don't supports this type"
	}
}

func Multiply(a, b interface{}) {
	aType := getType(a)
	bType := getType(b)

	switch {
	case aType == "float" && bType == "int":
		a = new(big.Float).SetFloat64(float64(a.(float64)))          //create big float  from float64
		b = new(big.Float).SetFloat64(float64(b.(int)))              //create big float from int
		result := new(big.Float).Mul(a.(*big.Float), b.(*big.Float)) //multiply big float
		fmt.Printf("Multiply: %1.3f\n", result)
	case aType == "float" && bType == "float":
		a = new(big.Float).SetFloat64(float64(a.(float64)))
		b = new(big.Float).SetFloat64(float64(b.(float64)))
		result := new(big.Float).Mul(a.(*big.Float), b.(*big.Float))
		fmt.Printf("Multiply: %1.3f\n", result)
	case aType == "int" && bType == "float":
		a = new(big.Float).SetFloat64(float64(a.(int)))
		b = new(big.Float).SetFloat64(float64(b.(float64)))
		result := new(big.Float).Mul(a.(*big.Float), b.(*big.Float))
		fmt.Printf("Multiply: %1.3f\n", result)
	case aType == "int" && bType == "int":
		a = new(big.Int).SetInt64(int64(a.(int)))
		b = new(big.Int).SetInt64(int64(b.(int)))
		result := new(big.Int).Mul(a.(*big.Int), b.(*big.Int))
		fmt.Printf("Multiply: %d\n", result)
	default:
		fmt.Println("Error")
	}
}

func Addition(a, b interface{}) {
	aType := getType(a)
	bType := getType(b)

	switch {
	case aType == "float" && bType == "int":
		a = new(big.Float).SetFloat64(float64(a.(float64)))
		b = new(big.Float).SetFloat64(float64(b.(int)))
		result := new(big.Float).Add(a.(*big.Float), b.(*big.Float))
		fmt.Printf("Addition: %1.3f\n", result)
	case aType == "float" && bType == "float":
		a = new(big.Float).SetFloat64(float64(a.(float64)))
		b = new(big.Float).SetFloat64(float64(b.(float64)))
		result := new(big.Float).Add(a.(*big.Float), b.(*big.Float))
		fmt.Printf("Addition: %1.3f\n", result)
	case aType == "int" && bType == "float":
		a = new(big.Float).SetFloat64(float64(a.(int)))
		b = new(big.Float).SetFloat64(float64(b.(float64)))
		result := new(big.Float).Add(a.(*big.Float), b.(*big.Float))
		fmt.Printf("Addition: %1.3f\n", result)
	case aType == "int" && bType == "int":
		a = new(big.Int).SetInt64(int64(a.(int)))
		b = new(big.Int).SetInt64(int64(b.(int)))
		result := new(big.Int).Add(a.(*big.Int), b.(*big.Int))
		fmt.Printf("Addition: %d\n", result)
	default:
		fmt.Println("Error")
	}

}
func Subtraction(a, b interface{}) {
	aType := getType(a)
	bType := getType(b)

	switch {
	case aType == "float" && bType == "int":
		a = new(big.Float).SetFloat64(float64(a.(float64)))
		b = new(big.Float).SetFloat64(float64(b.(int)))
		result := new(big.Float).Sub(a.(*big.Float), b.(*big.Float))
		fmt.Printf("Subtraction: %1.3f\n", result)
	case aType == "float" && bType == "float":
		a = new(big.Float).SetFloat64(float64(a.(float64)))
		b = new(big.Float).SetFloat64(float64(b.(float64)))
		result := new(big.Float).Sub(a.(*big.Float), b.(*big.Float))
		fmt.Printf("Subtraction: %1.3f\n", result)
	case aType == "int" && bType == "float":
		a = new(big.Float).SetFloat64(float64(a.(int)))
		b = new(big.Float).SetFloat64(float64(b.(float64)))
		result := new(big.Float).Sub(a.(*big.Float), b.(*big.Float))
		fmt.Printf("Subtraction: %1.3f\n", result)
	case aType == "int" && bType == "int":
		a = new(big.Int).SetInt64(int64(a.(int)))
		b = new(big.Int).SetInt64(int64(b.(int)))
		result := new(big.Int).Sub(a.(*big.Int), b.(*big.Int))
		fmt.Printf("Subtraction: %d\n", result)
	default:
		fmt.Println("Error")
	}

}

func Divide(a, b interface{}) {
	aType := getType(a)
	bType := getType(b)
	if CheckDivNull(b) {
		fmt.Println("Divide null")
		return
	}
	switch {
	case aType == "float" && bType == "int":
		a = new(big.Float).SetFloat64(float64(a.(float64)))
		b = new(big.Float).SetFloat64(float64(b.(int)))
		result := new(big.Float).Quo(a.(*big.Float), b.(*big.Float))
		fmt.Printf("Divide: %1.3f\n", result)
	case aType == "float" && bType == "float":
		a = new(big.Float).SetFloat64(float64(a.(float64)))
		b = new(big.Float).SetFloat64(float64(b.(float64)))
		result := new(big.Float).Quo(a.(*big.Float), b.(*big.Float))
		fmt.Printf("Divide: %1.3f\n", result)
	case aType == "int" && bType == "float":
		a = new(big.Float).SetFloat64(float64(a.(int)))
		b = new(big.Float).SetFloat64(float64(b.(float64)))
		result := new(big.Float).Quo(a.(*big.Float), b.(*big.Float))
		fmt.Printf("Divide: %1.3f\n", result)
	case aType == "int" && bType == "int":
		a = new(big.Float).SetFloat64(float64(a.(int)))
		b = new(big.Float).SetFloat64(float64(b.(int)))
		result := new(big.Float).Quo(a.(*big.Float), b.(*big.Float))
		fmt.Printf("Divide: %1.3f\n", result)
	default:
		fmt.Println("Error")
	}
}

func CheckDivNull(in interface{}) bool {
	if in == 0 {
		return true
	}
	return false
}

func main() {
	Addition(40417401740.0, 57240862083.0)
	Subtraction(40417401740, 57240862083.0)
	Multiply(40417401740, 57240862083)
	Divide(40417401740.0, 57240862083)

}
