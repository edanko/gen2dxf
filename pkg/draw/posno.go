package draw

import (
	"fmt"

	"github.com/edanko/dxf/color"
	"github.com/edanko/dxf/entity"
	"github.com/edanko/gen"
)

func (d *Drawer) addPosNo(p *gen.PartData) {
	text := fmt.Sprintf("%s-%s-%s", p.ShipNo, p.BlockNo, p.PosNo)

	if p.Mirrored {
		text += "(m)"
	}

	t := entity.NewText()
	t.Value = text
	t.Coord1 = []float64{p.PartCog.X, p.PartCog.Y, 0}
	t.Height = posNoTextHeight
	t.SetColor(color.Red)
	t.Anchor(entity.CENTER_TOP)
	d.drawing.AddEntity(t)
}
