package pipewire

/*
#cgo pkg-config: libpipewire-0.3
#include <pipewire/pipewire.h>
*/
import "C"
import "unsafe"

// Init Initialize the PipeWire system, parse and modify any parameters given
// by argc and argv and set up debugging.
func Init(argv []string) Pipewire {

	pipewireInit := Pipewire{
		cArgs:    makeStrings(argv),
		cArgsLen: C.int(len(argv)),
	}

	C.pw_init(
		&pipewireInit.cArgsLen, &pipewireInit.cArgs,
	)
	return pipewireInit
}

// DeInit Deinitialize PipeWire.
//
// Deinitialize the PipeWire system and free up all resources allocated by
// Init().
//
// Before 0.3.49 this function can only be called once after which the pipewire
// library can not be used again. This is usually called by test programs to
// check for memory leaks.
//
// Since 0.3.49 this function must be paired with an equal amount of Init()
// calls to deinitialize the PipeWire library. The PipeWire library can be used
// again after being deinitialized with a new Init() call.
func (p *Pipewire) DeInit() {
	C.pw_deinit()
	destroyStrings(p.cArgs, int(p.cArgsLen))
}

// DebugIsCategoryEnabled check if a debug category is enabled.
//
// Debugging categories can be enabled by using the PIPEWIRE_DEBUG environment
// variable.
func (p Pipewire) DebugIsCategoryEnabled(name string) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.pw_debug_is_category_enabled(cName) == true
}

// GetApplicationName get the application name.
func (p Pipewire) GetApplicationName() string {
	name := C.pw_get_application_name()
	return C.GoString(name)
}

// GetProgramName get the program name.
func (p Pipewire) GetProgramName() string {
	name := C.pw_get_prgname()
	return C.GoString(name)
}

// GetUserName get the user name.
func (p Pipewire) GetUserName() string {
	name := C.pw_get_user_name()
	return C.GoString(name)
}

// GetHostName get the host name
func (p Pipewire) GetHostName() string {
	host := C.pw_get_host_name()
	return C.GoString(host)
}

// GetClientName get the client name.
//
// Make a new PipeWire client name that can be used to construct a remote.
func (p Pipewire) GetClientName() string {
	name := C.pw_get_client_name()
	return C.GoString(name)
}

// InValgrind return true when process is executed inside valgrind
func (p Pipewire) InValgrind() bool {
	return C.pw_in_valgrind() == true
}

func (p Pipewire) CheckOption(option, value string) bool {
	opts := C.CString(option)
	val := C.CString(value)
	defer C.free(unsafe.Pointer(opts))
	defer C.free(unsafe.Pointer(val))
	return C.pw_check_option(opts, val) == true
}
