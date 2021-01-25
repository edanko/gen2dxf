package gen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type MarkingData struct {
	MarkingSide             string
	MarkingType             string
	MarkingName             string
	MarkingUserType         int
	MarkingAttri            string
	MarkingAttriSb          string
	MarkingShipside         string
	MarkingAssemblyLow      string
	MarkingProfileThickness float64
	MarkingWeld             float64
	MarkingDirection        int
	NumberOfHeads           int
	GeometryValidFor        int
	DistanceY1Y2            float64
	MtrlSideU               float64
	MtrlSideV               float64
	Comment                 string
	InclinationAngle        float64
	Contour                 *Contour
}

func readMarkingData(s *bufio.Scanner) *MarkingData {
	m := &MarkingData{}
	for s.Scan() {
		l := strings.SplitN(s.Text(), "=", 2)

		switch l[0] {
		case "END_OF_MARKING_DATA":
			return m

		case "START_OF_CONTOUR":
			m.Contour = readContour(s)

		case "MARKING_SIDE":
			m.MarkingSide = l[1]
		case "MARKING_TYPE":
			m.MarkingType = l[1]
		case "MARKING_NAME":
			m.MarkingName = l[1]
		case "MARKING_USER_TYPE":
			m.MarkingUserType, _ = strconv.Atoi(l[1])
		case "MARKING_ATTRI":
			m.MarkingAttri = l[1]
		case "MARKING_ATTRI_SB":
			m.MarkingAttriSb = l[1]
		case "MARKING_SHIPSIDE":
			m.MarkingShipside = l[1]
		case "MARKING_ASSEMBLY_LOW":
			m.MarkingAssemblyLow = l[1]
		case "MARKING_PROFILE_THICKNESS":
			m.MarkingProfileThickness, _ = strconv.ParseFloat(l[1], 64)
		case "MARKING_WELD":
			m.MarkingWeld, _ = strconv.ParseFloat(l[1], 64)
		case "MARKING_DIRECTION":
			m.MarkingDirection, _ = strconv.Atoi(l[1])
		case "NUMBER_OF_HEADS":
			m.NumberOfHeads, _ = strconv.Atoi(l[1])
		case "GEOMETRY_VALID_FOR":
			m.GeometryValidFor, _ = strconv.Atoi(l[1])
		case "DISTANCE_Y1_Y2":
			m.DistanceY1Y2, _ = strconv.ParseFloat(l[1], 64)
		case "MTRL_SIDE_U":
			m.MtrlSideU, _ = strconv.ParseFloat(l[1], 64)
		case "MTRL_SIDE_V":
			m.MtrlSideV, _ = strconv.ParseFloat(l[1], 64)
		case "COMMENT":
			m.Comment = l[1]
		case "INCLINATION_ANGLE":
			m.InclinationAngle, _ = strconv.ParseFloat(l[1], 64)

		default:
			fmt.Println("unknown field in marking data:", l[0])
		}

	}
	return m
}
