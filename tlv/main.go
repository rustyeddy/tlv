package main

import (
	"flag"
	"io"
	"log"
	"strconv"
	"strings"

	// switch to the term package?
	"github.com/tarm/serial"
)

type Config struct {
	Baud    int
	Debug   bool
	Serial  string
	Serial2 string

	PinX, PinY, PinSW int
}

var (
	ser1, ser2 io.ReadWriteCloser
	config     Config
	joy        *Joystick
)

func init() {
	//flag.StringVar(&config.Port, "port", "/dev/tty.wchusbserial1420", "serial port")
	flag.StringVar(&config.Serial, "serial", "/dev/tty.usbserial-14110", "serial port")
	flag.StringVar(&config.Serial2, "serial2", "/dev/tty.usbserial15410", "serial port2")
	flag.IntVar(&config.Baud, "baud", 115200, "baud rate")
	flag.BoolVar(&config.Debug, "debug", true, "turn debug on or off")
	joy = NewJoystick()
}

func main() {
	flag.Parse()

	// switch to term?
	//t, err := term.Open("/dev/ttyUSB0", Speed(19200), RawMode)
	ser1 := getSerialStream(config.Serial)
	ser2 := getSerialStream(config.Serial2)
	readStreams(ser1, ser2)
}

func getSerialStream(port string) io.ReadWriteCloser {

	sercfg := &serial.Config{
		Name:        port,
		Baud:        config.Baud,
		ReadTimeout: 1,
		Size:        8,
	}
	stream, err := serial.OpenPort(sercfg)
	if err != nil {
		log.Fatal(err)
	}
	return stream
}

func readStreams(s1, s2 io.ReadWriteCloser) {
	buf := make([]byte, 16)
	for {
		n, err := s1.Read(buf)
		if err != nil && err != io.EOF {
			log.Println(err)
		}
		if n == 0 {
			continue
		}
		incoming := string(buf[:n])
		incoming = strings.Trim(incoming, "\n\r")
		parsed := strings.Split(incoming, ":")

		if config.Debug {
			//log.Printf("incoming %+v", parsed)
		}
		switch parsed[0] {
		case "cm":
			doDistance(parsed)

		case "j":
			doJoystick(parsed)

		default:
			log.Printf("unknown command %v", parsed[0])
		}
	}
}

func doJoystick(parsed []string) {
	var sw, x, y int
	var err error

	sw, err = strconv.Atoi(parsed[1])
	if err == nil {
		x, err = strconv.Atoi(parsed[2])
		if err == nil {
			y, err = strconv.Atoi(parsed[3])
		}
	}
	// Should put an error counter here
	if err != nil {
		log.Print(err)
		return
	}
	joy.Update(x, y, sw)
}

func doDistance(cmd []string) {
	dist, err := strconv.Atoi(cmd[1])
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Distance: %d cm", dist)
}
