package main
import(
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i  { 
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() { 
	case time.Saturday, time.Sunday: 
		fmt.Println("weekend baby!")
	default: 
		fmt.Println("weekday :(")
	}

	t := time.Now()
	switch { 
	case t.Hour() < 12: 
		fmt.Println("< noon")
	default:
		fmt.Println("afternoon")	
	}
	
	whatAmI := func(i interface{}) { 
		switch t := i.(type) { 
		case bool:
			fmt.Println("bool!")
		case int:
			fmt.Println("int!")
		default:
			fmt.Printf("dunno type %T, try again.\n",t)
		}
	}
	whatAmI(true)
	whatAmI("hey")
}
