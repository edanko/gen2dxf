package draw

import (
	"github.com/edanko/dxf/color"
	"github.com/edanko/dxf/drawing"
	"github.com/edanko/dxf/entity"
	"github.com/edanko/gen2dxf/pkg/gen"
)

func addPosNo(d *drawing.Drawing, p *gen.PartData) {

	id := OrderString + "-" + p.BlockNo + "-" + p.PosNo
	if p.Mirrored == 1 {
		id += "m"
	}

	/* 	var x, y float64

	   	sds := p.StringData
	   	if sds == nil {
	   		x = p.PartCogU
	   		y = p.PartCogV
	   	} else {
	   		for _, v := range sds {
	   			if v.Type == "PART_NAME" {
	   				x = v.PosU
	   				y = v.PosV
	   				break
	   			}
	   		}
	   	} */

	x := p.PartCogU
	y := p.PartCogV - 15

	t := entity.NewText()
	t.Value = id
	t.Coord1 = []float64{x, y, 0}
	t.Height = float64(PosNoTextHeight)
	t.SetColor(color.Red)
	t.Anchor(entity.CENTER_TOP)
	d.AddEntity(t)
}
