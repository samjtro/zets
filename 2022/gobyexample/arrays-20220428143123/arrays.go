package main
import "fmt"

func main() {
	var a [5]int //0
	fmt.Println("emp:" ,a) //1

	a[4] = 100 //2
	fmt.Println("set:" ,a)
	fmt.Println("get:", a[4]) 
	fmt.Println("len:", len(a)) //3
	b := [5]int{ 1,2,3,4,5 } 
	fmt.Println("arr2:", b)

	var twoD [2][3]int //4
	for i:=0; i<2; i++ {
		for j:=0; j<3; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2d: ",twoD) //5
}
