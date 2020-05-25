package main

import (
	"fmt"
	"strings"
)

var x1, y1, x2, y2 int

func main() {
	var W, H, N, X0, Y0, batmanX, batmanY int

	// Width and Height of the building
	fmt.Scan(&W, &H)

	// N: maximum number of turns before game over.
	fmt.Scan(&N)

	// Batman starting position
	fmt.Scan(&X0, &Y0)

	batmanX, batmanY = X0, Y0

	x1, y1 = 0, 0
	x2, y2 = W-1, H-1

	for {
		// bombDir: the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
		var bombDir string
		fmt.Scan(&bombDir)

		if strings.Contains(bombDir, "U") {
			y2 = batmanY - 1
		} else if strings.Contains(bombDir, "D") {
			y1 = batmanY + 1
		}

		if strings.Contains(bombDir, "L") {
			x2 = batmanX - 1
		} else if strings.Contains(bombDir, "R") {
			x1 = batmanX + 1
		}

		batmanX = x1 + (x2-x1)/2
		batmanY = y1 + (y2-y1)/2

		// the location of the next window Batman should jump to.
		fmt.Println(batmanX, batmanY)
	}
}
