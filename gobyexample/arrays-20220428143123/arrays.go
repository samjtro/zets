package main
import "fmt"

//arrays are numbered sequences of elements of a specific length

func main() {
	var a [5]int //creates an array that holds exactly 5 ints: type of element + length are both part of type
	fmt.Println("emp:" ,a) //by default arrays are zero-valued

	a[4] = 100 //set value at index using array[index]
	fmt.Println("set:" ,a)
	fmt.Println("get:", a[4]) //get value at index using array[index]
	fmt.Println("len:", len(a)) //use len to return the length of an array
	b := [5]int{ 1,2,3,4,5 } //declare + init on the same line
	fmt.Println("arr2:", b)

	var twoD [2][3]int //multi-dimensional array
	for i:=0; i<2; i++ {
		for j:=0; j<3; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2d: ",twoD) //arrays appear in the form [1 2 ...] when using Println
}
