package gen

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

type Contour struct {
	NoOfSeg int
	StartU  float64
	StartV  float64
	//BevelData    *BevelData
	ExcessData   *ExcessData
	Segments     []*Segment
	AuxilaryData []*AuxilaryData
}

// for profile segment used only amp*, sweep, u and v
// Length used only in connection trace
type Segment struct {
	AmpU      float64
	AmpV      float64
	Amp       float64
	Radius    float64
	Sweep     float64
	Length    float64
	OriginU   float64
	OriginV   float64
	StartU    float64
	StartV    float64
	U         float64
	V         float64
	BevelData *BevelData
}

func readContour(s *bufio.Scanner) *Contour {
	con := &Contour{}

	var currentBevel *BevelData

	for s.Scan() {
		l := strings.Split(s.Text(), "=")

		switch l[0] {
		case "END_OF_CONTOUR", "END_OF_FCONTOUR", "END_OF_FCONTOUR2":
			return con

		case "BEVEL_DATA":
			currentBevel = readBevelData(s)
			//con.BevelData = readBevelData(s)

		case "EXCESS_DATA":
			con.ExcessData = readExcessData(s)

		case "AUXILIARY_DATA":
			con.AuxilaryData = append(con.AuxilaryData, readAuxilaryData(s))

		case "NO_OF_SEG":
			con.NoOfSeg, _ = strconv.Atoi(l[1])

		case "START_U":
			con.StartU, _ = strconv.ParseFloat(l[1], 64)

		case "START_V":
			con.StartV, _ = strconv.ParseFloat(l[1], 64)

		default:
			seg := &Segment{}

			if l[0] != "AMP_U" {
				log.Fatalf("something wrong with contour: first element (%s) != AMP_U\n", l[0])
			}
			seg.AmpU, _ = strconv.ParseFloat(l[1], 64)
			seg = readRestOfSegment(s, seg)

			if len(con.Segments) == 0 {
				seg.StartU = con.StartU
				seg.StartV = con.StartV
			} else {
				seg.StartU = con.Segments[len(con.Segments)-1].U
				seg.StartV = con.Segments[len(con.Segments)-1].V
			}

			seg.BevelData = currentBevel
			con.Segments = append(con.Segments, seg)
		}

	}
	return con
}

func readRestOfSegment(s *bufio.Scanner, seg *Segment) *Segment {
	for s.Scan() {

		l := strings.Split(s.Text(), "=")

		switch l[0] {
		case "AMP_V":
			seg.AmpV, _ = strconv.ParseFloat(l[1], 64)

		case "AMP":
			seg.Amp, _ = strconv.ParseFloat(l[1], 64)

		case "RADIUS":
			seg.Radius, _ = strconv.ParseFloat(l[1], 64)

		case "SWEEP":
			seg.Sweep, _ = strconv.ParseFloat(l[1], 64)

		case "LENGTH":
			seg.Length, _ = strconv.ParseFloat(l[1], 64)

		case "ORIGIN_U":
			seg.OriginU, _ = strconv.ParseFloat(l[1], 64)

		case "ORIGIN_V":
			seg.OriginV, _ = strconv.ParseFloat(l[1], 64)

		case "U":
			seg.U, _ = strconv.ParseFloat(l[1], 64)

		case "V":
			seg.V, _ = strconv.ParseFloat(l[1], 64)
			return seg
		}

	}
	return seg
}
