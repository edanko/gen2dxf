package draw

import (
	"gen2dxf/pkg/gen"

	"github.com/edanko/dxf/color"
	"github.com/edanko/dxf/drawing"
	"github.com/edanko/dxf/entity"
)

func addPosNo(d *drawing.Drawing, p *gen.PartData) {
	id := OrderString + "-" + p.BlockNo + "-" + p.PosNo

	if p.Mirrored == 1 {
		id += "(m)"
	}

	t := entity.NewText()
	t.Value = id
	t.Coord1 = []float64{p.PartCogU, p.PartCogV, 0}
	t.Height = float64(PosNoTextHeight)
	t.SetColor(color.Red)
	t.Anchor(entity.CENTER_TOP)
	d.AddEntity(t)
}
