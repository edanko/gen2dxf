package draw

import (
	"math"

	"github.com/edanko/dxf/drawing"
	"github.com/edanko/dxf/entity"
	"github.com/edanko/gen2dxf/pkg/gen"
	"github.com/edanko/gen2dxf/pkg/math2"
)

func addEdgeData(d *drawing.Drawing, gds []*gen.EdgeData) {
	if gds == nil {
		return
	}
	for _, gd := range gds {
		addEdge(d, gd)
	}
}

func addEdge(d *drawing.Drawing, gd *gen.EdgeData) {
	l := entity.NewLine()
	l.Start = []float64{gd.StartU, gd.StartV, 0}
	l.End = []float64{gd.EndU, gd.EndV, 0}
	l.SetColor(EdgeDataColor)
	d.AddEntity(l)

	x, y := math2.LineMidPoint(gd.StartU, gd.StartV, gd.EndU, gd.EndV)

	rot := math2.Angle(gd.StartU, gd.StartV, gd.EndU, gd.EndV) * 180.0 / math.Pi
	x, y = math2.AddAngle(x, y, rot+90.0, AddAngleLength)

	drawText(d, x, y, float64(MarkingTextHeight), rot, gd.Name, EdgeNameColor)
}
