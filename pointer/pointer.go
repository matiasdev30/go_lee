package pointer


func TryChangeValueToZeroWithPassValue(i int){
	i = 0
}

func TryChangeValueToZeroWithPointerValue(i * int){
	*i = 0
}