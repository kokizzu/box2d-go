package box2d

import (
	"fmt"
	"math"
	"math/bits"
	"runtime"
	"strings"
	"sync/atomic"
	"time"
	"unsafe"
)

func sliceOf(p uintptr, n size_t) []byte {
	arr := (*[math.MaxInt32]byte)(unsafe.Pointer(p))
	return (*arr)[:n]
}

func __atomic_fetch_addInt32(tls *TLS, addr uintptr, inc int32, mode int32) int32 {
	value := (*int32)(unsafe.Pointer(addr))

	prev := *value
	*value += inc

	return prev
}

func __builtin___atomic_compare_exchange_n(tls *TLS, addr, expected uintptr, desired int32, weak uint8, success_memorder, failure_memorder int32) uint8 {
	ptrValue := (*int32)(unsafe.Pointer(addr))
	ptrExpected := (*int32)(unsafe.Pointer(expected))

	if atomic.CompareAndSwapInt32(ptrValue, *ptrExpected, desired) {
		return true1
	} else {
		return false1
	}
}

func __builtin_remainderf(_ *TLS, a, b float32) float32 {
	return float32(math.Remainder(float64(a), float64(b)))
}

func __builtin_trap(_ *TLS) {
	panic("trap")
}

func __builtin_clz(t *TLS, n uint32) int32 {
	return int32(bits.LeadingZeros32(n))
}

func __builtin_ctzll(_ *TLS, value uint64) int32 {
	return int32(bits.TrailingZeros64(value))
}
func memset(_ *TLS, dest uintptr, c int32, n size_t) (r uintptr) {
	slice := sliceOf(dest, n)

	if c == 0 || false {
		clear(slice)
	} else {
		for idx := range slice {
			slice[idx] = byte(c)
		}
	}

	return dest
}

func memcpy(_ *TLS, a uintptr, b uintptr, n size_t) uintptr {
	dst := sliceOf(a, n)
	src := sliceOf(b, n)
	copy(dst, src)
	return a
}

var heap = map[uintptr]unsafe.Pointer{}

func malloc(_ *TLS, size size_t) uintptr {
	buf := unsafe.Pointer(unsafe.SliceData(make([]byte, size+32)))

	// always return a value aligned to 32 byte
	if uintptr(buf)%32 != 0 {
		buf = unsafe.Add(buf, 32-uintptr(buf)%32)
	}

	heap[uintptr(buf)] = buf

	fmt.Printf("Alloc %d bytes at %#x\n", size, uintptr(buf))

	return uintptr(buf)
}

func free(_ *TLS, ptr uintptr) {
	delete(heap, ptr)
}

func sqrtf(_ *TLS, f float32) float32 {
	return float32(math.Sqrt(float64(f)))
}

func printf(tls *TLS, fmt uintptr, va uintptr) (r int32) {
	panic("todo")
}

func qsort(tls *TLS, base uintptr, nel size_t, width size_t, cmp uintptr) {
	panic("todo")
}

func sched_yield(tls *TLS) int32 {
	runtime.Gosched()
	return 0
}

func clock_gettime(tls *TLS, clk clockid_t, ts uintptr) (r int32) {
	now := time.Now()

	(*timespec)(unsafe.Pointer(ts)).Tv_sec = now.Unix()
	(*timespec)(unsafe.Pointer(ts)).Tv_nsec = now.UnixNano() - now.Unix()*1_000_000_000

	return 0
}

func __builtin_snprintf(t *TLS, str uintptr, size size_t, format, args uintptr) int32 {
	panic("todo")
}

func __builtin_Gosched(tls *TLS) {
	runtime.Gosched()
}

func fprintf(tls *TLS, f uintptr, fmt uintptr, va uintptr) (r int32) {
	panic("todo")
}

func fopen(tls *TLS, filename uintptr, mode uintptr) (r uintptr) {
	panic("todo")
}

func fclose(tls *TLS, f uintptr) (r1 int32) {
	panic("todo")
}

// Dummy annotation marking that the value x escapes,
// for use in cases where the reflect code is so clever that
// the compiler cannot follow.
func escapes(x any) {
	if dummy.b {
		dummy.x = x
	}
}

var dummy struct {
	b bool
	x any
}

func addrOf[T any](val *T) uintptr {
	escapes(val)
	return uintptr(unsafe.Pointer(val))
}

func toString(base uintptr) string {
	var str strings.Builder

	for n := 0; ; n += 1 {
		ch := *(*byte)(unsafe.Pointer(base + uintptr(n)))
		if ch == 0 {
			return str.String()
		}

		str.WriteByte(ch)
	}
}

func b2DefaultAssertFcnGo(tls *TLS, condition uintptr, fileName uintptr, lineNumber int32) (r int32) {
	err := fmt.Errorf("%s:%d: %s", toString(fileName), lineNumber, toString(condition))
	panic(err)
}
