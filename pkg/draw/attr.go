package draw

import (
	"time"

	"github.com/edanko/dxf/color"
	"github.com/edanko/dxf/drawing"
)

const (
	attrSpacing = 10
)

// add nestix attributes
func addNestixAttrs(d *drawing.Drawing, quality, quantity, thickness, order, section, partid, ina string) {
	t, _ := d.Text("[M]"+quality, 0, 0, 0, NestixAttrTextHeight)
	t.SetColor(color.White)
	t, _ = d.Text("[Q]"+quantity, 0, -attrSpacing, 0, NestixAttrTextHeight)
	t.SetColor(color.White)
	t, _ = d.Text("[T]"+thickness, 0, -2*attrSpacing, 0, NestixAttrTextHeight)
	t.SetColor(color.White)
	t, _ = d.Text("[B]"+order, 0, -3*attrSpacing, 0, NestixAttrTextHeight)
	t.SetColor(color.White)
	t, _ = d.Text("[S]"+section, 0, -4*attrSpacing, 0, NestixAttrTextHeight)
	t.SetColor(color.White)
	t, _ = d.Text("[P]"+partid, 0, -5*attrSpacing, 0, NestixAttrTextHeight)
	t.SetColor(color.White)
	t, _ = d.Text("[I]"+ina, 0, -6*attrSpacing, 0, NestixAttrTextHeight)
	t.SetColor(color.White)
	t, _ = d.Text("[N]"+ina, 0, -7*attrSpacing, 0, NestixAttrTextHeight)
	t.SetColor(color.White)
	t, _ = d.Text("[A]"+ina, 0, -8*attrSpacing, 0, NestixAttrTextHeight)
	t.SetColor(color.White)
	t, _ = d.Text("[O]"+time.Now().Format("02.01.06"), 0, -9*attrSpacing, 0, NestixAttrTextHeight)
	t.SetColor(color.White)
}
