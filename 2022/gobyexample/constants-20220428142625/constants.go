package main
import(
	"fmt"
	"math"
)

const CONSTANT string = "hello, darwin"

func main() {
	fmt.Println(CONSTANT)
	const n = 500000000
	const d = 3e20 / n

	fmt.Println(d,int64(d),math.Sin(n))
}
