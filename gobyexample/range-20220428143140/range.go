package main
import "fmt"

func main() {
	nums := []int{ 2,3,4 }
	sum := 0

	for _,num := range nums { //0
		sum += num
	}
	for i,num := range nums { //1
		if num == 3 {
			fmt.Println(i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k,v := range kvs { //2
		fmt.Printf("%s -> %s\n",k,v)
	}
	for k := range kvs { //3
		fmt.Printf("%s\n",k)
	}
	for _,v := range kvs {
		fmt.Printf("%s\n",v)
	}

	for i,c := range "go" { //4
		fmt.Println(i,c)
	}
}
