package draw

import (
	"fmt"

	"github.com/edanko/dxf/color"
	"github.com/edanko/dxf/drawing"
	"github.com/edanko/gen2dxf/pkg/gen"
)

const (
	MarkingTSColor     = color.Cyan
	MarkingTextTSColor = color.Red
	MarkingOSColor     = 30
	MarkingTextOSColor = 30
)

func addMarkingData(d *drawing.Drawing, mds []*gen.MarkingData) {
	if mds == nil {
		return
	}
	for _, md := range mds {

		switch md.MarkingType {
		case "STIFFENER", "PANEL", "BRACKET_ON_INTERSEC._PANEL":
			addMarkingWithText(d, md)

		case "BRACKET":
			continue

		case "INCLINATION_ANGLE_LINE", "CUTOUT_CROSSING", "FLANGE_MARKING_LINE":
			addMarking(d, md)

		case "BENDING_MARK", "ADDED_CONTOUR", "FACE_PLATE_ALIGNMENT":
			addMarking(d, md)

		case "FOLDED_FLANGE":
			addMarking(d, md)

		case "ROLL_AXIS":
			addMarkingWithText(d, md)

		case "LONGITUDINAL":
			addMarkingWithText(d, md)

		case "":
			addMarking(d, md)

		default:
			fmt.Println("unprocessed marking type:", md.MarkingType)
		}
	}
}

// add general marking
func addMarking(d *drawing.Drawing, md *gen.MarkingData) {
	switch md.MarkingSide {
	case "TS":
		drawContour(d, md.Contour, MarkingTSColor)

	case "OS":
		drawContour(d, md.Contour, MarkingOSColor)

	case "BOTH":
		drawContour(d, md.Contour, MarkingTSColor)
		drawContour(d, md.Contour, MarkingOSColor)
	}
}

// add general marking
func addMarkingWithText(d *drawing.Drawing, md *gen.MarkingData) {
	switch md.MarkingSide {
	case "TS":
		drawContour(d, md.Contour, MarkingTSColor)
		drawMarkingText(d, md, MarkingTextTSColor)

	case "OS":
		drawContour(d, md.Contour, MarkingOSColor)
		drawMarkingText(d, md, MarkingTextOSColor)

	case "BOTH":
		drawContour(d, md.Contour, MarkingTSColor)
		drawMarkingText(d, md, MarkingTextTSColor)

		drawContour(d, md.Contour, MarkingOSColor)
		drawMarkingText(d, md, MarkingTextOSColor)
	}
}
