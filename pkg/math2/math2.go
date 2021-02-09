package math2

import (
	"math"

	"github.com/edanko/gen2dxf/pkg/gen"
)

func LineLength(sx, sy, ex, ey float64) float64 {
	return math.Hypot(sx-ex, sy-ey)
}

func ArcLength(sweep, r float64) float64 {
	return r * math.Abs(sweep)
}

func SegmentLength(s *gen.Segment) float64 {
	if s.Radius != 0 {
		return ArcLength(s.Sweep, s.Radius)
	}
	return LineLength(s.StartU, s.StartV, s.U, s.V)
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
	} else {
		return 0
	}
}

func AddAngle(x, y, angle, length float64) (float64, float64) {
	a := ToRad(angle)
	newX := length*math.Cos(a) + x
	newY := length*math.Sin(a) + y

	return newX, newY
}

func Angle(sx, sy, ex, ey float64) float64 {
	x := ex - sx
	y := ey - sy

	return math.Atan2(y, x)
}

func SegmentAngle(s *gen.Segment) float64 {
	return Angle(s.StartU, s.StartV, s.U, s.V)
}

func SegmentMidPoint(s *gen.Segment) (x float64, y float64) {
	if s.Radius != 0.0 {
		x, y = ArcMidPoint(s.StartU, s.StartV, s.U, s.V, s.OriginU, s.OriginV, s.Radius)
	} else {
		x, y = LineMidPoint(s.StartU, s.StartV, s.U, s.V)
	}
	return
}

func MidPoint(sx, sy, ex, ey, cx, cy, r float64) (x float64, y float64) {
	if r != 0.0 {
		x, y = ArcMidPoint(sx, sy, ex, ey, cx, cy, r)
	} else {
		x, y = LineMidPoint(sx, sy, ex, ey)
	}
	return
}

func LineMidPoint(sx, sy, ex, ey float64) (float64, float64) {
	x := sx + ((ex - sx) / 2.0)
	y := sy + ((ey - sy) / 2.0)

	return x, y
}

func ArcMidPoint(sx, sy, ex, ey, cx, cy, r float64) (float64, float64) {
	a1 := Angle(cx, cy, sx, sy)
	a2 := Angle(cx, cy, ex, ey)

	theta := (a1 + a2) / 2.0

	newx := r*math.Cos(theta) + cx
	newy := r*math.Sin(theta) + cy

	return newx, newy
}
