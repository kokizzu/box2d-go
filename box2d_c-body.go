package box2d

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

// C documentation
//
//	// Implement functions for b2BodyArray
func b2BodyArray_Create(tls *_Stack, capacity int32) (r b2BodyArray) {
	var a b2BodyArray
	_ = a
	a = b2BodyArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(128)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2BodyArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2BodyArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2BodyArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2BodyArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2BodyArray)(unsafe.Pointer(a)).Capacity)*uint64(128)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(128)))
	(*b2BodyArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2BodyArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2BodyArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2BodyArray)(unsafe.Pointer(a)).Capacity)*uint64(128)))
	(*b2BodyArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2BodyArray)(unsafe.Pointer(a)).Count = 0
	(*b2BodyArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2BodySimArray_Create(tls *_Stack, capacity int32) (r b2BodySimArray) {
	var a b2BodySimArray
	_ = a
	a = b2BodySimArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(100)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2BodySimArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2BodySimArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2BodySimArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2BodySimArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2BodySimArray)(unsafe.Pointer(a)).Capacity)*uint64(100)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(100)))
	(*b2BodySimArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2BodySimArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2BodySimArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2BodySimArray)(unsafe.Pointer(a)).Capacity)*uint64(100)))
	(*b2BodySimArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2BodySimArray)(unsafe.Pointer(a)).Count = 0
	(*b2BodySimArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2BodyStateArray_Create(tls *_Stack, capacity int32) (r b2BodyStateArray) {
	var a b2BodyStateArray
	_ = a
	a = b2BodyStateArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(32)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2BodyStateArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2BodyStateArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2BodyStateArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2BodyStateArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2BodyStateArray)(unsafe.Pointer(a)).Capacity)*uint64(32)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(32)))
	(*b2BodyStateArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2BodyStateArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2BodyStateArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2BodyStateArray)(unsafe.Pointer(a)).Capacity)*uint64(32)))
	(*b2BodyStateArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2BodyStateArray)(unsafe.Pointer(a)).Count = 0
	(*b2BodyStateArray)(unsafe.Pointer(a)).Capacity = 0
}

// C documentation
//
//	// Get a validated body from a world using an id.
func b2GetBodyFullId(tls *_Stack, world uintptr, bodyId BodyId) (r uintptr) {
	var v1 uintptr
	_ = v1
	// id index starts at one so that zero can represent null
	v1 = (*b2BodyArray)(unsafe.Pointer(world+1024)).Data + uintptr(bodyId.Index1-int32(1))*128
	goto _2
_2:
	return v1
}

func b2GetBodyTransformQuick(tls *_Stack, world uintptr, body uintptr) (r Transform) {
	var bodySim, set, v1, v3 uintptr
	_, _, _, _ = bodySim, set, v1, v3
	v1 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr((*b2Body)(unsafe.Pointer(body)).SetIndex)*88
	goto _2
_2:
	set = v1
	v3 = (*b2BodySimArray)(unsafe.Pointer(set)).Data + uintptr((*b2Body)(unsafe.Pointer(body)).LocalIndex)*100
	goto _4
_4:
	bodySim = v3
	return (*b2BodySim)(unsafe.Pointer(bodySim)).Transform
}

func b2GetBodyTransform(tls *_Stack, world uintptr, bodyId int32) (r Transform) {
	var body, v1 uintptr
	_, _ = body, v1
	v1 = (*b2BodyArray)(unsafe.Pointer(world+1024)).Data + uintptr(bodyId)*128
	goto _2
_2:
	body = v1
	return b2GetBodyTransformQuick(tls, world, body)
}

// C documentation
//
//	// Create a b2BodyId from a raw id.
func b2MakeBodyId(tls *_Stack, world uintptr, bodyId int32) (r BodyId) {
	var body, v1 uintptr
	_, _ = body, v1
	v1 = (*b2BodyArray)(unsafe.Pointer(world+1024)).Data + uintptr(bodyId)*128
	goto _2
_2:
	body = v1
	return BodyId{
		Index1:     bodyId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Body)(unsafe.Pointer(body)).Generation,
	}
}

func b2GetBodySim(tls *_Stack, world uintptr, body uintptr) (r uintptr) {
	var bodySim, set, v1, v3 uintptr
	_, _, _, _ = bodySim, set, v1, v3
	v1 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr((*b2Body)(unsafe.Pointer(body)).SetIndex)*88
	goto _2
_2:
	set = v1
	v3 = (*b2BodySimArray)(unsafe.Pointer(set)).Data + uintptr((*b2Body)(unsafe.Pointer(body)).LocalIndex)*100
	goto _4
_4:
	bodySim = v3
	return bodySim
}

func b2GetBodyState(tls *_Stack, world uintptr, body uintptr) (r uintptr) {
	var set, v1, v3 uintptr
	_, _, _ = set, v1, v3
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		v1 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr(int32(b2_awakeSet))*88
		goto _2
	_2:
		set = v1
		v3 = (*b2BodyStateArray)(unsafe.Pointer(set+16)).Data + uintptr((*b2Body)(unsafe.Pointer(body)).LocalIndex)*32
		goto _4
	_4:
		return v3
	}
	return uintptrFromInt32(0)
}

func b2CreateIslandForBody(tls *_Stack, world uintptr, setIndex int32, body uintptr) {
	var island uintptr
	_ = island
	island = b2CreateIsland(tls, world, setIndex)
	(*b2Body)(unsafe.Pointer(body)).IslandId = (*b2Island)(unsafe.Pointer(island)).IslandId
	(*b2Island)(unsafe.Pointer(island)).HeadBody = (*b2Body)(unsafe.Pointer(body)).Id
	(*b2Island)(unsafe.Pointer(island)).TailBody = (*b2Body)(unsafe.Pointer(body)).Id
	(*b2Island)(unsafe.Pointer(island)).BodyCount = int32(1)
}

func b2RemoveBodyFromIsland(tls *_Stack, world uintptr, body uintptr) {
	var island, nextBody, prevBody, v1, v3, v5 uintptr
	var islandDestroyed uint8
	var islandId int32
	_, _, _, _, _, _, _, _ = island, islandDestroyed, islandId, nextBody, prevBody, v1, v3, v5
	if (*b2Body)(unsafe.Pointer(body)).IslandId == -int32(1) {
		return
	}
	islandId = (*b2Body)(unsafe.Pointer(body)).IslandId
	v1 = (*b2IslandArray)(unsafe.Pointer(world+1184)).Data + uintptr(islandId)*56
	goto _2
_2:
	island = v1
	// Fix the island's linked list of sims
	if (*b2Body)(unsafe.Pointer(body)).IslandPrev != -int32(1) {
		v3 = (*b2BodyArray)(unsafe.Pointer(world+1024)).Data + uintptr((*b2Body)(unsafe.Pointer(body)).IslandPrev)*128
		goto _4
	_4:
		prevBody = v3
		(*b2Body)(unsafe.Pointer(prevBody)).IslandNext = (*b2Body)(unsafe.Pointer(body)).IslandNext
	}
	if (*b2Body)(unsafe.Pointer(body)).IslandNext != -int32(1) {
		v5 = (*b2BodyArray)(unsafe.Pointer(world+1024)).Data + uintptr((*b2Body)(unsafe.Pointer(body)).IslandNext)*128
		goto _6
	_6:
		nextBody = v5
		(*b2Body)(unsafe.Pointer(nextBody)).IslandPrev = (*b2Body)(unsafe.Pointer(body)).IslandPrev
	}
	*(*int32)(unsafe.Pointer(island + 20)) -= int32(1)
	islandDestroyed = uint8(false1)
	if (*b2Island)(unsafe.Pointer(island)).HeadBody == (*b2Body)(unsafe.Pointer(body)).Id {
		(*b2Island)(unsafe.Pointer(island)).HeadBody = (*b2Body)(unsafe.Pointer(body)).IslandNext
		if (*b2Island)(unsafe.Pointer(island)).HeadBody == -int32(1) {
			// Destroy empty island
			// Free the island
			b2DestroyIsland(tls, world, (*b2Island)(unsafe.Pointer(island)).IslandId)
			islandDestroyed = uint8(true1)
		}
	} else {
		if (*b2Island)(unsafe.Pointer(island)).TailBody == (*b2Body)(unsafe.Pointer(body)).Id {
			(*b2Island)(unsafe.Pointer(island)).TailBody = (*b2Body)(unsafe.Pointer(body)).IslandPrev
		}
	}
	if int32FromUint8(islandDestroyed) == false1 {
		b2ValidateIsland(tls, world, islandId)
	}
	(*b2Body)(unsafe.Pointer(body)).IslandId = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).IslandPrev = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).IslandNext = -int32(1)
}

func b2DestroyBodyContacts(tls *_Stack, world uintptr, body uintptr, wakeBodies uint8) {
	var contact, v1 uintptr
	var contactId, edgeIndex, edgeKey int32
	_, _, _, _, _ = contact, contactId, edgeIndex, edgeKey, v1
	// Destroy the attached contacts
	edgeKey = (*b2Body)(unsafe.Pointer(body)).HeadContactKey
	for edgeKey != -int32(1) {
		contactId = edgeKey >> int32(1)
		edgeIndex = edgeKey & int32(1)
		v1 = (*b2ContactArray)(unsafe.Pointer(world+1144)).Data + uintptr(contactId)*68
		goto _2
	_2:
		contact = v1
		edgeKey = (*(*b2ContactEdge)(unsafe.Pointer(contact + 12 + uintptr(edgeIndex)*12))).NextKey
		b2DestroyContact(tls, world, contact, wakeBodies)
	}
	b2ValidateSolverSets(tls, world)
}

func b2CreateBody(tls *_Stack, worldId WorldId, def uintptr) (r BodyId) {
	var body, bodySim, bodyState, set, world, v1, v11, v13, v15, v3, v5, v7, v9, p17 uintptr
	var bodyId, i, newCapacity, newCapacity1, newCapacity2, newCapacity3, setId, v10, v14, v2, v6 int32
	var id BodyId
	var isAwake uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodyId, bodySim, bodyState, i, id, isAwake, newCapacity, newCapacity1, newCapacity2, newCapacity3, set, setId, world, v1, v10, v11, v13, v14, v15, v2, v3, v5, v6, v7, v9, p17
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return b2_nullBodyId
	}
	isAwake = boolUint8(((*BodyDef)(unsafe.Pointer(def)).IsAwake != 0 || int32FromUint8((*BodyDef)(unsafe.Pointer(def)).EnableSleep) == false1) && (*BodyDef)(unsafe.Pointer(def)).IsEnabled != 0)
	if int32FromUint8((*BodyDef)(unsafe.Pointer(def)).IsEnabled) == false1 {
		// any body type can be disabled
		setId = int32(b2_disabledSet)
	} else {
		if (*BodyDef)(unsafe.Pointer(def)).Type1 == int32(b2_staticBody) {
			setId = int32(b2_staticSet)
		} else {
			if int32FromUint8(isAwake) == int32(true1) {
				setId = int32(b2_awakeSet)
			} else {
				// new set for a sleeping body in its own island
				setId = b2AllocId(tls, world+1040)
				if setId == (*b2World)(unsafe.Pointer(world)).SolverSets.Count {
					// Create a zero initialized solver set. All sub-arrays are also zero initialized.
					v1 = world + 1064
					if (*b2SolverSetArray)(unsafe.Pointer(v1)).Count == (*b2SolverSetArray)(unsafe.Pointer(v1)).Capacity {
						if (*b2SolverSetArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
							v2 = int32(2)
						} else {
							v2 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Capacity + (*b2SolverSetArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
						}
						newCapacity3 = v2
						b2SolverSetArray_Reserve(tls, v1, newCapacity3)
					}
					*(*b2SolverSet)(unsafe.Pointer((*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr((*b2SolverSetArray)(unsafe.Pointer(v1)).Count)*88)) = b2SolverSet{}
					*(*int32)(unsafe.Pointer(v1 + 8)) += int32(1)
				} else {
				}
				(*(*b2SolverSet)(unsafe.Pointer((*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setId)*88))).SetIndex = setId
			}
		}
	}
	bodyId = b2AllocId(tls, world+1000)
	v3 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr(setId)*88
	goto _4
_4:
	set = v3
	v5 = set
	if (*b2BodySimArray)(unsafe.Pointer(v5)).Count == (*b2BodySimArray)(unsafe.Pointer(v5)).Capacity {
		if (*b2BodySimArray)(unsafe.Pointer(v5)).Capacity < int32(2) {
			v6 = int32(2)
		} else {
			v6 = (*b2BodySimArray)(unsafe.Pointer(v5)).Capacity + (*b2BodySimArray)(unsafe.Pointer(v5)).Capacity>>int32(1)
		}
		newCapacity1 = v6
		b2BodySimArray_Reserve(tls, v5, newCapacity1)
	}
	*(*int32)(unsafe.Pointer(v5 + 8)) += int32(1)
	v7 = (*b2BodySimArray)(unsafe.Pointer(v5)).Data + uintptr((*b2BodySimArray)(unsafe.Pointer(v5)).Count-int32FromInt32(1))*100
	goto _8
_8:
	bodySim = v7
	*(*b2BodySim)(unsafe.Pointer(bodySim)) = b2BodySim{}
	(*b2BodySim)(unsafe.Pointer(bodySim)).Transform.P = (*BodyDef)(unsafe.Pointer(def)).Position
	(*b2BodySim)(unsafe.Pointer(bodySim)).Transform.Q = (*BodyDef)(unsafe.Pointer(def)).Rotation
	(*b2BodySim)(unsafe.Pointer(bodySim)).Center = (*BodyDef)(unsafe.Pointer(def)).Position
	(*b2BodySim)(unsafe.Pointer(bodySim)).Rotation0 = (*b2BodySim)(unsafe.Pointer(bodySim)).Transform.Q
	(*b2BodySim)(unsafe.Pointer(bodySim)).Center0 = (*b2BodySim)(unsafe.Pointer(bodySim)).Center
	(*b2BodySim)(unsafe.Pointer(bodySim)).MinExtent = float32(float32FromFloat32(100000) * b2_lengthUnitsPerMeter)
	(*b2BodySim)(unsafe.Pointer(bodySim)).MaxExtent = float32FromFloat32(0)
	(*b2BodySim)(unsafe.Pointer(bodySim)).LinearDamping = (*BodyDef)(unsafe.Pointer(def)).LinearDamping
	(*b2BodySim)(unsafe.Pointer(bodySim)).AngularDamping = (*BodyDef)(unsafe.Pointer(def)).AngularDamping
	(*b2BodySim)(unsafe.Pointer(bodySim)).GravityScale = (*BodyDef)(unsafe.Pointer(def)).GravityScale
	(*b2BodySim)(unsafe.Pointer(bodySim)).BodyId = bodyId
	(*b2BodySim)(unsafe.Pointer(bodySim)).IsBullet = (*BodyDef)(unsafe.Pointer(def)).IsBullet
	(*b2BodySim)(unsafe.Pointer(bodySim)).AllowFastRotation = (*BodyDef)(unsafe.Pointer(def)).AllowFastRotation
	if setId == int32(b2_awakeSet) {
		v9 = set + 16
		if (*b2BodyStateArray)(unsafe.Pointer(v9)).Count == (*b2BodyStateArray)(unsafe.Pointer(v9)).Capacity {
			if (*b2BodyStateArray)(unsafe.Pointer(v9)).Capacity < int32(2) {
				v10 = int32(2)
			} else {
				v10 = (*b2BodyStateArray)(unsafe.Pointer(v9)).Capacity + (*b2BodyStateArray)(unsafe.Pointer(v9)).Capacity>>int32(1)
			}
			newCapacity2 = v10
			b2BodyStateArray_Reserve(tls, v9, newCapacity2)
		}
		*(*int32)(unsafe.Pointer(v9 + 8)) += int32(1)
		v11 = (*b2BodyStateArray)(unsafe.Pointer(v9)).Data + uintptr((*b2BodyStateArray)(unsafe.Pointer(v9)).Count-int32FromInt32(1))*32
		goto _12
	_12:
		bodyState = v11
		*(*b2BodyState)(unsafe.Pointer(bodyState)) = b2BodyState{}
		(*b2BodyState)(unsafe.Pointer(bodyState)).LinearVelocity = (*BodyDef)(unsafe.Pointer(def)).LinearVelocity
		(*b2BodyState)(unsafe.Pointer(bodyState)).AngularVelocity = (*BodyDef)(unsafe.Pointer(def)).AngularVelocity
		(*b2BodyState)(unsafe.Pointer(bodyState)).DeltaRotation = b2Rot_identity
	}
	if bodyId == (*b2World)(unsafe.Pointer(world)).Bodies.Count {
		v13 = world + 1024
		if (*b2BodyArray)(unsafe.Pointer(v13)).Count == (*b2BodyArray)(unsafe.Pointer(v13)).Capacity {
			if (*b2BodyArray)(unsafe.Pointer(v13)).Capacity < int32(2) {
				v14 = int32(2)
			} else {
				v14 = (*b2BodyArray)(unsafe.Pointer(v13)).Capacity + (*b2BodyArray)(unsafe.Pointer(v13)).Capacity>>int32(1)
			}
			newCapacity = v14
			b2BodyArray_Reserve(tls, v13, newCapacity)
		}
		*(*b2Body)(unsafe.Pointer((*b2BodyArray)(unsafe.Pointer(v13)).Data + uintptr((*b2BodyArray)(unsafe.Pointer(v13)).Count)*128)) = b2Body{}
		*(*int32)(unsafe.Pointer(v13 + 8)) += int32(1)
	} else {
	}
	v15 = (*b2BodyArray)(unsafe.Pointer(world+1024)).Data + uintptr(bodyId)*128
	goto _16
_16:
	body = v15
	if (*BodyDef)(unsafe.Pointer(def)).Name != 0 {
		i = 0
		for i < int32(31) && int32FromUint8(*(*uint8)(unsafe.Pointer((*BodyDef)(unsafe.Pointer(def)).Name + uintptr(i)))) != 0 {
			*(*uint8)(unsafe.Pointer(body + uintptr(i))) = *(*uint8)(unsafe.Pointer((*BodyDef)(unsafe.Pointer(def)).Name + uintptr(i)))
			i += int32(1)
		}
		for i < int32(32) {
			*(*uint8)(unsafe.Pointer(body + uintptr(i))) = uint8(0)
			i += int32(1)
		}
	} else {
		memset(tls, body, 0, uint64FromInt32(32)*uint64FromInt64(1))
	}
	(*b2Body)(unsafe.Pointer(body)).UserData = (*BodyDef)(unsafe.Pointer(def)).UserData
	(*b2Body)(unsafe.Pointer(body)).SetIndex = setId
	(*b2Body)(unsafe.Pointer(body)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count - int32(1)
	p17 = body + 116
	*(*uint16_t)(unsafe.Pointer(p17)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(p17))) + int32FromInt32(1))
	(*b2Body)(unsafe.Pointer(body)).HeadShapeId = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).ShapeCount = 0
	(*b2Body)(unsafe.Pointer(body)).HeadChainId = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).HeadContactKey = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).ContactCount = 0
	(*b2Body)(unsafe.Pointer(body)).HeadJointKey = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).JointCount = 0
	(*b2Body)(unsafe.Pointer(body)).IslandId = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).IslandPrev = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).IslandNext = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).BodyMoveIndex = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).Id = bodyId
	(*b2Body)(unsafe.Pointer(body)).Mass = float32FromFloat32(0)
	(*b2Body)(unsafe.Pointer(body)).Inertia = float32FromFloat32(0)
	(*b2Body)(unsafe.Pointer(body)).SleepThreshold = (*BodyDef)(unsafe.Pointer(def)).SleepThreshold
	(*b2Body)(unsafe.Pointer(body)).SleepTime = float32FromFloat32(0)
	(*b2Body)(unsafe.Pointer(body)).Type1 = (*BodyDef)(unsafe.Pointer(def)).Type1
	(*b2Body)(unsafe.Pointer(body)).EnableSleep = (*BodyDef)(unsafe.Pointer(def)).EnableSleep
	(*b2Body)(unsafe.Pointer(body)).FixedRotation = (*BodyDef)(unsafe.Pointer(def)).FixedRotation
	(*b2Body)(unsafe.Pointer(body)).IsSpeedCapped = uint8(false1)
	(*b2Body)(unsafe.Pointer(body)).IsMarked = uint8(false1)
	// dynamic and kinematic bodies that are enabled need a island
	if setId >= int32(b2_awakeSet) {
		b2CreateIslandForBody(tls, world, setId, body)
	}
	b2ValidateSolverSets(tls, world)
	id = BodyId{
		Index1:     bodyId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Body)(unsafe.Pointer(body)).Generation,
	}
	return id
}

func b2WakeBody(tls *_Stack, world uintptr, body uintptr) (r uint8) {
	if (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeSolverSet(tls, world, (*b2Body)(unsafe.Pointer(body)).SetIndex)
		return uint8(true1)
	}
	return uint8(false1)
}

func b2DestroyBody(tls *_Stack, bodyId BodyId) {
	var body, chain, joint, movedBody, movedSim, set, shape, world, v1, v13, v15, v3, v5, v7, v9 uintptr
	var chainId, edgeIndex, edgeKey, jointId, movedId, movedIndex, movedIndex1, movedIndex2, result, shapeId, v10, v11, v16, v17 int32
	var wakeBodies uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, chain, chainId, edgeIndex, edgeKey, joint, jointId, movedBody, movedId, movedIndex, movedIndex1, movedIndex2, movedSim, result, set, shape, shapeId, wakeBodies, world, v1, v10, v11, v13, v15, v16, v17, v3, v5, v7, v9
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	// Wake bodies attached to this body, even if this body is static.
	wakeBodies = uint8(true1)
	// Destroy the attached joints
	edgeKey = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
	for edgeKey != -int32(1) {
		jointId = edgeKey >> int32(1)
		edgeIndex = edgeKey & int32(1)
		v1 = (*b2JointArray)(unsafe.Pointer(world+1104)).Data + uintptr(jointId)*72
		goto _2
	_2:
		joint = v1
		edgeKey = (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + uintptr(edgeIndex)*12))).NextKey
		// Careful because this modifies the list being traversed
		b2DestroyJointInternal(tls, world, joint, wakeBodies)
	}
	// Destroy all contacts attached to this body.
	b2DestroyBodyContacts(tls, world, body, wakeBodies)
	// Destroy the attached shapes and their broad-phase proxies.
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		v3 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr(shapeId)*288
		goto _4
	_4:
		shape = v3
		if (*b2Shape)(unsafe.Pointer(shape)).SensorIndex != -int32(1) {
			b2DestroySensor(tls, world, shape)
		}
		b2DestroyShapeProxy(tls, shape, world+40)
		// Return shape to free list.
		b2FreeId(tls, world+1200, shapeId)
		(*b2Shape)(unsafe.Pointer(shape)).Id = -int32(1)
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
	}
	// Destroy the attached chains. The associated shapes have already been destroyed above.
	chainId = (*b2Body)(unsafe.Pointer(body)).HeadChainId
	for chainId != -int32(1) {
		v5 = (*b2ChainShapeArray)(unsafe.Pointer(world+1264)).Data + uintptr(chainId)*48
		goto _6
	_6:
		chain = v5
		b2FreeChainData(tls, chain)
		// Return chain to free list.
		b2FreeId(tls, world+1224, chainId)
		(*b2ChainShape)(unsafe.Pointer(chain)).Id = -int32(1)
		chainId = (*b2ChainShape)(unsafe.Pointer(chain)).NextChainId
	}
	b2RemoveBodyFromIsland(tls, world, body)
	v7 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr((*b2Body)(unsafe.Pointer(body)).SetIndex)*88
	goto _8
_8:
	// Remove body sim from solver set that owns it
	set = v7
	v9 = set
	v10 = (*b2Body)(unsafe.Pointer(body)).LocalIndex
	movedIndex = -int32(1)
	if v10 != (*b2BodySimArray)(unsafe.Pointer(v9)).Count-int32FromInt32(1) {
		movedIndex = (*b2BodySimArray)(unsafe.Pointer(v9)).Count - int32(1)
		*(*b2BodySim)(unsafe.Pointer((*b2BodySimArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*100)) = *(*b2BodySim)(unsafe.Pointer((*b2BodySimArray)(unsafe.Pointer(v9)).Data + uintptr(movedIndex)*100))
	}
	*(*int32)(unsafe.Pointer(v9 + 8)) -= int32(1)
	v11 = movedIndex
	goto _12
_12:
	movedIndex2 = v11
	if movedIndex2 != -int32(1) {
		// Fix moved body index
		movedSim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(body)).LocalIndex)*100
		movedId = (*b2BodySim)(unsafe.Pointer(movedSim)).BodyId
		v13 = (*b2BodyArray)(unsafe.Pointer(world+1024)).Data + uintptr(movedId)*128
		goto _14
	_14:
		movedBody = v13
		(*b2Body)(unsafe.Pointer(movedBody)).LocalIndex = (*b2Body)(unsafe.Pointer(body)).LocalIndex
	}
	// Remove body state from awake set
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		v15 = set + 16
		v16 = (*b2Body)(unsafe.Pointer(body)).LocalIndex
		movedIndex1 = -int32(1)
		if v16 != (*b2BodyStateArray)(unsafe.Pointer(v15)).Count-int32FromInt32(1) {
			movedIndex1 = (*b2BodyStateArray)(unsafe.Pointer(v15)).Count - int32(1)
			*(*b2BodyState)(unsafe.Pointer((*b2BodyStateArray)(unsafe.Pointer(v15)).Data + uintptr(v16)*32)) = *(*b2BodyState)(unsafe.Pointer((*b2BodyStateArray)(unsafe.Pointer(v15)).Data + uintptr(movedIndex1)*32))
		}
		*(*int32)(unsafe.Pointer(v15 + 8)) -= int32(1)
		v17 = movedIndex1
		goto _18
	_18:
		result = v17
		_ = uint64FromInt64(4)
	} else {
		if (*b2SolverSet)(unsafe.Pointer(set)).SetIndex >= int32(b2_firstSleepingSet) && (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count == 0 {
			// Remove solver set if it's now an orphan.
			b2DestroySolverSet(tls, world, (*b2SolverSet)(unsafe.Pointer(set)).SetIndex)
		}
	}
	// Free body and id (preserve body generation)
	b2FreeId(tls, world+1000, (*b2Body)(unsafe.Pointer(body)).Id)
	(*b2Body)(unsafe.Pointer(body)).SetIndex = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).LocalIndex = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).Id = -int32(1)
	b2ValidateSolverSets(tls, world)
}

func b2Body_GetContactCapacity(tls *_Stack, bodyId BodyId) (r int32) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	// Conservative and fast
	return (*b2Body)(unsafe.Pointer(body)).ContactCount
}

func b2Body_GetContactData(tls *_Stack, bodyId BodyId, contactData uintptr, capacity int32) (r int32) {
	var body, contact, contactSim, shapeA, shapeB, world, v1, v3, v5 uintptr
	var contactId, contactKey, edgeIndex, index2 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _ = body, contact, contactId, contactKey, contactSim, edgeIndex, index2, shapeA, shapeB, world, v1, v3, v5
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	contactKey = (*b2Body)(unsafe.Pointer(body)).HeadContactKey
	index2 = 0
	for contactKey != -int32(1) && index2 < capacity {
		contactId = contactKey >> int32(1)
		edgeIndex = contactKey & int32(1)
		v1 = (*b2ContactArray)(unsafe.Pointer(world+1144)).Data + uintptr(contactId)*68
		goto _2
	_2:
		contact = v1
		// Is contact touching?
		if (*b2Contact)(unsafe.Pointer(contact)).Flags&uint32(b2_contactTouchingFlag) != 0 {
			v3 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr((*b2Contact)(unsafe.Pointer(contact)).ShapeIdA)*288
			goto _4
		_4:
			shapeA = v3
			v5 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr((*b2Contact)(unsafe.Pointer(contact)).ShapeIdB)*288
			goto _6
		_6:
			shapeB = v5
			(*(*ContactData)(unsafe.Pointer(contactData + uintptr(index2)*128))).ShapeIdA = ShapeId{
				Index1:     (*b2Shape)(unsafe.Pointer(shapeA)).Id + int32(1),
				World0:     bodyId.World0,
				Generation: (*b2Shape)(unsafe.Pointer(shapeA)).Generation,
			}
			(*(*ContactData)(unsafe.Pointer(contactData + uintptr(index2)*128))).ShapeIdB = ShapeId{
				Index1:     (*b2Shape)(unsafe.Pointer(shapeB)).Id + int32(1),
				World0:     bodyId.World0,
				Generation: (*b2Shape)(unsafe.Pointer(shapeB)).Generation,
			}
			contactSim = b2GetContactSim(tls, world, contact)
			(*(*ContactData)(unsafe.Pointer(contactData + uintptr(index2)*128))).Manifold = (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold
			index2 += int32(1)
		}
		contactKey = (*(*b2ContactEdge)(unsafe.Pointer(contact + 12 + uintptr(edgeIndex)*12))).NextKey
	}
	return index2
}

func b2Body_ComputeAABB(tls *_Stack, bodyId BodyId) (r AABB) {
	var aabb, c, v27, v5, v6 AABB
	var body, shape, world, v1, v3 uintptr
	var transform Transform
	var v11, v12, v13, v14, v16, v17, v18, v19, v21, v22, v23, v24, v26, v7, v8, v9 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, body, c, shape, transform, world, v1, v11, v12, v13, v14, v16, v17, v18, v19, v21, v22, v23, v24, v26, v27, v3, v5, v6, v7, v8, v9
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return AABB{}
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).HeadShapeId == -int32(1) {
		transform = b2GetBodyTransform(tls, world, (*b2Body)(unsafe.Pointer(body)).Id)
		return AABB{
			LowerBound: transform.P,
			UpperBound: transform.P,
		}
	}
	v1 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr((*b2Body)(unsafe.Pointer(body)).HeadShapeId)*288
	goto _2
_2:
	shape = v1
	aabb = (*b2Shape)(unsafe.Pointer(shape)).Aabb
	for (*b2Shape)(unsafe.Pointer(shape)).NextShapeId != -int32(1) {
		v3 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr((*b2Shape)(unsafe.Pointer(shape)).NextShapeId)*288
		goto _4
	_4:
		shape = v3
		v5 = aabb
		v6 = (*b2Shape)(unsafe.Pointer(shape)).Aabb
		v7 = v5.LowerBound.X
		v8 = v6.LowerBound.X
		if v7 < v8 {
			v11 = v7
		} else {
			v11 = v8
		}
		v9 = v11
		goto _10
	_10:
		c.LowerBound.X = v9
		v12 = v5.LowerBound.Y
		v13 = v6.LowerBound.Y
		if v12 < v13 {
			v16 = v12
		} else {
			v16 = v13
		}
		v14 = v16
		goto _15
	_15:
		c.LowerBound.Y = v14
		v17 = v5.UpperBound.X
		v18 = v6.UpperBound.X
		if v17 > v18 {
			v21 = v17
		} else {
			v21 = v18
		}
		v19 = v21
		goto _20
	_20:
		c.UpperBound.X = v19
		v22 = v5.UpperBound.Y
		v23 = v6.UpperBound.Y
		if v22 > v23 {
			v26 = v22
		} else {
			v26 = v23
		}
		v24 = v26
		goto _25
	_25:
		c.UpperBound.Y = v24
		v27 = c
		goto _28
	_28:
		aabb = v27
	}
	return aabb
}

func b2UpdateBodyMassData(tls *_Stack, world uintptr, body uintptr) {
	var bodySim, s3, s4, s5, state, v1, v13, v44 uintptr
	var deltaLinear, localCenter, oldCenter, v15, v17, v18, v21, v22, v24, v25, v29, v30, v32, v33, v34, v37, v38, v40, v41, v42 Vec2
	var extent, extent1 b2ShapeExtent
	var massData MassData
	var shapeId, shapeId1 int32
	var x, y, v10, v12, v16, v20, v26, v3, v36, v4, v46, v47, v48, v5, v50, v51, v52, v53, v55, v7, v8, v9 float32
	var v28 Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodySim, deltaLinear, extent, extent1, localCenter, massData, oldCenter, s3, s4, s5, shapeId, shapeId1, state, x, y, v1, v10, v12, v13, v15, v16, v17, v18, v20, v21, v22, v24, v25, v26, v28, v29, v3, v30, v32, v33, v34, v36, v37, v38, v4, v40, v41, v42, v44, v46, v47, v48, v5, v50, v51, v52, v53, v55, v7, v8, v9
	bodySim = b2GetBodySim(tls, world, body)
	// Compute mass data from shapes. Each shape has its own density.
	(*b2Body)(unsafe.Pointer(body)).Mass = float32FromFloat32(0)
	(*b2Body)(unsafe.Pointer(body)).Inertia = float32FromFloat32(0)
	(*b2BodySim)(unsafe.Pointer(bodySim)).InvMass = float32FromFloat32(0)
	(*b2BodySim)(unsafe.Pointer(bodySim)).InvInertia = float32FromFloat32(0)
	(*b2BodySim)(unsafe.Pointer(bodySim)).LocalCenter = b2Vec2_zero
	(*b2BodySim)(unsafe.Pointer(bodySim)).MinExtent = float32(float32FromFloat32(100000) * b2_lengthUnitsPerMeter)
	(*b2BodySim)(unsafe.Pointer(bodySim)).MaxExtent = float32FromFloat32(0)
	// Static and kinematic sims have zero mass.
	if (*b2Body)(unsafe.Pointer(body)).Type1 != int32(b2_dynamicBody) {
		(*b2BodySim)(unsafe.Pointer(bodySim)).Center = (*b2BodySim)(unsafe.Pointer(bodySim)).Transform.P
		// Need extents for kinematic bodies for sleeping to work correctly.
		if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_kinematicBody) {
			shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
			for shapeId != -int32(1) {
				v1 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr(shapeId)*288
				goto _2
			_2:
				s3 = v1
				extent = b2ComputeShapeExtent(tls, s3, b2Vec2_zero)
				v3 = (*b2BodySim)(unsafe.Pointer(bodySim)).MinExtent
				v4 = extent.MinExtent
				if v3 < v4 {
					v7 = v3
				} else {
					v7 = v4
				}
				v5 = v7
				goto _6
			_6:
				(*b2BodySim)(unsafe.Pointer(bodySim)).MinExtent = v5
				v8 = (*b2BodySim)(unsafe.Pointer(bodySim)).MaxExtent
				v9 = extent.MaxExtent
				if v8 > v9 {
					v12 = v8
				} else {
					v12 = v9
				}
				v10 = v12
				goto _11
			_11:
				(*b2BodySim)(unsafe.Pointer(bodySim)).MaxExtent = v10
				shapeId = (*b2Shape)(unsafe.Pointer(s3)).NextShapeId
			}
		}
		return
	}
	// Accumulate mass over all shapes.
	localCenter = b2Vec2_zero
	shapeId1 = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId1 != -int32(1) {
		v13 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr(shapeId1)*288
		goto _14
	_14:
		s4 = v13
		shapeId1 = (*b2Shape)(unsafe.Pointer(s4)).NextShapeId
		if (*b2Shape)(unsafe.Pointer(s4)).Density == float32FromFloat32(0) {
			continue
		}
		massData = b2ComputeShapeMass(tls, s4)
		*(*float32)(unsafe.Pointer(body + 88)) += massData.Mass
		v15 = localCenter
		v16 = massData.Mass
		v17 = massData.Center
		v18 = Vec2{
			X: v15.X + float32(v16*v17.X),
			Y: v15.Y + float32(v16*v17.Y),
		}
		goto _19
	_19:
		localCenter = v18
		*(*float32)(unsafe.Pointer(body + 92)) += massData.RotationalInertia
	}
	// Compute center of mass.
	if (*b2Body)(unsafe.Pointer(body)).Mass > float32FromFloat32(0) {
		(*b2BodySim)(unsafe.Pointer(bodySim)).InvMass = float32FromFloat32(1) / (*b2Body)(unsafe.Pointer(body)).Mass
		v20 = (*b2BodySim)(unsafe.Pointer(bodySim)).InvMass
		v21 = localCenter
		v22 = Vec2{
			X: float32(v20 * v21.X),
			Y: float32(v20 * v21.Y),
		}
		goto _23
	_23:
		localCenter = v22
	}
	if (*b2Body)(unsafe.Pointer(body)).Inertia > float32FromFloat32(0) && int32FromUint8((*b2Body)(unsafe.Pointer(body)).FixedRotation) == false1 {
		// Center the inertia about the center of mass.
		v24 = localCenter
		v25 = localCenter
		v26 = float32(v24.X*v25.X) + float32(v24.Y*v25.Y)
		goto _27
	_27:
		*(*float32)(unsafe.Pointer(body + 92)) -= float32((*b2Body)(unsafe.Pointer(body)).Mass * v26)
		(*b2BodySim)(unsafe.Pointer(bodySim)).InvInertia = float32FromFloat32(1) / (*b2Body)(unsafe.Pointer(body)).Inertia
	} else {
		(*b2Body)(unsafe.Pointer(body)).Inertia = float32FromFloat32(0)
		(*b2BodySim)(unsafe.Pointer(bodySim)).InvInertia = float32FromFloat32(0)
	}
	// Move center of mass.
	oldCenter = (*b2BodySim)(unsafe.Pointer(bodySim)).Center
	(*b2BodySim)(unsafe.Pointer(bodySim)).LocalCenter = localCenter
	v28 = (*b2BodySim)(unsafe.Pointer(bodySim)).Transform
	v29 = (*b2BodySim)(unsafe.Pointer(bodySim)).LocalCenter
	x = float32(v28.Q.C*v29.X) - float32(v28.Q.S*v29.Y) + v28.P.X
	y = float32(v28.Q.S*v29.X) + float32(v28.Q.C*v29.Y) + v28.P.Y
	v30 = Vec2{
		X: x,
		Y: y,
	}
	goto _31
_31:
	(*b2BodySim)(unsafe.Pointer(bodySim)).Center = v30
	(*b2BodySim)(unsafe.Pointer(bodySim)).Center0 = (*b2BodySim)(unsafe.Pointer(bodySim)).Center
	// Update center of mass velocity
	state = b2GetBodyState(tls, world, body)
	if state != uintptrFromInt32(0) {
		v32 = (*b2BodySim)(unsafe.Pointer(bodySim)).Center
		v33 = oldCenter
		v34 = Vec2{
			X: v32.X - v33.X,
			Y: v32.Y - v33.Y,
		}
		goto _35
	_35:
		v36 = (*b2BodyState)(unsafe.Pointer(state)).AngularVelocity
		v37 = v34
		v38 = Vec2{
			X: float32(-v36 * v37.Y),
			Y: float32(v36 * v37.X),
		}
		goto _39
	_39:
		deltaLinear = v38
		v40 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
		v41 = deltaLinear
		v42 = Vec2{
			X: v40.X + v41.X,
			Y: v40.Y + v41.Y,
		}
		goto _43
	_43:
		(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity = v42
	}
	// Compute body extents relative to center of mass
	shapeId1 = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId1 != -int32(1) {
		v44 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr(shapeId1)*288
		goto _45
	_45:
		s5 = v44
		extent1 = b2ComputeShapeExtent(tls, s5, localCenter)
		v46 = (*b2BodySim)(unsafe.Pointer(bodySim)).MinExtent
		v47 = extent1.MinExtent
		if v46 < v47 {
			v50 = v46
		} else {
			v50 = v47
		}
		v48 = v50
		goto _49
	_49:
		(*b2BodySim)(unsafe.Pointer(bodySim)).MinExtent = v48
		v51 = (*b2BodySim)(unsafe.Pointer(bodySim)).MaxExtent
		v52 = extent1.MaxExtent
		if v51 > v52 {
			v55 = v51
		} else {
			v55 = v52
		}
		v53 = v55
		goto _54
	_54:
		(*b2BodySim)(unsafe.Pointer(bodySim)).MaxExtent = v53
		shapeId1 = (*b2Shape)(unsafe.Pointer(s5)).NextShapeId
	}
}

func b2Body_GetPosition(tls *_Stack, bodyId BodyId) (r Vec2) {
	var body, world uintptr
	var transform Transform
	_, _, _ = body, transform, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	transform = b2GetBodyTransformQuick(tls, world, body)
	return transform.P
}

func b2Body_GetRotation(tls *_Stack, bodyId BodyId) (r Rot) {
	var body, world uintptr
	var transform Transform
	_, _, _ = body, transform, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	transform = b2GetBodyTransformQuick(tls, world, body)
	return transform.Q
}

func b2Body_GetTransform(tls *_Stack, bodyId BodyId) (r Transform) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return b2GetBodyTransformQuick(tls, world, body)
}

func b2Body_GetLocalPoint(tls *_Stack, bodyId BodyId, worldPoint Vec2) (r Vec2) {
	var body, world uintptr
	var transform, v1 Transform
	var vx, vy float32
	var v2, v3 Vec2
	_, _, _, _, _, _, _, _ = body, transform, vx, vy, world, v1, v2, v3
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	transform = b2GetBodyTransformQuick(tls, world, body)
	v1 = transform
	v2 = worldPoint
	vx = v2.X - v1.P.X
	vy = v2.Y - v1.P.Y
	v3 = Vec2{
		X: float32(v1.Q.C*vx) + float32(v1.Q.S*vy),
		Y: float32(-v1.Q.S*vx) + float32(v1.Q.C*vy),
	}
	goto _4
_4:
	return v3
}

func b2Body_GetWorldPoint(tls *_Stack, bodyId BodyId, localPoint Vec2) (r Vec2) {
	var body, world uintptr
	var transform, v1 Transform
	var x, y float32
	var v2, v3 Vec2
	_, _, _, _, _, _, _, _ = body, transform, world, x, y, v1, v2, v3
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	transform = b2GetBodyTransformQuick(tls, world, body)
	v1 = transform
	v2 = localPoint
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	return v3
}

func b2Body_GetLocalVector(tls *_Stack, bodyId BodyId, worldVector Vec2) (r Vec2) {
	var body, world uintptr
	var transform Transform
	var v1 Rot
	var v2, v3 Vec2
	_, _, _, _, _, _ = body, transform, world, v1, v2, v3
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	transform = b2GetBodyTransformQuick(tls, world, body)
	v1 = transform.Q
	v2 = worldVector
	v3 = Vec2{
		X: float32(v1.C*v2.X) + float32(v1.S*v2.Y),
		Y: float32(-v1.S*v2.X) + float32(v1.C*v2.Y),
	}
	goto _4
_4:
	return v3
}

func b2Body_GetWorldVector(tls *_Stack, bodyId BodyId, localVector Vec2) (r Vec2) {
	var body, world uintptr
	var transform Transform
	var v1 Rot
	var v2, v3 Vec2
	_, _, _, _, _, _ = body, transform, world, v1, v2, v3
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	transform = b2GetBodyTransformQuick(tls, world, body)
	v1 = transform.Q
	v2 = localVector
	v3 = Vec2{
		X: float32(v1.C*v2.X) - float32(v1.S*v2.Y),
		Y: float32(v1.S*v2.X) + float32(v1.C*v2.Y),
	}
	goto _4
_4:
	return v3
}

func b2Body_SetTransform(tls *_Stack, bodyId BodyId, position Vec2, rotation Rot) {
	var aabb, fatAABB, v7, v8 AABB
	var body, bodySim, broadPhase, shape, world, v5 uintptr
	var margin, speculativeDistance, x, y float32
	var s, v9 uint8
	var shapeId int32
	var transform, v1 Transform
	var v2, v3 Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, body, bodySim, broadPhase, fatAABB, margin, s, shape, shapeId, speculativeDistance, transform, world, x, y, v1, v2, v3, v5, v7, v8, v9
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	(*b2BodySim)(unsafe.Pointer(bodySim)).Transform.P = position
	(*b2BodySim)(unsafe.Pointer(bodySim)).Transform.Q = rotation
	v1 = (*b2BodySim)(unsafe.Pointer(bodySim)).Transform
	v2 = (*b2BodySim)(unsafe.Pointer(bodySim)).LocalCenter
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	(*b2BodySim)(unsafe.Pointer(bodySim)).Center = v3
	(*b2BodySim)(unsafe.Pointer(bodySim)).Rotation0 = (*b2BodySim)(unsafe.Pointer(bodySim)).Transform.Q
	(*b2BodySim)(unsafe.Pointer(bodySim)).Center0 = (*b2BodySim)(unsafe.Pointer(bodySim)).Center
	broadPhase = world + 40
	transform = (*b2BodySim)(unsafe.Pointer(bodySim)).Transform
	margin = float32(float32FromFloat32(0.05) * b2_lengthUnitsPerMeter)
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		v5 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr(shapeId)*288
		goto _6
	_6:
		shape = v5
		aabb = b2ComputeShapeAABB(tls, shape, transform)
		aabb.LowerBound.X -= speculativeDistance
		aabb.LowerBound.Y -= speculativeDistance
		aabb.UpperBound.X += speculativeDistance
		aabb.UpperBound.Y += speculativeDistance
		(*b2Shape)(unsafe.Pointer(shape)).Aabb = aabb
		v7 = (*b2Shape)(unsafe.Pointer(shape)).FatAABB
		v8 = aabb
		s = uint8(true1)
		s = boolUint8(s != 0 && v7.LowerBound.X <= v8.LowerBound.X)
		s = boolUint8(s != 0 && v7.LowerBound.Y <= v8.LowerBound.Y)
		s = boolUint8(s != 0 && v8.UpperBound.X <= v7.UpperBound.X)
		s = boolUint8(s != 0 && v8.UpperBound.Y <= v7.UpperBound.Y)
		v9 = s
		goto _10
	_10:
		if int32FromUint8(v9) == false1 {
			fatAABB.LowerBound.X = aabb.LowerBound.X - margin
			fatAABB.LowerBound.Y = aabb.LowerBound.Y - margin
			fatAABB.UpperBound.X = aabb.UpperBound.X + margin
			fatAABB.UpperBound.Y = aabb.UpperBound.Y + margin
			(*b2Shape)(unsafe.Pointer(shape)).FatAABB = fatAABB
			// They body could be disabled
			if (*b2Shape)(unsafe.Pointer(shape)).ProxyKey != -int32(1) {
				b2BroadPhase_MoveProxy(tls, broadPhase, (*b2Shape)(unsafe.Pointer(shape)).ProxyKey, fatAABB)
			}
		}
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
	}
}

func b2Body_GetLinearVelocity(tls *_Stack, bodyId BodyId) (r Vec2) {
	var body, state, world uintptr
	_, _, _ = body, state, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	state = b2GetBodyState(tls, world, body)
	if state != uintptrFromInt32(0) {
		return (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
	}
	return b2Vec2_zero
}

func b2Body_GetAngularVelocity(tls *_Stack, bodyId BodyId) (r float32) {
	var body, state, world uintptr
	_, _, _ = body, state, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	state = b2GetBodyState(tls, world, body)
	if state != uintptrFromInt32(0) {
		return (*b2BodyState)(unsafe.Pointer(state)).AngularVelocity
	}
	return float32(0)
}

func b2Body_SetLinearVelocity(tls *_Stack, bodyId BodyId, linearVelocity Vec2) {
	var body, state, world uintptr
	var v1 Vec2
	var v2 float32
	_, _, _, _, _ = body, state, world, v1, v2
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_staticBody) {
		return
	}
	v1 = linearVelocity
	v2 = float32(v1.X*v1.X) + float32(v1.Y*v1.Y)
	goto _3
_3:
	if v2 > float32FromFloat32(0) {
		b2WakeBody(tls, world, body)
	}
	state = b2GetBodyState(tls, world, body)
	if state == uintptrFromInt32(0) {
		return
	}
	(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity = linearVelocity
}

func b2Body_SetAngularVelocity(tls *_Stack, bodyId BodyId, angularVelocity float32) {
	var body, state, world uintptr
	_, _, _ = body, state, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_staticBody) || (*b2Body)(unsafe.Pointer(body)).FixedRotation != 0 {
		return
	}
	if angularVelocity != float32FromFloat32(0) {
		b2WakeBody(tls, world, body)
	}
	state = b2GetBodyState(tls, world, body)
	if state == uintptrFromInt32(0) {
		return
	}
	(*b2BodyState)(unsafe.Pointer(state)).AngularVelocity = angularVelocity
}

func b2Body_SetTargetTransform(tls *_Stack, bodyId BodyId, target Transform, timeStep float32) {
	var angularVelocity, c, deltaAngle, invTimeStep, maxVelocity, s1, x, y, v15, v18, v20, v21, v23, v9 float32
	var body, sim, state, world uintptr
	var center1, center2, linearVelocity, v10, v11, v17, v2, v3, v5, v6, v7 Vec2
	var q1, q2, v13, v14 Rot
	var v1 Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angularVelocity, body, c, center1, center2, deltaAngle, invTimeStep, linearVelocity, maxVelocity, q1, q2, s1, sim, state, world, x, y, v1, v10, v11, v13, v14, v15, v17, v18, v2, v20, v21, v23, v3, v5, v6, v7, v9
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_staticBody) || timeStep <= float32FromFloat32(0) {
		return
	}
	sim = b2GetBodySim(tls, world, body)
	// Compute linear velocity
	center1 = (*b2BodySim)(unsafe.Pointer(sim)).Center
	v1 = target
	v2 = (*b2BodySim)(unsafe.Pointer(sim)).LocalCenter
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	center2 = v3
	invTimeStep = float32FromFloat32(1) / timeStep
	v5 = center2
	v6 = center1
	v7 = Vec2{
		X: v5.X - v6.X,
		Y: v5.Y - v6.Y,
	}
	goto _8
_8:
	v9 = invTimeStep
	v10 = v7
	v11 = Vec2{
		X: float32(v9 * v10.X),
		Y: float32(v9 * v10.Y),
	}
	goto _12
_12:
	linearVelocity = v11
	// Compute angular velocity
	angularVelocity = float32FromFloat32(0)
	if int32FromUint8((*b2Body)(unsafe.Pointer(body)).FixedRotation) == false1 {
		q1 = (*b2BodySim)(unsafe.Pointer(sim)).Transform.Q
		q2 = target.Q
		v13 = q2
		v14 = q1
		s1 = float32(v13.S*v14.C) - float32(v13.C*v14.S)
		c = float32(v13.C*v14.C) + float32(v13.S*v14.S)
		v15 = b2Atan2(tls, s1, c)
		goto _16
	_16:
		deltaAngle = v15
		angularVelocity = float32(invTimeStep * deltaAngle)
	}
	v17 = linearVelocity
	v18 = sqrtf(tls, float32(v17.X*v17.X)+float32(v17.Y*v17.Y))
	goto _19
_19:
	v20 = angularVelocity
	if v20 < float32FromInt32(0) {
		v23 = -v20
	} else {
		v23 = v20
	}
	v21 = v23
	goto _22
_22:
	maxVelocity = v18 + float32(v21*(*b2BodySim)(unsafe.Pointer(sim)).MaxExtent)
	// Return if velocity would be sleepy
	if maxVelocity < (*b2Body)(unsafe.Pointer(body)).SleepThreshold {
		return
	}
	// Must wake for state to exist
	b2WakeBody(tls, world, body)
	state = b2GetBodyState(tls, world, body)
	if state == uintptrFromInt32(0) {
		return
	}
	(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity = linearVelocity
	(*b2BodyState)(unsafe.Pointer(state)).AngularVelocity = angularVelocity
}

func b2Body_GetLocalPointVelocity(tls *_Stack, bodyId BodyId, localPoint Vec2) (r1 Vec2) {
	var body, bodySim, set, state, world, v1, v3 uintptr
	var r, v2, v10, v11, v14, v15, v17, v18, v19, v5, v6, v7 Vec2
	var v13 float32
	var v9 Rot
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, r, set, state, v2, world, v1, v10, v11, v13, v14, v15, v17, v18, v19, v3, v5, v6, v7, v9
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	state = b2GetBodyState(tls, world, body)
	if state == uintptrFromInt32(0) {
		return b2Vec2_zero
	}
	v1 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr((*b2Body)(unsafe.Pointer(body)).SetIndex)*88
	goto _2
_2:
	set = v1
	v3 = (*b2BodySimArray)(unsafe.Pointer(set)).Data + uintptr((*b2Body)(unsafe.Pointer(body)).LocalIndex)*100
	goto _4
_4:
	bodySim = v3
	v5 = localPoint
	v6 = (*b2BodySim)(unsafe.Pointer(bodySim)).LocalCenter
	v7 = Vec2{
		X: v5.X - v6.X,
		Y: v5.Y - v6.Y,
	}
	goto _8
_8:
	v9 = (*b2BodySim)(unsafe.Pointer(bodySim)).Transform.Q
	v10 = v7
	v11 = Vec2{
		X: float32(v9.C*v10.X) - float32(v9.S*v10.Y),
		Y: float32(v9.S*v10.X) + float32(v9.C*v10.Y),
	}
	goto _12
_12:
	r = v11
	v13 = (*b2BodyState)(unsafe.Pointer(state)).AngularVelocity
	v14 = r
	v15 = Vec2{
		X: float32(-v13 * v14.Y),
		Y: float32(v13 * v14.X),
	}
	goto _16
_16:
	v17 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
	v18 = v15
	v19 = Vec2{
		X: v17.X + v18.X,
		Y: v17.Y + v18.Y,
	}
	goto _20
_20:
	v2 = v19
	return v2
}

func b2Body_GetWorldPointVelocity(tls *_Stack, bodyId BodyId, worldPoint Vec2) (r1 Vec2) {
	var body, bodySim, set, state, world, v11, v3 uintptr
	var r, v1, v10, v111, v13, v14, v15, v5, v6, v7 Vec2
	var v9 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, r, set, state, v1, world, v11, v10, v111, v13, v14, v15, v3, v5, v6, v7, v9
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	state = b2GetBodyState(tls, world, body)
	if state == uintptrFromInt32(0) {
		return b2Vec2_zero
	}
	v11 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr((*b2Body)(unsafe.Pointer(body)).SetIndex)*88
	goto _2
_2:
	set = v11
	v3 = (*b2BodySimArray)(unsafe.Pointer(set)).Data + uintptr((*b2Body)(unsafe.Pointer(body)).LocalIndex)*100
	goto _4
_4:
	bodySim = v3
	v5 = worldPoint
	v6 = (*b2BodySim)(unsafe.Pointer(bodySim)).Center
	v7 = Vec2{
		X: v5.X - v6.X,
		Y: v5.Y - v6.Y,
	}
	goto _8
_8:
	r = v7
	v9 = (*b2BodyState)(unsafe.Pointer(state)).AngularVelocity
	v10 = r
	v111 = Vec2{
		X: float32(-v9 * v10.Y),
		Y: float32(v9 * v10.X),
	}
	goto _12
_12:
	v13 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
	v14 = v111
	v15 = Vec2{
		X: v13.X + v14.X,
		Y: v13.Y + v14.Y,
	}
	goto _16
_16:
	v1 = v15
	return v1
}

func b2Body_ApplyForce(tls *_Stack, bodyId BodyId, force Vec2, point Vec2, wake uint8) {
	var body, bodySim, world uintptr
	var v1, v10, v2, v3, v5, v6, v7, v9 Vec2
	var v11 float32
	_, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, world, v1, v10, v11, v2, v3, v5, v6, v7, v9
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	if wake != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeBody(tls, world, body)
	}
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		bodySim = b2GetBodySim(tls, world, body)
		v1 = (*b2BodySim)(unsafe.Pointer(bodySim)).Force
		v2 = force
		v3 = Vec2{
			X: v1.X + v2.X,
			Y: v1.Y + v2.Y,
		}
		goto _4
	_4:
		(*b2BodySim)(unsafe.Pointer(bodySim)).Force = v3
		v5 = point
		v6 = (*b2BodySim)(unsafe.Pointer(bodySim)).Center
		v7 = Vec2{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
		}
		goto _8
	_8:
		v9 = v7
		v10 = force
		v11 = float32(v9.X*v10.Y) - float32(v9.Y*v10.X)
		goto _12
	_12:
		*(*float32)(unsafe.Pointer(bodySim + 56)) += v11
	}
}

func b2Body_ApplyForceToCenter(tls *_Stack, bodyId BodyId, force Vec2, wake uint8) {
	var body, bodySim, world uintptr
	var v1, v2, v3 Vec2
	_, _, _, _, _, _ = body, bodySim, world, v1, v2, v3
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	if wake != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeBody(tls, world, body)
	}
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		bodySim = b2GetBodySim(tls, world, body)
		v1 = (*b2BodySim)(unsafe.Pointer(bodySim)).Force
		v2 = force
		v3 = Vec2{
			X: v1.X + v2.X,
			Y: v1.Y + v2.Y,
		}
		goto _4
	_4:
		(*b2BodySim)(unsafe.Pointer(bodySim)).Force = v3
	}
}

func b2Body_ApplyTorque(tls *_Stack, bodyId BodyId, torque float32, wake uint8) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	if wake != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeBody(tls, world, body)
	}
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		bodySim = b2GetBodySim(tls, world, body)
		*(*float32)(unsafe.Pointer(bodySim + 56)) += torque
	}
}

func b2Body_ApplyLinearImpulse(tls *_Stack, bodyId BodyId, impulse Vec2, point Vec2, wake uint8) {
	var body, bodySim, set, state, world, v1, v3, v5 uintptr
	var localIndex int32
	var v10, v12, v13, v14, v16, v17, v7, v9 Vec2
	var v18, v8 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, localIndex, set, state, world, v1, v10, v12, v13, v14, v16, v17, v18, v3, v5, v7, v8, v9
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	if wake != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeBody(tls, world, body)
	}
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		localIndex = (*b2Body)(unsafe.Pointer(body)).LocalIndex
		v1 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr(int32(b2_awakeSet))*88
		goto _2
	_2:
		set = v1
		v3 = (*b2BodyStateArray)(unsafe.Pointer(set+16)).Data + uintptr(localIndex)*32
		goto _4
	_4:
		state = v3
		v5 = (*b2BodySimArray)(unsafe.Pointer(set)).Data + uintptr(localIndex)*100
		goto _6
	_6:
		bodySim = v5
		v7 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
		v8 = (*b2BodySim)(unsafe.Pointer(bodySim)).InvMass
		v9 = impulse
		v10 = Vec2{
			X: v7.X + float32(v8*v9.X),
			Y: v7.Y + float32(v8*v9.Y),
		}
		goto _11
	_11:
		(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity = v10
		v12 = point
		v13 = (*b2BodySim)(unsafe.Pointer(bodySim)).Center
		v14 = Vec2{
			X: v12.X - v13.X,
			Y: v12.Y - v13.Y,
		}
		goto _15
	_15:
		v16 = v14
		v17 = impulse
		v18 = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
		goto _19
	_19:
		*(*float32)(unsafe.Pointer(state + 8)) += float32((*b2BodySim)(unsafe.Pointer(bodySim)).InvInertia * v18)
		b2LimitVelocity(tls, state, (*b2World)(unsafe.Pointer(world)).MaxLinearSpeed)
	}
}

func b2Body_ApplyLinearImpulseToCenter(tls *_Stack, bodyId BodyId, impulse Vec2, wake uint8) {
	var body, bodySim, set, state, world, v1, v3, v5 uintptr
	var localIndex int32
	var v10, v7, v9 Vec2
	var v8 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, localIndex, set, state, world, v1, v10, v3, v5, v7, v8, v9
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	if wake != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeBody(tls, world, body)
	}
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		localIndex = (*b2Body)(unsafe.Pointer(body)).LocalIndex
		v1 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr(int32(b2_awakeSet))*88
		goto _2
	_2:
		set = v1
		v3 = (*b2BodyStateArray)(unsafe.Pointer(set+16)).Data + uintptr(localIndex)*32
		goto _4
	_4:
		state = v3
		v5 = (*b2BodySimArray)(unsafe.Pointer(set)).Data + uintptr(localIndex)*100
		goto _6
	_6:
		bodySim = v5
		v7 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
		v8 = (*b2BodySim)(unsafe.Pointer(bodySim)).InvMass
		v9 = impulse
		v10 = Vec2{
			X: v7.X + float32(v8*v9.X),
			Y: v7.Y + float32(v8*v9.Y),
		}
		goto _11
	_11:
		(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity = v10
		b2LimitVelocity(tls, state, (*b2World)(unsafe.Pointer(world)).MaxLinearSpeed)
	}
}

func b2Body_ApplyAngularImpulse(tls *_Stack, bodyId BodyId, impulse float32, wake uint8) {
	var body, bodySim, set, state, world, v1, v3, v5, v7 uintptr
	var id, localIndex int32
	_, _, _, _, _, _, _, _, _, _, _ = body, bodySim, id, localIndex, set, state, world, v1, v3, v5, v7
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	id = bodyId.Index1 - int32(1)
	v1 = (*b2BodyArray)(unsafe.Pointer(world+1024)).Data + uintptr(id)*128
	goto _2
_2:
	body = v1
	if wake != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		// this will not invalidate body pointer
		b2WakeBody(tls, world, body)
	}
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		localIndex = (*b2Body)(unsafe.Pointer(body)).LocalIndex
		v3 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr(int32(b2_awakeSet))*88
		goto _4
	_4:
		set = v3
		v5 = (*b2BodyStateArray)(unsafe.Pointer(set+16)).Data + uintptr(localIndex)*32
		goto _6
	_6:
		state = v5
		v7 = (*b2BodySimArray)(unsafe.Pointer(set)).Data + uintptr(localIndex)*100
		goto _8
	_8:
		bodySim = v7
		*(*float32)(unsafe.Pointer(state + 8)) += float32((*b2BodySim)(unsafe.Pointer(bodySim)).InvInertia * impulse)
	}
}

func b2Body_GetType(tls *_Stack, bodyId BodyId) (r BodyType) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return (*b2Body)(unsafe.Pointer(body)).Type1
}

// C documentation
//
//	// Changing the body type is quite complex mainly due to joints.
//	// Considerations:
//	// - body and joints must be moved to the correct set
//	// - islands must be updated
//	// - graph coloring must be correct
//	// - any body connected to a joint may be disabled
//	// - joints between static bodies must go into the static set
func b2Body_SetType(tls *_Stack, bodyId BodyId, type1 BodyType) {
	var awakeSet, awakeSet1, body, bodyA, bodyB, bodySim, joint, joint1, joint2, joint3, otherBody, otherBody1, shape, shape1, shape2, staticSet, staticSet1, world, v1, v11, v13, v15, v17, v19, v21, v23, v25, v27, v29, v3, v31, v5, v7, v9 uintptr
	var edgeIndex, edgeIndex1, edgeIndex2, edgeIndex3, jointId, jointId1, jointId2, jointId3, jointKey, jointKey1, jointKey2, jointKey3, otherBodyId, otherEdgeIndex, otherEdgeIndex1, shapeId, shapeId1, shapeId2 int32
	var forcePairCreation, forcePairCreation1, forcePairCreation2, wakeBodies uint8
	var originalType, proxyType, proxyType1 b2BodyType1
	var transform, transform1, transform2 Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = awakeSet, awakeSet1, body, bodyA, bodyB, bodySim, edgeIndex, edgeIndex1, edgeIndex2, edgeIndex3, forcePairCreation, forcePairCreation1, forcePairCreation2, joint, joint1, joint2, joint3, jointId, jointId1, jointId2, jointId3, jointKey, jointKey1, jointKey2, jointKey3, originalType, otherBody, otherBody1, otherBodyId, otherEdgeIndex, otherEdgeIndex1, proxyType, proxyType1, shape, shape1, shape2, shapeId, shapeId1, shapeId2, staticSet, staticSet1, transform, transform1, transform2, wakeBodies, world, v1, v11, v13, v15, v17, v19, v21, v23, v25, v27, v29, v3, v31, v5, v7, v9
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	originalType = (*b2Body)(unsafe.Pointer(body)).Type1
	if originalType == type1 {
		return
	}
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_disabledSet) {
		// Disabled bodies don't change solver sets or islands when they change type.
		(*b2Body)(unsafe.Pointer(body)).Type1 = type1
		// Body type affects the mass
		b2UpdateBodyMassData(tls, world, body)
		return
	}
	// Destroy all contacts but don't wake bodies.
	wakeBodies = uint8(false1)
	b2DestroyBodyContacts(tls, world, body, wakeBodies)
	// Wake this body because we assume below that it is awake or static.
	b2WakeBody(tls, world, body)
	// Unlink all joints and wake attached bodies.
	jointKey = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
	for jointKey != -int32(1) {
		jointId = jointKey >> int32(1)
		edgeIndex = jointKey & int32(1)
		v1 = (*b2JointArray)(unsafe.Pointer(world+1104)).Data + uintptr(jointId)*72
		goto _2
	_2:
		joint = v1
		if (*b2Joint)(unsafe.Pointer(joint)).IslandId != -int32(1) {
			b2UnlinkJoint(tls, world, joint)
		}
		v3 = (*b2BodyArray)(unsafe.Pointer(world+1024)).Data + uintptr((*(*b2JointEdge)(unsafe.Pointer(joint + 20))).BodyId)*128
		goto _4
	_4:
		// A body going from static to dynamic or kinematic goes to the awake set
		// and other attached bodies must be awake as well. For consistency, this is
		// done for all cases.
		bodyA = v3
		v5 = (*b2BodyArray)(unsafe.Pointer(world+1024)).Data + uintptr((*(*b2JointEdge)(unsafe.Pointer(joint + 20 + 1*12))).BodyId)*128
		goto _6
	_6:
		bodyB = v5
		b2WakeBody(tls, world, bodyA)
		b2WakeBody(tls, world, bodyB)
		jointKey = (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + uintptr(edgeIndex)*12))).NextKey
	}
	(*b2Body)(unsafe.Pointer(body)).Type1 = type1
	if originalType == int32(b2_staticBody) {
		// Body is going from static to dynamic or kinematic. It only makes sense to move it to the awake set.
		v7 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr(int32(b2_staticSet))*88
		goto _8
	_8:
		staticSet = v7
		v9 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr(int32(b2_awakeSet))*88
		goto _10
	_10:
		awakeSet = v9
		// Transfer body to awake set
		b2TransferBody(tls, world, awakeSet, staticSet, body)
		// Create island for body
		b2CreateIslandForBody(tls, world, int32(b2_awakeSet), body)
		// Transfer static joints to awake set
		jointKey1 = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
		for jointKey1 != -int32(1) {
			jointId1 = jointKey1 >> int32(1)
			edgeIndex1 = jointKey1 & int32(1)
			v11 = (*b2JointArray)(unsafe.Pointer(world+1104)).Data + uintptr(jointId1)*72
			goto _12
		_12:
			joint1 = v11
			// Transfer the joint if it is in the static set
			if (*b2Joint)(unsafe.Pointer(joint1)).SetIndex == int32(b2_staticSet) {
				b2TransferJoint(tls, world, awakeSet, staticSet, joint1)
			} else {
				if (*b2Joint)(unsafe.Pointer(joint1)).SetIndex == int32(b2_awakeSet) {
					// In this case the joint must be re-inserted into the constraint graph to ensure the correct
					// graph color.
					// First transfer to the static set.
					b2TransferJoint(tls, world, staticSet, awakeSet, joint1)
					// Now transfer it back to the awake set and into the graph coloring.
					b2TransferJoint(tls, world, awakeSet, staticSet, joint1)
				} else {
					// Otherwise the joint must be disabled.
				}
			}
			jointKey1 = (*(*b2JointEdge)(unsafe.Pointer(joint1 + 20 + uintptr(edgeIndex1)*12))).NextKey
		}
		// Recreate shape proxies in movable tree.
		transform = b2GetBodyTransformQuick(tls, world, body)
		shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
		for shapeId != -int32(1) {
			v13 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr(shapeId)*288
			goto _14
		_14:
			shape = v13
			shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
			b2DestroyShapeProxy(tls, shape, world+40)
			forcePairCreation = uint8(true1)
			proxyType = type1
			b2CreateShapeProxy(tls, shape, world+40, proxyType, transform, forcePairCreation)
		}
	} else {
		if type1 == int32(b2_staticBody) {
			// The body is going from dynamic/kinematic to static. It should be awake.
			v15 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr(int32(b2_staticSet))*88
			goto _16
		_16:
			staticSet1 = v15
			v17 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr(int32(b2_awakeSet))*88
			goto _18
		_18:
			awakeSet1 = v17
			// Transfer body to static set
			b2TransferBody(tls, world, staticSet1, awakeSet1, body)
			// Remove body from island.
			b2RemoveBodyFromIsland(tls, world, body)
			v19 = (*b2BodySimArray)(unsafe.Pointer(staticSet1)).Data + uintptr((*b2Body)(unsafe.Pointer(body)).LocalIndex)*100
			goto _20
		_20:
			bodySim = v19
			(*b2BodySim)(unsafe.Pointer(bodySim)).IsFast = uint8(false1)
			// Maybe transfer joints to static set.
			jointKey2 = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
			for jointKey2 != -int32(1) {
				jointId2 = jointKey2 >> int32(1)
				edgeIndex2 = jointKey2 & int32(1)
				v21 = (*b2JointArray)(unsafe.Pointer(world+1104)).Data + uintptr(jointId2)*72
				goto _22
			_22:
				joint2 = v21
				jointKey2 = (*(*b2JointEdge)(unsafe.Pointer(joint2 + 20 + uintptr(edgeIndex2)*12))).NextKey
				otherEdgeIndex = edgeIndex2 ^ int32(1)
				v23 = (*b2BodyArray)(unsafe.Pointer(world+1024)).Data + uintptr((*(*b2JointEdge)(unsafe.Pointer(joint2 + 20 + uintptr(otherEdgeIndex)*12))).BodyId)*128
				goto _24
			_24:
				otherBody = v23
				// Skip disabled joint
				if (*b2Joint)(unsafe.Pointer(joint2)).SetIndex == int32(b2_disabledSet) {
					// Joint is disable, should be connected to a disabled body
					continue
				}
				// Since the body was not static, the joint must be awake.
				// Only transfer joint to static set if both bodies are static.
				if (*b2Body)(unsafe.Pointer(otherBody)).SetIndex == int32(b2_staticSet) {
					b2TransferJoint(tls, world, staticSet1, awakeSet1, joint2)
				} else {
					// The other body must be awake.
					// The joint must live in a graph color.
					// In this case the joint must be re-inserted into the constraint graph to ensure the correct
					// graph color.
					// First transfer to the static set.
					b2TransferJoint(tls, world, staticSet1, awakeSet1, joint2)
					// Now transfer it back to the awake set and into the graph coloring.
					b2TransferJoint(tls, world, awakeSet1, staticSet1, joint2)
				}
			}
			// Recreate shape proxies in static tree.
			transform1 = b2GetBodyTransformQuick(tls, world, body)
			shapeId1 = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
			for shapeId1 != -int32(1) {
				v25 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr(shapeId1)*288
				goto _26
			_26:
				shape1 = v25
				shapeId1 = (*b2Shape)(unsafe.Pointer(shape1)).NextShapeId
				b2DestroyShapeProxy(tls, shape1, world+40)
				forcePairCreation1 = uint8(true1)
				b2CreateShapeProxy(tls, shape1, world+40, int32(b2_staticBody), transform1, forcePairCreation1)
			}
		} else {
			// Recreate shape proxies in static tree.
			transform2 = b2GetBodyTransformQuick(tls, world, body)
			shapeId2 = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
			for shapeId2 != -int32(1) {
				v27 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr(shapeId2)*288
				goto _28
			_28:
				shape2 = v27
				shapeId2 = (*b2Shape)(unsafe.Pointer(shape2)).NextShapeId
				b2DestroyShapeProxy(tls, shape2, world+40)
				proxyType1 = type1
				forcePairCreation2 = uint8(true1)
				b2CreateShapeProxy(tls, shape2, world+40, proxyType1, transform2, forcePairCreation2)
			}
		}
	}
	// Relink all joints
	jointKey3 = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
	for jointKey3 != -int32(1) {
		jointId3 = jointKey3 >> int32(1)
		edgeIndex3 = jointKey3 & int32(1)
		v29 = (*b2JointArray)(unsafe.Pointer(world+1104)).Data + uintptr(jointId3)*72
		goto _30
	_30:
		joint3 = v29
		jointKey3 = (*(*b2JointEdge)(unsafe.Pointer(joint3 + 20 + uintptr(edgeIndex3)*12))).NextKey
		otherEdgeIndex1 = edgeIndex3 ^ int32(1)
		otherBodyId = (*(*b2JointEdge)(unsafe.Pointer(joint3 + 20 + uintptr(otherEdgeIndex1)*12))).BodyId
		v31 = (*b2BodyArray)(unsafe.Pointer(world+1024)).Data + uintptr(otherBodyId)*128
		goto _32
	_32:
		otherBody1 = v31
		if (*b2Body)(unsafe.Pointer(otherBody1)).SetIndex == int32(b2_disabledSet) {
			continue
		}
		if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_staticBody) && (*b2Body)(unsafe.Pointer(otherBody1)).Type1 == int32(b2_staticBody) {
			continue
		}
		b2LinkJoint(tls, world, joint3, uint8(false1))
	}
	b2MergeAwakeIslands(tls, world)
	// Body type affects the mass
	b2UpdateBodyMassData(tls, world, body)
	b2ValidateSolverSets(tls, world)
}

func b2Body_SetName(tls *_Stack, bodyId BodyId, name uintptr) {
	var body, world uintptr
	var i int32
	_, _, _ = body, i, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	if name != uintptrFromInt32(0) {
		i = 0
		for {
			if !(i < int32(31)) {
				break
			}
			*(*uint8)(unsafe.Pointer(body + uintptr(i))) = *(*uint8)(unsafe.Pointer(name + uintptr(i)))
			goto _1
		_1:
			;
			i++
		}
		*(*uint8)(unsafe.Pointer(body + 31)) = uint8(0)
	} else {
		memset(tls, body, 0, uint64FromInt32(32)*uint64FromInt64(1))
	}
}

func b2Body_GetName(tls *_Stack, bodyId BodyId) (r uintptr) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return body
}

func b2Body_SetUserData(tls *_Stack, bodyId BodyId, userData uintptr) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	(*b2Body)(unsafe.Pointer(body)).UserData = userData
}

func b2Body_GetUserData(tls *_Stack, bodyId BodyId) (r uintptr) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return (*b2Body)(unsafe.Pointer(body)).UserData
}

func b2Body_GetMass(tls *_Stack, bodyId BodyId) (r float32) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return (*b2Body)(unsafe.Pointer(body)).Mass
}

func b2Body_GetRotationalInertia(tls *_Stack, bodyId BodyId) (r float32) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return (*b2Body)(unsafe.Pointer(body)).Inertia
}

func b2Body_GetLocalCenterOfMass(tls *_Stack, bodyId BodyId) (r Vec2) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	return (*b2BodySim)(unsafe.Pointer(bodySim)).LocalCenter
}

func b2Body_GetWorldCenterOfMass(tls *_Stack, bodyId BodyId) (r Vec2) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	return (*b2BodySim)(unsafe.Pointer(bodySim)).Center
}

func b2Body_SetMassData(tls *_Stack, bodyId BodyId, massData MassData) {
	var body, bodySim, world uintptr
	var center, v2, v3 Vec2
	var x, y, v5, v6 float32
	var v1 Transform
	_, _, _, _, _, _, _, _, _, _, _ = body, bodySim, center, world, x, y, v1, v2, v3, v5, v6
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	(*b2Body)(unsafe.Pointer(body)).Mass = massData.Mass
	(*b2Body)(unsafe.Pointer(body)).Inertia = massData.RotationalInertia
	(*b2BodySim)(unsafe.Pointer(bodySim)).LocalCenter = massData.Center
	v1 = (*b2BodySim)(unsafe.Pointer(bodySim)).Transform
	v2 = massData.Center
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	center = v3
	(*b2BodySim)(unsafe.Pointer(bodySim)).Center = center
	(*b2BodySim)(unsafe.Pointer(bodySim)).Center0 = center
	if (*b2Body)(unsafe.Pointer(body)).Mass > float32FromFloat32(0) {
		v5 = float32FromFloat32(1) / (*b2Body)(unsafe.Pointer(body)).Mass
	} else {
		v5 = float32FromFloat32(0)
	}
	(*b2BodySim)(unsafe.Pointer(bodySim)).InvMass = v5
	if (*b2Body)(unsafe.Pointer(body)).Inertia > float32FromFloat32(0) {
		v6 = float32FromFloat32(1) / (*b2Body)(unsafe.Pointer(body)).Inertia
	} else {
		v6 = float32FromFloat32(0)
	}
	(*b2BodySim)(unsafe.Pointer(bodySim)).InvInertia = v6
}

func b2Body_GetMassData(tls *_Stack, bodyId BodyId) (r MassData) {
	var body, bodySim, world uintptr
	var massData MassData
	_, _, _, _ = body, bodySim, massData, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	massData = MassData{
		Mass:              (*b2Body)(unsafe.Pointer(body)).Mass,
		Center:            (*b2BodySim)(unsafe.Pointer(bodySim)).LocalCenter,
		RotationalInertia: (*b2Body)(unsafe.Pointer(body)).Inertia,
	}
	return massData
}

func b2Body_ApplyMassFromShapes(tls *_Stack, bodyId BodyId) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	b2UpdateBodyMassData(tls, world, body)
}

func b2Body_SetLinearDamping(tls *_Stack, bodyId BodyId, linearDamping float32) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	(*b2BodySim)(unsafe.Pointer(bodySim)).LinearDamping = linearDamping
}

func b2Body_GetLinearDamping(tls *_Stack, bodyId BodyId) (r float32) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	return (*b2BodySim)(unsafe.Pointer(bodySim)).LinearDamping
}

func b2Body_SetAngularDamping(tls *_Stack, bodyId BodyId, angularDamping float32) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	(*b2BodySim)(unsafe.Pointer(bodySim)).AngularDamping = angularDamping
}

func b2Body_GetAngularDamping(tls *_Stack, bodyId BodyId) (r float32) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	return (*b2BodySim)(unsafe.Pointer(bodySim)).AngularDamping
}

func b2Body_SetGravityScale(tls *_Stack, bodyId BodyId, gravityScale float32) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	(*b2BodySim)(unsafe.Pointer(bodySim)).GravityScale = gravityScale
}

func b2Body_GetGravityScale(tls *_Stack, bodyId BodyId) (r float32) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	return (*b2BodySim)(unsafe.Pointer(bodySim)).GravityScale
}

func b2Body_IsAwake(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return boolUint8((*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet))
}

func b2Body_SetAwake(tls *_Stack, bodyId BodyId, awake uint8) {
	var body, island, world, v1 uintptr
	_, _, _, _ = body, island, world, v1
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	if awake != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeBody(tls, world, body)
	} else {
		if int32FromUint8(awake) == false1 && (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
			v1 = (*b2IslandArray)(unsafe.Pointer(world+1184)).Data + uintptr((*b2Body)(unsafe.Pointer(body)).IslandId)*56
			goto _2
		_2:
			island = v1
			if (*b2Island)(unsafe.Pointer(island)).ConstraintRemoveCount > 0 {
				// Must split the island before sleeping. This is expensive.
				b2SplitIsland(tls, world, (*b2Body)(unsafe.Pointer(body)).IslandId)
			}
			b2TrySleepIsland(tls, world, (*b2Body)(unsafe.Pointer(body)).IslandId)
		}
	}
}

func b2Body_IsEnabled(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return boolUint8((*b2Body)(unsafe.Pointer(body)).SetIndex != int32(b2_disabledSet))
}

func b2Body_IsSleepEnabled(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return (*b2Body)(unsafe.Pointer(body)).EnableSleep
}

func b2Body_SetSleepThreshold(tls *_Stack, bodyId BodyId, sleepThreshold float32) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	(*b2Body)(unsafe.Pointer(body)).SleepThreshold = sleepThreshold
}

func b2Body_GetSleepThreshold(tls *_Stack, bodyId BodyId) (r float32) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return (*b2Body)(unsafe.Pointer(body)).SleepThreshold
}

func b2Body_EnableSleep(tls *_Stack, bodyId BodyId, enableSleep uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	(*b2Body)(unsafe.Pointer(body)).EnableSleep = enableSleep
	if int32FromUint8(enableSleep) == false1 {
		b2WakeBody(tls, world, body)
	}
}

// C documentation
//
//	// Disabling a body requires a lot of detailed bookkeeping, but it is a valuable feature.
//	// The most challenging aspect that joints may connect to bodies that are not disabled.
func b2Body_Disable(tls *_Stack, bodyId BodyId) {
	var body, disabledSet, joint, jointSet, set, shape, world, v1, v3, v5, v7, v9 uintptr
	var edgeIndex, jointId, jointKey, shapeId int32
	var wakeBodies uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, disabledSet, edgeIndex, joint, jointId, jointKey, jointSet, set, shape, shapeId, wakeBodies, world, v1, v3, v5, v7, v9
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_disabledSet) {
		return
	}
	// Destroy contacts and wake bodies touching this body. This avoid floating bodies.
	// This is necessary even for static bodies.
	wakeBodies = uint8(true1)
	b2DestroyBodyContacts(tls, world, body, wakeBodies)
	// Disabled bodies are not in an island.
	b2RemoveBodyFromIsland(tls, world, body)
	// Remove shapes from broad-phase
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		v1 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr(shapeId)*288
		goto _2
	_2:
		shape = v1
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
		b2DestroyShapeProxy(tls, shape, world+40)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr((*b2Body)(unsafe.Pointer(body)).SetIndex)*88
	goto _4
_4:
	// Transfer simulation data to disabled set
	set = v3
	v5 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr(int32(b2_disabledSet))*88
	goto _6
_6:
	disabledSet = v5
	// Transfer body sim
	b2TransferBody(tls, world, disabledSet, set, body)
	// Unlink joints and transfer
	jointKey = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
	for jointKey != -int32(1) {
		jointId = jointKey >> int32(1)
		edgeIndex = jointKey & int32(1)
		v7 = (*b2JointArray)(unsafe.Pointer(world+1104)).Data + uintptr(jointId)*72
		goto _8
	_8:
		joint = v7
		jointKey = (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + uintptr(edgeIndex)*12))).NextKey
		// joint may already be disabled by other body
		if (*b2Joint)(unsafe.Pointer(joint)).SetIndex == int32(b2_disabledSet) {
			continue
		}
		// Remove joint from island
		if (*b2Joint)(unsafe.Pointer(joint)).IslandId != -int32(1) {
			b2UnlinkJoint(tls, world, joint)
		}
		v9 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr((*b2Joint)(unsafe.Pointer(joint)).SetIndex)*88
		goto _10
	_10:
		// Transfer joint to disabled set
		jointSet = v9
		b2TransferJoint(tls, world, disabledSet, jointSet, joint)
	}
	b2ValidateConnectivity(tls, world)
	b2ValidateSolverSets(tls, world)
}

func b2Body_Enable(tls *_Stack, bodyId BodyId) {
	var body, bodyA, bodyB, disabledSet, joint, jointSet, shape, targetSet, world, v1, v10, v12, v14, v4, v6, v8 uintptr
	var edgeIndex, jointId, jointKey, jointSetId, setId, shapeId, v3 int32
	var forcePairCreation, mergeIslands uint8
	var proxyType b2BodyType1
	var transform Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodyA, bodyB, disabledSet, edgeIndex, forcePairCreation, joint, jointId, jointKey, jointSet, jointSetId, mergeIslands, proxyType, setId, shape, shapeId, targetSet, transform, world, v1, v10, v12, v14, v3, v4, v6, v8
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).SetIndex != int32(b2_disabledSet) {
		return
	}
	v1 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr(int32(b2_disabledSet))*88
	goto _2
_2:
	disabledSet = v1
	if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_staticBody) {
		v3 = int32(b2_staticSet)
	} else {
		v3 = int32(b2_awakeSet)
	}
	setId = v3
	v4 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr(setId)*88
	goto _5
_5:
	targetSet = v4
	b2TransferBody(tls, world, targetSet, disabledSet, body)
	transform = b2GetBodyTransformQuick(tls, world, body)
	// Add shapes to broad-phase
	proxyType = (*b2Body)(unsafe.Pointer(body)).Type1
	forcePairCreation = uint8(true1)
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		v6 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr(shapeId)*288
		goto _7
	_7:
		shape = v6
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
		b2CreateShapeProxy(tls, shape, world+40, proxyType, transform, forcePairCreation)
	}
	if setId != int32(b2_staticSet) {
		b2CreateIslandForBody(tls, world, setId, body)
	}
	// Transfer joints. If the other body is disabled, don't transfer.
	// If the other body is sleeping, wake it.
	mergeIslands = uint8(false1)
	jointKey = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
	for jointKey != -int32(1) {
		jointId = jointKey >> int32(1)
		edgeIndex = jointKey & int32(1)
		v8 = (*b2JointArray)(unsafe.Pointer(world+1104)).Data + uintptr(jointId)*72
		goto _9
	_9:
		joint = v8
		jointKey = (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + uintptr(edgeIndex)*12))).NextKey
		v10 = (*b2BodyArray)(unsafe.Pointer(world+1024)).Data + uintptr((*(*b2JointEdge)(unsafe.Pointer(joint + 20))).BodyId)*128
		goto _11
	_11:
		bodyA = v10
		v12 = (*b2BodyArray)(unsafe.Pointer(world+1024)).Data + uintptr((*(*b2JointEdge)(unsafe.Pointer(joint + 20 + 1*12))).BodyId)*128
		goto _13
	_13:
		bodyB = v12
		if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_disabledSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_disabledSet) {
			// one body is still disabled
			continue
		}
		if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_staticSet) && (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_staticSet) {
			jointSetId = int32(b2_staticSet)
		} else {
			if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_staticSet) {
				jointSetId = (*b2Body)(unsafe.Pointer(bodyB)).SetIndex
			} else {
				jointSetId = (*b2Body)(unsafe.Pointer(bodyA)).SetIndex
			}
		}
		v14 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr(jointSetId)*88
		goto _15
	_15:
		jointSet = v14
		b2TransferJoint(tls, world, jointSet, disabledSet, joint)
		// Now that the joint is in the correct set, I can link the joint in the island.
		if jointSetId != int32(b2_staticSet) {
			b2LinkJoint(tls, world, joint, mergeIslands)
		}
	}
	// Now merge islands
	b2MergeAwakeIslands(tls, world)
	b2ValidateSolverSets(tls, world)
}

func b2Body_SetFixedRotation(tls *_Stack, bodyId BodyId, flag uint8) {
	var body, state, world uintptr
	_, _, _ = body, state, world
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).FixedRotation != flag {
		(*b2Body)(unsafe.Pointer(body)).FixedRotation = flag
		state = b2GetBodyState(tls, world, body)
		if state != uintptrFromInt32(0) {
			(*b2BodyState)(unsafe.Pointer(state)).AngularVelocity = float32FromFloat32(0)
		}
		b2UpdateBodyMassData(tls, world, body)
	}
}

func b2Body_IsFixedRotation(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return (*b2Body)(unsafe.Pointer(body)).FixedRotation
}

func b2Body_SetBullet(tls *_Stack, bodyId BodyId, flag uint8) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	(*b2BodySim)(unsafe.Pointer(bodySim)).IsBullet = flag
}

func b2Body_IsBullet(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	return (*b2BodySim)(unsafe.Pointer(bodySim)).IsBullet
}

func b2Body_EnableContactEvents(tls *_Stack, bodyId BodyId, flag uint8) {
	var body, shape, world, v1 uintptr
	var shapeId int32
	_, _, _, _, _ = body, shape, shapeId, world, v1
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		v1 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr(shapeId)*288
		goto _2
	_2:
		shape = v1
		(*b2Shape)(unsafe.Pointer(shape)).EnableContactEvents = flag
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
	}
}

func b2Body_EnableHitEvents(tls *_Stack, bodyId BodyId, flag uint8) {
	var body, shape, world, v1 uintptr
	var shapeId int32
	_, _, _, _, _ = body, shape, shapeId, world, v1
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		v1 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr(shapeId)*288
		goto _2
	_2:
		shape = v1
		(*b2Shape)(unsafe.Pointer(shape)).EnableHitEvents = flag
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
	}
}

func b2Body_GetWorld(tls *_Stack, bodyId BodyId) (r WorldId) {
	var world uintptr
	_ = world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	return WorldId{
		Index1:     uint16FromInt32(int32FromUint16(bodyId.World0) + int32(1)),
		Generation: (*b2World)(unsafe.Pointer(world)).Generation,
	}
}

func b2Body_GetShapeCount(tls *_Stack, bodyId BodyId) (r int32) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return (*b2Body)(unsafe.Pointer(body)).ShapeCount
}

func b2Body_GetShapes(tls *_Stack, bodyId BodyId, shapeArray uintptr, capacity int32) (r int32) {
	var body, shape, world, v1 uintptr
	var id ShapeId
	var shapeCount, shapeId int32
	_, _, _, _, _, _, _ = body, id, shape, shapeCount, shapeId, world, v1
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	shapeCount = 0
	for shapeId != -int32(1) && shapeCount < capacity {
		v1 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr(shapeId)*288
		goto _2
	_2:
		shape = v1
		id = ShapeId{
			Index1:     (*b2Shape)(unsafe.Pointer(shape)).Id + int32(1),
			World0:     bodyId.World0,
			Generation: (*b2Shape)(unsafe.Pointer(shape)).Generation,
		}
		*(*ShapeId)(unsafe.Pointer(shapeArray + uintptr(shapeCount)*8)) = id
		shapeCount += int32(1)
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
	}
	return shapeCount
}

func b2Body_GetJointCount(tls *_Stack, bodyId BodyId) (r int32) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return (*b2Body)(unsafe.Pointer(body)).JointCount
}

func b2Body_GetJoints(tls *_Stack, bodyId BodyId, jointArray uintptr, capacity int32) (r int32) {
	var body, joint, world, v1 uintptr
	var edgeIndex, jointCount, jointId, jointKey int32
	var id JointId
	_, _, _, _, _, _, _, _, _ = body, edgeIndex, id, joint, jointCount, jointId, jointKey, world, v1
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	jointKey = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
	jointCount = 0
	for jointKey != -int32(1) && jointCount < capacity {
		jointId = jointKey >> int32(1)
		edgeIndex = jointKey & int32(1)
		v1 = (*b2JointArray)(unsafe.Pointer(world+1104)).Data + uintptr(jointId)*72
		goto _2
	_2:
		joint = v1
		id = JointId{
			Index1:     jointId + int32(1),
			World0:     bodyId.World0,
			Generation: (*b2Joint)(unsafe.Pointer(joint)).Generation,
		}
		*(*JointId)(unsafe.Pointer(jointArray + uintptr(jointCount)*8)) = id
		jointCount += int32(1)
		jointKey = (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + uintptr(edgeIndex)*12))).NextKey
	}
	return jointCount
}

func b2BulletBodyTask(tls *_Stack, startIndex int32, endIndex int32, threadIndex uint32, taskContext uintptr) {
	var i, simIndex int32
	var stepContext uintptr
	_, _, _ = i, simIndex, stepContext
	_ = uint64FromInt64(4)
	stepContext = taskContext
	i = startIndex
	for {
		if !(i < endIndex) {
			break
		}
		simIndex = *(*int32)(unsafe.Pointer((*b2StepContext)(unsafe.Pointer(stepContext)).BulletBodies + uintptr(i)*4))
		b2SolveContinuous(tls, (*b2StepContext)(unsafe.Pointer(stepContext)).World, simIndex)
		goto _1
	_1:
		;
		i++
	}
}

func b2TransferBody(tls *_Stack, world uintptr, targetSet uintptr, sourceSet uintptr, body uintptr) {
	var movedBody, movedSim, sourceSim, state, targetSim, v1, v11, v13, v16, v18, v3, v5, v7 uintptr
	var movedId, movedIndex, movedIndex1, movedIndex2, newCapacity, newCapacity1, sourceIndex, targetIndex, v14, v17, v4, v8, v9 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = movedBody, movedId, movedIndex, movedIndex1, movedIndex2, movedSim, newCapacity, newCapacity1, sourceIndex, sourceSim, state, targetIndex, targetSim, v1, v11, v13, v14, v16, v17, v18, v3, v4, v5, v7, v8, v9
	sourceIndex = (*b2Body)(unsafe.Pointer(body)).LocalIndex
	v1 = (*b2BodySimArray)(unsafe.Pointer(sourceSet)).Data + uintptr(sourceIndex)*100
	goto _2
_2:
	sourceSim = v1
	targetIndex = (*b2SolverSet)(unsafe.Pointer(targetSet)).BodySims.Count
	v3 = targetSet
	if (*b2BodySimArray)(unsafe.Pointer(v3)).Count == (*b2BodySimArray)(unsafe.Pointer(v3)).Capacity {
		if (*b2BodySimArray)(unsafe.Pointer(v3)).Capacity < int32(2) {
			v4 = int32(2)
		} else {
			v4 = (*b2BodySimArray)(unsafe.Pointer(v3)).Capacity + (*b2BodySimArray)(unsafe.Pointer(v3)).Capacity>>int32(1)
		}
		newCapacity = v4
		b2BodySimArray_Reserve(tls, v3, newCapacity)
	}
	*(*int32)(unsafe.Pointer(v3 + 8)) += int32(1)
	v5 = (*b2BodySimArray)(unsafe.Pointer(v3)).Data + uintptr((*b2BodySimArray)(unsafe.Pointer(v3)).Count-int32FromInt32(1))*100
	goto _6
_6:
	targetSim = v5
	memcpy(tls, targetSim, sourceSim, uint64(100))
	v7 = sourceSet
	v8 = sourceIndex
	movedIndex = -int32(1)
	if v8 != (*b2BodySimArray)(unsafe.Pointer(v7)).Count-int32FromInt32(1) {
		movedIndex = (*b2BodySimArray)(unsafe.Pointer(v7)).Count - int32(1)
		*(*b2BodySim)(unsafe.Pointer((*b2BodySimArray)(unsafe.Pointer(v7)).Data + uintptr(v8)*100)) = *(*b2BodySim)(unsafe.Pointer((*b2BodySimArray)(unsafe.Pointer(v7)).Data + uintptr(movedIndex)*100))
	}
	*(*int32)(unsafe.Pointer(v7 + 8)) -= int32(1)
	v9 = movedIndex
	goto _10
_10:
	// Remove body sim from solver set that owns it
	movedIndex2 = v9
	if movedIndex2 != -int32(1) {
		// Fix moved body index
		movedSim = (*b2SolverSet)(unsafe.Pointer(sourceSet)).BodySims.Data + uintptr(sourceIndex)*100
		movedId = (*b2BodySim)(unsafe.Pointer(movedSim)).BodyId
		v11 = (*b2BodyArray)(unsafe.Pointer(world+1024)).Data + uintptr(movedId)*128
		goto _12
	_12:
		movedBody = v11
		(*b2Body)(unsafe.Pointer(movedBody)).LocalIndex = sourceIndex
	}
	if (*b2SolverSet)(unsafe.Pointer(sourceSet)).SetIndex == int32(b2_awakeSet) {
		v13 = sourceSet + 16
		v14 = sourceIndex
		movedIndex1 = -int32(1)
		if v14 != (*b2BodyStateArray)(unsafe.Pointer(v13)).Count-int32FromInt32(1) {
			movedIndex1 = (*b2BodyStateArray)(unsafe.Pointer(v13)).Count - int32(1)
			*(*b2BodyState)(unsafe.Pointer((*b2BodyStateArray)(unsafe.Pointer(v13)).Data + uintptr(v14)*32)) = *(*b2BodyState)(unsafe.Pointer((*b2BodyStateArray)(unsafe.Pointer(v13)).Data + uintptr(movedIndex1)*32))
		}
		*(*int32)(unsafe.Pointer(v13 + 8)) -= int32(1)
		_ = movedIndex1
		goto _15
	_15:
	} else {
		if (*b2SolverSet)(unsafe.Pointer(targetSet)).SetIndex == int32(b2_awakeSet) {
			v16 = targetSet + 16
			if (*b2BodyStateArray)(unsafe.Pointer(v16)).Count == (*b2BodyStateArray)(unsafe.Pointer(v16)).Capacity {
				if (*b2BodyStateArray)(unsafe.Pointer(v16)).Capacity < int32(2) {
					v17 = int32(2)
				} else {
					v17 = (*b2BodyStateArray)(unsafe.Pointer(v16)).Capacity + (*b2BodyStateArray)(unsafe.Pointer(v16)).Capacity>>int32(1)
				}
				newCapacity1 = v17
				b2BodyStateArray_Reserve(tls, v16, newCapacity1)
			}
			*(*int32)(unsafe.Pointer(v16 + 8)) += int32(1)
			v18 = (*b2BodyStateArray)(unsafe.Pointer(v16)).Data + uintptr((*b2BodyStateArray)(unsafe.Pointer(v16)).Count-int32FromInt32(1))*32
			goto _19
		_19:
			state = v18
			*(*b2BodyState)(unsafe.Pointer(state)) = b2_identityBodyState15
		}
	}
	(*b2Body)(unsafe.Pointer(body)).SetIndex = (*b2SolverSet)(unsafe.Pointer(targetSet)).SetIndex
	(*b2Body)(unsafe.Pointer(body)).LocalIndex = targetIndex
}

func b2DefaultBodyDef(tls *_Stack) (r BodyDef) {
	var def BodyDef
	_ = def
	def = BodyDef{}
	def.Type1 = int32(b2_staticBody)
	def.Rotation = b2Rot_identity
	def.SleepThreshold = float32(float32FromFloat32(0.05) * b2_lengthUnitsPerMeter)
	def.GravityScale = float32FromFloat32(1)
	def.EnableSleep = uint8(true1)
	def.IsAwake = uint8(true1)
	def.IsEnabled = uint8(true1)
	def.InternalValue = int32(B2_SECRET_COOKIE)
	return def
}

func b2BodyMoveEventArray_Create(tls *_Stack, capacity int32) (r b2BodyMoveEventArray) {
	var a b2BodyMoveEventArray
	_ = a
	a = b2BodyMoveEventArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(40)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2BodyMoveEventArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2BodyMoveEventArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2BodyMoveEventArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2BodyMoveEventArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2BodyMoveEventArray)(unsafe.Pointer(a)).Capacity)*uint64(40)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(40)))
	(*b2BodyMoveEventArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2BodyMoveEventArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2BodyMoveEventArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2BodyMoveEventArray)(unsafe.Pointer(a)).Capacity)*uint64(40)))
	(*b2BodyMoveEventArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2BodyMoveEventArray)(unsafe.Pointer(a)).Count = 0
	(*b2BodyMoveEventArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2Body_IsValid(tls *_Stack, id BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	if int32(B2_MAX_WORLDS) <= int32FromUint16(id.World0) {
		// invalid world
		return uint8(false1)
	}
	world = uintptr(unsafe.Pointer(&b2_worlds)) + uintptr(id.World0)*1792
	if int32FromUint16((*b2World)(unsafe.Pointer(world)).WorldId) != int32FromUint16(id.World0) {
		// world is free
		return uint8(false1)
	}
	if id.Index1 < int32(1) || (*b2World)(unsafe.Pointer(world)).Bodies.Count < id.Index1 {
		// invalid index
		return uint8(false1)
	}
	body = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(id.Index1-int32FromInt32(1))*128
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == -int32(1) {
		// this was freed
		return uint8(false1)
	}
	if int32FromUint16((*b2Body)(unsafe.Pointer(body)).Generation) != int32FromUint16(id.Generation) {
		// this id is orphaned
		return uint8(false1)
	}
	return uint8(true1)
}
