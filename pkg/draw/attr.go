package draw

import (
	"time"

	"github.com/edanko/dxf/drawing"
)

func addNestixAttrs(d *drawing.Drawing, quality, quantity, thickness, order, section, partid, ina string) {
	t, _ := d.Text("[M]"+quality, 0, 0, 0, NestixAttrTextHeight)
	t.SetColor(AttrColor)
	t, _ = d.Text("[Q]"+quantity, 0, -AttrSpacing, 0, NestixAttrTextHeight)
	t.SetColor(AttrColor)
	t, _ = d.Text("[T]"+thickness, 0, -2*AttrSpacing, 0, NestixAttrTextHeight)
	t.SetColor(AttrColor)
	t, _ = d.Text("[B]"+order, 0, -3*AttrSpacing, 0, NestixAttrTextHeight)
	t.SetColor(AttrColor)
	t, _ = d.Text("[S]"+section, 0, -4*AttrSpacing, 0, NestixAttrTextHeight)
	t.SetColor(AttrColor)
	t, _ = d.Text("[P]"+partid, 0, -5*AttrSpacing, 0, NestixAttrTextHeight)
	t.SetColor(AttrColor)
	t, _ = d.Text("[I]"+ina, 0, -6*AttrSpacing, 0, NestixAttrTextHeight)
	t.SetColor(AttrColor)
	t, _ = d.Text("[N]"+ina, 0, -7*AttrSpacing, 0, NestixAttrTextHeight)
	t.SetColor(AttrColor)
	t, _ = d.Text("[A]"+ina, 0, -8*AttrSpacing, 0, NestixAttrTextHeight)
	t.SetColor(AttrColor)
	t, _ = d.Text("[O]"+time.Now().Format("02.01.06"), 0, -9*AttrSpacing, 0, NestixAttrTextHeight)
	t.SetColor(AttrColor)
}
