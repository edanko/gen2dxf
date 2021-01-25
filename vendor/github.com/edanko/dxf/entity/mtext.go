package entity

import (
	"github.com/edanko/dxf/format"
	"github.com/edanko/dxf/table"
)

// MText represents MTEXT Entity.
type MText struct {
	*entity
	Coord1         []float64    // 10, 20, 30
	Coord2         []float64    // 11, 21, 31
	Height         float64      // 40
	Rotation       float64      // 50
	WidthFactor    float64      // 41
	ObliqueAngle   float64      // 51
	Value          string       // 1
	Style          *table.Style // 7
	GenFlag        int          // 71
	HorizontalFlag int          // 72
	VerticalFlag   int          // 73
}

// IsEntity is for Entity interface.
func (t *MText) IsEntity() bool {
	return true
}

// NewMText creates a new MText.
func NewMText() *MText {
	t := &MText{
		entity:         NewEntity(MTEXT),
		Coord1:         []float64{0.0, 0.0, 0.0},
		Coord2:         []float64{0.0, 0.0, 0.0},
		Height:         1.0,
		Value:          "",
		Style:          table.ST_STANDARD,
		GenFlag:        0,
		HorizontalFlag: 0,
		VerticalFlag:   0,
	}
	return t
}

// Format writes data to formatter.
func (t *MText) Format(f format.Formatter) {
	t.entity.Format(f)
	f.WriteString(100, "AcDbMText")
	for i := 0; i < 3; i++ {
		f.WriteFloat((i+1)*10, t.Coord1[i])
	}
	f.WriteFloat(40, t.Height)
	f.WriteFloat(50, t.Rotation)
	f.WriteFloat(41, t.WidthFactor)
	f.WriteFloat(51, t.ObliqueAngle)
	f.WriteString(1, t.Value)
	f.WriteString(7, t.Style.Name())
	if t.GenFlag != 0 {
		f.WriteInt(71, t.GenFlag)
	}
	if t.HorizontalFlag != 0 {
		f.WriteInt(72, t.HorizontalFlag)
		if t.VerticalFlag != 0 {
			for i := 0; i < 3; i++ {
				f.WriteFloat((i+1)*10+1, t.Coord1[i])
			}
		}
	}
	f.WriteString(100, "AcDbMText")
	if t.VerticalFlag != 0 {
		f.WriteInt(73, t.VerticalFlag)
	}
}

// String outputs data using default formatter.
func (t *MText) String() string {
	f := format.NewASCII()
	return t.FormatString(f)
}

// FormatString outputs data using given formatter.
func (t *MText) FormatString(f format.Formatter) string {
	t.Format(f)
	return f.Output()
}

// Anchor sets anchor point flags.
func (t *MText) Anchor(pos int) {
	switch pos {
	case LEFT_BASE:
		t.HorizontalFlag = 0
		t.VerticalFlag = 0
	case CENTER_BASE:
		t.HorizontalFlag = 1
		t.VerticalFlag = 0
	case RIGHT_BASE:
		t.HorizontalFlag = 2
		t.VerticalFlag = 0
	case LEFT_BOTTOM:
		t.HorizontalFlag = 0
		t.VerticalFlag = 1
	case CENTER_BOTTOM:
		t.HorizontalFlag = 1
		t.VerticalFlag = 1
	case RIGHT_BOTTOM:
		t.HorizontalFlag = 2
		t.VerticalFlag = 1
	case LEFT_CENTER:
		t.HorizontalFlag = 0
		t.VerticalFlag = 2
	case CENTER_CENTER:
		t.HorizontalFlag = 1
		t.VerticalFlag = 2
	case RIGHT_CENTER:
		t.HorizontalFlag = 2
		t.VerticalFlag = 2
	case LEFT_TOP:
		t.HorizontalFlag = 0
		t.VerticalFlag = 3
	case CENTER_TOP:
		t.HorizontalFlag = 1
		t.VerticalFlag = 3
	case RIGHT_TOP:
		t.HorizontalFlag = 2
		t.VerticalFlag = 3
	}
}

func (t *MText) BBox() ([]float64, []float64) {
	// TODO: text length, anchor point
	mins := []float64{t.Coord1[0], t.Coord1[1], t.Coord1[2]}
	maxs := []float64{t.Coord1[0], t.Coord1[1] + t.Height, t.Coord1[2]}
	return mins, maxs
}
