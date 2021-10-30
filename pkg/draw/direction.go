package draw

import (
	"gen2dxf/pkg/gen"

	"github.com/edanko/dxf/drawing"
	"github.com/edanko/dxf/entity"
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

	var mirr bool

	if p.Mirrored == 1 {
		mirr = true
	}

	drawDirection(d, p.PartCogU, p.PartCogV, top, right, mirr)
}

func drawDirection(d *drawing.Drawing, x, y float64, top, right string, mirr bool) {

	l1s := []float64{x + 0.0000, y + 90.9180, 0.0000}
	l1e := []float64{x - 13.6377, y + 77.2803, 0.0000}

	l2s := []float64{x + 0.0000, y + 90.9180, 0.0000}
	l2e := []float64{x + 13.6377, y + 77.2803, 0.0000}

	l3s := []float64{x + 0.0000, y + 13.6377, 0.0000}
	l3e := []float64{x + 0.0000, y + 90.9180, 0.0000}

	var cc, ca, l4s, l4e, l5s, l5e, l6s, l6e, tCoord []float64
	if mirr {
		cc = []float64{x - 13.6377, y + 13.6377, 0.0000}
		ca = []float64{270, 0}

		l4s = []float64{x - 13.6377, y + 0.0000, 0.0000}
		l4e = []float64{x - 90.9180, y + 0.0000, 0.0000}

		l5s = []float64{x - 77.2803, y + 13.6377, 0.0000}
		l5e = []float64{x - 90.9180, y + 0.0000, 0.0000}

		l6s = []float64{x - 77.2803, y - 13.6377, 0.0000}
		l6e = []float64{x - 90.9180, y + 0.0000, 0.0000}

		tCoord = []float64{x - 105, y + 12.1224, 0.0000}
	} else {
		cc = []float64{x + 13.6377, y + 13.6377, 0.0000}
		ca = []float64{180, 270}

		l4s = []float64{x + 13.6377, y + 0.0000, 0.0000}
		l4e = []float64{x + 90.9180, y + 0.0000, 0.0000}

		l5s = []float64{x + 77.2803, y + 13.6377, 0.0000}
		l5e = []float64{x + 90.9180, y + 0.0000, 0.0000}

		l6s = []float64{x + 77.2803, y - 13.6377, 0.0000}
		l6e = []float64{x + 90.9180, y + 0.0000, 0.0000}

		tCoord = []float64{x + 40.9131, y + 12.1224, 0.0000}
	}

	l := entity.NewLine()
	l.Start = l1s
	l.End = l1e
	l.SetColor(DirectionMarkingColor)
	d.AddEntity(l)

	l = entity.NewLine()
	l.Start = l2s
	l.End = l2e
	l.SetColor(DirectionMarkingColor)
	d.AddEntity(l)

	l = entity.NewLine()
	l.Start = l3s
	l.End = l3e
	l.SetColor(DirectionMarkingColor)
	d.AddEntity(l)

	c := entity.NewCircle()
	c.Center = cc
	c.Radius = 13.6377
	a := entity.NewArc(c)
	a.Angle = ca
	a.SetColor(DirectionMarkingColor)
	d.AddEntity(a)

	l = entity.NewLine()
	l.Start = l4s
	l.End = l4e
	l.SetColor(DirectionMarkingColor)
	d.AddEntity(l)

	l = entity.NewLine()
	l.Start = l5s
	l.End = l5e
	l.SetColor(DirectionMarkingColor)
	d.AddEntity(l)

	l = entity.NewLine()
	l.Start = l6s
	l.End = l6e
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
	t.Coord1 = tCoord
	t.SetColor(DirectionTextColor)
	d.AddEntity(t)
}
