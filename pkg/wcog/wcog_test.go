package wcog

import (
	"io"
	"strings"
	"sync"
	"testing"
)

func Test_filterPos(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{s: "06103-S06103-UZ_ST_993480-4949P"}, "4949"},
		{"", args{s: "06103-S06103-UZ_ST_993481-B4957P"}, "4957"},
		{"", args{s: "06103-S06103-ROS_SEC_884001-C4785P"}, "4785"},
		{"", args{s: "06104-S06104-ROS_SEC_884001-B4831S"}, "4831"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterPos(tt.args.s); got != tt.want {
				t.Errorf("filterPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadWCOG(t *testing.T) {
	type args struct {
		r io.Reader
		m *sync.Map
	}

	csv := `Part name,Weight,X,Y,Z,Panel,Block,Part,Type,Side,Stock Number,Quality,GPS1,GPS2,GPS3,GPS4,Ship,Ident,Nested on,Area,Circ. Length,Circ. Width,Thickness,Shape,Dimension,Total Length,Moulded Length
06103,222427.2,-20602.6,18694.4,8516.6
06103-S06103-PS10-B4101P,1.1,-23244.4,19589.0,12508.9,06103-3PBIM137_3P,06103,06103-3PBIM137_3P-7BP,PLANAR PLATE BRACKET,PS,,A40,,,,,10510,,,0.023,241.2,155.0,6.00000,,,,
06103-S06103-PS10-B4101P,1.1,-23244.4,16109.0,12509.3,06103-3PBIM137_3P,06103,06103-3PBIM137_3P-5BP,PLANAR PLATE BRACKET,PS,,A40,,,,,10510,,,0.022,240.0,155.0,6.00000,,,,
06103-S06103-PS10-8033P,46.4,-23305.0,20807.1,12522.7,06103-3PBIM137_5P,06103,06103-3PBIM137_5P-1P,PLANAR PLATE PLANE,PS,,D,,,,,10510,,,0.592,1922.7,340.0,10.00000,,,,
`

	r := strings.NewReader(csv)
	firstErrReader := strings.NewReader("")
	secondErrReader := strings.NewReader("123\n")

	m := new(sync.Map)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Reader test", args{r: r, m: m}, false},
		{"First error test", args{r: firstErrReader, m: m}, true},
		{"Second error test", args{r: secondErrReader, m: m}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadWCOG(tt.args.r, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("ReadWCOG() error = %v, wantErr %v", err, tt.wantErr)
			}
			v, ok := tt.args.m.Load("061034101")
			if !ok {
				t.Errorf("ReadWCOG() error: key \"061034101\" not found in map")
			}
			if v.(uint) != 2 {
				t.Errorf("ReadWCOG() error: key \"061034101\" has wrong value")
			}
		})
	}
}
