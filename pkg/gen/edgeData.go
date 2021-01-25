package gen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type EdgeData struct {
	Name    string
	StartU  float64
	StartV  float64
	EndU    float64
	EndV    float64
	ALength float64
	MLength float64
}

func readEdgeData(s *bufio.Scanner) *EdgeData {
	m := &EdgeData{}

	for s.Scan() {
		l := strings.SplitN(s.Text(), "=", 2)

		switch l[0] {
		case "END_OF_EDGE_DATA":
			return m

		case "NAME":
			m.Name = l[1]

		case "START_U":
			m.StartU, _ = strconv.ParseFloat(l[1], 64)

		case "START_V":
			m.StartV, _ = strconv.ParseFloat(l[1], 64)

		case "END_U":
			m.EndU, _ = strconv.ParseFloat(l[1], 64)

		case "END_V":
			m.EndV, _ = strconv.ParseFloat(l[1], 64)

		case "ALENGTH":
			m.ALength, _ = strconv.ParseFloat(l[1], 64)

		case "MLENGTH":
			m.MLength, _ = strconv.ParseFloat(l[1], 64)

		default:
			fmt.Println("unknown field in edge data:", l[0])
		}
	}
	return m
}
