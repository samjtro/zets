package main
import "fmt"

func zeroval(ival int) int { //0
	ival = 0
	return ival
}

func zeroptr(iptr *int) { //1
	*iptr = 0
}

func main() {
	i := 1
	b := zeroval(i) //3
	zeroptr(&i) //2
	fmt.Println("zerotptr: ",b)
	fmt.Println("pointer: ",i)
}
