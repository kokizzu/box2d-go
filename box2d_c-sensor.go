package box2d

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b2SensorArray_Create(tls *_Stack, capacity int32) (r b2SensorArray) {
	var a b2SensorArray
	_ = a
	a = b2SensorArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(40)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2SensorArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2SensorArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2SensorArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2SensorArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2SensorArray)(unsafe.Pointer(a)).Capacity)*uint64(40)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(40)))
	(*b2SensorArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2SensorArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2SensorArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2SensorArray)(unsafe.Pointer(a)).Capacity)*uint64(40)))
	(*b2SensorArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2SensorArray)(unsafe.Pointer(a)).Count = 0
	(*b2SensorArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2SensorTaskContextArray_Create(tls *_Stack, capacity int32) (r b2SensorTaskContextArray) {
	var a b2SensorTaskContextArray
	_ = a
	a = b2SensorTaskContextArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(16)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2SensorTaskContextArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2SensorTaskContextArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2SensorTaskContextArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2SensorTaskContextArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2SensorTaskContextArray)(unsafe.Pointer(a)).Capacity)*uint64(16)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(16)))
	(*b2SensorTaskContextArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2SensorTaskContextArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2SensorTaskContextArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2SensorTaskContextArray)(unsafe.Pointer(a)).Capacity)*uint64(16)))
	(*b2SensorTaskContextArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2SensorTaskContextArray)(unsafe.Pointer(a)).Count = 0
	(*b2SensorTaskContextArray)(unsafe.Pointer(a)).Capacity = 0
}

// Sensor shapes need to
// - detect begin and end overlap events
// - events must be reported in deterministic order
// - maintain an active list of overlaps for query

// Assumption
// - sensors don't detect shapes on the same body

// Algorithm
// Query all sensors for overlaps
// Check against previous overlaps

// Data structures
// Each sensor has an double buffered array of overlaps
// These overlaps use a shape reference with index and generation

func b2SensorQueryCallback(tls *_Stack, proxyId int32, userData uint64, context uintptr) (r uint8) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var newCapacity, sensorShapeId, shapeId, v8 int32
	var otherShape, queryContext, sensor, sensorShape, shapeRef, world, v1, v7, v9 uintptr
	var otherTransform Transform
	var output DistanceOutput
	var overlaps, v5 uint8
	var v3, v4 Filter
	var _ /* cache at bp+184 */ SimplexCache
	var _ /* input at bp+0 */ DistanceInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = newCapacity, otherShape, otherTransform, output, overlaps, queryContext, sensor, sensorShape, sensorShapeId, shapeId, shapeRef, world, v1, v3, v4, v5, v7, v8, v9
	_ = uint64FromInt64(4)
	shapeId = int32FromUint64(userData)
	queryContext = context
	sensorShape = (*b2SensorQueryContext)(unsafe.Pointer(queryContext)).SensorShape
	sensorShapeId = (*b2Shape)(unsafe.Pointer(sensorShape)).Id
	if shapeId == sensorShapeId {
		return uint8(true1)
	}
	world = (*b2SensorQueryContext)(unsafe.Pointer(queryContext)).World
	v1 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr(shapeId)*288
	goto _2
_2:
	otherShape = v1
	// Are sensor events enabled on the other shape?
	if int32FromUint8((*b2Shape)(unsafe.Pointer(otherShape)).EnableSensorEvents) == false1 {
		return uint8(true1)
	}
	// Skip shapes on the same body
	if (*b2Shape)(unsafe.Pointer(otherShape)).BodyId == (*b2Shape)(unsafe.Pointer(sensorShape)).BodyId {
		return uint8(true1)
	}
	// Check filter
	v3 = (*b2Shape)(unsafe.Pointer(sensorShape)).Filter
	v4 = (*b2Shape)(unsafe.Pointer(otherShape)).Filter
	if v3.GroupIndex == v4.GroupIndex && v3.GroupIndex != 0 {
		v5 = boolUint8(v3.GroupIndex > 0)
		goto _6
	}
	v5 = boolUint8(v3.MaskBits&v4.CategoryBits != uint64(0) && v3.CategoryBits&v4.MaskBits != uint64(0))
	goto _6
_6:
	if int32FromUint8(v5) == false1 {
		return uint8(true1)
	}
	otherTransform = b2GetBodyTransform(tls, world, (*b2Shape)(unsafe.Pointer(otherShape)).BodyId)
	(*(*DistanceInput)(unsafe.Pointer(bp))).ProxyA = b2MakeShapeDistanceProxy(tls, sensorShape)
	(*(*DistanceInput)(unsafe.Pointer(bp))).ProxyB = b2MakeShapeDistanceProxy(tls, otherShape)
	(*(*DistanceInput)(unsafe.Pointer(bp))).TransformA = (*b2SensorQueryContext)(unsafe.Pointer(queryContext)).Transform
	(*(*DistanceInput)(unsafe.Pointer(bp))).TransformB = otherTransform
	(*(*DistanceInput)(unsafe.Pointer(bp))).UseRadii = uint8(true1)
	*(*SimplexCache)(unsafe.Pointer(bp + 184)) = SimplexCache{}
	output = b2ShapeDistance(tls, bp, bp+184, uintptrFromInt32(0), 0)
	overlaps = boolUint8(output.Distance < float32(float32FromFloat32(10)*float32FromFloat32(1.1920928955078125e-07)))
	if int32FromUint8(overlaps) == false1 {
		return uint8(true1)
	}
	// Record the overlap
	sensor = (*b2SensorQueryContext)(unsafe.Pointer(queryContext)).Sensor
	v7 = sensor + 16
	if (*b2ShapeRefArray)(unsafe.Pointer(v7)).Count == (*b2ShapeRefArray)(unsafe.Pointer(v7)).Capacity {
		if (*b2ShapeRefArray)(unsafe.Pointer(v7)).Capacity < int32(2) {
			v8 = int32(2)
		} else {
			v8 = (*b2ShapeRefArray)(unsafe.Pointer(v7)).Capacity + (*b2ShapeRefArray)(unsafe.Pointer(v7)).Capacity>>int32(1)
		}
		newCapacity = v8
		b2ShapeRefArray_Reserve(tls, v7, newCapacity)
	}
	*(*int32)(unsafe.Pointer(v7 + 8)) += int32(1)
	v9 = (*b2ShapeRefArray)(unsafe.Pointer(v7)).Data + uintptr((*b2ShapeRefArray)(unsafe.Pointer(v7)).Count-int32FromInt32(1))*8
	goto _10
_10:
	shapeRef = v9
	(*b2ShapeRef)(unsafe.Pointer(shapeRef)).ShapeId = shapeId
	(*b2ShapeRef)(unsafe.Pointer(shapeRef)).Generation = (*b2Shape)(unsafe.Pointer(otherShape)).Generation
	return uint8(true1)
}

func b2SensorTask(tls *_Stack, startIndex int32, endIndex int32, threadIndex uint32, context uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var blockIndex, v11, v8, v9 uint32_t
	var body, s1, s2, sensor, sensorShape, taskContext, trees, world, v2, v4, v6 uintptr
	var count1, count2, i, sensorIndex int32
	var queryBounds AABB
	var temp b2ShapeRefArray
	var transform Transform
	var _ /* queryContext at bp+0 */ b2SensorQueryContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = blockIndex, body, count1, count2, i, queryBounds, s1, s2, sensor, sensorIndex, sensorShape, taskContext, temp, transform, trees, world, v11, v2, v4, v6, v8, v9
	world = context
	taskContext = (*b2World)(unsafe.Pointer(world)).SensorTaskContexts.Data + uintptr(threadIndex)*16
	trees = world + 40
	sensorIndex = startIndex
	for {
		if !(sensorIndex < endIndex) {
			break
		}
		v2 = (*b2SensorArray)(unsafe.Pointer(world+1280)).Data + uintptr(sensorIndex)*40
		goto _3
	_3:
		sensor = v2
		v4 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr((*b2Sensor)(unsafe.Pointer(sensor)).ShapeId)*288
		goto _5
	_5:
		sensorShape = v4
		// swap overlap arrays
		temp = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1
		(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2
		(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2 = temp
		(*b2ShapeRefArray)(unsafe.Pointer(sensor + 16)).Count = 0
		v6 = (*b2BodyArray)(unsafe.Pointer(world+1024)).Data + uintptr((*b2Shape)(unsafe.Pointer(sensorShape)).BodyId)*128
		goto _7
	_7:
		body = v6
		if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_disabledSet) || int32FromUint8((*b2Shape)(unsafe.Pointer(sensorShape)).EnableSensorEvents) == false1 {
			if (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Count != 0 {
				v8 = uint32FromInt32(sensorIndex)
				blockIndex = v8 / uint32(64)
				*(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(taskContext)).Bits + uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v8 % uint32FromInt32(64))
			}
			goto _1
		}
		transform = b2GetBodyTransformQuick(tls, world, body)
		*(*b2SensorQueryContext)(unsafe.Pointer(bp)) = b2SensorQueryContext{
			World:       world,
			TaskContext: taskContext,
			Sensor:      sensor,
			SensorShape: sensorShape,
			Transform:   transform,
		}
		queryBounds = (*b2Shape)(unsafe.Pointer(sensorShape)).Aabb
		// Query all trees
		b2DynamicTree_Query(tls, trees+uintptr(0)*72, queryBounds, (*b2Shape)(unsafe.Pointer(sensorShape)).Filter.MaskBits, __ccgo_fp(b2SensorQueryCallback), bp)
		b2DynamicTree_Query(tls, trees+uintptr(1)*72, queryBounds, (*b2Shape)(unsafe.Pointer(sensorShape)).Filter.MaskBits, __ccgo_fp(b2SensorQueryCallback), bp)
		b2DynamicTree_Query(tls, trees+uintptr(2)*72, queryBounds, (*b2Shape)(unsafe.Pointer(sensorShape)).Filter.MaskBits, __ccgo_fp(b2SensorQueryCallback), bp)
		// Sort the overlaps to enable finding begin and end events.
		qsort(tls, (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data, uint64FromInt32((*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count), uint64(8), b2CompareShapeRefs)
		count1 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Count
		count2 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count
		if count1 != count2 {
			// something changed
			v9 = uint32FromInt32(sensorIndex)
			blockIndex = v9 / uint32(64)
			*(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(taskContext)).Bits + uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v9 % uint32FromInt32(64))
		} else {
			i = 0
			for {
				if !(i < count1) {
					break
				}
				s1 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Data + uintptr(i)*8
				s2 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data + uintptr(i)*8
				if (*b2ShapeRef)(unsafe.Pointer(s1)).ShapeId != (*b2ShapeRef)(unsafe.Pointer(s2)).ShapeId || int32FromUint16((*b2ShapeRef)(unsafe.Pointer(s1)).Generation) != int32FromUint16((*b2ShapeRef)(unsafe.Pointer(s2)).Generation) {
					// something changed
					v11 = uint32FromInt32(sensorIndex)
					blockIndex = v11 / uint32(64)
					*(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(taskContext)).Bits + uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v11 % uint32FromInt32(64))
					break
				}
				goto _10
			_10:
				;
				i++
			}
		}
		goto _1
	_1:
		;
		sensorIndex++
	}
}

func b2OverlapSensors(tls *_Stack, world uintptr) {
	var bitSet, bits, r1, r11, r2, r21, refs1, refs2, sensor, sensorShape, userSensorTask, v10, v12, v14, v16, v18, v20, v6, v8 uintptr
	var blockCount, ctz, k, v4 uint32_t
	var count1, count2, i, i1, index11, index2, minRange, newCapacity, newCapacity1, sensorCount, sensorIndex, v11, v13, v15, v17, v19, v21 int32
	var event, event2, event4 SensorEndTouchEvent
	var event1, event3, event5 SensorBeginTouchEvent
	var sensorId, visitorId, visitorId1, visitorId2, visitorId3, visitorId4, visitorId5 ShapeId
	var word uint64_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bitSet, bits, blockCount, count1, count2, ctz, event, event1, event2, event3, event4, event5, i, i1, index11, index2, k, minRange, newCapacity, newCapacity1, r1, r11, r2, r21, refs1, refs2, sensor, sensorCount, sensorId, sensorIndex, sensorShape, userSensorTask, visitorId, visitorId1, visitorId2, visitorId3, visitorId4, visitorId5, word, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v4, v6, v8
	sensorCount = (*b2World)(unsafe.Pointer(world)).Sensors.Count
	if sensorCount == 0 {
		return
	}
	i = 0
	for {
		if !(i < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		b2SetBitCountAndClear(tls, (*b2World)(unsafe.Pointer(world)).SensorTaskContexts.Data+uintptr(i)*16, uint32FromInt32(sensorCount))
		goto _1
	_1:
		;
		i++
	}
	// Parallel-for sensors overlaps
	minRange = int32(16)
	userSensorTask = (*(*func(*_Stack, uintptr, int32, int32, uintptr, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).EnqueueTaskFcn})))(tls, __ccgo_fp(b2SensorTask), sensorCount, minRange, world, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
	*(*int32)(unsafe.Pointer(world + 1776)) += int32(1)
	if userSensorTask != uintptrFromInt32(0) {
		(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).FinishTaskFcn})))(tls, userSensorTask, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
	}
	bitSet = (*b2World)(unsafe.Pointer(world)).SensorTaskContexts.Data
	i1 = int32(1)
	for {
		if !(i1 < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		b2InPlaceUnion(tls, bitSet, (*b2World)(unsafe.Pointer(world)).SensorTaskContexts.Data+uintptr(i1)*16)
		goto _2
	_2:
		;
		i1++
	}
	// Iterate sensors bits and publish events
	// Process contact state changes. Iterate over set bits
	bits = (*b2BitSet)(unsafe.Pointer(bitSet)).Bits
	blockCount = (*b2BitSet)(unsafe.Pointer(bitSet)).BlockCount
	k = uint32(0)
	for {
		if !(k < blockCount) {
			break
		}
		word = *(*uint64_t)(unsafe.Pointer(bits + uintptr(k)*8))
		for word != uint64(0) {
			v4 = uint32FromInt32(__builtin_ctzll(tls, word))
			goto _5
		_5:
			ctz = v4
			sensorIndex = int32FromUint32(uint32FromInt32(64)*k + ctz)
			v6 = (*b2SensorArray)(unsafe.Pointer(world+1280)).Data + uintptr(sensorIndex)*40
			goto _7
		_7:
			sensor = v6
			v8 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr((*b2Sensor)(unsafe.Pointer(sensor)).ShapeId)*288
			goto _9
		_9:
			sensorShape = v8
			sensorId = ShapeId{
				Index1:     (*b2Sensor)(unsafe.Pointer(sensor)).ShapeId + int32(1),
				World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
				Generation: (*b2Shape)(unsafe.Pointer(sensorShape)).Generation,
			}
			count1 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Count
			count2 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count
			refs1 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Data
			refs2 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data
			// overlaps1 can have overlaps that end
			// overlaps2 can have overlaps that begin
			index11 = 0
			index2 = 0
			for index11 < count1 && index2 < count2 {
				r1 = refs1 + uintptr(index11)*8
				r2 = refs2 + uintptr(index2)*8
				if (*b2ShapeRef)(unsafe.Pointer(r1)).ShapeId == (*b2ShapeRef)(unsafe.Pointer(r2)).ShapeId {
					if int32FromUint16((*b2ShapeRef)(unsafe.Pointer(r1)).Generation) < int32FromUint16((*b2ShapeRef)(unsafe.Pointer(r2)).Generation) {
						// end
						visitorId = ShapeId{
							Index1:     (*b2ShapeRef)(unsafe.Pointer(r1)).ShapeId + int32(1),
							World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
							Generation: (*b2ShapeRef)(unsafe.Pointer(r1)).Generation,
						}
						event = SensorEndTouchEvent{
							SensorShapeId:  sensorId,
							VisitorShapeId: visitorId,
						}
						v10 = world + 1376 + uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16
						if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v10)).Count == (*b2SensorEndTouchEventArray)(unsafe.Pointer(v10)).Capacity {
							if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v10)).Capacity < int32(2) {
								v11 = int32(2)
							} else {
								v11 = (*b2SensorEndTouchEventArray)(unsafe.Pointer(v10)).Capacity + (*b2SensorEndTouchEventArray)(unsafe.Pointer(v10)).Capacity>>int32(1)
							}
							newCapacity1 = v11
							b2SensorEndTouchEventArray_Reserve(tls, v10, newCapacity1)
						}
						*(*SensorEndTouchEvent)(unsafe.Pointer((*b2SensorEndTouchEventArray)(unsafe.Pointer(v10)).Data + uintptr((*b2SensorEndTouchEventArray)(unsafe.Pointer(v10)).Count)*16)) = event
						*(*int32)(unsafe.Pointer(v10 + 8)) += int32(1)
						index11 += int32(1)
					} else {
						if int32FromUint16((*b2ShapeRef)(unsafe.Pointer(r1)).Generation) > int32FromUint16((*b2ShapeRef)(unsafe.Pointer(r2)).Generation) {
							// begin
							visitorId1 = ShapeId{
								Index1:     (*b2ShapeRef)(unsafe.Pointer(r2)).ShapeId + int32(1),
								World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
								Generation: (*b2ShapeRef)(unsafe.Pointer(r2)).Generation,
							}
							event1 = SensorBeginTouchEvent{
								SensorShapeId:  sensorId,
								VisitorShapeId: visitorId1,
							}
							v12 = world + 1344
							if (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v12)).Count == (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v12)).Capacity {
								if (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v12)).Capacity < int32(2) {
									v13 = int32(2)
								} else {
									v13 = (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v12)).Capacity + (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v12)).Capacity>>int32(1)
								}
								newCapacity = v13
								b2SensorBeginTouchEventArray_Reserve(tls, v12, newCapacity)
							}
							*(*SensorBeginTouchEvent)(unsafe.Pointer((*b2SensorBeginTouchEventArray)(unsafe.Pointer(v12)).Data + uintptr((*b2SensorBeginTouchEventArray)(unsafe.Pointer(v12)).Count)*16)) = event1
							*(*int32)(unsafe.Pointer(v12 + 8)) += int32(1)
							index2 += int32(1)
						} else {
							// persisted
							index11 += int32(1)
							index2 += int32(1)
						}
					}
				} else {
					if (*b2ShapeRef)(unsafe.Pointer(r1)).ShapeId < (*b2ShapeRef)(unsafe.Pointer(r2)).ShapeId {
						// end
						visitorId2 = ShapeId{
							Index1:     (*b2ShapeRef)(unsafe.Pointer(r1)).ShapeId + int32(1),
							World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
							Generation: (*b2ShapeRef)(unsafe.Pointer(r1)).Generation,
						}
						event2 = SensorEndTouchEvent{
							SensorShapeId:  sensorId,
							VisitorShapeId: visitorId2,
						}
						v14 = world + 1376 + uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16
						if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v14)).Count == (*b2SensorEndTouchEventArray)(unsafe.Pointer(v14)).Capacity {
							if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v14)).Capacity < int32(2) {
								v15 = int32(2)
							} else {
								v15 = (*b2SensorEndTouchEventArray)(unsafe.Pointer(v14)).Capacity + (*b2SensorEndTouchEventArray)(unsafe.Pointer(v14)).Capacity>>int32(1)
							}
							newCapacity1 = v15
							b2SensorEndTouchEventArray_Reserve(tls, v14, newCapacity1)
						}
						*(*SensorEndTouchEvent)(unsafe.Pointer((*b2SensorEndTouchEventArray)(unsafe.Pointer(v14)).Data + uintptr((*b2SensorEndTouchEventArray)(unsafe.Pointer(v14)).Count)*16)) = event2
						*(*int32)(unsafe.Pointer(v14 + 8)) += int32(1)
						index11 += int32(1)
					} else {
						// begin
						visitorId3 = ShapeId{
							Index1:     (*b2ShapeRef)(unsafe.Pointer(r2)).ShapeId + int32(1),
							World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
							Generation: (*b2ShapeRef)(unsafe.Pointer(r2)).Generation,
						}
						event3 = SensorBeginTouchEvent{
							SensorShapeId:  sensorId,
							VisitorShapeId: visitorId3,
						}
						v16 = world + 1344
						if (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v16)).Count == (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v16)).Capacity {
							if (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v16)).Capacity < int32(2) {
								v17 = int32(2)
							} else {
								v17 = (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v16)).Capacity + (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v16)).Capacity>>int32(1)
							}
							newCapacity = v17
							b2SensorBeginTouchEventArray_Reserve(tls, v16, newCapacity)
						}
						*(*SensorBeginTouchEvent)(unsafe.Pointer((*b2SensorBeginTouchEventArray)(unsafe.Pointer(v16)).Data + uintptr((*b2SensorBeginTouchEventArray)(unsafe.Pointer(v16)).Count)*16)) = event3
						*(*int32)(unsafe.Pointer(v16 + 8)) += int32(1)
						index2 += int32(1)
					}
				}
			}
			for index11 < count1 {
				// end
				r11 = refs1 + uintptr(index11)*8
				visitorId4 = ShapeId{
					Index1:     (*b2ShapeRef)(unsafe.Pointer(r11)).ShapeId + int32(1),
					World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
					Generation: (*b2ShapeRef)(unsafe.Pointer(r11)).Generation,
				}
				event4 = SensorEndTouchEvent{
					SensorShapeId:  sensorId,
					VisitorShapeId: visitorId4,
				}
				v18 = world + 1376 + uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16
				if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Count == (*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Capacity {
					if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Capacity < int32(2) {
						v19 = int32(2)
					} else {
						v19 = (*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Capacity + (*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Capacity>>int32(1)
					}
					newCapacity1 = v19
					b2SensorEndTouchEventArray_Reserve(tls, v18, newCapacity1)
				}
				*(*SensorEndTouchEvent)(unsafe.Pointer((*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Data + uintptr((*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Count)*16)) = event4
				*(*int32)(unsafe.Pointer(v18 + 8)) += int32(1)
				index11 += int32(1)
			}
			for index2 < count2 {
				// begin
				r21 = refs2 + uintptr(index2)*8
				visitorId5 = ShapeId{
					Index1:     (*b2ShapeRef)(unsafe.Pointer(r21)).ShapeId + int32(1),
					World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
					Generation: (*b2ShapeRef)(unsafe.Pointer(r21)).Generation,
				}
				event5 = SensorBeginTouchEvent{
					SensorShapeId:  sensorId,
					VisitorShapeId: visitorId5,
				}
				v20 = world + 1344
				if (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v20)).Count == (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v20)).Capacity {
					if (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v20)).Capacity < int32(2) {
						v21 = int32(2)
					} else {
						v21 = (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v20)).Capacity + (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v20)).Capacity>>int32(1)
					}
					newCapacity = v21
					b2SensorBeginTouchEventArray_Reserve(tls, v20, newCapacity)
				}
				*(*SensorBeginTouchEvent)(unsafe.Pointer((*b2SensorBeginTouchEventArray)(unsafe.Pointer(v20)).Data + uintptr((*b2SensorBeginTouchEventArray)(unsafe.Pointer(v20)).Count)*16)) = event5
				*(*int32)(unsafe.Pointer(v20 + 8)) += int32(1)
				index2 += int32(1)
			}
			// Clear the smallest set bit
			word = word & (word - uint64(1))
		}
		goto _3
	_3:
		;
		k++
	}
}

func b2DestroySensor(tls *_Stack, world uintptr, sensorShape uintptr) {
	var event SensorEndTouchEvent
	var i, movedIndex, movedIndex1, newCapacity, v5, v7, v8 int32
	var movedSensor, otherSensorShape, ref, sensor, v1, v10, v12, v4, v6 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = event, i, movedIndex, movedIndex1, movedSensor, newCapacity, otherSensorShape, ref, sensor, v1, v10, v12, v4, v5, v6, v7, v8
	v1 = (*b2SensorArray)(unsafe.Pointer(world+1280)).Data + uintptr((*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex)*40
	goto _2
_2:
	sensor = v1
	i = 0
	for {
		if !(i < (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count) {
			break
		}
		ref = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data + uintptr(i)*8
		event = SensorEndTouchEvent{
			SensorShapeId: ShapeId{
				Index1:     (*b2Shape)(unsafe.Pointer(sensorShape)).Id + int32(1),
				World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
				Generation: (*b2Shape)(unsafe.Pointer(sensorShape)).Generation,
			},
			VisitorShapeId: ShapeId{
				Index1:     (*b2ShapeRef)(unsafe.Pointer(ref)).ShapeId + int32(1),
				World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
				Generation: (*b2ShapeRef)(unsafe.Pointer(ref)).Generation,
			},
		}
		v4 = world + 1376 + uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16
		if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v4)).Count == (*b2SensorEndTouchEventArray)(unsafe.Pointer(v4)).Capacity {
			if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v4)).Capacity < int32(2) {
				v5 = int32(2)
			} else {
				v5 = (*b2SensorEndTouchEventArray)(unsafe.Pointer(v4)).Capacity + (*b2SensorEndTouchEventArray)(unsafe.Pointer(v4)).Capacity>>int32(1)
			}
			newCapacity = v5
			b2SensorEndTouchEventArray_Reserve(tls, v4, newCapacity)
		}
		*(*SensorEndTouchEvent)(unsafe.Pointer((*b2SensorEndTouchEventArray)(unsafe.Pointer(v4)).Data + uintptr((*b2SensorEndTouchEventArray)(unsafe.Pointer(v4)).Count)*16)) = event
		*(*int32)(unsafe.Pointer(v4 + 8)) += int32(1)
		goto _3
	_3:
		;
		i++
	}
	// Destroy sensor
	b2ShapeRefArray_Destroy(tls, sensor)
	b2ShapeRefArray_Destroy(tls, sensor+16)
	v6 = world + 1280
	v7 = (*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex
	movedIndex = -int32(1)
	if v7 != (*b2SensorArray)(unsafe.Pointer(v6)).Count-int32FromInt32(1) {
		movedIndex = (*b2SensorArray)(unsafe.Pointer(v6)).Count - int32(1)
		*(*b2Sensor)(unsafe.Pointer((*b2SensorArray)(unsafe.Pointer(v6)).Data + uintptr(v7)*40)) = *(*b2Sensor)(unsafe.Pointer((*b2SensorArray)(unsafe.Pointer(v6)).Data + uintptr(movedIndex)*40))
	}
	*(*int32)(unsafe.Pointer(v6 + 8)) -= int32(1)
	v8 = movedIndex
	goto _9
_9:
	movedIndex1 = v8
	if movedIndex1 != -int32(1) {
		v10 = (*b2SensorArray)(unsafe.Pointer(world+1280)).Data + uintptr((*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex)*40
		goto _11
	_11:
		// Fixup moved sensor
		movedSensor = v10
		v12 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr((*b2Sensor)(unsafe.Pointer(movedSensor)).ShapeId)*288
		goto _13
	_13:
		otherSensorShape = v12
		(*b2Shape)(unsafe.Pointer(otherSensorShape)).SensorIndex = (*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex
	}
}

func b2SensorBeginTouchEventArray_Create(tls *_Stack, capacity int32) (r b2SensorBeginTouchEventArray) {
	var a b2SensorBeginTouchEventArray
	_ = a
	a = b2SensorBeginTouchEventArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(16)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2SensorBeginTouchEventArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Capacity)*uint64(16)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(16)))
	(*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2SensorBeginTouchEventArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Capacity)*uint64(16)))
	(*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Count = 0
	(*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2SensorEndTouchEventArray_Create(tls *_Stack, capacity int32) (r b2SensorEndTouchEventArray) {
	var a b2SensorEndTouchEventArray
	_ = a
	a = b2SensorEndTouchEventArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(16)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2SensorEndTouchEventArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Capacity)*uint64(16)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(16)))
	(*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2SensorEndTouchEventArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Capacity)*uint64(16)))
	(*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Count = 0
	(*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Capacity = 0
}
