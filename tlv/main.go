package main

import "flag"
import "log"
import "github.com/tarm/serial"
import "github.com/mobilerobot-io/tlv"

type Config struct {
	Port string
	Baud int
}

const (
	Joystick = 0x10
)

var (
	config Config
)

func init() {
	flag.StringVar(&config.Port, "port", "/dev/tty.wchusbserial1420", "serial port")
	flag.IntVar(&config.Baud, "baud", 9600, "baud rate")
}

func main() {
	flag.Parse()
	config := &serial.Config{
		// Name: "/dev/ttyAMA0",
		Name: config.Port,
		Baud: config.Baud,
		ReadTimeout: 1,
		Size: 8,
	}

	stream, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}

	tlv.AddHandler(Joystick, JoystickCB)
	tlv.ReadLoop(stream)
}

func JoystickCB(tlv *TLV) (err error) {
	if tlv.Type != Joystick {
		panic("type expected 0x1 got " + string(tlv.Type))
	}
	if tlv.Len != 8 {
		panic("length expected 8 got (%d) " + string(tlv.Len))
	}

	
	return nil
}
