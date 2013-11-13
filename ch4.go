package main

import "fmt"

var evilGlobalFahrenheit float64 
var evilGlobalFeet float64

func main() {

	// Question 2
	x := 5
	x += 1
	fmt.Println(x)	

	// Question 5
	fmt.Print("Enter a fahrenheit value: ")
	fmt.Scanf("%f\n", &evilGlobalFahrenheit)
	fmt.Println("Celsius: ", calcCelsius(evilGlobalFahrenheit))

	// Question 6
	fmt.Print("Enter a quantity of feet: ")
	fmt.Scanf("%f\n", &evilGlobalFeet)
	fmt.Println("Feet: ", calcMeters(evilGlobalFeet))

}

func calcCelsius(fahr float64) float64 {
	x := (fahr - 32) * 5 / 9
	return x
}

func calcMeters(feet float64) float64 {
	return feet * 0.3048
}
