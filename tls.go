package box2d

import (
	"runtime"
	"unsafe"
)

var theStack *_Stack = newStack(1024 * 256)

type _Stack struct {
	Stack  []byte
	allocs []int
	pos    int

	heap map[uintptr]unsafe.Pointer
}

//go:noinline
func newStack(stackSize int) *_Stack {
	return &_Stack{
		Stack: make([]byte, stackSize),
		heap:  make(map[uintptr]unsafe.Pointer, 4096),
	}
}

func (tls *_Stack) Alloc(n int) uintptr {
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

func (tls *_Stack) Free(n int) uintptr {
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

func (tls *_Stack) toCString(src string) alloc {
	n := len(src) + 1
	mem := tls.Alloc(n)

	buf := sliceOf(mem, size_t(n))
	copy(buf, src)
	buf[n-1] = 0

	return alloc{Addr: mem, tls: tls, size: n}
}

type alloc struct {
	Addr uintptr
	tls  *_Stack
	size int
}

func (a alloc) Free() {
	a.tls.Free(a.size)
}

func copyToStack[T any](tls *_Stack, value T) alloc {
	n := unsafe.Sizeof(value)
	stack := tls.Alloc(int(n))
	memcpy(tls, stack, uintptr(unsafe.Pointer(&value)), size_t(n))

	runtime.KeepAlive(value)

	return alloc{Addr: stack, tls: tls, size: int(n)}
}
