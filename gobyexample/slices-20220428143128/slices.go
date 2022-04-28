package main
import "fmt"

// slices are a key data type in go
// more powerful interface to sequences than arrays
// slices are typed only by the elements they contain (not num of elements)
// blog post re: design+implementation of slices in go
// http://blog.golang.org/2011/01/go-slices-usage-and-internals.html

func main() {
	s := make([]string, 3) //create an empty slice with non-zero length using make()
	s[0] = "a" //set & get just like arrays
	s[1] = "B"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[0], s[1])
	fmt.Println("len:", len(s)) //len() works as well

	s = append(s, "d") //append returns a slice containing one or more new values *accepted by var, as we will get a new slice value
	s = append(s, "e", "f")
	fmt.Println("apd:",s)
	
	c := make([]string,len(s)) //slices can be copy'd; first by creating empy slice c (with len == s)
	copy(c,s) //then copy into c from s
	fmt.Println("cpy:",c)

	l := s[:5] //slices also support slicing, slice[lo:hi]
	fmt.Println("sl:",l)

	t := []string{ "g","h","i" } //slices can be declared+init on the same line as well
	fmt.Println("dcl:",t)

	twoD := make([][]int,3) //slices can also be multi-dimensional
	for i:=0; i<3; i++ {
		innerLen := i+1// the len of the inner slice my vary, unlike with multi-d arrays
		twoD[i] = make([]int,innerLen)
		for j:=0; j<innerLen; j++ {
			twoD[i][j] = i+j
		}	
	}
	fmt.Println("2d:",twoD)
}
