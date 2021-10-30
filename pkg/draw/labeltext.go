package draw

import (
	"fmt"

	"gen2dxf/pkg/gen"

	"github.com/edanko/dxf/color"
	"github.com/edanko/dxf/drawing"
)

func addLabeltextData(d *drawing.Drawing, lts []*gen.LabeltextData) {
	if lts == nil {
		return
	}
	for _, lt := range lts {
		addLabeltext(d, lt)
	}
}

func addLabeltext(d *drawing.Drawing, lt *gen.LabeltextData) {
	switch lt.TextType {
	case "PANEL_NAME", "PART_ID", "THICKNESS", "DIMENSIONS", "POSNO", "SIDE", "BEVEL", "EXCESS", "BEND_RADIUS":

	default:
		fmt.Println("labeltext type:", lt.TextType)
		drawText(d, lt.TextPositionU, lt.TextPositionV, lt.TextHeight, lt.TextAngle, lt.Text, color.Red)
	}
}
