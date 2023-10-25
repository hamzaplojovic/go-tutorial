package variables

import "fmt"

func main() {
	// ==================================================================================================
	// ==================================================================================================
	// DECLARING VARIABLES, VARIABLE TYPES

	// var name string = "Hamza"
	// var age int = 15
	// fmt.Print(name, " ", age)
	// ==================================================================================================
	// ==================================================================================================
	// DECLARING VARIABLES WITH :=

	// var name string = "Hamza"
	// is the same as
	// name := "Hamza"

	// the second example is used instead of var and it can be used to support dynamic typing
	// where the type of the variable will be assuumed based on the value assigned
	// empty := ""
	// fmt.Print(empty)
	// ==================================================================================================
	// ==================================================================================================
	// TYPE INFERENCE

	// var i int = 10
	// j := i + 20
	// the type of j will be the type of i, because of type inference
	// fmt.Print(j)

	// ==================================================================================================
	// ==================================================================================================
	// SAME LINE DECLARATIONS
	// var name, surname string = "Hamza", "Plojovic"
	// name, surname := "Hamza", "Plojovic"

	// fmt.Println(name, surname)

	// ==================================================================================================
	// ==================================================================================================
	// CONSTRUCTOR FUNCTIONS, NUMERIC VALUES

	// floatNumber := 5.2
	// fmt.Println(int64(floatNumber))

	// ==================================================================================================
	// ==================================================================================================
	// CONSTANTS

	const constantValue string = "Hamza"

	fmt.Println(constantValue)

	// ==================================================================================================
	// ==================================================================================================

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
