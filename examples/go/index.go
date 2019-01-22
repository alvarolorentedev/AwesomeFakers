package main

import (
	"fmt"
)

type SomeStruct struct {
	Int      int
	Int8     int8
	Int16    int16
	Int32    int32
	Int64    int64
	String   string
	Bool     bool
	SString  []string
	SInt     []int
	SInt8    []int8
	SInt16   []int16
	SInt32   []int32
	SInt64   []int64
	SFloat32 []float32
	SFloat64 []float64
	SBool    []bool
	Struct   AStruct
}
type AStruct struct {
	Number        float32 `faker:"lat"`
	Height        float32 `faker:"long"`
}

func SomeAwesomeFunction(aStruct SomeStruct) {
	fmt.Printf("%+v\n", aStruct)
}
