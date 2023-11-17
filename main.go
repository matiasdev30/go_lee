package main

//import modules
import (
	//import variable modules
	"fmt"
	"net/http"

	"github.com/matiasdev30/go_lee/condition"
	exampleproject "github.com/matiasdev30/go_lee/exampleProject"
	"github.com/matiasdev30/go_lee/loop"
	"github.com/matiasdev30/go_lee/pointer"
	"github.com/matiasdev30/go_lee/variable"

	//import function module
	"github.com/matiasdev30/go_lee/function"
)

func main() {
	//acess public function in variable module
	variable.ShowVariable()

	//function module
	//function with parametre
	function.ShowRadonNumber(10)
	//Function without return
	function.ThisIsSimplePublicFunctionInGO()
	//function with return, show inside print function
	isPar, rest := function.CalculateDivisionAndRestOfDivision(8)
	fmt.Println("Is par", isPar, "rest ", rest )
	sl := []int{2,6,6,42,2,7,5,80,54,34}
	function.VariadicFunction(sl...)

	//loop module
	loop.MutiplicateValueWith12(4)
	loop.LoopForLikeWhile()
	loop.InfinityLoopWithBreak()

	//Condition module
	condition.GetSystemRunTimeUsedSwithCaseCondition()
	condition.GetSaudationWithConditionWithIfElse()

	//web service project
	
	http.HandleFunc("/", exampleproject.Index)
	http.ListenAndServe(":8080", nil)

	//Pointer module
	//Try to change value with pass value
	i := 4
	pointer.TryChangeValueToZeroWithPassValue(i)
	println("Try change value with pass value", i)
	//Try change value by reference
	pointer.TryChangeValueToZeroWithPointerValue(&i)
	//Try change value with pointer
	println("Try change value with pointer value", i, ", adress,", &i)
}
