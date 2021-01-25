package draw

import (
	"github.com/edanko/dxf/color"
	"github.com/edanko/dxf/drawing"
	"github.com/edanko/gen2dxf/pkg/gen"
)

func addBurningData(d *drawing.Drawing, bds []*gen.BurningData) {
	if bds == nil {
		return
	}
	for _, bd := range bds {
		addBurning(d, bd)
	}
}

func addBurning(d *drawing.Drawing, bd *gen.BurningData) {
	drawBurningContour(d, bd.Contour, color.Green)
}
