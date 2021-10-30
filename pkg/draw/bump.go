package draw

import (
	"gen2dxf/pkg/gen"

	"github.com/edanko/dxf/drawing"
)

func addBumpData(d *drawing.Drawing, gds []*gen.BumpData) {
	if gds == nil {
		return
	}
	for _, gd := range gds {
		addBump(d, gd)
	}
}

func addBump(d *drawing.Drawing, gd *gen.BumpData) {
	drawContour(d, gd.Contour, GeometryColor)
}
