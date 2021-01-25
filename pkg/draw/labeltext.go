package draw

import (
	"github.com/edanko/dxf/color"
	"github.com/edanko/dxf/drawing"
	"github.com/edanko/gen2dxf/pkg/gen"
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
	drawText(d, lt.TextPositionU, lt.TextPositionV, MarkingTextHeight, lt.TextAngle, lt.Text, color.Red)
}
