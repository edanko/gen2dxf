package gen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type TransformationData struct {
	TransformationType string
	Transformation1    float64
	Transformation2    float64
	Transformation3    float64
	Transformation4    float64
	Transformation5    float64
	Transformation6    float64
}

func readTransformationData(s *bufio.Scanner) *TransformationData {
	td := &TransformationData{}

	for s.Scan() {
		l := strings.SplitN(s.Text(), "=", 2)

		switch l[0] {
		case "END_OF_TRANSFORMATION_DATA":
			return td

		case "TRANSFORMATION_TYPE":
			td.TransformationType = l[1]

		case "TRANSFORMATION_1":
			td.Transformation1, _ = strconv.ParseFloat(l[1], 64)

		case "TRANSFORMATION_2":
			td.Transformation2, _ = strconv.ParseFloat(l[1], 64)

		case "TRANSFORMATION_3":
			td.Transformation3, _ = strconv.ParseFloat(l[1], 64)

		case "TRANSFORMATION_4":
			td.Transformation4, _ = strconv.ParseFloat(l[1], 64)

		case "TRANSFORMATION_5":
			td.Transformation5, _ = strconv.ParseFloat(l[1], 64)

		case "TRANSFORMATION_6":
			td.Transformation6, _ = strconv.ParseFloat(l[1], 64)

		default:
			fmt.Println("unknown field in transformation data:", l[0])
		}
	}
	return td
}
