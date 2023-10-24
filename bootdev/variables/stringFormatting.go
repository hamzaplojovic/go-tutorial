package main

import "fmt"

func main() {
	const name string = "Hamza"
	const age int = 15
	formattedName := fmt.Sprintf("My name is %s and I am learning go", name)
	formattedAge := fmt.Sprintf("My age is %d", age)
	formattedWeather := fmt.Sprintf("Its currently %f degrees outside", 30.2)
	formattedWeatherTwoDecimals := fmt.Sprintf("Its currently %.2f degrees outside", 30.2)

	fmt.Println(formattedName)
	fmt.Println(formattedAge)
	fmt.Println(formattedWeather)
	fmt.Println(formattedWeatherTwoDecimals)

}
