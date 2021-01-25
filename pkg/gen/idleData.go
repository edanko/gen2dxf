package gen

import (
	"bufio"
	"fmt"
)

type IdleData struct {
	Contour *Contour
}

func readIdleData(s *bufio.Scanner) *IdleData {
	id := &IdleData{}

	for s.Scan() {
		switch s.Text() {
		case "END_OF_IDLE_DATA":
			return id

		case "START_OF_CONTOUR":
			id.Contour = readContour(s)

		default:
			fmt.Println("unknown idle data section:", s.Text())
		}
	}

	return id
}
