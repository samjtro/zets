package main
import "fmt"

func zeroval(ival int) { //0
	ival = 0
}

func zeroptr(iptr *int) { //1
	*iptr = 0
}

func main() {
	i := 1
	zeroval(i) //3
	zeroptr(&i) //2
	fmt.Println("zerotptr: ",i)
	fmt.Println("pointer: ",&i)
}
