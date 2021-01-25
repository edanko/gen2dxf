package entity

import (
	"github.com/edanko/dxf/format"
)

type LwPolylineVertex struct {
	X     float64
	Y     float64
	Bulge float64
}

// LwPolyline represents LWPOLYLINE Entity.
type LwPolyline struct {
	*entity
	Num      int
	Closed   bool
	Vertices []*LwPolylineVertex
}

// IsEntity is for Entity interface.
func (l *LwPolyline) IsEntity() bool {
	return true
}

// NewLwPolyline creates a new LwPolyline.
func NewLwPolyline() *LwPolyline {
	vs := make([]*LwPolylineVertex, 0)
	l := &LwPolyline{
		entity:   NewEntity(LWPOLYLINE),
		Num:      0,
		Closed:   false,
		Vertices: vs,
	}
	return l
}

func (l *LwPolyline) AddVertex(x, y, b float64) *LwPolylineVertex {
	v := &LwPolylineVertex{X: x, Y: y, Bulge: b}
	l.Vertices = append(l.Vertices, v)
	l.Num++
	return v
}

// Format writes data to formatter.
func (l *LwPolyline) Format(f format.Formatter) {
	l.entity.Format(f)
	f.WriteString(100, "AcDbPolyline")
	f.WriteInt(90, l.Num)
	if l.Closed {
		f.WriteInt(70, 1)
	} else {
		f.WriteInt(70, 0)
	}
	for i := 0; i < l.Num; i++ {
		f.WriteFloat(10, l.Vertices[i].X)
		f.WriteFloat(20, l.Vertices[i].Y)
		f.WriteFloat(42, l.Vertices[i].Bulge)
	}
}

// String outputs data using default formatter.
func (l *LwPolyline) String() string {
	f := format.NewASCII()
	return l.FormatString(f)
}

// FormatString outputs data using given formatter.
func (l *LwPolyline) FormatString(f format.Formatter) string {
	l.Format(f)
	return f.Output()
}

// Close closes LwPolyline.
func (l *LwPolyline) Close() {
	l.Closed = true
}

func (l *LwPolyline) BBox() ([]float64, []float64) {
	mins := make([]float64, 3)
	maxs := make([]float64, 3)
	for i, p := range l.Vertices {
		//for i := 0; i < 2; i++ {
		if p.X < mins[i] {
			mins[i] = p.X
		}
		if p.Y < mins[i] {
			mins[i] = p.Y
		}
		if p.X > maxs[i] {
			maxs[i] = p.X
		}
		if p.Y > maxs[i] {
			maxs[i] = p.Y
		}
		//}
	}
	return mins, maxs
}
