package utils

import "C"

type SPAType uint32

// SPA Types - Data type information enumerations
const (
	// Basic types

	SPATypeEnumStart SPAType = iota
	SPATypeEnumNone
	SPATypeEnumBool
	SPATypeEnumID
	SPATypeEnumInt
	SPATypeEnumLong
	SPATypeEnumFloat
	SPATypeEnumDouble
	SPATypeEnumString
	SPATypeEnumBytes
	SPATypeEnumRectangle
	SPATypeEnumFraction
	SPATypeEnumBitmap
	SPATypeEnumArray
	SPATypeEnumStruct
	SPATypeEnumObject
	SPATypeEnumSequence
	SPATypeEnumPointer
	SPATypeEnumFd
	SPATypeEnumChoice
	SPATypeEnumPod
	SPATypeEnumLast /**< not part of ABI */

	/* Pointers */

	SPATypeEnumPointerStart SPAType = iota + 0x10000
	SPATypeEnumPointerBuffer
	SPATypeEnumPointerMeta
	SPATypeEnumPointerDict
	SPATypeEnumPointerLast /**< not part of ABI */

	/* Events */

	SPATypeEnumEventStart SPAType = iota + 0x20000
	SPATypeEnumEventDevice
	SPATypeEnumEventNode
	SPATypeEnumEventLast /**< not part of ABI */

	/* Commands */

	SPATypeEnumCommandStart SPAType = iota + 0x30000
	SPATypeEnumCommandDevice
	SPATypeEnumCommandNode
	SPATypeEnumCommandLast /**< not part of ABI */

	/* Objects */

	SPATypeEnumObjectStart SPAType = iota + 0x40000
	SPATypeEnumObjectPropInfo
	SPATypeEnumObjectProps
	SPATypeEnumObjectFormat
	SPATypeEnumObjectParamBuffers
	SPATypeEnumObjectParamMeta
	SPATypeEnumObjectParamIO
	SPATypeEnumObjectParamProfile
	SPATypeEnumObjectParamPortConfig
	SPATypeEnumObjectParamRoute
	SPATypeEnumObjectProfiler
	SPATypeEnumObjectParamLatency
	SPATypeEnumObjectParamProcessLatency
	SPATypeEnumObjectLast /**< not part of ABI */

	/* vendor extensions */

	SPATypeEnumVendorPipeWire SPAType = 0x02000000

	SPATypeEnumVendorOther SPAType = 0x7f000000
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
