package main
import "fmt"

func main() {
	s := make([]string, 3) 
	s[0] = "a" 
	s[1] = "B"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[0], s[1])
	fmt.Println("len:", len(s)) 

	s = append(s, "d") 
	s = append(s, "e", "f")
	fmt.Println("apd:",s)
	
	c := make([]string,len(s)) 
	copy(c,s) 
	fmt.Println("cpy:",c)

	l := s[:5] 
	fmt.Println("sl:",l)

	t := []string{ "g","h","i" } 
	fmt.Println("dcl:",t)

	twoD := make([][]int,3)
	for i:=0; i<3; i++ {
		innerLen := i+1
		twoD[i] = make([]int,innerLen)
		for j:=0; j<innerLen; j++ {
			twoD[i][j] = i+j
		}	
	}
	fmt.Println("2d:",twoD)
}
