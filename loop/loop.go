package loop

//In Go only have loop for
func MutiplicateValueWith12(value int){
	for c := 1; c < 13; c++ {
		println(value,"x", c , "= ", value * c)
	}
}

//Make loop for like a while
func LoopForLikeWhile(){
	c := 0
	for c<4 {
		println("@matiasdev")
		c ++
	}
}

//Infinity loop with break
func InfinityLoopWithBreak(){
	c := 0
	for{
		c ++
		println("@fvckingnullinnirvana")
		if(c == 8){
			break
		}
	}
}