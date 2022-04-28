package main
import(
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i  { //basic switch
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() { 
	case time.Saturday, time.Sunday: //using commas to seperate multiple expressions
		fmt.Println("weekend baby!")
	default: //and the optional default
		fmt.Println("weekday :(")
	}

	t := time.Now()
	switch { //switch without an expression, exemplefies if/else
	case t.Hour() < 12: //case expressions can be non-constants
		fmt.Println("< noon")
	default:
		fmt.Println("afternoon")	
	}
	
	whatAmI := func(i interface{}) { //use type switch to compare interface value
		switch t := i.(type) { //type switch compares types
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
