package main
import "fmt"

func main() {
	m := make(map[string]int) //0
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:",m) //1
	
	v1 := m["k1"]
	fmt.Println("k1:",v1)
	fmt.Println("len:",len(m)) //2
	
	delete(m,"k2") //3
	fmt.Println("map:",m)

	_, prs := m["k2"] //4
	fmt.Println("prs:",prs)

	n := map[string]int{ "one": 1, "two": 2 } //5
	fmt.Println("map:",n) //6
}
