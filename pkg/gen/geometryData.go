package gen

import (
	"bufio"
	"fmt"
	"strings"
)

type GeometryData struct {
	Type    string
	Contour *Contour
}

func readGeometryData(s *bufio.Scanner) *GeometryData {
	m := &GeometryData{}

	for s.Scan() {
		l := strings.SplitN(s.Text(), "=", 2)

		switch l[0] {
		case "END_OF_GEOMETRY_DATA":
			return m

		case "START_OF_CONTOUR":
			m.Contour = readContour(s)

		case "GEOMETRY_TYPE", "TYPE":
			m.Type = l[1]

		default:
			fmt.Println("unknown field in geometry data:", l[0])

		}
	}
	return m
}
