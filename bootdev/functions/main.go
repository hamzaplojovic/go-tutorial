package main

import "fmt"

func printName(name string) string {
	return "Hello " + name
}

// in a function, if two or more are the same type,
// we add the type on just the last same one parameters
func addition(a, b int) int {
	return a + b
}

// callbacks are supported in go
func printNameCallback(printName func(name string) string) (result string) {
	result = printName("Hamza")
	return result
}

func getNames() (name1 string, name2 string) {
	return "John", "Bob"
}

func getCords() (x, y int) {

	return // automatically returns x and y

}

func main() {
	fmt.Println(printName("Hamza"))
	fmt.Println(addition(10, 20))
	fmt.Print(printNameCallback(printName))

	name1, name2 := getNames()
	name3, _ := getNames() // ignore the name4 value

	fmt.Println(name1, " ", name2, " ", name3)

	fmt.Println(getCords())

}
