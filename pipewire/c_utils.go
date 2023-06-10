package pipewire

/*
#include "init.h"
*/
import "C"

func makeStrings(array []string) **C.char {
	cArray := C.make_strings(C.int(len(array)))
	for i, e := range array {
		cstr := C.CString(e)
		C.set_string(cArray, C.int(i), (*C.char)(cstr))
	}
	C.set_string(cArray, C.int(len(array)), nil)
	return cArray
}

func destroyStrings(strings **C.char, count int) {
	C.destroy_strings(strings, C.int(count))
}
