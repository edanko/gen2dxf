package gen

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type PlateGen struct {
	GeneralData *GeneralData
	PartData    []*PartData
}

func ParsePlateFile(fname string) *PlateGen {

	fmt.Println("[i] processing", filepath.Base(fname))

	f, err := os.Open(fname)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	return readPlate(f)
}

func readPlate(r io.Reader) *PlateGen {
	g := &PlateGen{}
	s := bufio.NewScanner(r)

	var pi *PartInformation

	addPartInfo := func() {
		if pi == nil {
			return
		}
		for _, p := range g.PartData {
			if p.Name == pi.PartName {

				if pi.IdleData != nil {
					p.IdleData = append(p.IdleData, pi.IdleData...)
				}
				if pi.MarkingData != nil {
					p.MarkingData = append(p.MarkingData, pi.MarkingData...)
				}
				if pi.BurningData != nil {
					p.BurningData = append(p.BurningData, pi.BurningData...)
				}
				if pi.LabeltextData != nil {
					p.LabeltextData = append(p.LabeltextData, pi.LabeltextData...)
				}
				break
			}
		}
		pi = nil
	}

	for s.Scan() {
		t := s.Text()

		switch t {
		case "TYPE_OF_GENERIC_FILE=LISTED_PROFILE":
			fmt.Println("[x] found listed profile - skipping")
			return nil

		case "GENERAL_DATA":
			g.GeneralData = readGeneralData(s)

		case "PART_DATA":
			p := readPartData(s)
			g.PartData = append(g.PartData, p)

		case "PART_INFORMATION":
			if pi != nil {
				addPartInfo()
			}
			pi = readPartInformation(s)

		case "IDLE_DATA":
			i := readIdleData(s)
			pi.IdleData = append(pi.IdleData, i)

		case "MARKING_DATA":
			m := readMarkingData(s)
			if pi == nil {
				g.PartData[len(g.PartData)-1].MarkingData = append(g.PartData[len(g.PartData)-1].MarkingData, m)
			} else {
				pi.MarkingData = append(pi.MarkingData, m)
			}

		case "GEOMETRY_DATA":
			gd := readGeometryData(s)
			g.PartData[len(g.PartData)-1].GeometryData = append(g.PartData[len(g.PartData)-1].GeometryData, gd)
			if pi != nil {
				fmt.Println("found geometry data after part information, check this shit")
			}

		case "STRING_DATA":
			sd := readStringData(s)
			g.PartData[len(g.PartData)-1].StringData = append(g.PartData[len(g.PartData)-1].StringData, sd)
			if pi != nil {
				fmt.Println("found string data after part information, check this shit")
			}

		case "BUMP_DATA":
			b := readBumpData(s)
			g.PartData[len(g.PartData)-1].BumpData = append(g.PartData[len(g.PartData)-1].BumpData, b)
			if pi != nil {
				fmt.Println("found bump data after part information, check this shit")
			}

		case "LABELTEXT_DATA":
			l := readLabeltextData(s)
			if pi == nil {
				g.PartData[len(g.PartData)-1].LabeltextData = append(g.PartData[len(g.PartData)-1].LabeltextData, l)
			} else {
				pi.LabeltextData = append(pi.LabeltextData, l)
			}

		case "BURNING_DATA":
			b := readBurningData(s)
			if pi == nil {
				g.PartData[len(g.PartData)-1].BurningData = append(g.PartData[len(g.PartData)-1].BurningData, b)
			} else {
				pi.BurningData = append(pi.BurningData, b)
			}

		case "EDGE_DATA":
			e := readEdgeData(s)
			g.PartData[len(g.PartData)-1].EdgeData = append(g.PartData[len(g.PartData)-1].EdgeData, e)
			if pi != nil {
				fmt.Println("found edge data after part information, check this shit")
			}

		default:
			fmt.Println("unknown plate section:", t)

		}
	}

	addPartInfo()

	return g
}

func (g *PlateGen) Mirror() {
	for _, p := range g.PartData {
		p.Mirror()
	}
}
