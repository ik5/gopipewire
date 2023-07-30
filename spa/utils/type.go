package utils

import "C"

type SPAType uint32

// SPA Types - Data type information enumerations
const (
	// Basic types

	SPATypeStart SPAType = iota
	SPATypeNone
	SPATypeBool
	SPATypeID
	SPATypeInt
	SPATypeLong
	SPATypeFloat
	SPATypeDouble
	SPATypeString
	SPATypeBytes
	SPATypeRectangle
	SPATypeFraction
	SPATypeBitmap
	SPATypeArray
	SPATypeStruct
	SPATypeObject
	SPATypeSequence
	SPATypePointer
	SPATypeFd
	SPATypeChoice
	SPATypePod
	SPATypeLast /**< not part of ABI */

	/* Pointers */

	SPATypePointerStart SPAType = iota + 0x10000
	SPATypePointerBuffer
	SPATypePointerMeta
	SPATypePointerDict
	SPATypePointerLast /**< not part of ABI */

	/* Events */

	SPATypeEventStart SPAType = iota + 0x20000
	SPATypeEventDevice
	SPATypeEventNode
	SPATypeEventLast /**< not part of ABI */

	/* Commands */

	SPATypeCommandStart SPAType = iota + 0x30000
	SPATypeCommandDevice
	SPATypeCommandNode
	SPATypeCommandLast /**< not part of ABI */

	/* Objects */

	SPATypeObjectStart SPAType = iota + 0x40000
	SPATypeObjectPropInfo
	SPATypeObjectProps
	SPATypeObjectFormat
	SPATypeObjectParamBuffers
	SPATypeObjectParamMeta
	SPATypeObjectParamIO
	SPATypeObjectParamProfile
	SPATypeObjectParamPortConfig
	SPATypeObjectParamRoute
	SPATypeObjectProfiler
	SPATypeObjectParamLatency
	SPATypeObjectParamProcessLatency
	SPATypeObjectLast /**< not part of ABI */

	/* vendor extensions */

	SPATypeVendorPipeWire SPAType = 0x02000000

	SPATypeVendorOther SPAType = 0x7f000000
)

type SPAStringType string

// SPA Type Information
const (
	SPATypeInfoBase          SPAStringType = "Spa:"
	SPATypeInfoFlags                       = SPATypeInfoBase + "Flags"
	SPATypeInfoFlagsBase                   = SPATypeInfoFlags + ":"
	SPATypeInfoEnum                        = SPATypeInfoBase + "Enum"
	SPATypeInfoEnumBase                    = SPATypeInfoEnum + ":"
	SPATypeInfoPod                         = SPATypeInfoBase + "Pod"
	SPATypeInfoPodBase                     = SPATypeInfoPod + ":"
	SPATypeInfoStruct                      = SPATypeInfoBase + "Struct"
	SPATypeInfoStructBase                  = SPATypeInfoStruct + ":"
	SPATypeInfoObject                      = SPATypeInfoBase + "Object"
	SPATypeInfoObjectBase                  = SPATypeInfoObject + ":"
	SPATypeInfoPointer                     = SPATypeInfoBase + "Pointer"
	SPATypeInfoPointerBase                 = SPATypeInfoPointer + ":"
	SPATypeInfoInterface                   = SPATypeInfoBase + "Interface"
	SPATypeInfoInterfaceBase               = SPATypeInfoInterface + ":"
	SPATypeInfoEvent                       = SPATypeInfoBase + "Event"
	SPATypeInfoEventBase                   = SPATypeInfoEvent + ":"
	SPATypeInfoCommand                     = SPATypeInfoBase + "Command"
	SPATypeInfoCommandBase                 = SPATypeInfoCommand + ":"
)

type SPATypeInfo struct {
	Type   uint32
	Parent uint32
	Name   *C.char
	Values *SPATypeInfo
}

type SPARectangle struct {
	width  int32
	height int32
}

type SPAPoint struct {
	x int32
	y int32
}
