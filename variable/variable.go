package variable

import "fmt"

//variable in go, inside scope
var isFortType bool = true

//const in go, 
// const in go is initialize wi Upper word, and dont need to user sort declaration
const Gravity = 9.8

//declare multiple variable
var name, age = "Matias", 23

func ShowVariable() {
	//variable in scope
	var goIsOrientedObject bool
	fmt.Println("Go is orieted objest?", goIsOrientedObject)

	//init variable with sort declaration :=, this is create a variable and init
	goIsFast := true
	fmt.Println("Go is fast?", goIsFast)

	//Print global variable
	fmt.Println("This are global variable\n name:", name ," age:", age )

	//ths is numerics primitive type declaration
	var numberPhone int = 922831485
	pi := 3.14
	fmt.Println("This is numercs primative types in Go\n numberPhone:", numberPhone, ", pi:", pi,)

	//const in go
	fmt.Println("This a go const\n gravity", Gravity)

}
