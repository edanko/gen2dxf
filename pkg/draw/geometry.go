package draw

import (
	"github.com/edanko/dxf/drawing"
	"github.com/edanko/gen2dxf/pkg/gen"
)

func addGeometryData(d *drawing.Drawing, gds []*gen.GeometryData) {
	if gds == nil {
		return
	}
	for _, gd := range gds {
		if gd.Type == "BEVEL_SYMBOL" {
			continue
		}
		addGeometry(d, gd)
	}
}

func addGeometry(d *drawing.Drawing, gd *gen.GeometryData) {
	drawContour(d, gd.Contour, GeometryColor)
}
