package condition

import (
	"runtime"
	"time"
)


func GetSystemRunTimeUsedSwithCaseCondition(){
	println(" get system with runtime module go")
	switch os := runtime.GOOS; os {
	case "darwin":
		print("OS X")
	case "linux" :
		println("linux")
	default:
		println(runtime.GOOS)
	} 
}

func GetSaudationWithConditionWithIfElse(){
	 hour := time.Now().Hour();
	if hour > 0 {
		println("Bom dia")
	}else if(hour > 12) {
		println("Boa tarde")
	}else{
		println("Boa noite")
		 println("deferimos")
	}
}
