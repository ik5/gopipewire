package pipewire

/*
#cgo pkg-config: libpipewire-0.3
#include <pipewire/pipewire.h>
#include <pipewire/version.h>

int check_version(int major, int minor, int micro) {
	return PW_CHECK_VERSION(major, minor, micro);
}

char * get_headers_version(void) {
	return pw_get_headers_version();
}

*/
import "C"

// GetHeaderVersion return the version of the header files. Keep in mind that
// this is a macro and not a function, so it is impossible to get the pointer
// of it.
func (p Pipewire) GetHeaderVersion() string {
	return C.GoString(C.get_headers_version())
}

// GetLibraryVersion return the version of the header files.
// Keep in mind that this is a macro and not a function, so it is impossible to
// get the pointer of it.
func (p Pipewire) GetLibraryVersion() string {
	return C.GoString(C.pw_get_library_version())
}

func (p Pipewire) CheckVersion(major, minor, micro int) bool {
	return C.check_version(
		C.int(major), C.int(minor), C.int(micro),
	) > 0
}
