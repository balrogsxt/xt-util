package xvar

import (
	"fmt"
	"strconv"
)

type Var struct {
	v any
}

func New(v any) *Var {
	return &Var{
		v: v,
	}
}
func (l *Var) Int() int {
	return int(l.Int64())
}
func (l *Var) Int8() int {
	v, _ := strconv.ParseInt(l.String(), 10, 8)
	return int(v)
}
func (l *Var) Int16() int16 {
	v, _ := strconv.ParseInt(l.String(), 10, 16)
	return int16(v)
}
func (l *Var) Int32() int32 {
	v, _ := strconv.ParseInt(l.String(), 10, 32)
	return int32(v)
}
func (l *Var) Int64() int64 {
	v, _ := strconv.ParseInt(l.String(), 10, 64)
	return v
}

func (l *Var) Float32() float32 {
	v, _ := strconv.ParseFloat(l.String(), 32)
	return float32(v)
}
func (l *Var) Float64() float64 {
	v, _ := strconv.ParseFloat(l.String(), 64)
	return v
}
func (l *Var) Uint() int {
	return int(l.Uint64())
}
func (l *Var) Uint8() uint {
	v, _ := strconv.ParseUint(l.String(), 10, 8)
	return uint(v)
}
func (l *Var) Uint16() uint16 {
	v, _ := strconv.ParseUint(l.String(), 10, 16)
	return uint16(v)
}
func (l *Var) Uint32() uint32 {
	v, _ := strconv.ParseUint(l.String(), 10, 32)
	return uint32(v)
}
func (l *Var) Uint64() uint64 {
	v, _ := strconv.ParseUint(l.String(), 10, 64)
	return v
}

func (l *Var) Bool() bool {
	v, _ := strconv.ParseBool(l.String())
	return v
}

func (l *Var) String() string {
	if l.v == nil {
		return ""
	}
	return fmt.Sprintf("%v", l.v)
}

func (l *Var) Any() any {
	return l.v
}
