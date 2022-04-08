// TODO: think about moving this to gen repo
package math2

import (
	"math"

	"github.com/edanko/gen"
)

func LineLength(start, end *gen.Point) float64 {
	return math.Hypot(start.X-end.X, start.Y-end.Y)
}

func ArcLength(sweep, r float64) float64 {
	return r * math.Abs(sweep)
}

func SegmentLength(s *gen.Segment) float64 {
	if s.Radius != 0 {
		return ArcLength(s.Sweep, s.Radius)
	}
	return LineLength(s.Start, s.End)
}

func ToRad(deg float64) float64 {
	return deg * math.Pi / 180.0
}

func ToDeg(rad float64) float64 {
	return rad * 180.0 / math.Pi
}

func SegmentBulge(s *gen.Segment) float64 {
	return Bulge(s.Radius, s.Amp, s.Sweep)
}

func Bulge(r, amp, sweep float64) float64 {
	if r != 0 {
		bulge := math.Tan(sweep / 4)

		if amp < 0 {
			return -bulge
		}
		return bulge
	}

	return 0
}

func AddAngle(x, y, angle, length float64) (float64, float64) {
	a := ToRad(angle)
	newX := length*math.Cos(a) + x
	newY := length*math.Sin(a) + y

	return newX, newY
}

func Angle(start, end *gen.Point) float64 {
	x := end.X - start.X
	y := end.Y - start.Y

	return math.Atan2(y, x)
}

func SegmentAngle(s *gen.Segment) float64 {
	return Angle(s.Start, s.End)
}

func SegmentMidPoint(s *gen.Segment) (x float64, y float64) {
	if s.Radius != 0.0 {
		x, y = ArcMidPoint(s.Start, s.End, s.Origin, s.Radius)
	} else {
		x, y = LineMidPoint(s.Start, s.End)
	}
	return
}

func MidPoint(start, end, center *gen.Point, radius float64) (x float64, y float64) {
	if radius != 0.0 {
		x, y = ArcMidPoint(start, end, center, radius)
	} else {
		x, y = LineMidPoint(start, end)
	}
	return
}

func LineMidPoint(start, end *gen.Point) (float64, float64) {
	x := start.X + ((end.X - start.X) / 2.0)
	y := start.Y + ((end.Y - start.Y) / 2.0)

	return x, y
}

func ArcMidPoint(start, end, center *gen.Point, radius float64) (float64, float64) {
	a1 := Angle(center, start)
	a2 := Angle(center, end)

	theta := (a1 + a2) / 2.0

	newx := radius*math.Cos(theta) + center.X
	newy := radius*math.Sin(theta) + center.Y

	return newx, newy
}

func IsClockwise(c *gen.Contour) bool {
	var sum float64

	for _, s := range c.Segments {
		sum += (s.End.X - s.Origin.X) * (s.End.Y + s.Origin.Y)
	}

	return sum > 0
}
