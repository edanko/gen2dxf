package gen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type LabeltextData struct {
	TextType      string
	TextPositionU float64
	TextPositionV float64
	TextAngle     float64
	TextHeight    float64
	TextDirection int
	TextSide      string
	Text          string
	TextElements  int
	GeometryData  []*GeometryData
}

func readLabeltextData(s *bufio.Scanner) *LabeltextData {
	lt := &LabeltextData{}
	for s.Scan() {
		l := strings.SplitN(s.Text(), "=", 2)

		switch l[0] {
		case "GEOMETRY_DATA":
			g := readGeometryData(s)
			lt.GeometryData = append(lt.GeometryData, g)

		case "END_OF_LABELTEXT_DATA":
			return lt

		case "TEXT_TYPE":
			lt.TextType = l[1]

		case "TEXT_POSITION_U":
			lt.TextPositionU, _ = strconv.ParseFloat(l[1], 64)

		case "TEXT_POSITION_V":
			lt.TextPositionV, _ = strconv.ParseFloat(l[1], 64)

		case "TEXT_ANGLE":
			lt.TextAngle, _ = strconv.ParseFloat(l[1], 64)

		case "TEXT_HEIGHT":
			lt.TextHeight, _ = strconv.ParseFloat(l[1], 64)

		case "TEXT_DIRECTION":
			lt.TextDirection, _ = strconv.Atoi(l[1])

		case "TEXT_SIDE":
			lt.TextSide = l[1]

		case "TEXT":
			lt.Text = l[1]

		case "TEXT_ELEMENTS":
			lt.TextElements, _ = strconv.Atoi(l[1])

		default:
			fmt.Println("unknown field in labeltext data:", l[0])

		}
	}
	return lt
}
