package dmx

// Modeled on the API docs http://www.enttec.com/docs/dmx_usb_pro_api_spec.pdf

import (
	"fmt"

	"github.com/hybridgroup/gobot"
)

// MaxUniverseSize is the most number of DMX addresses the device can output.
const MaxUniverseSize = 512

// MinUniverseSize is the lowest number of DMX addresses the device can output.
const MinUniverseSize = 24

// InvalidUniverseSizeError is raised when you call `SetUniverseSize` with a
// universe size below `MinUniverseSize` or above `MaxUniverseSize`
type InvalidUniverseSizeError struct {
	UniverseSize int
}

func (e *InvalidUniverseSizeError) Error() string {
	return fmt.Sprintf("dmx: universe size %v must be between (inclusive) %v and %v", e.UniverseSize, MinUniverseSize, MaxUniverseSize)
}

// verify that it fulfils the Driver interface when compiled
var _ gobot.Driver = (*Driver)(nil)

type packet struct {
	header   []uint8
	body     []uint8
	checksum uint8
}

// Driver represents a DMX USB Pro
type Driver struct {
	name          string
	adaptor       *Adaptor
	packetChannel chan *packet
	universeSize  int
	cmd           gobot.Commander
}

// NewDriver returns a new Driver given an Adaptor and name.
//
// Adds the following API Commands:
// 	"OutputDMX" - See Driver.OutputDMX
// 	"SetUniverseSize" - See Driver.SetUniverseSize
func NewDriver(a *Adaptor, name string) *Driver {
	s := &Driver{
		name:         name,
		adaptor:      a,
		cmd:          gobot.NewCommander(),
		universeSize: MaxUniverseSize,
	}

	s.cmd.AddCommand("OutputDMX", func(params map[string]interface{}) interface{} {
		return s.OutputDMX(params["data"].(map[int]byte))

	})

	s.cmd.AddCommand("SetUniverseSize", func(params map[string]interface{}) interface{} {
		return s.SetUniverseSize(params["universeSize"].(int))
	})

	return s
}

// Name returns the name
func (s *Driver) Name() string { return s.name }

// Connection returns the connection
func (s *Driver) Connection() gobot.Connection { return gobot.Connection(s.adaptor) }

// Halt stops the connection.
func (s *Driver) Halt() []error {
	return s.adaptor.Finalize()
}

// Start starts the connection
func (s *Driver) Start() []error {
	return s.adaptor.Connect()
}

// SetUniverseSize sets the universe size of the output.
//
// It has a minumum size of 12 and max size of 512.
func (s *Driver) SetUniverseSize(size int) error {
	if size < MinUniverseSize || size > MaxUniverseSize {
		return &InvalidUniverseSizeError{size}
	}
	s.universeSize = size
	return nil
}

// OutputDMX outputs the mapping of DMX addresses and values.
//
// All addresses not provided will be set to the lowest value, up the the
// maximum address, which is the universe size.
func (s *Driver) OutputDMX(data map[int]byte) error {
	return s.adaptor.OutputDMX(data, s.universeSize)
}
