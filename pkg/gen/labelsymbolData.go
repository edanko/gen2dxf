package gen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type LabelsymbolData struct {
	SymbolType      string
	SymbolPositionU float64
	SymbolPositionV float64
	SymbolAngle     float64
	SymbolHeight    float64
	SymbolFont      int
	SymbolNumber    int
}

func readLabelsymbolData(s *bufio.Scanner) *LabelsymbolData {
	ls := &LabelsymbolData{}
	for s.Scan() {
		l := strings.SplitN(s.Text(), "=", 2)

		switch l[0] {
		case "END_OF_LABELSYMBOL_DATA":
			return ls

		case "SYMBOL_TYPE":
			ls.SymbolType = l[1]

		case "SYMBOL_POSITION_U":
			ls.SymbolPositionU, _ = strconv.ParseFloat(l[1], 64)

		case "SYMBOL_POSITION_V":
			ls.SymbolPositionV, _ = strconv.ParseFloat(l[1], 64)

		case "SYMBOL_ANGLE":
			ls.SymbolAngle, _ = strconv.ParseFloat(l[1], 64)

		case "SYMBOL_HEIGHT":
			ls.SymbolHeight, _ = strconv.ParseFloat(l[1], 64)

		case "SYMBOL_FONT":
			ls.SymbolFont, _ = strconv.Atoi(l[1])

		case "SYMBOL_NUMBER":
			ls.SymbolNumber, _ = strconv.Atoi(l[1])

		default:
			fmt.Println("unknown field in labelsymbol data:", l[0])

		}
	}
	return ls
}
