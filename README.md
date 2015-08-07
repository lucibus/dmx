# DMX

DMX output control for [Gobot](https://github.com/lucibus/dmx).

> [DMX512](https://en.wikipedia.org/wiki/DMX512) (Digital Multiplex) is a standard for digital communication networks that are commonly used to control stage lighting and effects. It was originally intended as a standardized method for controlling light dimmers, which, prior to DMX512, had employed various incompatible proprietary protocols. It soon became the primary method for linking controllers (such as a lighting console) to dimmers and special effects devices such as fog machines and intelligent lights. DMX has also expanded to uses in non-theatrical interior and architectural lighting, at scales ranging from strings of Christmas lights to electronic billboards. DMX can now be used to control almost anything, reflecting its popularity in theatres and venues.


# Install

```bash
go get github.com/lucibus/dmx
```

# Usage

We currently only support the [ENTTEC DMX USB Pro](http://www.enttec.com/?main_menu=Products&pn=70304)
for outputting DMX signals. Please open issues or send PRs for other outputs
that you want to support.

You need to install the drivers for that device, so that it shows up as serial
device. On my mac I had to [follow this tutorial](http://www.mommosoft.com/blog/2014/10/24/ftdi-chip-and-os-x-10-10/)
and then it appeared as `/dev/tty.usbserial-EN158833`.

Check out `example/main.go` for example usage.

# Testing

To run the unit tests just use `go test`.

If you have the device working, you can do a full integration test with

```bash
go run example/main.go
```

You will likely have to modify the path to the port in the example file first.
This will randomly change the level of output address one every three seconds.