package main

import (
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/lucibus/dmx"
)

func main() {
	gbot := gobot.NewGobot()

	adaptor := dmx.NewUSBEnttecProAdaptor("dmx", "/dev/tty.usbserial-EN158833")
	driver := dmx.NewDriver(adaptor, "dmx")

	work := func() {
		gobot.Every(3*time.Millisecond, func() {
			driver.OutputDMX(map[int]byte{1: byte(gobot.Rand(255))})
		})
	}

	robot := gobot.NewRobot("dmx",
		[]gobot.Connection{adaptor},
		[]gobot.Device{driver},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
