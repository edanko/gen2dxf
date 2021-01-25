package gen

import (
	"bufio"
	"fmt"
	"strings"
)

type PartInformation struct {
	PartName      string
	TypeOfWork    string
	IdleData      []*IdleData
	MarkingData   []*MarkingData
	BurningData   []*BurningData
	LabeltextData []*LabeltextData
}

func readPartInformation(s *bufio.Scanner) *PartInformation {
	pi := &PartInformation{}
	for s.Scan() {
		l := strings.SplitN(s.Text(), "=", 2)

		switch l[0] {

		case "PART_NAME":
			pi.PartName = l[1]
		case "TYPE_OF_WORK":
			pi.TypeOfWork = l[1]

		case "END_OF_PART_INFORMATION":
			return pi

		default:
			fmt.Println("unknown part information section", l[0])
		}
	}
	return pi
}
