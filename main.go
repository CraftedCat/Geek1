package main

import (
	"fmt"
	"math"
)

func main() {
	var rub float64
	var kate1 float64
	var kate2 float64
	fmt.Println("Введите ваше кол-во рублей")
	fmt.Scanln(&rub)
	fmt.Println(currency(rub))
	fmt.Println("Введите длины катетов в сантиметрах через пробел:")
	fmt.Scanln(&kate1, &kate2)
	println("Площадь:", tectonic(kate1, kate2), "квадртаных сантиметров")

}

func Round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}

func currency(rub float64) float64 {
	curs := 65.5
	var usd float64
	usd = rub / curs
	return Round(usd, 2)
}

func tectonic(kate1, kate2 float64) float64 {
	println(kate1)
	println(kate2)
	return Round((kate1*kate2)/2*10, 2)
}
