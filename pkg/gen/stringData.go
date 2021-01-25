package gen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type StringData struct {
	Type   string
	PosU   float64
	PosV   float64
	Angle  float64
	Height float64
	String string
}

func readStringData(s *bufio.Scanner) *StringData {
	sd := &StringData{}
	for s.Scan() {
		l := strings.SplitN(s.Text(), "=", 2)

		switch l[0] {
		case "END_OF_STRING_DATA":
			return sd

		case "STRING_TYPE":
			sd.Type = l[1]

		case "STRING_POSITION_U":
			sd.PosU, _ = strconv.ParseFloat(l[1], 64)

		case "STRING_POSITION_V":
			sd.PosV, _ = strconv.ParseFloat(l[1], 64)

		case "STRING_ANGLE":
			sd.Angle, _ = strconv.ParseFloat(l[1], 64)

		case "STRING_HEIGHT":
			sd.Height, _ = strconv.ParseFloat(l[1], 64)

		case "STRING":
			sd.String = l[1]

		default:
			fmt.Println("unknown field in string data:", l[0])
		}
	}
	return sd
}
