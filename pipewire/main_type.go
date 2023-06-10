package pipewire

import "C"

// Pipewire is the main memory management struct for controlling base of
// pipewire
type Pipewire struct {
	cArgs    **C.char
	cArgsLen C.int
}
