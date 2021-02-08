package draw

import (
	"errors"
	"math"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/edanko/dxf/color"
	"github.com/edanko/dxf/drawing"
	"github.com/edanko/dxf/entity"
	"github.com/edanko/gen2dxf/pkg/gen"
	"github.com/edanko/gen2dxf/pkg/math2"
	"github.com/edanko/gen2dxf/pkg/wcog"
)

const (
	PosNoTextHeight           = 25
	MarkingTextHeight         = 20
	BevelMarkingTextHeight    = 20
	NestixBevelAttrTextHeight = 1
	NestixAttrTextHeight      = 5
	AddAngleLength            = 10
	OrderString               = "056001"
)

func PlateToDXF(g *gen.PlateGen, w *wcog.WCOG) error {
	if g == nil {
		return errors.New("[x] possible listed profile")
	}

	for _, p := range g.PartData {

		out := path.Join(time.Now().Format("06.01.02")+" "+p.BlockNo+" original", OrderString+"-"+p.BlockNo+"-"+p.PosNo+".dxf")
		out2 := path.Join(time.Now().Format("06.01.02")+" "+p.BlockNo+" original by thickness", strconv.FormatFloat(g.GeneralData.RawThickness, 'f', -1, 64), OrderString+"-"+p.BlockNo+"-"+p.PosNo+".dxf")

		if err := os.MkdirAll(filepath.Dir(out), 0755); err != nil {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(out2), 0755); err != nil {
			return err
		}

		if _, err := os.Stat(out); os.IsNotExist(err) {

			d := drawing.New()

			quality := g.GeneralData.Quality

			if quality == "RSA" {
				quality = "A"
			}

			quantity := strconv.Itoa(w.GetQuantity(p.BlockNo, p.PosNo))
			thickness := strconv.FormatFloat(g.GeneralData.RawThickness, 'f', -1, 64)
			order := OrderString
			section := p.BlockNo
			posno := p.PosNo
			ina := order + section + posno
			addNestixAttrs(d, quality, quantity, thickness, order, section, posno, ina)

			addPosNo(d, p)

			if p.PartArea > 300000 {
				addDirection(d, p)
			}
			//addIdleData(d, p.IdleData)
			addBurningData(d, p.BurningData)
			addMarkingData(d, p.MarkingData)
			addStringData(d, p.StringData)
			//addBumpData(d, p.BumpData)
			//addEdgeData(d, p.EdgeData)
			addLabeltextData(d, p.LabeltextData)
			addGeometryData(d, p.GeometryData)

			err := d.SaveAs(out)
			if err != nil {
				return err
			}

			err = d.SaveAs(out2)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func drawContour(d *drawing.Drawing, c *gen.Contour, cl color.ColorNumber) {
	/* 	p := entity.NewLwPolyline()

	   	for _, s := range c.Segments {
	   		p.AddVertex(s.StartU, s.StartV, math2.GetBulge(s.Radius, s.Amp, s.Sweep))
	   		p.AddVertex(s.U, s.V, math2.GetBulge(s.Radius, s.Amp, s.Sweep))
	   	}
	   	p.SetColor(cl)

	   	p.Closed = false
	   	d.AddEntity(p) */

	p := entity.NewLwPolyline()

	b := math2.GetSegmentBulge(c.Segments[0])
	p.AddVertex(c.Segments[0].StartU, c.Segments[0].StartV, b)

	for i := 0; i < len(c.Segments); i++ {

		b = 0.0
		if i < len(c.Segments)-1 {
			b = math2.GetSegmentBulge(c.Segments[i+1])
		}

		p.AddVertex(c.Segments[i].U, c.Segments[i].V, b)
	}
	//p.AddVertex(c.Segments[len(c.Segments)-1].U, c.Segments[len(c.Segments)-1].V, 0)

	p.SetColor(cl)
	d.AddEntity(p)
}

func addBevel(d *drawing.Drawing, c *gen.Contour, from, to int) {

	longestIdx := from
	longestLen := math2.LineLength(c.Segments[from].StartU, c.Segments[from].StartV, c.Segments[from].U, c.Segments[from].V)

	/* 	radiusExist := false
	   	// check radius segment existance
	   	for i := from + 1; i <= to; i++ {
	   		if to-from > 1 {
	   			if c.Segments[i].Radius != 0 {
	   				radiusExist = true
	   			}
	   		}
	   	} */

	var x, y, rotation float64

	//if radiusExist {
	for i := from + 1; i <= to; i++ {
		seg := c.Segments[i]

		l := math2.LineLength(seg.StartU, seg.StartV, seg.U, seg.V)

		if l > longestLen {
			longestLen = l
			longestIdx = i
		}
	}

	x, y = math2.GetSegmentMidPoint(c.Segments[longestIdx])
	rotation = math2.RadToDeg(math2.GetSegmentAngle(c.Segments[longestIdx]))
	/* 	} else {
		x, y = math2.LineMidPoint(c.Segments[from].StartU, c.Segments[from].StartV, c.Segments[to].U, c.Segments[to].V)
		rotation = math2.RadToDeg(math2.GetAngle(c.Segments[from].StartU, c.Segments[from].StartV, c.Segments[to].U, c.Segments[to].V))
	} */

	bevelText, bevelNestix := GetBevel(c.Segments[from].BevelData)

	// nestix
	n := entity.NewText()
	n.Value = bevelNestix
	n.Coord1 = []float64{x, y, 0}
	n.Height = float64(NestixAttrTextHeight)
	n.SetColor(color.White)
	n.Rotation = rotation + 180
	d.AddEntity(n)
	// ---

	x, y = math2.AddAngle(x, y, rotation-90.0, -AddAngleLength)

	t := entity.NewText()
	t.Value = bevelText
	t.Coord1 = []float64{x, y, 0}
	t.Height = float64(BevelMarkingTextHeight)
	t.SetColor(color.Red)
	t.Rotation = rotation //+ 180
	t.Anchor(entity.CENTER_BOTTOM)
	d.AddEntity(t)
}

/* func drawBurningContour(d *drawing.Drawing, c *gen.Contour, cl color.ColorNumber) {
	poly := entity.NewLwPolyline()
	lastBevelCode := c.Segments[0].BevelData.BevelCode

	from := 0
	to := 0

	for i, s := range c.Segments {
		if s.BevelData.BevelCode == lastBevelCode {
			b := math2.GetBulge(s.Radius, s.Amp, s.Sweep)
			poly.AddVertex(s.StartU, s.StartV, b)
			poly.AddVertex(s.U, s.V, b)
			to = i
		} else {
			poly.SetColor(cl)
			d.AddEntity(poly)
			addBevel(d, c, from, to)

			from = i
			to = i
			poly = entity.NewLwPolyline()
			lastBevelCode = s.BevelData.BevelCode

			b := math2.GetBulge(s.Radius, s.Amp, s.Sweep)
			poly.AddVertex(s.StartU, s.StartV, b)
			poly.AddVertex(s.U, s.V, b)
		}
	}

	poly.SetColor(cl)
	d.AddEntity(poly)
	addBevel(d, c, from, to)
} */

/*

func (o *out) draw(c *Contour, cl color.ColorNumber) {
	p := entity.NewLwPolyline()

	var sx, sy float64
	if c.StartU == 0 && c.StartV == 0 {
		sx = c.U
		sy = c.V
	} else {
		sx = c.StartU
		sy = c.StartV
	}

	for _, s := range c.Segments {
		p.AddVertex(sx, sy, getBulge(s.Radius, s.Amp, s.Sweep))

		sx = s.U
		sy = s.V
	}
	p.AddVertex(sx, sy, 0)
	p.SetColor(cl)
	// TODO: smth
	p.Closed = false
	o.AddEntity(p)
}

*/

func drawBurningContour(d *drawing.Drawing, c *gen.Contour, cl color.ColorNumber) {

	/* 	p := entity.NewLwPolyline()

	   	var sx, sy float64
	   	sx = c.StartU
	   	sy = c.StartV

	   	for _, s := range c.Segments {
	   		p.AddVertex(sx, sy, math2.GetBulge(s.Radius, s.Amp, s.Sweep))

	   		sx = s.U
	   		sy = s.V
	   	}
	   	p.AddVertex(sx, sy, 0)
	   	p.SetColor(cl)
	   	// TODO: smth
	   	p.Closed = false
	   	d.AddEntity(p)

	   	return */

	from := 0
	to := 0

	lastBevelCode := c.Segments[0].BevelData.BevelCode

	for i := 0; i < len(c.Segments); i++ {
		to = i

		if c.Segments[i].BevelData.BevelCode != lastBevelCode || i == len(c.Segments)-1 {

			p := entity.NewLwPolyline()

			b := math2.GetSegmentBulge(c.Segments[from])
			p.AddVertex(c.Segments[from].StartU, c.Segments[from].StartV, b)

			for j := from; j < to; j++ {
				b = 0.0
				if j < len(c.Segments)-1 {
					b = math2.GetSegmentBulge(c.Segments[j+1])
				}

				p.AddVertex(c.Segments[j].U, c.Segments[j].V, b)
			}

			if i == len(c.Segments)-1 {
				p.AddVertex(c.Segments[to].U, c.Segments[to].V, 0)
			}

			p.SetColor(cl)
			d.AddEntity(p)
			addBevel(d, c, from, to-1)

			lastBevelCode = c.Segments[i].BevelData.BevelCode
			from = i
		}
	}
}

func drawMarkingText(d *drawing.Drawing, md *gen.MarkingData, color color.ColorNumber) {
	x, y, rot := getMarkingTextPosition(md.Contour)

	if strings.Contains(md.MarkingName, "ROLL AXIS") {
		drawText(d, x, y, float64(MarkingTextHeight), rot, md.MarkingName, color)
		return
	}

	if len(md.MarkingAttri) > 0 {
		drawText(d, x, y, float64(MarkingTextHeight), rot, filterPos(md.MarkingAttri), color)
	} else if len(md.MarkingAttriSb) > 0 {
		drawText(d, x, y, float64(MarkingTextHeight), rot, filterPos(md.MarkingAttriSb), color)
	} else {
		return
	}
}

func drawText(d *drawing.Drawing, x, y, height, rotation float64, text string, color color.ColorNumber) {
	t := entity.NewText()
	t.Value = text
	t.Coord1 = []float64{x, y, 0}
	t.Height = height
	t.SetColor(color)
	t.Rotation = rotation
	t.Anchor(entity.CENTER_BOTTOM)
	d.AddEntity(t)
}

func getMarkingTextPosition(c *gen.Contour) (x float64, y float64, rot float64) {

	var sx, sy, ex, ey float64 //, cx, cy, r float64

	if len(c.Segments) > 1 {

		segmentsCount := len(c.Segments) - 1

		found := false

		for i := 0; i < len(c.Segments); i++ {
			if math.Round(c.Segments[i].Radius) == 10 {
				if i < segmentsCount-2 {
					if math.Round(c.Segments[i+1].Radius) == 30 || math.Round(c.Segments[i+1].Radius) == 35 || math.Round(c.Segments[i+1].Radius) == 50 {
						if math.Round(c.Segments[i+2].Radius) == 10 {
							found = true

							sx = c.Segments[i-1].U
							sy = c.Segments[i-1].V

							ex = c.Segments[i+2].U
							ey = c.Segments[i+2].V

							x, y = math2.LineMidPoint(sx, sy, ex, ey)

							rot = math2.GetAngle(sx, sy, ex, ey) * 180.0 / math.Pi

							break
						}
					}
				}
			}
		}

		if !found {
			longestIdx := 0
			longestLen := math2.LineLength(c.Segments[0].StartU, c.Segments[0].StartV, c.Segments[0].U, c.Segments[0].V)

			for i := 1; i < len(c.Segments); i++ {
				seg := c.Segments[i]

				l := math2.LineLength(seg.StartU, seg.StartV, seg.U, seg.V)

				if l > longestLen {
					longestLen = l
					longestIdx = i
				}
			}

			seg := c.Segments[longestIdx]

			x, y = math2.GetSegmentMidPoint(seg)

			rot = math2.GetAngle(seg.StartU, seg.StartV, seg.U, seg.V) * 180.0 / math.Pi

		}
	} else {
		seg := c.Segments[0]

		x, y = math2.GetSegmentMidPoint(seg)

		rot = math2.GetAngle(seg.StartU, seg.StartV, seg.U, seg.V) * 180.0 / math.Pi

	}

	x, y = math2.AddAngle(x, y, rot, AddAngleLength)

	return
}

func filterPos(s string) string {
	poss := strings.Split(s, "-")
	pos := poss[len(poss)-1]
	pos = strings.Replace(pos, "P", "", -1)
	pos = strings.Replace(pos, "S", "", -1)
	pos = strings.Replace(pos, "B", "", -1)
	pos = strings.Replace(pos, "C", "", -1)

	return pos
}
