package main
import "fmt"

func vals(a,b int) (int, int) { //0	
	c := a+b
	d := a*b
	return c,d
}

func main() {
	var a,b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	_,i := vals(a,b) //1
	fmt.Println(i)

}
