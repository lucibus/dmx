package dmx

import (
	"sync"

	"github.com/hybridgroup/gobot"
)

// verify that it fulfils the Adaptor interface when compiled
var _ Adaptor = (*DebugAdaptor)(nil)
var _ gobot.Adaptor = (*DebugAdaptor)(nil)

// DebugAdaptor is a debugging DMX adaptor. It will set the Use the GetLastOuput and
// GetLastUniverseSize attributes to get the last values set by `OutputDMX` is
// called, so you can check them for debugging purposes
type DebugAdaptor struct {
	lastOutput       map[int]byte
	lastUniverseSize int
	mutex            sync.RWMutex
}

// NewDebugAdaptor returns a new DebugAdaptor
func NewDebugAdaptor() *DebugAdaptor {
	return &DebugAdaptor{}
}

// Name is "Debug Adaptor".
func (a *DebugAdaptor) Name() string { return "Debug Adaptor" }

// Connect is a noop.
func (a *DebugAdaptor) Connect() (errs []error) { return }

// Finalize is a noop.
func (a *DebugAdaptor) Finalize() (errs []error) { return }

// OutputDMX sets just records the values.
func (a *DebugAdaptor) OutputDMX(output map[int]byte, universeSize int) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.lastOutput = output
	a.lastUniverseSize = universeSize
	return nil
}

// GetLastOutput returns the last output set.
func (a *DebugAdaptor) GetLastOutput() map[int]byte {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	return a.lastOutput
}

// GetLastUniverseSize returns the last univserse size set.
func (a *DebugAdaptor) GetLastUniverseSize() int {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	return a.lastUniverseSize
}
