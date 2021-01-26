package gen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type End struct {
	LeftOrigin               float64
	LeftClosestPoint         float64
	LeftFarthestPoint        float64
	LeftV0                   float64
	LeftClosestPointWeb      float64
	LeftFarthestPointWeb     float64
	BurnEcut                 int
	Stoss                    int
	FabricationExcess        float64
	EndcutType               int
	EndcutMask               string
	EndcutCode               int
	A                        float64
	B                        float64
	C                        float64
	R1                       float64
	R2                       float64
	V1                       float64
	V2                       float64
	V3                       float64
	V4                       float64
	Ks                       float64
	EndcutUserDefinedName    string
	EndcutUserDefinedParam1  int
	EndcutUserDefinedParam2  int
	EndcutUserDefinedParam3  int
	EndcutUserDefinedParam4  int
	EndcutUserDefinedParam5  int
	EndcutUserDefinedParam6  int
	EndcutUserDefinedParam7  int
	EndcutUserDefinedParam8  int
	EndcutUserDefinedParam9  int
	EndcutUserDefinedParam10 int
	EndcutUserDefinedParam11 int
	EndcutUserDefinedParam12 int
	EndcutUserDefinedParam13 int
	EndcutUserDefinedParam14 int
	EndcutUserDefinedParam15 int
	EndcutUserDefinedParam16 int
	EndcutUserDefinedParam17 int
	EndcutUserDefinedParam18 int
	EndcutUserDefinedParam19 int
	EndcutUserDefinedParam20 int
	EndcutUserDefinedData1   string
	EndcutUserDefinedData2   string
	EndcutUserDefinedData3   string
	EndcutUserDefinedData4   string
	EndcutUserDefinedData5   string
	EndcutUserDefinedData6   string
	EndcutUserDefinedData7   string
	EndcutUserDefinedData8   string
	EndcutUserDefinedData9   string
	EndcutUserDefinedData10  string
	EndcutUserDefinedData11  string
	EndcutUserDefinedData12  string
	EndcutUserDefinedData13  string
	EndcutUserDefinedData14  string
	EndcutUserDefinedData15  string
	EndcutUserDefinedData16  string
	EndcutUserDefinedData17  string
	EndcutUserDefinedData18  string
	EndcutUserDefinedData19  string
	EndcutUserDefinedData20  string
	Excess                   float64
	BevelDefined             int
	BevelCode                int
	BevelName                string
	BevelType                int
	BevelVariant             int
	E                        float64
	Gap                      float64
	Chamfer                  float64
	Alpha                    float64
	Beta                     float64
	Nose                     float64
	H                        float64
	HFact                    float64
	HFactAdjust              float64
	AngleTs                  float64
	AngleOs                  float64
	DepthTs                  float64
	DepthOs                  float64
	ChamferWidthTs           float64
	ChamferWidthOs           float64
	ChamferHeightTs          float64
	ChamferHeightOs          float64
	ConnectionAngle          float64
	WebSeg                   float64
	BevelDefinedFlange       float64
	BevelCodeFlange          float64
	BevelNameFlange          string
	BevelTypeFlange          float64
	BevelVariantFlange       float64
	EFlange                  float64
	GapFlange                float64
	ChamferFlange            float64
	AlphaFlange              float64
	BetaFlange               float64
	NoseFlange               float64
	HFlange                  float64
	HFactFlange              float64
	HFactAdjustFlange        float64
	AngleTsFlange            float64
	AngleOsFlange            float64
	DepthTsFlange            float64
	DepthOsFlange            float64
	ChamferWidthTsFlange     float64
	ChamferWidthOsFlange     float64
	ChamferHeightTsFlange    float64
	ChamferHeightOsFlange    float64
	ConnectionAngleFlange    float64
	FlaSeg                   float64
	BevelDefinedFlange2      float64
	BevelCodeFlange2         float64
	BevelNameFlange2         string
	BevelTypeFlange2         float64
	BevelVariantFlange2      float64
	EFlange2                 float64
	GapFlange2               float64
	ChamferFlange2           float64
	AlphaFlange2             float64
	BetaFlange2              float64
	NoseFlange2              float64
	HFlange2                 float64
	HFactFlange2             float64
	HFactAdjustFlange2       float64
	AngleTsFlange2           float64
	AngleOsFlange2           float64
	DepthTsFlange2           float64
	DepthOsFlange2           float64
	ChamferWidthTsFlange2    float64
	ChamferWidthOsFlange2    float64
	ChamferHeightTsFlange2   float64
	ChamferHeightOsFlange2   float64
	Fla2Seg                  float64
	Gsd                      float64
	GsdDist                  float64
	Contour                  *Contour
	FContour                 *Contour
}

func readEnd(s *bufio.Scanner) *End {
	e := &End{}

	bevelNamesFound := 0

	for s.Scan() {
		l := strings.SplitN(s.Text(), "=", 2)

		switch l[0] {
		case "END_OF_LEFT_END", "END_OF_RIGHT_END":
			return e

		case "START_OF_CONTOUR":
			e.Contour = readContour(s)

		case "START_OF_FCONTOUR":
			e.FContour = readContour(s)

		case "LEFT_ORIGIN":
			e.LeftOrigin, _ = strconv.ParseFloat(l[1], 64)

		case "LEFT_CLOSEST_POINT":
			e.LeftClosestPoint, _ = strconv.ParseFloat(l[1], 64)

		case "LEFT_FARTHEST_POINT":
			e.LeftFarthestPoint, _ = strconv.ParseFloat(l[1], 64)

		case "LEFT_V0":
			e.LeftV0, _ = strconv.ParseFloat(l[1], 64)

		case "LEFT_CLOSEST_POINT_WEB":
			e.LeftClosestPointWeb, _ = strconv.ParseFloat(l[1], 64)

		case "LEFT_FARTHEST_POINT_WEB":
			e.LeftFarthestPointWeb, _ = strconv.ParseFloat(l[1], 64)

		case "BURN_ECUT":
			e.BurnEcut, _ = strconv.Atoi(l[1])

		case "STOSS":
			e.Stoss, _ = strconv.Atoi(l[1])

		case "FABRICATION_EXCESS":
			e.FabricationExcess, _ = strconv.ParseFloat(l[1], 64)

		case "ENDCUT_TYPE":
			e.EndcutType, _ = strconv.Atoi(l[1])

		case "ENDCUT_MASK":
			e.EndcutMask = l[1]

		case "ENDCUT_CODE":
			e.EndcutCode, _ = strconv.Atoi(l[1])

		case "A":
			e.A, _ = strconv.ParseFloat(l[1], 64)
		case "B":
			e.B, _ = strconv.ParseFloat(l[1], 64)
		case "C":
			e.C, _ = strconv.ParseFloat(l[1], 64)
		case "R1":
			e.R1, _ = strconv.ParseFloat(l[1], 64)
		case "R2":
			e.R2, _ = strconv.ParseFloat(l[1], 64)
		case "V1":
			e.V1, _ = strconv.ParseFloat(l[1], 64)
		case "V2":
			e.V2, _ = strconv.ParseFloat(l[1], 64)
		case "V3":
			e.V3, _ = strconv.ParseFloat(l[1], 64)
		case "V4":
			e.V4, _ = strconv.ParseFloat(l[1], 64)
		case "KS":
			e.Ks, _ = strconv.ParseFloat(l[1], 64)

		case "ENDCUT_USER_DEFINED_NAME":
			e.EndcutUserDefinedName = l[1]
		case "ENDCUT_USER_DEFINED_PARAM1":
			e.EndcutUserDefinedParam1, _ = strconv.Atoi(l[1])
		case "ENDCUT_USER_DEFINED_PARAM2":
			e.EndcutUserDefinedParam2, _ = strconv.Atoi(l[1])
		case "ENDCUT_USER_DEFINED_PARAM3":
			e.EndcutUserDefinedParam3, _ = strconv.Atoi(l[1])
		case "ENDCUT_USER_DEFINED_PARAM4":
			e.EndcutUserDefinedParam4, _ = strconv.Atoi(l[1])
		case "ENDCUT_USER_DEFINED_PARAM5":
			e.EndcutUserDefinedParam5, _ = strconv.Atoi(l[1])
		case "ENDCUT_USER_DEFINED_PARAM6":
			e.EndcutUserDefinedParam6, _ = strconv.Atoi(l[1])
		case "ENDCUT_USER_DEFINED_PARAM7":
			e.EndcutUserDefinedParam7, _ = strconv.Atoi(l[1])
		case "ENDCUT_USER_DEFINED_PARAM8":
			e.EndcutUserDefinedParam8, _ = strconv.Atoi(l[1])
		case "ENDCUT_USER_DEFINED_PARAM9":
			e.EndcutUserDefinedParam9, _ = strconv.Atoi(l[1])
		case "ENDCUT_USER_DEFINED_PARAM10":
			e.EndcutUserDefinedParam10, _ = strconv.Atoi(l[1])
		case "ENDCUT_USER_DEFINED_PARAM11":
			e.EndcutUserDefinedParam11, _ = strconv.Atoi(l[1])
		case "ENDCUT_USER_DEFINED_PARAM12":
			e.EndcutUserDefinedParam12, _ = strconv.Atoi(l[1])
		case "ENDCUT_USER_DEFINED_PARAM13":
			e.EndcutUserDefinedParam13, _ = strconv.Atoi(l[1])
		case "ENDCUT_USER_DEFINED_PARAM14":
			e.EndcutUserDefinedParam14, _ = strconv.Atoi(l[1])
		case "ENDCUT_USER_DEFINED_PARAM15":
			e.EndcutUserDefinedParam15, _ = strconv.Atoi(l[1])
		case "ENDCUT_USER_DEFINED_PARAM16":
			e.EndcutUserDefinedParam16, _ = strconv.Atoi(l[1])
		case "ENDCUT_USER_DEFINED_PARAM17":
			e.EndcutUserDefinedParam17, _ = strconv.Atoi(l[1])
		case "ENDCUT_USER_DEFINED_PARAM18":
			e.EndcutUserDefinedParam18, _ = strconv.Atoi(l[1])
		case "ENDCUT_USER_DEFINED_PARAM19":
			e.EndcutUserDefinedParam19, _ = strconv.Atoi(l[1])
		case "ENDCUT_USER_DEFINED_PARAM20":
			e.EndcutUserDefinedParam20, _ = strconv.Atoi(l[1])

		case "ENDCUT_USER_DEFINED_DATA1":
			e.EndcutUserDefinedData1 = l[1]
		case "ENDCUT_USER_DEFINED_DATA2":
			e.EndcutUserDefinedData2 = l[1]
		case "ENDCUT_USER_DEFINED_DATA3":
			e.EndcutUserDefinedData3 = l[1]
		case "ENDCUT_USER_DEFINED_DATA4":
			e.EndcutUserDefinedData4 = l[1]
		case "ENDCUT_USER_DEFINED_DATA5":
			e.EndcutUserDefinedData5 = l[1]
		case "ENDCUT_USER_DEFINED_DATA6":
			e.EndcutUserDefinedData6 = l[1]
		case "ENDCUT_USER_DEFINED_DATA7":
			e.EndcutUserDefinedData7 = l[1]
		case "ENDCUT_USER_DEFINED_DATA8":
			e.EndcutUserDefinedData8 = l[1]
		case "ENDCUT_USER_DEFINED_DATA9":
			e.EndcutUserDefinedData9 = l[1]
		case "ENDCUT_USER_DEFINED_DATA10":
			e.EndcutUserDefinedData10 = l[1]
		case "ENDCUT_USER_DEFINED_DATA11":
			e.EndcutUserDefinedData11 = l[1]
		case "ENDCUT_USER_DEFINED_DATA12":
			e.EndcutUserDefinedData12 = l[1]
		case "ENDCUT_USER_DEFINED_DATA13":
			e.EndcutUserDefinedData13 = l[1]
		case "ENDCUT_USER_DEFINED_DATA14":
			e.EndcutUserDefinedData14 = l[1]
		case "ENDCUT_USER_DEFINED_DATA15":
			e.EndcutUserDefinedData15 = l[1]
		case "ENDCUT_USER_DEFINED_DATA16":
			e.EndcutUserDefinedData16 = l[1]
		case "ENDCUT_USER_DEFINED_DATA17":
			e.EndcutUserDefinedData17 = l[1]
		case "ENDCUT_USER_DEFINED_DATA18":
			e.EndcutUserDefinedData18 = l[1]
		case "ENDCUT_USER_DEFINED_DATA19":
			e.EndcutUserDefinedData19 = l[1]
		case "ENDCUT_USER_DEFINED_DATA20":
			e.EndcutUserDefinedData20 = l[1]

		case "EXCESS":
			e.Excess, _ = strconv.ParseFloat(l[1], 64)

		case "BEVEL_DEFINED":
			e.BevelDefined, _ = strconv.Atoi(l[1])

		case "BEVEL_CODE":
			e.BevelCode, _ = strconv.Atoi(l[1])

		case "BEVEL_TYPE":
			e.BevelType, _ = strconv.Atoi(l[1])

		case "BEVEL_VARIANT":
			e.BevelVariant, _ = strconv.Atoi(l[1])

		case "E":
			e.E, _ = strconv.ParseFloat(l[1], 64)

		case "GAP":
			e.Gap, _ = strconv.ParseFloat(l[1], 64)

		case "CHAMFER":
			e.Chamfer, _ = strconv.ParseFloat(l[1], 64)

		case "ALPHA":
			e.Alpha, _ = strconv.ParseFloat(l[1], 64)

		case "BETA":
			e.Beta, _ = strconv.ParseFloat(l[1], 64)

		case "NOSE":
			e.Nose, _ = strconv.ParseFloat(l[1], 64)

		case "H":
			e.H, _ = strconv.ParseFloat(l[1], 64)

		case "H_FACT":
			e.HFact, _ = strconv.ParseFloat(l[1], 64)

		case "H_FACT_ADJUST":
			e.HFactAdjust, _ = strconv.ParseFloat(l[1], 64)

		case "ANGLE_TS":
			e.AngleTs, _ = strconv.ParseFloat(l[1], 64)

		case "ANGLE_OS":
			e.AngleOs, _ = strconv.ParseFloat(l[1], 64)

		case "DEPTH_TS":
			e.DepthTs, _ = strconv.ParseFloat(l[1], 64)

		case "DEPTH_OS":
			e.DepthOs, _ = strconv.ParseFloat(l[1], 64)

		case "CHAMFER_WIDTH_TS":
			e.ChamferWidthTs, _ = strconv.ParseFloat(l[1], 64)

		case "CHAMFER_WIDTH_OS":
			e.ChamferWidthOs, _ = strconv.ParseFloat(l[1], 64)

		case "CHAMFER_HEIGHT_TS":
			e.ChamferHeightTs, _ = strconv.ParseFloat(l[1], 64)

		case "CHAMFER_HEIGHT_OS":
			e.ChamferHeightOs, _ = strconv.ParseFloat(l[1], 64)

		case "CONNECTION_ANGLE":
			e.ConnectionAngle, _ = strconv.ParseFloat(l[1], 64)

		case "WEB_SEG":
			e.WebSeg, _ = strconv.ParseFloat(l[1], 64)

		case "BEVEL_DEFINED_FLANGE":
			e.BevelDefinedFlange, _ = strconv.ParseFloat(l[1], 64)

		case "BEVEL_CODE_FLANGE":
			e.BevelCodeFlange, _ = strconv.ParseFloat(l[1], 64)

		case "BEVEL_NAME":
			switch bevelNamesFound {
			case 0:
				e.BevelName = l[1]
			case 1:
				e.BevelNameFlange = l[1]
			case 2:
				e.BevelNameFlange2 = l[1]
			}

			bevelNamesFound++

		case "BEVEL_TYPE_FLANGE":
			e.BevelTypeFlange, _ = strconv.ParseFloat(l[1], 64)

		case "BEVEL_VARIANT_FLANGE":
			e.BevelVariantFlange, _ = strconv.ParseFloat(l[1], 64)

		case "E_FLANGE":
			e.EFlange, _ = strconv.ParseFloat(l[1], 64)

		case "GAP_FLANGE":
			e.GapFlange, _ = strconv.ParseFloat(l[1], 64)

		case "CHAMFER_FLANGE":
			e.ChamferFlange, _ = strconv.ParseFloat(l[1], 64)

		case "ALPHA_FLANGE":
			e.AlphaFlange, _ = strconv.ParseFloat(l[1], 64)

		case "BETA_FLANGE":
			e.BetaFlange, _ = strconv.ParseFloat(l[1], 64)

		case "NOSE_FLANGE":
			e.NoseFlange, _ = strconv.ParseFloat(l[1], 64)

		case "H_FLANGE":
			e.HFlange, _ = strconv.ParseFloat(l[1], 64)

		case "H_FACT_FLANGE":
			e.HFactFlange, _ = strconv.ParseFloat(l[1], 64)

		case "H_FACT_ADJUST_FLANGE":
			e.HFactAdjustFlange, _ = strconv.ParseFloat(l[1], 64)

		case "ANGLE_TS_FLANGE":
			e.AngleTsFlange, _ = strconv.ParseFloat(l[1], 64)

		case "ANGLE_OS_FLANGE":
			e.AngleOsFlange, _ = strconv.ParseFloat(l[1], 64)

		case "DEPTH_TS_FLANGE":
			e.DepthTsFlange, _ = strconv.ParseFloat(l[1], 64)

		case "DEPTH_OS_FLANGE":
			e.DepthOsFlange, _ = strconv.ParseFloat(l[1], 64)

		case "CHAMFER_WIDTH_TS_FLANGE":
			e.ChamferWidthTsFlange, _ = strconv.ParseFloat(l[1], 64)

		case "CHAMFER_WIDTH_OS_FLANGE":
			e.ChamferWidthOsFlange, _ = strconv.ParseFloat(l[1], 64)

		case "CHAMFER_HEIGHT_TS_FLANGE":
			e.ChamferHeightTsFlange, _ = strconv.ParseFloat(l[1], 64)

		case "CHAMFER_HEIGHT_OS_FLANGE":
			e.ChamferHeightOsFlange, _ = strconv.ParseFloat(l[1], 64)

		case "CONNECTION_ANGLE_FLANGE":
			e.ConnectionAngleFlange, _ = strconv.ParseFloat(l[1], 64)

		case "FLA_SEG":
			e.FlaSeg, _ = strconv.ParseFloat(l[1], 64)

		case "BEVEL_DEFINED_FLANGE2":
			e.BevelDefinedFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "BEVEL_CODE_FLANGE2":
			e.BevelCodeFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "BEVEL_TYPE_FLANGE2":
			e.BevelTypeFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "BEVEL_VARIANT_FLANGE2":
			e.BevelVariantFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "E_FLANGE2":
			e.EFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "GAP_FLANGE2":
			e.GapFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "CHAMFER_FLANGE2":
			e.ChamferFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "ALPHA_FLANGE2":
			e.AlphaFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "BETA_FLANGE2":
			e.BetaFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "NOSE_FLANGE2":
			e.NoseFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "H_FLANGE2":
			e.HFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "H_FACT_FLANGE2":
			e.HFactFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "H_FACT_ADJUST_FLANGE2":
			e.HFactAdjustFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "ANGLE_TS_FLANGE2":
			e.AngleTsFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "ANGLE_OS_FLANGE2":
			e.AngleOsFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "DEPTH_TS_FLANGE2":
			e.DepthTsFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "DEPTH_OS_FLANGE2":
			e.DepthOsFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "CHAMFER_WIDTH_TS_FLANGE2":
			e.ChamferWidthTsFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "CHAMFER_WIDTH_OS_FLANGE2":
			e.ChamferWidthOsFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "CHAMFER_HEIGHT_TS_FLANGE2":
			e.ChamferHeightTsFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "CHAMFER_HEIGHT_OS_FLANGE2":
			e.ChamferHeightOsFlange2, _ = strconv.ParseFloat(l[1], 64)

		case "FLA2_SEG":
			e.Fla2Seg, _ = strconv.ParseFloat(l[1], 64)

		case "GSD":
			e.Gsd, _ = strconv.ParseFloat(l[1], 64)

		case "GSD_DIST":
			e.GsdDist, _ = strconv.ParseFloat(l[1], 64)

		default:
			fmt.Println("unknown field in end data:", l[0])
		}
	}
	return e
}
