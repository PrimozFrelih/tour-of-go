package mutex

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"testing"
	"time"
)

const key = "some_key"

func TestMutex(t *testing.T) {
	s := SafeStruct{
		m: map[string]int{},
	}

	// run 10k go routines that increment the key
	for i := 0; i < 10*1000; i++ {
		go s.Increment(key)
	}

	time.Sleep(time.Millisecond * 10) // go is actually pretty fast :)
	fmt.Printf(humanize.Comma(int64(s.Value(key))))
}
