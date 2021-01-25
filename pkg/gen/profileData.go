package gen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type ProfileData struct {
	Name                string
	GeneratedAs         string
	PosNo               string
	Mlength             float64
	MlengthManual       float64
	Tlength             float64
	TlengthManual       float64
	EndPointDist        float64
	EndPointDistProd    float64
	TraceLength         float64
	TraceLengthProd     float64
	TwistAngle          float64
	ApproxPartWeight    float64
	IdentString         string
	PartnameLong        string
	PartnameShort       string
	ShipNo              string
	ShipName            string
	ProfSide            string
	SideUp              string
	Direction           string
	DirectionSign       string
	DirectionTrace      string
	BevelTrace          int
	Grinding            int
	Assembly            string
	AssemblyHigh        string
	AssemblyTot         string
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
	SurfaceTreatment    string
	PlanningUnit        string
	WorkingLocation     string
	BuildingStrategy    string
	AssemblySequence    string
	Baugruppe           string
	Zeichnungsno        string
	Zaehlno             string
	BlockNo             string
	BlockName           string
	NoOfRefLines        int
	Mirror              string
	Form                string
	BendDirection       string
	AdditionalInfo1     string
	AdditionalInfo2     string
	AdditionalInfo3     string
	AdditionalInfo4     string
	AdditionalInfo5     string
	AdditionalInfo6     string
	AdditionalInfo7     string
	AdditionalInfo8     string
	AdditionalInfo9     string
	AdditionalInfo10    string
	AdditionalInfo11    string
	AdditionalInfo12    string
	MarkingLeftExcist   int
	MarkingRightExcist  int
	NoOfMacs            int
	NoOfMarks           int
	LeftEnd             *End
	RightEnd            *End
	HolesNotchesCutouts []*HolesNotchesCutouts
	ConnectionTrace     *ConnectionTrace
	GeometryData        []*GeometryData
	StringData          []*StringData
}

func readProfileData(s *bufio.Scanner) *ProfileData {
	p := &ProfileData{}

	for s.Scan() {
		l := strings.SplitN(s.Text(), "=", 2)

		switch l[0] {
		case "END_OF_PROFILE_DATA":
			return p

		case "NAME":
			p.Name = l[1]

		case "GENERATED_AS":
			p.GeneratedAs = l[1]

		case "POSNO":
			p.PosNo = l[1]

		case "MLENGTH":
			p.Mlength, _ = strconv.ParseFloat(l[1], 64)
		case "MLENGTH_MANUAL":
			p.MlengthManual, _ = strconv.ParseFloat(l[1], 64)
		case "TLENGTH":
			p.Tlength, _ = strconv.ParseFloat(l[1], 64)
		case "TLENGTH_MANUAL":
			p.TlengthManual, _ = strconv.ParseFloat(l[1], 64)

		case "END_POINT_DIST":
			p.EndPointDist, _ = strconv.ParseFloat(l[1], 64)

		case "END_POINT_DIST_PROD":
			p.EndPointDistProd, _ = strconv.ParseFloat(l[1], 64)

		case "TRACE_LENGTH":
			p.TraceLength, _ = strconv.ParseFloat(l[1], 64)

		case "TRACE_LENGTH_PROD":
			p.TraceLengthProd, _ = strconv.ParseFloat(l[1], 64)

		case "TWIST_ANGLE":
			p.TwistAngle, _ = strconv.ParseFloat(l[1], 64)

		case "APPROX_PART_WEIGHT":
			p.ApproxPartWeight, _ = strconv.ParseFloat(l[1], 64)

		case "IDENT_STRING":
			p.IdentString = l[1]

		case "PARTNAME_LONG":
			p.PartnameLong = l[1]
		case "PARTNAME_SHORT":
			p.PartnameShort = l[1]

		case "SHIP_NO":
			p.ShipNo = l[1]
		case "SHIP_NAME":
			p.ShipName = l[1]

		case "PROF_SIDE":
			p.ProfSide = l[1]

		case "SIDE_UP":
			p.SideUp = l[1]

		case "DIRECTION":
			p.Direction = l[1]
		case "DIRECTION_SIGN":
			p.DirectionSign = l[1]
		case "DIRECTION_TRACE":
			p.DirectionTrace = l[1]

		case "BEVEL_TRACE":
			p.BevelTrace, _ = strconv.Atoi(l[1])

		case "GRINDING":
			p.Grinding, _ = strconv.Atoi(l[1])

		case "ASSEMBLY":
			p.Assembly = l[1]
		case "ASSEMBLY_HIGH":
			p.AssemblyHigh = l[1]
		case "ASSEMBLY_TOT":
			p.AssemblyTot = l[1]

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

		case "SURFACE_TREATMENT":
			p.BlockNo = l[1]

		case "PLANNING_UNIT":
			p.PlanningUnit = l[1]

		case "WORKING_LOCATION":
			p.WorkingLocation = l[1]

		case "BUILDING_STRATEGY":
			p.BuildingStrategy = l[1]

		case "BAUGRUPPE":
			p.Baugruppe = l[1]

		case "ZEICHNUNGSNO":
			p.Zeichnungsno = l[1]

		case "ZAEHLNO":
			p.Zaehlno = l[1]

		case "ASSEMBLY_SEQUENCE":
			p.AssemblySequence = l[1]

		case "BLOCK_NO":
			p.BlockNo = l[1]
		case "BLOCK_NAME":
			p.BlockName = l[1]

		case "NO_OF_REF_LINES":
			p.NoOfRefLines, _ = strconv.Atoi(l[1])

		case "MIRROR":
			p.Mirror = l[1]

		case "FORM":
			p.Form = l[1]

		case "BEND_DIRECTION":
			p.BendDirection = l[1]

		case "ADDITIONAL_INFO1":
			p.AdditionalInfo1 = l[1]
		case "ADDITIONAL_INFO2":
			p.AdditionalInfo2 = l[1]
		case "ADDITIONAL_INFO3":
			p.AdditionalInfo3 = l[1]
		case "ADDITIONAL_INFO4":
			p.AdditionalInfo4 = l[1]
		case "ADDITIONAL_INFO5":
			p.AdditionalInfo5 = l[1]
		case "ADDITIONAL_INFO6":
			p.AdditionalInfo6 = l[1]
		case "ADDITIONAL_INFO7":
			p.AdditionalInfo7 = l[1]
		case "ADDITIONAL_INFO8":
			p.AdditionalInfo8 = l[1]
		case "ADDITIONAL_INFO9":
			p.AdditionalInfo9 = l[1]
		case "ADDITIONAL_INFO10":
			p.AdditionalInfo10 = l[1]
		case "ADDITIONAL_INFO11":
			p.AdditionalInfo11 = l[1]
		case "ADDITIONAL_INFO12":
			p.AdditionalInfo12 = l[1]

		case "MARKING_LEFT_EXCIST":
			p.MarkingLeftExcist, _ = strconv.Atoi(l[1])
		case "MARKING_RIGHT_EXCIST":
			p.MarkingRightExcist, _ = strconv.Atoi(l[1])

		case "NO_OF_MACS":
			p.NoOfMacs, _ = strconv.Atoi(l[1])
		case "NO_OF_MARKS":
			p.NoOfMarks, _ = strconv.Atoi(l[1])

		default:
			fmt.Println("unknown field in profile data:", l[0])

		}
	}
	return p
}
