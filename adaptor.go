package dmx

import "github.com/hybridgroup/gobot"

// Adaptor is anything that can output DMX.
type Adaptor interface {
	gobot.Adaptor
	OutputDMX(data map[int]byte, universeSize int) error
}
