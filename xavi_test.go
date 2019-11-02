package xavi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPass(t *testing.T) {
	{
		a := sampleA()
		var dstA DstA
		if assert.NoError(t, Pass(&dstA, a)) {
			assert.Equal(t, a.ID, dstA.ID)
			assert.Equal(t, a.Name, dstA.Name)
			assert.Equal(t, a.Bool, dstA.Bool)
			assert.Equal(t, a.I8, dstA.I8)
			assert.Equal(t, a.I16, dstA.I16)
			assert.Equal(t, a.I32, dstA.I32)
			assert.Equal(t, a.I64, dstA.I64)
			assert.Equal(t, a.UI, dstA.UI)
			assert.Equal(t, a.UI8, dstA.UI8)
			assert.Equal(t, a.UI16, dstA.UI16)
			assert.Equal(t, a.UI32, dstA.UI32)
			assert.Equal(t, a.UI64, dstA.UI64)
			assert.Equal(t, a.F32, dstA.F32)
			assert.Equal(t, a.F64, dstA.F64)

			assert.Equal(t, a.Sub.Name, dstA.Sub.Name)
			assert.Equal(t, a.SubA.Name, dstA.SubA.Name)
		}
	}

	{
		b := sampleB()
		var dstB DstB
		if assert.NoError(t, Pass(&dstB, b)) {
			assert.Equal(t, b.BID, dstB.ID)
			assert.Equal(t, b.BName, dstB.Name)
			assert.Equal(t, b.BStruct.Name, dstB.Struct.Name)
		}
	}
	{
		b := sampleB()
		var dstB DstB
		if assert.NoError(t, Pass(&dstB, &b)) {
			assert.Equal(t, b.BID, dstB.ID)
			assert.Equal(t, b.BName, dstB.Name)
			assert.Equal(t, b.BStruct.Name, dstB.Struct.Name)
		}
	}

	{
		c := sampleC()
		var dstC DstC
		if assert.NoError(t, Pass(&dstC, c)) {
			assert.NotEqual(t, c.Sub.Name, dstC.Sub.Name)
		}
	}

	//Error
	{
		a := sampleA()
		var dstA DstA
		assert.Error(t, Pass(dstA, a))
	}

	{
		var a *A
		var dstA DstA
		assert.Error(t, Pass(&dstA, a))
	}

	{
		a := sampleA()
		var ss string
		assert.Error(t, Pass(&ss, a))
	}
	{
		s := "hoge"
		var dstA DstA
		assert.Error(t, Pass(&dstA, s))
	}
}

type A struct {
	ID   int
	Name string
	Bool bool
	I8   int8
	I16  int16
	I32  int32
	I64  int64
	UI   uint
	UI8  uint8
	UI16 uint16
	UI32 uint32
	UI64 uint64
	F32  float32
	F64  float64
	Sub  SubA
	SubA
}

type SubA struct {
	Name string
}

func sampleA() A {
	return A{
		ID:   1,
		Name: "A",
		Bool: true,
		I8:   8, I16: 16, I32: 32, I64: 64,
		UI: 1, UI8: 8, UI16: 16, UI32: 32, UI64: 64,
		F32: 3.2, F64: 6.4,
		Sub: SubA{
			Name: "sub",
		},
		SubA: SubA{
			Name: "subA",
		},
	}
}

type DstA struct {
	ID   int
	Name string
	Bool bool
	I8   int8
	I16  int16
	I32  int32
	I64  int64
	UI   uint
	UI8  uint8
	UI16 uint16
	UI32 uint32
	UI64 uint64
	F32  float32
	F64  float64
	Sub  SubA
	SubA
}

type B struct {
	BID     int64  `xavi:"id"`
	BName   string `xavi:"name"`
	BStruct SubB   `xavi:"subB"`
}

type SubB struct {
	Name string
}

func sampleB() B {
	return B{
		BID:   1,
		BName: "B",
		BStruct: SubB{
			Name: "subB",
		},
	}
}

type DstB struct {
	ID     int64  `xavi:"id"`
	Name   string `xavi:"name"`
	Struct SubB   `xavi:"subB"`
}

type C struct {
	Sub SubC `xavi:"sub"`
}

type SubC struct {
	Name string
}

func sampleC() C {
	return C{
		Sub: SubC{
			Name: "sub",
		},
	}
}

type DstC struct {
	Sub DstSubC `xavi:"sub"`
}

type DstSubC struct {
	Name string
}
