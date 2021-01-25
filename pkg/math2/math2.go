package math2

import (
	"math"

	"github.com/edanko/gen2dxf/pkg/gen"
)

func LineLength(sx, sy, ex, ey float64) float64 {
	x := math.Pow((sx - ex), 2)
	y := math.Pow((sy - ey), 2)

	return math.Sqrt(x + y)
}

func DegToRad(x float64) float64 {
	return x * math.Pi / 180.0
}

func RadToDeg(x float64) float64 {
	return x * 180.0 / math.Pi
}

func GetSegmentBulge(s *gen.Segment) float64 {
	return GetBulge(s.Radius, s.Amp, s.Sweep)
}

func GetBulge(r, amp, sweep float64) float64 {
	if r != 0 { //&& r < 20000 {
		bulge := math.Tan(sweep / 4)

		if amp < 0 {
			return -bulge
		}
		return bulge
	} else {
		return 0
	}
}

/*func arcToBulge(sx, sy, ex, ey, cx, cy, r, amp, sweep float64) float64 {
	if r != 0 && r < 20000 {
		//angle1 := getAngle(cx, cy, sx, sy)
		//angle2 := getAngle(cx, cy, ex, ey)

		//phi := angle2 - angle1
		//phi = math.Mod((math.Pi*2 + phi), math.Pi*2)

		bulge := math.Tan(sweep / 4)

		if amp < 0 {
			return -bulge
		}
		return bulge
	} else {
		return 0
	}
}*/

// adds length to x,y coords by angle
func AddAngle(x, y, angle, length float64) (float64, float64) {
	angle = angle * math.Pi / 180.0
	newX := length*math.Cos(angle) + x
	newY := length*math.Sin(angle) + y

	return newX, newY
}

func GetAngle(sx, sy, ex, ey float64) float64 {
	x := ex - sx
	y := ey - sy

	return math.Atan2(y, x)
}

func GetSegmentAngle(s *gen.Segment) float64 {
	return GetAngle(s.StartU, s.StartV, s.U, s.V)
}

func GetSegmentMidPoint(s *gen.Segment) (x float64, y float64) {
	if s.Radius != 0.0 {
		x, y = ArcMidPoint(s.StartU, s.StartV, s.U, s.V, s.OriginU, s.OriginV, s.Radius)
	} else {
		x, y = LineMidPoint(s.StartU, s.StartV, s.U, s.V)
	}
	return
}

func GetMidPoint(sx, sy, ex, ey, cx, cy, r float64) (x float64, y float64) {
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
	x1 := sx - cx
	y1 := sy - cy

	x2 := ex - cx
	y2 := ey - cy

	theta1 := math.Atan2(y1, x1)
	theta2 := math.Atan2(y2, x2)

	theta := (theta1 + theta2) / 2.0

	newx := r*math.Cos(theta) + cx
	newy := r*math.Sin(theta) + cy

	return newx, newy
}
