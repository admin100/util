package console

// +----------------------------------------------------------------------
// | console [ golang operation of the console's ]
// +----------------------------------------------------------------------
// | Author: lemontea <36634584@qq.com> <https://github.com/admin100>
// +----------------------------------------------------------------------

import (
	"syscall"
	"unsafe"
)

//Set windows console title
func SetConsoleTitle(title string) (int, error) {
	handle, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		return 0, err
	}
	defer syscall.FreeLibrary(handle)
	proc, err := syscall.GetProcAddress(handle, "SetConsoleTitleW")
	if err != nil {
		return 0, err
	}
	r, _, err := syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
	return int(r), err
}
