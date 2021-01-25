package gen

import (
	"bufio"
	"fmt"
	"strings"
)

type BumpData struct {
	Shape   string
	Contour *Contour
}

func readBumpData(s *bufio.Scanner) *BumpData {
	m := &BumpData{}

	for s.Scan() {
		l := strings.SplitN(s.Text(), "=", 2)

		switch l[0] {
		case "END_OF_BUMP_DATA":
			return m

		case "START_OF_CONTOUR":
			m.Contour = readContour(s)

		case "SHAPE":
			m.Shape = l[1]

		default:
			fmt.Println("unknown field in bump data:", l[0])

		}
	}
	return m
}
