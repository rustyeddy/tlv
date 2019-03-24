package main

import (
	"flag"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/tarm/serial"
)

type Config struct {
	Baud  int
	Debug bool
	Port  string

	PinX, PinY, PinSW int
}

var (
	config Config
	joy    *Joystick
)

func init() {
	//flag.StringVar(&config.Port, "port", "/dev/tty.wchusbserial1420", "serial port")
	flag.StringVar(&config.Port, "port", "/dev/tty.usbserial-14110", "serial port")
	flag.IntVar(&config.Baud, "baud", 9600, "baud rate")
	flag.BoolVar(&config.Debug, "debug", true, "turn debug on or off")
	joy = NewJoystick()
}

func main() {
	flag.Parse()
	config := &serial.Config{
		Name:        config.Port,
		Baud:        config.Baud,
		ReadTimeout: 1,
		Size:        8,
	}

	stream, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}
	readLoop(stream)
}

func readLoop(stream io.Reader) {
	buf := make([]byte, 16)
	for {
		n, err := stream.Read(buf)
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
