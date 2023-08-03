package utils

import "C"
import "gitea.linesip.com/libraries/gopipewire/spa/pod"

const (
	SPATypeInfoDirection     = SPATypeInfoEnum + "Direction"
	SPATypeInfoDirectionBase = SPATypeInfoDirection + ":"
	SPATypeInfoChoice        = SPATypeInfoEnumBase + "Choice"
	SPATypeInfoChoiceBase    = SPATypeInfoChoice + ":"
)

var (
	SPATypeDirection = [...]SPATypeInfo{
		{
			Type:   uint32(SPADirectionInput),
			Parent: uint32(SPATypeEnumInt),
			Name:   C.CString(string(SPATypeInfoDirectionBase + "Input")),
			Values: nil,
		},
		{
			Type:   uint32(SPADirectionOutput),
			Parent: uint32(SPATypeEnumInt),
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

	SPATypeChoice = [...]SPATypeInfo{
		{pod.SPAChoiceNone},
		{},
		{},
		{},
		{},
		{
			Type:   0,
			Parent: 0,
			Name:   nil,
			Values: nil,
		},
	}
)
