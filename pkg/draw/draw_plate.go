package draw

import (
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/edanko/gen2dxf/pkg/math2"

	"github.com/edanko/dxf/color"
	"github.com/edanko/dxf/drawing"
	"github.com/edanko/dxf/entity"
	"github.com/edanko/gen"
)

const (
	posNoTextHeight        float64 = 25
	markingTextHeight      float64 = 20
	bevelMarkingTextHeight float64 = 20
	attrTextHeight         float64 = 5.0
	attrSpacing            float64 = 10
	addAngleLength         float64 = 10
	attrColor                      = color.White
	burningColor                   = color.Green
	geometryColor                  = color.Cyan
	markingTSColor                 = color.Cyan
	markingTextTSColor             = color.Red
	markingOSColor                 = 30 // orange
	markingTextOSColor             = 30 // orange
)

type Drawer struct {
	drawing *drawing.Drawing
	writer  io.Writer
}

func NewDrawer(writer io.Writer) *Drawer {
	return &Drawer{
		drawing: drawing.New(),
		writer:  writer,
	}
}

func (d *Drawer) PartToDXF(p *gen.PartData) error {
	var tsLength, osLength float64

	for _, m := range p.MarkingData {
		var l float64

		for _, s := range m.Contour.Segments {
			l += math2.SegmentLength(s)
		}

		switch m.MarkingSide {
		case "TS":
			tsLength += l
		case "OS":
			osLength += l
		case "BOTH":
			tsLength += l
			osLength += l
		}
	}

	if osLength > tsLength {
		p.Mirror()
	}

	quantity := strconv.Itoa(p.Quantity)
	thickness := strconv.FormatFloat(p.Thickness, 'g', -1, 64)
	order := p.ShipNo
	section := p.BlockNo
	posno := p.PosNo
	ina := order + section + posno
	err := d.addAttrs(p.Quality, quantity, thickness, order, section, posno, ina)
	if err != nil {
		return err
	}

	d.addPosNo(p)
	d.addBurningData(p.BurningData)
	d.addMarkingData(p.MarkingData)
	d.addStringData(p.StringData)
	d.addGeometryData(p.GeometryData)

	_, err = d.drawing.WriteTo(d.writer)
	if err != nil {
		return err
	}

	err = d.drawing.Close()
	if err != nil {
		return err
	}

	return nil
}

func (d *Drawer) drawContour(c *gen.Contour, cl color.ColorNumber) {
	p := entity.NewLwPolyline()

	b := math2.SegmentBulge(c.Segments[0])
	p.AddVertex(c.Segments[0].Start.X, c.Segments[0].Start.Y, b)

	for i := 0; i < len(c.Segments); i++ {

		b = 0
		if i < len(c.Segments)-1 {
			b = math2.SegmentBulge(c.Segments[i+1])
		}

		p.AddVertex(c.Segments[i].End.X, c.Segments[i].End.Y, b)
	}
	p.SetColor(cl)
	d.drawing.AddEntity(p)
}

func (d *Drawer) addBevel(c *gen.Contour, from, to int) {
	longestIdx := from
	longestLen := math2.SegmentLength(c.Segments[from])

	var x, y, rotation float64

	for i := from + 1; i <= to; i++ {
		seg := c.Segments[i]

		l := math2.LineLength(seg.Start, seg.End)

		if l > longestLen {
			longestLen = l
			longestIdx = i
		}
	}

	x, y = math2.SegmentMidPoint(c.Segments[longestIdx])
	rotation = math2.ToDeg(math2.SegmentAngle(c.Segments[longestIdx]))

	text, attrText := parseBevel(c.Segments[from].BevelData)

	// nestix
	n := entity.NewText()
	n.Value = attrText
	n.Coord1 = []float64{x, y, 0}
	n.Height = 1
	n.SetColor(color.White)
	n.Rotation = rotation
	d.drawing.AddEntity(n)
	// ---

	t := entity.NewText()
	t.Value = text

	if math2.IsClockwise(c) {
		x, y = math2.AddAngle(x, y, rotation-90, addAngleLength)
		// x, y = math2.AddAngle(x, y, rotation, AddAngleLength)

		t.Coord1 = []float64{x, y, 0}
		t.Rotation = rotation + 180
		// t.Rotation = rotation

	} else {
		// x, y = math2.AddAngle(x, y, rotation-90, AddAngleLength)
		x, y = math2.AddAngle(x, y, rotation+90, addAngleLength)

		t.Coord1 = []float64{x, y, 0}
		t.Rotation = rotation
		// t.Rotation = rotation + 180
	}

	t.Anchor(entity.CENTER_BOTTOM)
	t.Height = bevelMarkingTextHeight
	t.SetColor(color.Red)
	d.drawing.AddEntity(t)
}

func (d *Drawer) drawBurningContour(c *gen.Contour, cl color.ColorNumber) {
	from := 0
	to := 0

	lastBevelCode := c.Segments[0].BevelData.BevelCode

	for i := 0; i < len(c.Segments); i++ {
		to = i

		if c.Segments[i].BevelData.BevelCode != lastBevelCode || i == len(c.Segments)-1 {

			p := entity.NewLwPolyline()

			b := math2.SegmentBulge(c.Segments[from])
			p.AddVertex(c.Segments[from].Start.X, c.Segments[from].Start.Y, b)

			for j := from; j < to; j++ {
				b = 0.0
				if j < len(c.Segments)-1 {
					b = math2.SegmentBulge(c.Segments[j+1])
				}

				p.AddVertex(c.Segments[j].End.X, c.Segments[j].End.Y, b)
			}

			if i == len(c.Segments)-1 {
				p.AddVertex(c.Segments[to].End.X, c.Segments[to].End.Y, 0)
			}

			p.SetColor(cl)
			d.drawing.AddEntity(p)
			d.addBevel(c, from, to-1)

			lastBevelCode = c.Segments[i].BevelData.BevelCode
			from = i
		}
	}
}

func (d *Drawer) drawMarkingText(md *gen.MarkingData, color color.ColorNumber) {
	x, y, rot := getMarkingTextPosition(md.Contour)

	if strings.Contains(md.MarkingName, "ROLL AXIS") {
		d.drawText(x, y, markingTextHeight, rot, md.MarkingName, color)
		return
	}

	if md.MarkingType == "INCLINATION_ANGLE_LINE" {
		angle := strconv.FormatFloat(md.InclinationAngle, 'f', -1, 64)

		s := md.Contour.Segments[0].Start
		e := md.Contour.Segments[len(md.Contour.Segments)-1].End

		x, y = math2.LineMidPoint(s, e)
		rot = math2.Angle(s, e)
		rot = math2.ToDeg(rot)

		d.drawText(x, y, markingTextHeight, rot, angle, color)
		return
	}

	d.drawText(x, y, markingTextHeight, rot, filterPos(md.MarkingAttri), color)
}

func (d *Drawer) drawText(x, y, height, rotation float64, text string, color color.ColorNumber) {
	t := entity.NewText()
	t.Value = text
	t.Coord1 = []float64{x, y, 0}
	t.Height = height
	t.SetColor(color)
	t.Rotation = rotation
	t.Anchor(entity.CENTER_BOTTOM)
	d.drawing.AddEntity(t)
}

func getMarkingTextPosition(c *gen.Contour) (x float64, y float64, rot float64) {
	if len(c.Segments) > 1 {

		segmentsCount := len(c.Segments) - 1

		found := false

		for i := 0; i < len(c.Segments); i++ {
			if math.Round(c.Segments[i].Radius) == 10 {
				if i < segmentsCount-2 {
					if math.Round(c.Segments[i+1].Radius) == 30 || math.Round(c.Segments[i+1].Radius) == 35 || math.Round(c.Segments[i+1].Radius) == 50 {
						if math.Round(c.Segments[i+2].Radius) == 10 {
							found = true

							s := c.Segments[i-1].End
							e := c.Segments[i+2].End

							x, y = math2.LineMidPoint(s, e)
							rot = math2.Angle(s, e)
							rot = math2.ToDeg(rot)
							break
						}
					}
				}
			}
		}

		if !found {
			longestIdx := 0
			longestLen := math2.SegmentLength(c.Segments[0])

			for i := 1; i < len(c.Segments); i++ {
				seg := c.Segments[i]

				l := math2.SegmentLength(seg)

				if l > longestLen {
					longestLen = l
					longestIdx = i
				}
			}

			seg := c.Segments[longestIdx]
			x, y = math2.SegmentMidPoint(seg)
			rot = math2.SegmentAngle(seg)
			rot = math2.ToDeg(rot)
		}
	} else {
		seg := c.Segments[0]
		x, y = math2.SegmentMidPoint(seg)
		rot = math2.SegmentAngle(seg)
		rot = math2.ToDeg(rot)
	}

	x, y = math2.AddAngle(x, y, rot, addAngleLength)
	return
}

func filterPos(pos string) string {
	ss := strings.Split(pos, "-")
	s := ss[len(ss)-1]
	s = strings.ReplaceAll(s, "P", "")
	s = strings.ReplaceAll(s, "S", "")
	s = strings.ReplaceAll(s, "B", "")
	s = strings.ReplaceAll(s, "C", "")

	return s
}
