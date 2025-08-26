package box2d

import (
	"fmt"
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

func (b Body) SetName(name string) {
	nameC := b.tls.toCString(name)
	defer nameC.Free()

	b2Body_SetName(b.tls, b.Id, nameC.Addr)
}

func (b Body) GetShapes(reuse []Shape) []Shape {
	itemCount := b.GetShapeCount()

	// allocate memory on the stack to read the ids
	n := int(itemCount) * int(unsafe.Sizeof(ShapeId{}))
	ptr := b.tls.Alloc(n)
	defer b.tls.Free(n)

	// fill the data
	_ = b2Body_GetShapes(b.tls, b.Id, ptr, itemCount)

	// allocate memory for the result
	if cap(reuse) < int(itemCount) {
		// need to allocate more space
		reuse = make([]Shape, itemCount)
	}

	// ensure the data is on the heap
	items := reuse[:0]

	for _, itemId := range unsafe.Slice((*ShapeId)(unsafe.Pointer(ptr)), itemCount) {
		items = append(items, Shape{tls: b.tls, Id: itemId})
	}

	return items
}

func (b Body) GetJoints(reuse []Joint) []Joint {
	itemCount := b.GetJointCount()

	// allocate memory on the stack to read the ids
	n := int(itemCount) * int(unsafe.Sizeof(JointId{}))
	ptr := b.tls.Alloc(n)
	defer b.tls.Free(n)

	// fill the data
	_ = b2Body_GetJoints(b.tls, b.Id, ptr, itemCount)

	// allocate memory for the result
	if cap(reuse) < int(itemCount) {
		// need to allocate more space
		reuse = make([]Joint, itemCount)
	}

	items := reuse[:0]

	// copy the items
	for _, itemId := range unsafe.Slice((*ShapeId)(unsafe.Pointer(ptr)), itemCount) {
		items = append(items, Joint{tls: b.tls, Id: itemId})
	}

	return items
}

func (b Chain) GetSegments(reuse []Shape) []Shape {
	itemCount := b.GetSegmentCount()

	// allocate memory on the stack to read the ids
	n := int(itemCount) * int(unsafe.Sizeof(ShapeId{}))
	ptr := b.tls.Alloc(n)
	defer b.tls.Free(n)

	// fill the data
	_ = b2Chain_GetSegments(b.tls, b.Id, ptr, itemCount)

	// allocate memory for the result
	if cap(reuse) < int(itemCount) {
		// need to allocate more space
		reuse = make([]Shape, itemCount)
	}

	items := reuse[:0]

	// copy the items
	for _, itemId := range unsafe.Slice((*ShapeId)(unsafe.Pointer(ptr)), itemCount) {
		items = append(items, Shape{tls: b.tls, Id: itemId})
	}

	return items
}

func (b Shape) GetSensorOverlaps(reuse []Shape) []Shape {
	// allow up to 1024 overlaps to be returned
	itemsCap := int32(1024)

	// allocate memory on the stack to read the ids
	n := int(itemsCap) * int(unsafe.Sizeof(ShapeId{}))
	ptr := b.tls.Alloc(n)
	defer b.tls.Free(n)

	// fill the data
	itemCount := b2Shape_GetSensorOverlaps(b.tls, b.Id, ptr, itemsCap)

	// allocate memory for the result
	if cap(reuse) < int(itemCount) {
		// need to allocate more space
		reuse = make([]Shape, itemCount)
	}

	items := reuse[:0]

	// copy the items
	for _, itemId := range unsafe.Slice((*ShapeId)(unsafe.Pointer(ptr)), itemCount) {
		items = append(items, Shape{tls: b.tls, Id: itemId})
	}

	return items
}

// SetRestitutionCallback sets the worlds RestitutionCallback.
// CAUTION: Callback must not be a closure
func (b World) SetRestitutionCallback(callback RestitutionCallback) {
	b2World_SetRestitutionCallback(b.tls, b.Id, __ccgo_fp(callback))
}

// SetFrictionCallback sets the worlds FrictionCallback.
// CAUTION: Callback must not be a closure
func (b World) SetFrictionCallback(callback FrictionCallback) {
	b2World_SetFrictionCallback(b.tls, b.Id, __ccgo_fp(callback))
}

type OverlapResultFcn = func(shapeId ShapeId, context uintptr) bool

// OverlapAABB sets the worlds OverlapResultFcn
// CAUTION: fcn must not be a closure
func (b World) OverlapAABB(aabb AABB, filter QueryFilter, fcn OverlapResultFcn, context uintptr) TreeStats {
	return b2World_OverlapAABB(b.tls, b.Id, aabb, filter, __ccgo_fp(fcn), context)
}

type CustomFilterFcn = func(shapeIdA, shapeIdB ShapeId, context uintptr) bool

func (b World) SetCustomFilterCallback(fcn CustomFilterFcn, context uintptr) {
	b2World_SetCustomFilterCallback(b.tls, b.Id, __ccgo_fp(fcn), context)
}

type PreSolveFcn = func(shapeIdA, shapeIdB ShapeId, manifold uintptr, context uintptr)

func (b World) SetPreSolveCallback(fcn PreSolveFcn, context uintptr) {
	b2World_SetPreSolveCallback(b.tls, b.Id, __ccgo_fp(fcn), context)
}

type BodyEvents struct {
	MoveEvents []BodyMoveEvent
}

func (b World) GetBodyEvents() BodyEvents {
	events := b2World_GetBodyEvents(b.tls, b.Id)

	return BodyEvents{
		MoveEvents: copySlice[BodyMoveEvent](events.MoveEvents, events.MoveCount),
	}
}

type SensorEvents struct {
	BeginEvents []SensorBeginTouchEvent
	EndEvents   []SensorEndTouchEvent
}

func (b World) GetSensorEvents() SensorEvents {
	events := b2World_GetSensorEvents(b.tls, b.Id)

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

func (b World) GetContactEvents() ContactEvents {
	events := b2World_GetContactEvents(b.tls, b.Id)

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

func b2DefaultAssertFcnGo(tls *TLS, condition uintptr, fileName uintptr, lineNumber int32) (r int32) {
	err := fmt.Errorf("%s:%d: %s", toString(fileName), lineNumber, toString(condition))
	panic(err)
}
