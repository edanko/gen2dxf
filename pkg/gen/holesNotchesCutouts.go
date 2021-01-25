package gen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type HolesNotchesCutouts struct {
	Open              int
	DistOrigin        float64
	DistOriginToZeroP float64
	DistOriginV       float64
	Type              int
	Rotation          float64
	Mirror            int
	Name              string
	DistLeft          float64
	Contour           *Contour
}

func readHolesNotchesCutouts(s *bufio.Scanner) *HolesNotchesCutouts {
	h := &HolesNotchesCutouts{}
	for s.Scan() {
		l := strings.Split(s.Text(), "=")

		switch l[0] {
		case "END_OF_HOLES_NOTCHES_CUTOUTS":
			return h

		case "START_OF_CONTOUR":
			h.Contour = readContour(s)

		case "OPEN":
			h.Open, _ = strconv.Atoi(l[1])

		case "DIST_ORIGIN":
			h.DistOrigin, _ = strconv.ParseFloat(l[1], 64)

		case "DIST_ORIGIN_TO_ZEROP":
			h.DistOriginToZeroP, _ = strconv.ParseFloat(l[1], 64)

		case "DIST_ORIGIN_V":
			h.DistOriginV, _ = strconv.ParseFloat(l[1], 64)

		case "TYPE":
			h.Type, _ = strconv.Atoi(l[1])

		case "ROTATION":
			h.Rotation, _ = strconv.ParseFloat(l[1], 64)

		case "MIRROR":
			h.Mirror, _ = strconv.Atoi(l[1])

		case "NAME":
			h.Name = l[1]

		case "DIST_LEFT":
			h.DistLeft, _ = strconv.ParseFloat(l[1], 64)

		default:
			fmt.Println("unknown field in holes notches cutouts data:", l[0])
		}

	}
	return h
}
