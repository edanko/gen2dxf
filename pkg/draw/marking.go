package draw

import (
	"fmt"

	"github.com/edanko/gen"
)

func (d *Drawer) addMarkingData(mds []gen.MarkingData) {
	if mds == nil {
		return
	}
	for _, md := range mds {
		switch md.MarkingType {
		case "STIFFENER",
			"PANEL",
			"BRACKET_ON_INTERSEC._PANEL",
			"INCLINATION_ANGLE_LINE",
			"ROLL_AXIS",
			"LONGITUDINAL",
			"PANEL_GSD",
			"STIFFENER_GSD":
			d.addMarkingWithText(md)

		case "CUTOUT_CROSSING",
			"FLANGE_MARKING_LINE",
			"BENDING_MARK",
			"ADDED_CONTOUR",
			"FACE_PLATE_ALIGNMENT",
			"FOLDED_FLANGE",
			"HULL_MARK",
			"CORNER_GSD",
			"CL_GSD",
			"":
			d.addMarking(md)

		case "BRACKET":
			continue

		default:
			fmt.Println("unprocessed marking type:", md.MarkingType)
		}
	}
}

func (d *Drawer) addMarking(md gen.MarkingData) {
	switch md.MarkingSide {
	case "TS":
		d.drawContour(md.Contour, markingTSColor)

	case "OS":
		d.drawContour(md.Contour, markingOSColor)

	case "BOTH":
		d.drawContour(md.Contour, markingTSColor)
		d.drawContour(md.Contour, markingOSColor)
	}
}

func (d *Drawer) addMarkingWithText(md gen.MarkingData) {
	switch md.MarkingSide {
	case "TS":
		d.drawContour(md.Contour, markingTSColor)
		d.drawMarkingText(md, markingTextTSColor)

	case "OS":
		d.drawContour(md.Contour, markingOSColor)
		d.drawMarkingText(md, markingTextOSColor)

	case "BOTH":
		d.drawContour(md.Contour, markingTSColor)
		d.drawMarkingText(md, markingTextTSColor)

		d.drawContour(md.Contour, markingOSColor)
		d.drawMarkingText(md, markingTextOSColor)
	}
}
