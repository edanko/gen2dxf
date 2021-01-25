package table

import (
	"github.com/edanko/dxf/format"
	"github.com/edanko/dxf/handle"
)

// SymbolTable is interface for AcDbSymbolTableRecord.
type SymbolTable interface {
	IsSymbolTable() bool
	Format(format.Formatter)
	Handle() int
	SetHandle(*int)
	SetOwner(handle.Handler)
	Name() string
}
