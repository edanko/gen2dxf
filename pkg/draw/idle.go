package draw

import (
	"gen2dxf/pkg/gen"

	"github.com/edanko/dxf/drawing"
)

func addIdleData(d *drawing.Drawing, ids []*gen.IdleData) {
	if ids == nil {
		return
	}
	for _, id := range ids {
		addIdle(d, id)
	}
}

func addIdle(d *drawing.Drawing, sd *gen.IdleData) {
	drawContour(d, sd.Contour, IdleColor)
}
