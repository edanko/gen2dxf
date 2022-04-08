package draw

import (
	"github.com/edanko/gen"
)

func (d *Drawer) addBurningData(bds []gen.BurningData) {
	if bds == nil {
		return
	}
	for _, bd := range bds {
		d.drawBurningContour(bd.Contour, burningColor)
	}
}
