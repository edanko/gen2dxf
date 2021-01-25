package draw

import (
	"github.com/edanko/dxf/color"
	"github.com/edanko/dxf/drawing"
	"github.com/edanko/gen2dxf/pkg/gen"
)

const (
	IdleColor = color.Yellow
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
