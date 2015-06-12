package output

import (
	"fmt"
)

type DummyOutput struct{}

func (do *DummyOutput) Set(state State) (err error) {
	fmt.Printf("%v\n", state)
	return
}
