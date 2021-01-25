package gen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type BurningData struct {
	Shape            string
	StartEndInGap    string
	BevelDefined     string
	Direction        int
	Designation      string
	PartID           int
	Comment          string
	NumberOfHeads    int
	GeometryValidFor int
	DistanceY1Y2     float64
	GeometryData     *GeometryData
	Contour          *Contour
}

func readBurningData(s *bufio.Scanner) *BurningData {
	b := &BurningData{}
	for s.Scan() {
		l := strings.SplitN(s.Text(), "=", 2)

		switch l[0] {
		case "END_OF_BURNING_DATA":
			return b

		case "GEOMETRY_DATA":
			b.GeometryData = readGeometryData(s)

		case "START_OF_CONTOUR":
			b.Contour = readContour(s)

		case "SHAPE":
			b.Shape = l[1]
		case "START_END_IN_GAP":
			b.StartEndInGap = l[1]
		case "BEVEL_DEFINED":
			b.BevelDefined = l[1]
		case "DIRECTION":
			b.Direction, _ = strconv.Atoi(l[1])
		case "DESIGNATION":
			b.Designation = l[1]
		case "PART_ID":
			b.PartID, _ = strconv.Atoi(l[1])
		case "COMMENT":
			b.Comment = l[1]
		case "NUMBER_OF_HEADS":
			b.NumberOfHeads, _ = strconv.Atoi(l[1])
		case "GEOMETRY_VALID_FOR":
			b.GeometryValidFor, _ = strconv.Atoi(l[1])
		case "DISTANCE_Y1_Y2":
			b.DistanceY1Y2, _ = strconv.ParseFloat(l[1], 64)

		default:
			fmt.Println("unknown field in burning data:", l[0])
		}
	}
	return b
}
