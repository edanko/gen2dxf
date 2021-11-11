package draw

import (
	"github.com/edanko/dxf/drawing"
)

func addNestixAttrs(d *drawing.Drawing, quality, quantity, thickness, order, section, partid, ina string) {
	as := []string{
		"[M]" + quality,
		"[Q]" + quantity,
		"[T]" + thickness,
		"[B]" + order,
		"[S]" + section,
		"[P]" + partid,
		"[I]" + ina,
		"[N]" + ina,
		"[A]" + ina,
	}

	for i, a := range as {
		t, _ := d.Text(a, 0, -float64(i)*AttrSpacing, 0, NestixAttrTextHeight)
		t.SetColor(AttrColor)
	}
}
