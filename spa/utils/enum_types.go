package utils

import "C"

const (
	SPATypeInfoDirection     = SPATypeInfoEnum + "Direction"
	SPATypeInfoDirectionBase = SPATypeInfoDirection + ":"
)

var (
	SPATypeDirection = [...]SPATypeInfo{
		{
			Type:   uint32(SPADirectionInput),
			Parent: uint32(SPATypeInt),
			Name:   C.CString(string(SPATypeInfoDirectionBase + "Input")),
			Values: nil,
		},
		{
			Type:   uint32(SPADirectionOutput),
			Parent: uint32(SPATypeInt),
			Name:   C.CString(string(SPATypeInfoDirectionBase + "Output")),
			Values: &SPATypeInfo{},
		},
		{
			Type:   0,
			Parent: 0,
			Name:   nil,
			Values: nil,
		},
	}
)
