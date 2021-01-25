package gen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type GeneralData struct {
	TypeOfGenericFile   string
	TypeOfManufact      string
	BurnerID            string
	NestName            string
	RawName             string
	RawLength           float64
	RawWidth            float64
	RawThickness        float64
	RawArea             float64
	Quality             string
	Density             float64
	SurfTreatTs         string
	SurfTreatOs         string
	RestLength          float64
	SplitDate           string
	ProductionDate      string
	NoOfParts           int
	TotalIdle           float64
	TotalMarking        float64
	TotalBurning        float64
	TotalBlasting       float64
	TotalMarkBlast      float64
	TotalNoOfStarts     int
	NoOfBurningStarts   int
	NoOfMarkingStarts   int
	NoOfBlastingStarts  int
	NoOfMarkBlastStarts int
	StartingPosX        float64
	StartingPosY        float64
	QuantityNormal      int
	QuantityMirrored    int
	ToolID              int
	ProcessID           int
	Sequence            string
	MaterialCode        string
	PurchaseInfo        string
	VerificationDate    string
	PlateAlignment      string
}

func readGeneralData(s *bufio.Scanner) *GeneralData {
	g := &GeneralData{}
	for s.Scan() {
		l := strings.SplitN(s.Text(), "=", 2)

		switch l[0] {
		case "END_OF_GENERAL_DATA":
			return g

		case "TYPE_OF_GENERIC_FILE":
			g.TypeOfGenericFile = l[1]
		case "TYPE_OF_MANUFACT":
			g.TypeOfManufact = l[1]
		case "BURNER_ID":
			g.BurnerID = l[1]
		case "NEST_NAME":
			g.NestName = l[1]
		case "RAW_NAME":
			g.RawName = l[1]
		case "RAW_LENGTH":
			g.RawLength, _ = strconv.ParseFloat(l[1], 64)
		case "RAW_WIDTH":
			g.RawWidth, _ = strconv.ParseFloat(l[1], 64)
		case "RAW_THICKNESS":
			g.RawThickness, _ = strconv.ParseFloat(l[1], 64)
		case "RAW_AREA":
			g.RawArea, _ = strconv.ParseFloat(l[1], 64)
		case "QUALITY":
			g.Quality = l[1]
		case "DENSITY":
			g.Density, _ = strconv.ParseFloat(l[1], 64)
		case "SURF_TREAT_TS":
			g.SurfTreatTs = l[1]
		case "SURF_TREAT_OS":
			g.SurfTreatOs = l[1]
		case "REST_LENGTH":
			g.RestLength, _ = strconv.ParseFloat(l[1], 64)
		case "SPLIT_DATE":
			g.SplitDate = l[1]
		case "PRODUCTION_DATE":
			g.ProductionDate = l[1]
		case "NO_OF_PARTS":
			g.NoOfParts, _ = strconv.Atoi(l[1])
		case "TOTAL_IDLE":
			g.TotalIdle, _ = strconv.ParseFloat(l[1], 64)
		case "TOTAL_MARKING":
			g.TotalMarking, _ = strconv.ParseFloat(l[1], 64)
		case "TOTAL_BURNING":
			g.TotalBurning, _ = strconv.ParseFloat(l[1], 64)
		case "TOTAL_BLASTING":
			g.TotalBlasting, _ = strconv.ParseFloat(l[1], 64)
		case "TOTAL_MARK_BLAST":
			g.TotalMarkBlast, _ = strconv.ParseFloat(l[1], 64)
		case "TOTAL_NO_OF_STARTS":
			g.TotalNoOfStarts, _ = strconv.Atoi(l[1])
		case "NO_OF_BURNING_STARTS":
			g.NoOfBurningStarts, _ = strconv.Atoi(l[1])
		case "NO_OF_MARKING_STARTS":
			g.NoOfMarkingStarts, _ = strconv.Atoi(l[1])
		case "NO_OF_BLASTING_STARTS":
			g.NoOfBlastingStarts, _ = strconv.Atoi(l[1])
		case "NO_OF_MARK_BLAST_STARTS":
			g.NoOfMarkBlastStarts, _ = strconv.Atoi(l[1])
		case "STARTING_POS_X":
			g.StartingPosX, _ = strconv.ParseFloat(l[1], 64)
		case "STARTING_POS_Y":
			g.StartingPosY, _ = strconv.ParseFloat(l[1], 64)
		case "QUANTITY_NORMAL":
			g.QuantityNormal, _ = strconv.Atoi(l[1])
		case "QUANTITY_MIRRORED":
			g.QuantityMirrored, _ = strconv.Atoi(l[1])
		case "TOOL_ID":
			g.ToolID, _ = strconv.Atoi(l[1])
		case "PROCESS_ID":
			g.ProcessID, _ = strconv.Atoi(l[1])
		case "SEQUENCE":
			g.Sequence = l[1]
		case "MATERIAL_CODE":
			g.MaterialCode = l[1]
		case "PURCHASE_INFO":
			g.PurchaseInfo = l[1]
		case "VERIFICATION_DATE":
			g.VerificationDate = l[1]
		case "PLATE_ALIGNMENT":
			g.PlateAlignment = l[1]

		default:
			fmt.Println("unknown field in general data:", l[0])
		}

	}
	return g
}
