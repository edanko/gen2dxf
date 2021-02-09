package draw

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/edanko/gen2dxf/pkg/gen"
)

/*
	0 NONE
	22 C25d-R
	103 C7-AF
	104 C21-AF
	113 C7-R
	114 C8-R
	116 C12-R
	117 C15-R
	121 C21-R
	122 C25-R
	125 C39-R
	133 C62-P3
	143 T3-R
	144 T6-R
	145 T7-R
	146 T8-R
	150 T12-R
	153 T20-R
	154 T21-R
	155 T22-R
	172 U6-R
	173 U7-R
	177 U30-R
	180 U40-R
	181 U41-R
	183 U43-R
	203 C7e-AF
	204 C21e-AF
	205 C21e-AFh
	206 C25e-AF
	207 C25e-AFh
	215 C10e-R
	216 C12e-R
	217 C15e-R
	219 C19e-R
	221 C21e-R
	222 C25e-K
	227 C43e-R
	233 C62e-P3
	243 T3E-R
	244 T6E-R
	245 T7E-R
	246 T8E-R
	248 T10E-R
	253 T20E-R
	254 T21e-R
	255 T22e-R
	273 U7e-R
	274 U8e-R
	277 C30e-R
	293 U7d-R
	305 C25n-AF
	310 C12n-R
	312 C15n-P
	315 C21n-R
	318 C25n-P
	331 C70n-R
	346 U7n-R
	348 U8n-R
	391 T6n-R
	392 T7n-R
	393 T8n-R
	616 C12-R-
	716 C12e-R-
	717 C15e-R-
	727 C43e-R-
	103106 C7-AF
	103110 C7-AF
	103113 C7-AF
	103210 C7-AF
	104210 C21-AF
	121110 C21-R
	121214 C21-R
	121413 C21-R
	203112 C7e-AF + ласка
	206122 C25e-AF + ласка
	207120 C25e-AFh
	207140 C25e-AFh
	207338 C25e-AFh
	222120 C25e-K
	222132 C25e-K

	-103 C7-AFO
	-104 C21-AFO
	-107 C25-AFkO
	-113 C7-R
	-114 C8-RO
	-116 C12-RO
	-121 C21-RO
	-122 C25-RO
	-125 C39-RO
	-127 C43-RO
	-133 C62-P3O
	-135 C64-P3O
	-143 T3-R
	-144 T6-RO
	-145 T7-RO
	-146 T8-RO
	-153 T20-R
	-154 T21-R
	-155 T22-R
	-172 U6-RO
	-173 U7-RO
	-177 U30-R
	-203 C7e-AF
	-204 C21e-AFO
	-205 C21e-AFhO
	-206 C25e-AF
	-207 C25e-AFhO
	-216 C12e-RO
	-221 C21e-RO
	-222 C25e-KO
	-227 C43e-RO
	-233 C62e-P3O
	-234 C63e-P3O
	-243 T3E-R
	-244 T6E-RO
	-245 T7E-RO
	-246 T8E-R
	-248 T10E-RO
	-253 T20E-R
	-254 T21e-R
	-255 T22e-R
	-273 U7e-RO
	-277 C30e-R
	-315 C21n-R
	-318 C25n-P
	-346 U7n-R
	-348 U8n-R
	-391 T6n-RO
	-392 T7n-RO
	-393 T8n-RO
	-398 T13n-P3O
	-716 C12e-R-
	-728 C43e-R-
	-734 C63e-P3-O
	-121214 C21-RO
	-221213 C21e-RO
	-221214 C21e-RO
	-221413 C21e-RO
*/

func Bevel(b *gen.BevelData) (name, nestix string) {
	if b == nil {
		return "", ""
	}

	switch b.BevelCode {
	// NONE
	case -728, 727, 253, 103113, 103110, 103106, 103210, 203112, 0, 180, -253, -153, 113, -113, 153, -154, 181, 277, 717, 616, 177, -177, 143, -143, -243, 243, -203, 716, -716, -103, 203, -227, 227, 254, -254, -277, 255, -255, 103, 183:
		name = ""
		nestix = ""

	// TS
	case 331:
		name = "TSV15"
		nestix = fmt.Sprintf("[V/15.0][S%.1f]", b.PlateThickness)

	case 233, 133:
		name = "TSV20"
		nestix = fmt.Sprintf("[V/20.0][S%.1f]", b.PlateThickness)

	case 221, 121, 219, 315, 121413, 121214, 121110:
		name = "TSV25"
		nestix = fmt.Sprintf("[V/25.0][S%.1f]", b.PlateThickness)

	case 116, 145, 173, 114, 144, 172:
		name = "TSV45"
		nestix = fmt.Sprintf("[V/45.0][S%.1f]", b.PlateThickness)

	case 245, 273, 216, 248, 215, 293, 310, 392, 346, 244:
		name = "TSV50"
		nestix = fmt.Sprintf("[V/45.0][S%.1f]", b.PlateThickness)

	case 205:
		name = fmt.Sprintf("TSY25/%g", b.DepthTS)
		nestix = fmt.Sprintf("[Y/25.0;%g][S%.1f]", b.DepthTS, b.PlateThickness)

	case 204, 104, 104210:
		name = fmt.Sprintf("TSY30/%g", b.DepthTS)
		nestix = fmt.Sprintf("[Y/30.0;%g][S%.1f]", b.DepthTS, b.PlateThickness)

	case 391:
		name = fmt.Sprintf("TSY50/%g", b.DepthTS)
		nestix = fmt.Sprintf("[Y/45.0;%g][S%.1f]", b.DepthTS, b.PlateThickness)

	// OS
	case -234:
		name = "OSV10"
		nestix = fmt.Sprintf("[V\\10.0][S%.1f]", b.PlateThickness)

	case -133, -233:
		name = "OSV20"
		nestix = fmt.Sprintf("[V\\20.0][S%.1f]", b.PlateThickness)

	case -121, -221, -221213, -221214, -221413, -121214:
		name = "OSV25"
		nestix = fmt.Sprintf("[V\\25.0][S%.1f]", b.PlateThickness)

	case -734:
		name = "OSV35"
		nestix = fmt.Sprintf("[V\\35.0][S%.1f]", b.PlateThickness)

	case -116, -145, -173, -144, -114, -172:
		name = "OSV45"
		nestix = fmt.Sprintf("[V\\45.0][S%.1f]", b.PlateThickness)

	case -245, -273, -216, -248, -346, -392, -244:
		name = "OSV50"
		nestix = fmt.Sprintf("[V\\45.0][S%.1f]", b.PlateThickness)

	case -315, -205:
		name = fmt.Sprintf("OSY25/%g", b.PlateThickness-b.DepthOS)
		nestix = fmt.Sprintf("[Y\\25.0;%g][S%.1f]", b.PlateThickness-b.DepthOS, b.PlateThickness)

	case -104, -204:
		name = "OSY30/7"
		nestix = fmt.Sprintf("[Y\\30.0;7.0][S%.1f]", b.PlateThickness)

	case -391, -398:
		name = fmt.Sprintf("OSY50/%g", b.PlateThickness-b.DepthTS)
		nestix = fmt.Sprintf("[Y\\45.0;%g][S%.1f]", b.PlateThickness-b.DepthOS, b.PlateThickness)

	//BOTH
	case -222, 222, 206, -206, 122, -122, -207, -135, 207, 207338, 206122, 207140, 207120, 222132, 222120:
		name = "X25"
		//nestix = fmt.Sprintf("[X25;%.1f;25][S%.1f]", b.PlateThickness/2.0, b.PlateThickness)

	case -107:
		if b.PlateThickness > 38 {
			name = "K25-2"
		} else {
			name = "K30-2"
		}

	case -318, 318:
		name = "X27"

	case 305:
		name = "K27-2"

	case 146, -146, 117, 150:
		name = "X45"

	case 246, -246, 155, 274, -155, 22, 217:
		name = "X50"

	case -348, 348, 393, -393:
		name = "K50-4"

	case 312:
		name = "K50-2"

	default:
		name = ""
		nestix = ""
		spew.Dump(b)
	}

	if b.ChamferHeightTS != 0.0 {
		name = fmt.Sprintf("TSL%g-%g (%s)", b.ChamferWidthTS, b.PlateThickness-b.ChamferHeightTS, name)
		nestix = ""
	}
	if b.ChamferHeightOS != 0.0 {
		name = fmt.Sprintf("OSL%g-%g (%s)", b.ChamferWidthOS, b.PlateThickness-b.ChamferHeightOS, name)
		nestix = ""
	}

	return
}
