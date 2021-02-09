package draw

import (
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
	PosNoTextHeight             = 25
	MarkingTextHeight           = 20
	BevelMarkingTextHeight      = 20
	NestixBevelAttrTextHeight   = 1
	NestixAttrTextHeight        = 5.0
	AddAngleLength              = 10
	OrderString                 = "056001"
	AttrSpacing                 = 10
	AttrColor                   = color.White
	DirectionTextColor          = color.Cyan
	DirectionMarkingColor       = color.Yellow
	EdgeDataColor               = color.Cyan
	EdgeNameColor               = color.Red
	GeometryColor               = color.Cyan
	IdleColor                   = color.Yellow
	MarkingTSColor              = color.Cyan
	MarkingTextTSColor          = color.Red
	MarkingOSColor              = 30
	MarkingTextOSColor          = 30
	MinPartAreaForDirectionMark = 300000
)

// do smth with it
var (
	partMirrored = false
)

func PlateToDXF(g *gen.PlateGen, w *wcog.WCOG) error {
	for _, p := range g.PartData {

		dir := time.Now().Format("06.01.02") + " " + p.BlockNo
		subdir := strconv.FormatFloat(g.GeneralData.RawThickness, 'f', -1, 64)
		filename := OrderString + "-" + p.BlockNo + "-" + p.PosNo + ".dxf"

		out := path.Join(dir, subdir, filename)

		if err := os.MkdirAll(filepath.Dir(out), 0755); err != nil {
			return err
		}

		if _, err := os.Stat(out); os.IsNotExist(err) {

			var tsLength, osLength float64

			for _, p := range g.PartData {
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
			}

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

			if p.Mirrored == 1 {
				partMirrored = true
			} else {
				partMirrored = false
			}

			addPosNo(d, p)

			if p.PartArea > MinPartAreaForDirectionMark {
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
		}
	}
	return nil
}

func drawContour(d *drawing.Drawing, c *gen.Contour, cl color.ColorNumber) {
	p := entity.NewLwPolyline()

	b := math2.SegmentBulge(c.Segments[0])
	p.AddVertex(c.Segments[0].StartU, c.Segments[0].StartV, b)

	for i := 0; i < len(c.Segments); i++ {

		b = 0.0
		if i < len(c.Segments)-1 {
			b = math2.SegmentBulge(c.Segments[i+1])
		}

		p.AddVertex(c.Segments[i].U, c.Segments[i].V, b)
	}
	p.SetColor(cl)
	d.AddEntity(p)
}

func addBevel(d *drawing.Drawing, c *gen.Contour, from, to int) {

	longestIdx := from
	longestLen := math2.SegmentLength(c.Segments[from])

	var x, y, rotation float64

	for i := from + 1; i <= to; i++ {
		seg := c.Segments[i]

		l := math2.LineLength(seg.StartU, seg.StartV, seg.U, seg.V)

		if l > longestLen {
			longestLen = l
			longestIdx = i
		}
	}

	x, y = math2.SegmentMidPoint(c.Segments[longestIdx])
	rotation = math2.ToDeg(math2.SegmentAngle(c.Segments[longestIdx]))

	bevelText, bevelNestix := Bevel(c.Segments[from].BevelData)

	// nestix
	n := entity.NewText()
	n.Value = bevelNestix
	n.Coord1 = []float64{x, y, 0}
	n.Height = NestixAttrTextHeight
	n.SetColor(color.White)

	if partMirrored {
		n.Rotation = rotation
	} else {
		n.Rotation = rotation + 180
	}

	d.AddEntity(n)
	// ---

	if partMirrored {
		x, y = math2.AddAngle(x, y, rotation, AddAngleLength)

		t := entity.NewText()
		t.Value = bevelText
		t.Coord1 = []float64{x, y, 0}
		t.Height = float64(BevelMarkingTextHeight)
		t.SetColor(color.Red)
		t.Rotation = rotation + 180
		t.Anchor(entity.CENTER_BOTTOM)
		d.AddEntity(t)
	} else {
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
}

func drawBurningContour(d *drawing.Drawing, c *gen.Contour, cl color.ColorNumber) {

	from := 0
	to := 0

	lastBevelCode := c.Segments[0].BevelData.BevelCode

	for i := 0; i < len(c.Segments); i++ {
		to = i

		if c.Segments[i].BevelData.BevelCode != lastBevelCode || i == len(c.Segments)-1 {

			p := entity.NewLwPolyline()

			b := math2.SegmentBulge(c.Segments[from])
			p.AddVertex(c.Segments[from].StartU, c.Segments[from].StartV, b)

			for j := from; j < to; j++ {
				b = 0.0
				if j < len(c.Segments)-1 {
					b = math2.SegmentBulge(c.Segments[j+1])
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

	if md.MarkingType == "INCLINATION_ANGLE_LINE" {
		angle := strconv.FormatFloat(md.InclinationAngle, 'f', -1, 64)

		sx := md.Contour.Segments[0].StartU
		sy := md.Contour.Segments[0].StartV

		ex := md.Contour.Segments[1].U
		ey := md.Contour.Segments[1].V

		x, y = math2.LineMidPoint(sx, sy, ex, ey)
		rot = math2.Angle(sx, sy, ex, ey)
		rot = math2.ToDeg(rot)

		drawText(d, x, y, float64(MarkingTextHeight), rot, angle, color)
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

	var sx, sy, ex, ey float64

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
							rot = math2.Angle(sx, sy, ex, ey)
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

	x, y = math2.AddAngle(x, y, rot, AddAngleLength)
	return
}

func filterPos(pos string) string {
	ss := strings.Split(pos, "-")
	s := ss[len(ss)-1]
	s = strings.Replace(s, "P", "", -1)
	s = strings.Replace(s, "S", "", -1)
	s = strings.Replace(s, "B", "", -1)
	s = strings.Replace(s, "C", "", -1)

	return s
}
