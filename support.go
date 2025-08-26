package box2d

import (
	"fmt"
	"math"
	"math/bits"
	"runtime"
	"sort"
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

func __builtin_aligned_alloc(tls *TLS, align, size uint64) uintptr {
	buf := unsafe.Pointer(unsafe.SliceData(make([]byte, size+align)))
	addr := uintptr(buf)

	if addr%uintptr(align) != 0 {
		addr += uintptr(align) - addr%uintptr(align)
	}

	// address is now aligned
	if addr%uintptr(align) != 0 {
		panic("not aligned")
	}

	tls.heap[addr] = buf
	return addr
}

func memset(_ *TLS, dest uintptr, c int32, n size_t) (r uintptr) {
	slice := sliceOf(dest, n)

	if c == 0 {
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

func free(tls *TLS, ptr uintptr) {
	if _, ok := tls.heap[ptr]; !ok {
		err := fmt.Errorf("free(%#x): unknown address", ptr)
		panic(err)
	}

	delete(tls.heap, ptr)
}

func sqrtf(_ *TLS, f float32) float32 {
	return float32(math.Sqrt(float64(f)))
}

type _qsort struct {
	tls      *TLS
	memory   []byte
	compare  func(_ *TLS, a, b uintptr) int32
	itemSize int
	len      int
}

func (q _qsort) sliceOf(idx int) []byte {
	start := idx * q.itemSize
	return q.memory[start:][:q.itemSize]
}

func (q _qsort) Len() int {
	return q.len
}

func (q _qsort) Less(i, j int) bool {
	lhs := q.sliceOf(i)
	rhs := q.sliceOf(j)

	rc := q.compare(
		q.tls,
		uintptr(unsafe.Pointer(unsafe.SliceData(lhs))),
		uintptr(unsafe.Pointer(unsafe.SliceData(rhs))),
	)

	return rc < 0
}

func (q _qsort) Swap(i, j int) {
	lhs := q.sliceOf(i)
	rhs := q.sliceOf(j)

	for idx := range len(lhs) {
		// swap the bytes
		lhs[idx], rhs[idx] = rhs[idx], lhs[idx]
	}
}

func qsort(tls *TLS, base uintptr, nel size_t, width size_t, compare func(_ *TLS, a, b uintptr) int32) {
	sort.Sort(_qsort{
		tls:      tls,
		memory:   sliceOf(base, nel*width),
		compare:  compare,
		itemSize: int(width),
		len:      int(nel),
	})
}

func clock_gettime(tls *TLS, clk clockid_t, ts uintptr) (r int32) {
	now := time.Now()

	(*timespec)(unsafe.Pointer(ts)).Tv_sec = now.Unix()
	(*timespec)(unsafe.Pointer(ts)).Tv_nsec = now.UnixNano() - now.Unix()*1_000_000_000

	return 0
}

func sched_yield(tls *TLS) int32 {
	runtime.Gosched()
	return 0
}

func __builtin_Gosched(tls *TLS) {
	runtime.Gosched()
}

func printf(tls *TLS, fmt uintptr, va uintptr) (r int32) {
	panic("printf: not implemented")
}

func __builtin_snprintf(t *TLS, str uintptr, size size_t, format, args uintptr) int32 {
	panic("snprintf: not implemented")
}

func fprintf(tls *TLS, f uintptr, fmt uintptr, va uintptr) (r int32) {
	panic("fprintf: not implemented")
}

func fopen(tls *TLS, filename uintptr, mode uintptr) (r uintptr) {
	panic("fopen: not implemented")
}

func fclose(tls *TLS, f uintptr) (r1 int32) {
	panic("fclose: not implemented")
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
