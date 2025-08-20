package box2d

import (
	"runtime"
	"slices"
	"unsafe"
)

type RestitutionCallback = func(restitutionA float32, userMaterialIdA int32, restitutionB float32, userMaterialIdB int32) float32
type FrictionCallback = func(frictionA float32, userMaterialIdA int32, frictionB float32, userMaterialIdB int32) float32

type Box2D struct {
	tls *TLS
}

func NewBox2D() *Box2D {
	box := &Box2D{
		tls: NewTLS(256 * 1024),
	}

	b2SetAssertFcn(box.tls, __ccgo_fp(b2DefaultAssertFcnGo))

	return box
}

func (b *Box2D) BodySetName(bodyId BodyId, name string) {
	nameC := b.toCString(name)
	defer nameC.Free()

	b2Body_SetName(b.tls, bodyId, nameC.Addr)
}

func (b *Box2D) BodyGetShapes(bodyId BodyId, reuse []ShapeId) []ShapeId {
	itemCount := b.BodyGetShapeCount(bodyId)

	if cap(reuse) < int(itemCount) {
		// need to allocate more space
		reuse = append(reuse[:0], make([]ShapeId, itemCount)...)
	}

	// ensure the data is on the heap
	items := reuse[:0]
	escapes(items)

	// fill the data
	ptr := addrOf(unsafe.SliceData(items))
	n := b2Body_GetShapes(b.tls, bodyId, ptr, int32(cap(items)))

	runtime.KeepAlive(items)
	return items[:n]
}

func (b *Box2D) BodyGetJoints(bodyId BodyId, reuse []JointId) []JointId {
	itemCount := b.BodyGetJointCount(bodyId)

	if cap(reuse) < int(itemCount) {
		// need to allocate more space
		reuse = append(reuse[:0], make([]JointId, itemCount)...)
	}

	// ensure the data is on the heap
	items := reuse[:0]
	escapes(items)

	// fill the data
	ptr := addrOf(unsafe.SliceData(items))
	n := b2Body_GetJoints(b.tls, bodyId, ptr, int32(cap(items)))

	runtime.KeepAlive(items)
	return items[:n]
}

func (b *Box2D) ChainGetSegments(chainId ChainId, reuse []ShapeId) []ShapeId {
	itemCount := b.ChainGetSegmentCount(chainId)

	if cap(reuse) < int(itemCount) {
		// need to allocate more space
		reuse = append(reuse[:0], make([]ShapeId, itemCount)...)
	}

	// ensure the data is on the heap
	items := reuse[:0]
	escapes(items)

	// fill the data
	ptr := addrOf(unsafe.SliceData(items))
	n := b2Chain_GetSegments(b.tls, chainId, ptr, int32(cap(items)))

	runtime.KeepAlive(items)
	return items[:n]
}

// WorldSetRestitutionCallback sets the worlds RestitutionCallback.
// CAUTION: Callback must not be a closure
func (b *Box2D) WorldSetRestitutionCallback(worldId WorldId, callback RestitutionCallback) {
	b2World_SetRestitutionCallback(b.tls, worldId, __ccgo_fp(callback))
}

// WorldSetFrictionCallback sets the worlds FrictionCallback.
// CAUTION: Callback must not be a closure
func (b *Box2D) WorldSetFrictionCallback(worldId WorldId, callback FrictionCallback) {
	b2World_SetFrictionCallback(b.tls, worldId, __ccgo_fp(callback))
}

type OverlapResultFcn = func(shapeId ShapeId, context uintptr) bool

// WorldOverlapAABB sets the worlds OverlapResultFcn
// CAUTION: fcn must not be a closure
func (b *Box2D) WorldOverlapAABB(worldId WorldId, aabb AABB, filter QueryFilter, fcn OverlapResultFcn, context uintptr) TreeStats {
	return b2World_OverlapAABB(b.tls, worldId, aabb, filter, __ccgo_fp(fcn), context)
}

type CustomFilterFcn = func(shapeIdA, shapeIdB ShapeId, context uintptr) bool

func (b *Box2D) WorldSetCustomFilterCallback(worldId WorldId, fcn CustomFilterFcn, context uintptr) {
	b2World_SetCustomFilterCallback(b.tls, worldId, __ccgo_fp(fcn), context)
}

type PreSolveFcn = func(shapeIdA, shapeIdB ShapeId, manifold uintptr, context uintptr)

func (b *Box2D) WorldSetPreSolveCallback(worldId WorldId, fcn PreSolveFcn, context uintptr) {
	b2World_SetPreSolveCallback(b.tls, worldId, __ccgo_fp(fcn), context)
}

func (b *Box2D) toCString(src string) alloc {
	n := len(src) + 1
	mem := b.tls.Alloc(n)

	buf := sliceOf(mem, size_t(n))
	copy(buf, src)
	buf[n-1] = 0

	return alloc{Addr: mem, tls: b.tls, size: n}
}

type BodyEvents struct {
	MoveEvents []BodyMoveEvent
}

func (b *Box2D) WorldGetBodyEvents(w WorldId) BodyEvents {
	events := b2World_GetBodyEvents(b.tls, w)

	return BodyEvents{
		MoveEvents: copySlice[BodyMoveEvent](events.MoveEvents, events.MoveCount),
	}
}

type SensorEvents struct {
	BeginEvents []SensorBeginTouchEvent
	EndEvents   []SensorEndTouchEvent
}

func (b *Box2D) WorldGetSensorEvents(w WorldId) SensorEvents {
	events := b2World_GetSensorEvents(b.tls, w)

	return SensorEvents{
		BeginEvents: copySlice[SensorBeginTouchEvent](events.BeginEvents, events.BeginCount),
		EndEvents:   copySlice[SensorEndTouchEvent](events.EndEvents, events.EndCount),
	}
}

type ContactEvents struct {
	BeginEvents []ContactBeginTouchEvent
	EndEvents   []ContactEndTouchEvent
	HitEvents   []ContactHitEvent
}

func (b *Box2D) WorldGetContactEvents(w WorldId) ContactEvents {
	events := b2World_GetContactEvents(b.tls, w)

	return ContactEvents{
		BeginEvents: copySlice[ContactBeginTouchEvent](events.BeginEvents, events.BeginCount),
		EndEvents:   copySlice[ContactEndTouchEvent](events.EndEvents, events.EndCount),
		HitEvents:   copySlice[ContactHitEvent](events.HitEvents, events.HitCount),
	}
}

func copySlice[T any](data uintptr, n int32) []T {
	event := (*T)(unsafe.Pointer(data))
	bodyMoveEvents := unsafe.Slice(event, n)
	return slices.Clone(bodyMoveEvents)
}
