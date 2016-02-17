package types

import "fmt"

func processSwitch(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Printf("%d is int\n", i)
	case string:
		fmt.Printf("%s is string\n", i)
	default:
		fmt.Printf("%v is unknown type\n", i)
	}
}
