package pod

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
