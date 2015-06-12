package output

import (
	"fmt"
	"time"

	"github.com/tarm/serial"
)

func leastSignificantByte(b int) byte {
	return byte(b & 0xFF)
}

// will only work for ints that are < 1024
func mostSignificantByte(b int) byte {
	return byte((b >> 8) & 0xFF)
}

/*
ENTTECUSBProOutput should be used with an ENTEC DMXUSB PRO device.
First you must install the drivers for the device and make it is available
on your filesystem. Then pass in the COM Port. (like /dev/tty.usbserial-EN158833)
*/
type ENTTECUSBProOutput struct {
	COMPort string
	port    *serial.Port // don't pass in when you init
}

func (eo *ENTTECUSBProOutput) init() (err error) {
	c := &serial.Config{Name: eo.COMPort, Baud: 57600, ReadTimeout: time.Second}
	eo.port, err = serial.OpenPort(c)
	return
}

func (eo *ENTTECUSBProOutput) transmit(label byte, data []byte) (err error) {
	output := []byte{
		0x7E,  // start of a message delimiter
		label, // Label to identify type of message "Output Only Send DMX Packet Request"
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
	fmt.Printf("%v\n", output)
	_, err = eo.port.Write(output)
	return
}

/*
Set will send the that state out to ENTTeC DMX Pro via serial over USB.

It uses this specification: https://www.enttec.com/docs/dmx_usb_pro_api_spec.pdf

And a sample implementation from python: https://github.com/c0z3n/pySimpleDMX/blob/master/pysimpledmx.py
*/
func (eo *ENTTECUSBProOutput) Set(state State) (err error) {
	if eo.port == nil {
		err = eo.init()
		if err != nil {
			return
		}
	}

	minPassed, maxPassed := IntRange(state.Keys())
	if minPassed < 0 {
		return fmt.Errorf("Can't send DMX channels less than 0")
	}
	const MaxChannel = 512
	if minPassed > MaxChannel {
		return fmt.Errorf("Can't send DMX channels greater than %v", MaxChannel)
	}

	data := []byte{0}
	for i := 1; i < MaxInt(maxPassed, MaxChannel); i++ {
		// if the channel `i` is not in `state` then it will simply
		// return `0` when trying to get the value, which is great,
		// so all channels not in `state` will be set to 0
		data = append(data, byte(state[i]))
	}
	err = eo.transmit(6, data)
	return
}
