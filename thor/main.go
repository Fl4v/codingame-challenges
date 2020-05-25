package main

import (
	"fmt"
)

var lightX, lightY, initialTX, initialTY int
var directionX string
var directionY string
var thorX int
var thorY int

func main() {

	fmt.Scan(&lightX, &lightY, &initialTX, &initialTY)
	thorX = initialTX
	thorY = initialTY

	for {
		//var remainingTurns int
		//fmt.Scan(&remainingTurns)

		directionX = ""
		directionY = ""

		if thorX > lightX {
			directionX = "W"
			thorX--
		} else if thorX < lightX {
			directionX = "E"
			thorX++
		}
		if thorY > lightY {
			directionY = "N"
			thorY--
		} else if thorY < lightY {
			directionY = "S"
			thorY++
		}
		fmt.Println(directionY+directionX, thorX, thorY)
		if thorX == lightX && thorY == lightY {
			break
		}
	}
}
