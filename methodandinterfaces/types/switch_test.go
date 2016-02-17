package types

import "testing"

func TestSwitch(t *testing.T) {
	processSwitch(5)
	processSwitch("test")
	processSwitch(false)
}
