package Windows

import "unsafe"

type (
	HMOUDLE = unsafe.Pointer
)

const (
	DllProcessAttach = 1
	DllThreadAttach  = 2
	DllThreadDetach  = 3
	DllProcessDetach = 4
)
