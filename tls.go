package box2d

import (
	"runtime"
	"unsafe"
)

type TLS struct {
	Stack  []byte
	allocs []int
	pos    int
}

//go:noinline
func NewTLS(stackSize int) *TLS {
	return &TLS{
		Stack: make([]byte, stackSize),
	}
}

func StackAlloc[T any](tls *TLS) *T {
	var tZero T
	ptr := tls.Alloc(int(unsafe.Sizeof(tZero)))
	return (*T)(unsafe.Pointer(ptr))
}

func (tls *TLS) Alloc(n int) uintptr {
	if n%8 > 0 {
		// round up to a multiple of 8
		n += 8 - n%8
	}

	if tls.pos+n > len(tls.Stack) {
		panic("out of stack space")
	}

	buf := uintptr(unsafe.Pointer(unsafe.SliceData(tls.Stack[tls.pos:])))
	tls.pos += n
	tls.allocs = append(tls.allocs, n)

	return buf
}

func (tls *TLS) Free(n int) uintptr {
	if n%8 > 0 {
		// round up to a multiple of 8
		n += 8 - n%8
	}

	lastAlloc := tls.allocs[len(tls.allocs)-1]
	if lastAlloc != n {
		panic("alloc/free is not paired")
	}

	tls.allocs = tls.allocs[:len(tls.allocs)-1]
	tls.pos -= n

	return 0
}

type alloc struct {
	Addr uintptr
	tls  *TLS
	size int
}

func (a alloc) Free() {
	a.tls.Free(a.size)
}

func copyToStack[T any](tls *TLS, value T) alloc {
	n := unsafe.Sizeof(value)
	stack := tls.Alloc(int(n))
	memcpy(tls, stack, uintptr(unsafe.Pointer(&value)), size_t(n))

	runtime.KeepAlive(value)

	return alloc{Addr: stack, tls: tls, size: int(n)}
}
