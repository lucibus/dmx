package dmx

import "github.com/hybridgroup/gobot"

// verify that it fulfils the Adaptor interface when compiled
var _ Adaptor = (*DebugAdaptor)(nil)
var _ gobot.Adaptor = (*DebugAdaptor)(nil)

// DebugAdaptor is a debugging DMX adaptor. It will set the LastOuput and
// LastUniverseSize attributes whenever `OutputDMX` is called, so you
// can check them for debugging purposes
type DebugAdaptor struct {
	LastOutput       map[int]byte
	LastUniverseSize int
}

// NewDebugAdaptor returns a new DebugAdaptor
func NewDebugAdaptor() *DebugAdaptor {
	return &DebugAdaptor{
		LastOutput: map[int]byte{},
	}
}

// Name is "Debug Adaptor".
func (a *DebugAdaptor) Name() string { return "Debug Adaptor" }

// Connect is a noop.
func (a *DebugAdaptor) Connect() (errs []error) { return }

// Finalize is a noop.
func (a *DebugAdaptor) Finalize() (errs []error) { return }

// OutputDMX sets
func (a *DebugAdaptor) OutputDMX(data map[int]byte, universeSize int) error {
	a.LastOutput = data
	a.LastUniverseSize = universeSize
	return nil
}
