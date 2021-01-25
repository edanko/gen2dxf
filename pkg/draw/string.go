package draw

import (
	"fmt"

	"github.com/edanko/dxf/color"
	"github.com/edanko/dxf/drawing"
	"github.com/edanko/gen2dxf/pkg/gen"
)

func addStringData(d *drawing.Drawing, sds []*gen.StringData) {
	if sds == nil {
		return
	}
	for _, sd := range sds {
		addString(d, sd)
	}
}

func addString(d *drawing.Drawing, sd *gen.StringData) {
	switch sd.Type {

	case "DIR_RIGHT":
	case "DIR_TOP":
	case "THICKNESS_&_QUALITY":
	case "EXCESS_LIMIT": //same as excess_geometry
	case "PART_NAME":
	case "POSNO":

	case "EXCESS_GEOMETRY":
		drawText(d, sd.PosU, sd.PosV, 25, sd.Angle, sd.String, color.Red)

	case "BEVEL_ANGLE":
		//drawText(d, sd.PosU, sd.PosV, 25, sd.Angle, sd.String, color.Red)

	default:
		fmt.Println("string data type:", sd.Type)
		drawText(d, sd.PosU, sd.PosV, sd.Height, sd.Angle, sd.String, color.Red)
	}
}
