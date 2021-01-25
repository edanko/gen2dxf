package gen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type PartData struct {
	Name                string
	PosNo               string
	PartnameLong        string
	PartnameLongSb      string
	ShipNo              string
	AssemblyLow         string
	AssemblyLowSb       string
	AssemblyHigh        string
	AssemblyHighSb      string
	BlockNo             string
	PartArea            float64
	Mirrored            int
	ClipGrindingSide    string
	CustData1           string
	CustData2           string
	CustData3           string
	CustData4           string
	CustText1           string
	CustText2           string
	CustText3           string
	CustText4           string
	CustText5           string
	CustText6           string
	CustText7           string
	CustText8           string
	CustText9           string
	CustText10          string
	WorkingLocation     string
	BuildingStrategy    string
	PlanningUnit        string
	ExtensionU          float64
	ExtensionV          float64
	PlateSide           string
	PanelName           string
	PartCogU            float64
	PartCogV            float64
	NestedOn            string
	Knuckled            string
	RawName             string
	NoIntervalsExcess1  int
	LengthExcess1       float64
	NoIntervalsExcess2  int
	LengthExcess2       float64
	NoIntervalsExcess3  int
	LengthExcess3       float64
	TransformationData  *TransformationData
	Destination         string
	Comment             string
	FunctionalDescrCode int
	FunctionalDescrText string
	UniversalID         string
	StringData          []*StringData
	IdleData            []*IdleData
	BurningData         []*BurningData
	MarkingData         []*MarkingData
	GeometryData        []*GeometryData
	LabeltextData       []*LabeltextData
	BumpData            []*BumpData
	EdgeData            []*EdgeData
}

func readPartData(s *bufio.Scanner) *PartData {
	p := &PartData{}
	for s.Scan() {
		l := strings.SplitN(s.Text(), "=", 2)

		switch l[0] {
		case "END_OF_PART_DATA":
			return p

		case "TRANSFORMATION_DATA":
			p.TransformationData = readTransformationData(s)

		case "NAME":
			p.Name = l[1]
		case "POSNO":
			p.PosNo = l[1]
		case "PARTNAME_LONG":
			p.PartnameLong = l[1]
		case "PARTNAME_LONG_SB":
			p.PartnameLongSb = l[1]
		case "SHIP_NO":
			p.ShipNo = l[1]
		case "ASSEMBLY_LOW":
			p.AssemblyLow = l[1]
		case "ASSEMBLY_LOW_SB":
			p.AssemblyLowSb = l[1]
		case "ASSEMBLY_HIGH":
			p.AssemblyHigh = l[1]
		case "ASSEMBLY_HIGH_SB":
			p.AssemblyHighSb = l[1]
		case "BLOCK_NO":
			p.BlockNo = l[1]
		case "PART_AREA":
			p.PartArea, _ = strconv.ParseFloat(l[1], 64)
		case "MIRRORED":
			p.Mirrored, _ = strconv.Atoi(l[1])
		case "CLIP_GRINDING_SIDE":
			p.ClipGrindingSide = l[1]
		case "CUST_DATA_1":
			p.CustData1 = l[1]
		case "CUST_DATA_2":
			p.CustData2 = l[1]
		case "CUST_DATA_3":
			p.CustData3 = l[1]
		case "CUST_DATA_4":
			p.CustData4 = l[1]
		case "CUST_TEXT_1":
			p.CustText1 = l[1]
		case "CUST_TEXT_2":
			p.CustText2 = l[1]
		case "CUST_TEXT_3":
			p.CustText3 = l[1]
		case "CUST_TEXT_4":
			p.CustText4 = l[1]
		case "CUST_TEXT_5":
			p.CustText5 = l[1]
		case "CUST_TEXT_6":
			p.CustText6 = l[1]
		case "CUST_TEXT_7":
			p.CustText7 = l[1]
		case "CUST_TEXT_8":
			p.CustText8 = l[1]
		case "CUST_TEXT_9":
			p.CustText9 = l[1]
		case "CUST_TEXT_10":
			p.CustText10 = l[1]
		case "WORKING_LOCATION":
			p.WorkingLocation = l[1]
		case "BUILDING_STRATEGY":
			p.BuildingStrategy = l[1]
		case "PLANNING_UNIT":
			p.PlanningUnit = l[1]
		case "EXTENSION_U":
			p.ExtensionU, _ = strconv.ParseFloat(l[1], 64)
		case "EXTENSION_V":
			p.ExtensionV, _ = strconv.ParseFloat(l[1], 64)
		case "PLATE_SIDE":
			p.PlateSide = l[1]
		case "PANEL_NAME":
			p.PanelName = l[1]
		case "PART_COG_U":
			p.PartCogU, _ = strconv.ParseFloat(l[1], 64)
		case "PART_COG_V":
			p.PartCogV, _ = strconv.ParseFloat(l[1], 64)
		case "NESTED_ON":
			p.NestedOn = l[1]
		case "KNUCKLED":
			p.Knuckled = l[1]
		case "RAW_NAME":
			p.RawName = l[1]
		case "NO_INTERVALS_EXCESS_1":
			p.NoIntervalsExcess1, _ = strconv.Atoi(l[1])
		case "LENGTH_EXCESS_1":
			p.LengthExcess1, _ = strconv.ParseFloat(l[1], 64)
		case "NO_INTERVALS_EXCESS_2":
			p.NoIntervalsExcess2, _ = strconv.Atoi(l[1])
		case "LENGTH_EXCESS_2":
			p.LengthExcess2, _ = strconv.ParseFloat(l[1], 64)
		case "NO_INTERVALS_EXCESS_3":
			p.NoIntervalsExcess3, _ = strconv.Atoi(l[1])
		case "LENGTH_EXCESS_3":
			p.LengthExcess3, _ = strconv.ParseFloat(l[1], 64)
		case "DESTINATION":
			p.Destination = l[1]
		case "COMMENT":
			p.Comment = l[1]
		case "FUNCTIONAL_DESCR_CODE":
			p.FunctionalDescrCode, _ = strconv.Atoi(l[1])
		case "FUNCTIONAL_DESCR_TEXT":
			p.FunctionalDescrText = l[1]
		case "UNIVERSAL_ID":
			p.UniversalID = l[1]

		default:
			fmt.Println("unknown field in part data:", l[0])

		}
	}
	return p
}
