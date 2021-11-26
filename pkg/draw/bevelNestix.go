package draw

import (
	"fmt"
	"math"

	"gen2dxf/pkg/gen"
	"gen2dxf/pkg/math2"
)

func BevelNestix(b *gen.BevelData) (name, nestix string) {
	if b == nil {
		return "", ""
	}

	var t float64

	if b.E == -1 {
		t = b.PlateThickness
	} else {
		t = b.PlateThickness - b.E
	}

	switch {
	case b.AngleTS != 0 && b.AngleOS == 0 && (t-b.DepthTS <= 2 || b.DepthTS == 0):

		name = fmt.Sprintf("TSV%g", b.AngleTS)
		nestix = fmt.Sprintf("[V/%g][S%.1f]", b.AngleTS, t)

	case b.AngleOS != 0 && b.AngleTS == 0 && (t-b.DepthOS <= 2 || b.DepthOS == 0):

		name = fmt.Sprintf("OSV%g", b.AngleOS)
		nestix = fmt.Sprintf("[V\\%g][S%.1f]", b.AngleOS, t)

	case b.AngleTS != 0 && b.AngleOS == 0 && t-b.DepthTS > 2:

		name = fmt.Sprintf("TSY%g/%g", b.AngleTS, b.DepthTS)
		nestix = fmt.Sprintf("[Y/%g;%g][S%.1f]", b.AngleTS, b.DepthTS, t)

	case b.AngleOS != 0 && b.AngleTS == 0 && t-b.DepthOS > 2:

		name = fmt.Sprintf("OSY%g/%g", b.AngleOS, t-b.DepthOS)
		nestix = fmt.Sprintf("[Y\\%g;%g][S%.1f]", b.AngleOS, t-b.DepthOS, t)

	case b.AngleTS != 0 && b.AngleOS != 0:

		if b.PlateThickness-(b.DepthTS+b.DepthOS) <= 2 {

			name = fmt.Sprintf("X%g", b.AngleTS)
			nestix = fmt.Sprintf("[X%g;%.1f;%g][S%.1f]", b.AngleTS, t/2.0, b.AngleOS, t)

		} else {

			var depthts float64
			if b.E == -1 {
				depthts = b.DepthOS
			} else {
				depthts = b.PlateThickness - b.DepthTS
			}

			name = fmt.Sprintf("K%g-%g", b.AngleTS, b.PlateThickness-(b.DepthTS+b.DepthOS))
			nestix = fmt.Sprintf("[K%g;%.1f;%g;%g][S%.1f]", b.AngleTS, depthts, b.PlateThickness-(b.DepthTS+b.DepthOS), b.AngleOS, t)

		}
	}

	if b.ChamferHeightTS != 0.0 {
		angle := math2.ToDeg(math.Atan(b.ChamferWidthTS / (b.PlateThickness - b.ChamferHeightTS)))

		name = fmt.Sprintf("TSL%.0f-%g (%s)", angle, b.PlateThickness-b.ChamferHeightTS, name)
		nestix = ""
	}
	if b.ChamferHeightOS != 0.0 {
		angle := math2.ToDeg(math.Atan(b.ChamferWidthOS / (b.PlateThickness - b.ChamferHeightOS)))

		name = fmt.Sprintf("OSL%.0f-%g (%s)", angle, b.PlateThickness-b.ChamferHeightOS, name)
		nestix = ""
	}

	return
}
