package dmx

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func initTestDriver() *Driver {
	a := NewAdaptor("bot", "/dev/null")
	a.sp = testReadWriteCloser{}
	return NewDriver(a, "bot")
}

func TestDriver(t *testing.T) {
	Convey("Given a driver", t, func() {
		a := NewAdaptor("bot", "/dev/null")
		a.sp = testReadWriteCloser{}
		d := NewDriver(a, "bot")

		Convey("The name should be correct", func() {
			So(d.Name(), ShouldEqual, "bot")
		})

		Convey("SetUniverseSize", func() {
			Convey("Should set the universe size", func() {
				d.SetUniverseSize(100)
				So(d.universeSize, ShouldEqual, 100)
			})
			Convey("should raise an error if too low", func() {
				So(d.SetUniverseSize(10), ShouldResemble, &InvalidUniverseSizeError{10})
			})
			Convey("should raise an error if too high", func() {
				So(d.SetUniverseSize(513), ShouldResemble, &InvalidUniverseSizeError{513})
			})

		})
	})
}
