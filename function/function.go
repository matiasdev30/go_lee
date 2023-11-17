package function

import (
	"fmt"
	"math/rand"
)

// this a simple public function in GO, without parametres and returns
func ThisIsSimplePublicFunctionInGO() {
	println("\nThis is a public simpler function in go")
}

//this is a private function in Go, with return
func returnRadonNumer(rangeValue int) int {
	return rand.Intn(10)
}

//function with funcion give in parameter
func ShowRadonNumber(rangeValue int){
	fmt.Println("Radon number", returnRadonNumer(rangeValue))
}

//Function with multiple return
func CalculateDivisionAndRestOfDivision(number int) (bool, int) {
	var isPar bool

	//condition operators in GO
	if number % 2 == 0 {
		isPar = true
	}else{
		isPar = false
	}

	return  isPar, number % 2
}

//function returning function
func IntSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

//function variadics
func VariadicFunction(num ...int){
	for c := 0; c < len(num); c++{
		println(num[c])
	}
}