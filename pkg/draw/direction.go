package draw

import (
	"github.com/edanko/dxf/color"
	"github.com/edanko/dxf/drawing"
	"github.com/edanko/dxf/entity"
	"github.com/edanko/gen2dxf/pkg/gen"
)

const (
	DirectionTextColor    = color.Cyan
	DirectionMarkingColor = color.Yellow
)

func addDirection(d *drawing.Drawing, p *gen.PartData) {
	sds := p.StringData
	if sds == nil {
		return
	}

	var top, right string

	for _, v := range sds {
		switch v.Type {
		case "DIR_TOP":
			top = v.String
		case "DIR_RIGHT":
			right = v.String
		}
	}
	drawDirection(d, top, right, p.PartCogU, p.PartCogV)
}

func drawDirection(d *drawing.Drawing, top, right string, x, y float64) {
	l := entity.NewLine()
	l.Start = []float64{x + 0.0000, y + 90.9180, 0.0000}
	l.End = []float64{x - 13.6377, y + 77.2803, 0.0000}
	l.SetColor(DirectionMarkingColor)
	d.AddEntity(l)

	l = entity.NewLine()
	l.Start = []float64{x + 0.0000, y + 90.9180, 0.0000}
	l.End = []float64{x + 13.6377, y + 77.2803, 0.0000}
	l.SetColor(DirectionMarkingColor)
	d.AddEntity(l)

	l = entity.NewLine()
	l.Start = []float64{x + 0.0000, y + 13.6377, 0.0000}
	l.End = []float64{x + 0.0000, y + 90.9180, 0.0000}
	l.SetColor(DirectionMarkingColor)
	d.AddEntity(l)

	c := entity.NewCircle()
	c.Center = []float64{x + 13.6377, y + 13.6377, 0.0000}
	c.Radius = 13.6377
	a := entity.NewArc(c)
	a.Angle = []float64{180, 270}
	a.SetColor(DirectionMarkingColor)
	d.AddEntity(a)

	l = entity.NewLine()
	l.Start = []float64{x + 13.6377, y + 0.0000, 0.0000}
	l.End = []float64{x + 90.9180, y + 0.0000, 0.0000}
	l.SetColor(DirectionMarkingColor)
	d.AddEntity(l)

	l = entity.NewLine()
	l.Start = []float64{x + 77.2803, y + 13.6377, 0.0000}
	l.End = []float64{x + 90.9180, y + 0.0000, 0.0000}
	l.SetColor(DirectionMarkingColor)
	d.AddEntity(l)

	l = entity.NewLine()
	l.Start = []float64{x + 77.2803, y - 13.6377, 0.0000}
	l.End = []float64{x + 90.9180, y + 0.0000, 0.0000}
	l.SetColor(DirectionMarkingColor)
	d.AddEntity(l)

	t := entity.NewText()
	t.Value = top
	t.Height = 22
	t.Rotation = 270
	t.Coord1 = []float64{x + 4.5459, y + 78.7956, 0.0000}
	t.SetColor(DirectionTextColor)
	d.AddEntity(t)

	t = entity.NewText()
	t.Value = right
	t.Height = 22
	t.Rotation = 0
	t.Coord1 = []float64{x + 40.9131, y + 12.1224, 0.0000}
	t.SetColor(DirectionTextColor)
	d.AddEntity(t)
}
