package draw

import (
	"fmt"

	"gen2dxf/pkg/gen"

	"github.com/edanko/dxf/drawing"
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

		case "INCLINATION_ANGLE_LINE":
			addMarkingWithText(d, md)

		case "CUTOUT_CROSSING", "FLANGE_MARKING_LINE":
			addMarking(d, md)

		case "BENDING_MARK", "ADDED_CONTOUR", "FACE_PLATE_ALIGNMENT":
			addMarking(d, md)

		case "FOLDED_FLANGE", "HULL_MARK":
			addMarking(d, md)

		case "ROLL_AXIS":
			addMarkingWithText(d, md)

		case "LONGITUDINAL":
			addMarkingWithText(d, md)

		case "CORNER_GSD", "CL_GSD":
			addMarking(d, md)

		case "PANEL_GSD", "STIFFENER_GSD":
			addMarkingWithText(d, md)

		case "":
			addMarking(d, md)

		default:
			fmt.Println("unprocessed marking type:", md.MarkingType)
		}
	}
}

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
