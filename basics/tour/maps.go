package tour
import "strings"

func WordCount(s string) map[string]int {
	// init map
	var m = make(map[string]int)

	// split into fields
	fields := strings.Fields(s)

	// go through fields
	for _, value := range fields {
		m[value] = m[value] + 1
	}

	return m
}
