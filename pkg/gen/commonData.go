package gen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type CommonData struct {
	NestName        string
	Shape           string
	Contour         *Contour
	Dimension       string
	WebHeight       float64
	WebThickness    float64
	MaxWebHeight    float64
	FlangeWidth     float64
	FlangeThickness float64
	RawLength       float64
	RestLength      float64
	UsedLength      float64
	LastScrapLength float64
	Quality         string
	Density         float64
	Area            float64
	TotalWeight     float64
	RawBar          string
	BuyingMark      string
	CuttingStn      string
	TextHeight      float64
	TextWidth       float64
	TextPlacing     int
	TextDistTop     float64
	TextMax1        int
	TextMax2        int
	TextHor1        int
	TextVer1        int
	TextHor2        int
	TextVer2        int
	TextHor3        int
	TextVer3        int
	Tscrap          int
	NoOfProfs       int
	Rest            int
	Gsd             int
	GsdDist         int
	GsdOneLim       int
}

func readCommonData(s *bufio.Scanner) *CommonData {
	cd := &CommonData{}

	for s.Scan() {
		l := strings.SplitN(s.Text(), "=", 2)

		switch l[0] {
		case "END_OF_COMMON_DATA":
			return cd

		case "START_OF_CONTOUR":
			cd.Contour = readContour(s)

		case "NEST_NAME":
			cd.NestName = l[1]

		case "SHAPE":
			cd.Shape = l[1]

		case "DIMENSION":
			cd.Dimension = l[1]

		case "WEB_HEIGHT":
			cd.WebHeight, _ = strconv.ParseFloat(l[1], 64)

		case "WEB_THICKNESS":
			cd.WebThickness, _ = strconv.ParseFloat(l[1], 64)

		case "MAX_WEB_HEIGHT":
			cd.MaxWebHeight, _ = strconv.ParseFloat(l[1], 64)

		case "FLANGE_WIDTH":
			cd.FlangeWidth, _ = strconv.ParseFloat(l[1], 64)

		case "FLANGE_THICKNESS":
			cd.FlangeThickness, _ = strconv.ParseFloat(l[1], 64)

		case "RAW_LENGTH":
			cd.RawLength, _ = strconv.ParseFloat(l[1], 64)

		case "REST_LENGTH":
			cd.RestLength, _ = strconv.ParseFloat(l[1], 64)

		case "USED_LENGTH":
			cd.UsedLength, _ = strconv.ParseFloat(l[1], 64)

		case "LAST_SCRAP_LENGTH":
			cd.LastScrapLength, _ = strconv.ParseFloat(l[1], 64)

		case "QUALITY":
			cd.Quality = l[1]

		case "DENSITY":
			cd.Density, _ = strconv.ParseFloat(l[1], 64)

		case "AREA":
			cd.Area, _ = strconv.ParseFloat(l[1], 64)

		case "TOTAL_WEIGHT":
			cd.TotalWeight, _ = strconv.ParseFloat(l[1], 64)

		case "RAW_BAR":
			cd.RawBar = l[1]

		case "BUYING_MARK":
			cd.BuyingMark = l[1]

		case "CUTTING_STN":
			cd.CuttingStn = l[1]

		case "TEXT_HEIGHT":
			cd.TextHeight, _ = strconv.ParseFloat(l[1], 64)

		case "TEXT_WIDTH":
			cd.TextWidth, _ = strconv.ParseFloat(l[1], 64)

		case "TEXT_PLACING":
			cd.TextPlacing, _ = strconv.Atoi(l[1])

		case "TEXT_DIST_TOP":
			cd.TextDistTop, _ = strconv.ParseFloat(l[1], 64)

		case "TEXT_MAX1":
			cd.TextMax1, _ = strconv.Atoi(l[1])
		case "TEXT_MAX2":
			cd.TextMax2, _ = strconv.Atoi(l[1])
		case "TEXT_HOR1":
			cd.TextHor1, _ = strconv.Atoi(l[1])
		case "TEXT_VER1":
			cd.TextVer1, _ = strconv.Atoi(l[1])
		case "TEXT_HOR2":
			cd.TextHor2, _ = strconv.Atoi(l[1])
		case "TEXT_VER2":
			cd.TextVer2, _ = strconv.Atoi(l[1])
		case "TEXT_HOR3":
			cd.TextHor3, _ = strconv.Atoi(l[1])
		case "TEXT_VER3":
			cd.TextVer3, _ = strconv.Atoi(l[1])

		case "TSCRAP":
			cd.Tscrap, _ = strconv.Atoi(l[1])

		case "NO_OF_PROFS":
			cd.NoOfProfs, _ = strconv.Atoi(l[1])

		case "REST":
			cd.Rest, _ = strconv.Atoi(l[1])

		case "GSD":
			cd.Gsd, _ = strconv.Atoi(l[1])

		case "GSD_DIST":
			cd.GsdDist, _ = strconv.Atoi(l[1])

		case "GSD_ONE_LIM":
			cd.GsdOneLim, _ = strconv.Atoi(l[1])

		default:
			fmt.Println("unknown field in common data:", l[0])

		}
	}
	return cd
}
