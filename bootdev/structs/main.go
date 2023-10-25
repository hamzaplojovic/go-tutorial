package main

import "fmt"

type Car struct {
	name  string
	model string
	color string
}

func printCar(carStruct Car) string {
	return fmt.Sprintf("\nThis is %s and the model is %s in the beautiful color of %s", carStruct.name, carStruct.model, carStruct.color)
}

// we can also make a method for a struct

func (car Car) printCarMethod() {
	formatted := fmt.Sprintf("\nFROM THE METHOD: This is %s and the model is %s in the beautiful color of %s", car.name, car.model, car.color)
	fmt.Println(formatted)

}

func main() {

	car1 := Car{
		name:  "Ford",
		model: "Mustang",
		color: "red",
	}

	fmt.Println(printCar(car1))
	car1.printCarMethod()

	// go supports anonymous structs

	anonymousStruct := struct {
		name  string
		model string
		color string
	}{
		name:  "Ford",
		model: "Mustang",
		color: "red",
	}

	fmt.Printf("\n%s", anonymousStruct)

	// embeded structs are kind of like a spred operator
	type notEmbededStruct struct {
		name string
	}

	type embededStruct struct {
		notEmbededStruct
		model string
	}

}
