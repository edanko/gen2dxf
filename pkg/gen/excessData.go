package gen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type ExcessData struct {
	ExcessValue float64
}

func readExcessData(s *bufio.Scanner) *ExcessData {
	e := &ExcessData{}

	for s.Scan() {
		l := strings.SplitN(s.Text(), "=", 2)

		switch l[0] {
		case "END_OF_EXCESS_DATA":
			return e

		case "EXCESS_VALUE":
			e.ExcessValue, _ = strconv.ParseFloat(l[1], 64)

		default:
			fmt.Println("unknown field in excess data:", l[0])
		}
	}
	return e
}
