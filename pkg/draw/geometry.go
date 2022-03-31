package draw

import (
	"github.com/edanko/gen"
)

func (d *Drawer) addGeometryData(gds []*gen.GeometryData) {
	if gds == nil {
		return
	}
	for _, gd := range gds {
		if gd.Type == "BEVEL_SYMBOL" {
			continue
		}
		d.drawContour(gd.Contour, geometryColor)
	}
}
