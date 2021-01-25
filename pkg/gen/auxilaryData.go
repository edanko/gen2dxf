package gen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type AuxilaryData struct {
	AuxilaryFunction int
}

func readAuxilaryData(s *bufio.Scanner) *AuxilaryData {
	aux := &AuxilaryData{}

	for s.Scan() {
		l := strings.SplitN(s.Text(), "=", 2)

		switch l[0] {
		case "END_OF_AUXILIARY_DATA":
			return aux

		case "AUXILIARY_FUNCTION":
			aux.AuxilaryFunction, _ = strconv.Atoi(l[1])

		default:
			fmt.Println("unknown field in aux function:", l[0])
		}
	}
	return aux
}
