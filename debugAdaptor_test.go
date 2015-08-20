package dmx

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDebugAdaptor(t *testing.T) {
	Convey("Given an adaptor", t, func() {
		a := NewDebugAdaptor()

		Convey("Connect should raise no errors", func() {
			So(a.Connect(), ShouldBeEmpty)
		})

		Convey("Finalize should raise no errors", func() {
			a.Connect()
			So(a.Finalize(), ShouldBeEmpty)
		})

		Convey("OutputDMX", func() {
			Convey("Should set LastOutput", func() {
				So(a.OutputDMX(map[int]byte{1: 255}, 512), ShouldBeNil)
				So(a.GetLastOutput(), ShouldResemble, map[int]byte{1: 255})
			})
			Convey("Should set LastUniverseSize", func() {
				So(a.OutputDMX(map[int]byte{1: 255}, 512), ShouldBeNil)
				So(a.GetLastUniverseSize(), ShouldEqual, 512)
			})
			Convey("Shouldn't cause a race condition", func() {
				go a.OutputDMX(map[int]byte{1: 255}, 512)
				go a.OutputDMX(map[int]byte{1: 10}, 512)
			})
		})
	})
}
