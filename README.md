# lige

This produces an HTTP API to control an [ENTTEC DMXUSB PRO device](http://www.enttec.com/?main_menu=Products&pn=70304&show=description).

It communicates with the device over USB through a serial interface.


# Installing

First install the library

```shell
go install github.com/saulshanabrook/lige
```

Then you have to make sure you can communicate to your device over serial.

For my mac, I had to follow [this tutorial](http://www.mommosoft.com/blog/2014/10/24/ftdi-chip-and-os-x-10-10/)
to be able for it to show up.

To tell if the driver is working, plug in the USB device make sure it shows
up  under `/dev/tty.<something>`

```
$ ls /dev/tty.*
/dev/tty.usbserial-EN158833
```

# Running

```
lige <COM path> # like /dev/tty.usbserial-EN158833
```

# Testing
Using [HTTTPie](https://github.com/jakubroztocil/httpie)

```
# set channel 1 to 100
http POST localhost:8080 1=100
# set channel 1 to full
http POST localhost:8080 1=255
```
