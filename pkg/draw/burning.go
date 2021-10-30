package draw

import (
	"gen2dxf/pkg/gen"

	"github.com/edanko/dxf/color"
	"github.com/edanko/dxf/drawing"
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
