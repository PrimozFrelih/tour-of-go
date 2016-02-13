package switch1
import (
	"testing"
	"runtime"
	"fmt"
	"time"
)

func TestSwitch(t *testing.T) {

	// set a var and use it in a switch
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println(os)
	}
}

// NOTE: switch in go breaks automatically, fallthrough keyword is used so the switch doesn't break automatically
func TestFallthrough(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "darwin":
		fallthrough
	case "linux":
		fmt.Println("Good")
	case "windows":
		fmt.Println("Bad, very bad!")
	}
}

func TestSwitch2(t *testing.T) {
	today := time.Now().Weekday()

	// switch on a constant
	switch time.Saturday {
	case today:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("After tomorrow")
	default:
		fmt.Println("Some day")
	}
}

func TestSwitch3(t *testing.T) {
	now := time.Now()

	// no condition
	switch {
	case now.Hour() < 12:
		fmt.Println("Morning")
	case now.Hour() < 18:
		fmt.Println("Afternoon")
	default:
		fmt.Println("Evening")
	}
}
