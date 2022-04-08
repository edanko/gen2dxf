package draw

import (
	"github.com/edanko/dxf/color"
	"github.com/edanko/gen"
)

func (d *Drawer) addStringData(sds []*gen.StringData) {
	if sds == nil {
		return
	}
	for _, sd := range sds {
		if sd.Type == "EXCESS_GEOMETRY" {
			d.drawText(sd.Pos.X, sd.Pos.Y, 25, sd.Angle, sd.String, color.Red)
		}
	}
}
