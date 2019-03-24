package main

import "log"

// Joystick represents a dual axisis potentiometer with a switch.
type Joystick struct {
	X, Y, SW             int
	zeroX, zeroY, zeroSW int
}

func NewJoystick() *Joystick {
	j := &Joystick{}
	return j
}

func (j *Joystick) Calibrate(x, y, sw int) {
	j.zeroSW = sw
	j.zeroX = x
	j.zeroY = y
}

func (j *Joystick) Update(x, y, sw int) {

	// TODO must do a better job of calibrating, this assumes the first packet 
	// we read will be an accurate reading of the joystick in its nutural(sp?) 
	// position
	if j.zeroX == 0 || j.zeroY == 0 {
		j.Calibrate(x, y, sw)
		return
	}

	if config.Debug {
		//log.Printf("update %d %d %d\n", x, y, sw)
	}

	// check for movement along the X
	deltaX := x - j.zeroX
	deltaY := y - j.zeroY
	deltaSW := sw - j.SW
	if deltaSW != 0 {
		log.Printf("changed click state: %d", deltaSW)
	}

	if deltaX != 0 || deltaY != 0 {
		log.Printf("  change vector %d / %d ", deltaX, deltaY)
	}

	// Update old variables
	j.SW = sw
	j.X = x
	j.Y = y
}
