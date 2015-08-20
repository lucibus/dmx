package dmx

import (
	"fmt"
	"io"
	"sync"

	"github.com/hybridgroup/gobot"
	"github.com/tarm/goserial"
)

const baud = 57600

// InvalidAddressError is raised when you call `OutputDMX` with an address
// that is out of range. Either below 1 or above the universe size.
type InvalidAddressError struct {
	Address      int
	UniverseSize int
}

// NotConnectedError is raised when you call `OutputDMX` before calling
// `Connect`
type NotConnectedError struct{}

func (*NotConnectedError) Error() string {
	return "dmx: must call Connect before calling OutputDMX"
}

func (e *InvalidAddressError) Error() string {
	return fmt.Sprintf("dmx: address %v must be between (inclusive) %v and %v", e.Address, 1, e.UniverseSize)
}

// verify that it fulfils the Adaptor interface when compiled
var _ Adaptor = (*USBEnttecProAdaptor)(nil)
var _ gobot.Adaptor = (*USBEnttecProAdaptor)(nil)

// USBEnttecProAdaptor represents a Connection to a DMX USB Pro
type USBEnttecProAdaptor struct {
	name    string
	port    string
	sp      io.ReadWriteCloser
	connect func(port string) (io.ReadWriteCloser, error)
	mutex   sync.Mutex
}

// NewUSBEnttecProAdaptor returns a new Adaptor given a name and port
func NewUSBEnttecProAdaptor(name string, port string) *USBEnttecProAdaptor {
	return &USBEnttecProAdaptor{
		name: name,
		port: port,
		connect: func(port string) (io.ReadWriteCloser, error) {
			return serial.OpenPort(&serial.Config{Name: port, Baud: baud})
		},
	}
}

// Name returns the name.
func (a *USBEnttecProAdaptor) Name() string { return a.name }

// Port returns the port.
func (a *USBEnttecProAdaptor) Port() string { return a.port }

// Connect initiates a connection.
func (a *USBEnttecProAdaptor) Connect() (errs []error) {
	sp, err := a.connect(a.Port())
	if err != nil {
		return []error{err}
	}
	a.sp = sp
	return
}

// Finalize closes the connection.
func (a *USBEnttecProAdaptor) Finalize() (errs []error) {
	err := a.sp.Close()
	if err != nil {
		return []error{err}
	}
	return
}

// OutputDMX outputs the mapping of DMX addresses and values.
//
// All addresses not provided will be set to the lowest value, up the the
// maximum address, which is the universe size.
func (a *USBEnttecProAdaptor) OutputDMX(data map[int]byte, universeSize int) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if a.sp == nil {
		return &NotConnectedError{}
	}

	// The data we want to send out starts with the "start code" (0) and then
	// continues for as long as the universe size
	dataB := make([]byte, universeSize+1)
	for address, value := range data {
		if address > universeSize || address < 1 {
			return &InvalidAddressError{Address: address, UniverseSize: universeSize}
		}
		dataB[address] = value
	}
	return a.sendMessage(
		6, // message type code, corresponds to "Output Only Send DMX Packet Request"
		dataB,
	)
}

func leastSignificantByte(b int) byte {
	return byte(b & 0xFF)
}

// will only work for ints that are < 1024
func mostSignificantByte(b int) byte {
	return byte((b >> 8) & 0xFF)
}

func (a *USBEnttecProAdaptor) sendMessage(label int, data []byte) error {
	output := []byte{
		0x7E,                            // start of a message delimiter
		byte(label),                     // Label to identify type of message "Output Only Send DMX Packet Request"
		leastSignificantByte(len(data)), // Data length LSB
		mostSignificantByte(len(data)),  // Data length MSB
	}
	output = append(
		output,
		data...,
	)
	output = append(
		output,
		0xE7, // End of message delimiter
	)
	_, err := a.sp.Write(output)
	return err
}
