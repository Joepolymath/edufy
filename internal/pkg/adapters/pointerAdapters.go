package adapters

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PointerInterface interface {
	StringPointer(value string) *string
	IntPointer(value int) *int
	UIntPointer(value uint) *uint
	DecimalPointer(value decimal.Decimal) *decimal.Decimal
	BoolPointer(value bool) *bool
	Int64Pointer(value int64) *int64
	UUIDPointer(value uuid.UUID) *uuid.UUID
	Float64Pointer(value float64) *float64
}

type PointerStruct struct {
}

func NewPointer() PointerInterface {
	return &PointerStruct{}
}

func (p *PointerStruct) StringPointer(value string) *string {
	return &value
}

func (p *PointerStruct) IntPointer(value int) *int {
	return &value
}

func (p *PointerStruct) UIntPointer(value uint) *uint {
	return &value
}

func (p *PointerStruct) DecimalPointer(value decimal.Decimal) *decimal.Decimal {
	return &value
}

func (p *PointerStruct) BoolPointer(value bool) *bool {
	return &value
}

func (p *PointerStruct) Int64Pointer(value int64) *int64 {
	return &value
}

func (p *PointerStruct) Float64Pointer(value float64) *float64 {
	return &value
}

func (p *PointerStruct) UUIDPointer(value uuid.UUID) *uuid.UUID {
	return &value
}
