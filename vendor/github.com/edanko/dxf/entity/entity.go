package entity

import (
	"github.com/edanko/dxf/color"
	"github.com/edanko/dxf/format"
	"github.com/edanko/dxf/handle"
	"github.com/edanko/dxf/table"
)

// Entity is interface for DXF Entities.
type Entity interface {
	IsEntity() bool
	Format(format.Formatter)
	Handle() int
	SetHandle(*int)
	SetBlockRecord(handle.Handler)
	Layer() *table.Layer
	SetLayer(*table.Layer)
	SetLtscale(float64)
	BBox() ([]float64, []float64)
	SetColor(color.ColorNumber)
}

// entity is common part of Entities.
// It is embedded in each entities to implement Entity interface.
type entity struct {
	Type        EntityType        // 0
	handle      int               // 5
	blockRecord handle.Handler    // 102 330
	owner       handle.Handler    // 330
	layer       *table.Layer      // 8
	ltscale     float64           // 48
	color       color.ColorNumber // 62
}

// NewEntity creates a new entity.
func NewEntity(t EntityType) *entity {
	e := &entity{
		Type:        t,
		handle:      0,
		blockRecord: nil,
		owner:       nil,
		layer:       table.LY_0,
		ltscale:     1.0,
		color:       0,
	}
	return e
}

func (e *entity) SetColor(cl color.ColorNumber) {
	e.color = cl
}

// Format writes data to formatter.
func (e *entity) Format(f format.Formatter) {
	f.WriteString(0, EntityTypeString(e.Type))
	f.WriteHex(5, e.handle)
	if e.blockRecord != nil {
		f.WriteString(102, "{ACAD_REACTORS")
		f.WriteHex(330, e.blockRecord.Handle())
		f.WriteString(102, "}")
	}
	if e.owner != nil {
		f.WriteHex(330, e.owner.Handle())
	}
	f.WriteString(100, "AcDbEntity")
	f.WriteString(8, e.layer.Name())
	if e.ltscale != 1.0 {
		f.WriteFloat(48, e.ltscale)
	}
	f.WriteInt(62, int(e.color))
}

// String outputs data using default formatter.
func (e *entity) String() string {
	f := format.NewASCII()
	return e.FormatString(f)
}

// FormatString outputs data using given formatter.
func (e *entity) FormatString(f format.Formatter) string {
	e.Format(f)
	return f.Output()
}

// Handle returns a handle value of TABLE.
func (e *entity) Handle() int {
	return e.handle
}

// SetHandle sets handles to TABLE itself and each SymbolTable.
func (e *entity) SetHandle(v *int) {
	e.handle = *v
	(*v)++
}

// SetBlockRecord sets BLOCK_RECORD to entity (code 330).
func (e *entity) SetBlockRecord(h handle.Handler) {
	e.blockRecord = h
}

// SetOwner sets an owner.
func (e *entity) SetOwner(h handle.Handler) {
	e.owner = h
}

// Layer returns entity's Layer.
func (e *entity) Layer() *table.Layer {
	return e.layer
}

// SetLayer sets Layer to entity.
func (e *entity) SetLayer(l *table.Layer) {
	e.layer = l
}

// SetLtscale sets Layer to entity.
func (e *entity) SetLtscale(v float64) {
	e.ltscale = v
}

// SetEntityType sets entity type.
func (e *entity) SetEntityType(t EntityType) {
	e.Type = t
}
