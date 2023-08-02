package control

type SPAControlType uint32

const (
	SPAControlInvalid    SPAControlType = iota
	SPAControlProperties                // data contains a SPA_TYPE_OBJECT_Props
	SPAControlMidi                      // data contains a spa_pod_bytes with raw midi data
	SPAControlOSC                       // data contains a spa_pod_bytes with an OSC packet
	SPAControlLast                      // not part of ABI
)
