package gen

import (
	"bufio"
	"fmt"
)

type ConnectionTrace struct {
	Contour *Contour
}

func readConnectionTrace(s *bufio.Scanner) *ConnectionTrace {
	ct := &ConnectionTrace{}

	for s.Scan() {
		c := s.Text()

		switch c {
		case "END_OF_CONNECTION_TRACE":
			return ct

		case "START_OF_CONTOUR":
			ct.Contour = readContour(s)

		default:
			fmt.Println("unknown connection trace field:", c)
		}
	}
	return ct
}
