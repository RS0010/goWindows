package Windows

import "golang.org/x/sys/windows"

var (
	libKernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procDisableThreadLibraryCalls = libKernel32.NewProc("DisableThreadLibraryCalls")
)

func DisableThreadLibraryCalls(hLibModule HMOUDLE) bool {
	r0, _, _ := procDisableThreadLibraryCalls.Call(uintptr(hLibModule))
	if r0 == 0 {
		return false
	}
	return true
}