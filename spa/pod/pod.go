package pod

import (
	"unsafe"

	"gitea.linesip.com/libraries/gopipewire/spa/utils"
)

type SPAPod struct {
	Size uint32 // size of the body
	Type uint32 // a basic id of enum spa_type
}

type SPAPodBool struct {
	Pod     SPAPod
	Value   int32
	padding int32
}

type SPAPodID struct {
	Pod     SPAPod
	Value   uint32
	padding int32
}

type SPAPodInt struct {
	Pod     SPAPod
	Value   int32
	padding int32
}

type SPAPodLong struct {
	Pod   SPAPod
	Value int64
}

type SPAPodFloat struct {
	Pod     SPAPod
	Value   float32
	padding int32
}

type SPAPodDouble struct {
	Pod   SPAPod
	Value float64
}

type SPAPodString struct {
	Pod SPAPod
}

type SPAPodBytes struct {
	Pod SPAPod
}

type SPAPodRectangle struct {
	Pod   SPAPod
	Value utils.Rectangle
}

type SPAPodFraction struct {
	Pod   SPAPod
	Value utils.Fraction
}

type SPAPodBitmap struct {
	Pod SPAPod
}

type SPAPodArrayBody struct {
	Child SPAPod
}

type SPAPodArray struct {
	Pod  SPAPod
	Body SPAPodArrayBody
}

type SPAPodChoiceType uint32

const (
	// no choice, first value is current
	SPAChoiceNone SPAPodChoiceType = iota
	// range: default, min, max
	SPAChoiceRange
	// range with step: default, min, max, step
	SPAChoiceStep
	// list: default, alternative,...
	SPAChoiceEnum
	// flags: default, possible flags,...
	SPAChoiceFlags
)

type SPAPodChoiceBody struct {
	// type of choice, one of enum spa_choice_type
	Type SPAPodChoiceType
	// extra flags
	Flags uint32
	Child SPAPod
}

type SPAPodChoice struct {
	Pod  SPAPod
	Body SPAPodChoiceBody
}

type SPAPodStruct struct {
	Pod SPAPod
}

type SPAPodObjectBody struct {
	Type utils.SPAType
	ID   uint32
}

type SPAPodObject struct {
	Pod  SPAPod
	Body SPAPodObjectBody
}

type SPAPodPointerBody struct {
	Pod     SPAPod
	padding uint32
	Body    unsafe.Pointer
}

type SPAPodPointer struct {
	Pod  SPAPod
	Body SPAPodPointerBody
}

type SPAPodFD struct {
	Pod   SPAPod
	Value int64
}
