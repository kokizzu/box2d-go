package box2d

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b2DefaultWorldDef(tls *_Stack) (r WorldDef) {
	var def WorldDef
	_ = def
	def = WorldDef{}
	def.Gravity.X = float32FromFloat32(0)
	def.Gravity.Y = -float32FromFloat32(10)
	def.HitEventThreshold = float32(float32FromFloat32(1) * b2_lengthUnitsPerMeter)
	def.RestitutionThreshold = float32(float32FromFloat32(1) * b2_lengthUnitsPerMeter)
	def.MaxContactPushSpeed = float32(float32FromFloat32(3) * b2_lengthUnitsPerMeter)
	def.ContactHertz = float32(30)
	def.ContactDampingRatio = float32FromFloat32(10)
	// 400 meters per second, faster than the speed of sound
	def.MaximumLinearSpeed = float32(float32FromFloat32(400) * b2_lengthUnitsPerMeter)
	def.EnableSleep = uint8(true1)
	def.EnableContinuous = uint8(true1)
	def.InternalValue = int32(B2_SECRET_COOKIE)
	return def
}

func b2GetWorldFromId(tls *_Stack, id WorldId) (r uintptr) {
	var world uintptr
	_ = world
	world = uintptr(unsafe.Pointer(&b2_worlds)) + uintptr(int32FromUint16(id.Index1)-int32FromInt32(1))*1792
	return world
}

func b2GetWorld(tls *_Stack, index int32) (r uintptr) {
	var world uintptr
	_ = world
	world = uintptr(unsafe.Pointer(&b2_worlds)) + uintptr(index)*1792
	return world
}

func b2GetWorldLocked(tls *_Stack, index int32) (r uintptr) {
	var world uintptr
	_ = world
	world = uintptr(unsafe.Pointer(&b2_worlds)) + uintptr(index)*1792
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return uintptrFromInt32(0)
	}
	return world
}

func b2CreateWorld(tls *_Stack, def uintptr) (r WorldId) {
	var generation uint16_t
	var i, i1, newCapacity, worldId, v10, v12, v14, v16, v3, v5, v7, v8, v9 int32
	var set b2SolverSet
	var world, v13, v15, v2, v4, v6 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = generation, i, i1, newCapacity, set, world, worldId, v10, v12, v13, v14, v15, v16, v2, v3, v4, v5, v6, v7, v8, v9
	worldId = -int32(1)
	i = 0
	for {
		if !(i < int32(B2_MAX_WORLDS)) {
			break
		}
		if int32FromUint8(b2_worlds[i].InUse) == false1 {
			worldId = i
			break
		}
		goto _1
	_1:
		;
		i++
	}
	if worldId == -int32(1) {
		return WorldId{}
	}
	b2InitializeContactRegisters(tls)
	world = uintptr(unsafe.Pointer(&b2_worlds)) + uintptr(worldId)*1792
	generation = (*b2World)(unsafe.Pointer(world)).Generation
	*(*b2World)(unsafe.Pointer(world)) = b2World{}
	(*b2World)(unsafe.Pointer(world)).WorldId = uint16FromInt32(worldId)
	(*b2World)(unsafe.Pointer(world)).Generation = generation
	(*b2World)(unsafe.Pointer(world)).InUse = uint8(true1)
	(*b2World)(unsafe.Pointer(world)).Arena = b2CreateArenaAllocator(tls, int32(2048))
	b2CreateBroadPhase(tls, world+40)
	b2CreateGraph(tls, world+328, int32(16))
	// pools
	(*b2World)(unsafe.Pointer(world)).BodyIdPool = b2CreateIdPool(tls)
	(*b2World)(unsafe.Pointer(world)).Bodies = b2BodyArray_Create(tls, int32(16))
	(*b2World)(unsafe.Pointer(world)).SolverSets = b2SolverSetArray_Create(tls, int32(8))
	// add empty static, active, and disabled body sets
	(*b2World)(unsafe.Pointer(world)).SolverSetIdPool = b2CreateIdPool(tls)
	set = b2SolverSet{}
	// static set
	set.SetIndex = b2AllocId(tls, world+1040)
	v2 = world + 1064
	if (*b2SolverSetArray)(unsafe.Pointer(v2)).Count == (*b2SolverSetArray)(unsafe.Pointer(v2)).Capacity {
		if (*b2SolverSetArray)(unsafe.Pointer(v2)).Capacity < int32(2) {
			v3 = int32(2)
		} else {
			v3 = (*b2SolverSetArray)(unsafe.Pointer(v2)).Capacity + (*b2SolverSetArray)(unsafe.Pointer(v2)).Capacity>>int32(1)
		}
		newCapacity = v3
		b2SolverSetArray_Reserve(tls, v2, newCapacity)
	}
	*(*b2SolverSet)(unsafe.Pointer((*b2SolverSetArray)(unsafe.Pointer(v2)).Data + uintptr((*b2SolverSetArray)(unsafe.Pointer(v2)).Count)*88)) = set
	*(*int32)(unsafe.Pointer(v2 + 8)) += int32(1)
	// disabled set
	set.SetIndex = b2AllocId(tls, world+1040)
	v4 = world + 1064
	if (*b2SolverSetArray)(unsafe.Pointer(v4)).Count == (*b2SolverSetArray)(unsafe.Pointer(v4)).Capacity {
		if (*b2SolverSetArray)(unsafe.Pointer(v4)).Capacity < int32(2) {
			v5 = int32(2)
		} else {
			v5 = (*b2SolverSetArray)(unsafe.Pointer(v4)).Capacity + (*b2SolverSetArray)(unsafe.Pointer(v4)).Capacity>>int32(1)
		}
		newCapacity = v5
		b2SolverSetArray_Reserve(tls, v4, newCapacity)
	}
	*(*b2SolverSet)(unsafe.Pointer((*b2SolverSetArray)(unsafe.Pointer(v4)).Data + uintptr((*b2SolverSetArray)(unsafe.Pointer(v4)).Count)*88)) = set
	*(*int32)(unsafe.Pointer(v4 + 8)) += int32(1)
	// awake set
	set.SetIndex = b2AllocId(tls, world+1040)
	v6 = world + 1064
	if (*b2SolverSetArray)(unsafe.Pointer(v6)).Count == (*b2SolverSetArray)(unsafe.Pointer(v6)).Capacity {
		if (*b2SolverSetArray)(unsafe.Pointer(v6)).Capacity < int32(2) {
			v7 = int32(2)
		} else {
			v7 = (*b2SolverSetArray)(unsafe.Pointer(v6)).Capacity + (*b2SolverSetArray)(unsafe.Pointer(v6)).Capacity>>int32(1)
		}
		newCapacity = v7
		b2SolverSetArray_Reserve(tls, v6, newCapacity)
	}
	*(*b2SolverSet)(unsafe.Pointer((*b2SolverSetArray)(unsafe.Pointer(v6)).Data + uintptr((*b2SolverSetArray)(unsafe.Pointer(v6)).Count)*88)) = set
	*(*int32)(unsafe.Pointer(v6 + 8)) += int32(1)
	(*b2World)(unsafe.Pointer(world)).ShapeIdPool = b2CreateIdPool(tls)
	(*b2World)(unsafe.Pointer(world)).Shapes = b2ShapeArray_Create(tls, int32(16))
	(*b2World)(unsafe.Pointer(world)).ChainIdPool = b2CreateIdPool(tls)
	(*b2World)(unsafe.Pointer(world)).ChainShapes = b2ChainShapeArray_Create(tls, int32(4))
	(*b2World)(unsafe.Pointer(world)).ContactIdPool = b2CreateIdPool(tls)
	(*b2World)(unsafe.Pointer(world)).Contacts = b2ContactArray_Create(tls, int32(16))
	(*b2World)(unsafe.Pointer(world)).JointIdPool = b2CreateIdPool(tls)
	(*b2World)(unsafe.Pointer(world)).Joints = b2JointArray_Create(tls, int32(16))
	(*b2World)(unsafe.Pointer(world)).IslandIdPool = b2CreateIdPool(tls)
	(*b2World)(unsafe.Pointer(world)).Islands = b2IslandArray_Create(tls, int32(8))
	(*b2World)(unsafe.Pointer(world)).Sensors = b2SensorArray_Create(tls, int32(4))
	(*b2World)(unsafe.Pointer(world)).BodyMoveEvents = b2BodyMoveEventArray_Create(tls, int32(4))
	(*b2World)(unsafe.Pointer(world)).SensorBeginEvents = b2SensorBeginTouchEventArray_Create(tls, int32(4))
	*(*b2SensorEndTouchEventArray)(unsafe.Pointer(world + 1376)) = b2SensorEndTouchEventArray_Create(tls, int32(4))
	*(*b2SensorEndTouchEventArray)(unsafe.Pointer(world + 1376 + 1*16)) = b2SensorEndTouchEventArray_Create(tls, int32(4))
	(*b2World)(unsafe.Pointer(world)).ContactBeginEvents = b2ContactBeginTouchEventArray_Create(tls, int32(4))
	*(*b2ContactEndTouchEventArray)(unsafe.Pointer(world + 1408)) = b2ContactEndTouchEventArray_Create(tls, int32(4))
	*(*b2ContactEndTouchEventArray)(unsafe.Pointer(world + 1408 + 1*16)) = b2ContactEndTouchEventArray_Create(tls, int32(4))
	(*b2World)(unsafe.Pointer(world)).ContactHitEvents = b2ContactHitEventArray_Create(tls, int32(4))
	(*b2World)(unsafe.Pointer(world)).EndEventArrayIndex = 0
	(*b2World)(unsafe.Pointer(world)).StepIndex = uint64(0)
	(*b2World)(unsafe.Pointer(world)).SplitIslandId = -int32(1)
	(*b2World)(unsafe.Pointer(world)).ActiveTaskCount = 0
	(*b2World)(unsafe.Pointer(world)).TaskCount = 0
	(*b2World)(unsafe.Pointer(world)).Gravity = (*WorldDef)(unsafe.Pointer(def)).Gravity
	(*b2World)(unsafe.Pointer(world)).HitEventThreshold = (*WorldDef)(unsafe.Pointer(def)).HitEventThreshold
	(*b2World)(unsafe.Pointer(world)).RestitutionThreshold = (*WorldDef)(unsafe.Pointer(def)).RestitutionThreshold
	(*b2World)(unsafe.Pointer(world)).MaxLinearSpeed = (*WorldDef)(unsafe.Pointer(def)).MaximumLinearSpeed
	(*b2World)(unsafe.Pointer(world)).MaxContactPushSpeed = (*WorldDef)(unsafe.Pointer(def)).MaxContactPushSpeed
	(*b2World)(unsafe.Pointer(world)).ContactHertz = (*WorldDef)(unsafe.Pointer(def)).ContactHertz
	(*b2World)(unsafe.Pointer(world)).ContactDampingRatio = (*WorldDef)(unsafe.Pointer(def)).ContactDampingRatio
	if (*WorldDef)(unsafe.Pointer(def)).FrictionCallback == uintptrFromInt32(0) {
		(*b2World)(unsafe.Pointer(world)).FrictionCallback = __ccgo_fp(b2DefaultFrictionCallback)
	} else {
		(*b2World)(unsafe.Pointer(world)).FrictionCallback = (*WorldDef)(unsafe.Pointer(def)).FrictionCallback
	}
	if (*WorldDef)(unsafe.Pointer(def)).RestitutionCallback == uintptrFromInt32(0) {
		(*b2World)(unsafe.Pointer(world)).RestitutionCallback = __ccgo_fp(b2DefaultRestitutionCallback)
	} else {
		(*b2World)(unsafe.Pointer(world)).RestitutionCallback = (*WorldDef)(unsafe.Pointer(def)).RestitutionCallback
	}
	(*b2World)(unsafe.Pointer(world)).EnableSleep = (*WorldDef)(unsafe.Pointer(def)).EnableSleep
	(*b2World)(unsafe.Pointer(world)).Locked = uint8(false1)
	(*b2World)(unsafe.Pointer(world)).EnableWarmStarting = uint8(true1)
	(*b2World)(unsafe.Pointer(world)).EnableContinuous = (*WorldDef)(unsafe.Pointer(def)).EnableContinuous
	(*b2World)(unsafe.Pointer(world)).EnableSpeculative = uint8(true1)
	(*b2World)(unsafe.Pointer(world)).UserTreeTask = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).UserData = (*WorldDef)(unsafe.Pointer(def)).UserData
	if (*WorldDef)(unsafe.Pointer(def)).WorkerCount > 0 && (*WorldDef)(unsafe.Pointer(def)).EnqueueTask != uintptrFromInt32(0) && (*WorldDef)(unsafe.Pointer(def)).FinishTask != uintptrFromInt32(0) {
		v8 = (*WorldDef)(unsafe.Pointer(def)).WorkerCount
		v9 = int32(B2_MAX_WORKERS)
		if v8 < v9 {
			v12 = v8
		} else {
			v12 = v9
		}
		v10 = v12
		goto _11
	_11:
		(*b2World)(unsafe.Pointer(world)).WorkerCount = v10
		(*b2World)(unsafe.Pointer(world)).EnqueueTaskFcn = (*WorldDef)(unsafe.Pointer(def)).EnqueueTask
		(*b2World)(unsafe.Pointer(world)).FinishTaskFcn = (*WorldDef)(unsafe.Pointer(def)).FinishTask
		(*b2World)(unsafe.Pointer(world)).UserTaskContext = (*WorldDef)(unsafe.Pointer(def)).UserTaskContext
	} else {
		(*b2World)(unsafe.Pointer(world)).WorkerCount = int32(1)
		(*b2World)(unsafe.Pointer(world)).EnqueueTaskFcn = __ccgo_fp(b2DefaultAddTaskFcn)
		(*b2World)(unsafe.Pointer(world)).FinishTaskFcn = __ccgo_fp(b2DefaultFinishTaskFcn)
		(*b2World)(unsafe.Pointer(world)).UserTaskContext = uintptrFromInt32(0)
	}
	(*b2World)(unsafe.Pointer(world)).TaskContexts = b2TaskContextArray_Create(tls, (*b2World)(unsafe.Pointer(world)).WorkerCount)
	v13 = world + 1296
	v14 = (*b2World)(unsafe.Pointer(world)).WorkerCount
	b2TaskContextArray_Reserve(tls, v13, v14)
	(*b2TaskContextArray)(unsafe.Pointer(v13)).Count = v14
	(*b2World)(unsafe.Pointer(world)).SensorTaskContexts = b2SensorTaskContextArray_Create(tls, (*b2World)(unsafe.Pointer(world)).WorkerCount)
	v15 = world + 1312
	v16 = (*b2World)(unsafe.Pointer(world)).WorkerCount
	b2SensorTaskContextArray_Reserve(tls, v15, v16)
	(*b2SensorTaskContextArray)(unsafe.Pointer(v15)).Count = v16
	i1 = 0
	for {
		if !(i1 < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		(*(*b2TaskContext)(unsafe.Pointer((*b2World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(i1)*56))).ContactStateBitSet = b2CreateBitSet(tls, uint32(1024))
		(*(*b2TaskContext)(unsafe.Pointer((*b2World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(i1)*56))).EnlargedSimBitSet = b2CreateBitSet(tls, uint32(256))
		(*(*b2TaskContext)(unsafe.Pointer((*b2World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(i1)*56))).AwakeIslandBitSet = b2CreateBitSet(tls, uint32(256))
		(*(*b2SensorTaskContext)(unsafe.Pointer((*b2World)(unsafe.Pointer(world)).SensorTaskContexts.Data + uintptr(i1)*16))).EventBits = b2CreateBitSet(tls, uint32(128))
		goto _17
	_17:
		;
		i1++
	}
	(*b2World)(unsafe.Pointer(world)).DebugBodySet = b2CreateBitSet(tls, uint32(256))
	(*b2World)(unsafe.Pointer(world)).DebugJointSet = b2CreateBitSet(tls, uint32(256))
	(*b2World)(unsafe.Pointer(world)).DebugContactSet = b2CreateBitSet(tls, uint32(256))
	(*b2World)(unsafe.Pointer(world)).DebugIslandSet = b2CreateBitSet(tls, uint32(256))
	// add one to worldId so that 0 represents a null b2WorldId
	return WorldId{
		Index1:     uint16FromInt32(worldId + int32FromInt32(1)),
		Generation: (*b2World)(unsafe.Pointer(world)).Generation,
	}
}

func b2DestroyWorld(tls *_Stack, worldId WorldId) {
	var chain, set, world uintptr
	var chainCapacity, i, i1, i2, i3, sensorCount, setCapacity int32
	var generation uint16_t
	_, _, _, _, _, _, _, _, _, _, _ = chain, chainCapacity, generation, i, i1, i2, i3, sensorCount, set, setCapacity, world
	world = b2GetWorldFromId(tls, worldId)
	b2DestroyBitSet(tls, world+1464)
	b2DestroyBitSet(tls, world+1480)
	b2DestroyBitSet(tls, world+1496)
	b2DestroyBitSet(tls, world+1512)
	i = 0
	for {
		if !(i < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		b2DestroyBitSet(tls, (*b2World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i)*56)
		b2DestroyBitSet(tls, (*b2World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i)*56+16)
		b2DestroyBitSet(tls, (*b2World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i)*56+32)
		b2DestroyBitSet(tls, (*b2World)(unsafe.Pointer(world)).SensorTaskContexts.Data+uintptr(i)*16)
		goto _1
	_1:
		;
		i++
	}
	b2TaskContextArray_Destroy(tls, world+1296)
	b2SensorTaskContextArray_Destroy(tls, world+1312)
	b2BodyMoveEventArray_Destroy(tls, world+1328)
	b2SensorBeginTouchEventArray_Destroy(tls, world+1344)
	b2SensorEndTouchEventArray_Destroy(tls, world+1376+uintptr(0)*16)
	b2SensorEndTouchEventArray_Destroy(tls, world+1376+uintptr(1)*16)
	b2ContactBeginTouchEventArray_Destroy(tls, world+1360)
	b2ContactEndTouchEventArray_Destroy(tls, world+1408+uintptr(0)*16)
	b2ContactEndTouchEventArray_Destroy(tls, world+1408+uintptr(1)*16)
	b2ContactHitEventArray_Destroy(tls, world+1448)
	chainCapacity = (*b2World)(unsafe.Pointer(world)).ChainShapes.Count
	i1 = 0
	for {
		if !(i1 < chainCapacity) {
			break
		}
		chain = (*b2World)(unsafe.Pointer(world)).ChainShapes.Data + uintptr(i1)*48
		if (*b2ChainShape)(unsafe.Pointer(chain)).Id != -int32(1) {
			b2FreeChainData(tls, chain)
		} else {
		}
		goto _2
	_2:
		;
		i1++
	}
	sensorCount = (*b2World)(unsafe.Pointer(world)).Sensors.Count
	i2 = 0
	for {
		if !(i2 < sensorCount) {
			break
		}
		b2ShapeRefArray_Destroy(tls, (*b2World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i2)*40)
		b2ShapeRefArray_Destroy(tls, (*b2World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i2)*40+16)
		goto _3
	_3:
		;
		i2++
	}
	b2SensorArray_Destroy(tls, world+1280)
	b2BodyArray_Destroy(tls, world+1024)
	b2ShapeArray_Destroy(tls, world+1248)
	b2ChainShapeArray_Destroy(tls, world+1264)
	b2ContactArray_Destroy(tls, world+1144)
	b2JointArray_Destroy(tls, world+1104)
	b2IslandArray_Destroy(tls, world+1184)
	// Destroy solver sets
	setCapacity = (*b2World)(unsafe.Pointer(world)).SolverSets.Count
	i3 = 0
	for {
		if !(i3 < setCapacity) {
			break
		}
		set = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(i3)*88
		if (*b2SolverSet)(unsafe.Pointer(set)).SetIndex != -int32(1) {
			b2DestroySolverSet(tls, world, i3)
		}
		goto _4
	_4:
		;
		i3++
	}
	b2SolverSetArray_Destroy(tls, world+1064)
	b2DestroyGraph(tls, world+328)
	b2DestroyBroadPhase(tls, world+40)
	b2DestroyIdPool(tls, world+1000)
	b2DestroyIdPool(tls, world+1200)
	b2DestroyIdPool(tls, world+1224)
	b2DestroyIdPool(tls, world+1120)
	b2DestroyIdPool(tls, world+1080)
	b2DestroyIdPool(tls, world+1160)
	b2DestroyIdPool(tls, world+1040)
	b2DestroyArenaAllocator(tls, world)
	// Wipe world but preserve generation
	generation = (*b2World)(unsafe.Pointer(world)).Generation
	*(*b2World)(unsafe.Pointer(world)) = b2World{}
	(*b2World)(unsafe.Pointer(world)).WorldId = uint16(0)
	(*b2World)(unsafe.Pointer(world)).Generation = uint16FromInt32(int32FromUint16(generation) + int32(1))
}

func b2World_Step(tls *_Stack, worldId WorldId, timeStep float32, subStepCount int32) {
	bp := tls.Alloc(304)
	defer tls.Free(304)
	var a11, a21, a31, contactHertz, omega, v10, v11, v12, v15, v16, v6, v7, v8 float32
	var collideTicks, pairTicks, sensorTicks, solveTicks, stepTicks uint64_t
	var world uintptr
	var v1, v2, v3, v5 int32
	var v13, v17 b2Softness
	var _ /* context at bp+0 */ b2StepContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a21, a31, collideTicks, contactHertz, omega, pairTicks, sensorTicks, solveTicks, stepTicks, world, v1, v10, v11, v12, v13, v15, v16, v17, v2, v3, v5, v6, v7, v8
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	// Prepare to capture events
	// Ensure user does not access stale data if there is an early return
	(*b2BodyMoveEventArray)(unsafe.Pointer(world + 1328)).Count = 0
	(*b2SensorBeginTouchEventArray)(unsafe.Pointer(world + 1344)).Count = 0
	(*b2ContactBeginTouchEventArray)(unsafe.Pointer(world + 1360)).Count = 0
	(*b2ContactHitEventArray)(unsafe.Pointer(world + 1448)).Count = 0
	(*b2World)(unsafe.Pointer(world)).Profile = Profile{}
	if timeStep == float32FromFloat32(0) {
		// Swap end event array buffers
		(*b2World)(unsafe.Pointer(world)).EndEventArrayIndex = int32(1) - (*b2World)(unsafe.Pointer(world)).EndEventArrayIndex
		(*b2SensorEndTouchEventArray)(unsafe.Pointer(world + 1376 + uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16)).Count = 0
		(*b2ContactEndTouchEventArray)(unsafe.Pointer(world + 1408 + uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16)).Count = 0
		// todo_erin would be useful to still process collision while paused
		return
	}
	(*b2World)(unsafe.Pointer(world)).Locked = uint8(true1)
	(*b2World)(unsafe.Pointer(world)).ActiveTaskCount = 0
	(*b2World)(unsafe.Pointer(world)).TaskCount = 0
	stepTicks = b2GetTicks(tls)
	// Update collision pairs and create contacts
	pairTicks = b2GetTicks(tls)
	b2UpdateBroadPhasePairs(tls, world)
	(*b2World)(unsafe.Pointer(world)).Profile.Pairs = b2GetMilliseconds(tls, pairTicks)
	*(*b2StepContext)(unsafe.Pointer(bp)) = b2StepContext{}
	(*(*b2StepContext)(unsafe.Pointer(bp))).World = world
	(*(*b2StepContext)(unsafe.Pointer(bp))).Dt = timeStep
	v1 = int32(1)
	v2 = subStepCount
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	(*(*b2StepContext)(unsafe.Pointer(bp))).SubStepCount = v3
	if timeStep > float32FromFloat32(0) {
		(*(*b2StepContext)(unsafe.Pointer(bp))).Inv_dt = float32FromFloat32(1) / timeStep
		(*(*b2StepContext)(unsafe.Pointer(bp))).H = timeStep / float32((*(*b2StepContext)(unsafe.Pointer(bp))).SubStepCount)
		(*(*b2StepContext)(unsafe.Pointer(bp))).Inv_h = float32(float32((*(*b2StepContext)(unsafe.Pointer(bp))).SubStepCount) * (*(*b2StepContext)(unsafe.Pointer(bp))).Inv_dt)
	} else {
		(*(*b2StepContext)(unsafe.Pointer(bp))).Inv_dt = float32FromFloat32(0)
		(*(*b2StepContext)(unsafe.Pointer(bp))).H = float32FromFloat32(0)
		(*(*b2StepContext)(unsafe.Pointer(bp))).Inv_h = float32FromFloat32(0)
	}
	(*b2World)(unsafe.Pointer(world)).Inv_h = (*(*b2StepContext)(unsafe.Pointer(bp))).Inv_h
	v6 = (*b2World)(unsafe.Pointer(world)).ContactHertz
	v7 = float32(float32FromFloat32(0.125) * (*(*b2StepContext)(unsafe.Pointer(bp))).Inv_h)
	if v6 < v7 {
		v10 = v6
	} else {
		v10 = v7
	}
	v8 = v10
	goto _9
_9:
	// Hertz values get reduced for large time steps
	contactHertz = v8
	v11 = contactHertz
	v12 = (*(*b2StepContext)(unsafe.Pointer(bp))).H
	if v11 == float32FromFloat32(0) {
		v13 = b2Softness{}
		goto _14
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v11)
	a11 = float32(float32FromFloat32(2)*(*b2World)(unsafe.Pointer(world)).ContactDampingRatio) + float32(v12*omega)
	a21 = float32(float32(v12*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v13 = b2Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _14
_14:
	(*(*b2StepContext)(unsafe.Pointer(bp))).ContactSoftness = v13
	v15 = float32(float32FromFloat32(2) * contactHertz)
	v16 = (*(*b2StepContext)(unsafe.Pointer(bp))).H
	if v15 == float32FromFloat32(0) {
		v17 = b2Softness{}
		goto _18
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v15)
	a11 = float32(float32FromFloat32(2)*(*b2World)(unsafe.Pointer(world)).ContactDampingRatio) + float32(v16*omega)
	a21 = float32(float32(v16*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v17 = b2Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _18
_18:
	(*(*b2StepContext)(unsafe.Pointer(bp))).StaticSoftness = v17
	(*b2World)(unsafe.Pointer(world)).ContactSpeed = (*b2World)(unsafe.Pointer(world)).MaxContactPushSpeed / (*(*b2StepContext)(unsafe.Pointer(bp))).StaticSoftness.MassScale
	(*(*b2StepContext)(unsafe.Pointer(bp))).RestitutionThreshold = (*b2World)(unsafe.Pointer(world)).RestitutionThreshold
	(*(*b2StepContext)(unsafe.Pointer(bp))).MaxLinearVelocity = (*b2World)(unsafe.Pointer(world)).MaxLinearSpeed
	(*(*b2StepContext)(unsafe.Pointer(bp))).EnableWarmStarting = (*b2World)(unsafe.Pointer(world)).EnableWarmStarting
	// Update contacts
	collideTicks = b2GetTicks(tls)
	b2Collide(tls, bp)
	(*b2World)(unsafe.Pointer(world)).Profile.Collide = b2GetMilliseconds(tls, collideTicks)
	// Integrate velocities, solve velocity constraints, and integrate positions.
	if (*(*b2StepContext)(unsafe.Pointer(bp))).Dt > float32FromFloat32(0) {
		solveTicks = b2GetTicks(tls)
		b2Solve(tls, world, bp)
		(*b2World)(unsafe.Pointer(world)).Profile.Solve = b2GetMilliseconds(tls, solveTicks)
	}
	// Update sensors
	sensorTicks = b2GetTicks(tls)
	b2OverlapSensors(tls, world)
	(*b2World)(unsafe.Pointer(world)).Profile.Sensors = b2GetMilliseconds(tls, sensorTicks)
	(*b2World)(unsafe.Pointer(world)).Profile.Step = b2GetMilliseconds(tls, stepTicks)
	// Ensure stack is large enough
	b2GrowArena(tls, world)
	// Make sure all tasks that were started were also finished
	// Swap end event array buffers
	(*b2World)(unsafe.Pointer(world)).EndEventArrayIndex = int32(1) - (*b2World)(unsafe.Pointer(world)).EndEventArrayIndex
	(*b2SensorEndTouchEventArray)(unsafe.Pointer(world + 1376 + uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16)).Count = 0
	(*b2ContactEndTouchEventArray)(unsafe.Pointer(world + 1408 + uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16)).Count = 0
	(*b2World)(unsafe.Pointer(world)).Locked = uint8(false1)
}

func b2World_Draw(tls *_Stack, worldId WorldId, draw uintptr) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var aabb, aabb1, c, v55, v56, v77 AABB
	var addColor, color, color1, frictionColor, impulseColor, normalColor, persistColor, speculativeColor b2HexColor1
	var body, body1, body2, body3, bodySim, bodySim1, bodySim2, bodySim3, contact, graphColor, island, joint, point, set, set1, set2, shape, shape1, shape2, world, v12, v2, v20, v5, v51, v53, v9 uintptr
	var bodyCount, bodyCount1, bodyCount2, bodyId, bodyIndex, bodyIndex1, bodyIndex2, colorIndex, contactCount, contactIndex, count, count1, count2, i, i1, i2, j, pointCount, setCount, setCount1, setCount2, setIndex, setIndex1, setIndex2, shapeCount, shapeId, shapeId1, shapeId2 int32
	var colors [12]b2HexColor1
	var k_axisScale, k_impulseScale, linearSlop, mass, pointSize, x, y, v27, v31, v33, v38, v46, v57, v58, v59, v61, v62, v63, v64, v66, v67, v68, v69, v71, v72, v73, v74, v76 float32
	var normal, offset, offset1, p1, p11, p12, p13, p2, p21, p22, p23, tangent, v16, v17, v24, v25, v32, v34, v35, v37, v39, v40, v42, v43, v45, v47, v48 Vec2
	var transform, transform1, xf, v15, v23 Transform
	var _ /* buffer at bp+0 */ [32]uint8
	var _ /* buffer at bp+64 */ [32]uint8
	var _ /* buffer at bp+96 */ [32]uint8
	var _ /* vs at bp+128 */ [4]Vec2
	var _ /* vs at bp+32 */ [4]Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, aabb1, addColor, body, body1, body2, body3, bodyCount, bodyCount1, bodyCount2, bodyId, bodyIndex, bodyIndex1, bodyIndex2, bodySim, bodySim1, bodySim2, bodySim3, c, color, color1, colorIndex, colors, contact, contactCount, contactIndex, count, count1, count2, frictionColor, graphColor, i, i1, i2, impulseColor, island, j, joint, k_axisScale, k_impulseScale, linearSlop, mass, normal, normalColor, offset, offset1, p1, p11, p12, p13, p2, p21, p22, p23, persistColor, point, pointCount, pointSize, set, set1, set2, setCount, setCount1, setCount2, setIndex, setIndex1, setIndex2, shape, shape1, shape2, shapeCount, shapeId, shapeId1, shapeId2, speculativeColor, tangent, transform, transform1, world, x, xf, y, v12, v15, v16, v17, v2, v20, v23, v24, v25, v27, v31, v32, v33, v34, v35, v37, v38, v39, v40, v42, v43, v45, v46, v47, v48, v5, v51, v53, v55, v56, v57, v58, v59, v61, v62, v63, v64, v66, v67, v68, v69, v71, v72, v73, v74, v76, v77, v9
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	// todo it seems bounds drawing is fast enough for regular usage
	if (*DebugDraw)(unsafe.Pointer(draw)).UseDrawingBounds != 0 {
		b2DrawWithBounds(tls, world, draw)
		return
	}
	if (*DebugDraw)(unsafe.Pointer(draw)).DrawShapes != 0 {
		setCount = (*b2World)(unsafe.Pointer(world)).SolverSets.Count
		setIndex = 0
		for {
			if !(setIndex < setCount) {
				break
			}
			v2 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr(setIndex)*88
			goto _3
		_3:
			set = v2
			bodyCount = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count
			bodyIndex = 0
			for {
				if !(bodyIndex < bodyCount) {
					break
				}
				bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr(bodyIndex)*100
				v5 = (*b2BodyArray)(unsafe.Pointer(world+1024)).Data + uintptr((*b2BodySim)(unsafe.Pointer(bodySim)).BodyId)*128
				goto _6
			_6:
				body = v5
				xf = (*b2BodySim)(unsafe.Pointer(bodySim)).Transform
				shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
				for shapeId != -int32(1) {
					shape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*288
					if (*b2Shape)(unsafe.Pointer(shape)).CustomColor != uint32(0) {
						color = int32FromUint32((*b2Shape)(unsafe.Pointer(shape)).CustomColor)
					} else {
						if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_dynamicBody) && (*b2Body)(unsafe.Pointer(body)).Mass == float32FromFloat32(0) {
							// Bad body
							color = int32(b2_colorRed)
						} else {
							if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_disabledSet) {
								color = int32(b2_colorSlateGray)
							} else {
								if (*b2Shape)(unsafe.Pointer(shape)).SensorIndex != -int32(1) {
									color = int32(b2_colorWheat)
								} else {
									if (*b2BodySim)(unsafe.Pointer(bodySim)).IsBullet != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
										color = int32(b2_colorTurquoise)
									} else {
										if (*b2Body)(unsafe.Pointer(body)).IsSpeedCapped != 0 {
											color = int32(b2_colorYellow)
										} else {
											if (*b2BodySim)(unsafe.Pointer(bodySim)).IsFast != 0 {
												color = int32(b2_colorSalmon)
											} else {
												if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_staticBody) {
													color = int32(b2_colorPaleGreen)
												} else {
													if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_kinematicBody) {
														color = int32(b2_colorRoyalBlue)
													} else {
														if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
															color = int32(b2_colorPink)
														} else {
															color = int32(b2_colorGray)
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
					b2DrawShape(tls, draw, shape, xf, color)
					shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
				}
				goto _4
			_4:
				;
				bodyIndex++
			}
			goto _1
		_1:
			;
			setIndex++
		}
	}
	if (*DebugDraw)(unsafe.Pointer(draw)).DrawJoints != 0 {
		count = (*b2World)(unsafe.Pointer(world)).Joints.Count
		i = 0
		for {
			if !(i < count) {
				break
			}
			joint = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr(i)*72
			if (*b2Joint)(unsafe.Pointer(joint)).SetIndex == -int32(1) {
				goto _7
			}
			b2DrawJoint(tls, draw, world, joint)
			goto _7
		_7:
			;
			i++
		}
	}
	if (*DebugDraw)(unsafe.Pointer(draw)).DrawBounds != 0 {
		color1 = int32(b2_colorGold)
		setCount1 = (*b2World)(unsafe.Pointer(world)).SolverSets.Count
		setIndex1 = 0
		for {
			if !(setIndex1 < setCount1) {
				break
			}
			v9 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr(setIndex1)*88
			goto _10
		_10:
			set1 = v9
			bodyCount1 = (*b2SolverSet)(unsafe.Pointer(set1)).BodySims.Count
			bodyIndex1 = 0
			for {
				if !(bodyIndex1 < bodyCount1) {
					break
				}
				bodySim1 = (*b2SolverSet)(unsafe.Pointer(set1)).BodySims.Data + uintptr(bodyIndex1)*100
				__builtin_snprintf(tls, bp, uint64(32), __ccgo_ts+264, vaList(bp+168, (*b2BodySim)(unsafe.Pointer(bodySim1)).BodyId))
				(*(*func(*_Stack, Vec2, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*DebugDraw)(unsafe.Pointer(draw)).ｆDrawStringFcn})))(tls, (*b2BodySim)(unsafe.Pointer(bodySim1)).Center, bp, int32(b2_colorWhite), (*DebugDraw)(unsafe.Pointer(draw)).Context)
				v12 = (*b2BodyArray)(unsafe.Pointer(world+1024)).Data + uintptr((*b2BodySim)(unsafe.Pointer(bodySim1)).BodyId)*128
				goto _13
			_13:
				body1 = v12
				shapeId1 = (*b2Body)(unsafe.Pointer(body1)).HeadShapeId
				for shapeId1 != -int32(1) {
					shape1 = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId1)*288
					aabb = (*b2Shape)(unsafe.Pointer(shape1)).FatAABB
					*(*[4]Vec2)(unsafe.Pointer(bp + 32)) = [4]Vec2{
						0: {
							X: aabb.LowerBound.X,
							Y: aabb.LowerBound.Y,
						},
						1: {
							X: aabb.UpperBound.X,
							Y: aabb.LowerBound.Y,
						},
						2: {
							X: aabb.UpperBound.X,
							Y: aabb.UpperBound.Y,
						},
						3: {
							X: aabb.LowerBound.X,
							Y: aabb.UpperBound.Y,
						},
					}
					(*(*func(*_Stack, uintptr, int32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*DebugDraw)(unsafe.Pointer(draw)).ｆDrawPolygonFcn})))(tls, bp+32, int32(4), color1, (*DebugDraw)(unsafe.Pointer(draw)).Context)
					shapeId1 = (*b2Shape)(unsafe.Pointer(shape1)).NextShapeId
				}
				goto _11
			_11:
				;
				bodyIndex1++
			}
			goto _8
		_8:
			;
			setIndex1++
		}
	}
	if (*DebugDraw)(unsafe.Pointer(draw)).DrawBodyNames != 0 {
		offset = Vec2{
			X: float32FromFloat32(0.05),
			Y: float32FromFloat32(0.05),
		}
		count1 = (*b2World)(unsafe.Pointer(world)).Bodies.Count
		i1 = 0
		for {
			if !(i1 < count1) {
				break
			}
			body2 = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(i1)*128
			if (*b2Body)(unsafe.Pointer(body2)).SetIndex == -int32(1) {
				goto _14
			}
			if int32FromUint8(*(*uint8)(unsafe.Pointer(body2))) == 0 {
				goto _14
			}
			bodySim2 = b2GetBodySim(tls, world, body2)
			transform = Transform{
				P: (*b2BodySim)(unsafe.Pointer(bodySim2)).Center,
				Q: (*b2BodySim)(unsafe.Pointer(bodySim2)).Transform.Q,
			}
			v15 = transform
			v16 = offset
			x = float32(v15.Q.C*v16.X) - float32(v15.Q.S*v16.Y) + v15.P.X
			y = float32(v15.Q.S*v16.X) + float32(v15.Q.C*v16.Y) + v15.P.Y
			v17 = Vec2{
				X: x,
				Y: y,
			}
			goto _18
		_18:
			p1 = v17
			(*(*func(*_Stack, Vec2, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*DebugDraw)(unsafe.Pointer(draw)).ｆDrawStringFcn})))(tls, p1, body2, int32(b2_colorBlueViolet), (*DebugDraw)(unsafe.Pointer(draw)).Context)
			goto _14
		_14:
			;
			i1++
		}
	}
	if (*DebugDraw)(unsafe.Pointer(draw)).DrawMass != 0 {
		offset1 = Vec2{
			X: float32FromFloat32(0.1),
			Y: float32FromFloat32(0.1),
		}
		setCount2 = (*b2World)(unsafe.Pointer(world)).SolverSets.Count
		setIndex2 = 0
		for {
			if !(setIndex2 < setCount2) {
				break
			}
			v20 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr(setIndex2)*88
			goto _21
		_21:
			set2 = v20
			bodyCount2 = (*b2SolverSet)(unsafe.Pointer(set2)).BodySims.Count
			bodyIndex2 = 0
			for {
				if !(bodyIndex2 < bodyCount2) {
					break
				}
				bodySim3 = (*b2SolverSet)(unsafe.Pointer(set2)).BodySims.Data + uintptr(bodyIndex2)*100
				transform1 = Transform{
					P: (*b2BodySim)(unsafe.Pointer(bodySim3)).Center,
					Q: (*b2BodySim)(unsafe.Pointer(bodySim3)).Transform.Q,
				}
				(*(*func(*_Stack, Transform, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*DebugDraw)(unsafe.Pointer(draw)).ｆDrawTransformFcn})))(tls, transform1, (*DebugDraw)(unsafe.Pointer(draw)).Context)
				v23 = transform1
				v24 = offset1
				x = float32(v23.Q.C*v24.X) - float32(v23.Q.S*v24.Y) + v23.P.X
				y = float32(v23.Q.S*v24.X) + float32(v23.Q.C*v24.Y) + v23.P.Y
				v25 = Vec2{
					X: x,
					Y: y,
				}
				goto _26
			_26:
				p2 = v25
				if (*b2BodySim)(unsafe.Pointer(bodySim3)).InvMass > float32FromFloat32(0) {
					v27 = float32FromFloat32(1) / (*b2BodySim)(unsafe.Pointer(bodySim3)).InvMass
				} else {
					v27 = float32FromFloat32(0)
				}
				mass = v27
				__builtin_snprintf(tls, bp+64, uint64(32), __ccgo_ts+252, vaList(bp+168, float64(mass)))
				(*(*func(*_Stack, Vec2, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*DebugDraw)(unsafe.Pointer(draw)).ｆDrawStringFcn})))(tls, p2, bp+64, int32(b2_colorWhite), (*DebugDraw)(unsafe.Pointer(draw)).Context)
				goto _22
			_22:
				;
				bodyIndex2++
			}
			goto _19
		_19:
			;
			setIndex2++
		}
	}
	if (*DebugDraw)(unsafe.Pointer(draw)).DrawContacts != 0 {
		k_impulseScale = float32FromFloat32(1)
		k_axisScale = float32FromFloat32(0.3)
		linearSlop = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
		speculativeColor = int32(b2_colorLightGray)
		addColor = int32(b2_colorGreen)
		persistColor = int32(b2_colorBlue)
		normalColor = int32(b2_colorDimGray)
		impulseColor = int32(b2_colorMagenta)
		frictionColor = int32(b2_colorYellow)
		colors = [12]b2HexColor1{
			0:  int32(b2_colorRed),
			1:  int32(b2_colorOrange),
			2:  int32(b2_colorYellow),
			3:  int32(b2_colorGreen),
			4:  int32(b2_colorCyan),
			5:  int32(b2_colorBlue),
			6:  int32(b2_colorViolet),
			7:  int32(b2_colorPink),
			8:  int32(b2_colorChocolate),
			9:  int32(b2_colorGoldenRod),
			10: int32(b2_colorCoral),
		}
		colorIndex = 0
		for {
			if !(colorIndex < int32(B2_GRAPH_COLOR_COUNT)) {
				break
			}
			graphColor = world + 328 + uintptr(colorIndex)*56
			contactCount = (*b2GraphColor)(unsafe.Pointer(graphColor)).ContactSims.Count
			contactIndex = 0
			for {
				if !(contactIndex < contactCount) {
					break
				}
				contact = (*b2GraphColor)(unsafe.Pointer(graphColor)).ContactSims.Data + uintptr(contactIndex)*176
				pointCount = (*b2ContactSim)(unsafe.Pointer(contact)).Manifold.PointCount
				normal = (*b2ContactSim)(unsafe.Pointer(contact)).Manifold.Normal
				j = 0
				for {
					if !(j < pointCount) {
						break
					}
					point = contact + 36 + 12 + uintptr(j)*48
					if (*DebugDraw)(unsafe.Pointer(draw)).DrawGraphColors != 0 && 0 <= colorIndex && colorIndex <= int32(B2_GRAPH_COLOR_COUNT) {
						if colorIndex == int32FromInt32(B2_GRAPH_COLOR_COUNT)-int32FromInt32(1) {
							v31 = float32FromFloat32(7.5)
						} else {
							v31 = float32FromFloat32(5)
						}
						// graph color
						pointSize = v31
						(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*DebugDraw)(unsafe.Pointer(draw)).ｆDrawPointFcn})))(tls, (*ManifoldPoint)(unsafe.Pointer(point)).Point, pointSize, colors[colorIndex], (*DebugDraw)(unsafe.Pointer(draw)).Context)
						// m_context->draw.DrawString(point->position, "%d", point->color);
					} else {
						if (*ManifoldPoint)(unsafe.Pointer(point)).Separation > linearSlop {
							// Speculative
							(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*DebugDraw)(unsafe.Pointer(draw)).ｆDrawPointFcn})))(tls, (*ManifoldPoint)(unsafe.Pointer(point)).Point, float32FromFloat32(5), speculativeColor, (*DebugDraw)(unsafe.Pointer(draw)).Context)
						} else {
							if int32FromUint8((*ManifoldPoint)(unsafe.Pointer(point)).Persisted) == false1 {
								// Add
								(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*DebugDraw)(unsafe.Pointer(draw)).ｆDrawPointFcn})))(tls, (*ManifoldPoint)(unsafe.Pointer(point)).Point, float32FromFloat32(10), addColor, (*DebugDraw)(unsafe.Pointer(draw)).Context)
							} else {
								if int32FromUint8((*ManifoldPoint)(unsafe.Pointer(point)).Persisted) == int32(true1) {
									// Persist
									(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*DebugDraw)(unsafe.Pointer(draw)).ｆDrawPointFcn})))(tls, (*ManifoldPoint)(unsafe.Pointer(point)).Point, float32FromFloat32(5), persistColor, (*DebugDraw)(unsafe.Pointer(draw)).Context)
								}
							}
						}
					}
					if (*DebugDraw)(unsafe.Pointer(draw)).DrawContactNormals != 0 {
						p11 = (*ManifoldPoint)(unsafe.Pointer(point)).Point
						v32 = p11
						v33 = k_axisScale
						v34 = normal
						v35 = Vec2{
							X: v32.X + float32(v33*v34.X),
							Y: v32.Y + float32(v33*v34.Y),
						}
						goto _36
					_36:
						p21 = v35
						(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*DebugDraw)(unsafe.Pointer(draw)).ｆDrawSegmentFcn})))(tls, p11, p21, normalColor, (*DebugDraw)(unsafe.Pointer(draw)).Context)
					} else {
						if (*DebugDraw)(unsafe.Pointer(draw)).DrawContactImpulses != 0 {
							p12 = (*ManifoldPoint)(unsafe.Pointer(point)).Point
							v37 = p12
							v38 = float32(k_impulseScale * (*ManifoldPoint)(unsafe.Pointer(point)).TotalNormalImpulse)
							v39 = normal
							v40 = Vec2{
								X: v37.X + float32(v38*v39.X),
								Y: v37.Y + float32(v38*v39.Y),
							}
							goto _41
						_41:
							p22 = v40
							(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*DebugDraw)(unsafe.Pointer(draw)).ｆDrawSegmentFcn})))(tls, p12, p22, impulseColor, (*DebugDraw)(unsafe.Pointer(draw)).Context)
							__builtin_snprintf(tls, bp+96, uint64FromInt32(int32FromUint64(uint64FromInt64(32)/uint64FromInt64(1))), __ccgo_ts+267, vaList(bp+168, float64(float32FromFloat32(1000)*(*ManifoldPoint)(unsafe.Pointer(point)).TotalNormalImpulse)))
							(*(*func(*_Stack, Vec2, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*DebugDraw)(unsafe.Pointer(draw)).ｆDrawStringFcn})))(tls, p12, bp+96, int32(b2_colorWhite), (*DebugDraw)(unsafe.Pointer(draw)).Context)
						}
					}
					if (*DebugDraw)(unsafe.Pointer(draw)).DrawContactFeatures != 0 {
						__builtin_snprintf(tls, bp+96, uint64FromInt32(int32FromUint64(uint64FromInt64(32)/uint64FromInt64(1))), __ccgo_ts+264, vaList(bp+168, int32FromUint16((*ManifoldPoint)(unsafe.Pointer(point)).Id)))
						(*(*func(*_Stack, Vec2, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*DebugDraw)(unsafe.Pointer(draw)).ｆDrawStringFcn})))(tls, (*ManifoldPoint)(unsafe.Pointer(point)).Point, bp+96, int32(b2_colorOrange), (*DebugDraw)(unsafe.Pointer(draw)).Context)
					}
					if (*DebugDraw)(unsafe.Pointer(draw)).DrawFrictionImpulses != 0 {
						v42 = normal
						v43 = Vec2{
							X: v42.Y,
							Y: -v42.X,
						}
						goto _44
					_44:
						tangent = v43
						p13 = (*ManifoldPoint)(unsafe.Pointer(point)).Point
						v45 = p13
						v46 = float32(k_impulseScale * (*ManifoldPoint)(unsafe.Pointer(point)).TangentImpulse)
						v47 = tangent
						v48 = Vec2{
							X: v45.X + float32(v46*v47.X),
							Y: v45.Y + float32(v46*v47.Y),
						}
						goto _49
					_49:
						p23 = v48
						(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*DebugDraw)(unsafe.Pointer(draw)).ｆDrawSegmentFcn})))(tls, p13, p23, frictionColor, (*DebugDraw)(unsafe.Pointer(draw)).Context)
						__builtin_snprintf(tls, bp+96, uint64FromInt32(int32FromUint64(uint64FromInt64(32)/uint64FromInt64(1))), __ccgo_ts+267, vaList(bp+168, float64((*ManifoldPoint)(unsafe.Pointer(point)).TangentImpulse)))
						(*(*func(*_Stack, Vec2, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*DebugDraw)(unsafe.Pointer(draw)).ｆDrawStringFcn})))(tls, p13, bp+96, int32(b2_colorWhite), (*DebugDraw)(unsafe.Pointer(draw)).Context)
					}
					goto _30
				_30:
					;
					j++
				}
				goto _29
			_29:
				;
				contactIndex++
			}
			goto _28
		_28:
			;
			colorIndex++
		}
	}
	if (*DebugDraw)(unsafe.Pointer(draw)).DrawIslands != 0 {
		count2 = (*b2World)(unsafe.Pointer(world)).Islands.Count
		i2 = 0
		for {
			if !(i2 < count2) {
				break
			}
			island = (*b2World)(unsafe.Pointer(world)).Islands.Data + uintptr(i2)*56
			if (*b2Island)(unsafe.Pointer(island)).SetIndex == -int32(1) {
				goto _50
			}
			shapeCount = 0
			aabb1 = AABB{
				LowerBound: Vec2{
					X: float32FromFloat32(3.4028234663852886e+38),
					Y: float32FromFloat32(3.4028234663852886e+38),
				},
				UpperBound: Vec2{
					X: -float32FromFloat32(3.4028234663852886e+38),
					Y: -float32FromFloat32(3.4028234663852886e+38),
				},
			}
			bodyId = (*b2Island)(unsafe.Pointer(island)).HeadBody
			for bodyId != -int32(1) {
				v51 = (*b2BodyArray)(unsafe.Pointer(world+1024)).Data + uintptr(bodyId)*128
				goto _52
			_52:
				body3 = v51
				shapeId2 = (*b2Body)(unsafe.Pointer(body3)).HeadShapeId
				for shapeId2 != -int32(1) {
					v53 = (*b2ShapeArray)(unsafe.Pointer(world+1248)).Data + uintptr(shapeId2)*288
					goto _54
				_54:
					shape2 = v53
					v55 = aabb1
					v56 = (*b2Shape)(unsafe.Pointer(shape2)).FatAABB
					v57 = v55.LowerBound.X
					v58 = v56.LowerBound.X
					if v57 < v58 {
						v61 = v57
					} else {
						v61 = v58
					}
					v59 = v61
					goto _60
				_60:
					c.LowerBound.X = v59
					v62 = v55.LowerBound.Y
					v63 = v56.LowerBound.Y
					if v62 < v63 {
						v66 = v62
					} else {
						v66 = v63
					}
					v64 = v66
					goto _65
				_65:
					c.LowerBound.Y = v64
					v67 = v55.UpperBound.X
					v68 = v56.UpperBound.X
					if v67 > v68 {
						v71 = v67
					} else {
						v71 = v68
					}
					v69 = v71
					goto _70
				_70:
					c.UpperBound.X = v69
					v72 = v55.UpperBound.Y
					v73 = v56.UpperBound.Y
					if v72 > v73 {
						v76 = v72
					} else {
						v76 = v73
					}
					v74 = v76
					goto _75
				_75:
					c.UpperBound.Y = v74
					v77 = c
					goto _78
				_78:
					aabb1 = v77
					shapeCount += int32(1)
					shapeId2 = (*b2Shape)(unsafe.Pointer(shape2)).NextShapeId
				}
				bodyId = (*b2Body)(unsafe.Pointer(body3)).IslandNext
			}
			if shapeCount > 0 {
				*(*[4]Vec2)(unsafe.Pointer(bp + 128)) = [4]Vec2{
					0: {
						X: aabb1.LowerBound.X,
						Y: aabb1.LowerBound.Y,
					},
					1: {
						X: aabb1.UpperBound.X,
						Y: aabb1.LowerBound.Y,
					},
					2: {
						X: aabb1.UpperBound.X,
						Y: aabb1.UpperBound.Y,
					},
					3: {
						X: aabb1.LowerBound.X,
						Y: aabb1.UpperBound.Y,
					},
				}
				(*(*func(*_Stack, uintptr, int32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*DebugDraw)(unsafe.Pointer(draw)).ｆDrawPolygonFcn})))(tls, bp+128, int32(4), int32(b2_colorOrangeRed), (*DebugDraw)(unsafe.Pointer(draw)).Context)
			}
			goto _50
		_50:
			;
			i2++
		}
	}
}

func b2World_GetBodyEvents(tls *_Stack, worldId WorldId) (r b2BodyEvents) {
	var count int32
	var events b2BodyEvents
	var world uintptr
	_, _, _ = count, events, world
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return b2BodyEvents{}
	}
	count = (*b2World)(unsafe.Pointer(world)).BodyMoveEvents.Count
	events = b2BodyEvents{
		MoveEvents: (*b2World)(unsafe.Pointer(world)).BodyMoveEvents.Data,
		MoveCount:  count,
	}
	return events
}

func b2World_GetSensorEvents(tls *_Stack, worldId WorldId) (r b2SensorEvents) {
	var beginCount, endCount, endEventArrayIndex int32
	var events b2SensorEvents
	var world uintptr
	_, _, _, _, _ = beginCount, endCount, endEventArrayIndex, events, world
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return b2SensorEvents{}
	}
	// Careful to use previous buffer
	endEventArrayIndex = int32(1) - (*b2World)(unsafe.Pointer(world)).EndEventArrayIndex
	beginCount = (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Count
	endCount = (*(*b2SensorEndTouchEventArray)(unsafe.Pointer(world + 1376 + uintptr(endEventArrayIndex)*16))).Count
	events = b2SensorEvents{
		BeginEvents: (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Data,
		EndEvents:   (*(*b2SensorEndTouchEventArray)(unsafe.Pointer(world + 1376 + uintptr(endEventArrayIndex)*16))).Data,
		BeginCount:  beginCount,
		EndCount:    endCount,
	}
	return events
}

func b2World_GetContactEvents(tls *_Stack, worldId WorldId) (r b2ContactEvents) {
	var beginCount, endCount, endEventArrayIndex, hitCount int32
	var events b2ContactEvents
	var world uintptr
	_, _, _, _, _, _ = beginCount, endCount, endEventArrayIndex, events, hitCount, world
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return b2ContactEvents{}
	}
	// Careful to use previous buffer
	endEventArrayIndex = int32(1) - (*b2World)(unsafe.Pointer(world)).EndEventArrayIndex
	beginCount = (*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Count
	endCount = (*(*b2ContactEndTouchEventArray)(unsafe.Pointer(world + 1408 + uintptr(endEventArrayIndex)*16))).Count
	hitCount = (*b2World)(unsafe.Pointer(world)).ContactHitEvents.Count
	events = b2ContactEvents{
		BeginEvents: (*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Data,
		EndEvents:   (*(*b2ContactEndTouchEventArray)(unsafe.Pointer(world + 1408 + uintptr(endEventArrayIndex)*16))).Data,
		HitEvents:   (*b2World)(unsafe.Pointer(world)).ContactHitEvents.Data,
		BeginCount:  beginCount,
		EndCount:    endCount,
		HitCount:    hitCount,
	}
	return events
}

func b2World_IsValid(tls *_Stack, id WorldId) (r uint8) {
	var world uintptr
	_ = world
	if int32FromUint16(id.Index1) < int32(1) || int32(B2_MAX_WORLDS) < int32FromUint16(id.Index1) {
		return uint8(false1)
	}
	world = uintptr(unsafe.Pointer(&b2_worlds)) + uintptr(int32FromUint16(id.Index1)-int32FromInt32(1))*1792
	if int32FromUint16((*b2World)(unsafe.Pointer(world)).WorldId) != int32FromUint16(id.Index1)-int32(1) {
		// world is not allocated
		return uint8(false1)
	}
	return boolUint8(int32FromUint16(id.Generation) == int32FromUint16((*b2World)(unsafe.Pointer(world)).Generation))
}

func b2World_EnableSleeping(tls *_Stack, worldId WorldId, flag uint8) {
	var i, setCount int32
	var set, world, v2 uintptr
	_, _, _, _, _ = i, set, setCount, world, v2
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if flag == (*b2World)(unsafe.Pointer(world)).EnableSleep {
		return
	}
	(*b2World)(unsafe.Pointer(world)).EnableSleep = flag
	if int32FromUint8(flag) == false1 {
		setCount = (*b2World)(unsafe.Pointer(world)).SolverSets.Count
		i = int32(b2_firstSleepingSet)
		for {
			if !(i < setCount) {
				break
			}
			v2 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr(i)*88
			goto _3
		_3:
			set = v2
			if (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count > 0 {
				b2WakeSolverSet(tls, world, i)
			}
			goto _1
		_1:
			;
			i++
		}
	}
}

func b2World_IsSleepingEnabled(tls *_Stack, worldId WorldId) (r uint8) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).EnableSleep
}

func b2World_EnableWarmStarting(tls *_Stack, worldId WorldId, flag uint8) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	(*b2World)(unsafe.Pointer(world)).EnableWarmStarting = flag
}

func b2World_IsWarmStartingEnabled(tls *_Stack, worldId WorldId) (r uint8) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).EnableWarmStarting
}

func b2World_GetAwakeBodyCount(tls *_Stack, worldId WorldId) (r int32) {
	var awakeSet, world, v1 uintptr
	_, _, _ = awakeSet, world, v1
	world = b2GetWorldFromId(tls, worldId)
	v1 = (*b2SolverSetArray)(unsafe.Pointer(world+1064)).Data + uintptr(int32(b2_awakeSet))*88
	goto _2
_2:
	awakeSet = v1
	return (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Count
}

func b2World_EnableContinuous(tls *_Stack, worldId WorldId, flag uint8) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	(*b2World)(unsafe.Pointer(world)).EnableContinuous = flag
}

func b2World_IsContinuousEnabled(tls *_Stack, worldId WorldId) (r uint8) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).EnableContinuous
}

func b2World_SetRestitutionThreshold(tls *_Stack, worldId WorldId, value float32) {
	var world uintptr
	var v1, v2, v3, v4, v6, v7 float32
	_, _, _, _, _, _, _ = world, v1, v2, v3, v4, v6, v7
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	v1 = value
	v2 = float32FromFloat32(0)
	v3 = float32FromFloat32(3.4028234663852886e+38)
	if v1 < v2 {
		v6 = v2
	} else {
		if v1 > v3 {
			v7 = v3
		} else {
			v7 = v1
		}
		v6 = v7
	}
	v4 = v6
	goto _5
_5:
	(*b2World)(unsafe.Pointer(world)).RestitutionThreshold = v4
}

func b2World_GetRestitutionThreshold(tls *_Stack, worldId WorldId) (r float32) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).RestitutionThreshold
}

func b2World_SetHitEventThreshold(tls *_Stack, worldId WorldId, value float32) {
	var world uintptr
	var v1, v2, v3, v4, v6, v7 float32
	_, _, _, _, _, _, _ = world, v1, v2, v3, v4, v6, v7
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	v1 = value
	v2 = float32FromFloat32(0)
	v3 = float32FromFloat32(3.4028234663852886e+38)
	if v1 < v2 {
		v6 = v2
	} else {
		if v1 > v3 {
			v7 = v3
		} else {
			v7 = v1
		}
		v6 = v7
	}
	v4 = v6
	goto _5
_5:
	(*b2World)(unsafe.Pointer(world)).HitEventThreshold = v4
}

func b2World_GetHitEventThreshold(tls *_Stack, worldId WorldId) (r float32) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).HitEventThreshold
}

func b2World_SetContactTuning(tls *_Stack, worldId WorldId, hertz float32, dampingRatio float32, pushSpeed float32) {
	var world uintptr
	var v1, v10, v11, v13, v14, v15, v16, v17, v18, v2, v20, v21, v3, v4, v6, v7, v8, v9 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = world, v1, v10, v11, v13, v14, v15, v16, v17, v18, v2, v20, v21, v3, v4, v6, v7, v8, v9
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	v1 = hertz
	v2 = float32FromFloat32(0)
	v3 = float32FromFloat32(3.4028234663852886e+38)
	if v1 < v2 {
		v6 = v2
	} else {
		if v1 > v3 {
			v7 = v3
		} else {
			v7 = v1
		}
		v6 = v7
	}
	v4 = v6
	goto _5
_5:
	(*b2World)(unsafe.Pointer(world)).ContactHertz = v4
	v8 = dampingRatio
	v9 = float32FromFloat32(0)
	v10 = float32FromFloat32(3.4028234663852886e+38)
	if v8 < v9 {
		v13 = v9
	} else {
		if v8 > v10 {
			v14 = v10
		} else {
			v14 = v8
		}
		v13 = v14
	}
	v11 = v13
	goto _12
_12:
	(*b2World)(unsafe.Pointer(world)).ContactDampingRatio = v11
	v15 = pushSpeed
	v16 = float32FromFloat32(0)
	v17 = float32FromFloat32(3.4028234663852886e+38)
	if v15 < v16 {
		v20 = v16
	} else {
		if v15 > v17 {
			v21 = v17
		} else {
			v21 = v15
		}
		v20 = v21
	}
	v18 = v20
	goto _19
_19:
	(*b2World)(unsafe.Pointer(world)).MaxContactPushSpeed = v18
}

func b2World_SetMaximumLinearSpeed(tls *_Stack, worldId WorldId, maximumLinearSpeed float32) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	(*b2World)(unsafe.Pointer(world)).MaxLinearSpeed = maximumLinearSpeed
}

func b2World_GetMaximumLinearSpeed(tls *_Stack, worldId WorldId) (r float32) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).MaxLinearSpeed
}

func b2World_GetProfile(tls *_Stack, worldId WorldId) (r Profile) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).Profile
}

func b2World_GetCounters(tls *_Stack, worldId WorldId) (r Counters) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var dynamicTree, kinematicTree, staticTree, world, v1, v10, v13, v4, v7 uintptr
	var i, v11, v14, v16, v17, v18, v2, v20, v5, v8 int32
	var _ /* s at bp+0 */ Counters
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = dynamicTree, i, kinematicTree, staticTree, world, v1, v10, v11, v13, v14, v16, v17, v18, v2, v20, v4, v5, v7, v8
	world = b2GetWorldFromId(tls, worldId)
	*(*Counters)(unsafe.Pointer(bp)) = Counters{}
	v1 = world + 1000
	v2 = (*b2IdPool)(unsafe.Pointer(v1)).NextIndex - (*b2IdPool)(unsafe.Pointer(v1)).FreeArray.Count
	goto _3
_3:
	(*(*Counters)(unsafe.Pointer(bp))).BodyCount = v2
	v4 = world + 1200
	v5 = (*b2IdPool)(unsafe.Pointer(v4)).NextIndex - (*b2IdPool)(unsafe.Pointer(v4)).FreeArray.Count
	goto _6
_6:
	(*(*Counters)(unsafe.Pointer(bp))).ShapeCount = v5
	v7 = world + 1120
	v8 = (*b2IdPool)(unsafe.Pointer(v7)).NextIndex - (*b2IdPool)(unsafe.Pointer(v7)).FreeArray.Count
	goto _9
_9:
	(*(*Counters)(unsafe.Pointer(bp))).ContactCount = v8
	v10 = world + 1080
	v11 = (*b2IdPool)(unsafe.Pointer(v10)).NextIndex - (*b2IdPool)(unsafe.Pointer(v10)).FreeArray.Count
	goto _12
_12:
	(*(*Counters)(unsafe.Pointer(bp))).JointCount = v11
	v13 = world + 1160
	v14 = (*b2IdPool)(unsafe.Pointer(v13)).NextIndex - (*b2IdPool)(unsafe.Pointer(v13)).FreeArray.Count
	goto _15
_15:
	(*(*Counters)(unsafe.Pointer(bp))).IslandCount = v14
	staticTree = world + 40 + uintptr(b2_staticBody)*72
	(*(*Counters)(unsafe.Pointer(bp))).StaticTreeHeight = b2DynamicTree_GetHeight(tls, staticTree)
	dynamicTree = world + 40 + uintptr(b2_dynamicBody)*72
	kinematicTree = world + 40 + uintptr(b2_kinematicBody)*72
	v16 = b2DynamicTree_GetHeight(tls, dynamicTree)
	v17 = b2DynamicTree_GetHeight(tls, kinematicTree)
	if v16 > v17 {
		v20 = v16
	} else {
		v20 = v17
	}
	v18 = v20
	goto _19
_19:
	(*(*Counters)(unsafe.Pointer(bp))).TreeHeight = v18
	(*(*Counters)(unsafe.Pointer(bp))).StackUsed = b2GetMaxArenaAllocation(tls, world)
	(*(*Counters)(unsafe.Pointer(bp))).ByteCount = b2GetByteCount(tls)
	(*(*Counters)(unsafe.Pointer(bp))).TaskCount = (*b2World)(unsafe.Pointer(world)).TaskCount
	i = 0
	for {
		if !(i < int32(B2_GRAPH_COLOR_COUNT)) {
			break
		}
		*(*int32)(unsafe.Pointer(bp + 40 + uintptr(i)*4)) = (*(*b2GraphColor)(unsafe.Pointer(world + 328 + uintptr(i)*56))).ContactSims.Count + (*(*b2GraphColor)(unsafe.Pointer(world + 328 + uintptr(i)*56))).JointSims.Count
		goto _21
	_21:
		;
		i++
	}
	return *(*Counters)(unsafe.Pointer(bp))
}

func b2World_SetUserData(tls *_Stack, worldId WorldId, userData uintptr) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	(*b2World)(unsafe.Pointer(world)).UserData = userData
}

func b2World_GetUserData(tls *_Stack, worldId WorldId) (r uintptr) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).UserData
}

func b2World_SetFrictionCallback(tls *_Stack, worldId WorldId, callback uintptr) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if callback != uintptrFromInt32(0) {
		(*b2World)(unsafe.Pointer(world)).FrictionCallback = callback
	} else {
		(*b2World)(unsafe.Pointer(world)).FrictionCallback = __ccgo_fp(b2DefaultFrictionCallback)
	}
}

func b2World_SetRestitutionCallback(tls *_Stack, worldId WorldId, callback uintptr) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if callback != uintptrFromInt32(0) {
		(*b2World)(unsafe.Pointer(world)).RestitutionCallback = callback
	} else {
		(*b2World)(unsafe.Pointer(world)).RestitutionCallback = __ccgo_fp(b2DefaultRestitutionCallback)
	}
}

func b2World_DumpMemoryStats(tls *_Stack, worldId WorldId) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var bodyBitSetBytes, bodySimCapacity, bodyStateCapacity, contactSimCapacity, i, i1, islandSimCapacity, jointSimCapacity, solverSetCapacity, v1, v11, v13, v15, v17, v19, v21, v23, v25, v27, v29, v3, v31, v33, v35, v37, v39, v41, v43, v47, v5, v7, v9 int32
	var c, file, moveSet, pairSet, set, world uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyBitSetBytes, bodySimCapacity, bodyStateCapacity, c, contactSimCapacity, file, i, i1, islandSimCapacity, jointSimCapacity, moveSet, pairSet, set, solverSetCapacity, world, v1, v11, v13, v15, v17, v19, v21, v23, v25, v27, v29, v3, v31, v33, v35, v37, v39, v41, v43, v47, v5, v7, v9
	file = fopen(tls, __ccgo_ts+272, __ccgo_ts+289)
	if file == uintptrFromInt32(0) {
		return
	}
	world = b2GetWorldFromId(tls, worldId)
	// id pools
	fprintf(tls, file, __ccgo_ts+291, 0)
	v3 = int32FromUint64(uint64FromInt32((*b2IntArray)(unsafe.Pointer(world+1000)).Capacity) * uint64FromInt64(4))
	goto _4
_4:
	v1 = v3
	goto _2
_2:
	fprintf(tls, file, __ccgo_ts+301, vaList(bp+8, v1))
	v7 = int32FromUint64(uint64FromInt32((*b2IntArray)(unsafe.Pointer(world+1040)).Capacity) * uint64FromInt64(4))
	goto _8
_8:
	v5 = v7
	goto _6
_6:
	fprintf(tls, file, __ccgo_ts+315, vaList(bp+8, v5))
	v11 = int32FromUint64(uint64FromInt32((*b2IntArray)(unsafe.Pointer(world+1080)).Capacity) * uint64FromInt64(4))
	goto _12
_12:
	v9 = v11
	goto _10
_10:
	fprintf(tls, file, __ccgo_ts+335, vaList(bp+8, v9))
	v15 = int32FromUint64(uint64FromInt32((*b2IntArray)(unsafe.Pointer(world+1120)).Capacity) * uint64FromInt64(4))
	goto _16
_16:
	v13 = v15
	goto _14
_14:
	fprintf(tls, file, __ccgo_ts+350, vaList(bp+8, v13))
	v19 = int32FromUint64(uint64FromInt32((*b2IntArray)(unsafe.Pointer(world+1160)).Capacity) * uint64FromInt64(4))
	goto _20
_20:
	v17 = v19
	goto _18
_18:
	fprintf(tls, file, __ccgo_ts+367, vaList(bp+8, v17))
	v23 = int32FromUint64(uint64FromInt32((*b2IntArray)(unsafe.Pointer(world+1200)).Capacity) * uint64FromInt64(4))
	goto _24
_24:
	v21 = v23
	goto _22
_22:
	fprintf(tls, file, __ccgo_ts+383, vaList(bp+8, v21))
	v27 = int32FromUint64(uint64FromInt32((*b2IntArray)(unsafe.Pointer(world+1224)).Capacity) * uint64FromInt64(4))
	goto _28
_28:
	v25 = v27
	goto _26
_26:
	fprintf(tls, file, __ccgo_ts+398, vaList(bp+8, v25))
	fprintf(tls, file, __ccgo_ts+413, 0)
	// world arrays
	fprintf(tls, file, __ccgo_ts+415, 0)
	v29 = int32FromUint64(uint64FromInt32((*b2BodyArray)(unsafe.Pointer(world+1024)).Capacity) * uint64FromInt64(128))
	goto _30
_30:
	fprintf(tls, file, __ccgo_ts+429, vaList(bp+8, v29))
	v31 = int32FromUint64(uint64FromInt32((*b2SolverSetArray)(unsafe.Pointer(world+1064)).Capacity) * uint64FromInt64(88))
	goto _32
_32:
	fprintf(tls, file, __ccgo_ts+441, vaList(bp+8, v31))
	v33 = int32FromUint64(uint64FromInt32((*b2JointArray)(unsafe.Pointer(world+1104)).Capacity) * uint64FromInt64(72))
	goto _34
_34:
	fprintf(tls, file, __ccgo_ts+458, vaList(bp+8, v33))
	v35 = int32FromUint64(uint64FromInt32((*b2ContactArray)(unsafe.Pointer(world+1144)).Capacity) * uint64FromInt64(68))
	goto _36
_36:
	fprintf(tls, file, __ccgo_ts+470, vaList(bp+8, v35))
	v37 = int32FromUint64(uint64FromInt32((*b2IslandArray)(unsafe.Pointer(world+1184)).Capacity) * uint64FromInt64(56))
	goto _38
_38:
	fprintf(tls, file, __ccgo_ts+484, vaList(bp+8, v37))
	v39 = int32FromUint64(uint64FromInt32((*b2ShapeArray)(unsafe.Pointer(world+1248)).Capacity) * uint64FromInt64(288))
	goto _40
_40:
	fprintf(tls, file, __ccgo_ts+497, vaList(bp+8, v39))
	v41 = int32FromUint64(uint64FromInt32((*b2ChainShapeArray)(unsafe.Pointer(world+1264)).Capacity) * uint64FromInt64(48))
	goto _42
_42:
	fprintf(tls, file, __ccgo_ts+509, vaList(bp+8, v41))
	fprintf(tls, file, __ccgo_ts+413, 0)
	// broad-phase
	fprintf(tls, file, __ccgo_ts+521, 0)
	fprintf(tls, file, __ccgo_ts+534, vaList(bp+8, b2DynamicTree_GetByteCount(tls, world+40+uintptr(b2_staticBody)*72)))
	fprintf(tls, file, __ccgo_ts+551, vaList(bp+8, b2DynamicTree_GetByteCount(tls, world+40+uintptr(b2_kinematicBody)*72)))
	fprintf(tls, file, __ccgo_ts+571, vaList(bp+8, b2DynamicTree_GetByteCount(tls, world+40+uintptr(b2_dynamicBody)*72)))
	moveSet = world + 40 + 216
	fprintf(tls, file, __ccgo_ts+589, vaList(bp+8, b2GetHashSetBytes(tls, moveSet), (*b2HashSet)(unsafe.Pointer(moveSet)).Count, (*b2HashSet)(unsafe.Pointer(moveSet)).Capacity))
	v43 = int32FromUint64(uint64FromInt32((*b2IntArray)(unsafe.Pointer(world+40+232)).Capacity) * uint64FromInt64(4))
	goto _44
_44:
	fprintf(tls, file, __ccgo_ts+611, vaList(bp+8, v43))
	pairSet = world + 40 + 272
	fprintf(tls, file, __ccgo_ts+626, vaList(bp+8, b2GetHashSetBytes(tls, pairSet), (*b2HashSet)(unsafe.Pointer(pairSet)).Count, (*b2HashSet)(unsafe.Pointer(pairSet)).Capacity))
	fprintf(tls, file, __ccgo_ts+413, 0)
	// solver sets
	bodySimCapacity = 0
	bodyStateCapacity = 0
	jointSimCapacity = 0
	contactSimCapacity = 0
	islandSimCapacity = 0
	solverSetCapacity = (*b2World)(unsafe.Pointer(world)).SolverSets.Count
	i = 0
	for {
		if !(i < solverSetCapacity) {
			break
		}
		set = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(i)*88
		if (*b2SolverSet)(unsafe.Pointer(set)).SetIndex == -int32(1) {
			goto _45
		}
		bodySimCapacity += (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Capacity
		bodyStateCapacity += (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Capacity
		jointSimCapacity += (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Capacity
		contactSimCapacity += (*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Capacity
		islandSimCapacity += (*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Capacity
		goto _45
	_45:
		;
		i++
	}
	fprintf(tls, file, __ccgo_ts+648, 0)
	fprintf(tls, file, __ccgo_ts+661, vaList(bp+8, bodySimCapacity*int32FromInt64(100)))
	fprintf(tls, file, __ccgo_ts+675, vaList(bp+8, bodyStateCapacity*int32FromInt64(32)))
	fprintf(tls, file, __ccgo_ts+691, vaList(bp+8, jointSimCapacity*int32FromInt64(196)))
	fprintf(tls, file, __ccgo_ts+706, vaList(bp+8, contactSimCapacity*int32FromInt64(176)))
	fprintf(tls, file, __ccgo_ts+723, vaList(bp+8, islandSimCapacity*int32FromInt64(4)))
	fprintf(tls, file, __ccgo_ts+413, 0)
	// constraint graph
	bodyBitSetBytes = 0
	contactSimCapacity = 0
	jointSimCapacity = 0
	i1 = 0
	for {
		if !(i1 < int32(B2_GRAPH_COLOR_COUNT)) {
			break
		}
		c = world + 328 + uintptr(i1)*56
		v47 = int32FromUint64(uint64((*b2BitSet)(unsafe.Pointer(c)).BlockCapacity) * uint64(8))
		goto _48
	_48:
		bodyBitSetBytes += v47
		contactSimCapacity += (*b2GraphColor)(unsafe.Pointer(c)).ContactSims.Capacity
		jointSimCapacity += (*b2GraphColor)(unsafe.Pointer(c)).JointSims.Capacity
		goto _46
	_46:
		;
		i1++
	}
	fprintf(tls, file, __ccgo_ts+739, 0)
	fprintf(tls, file, __ccgo_ts+757, vaList(bp+8, bodyBitSetBytes))
	fprintf(tls, file, __ccgo_ts+691, vaList(bp+8, jointSimCapacity*int32FromInt64(196)))
	fprintf(tls, file, __ccgo_ts+706, vaList(bp+8, contactSimCapacity*int32FromInt64(176)))
	fprintf(tls, file, __ccgo_ts+413, 0)
	// stack allocator
	fprintf(tls, file, __ccgo_ts+776, vaList(bp+8, (*b2World)(unsafe.Pointer(world)).Arena.Capacity))
	// chain shapes
	// todo
	fclose(tls, file)
}

func b2World_OverlapAABB(tls *_Stack, worldId WorldId, aabb AABB, filter QueryFilter, fcn uintptr, context uintptr) (r TreeStats) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var i int32
	var treeResult, treeStats TreeStats
	var world uintptr
	var _ /* worldContext at bp+0 */ WorldQueryContext
	_, _, _, _ = i, treeResult, treeStats, world
	treeStats = TreeStats{}
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return treeStats
	}
	*(*WorldQueryContext)(unsafe.Pointer(bp)) = WorldQueryContext{
		World:       world,
		Fcn:         fcn,
		Filter:      filter,
		UserContext: context,
	}
	i = 0
	for {
		if !(i < int32(b2_bodyTypeCount)) {
			break
		}
		treeResult = b2DynamicTree_Query(tls, world+40+uintptr(i)*72, aabb, filter.MaskBits, __ccgo_fp(TreeQueryCallback), bp)
		treeStats.NodeVisits += treeResult.NodeVisits
		treeStats.LeafVisits += treeResult.LeafVisits
		goto _1
	_1:
		;
		i++
	}
	return treeStats
}

func b2World_OverlapShape(tls *_Stack, worldId WorldId, proxy uintptr, filter QueryFilter, fcn uintptr, context uintptr) (r1 TreeStats) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var a6, aabb, v40 AABB
	var c, c1, r, v16, v18, v19, v30, v32, v33, v34, v36, v37, v38, v4, v5 Vec2
	var i, i1 int32
	var treeResult, treeStats TreeStats
	var world, v1 uintptr
	var v10, v11, v12, v13, v15, v2, v20, v21, v22, v24, v25, v26, v27, v29, v6, v7, v8 float32
	var _ /* worldContext at bp+0 */ WorldOverlapContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a6, aabb, c, c1, i, i1, r, treeResult, treeStats, world, v1, v10, v11, v12, v13, v15, v16, v18, v19, v2, v20, v21, v22, v24, v25, v26, v27, v29, v30, v32, v33, v34, v36, v37, v38, v4, v40, v5, v6, v7, v8
	treeStats = TreeStats{}
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return treeStats
	}
	v1 = proxy
	v2 = (*ShapeProxy)(unsafe.Pointer(proxy)).Radius
	a6 = AABB{
		LowerBound: *(*Vec2)(unsafe.Pointer(v1)),
		UpperBound: *(*Vec2)(unsafe.Pointer(v1)),
	}
	i = int32(1)
	for {
		if !(i < (*ShapeProxy)(unsafe.Pointer(proxy)).Count) {
			break
		}
		v4 = a6.LowerBound
		v5 = *(*Vec2)(unsafe.Pointer(v1 + uintptr(i)*8))
		v6 = v4.X
		v7 = v5.X
		if v6 < v7 {
			v10 = v6
		} else {
			v10 = v7
		}
		v8 = v10
		goto _9
	_9:
		c.X = v8
		v11 = v4.Y
		v12 = v5.Y
		if v11 < v12 {
			v15 = v11
		} else {
			v15 = v12
		}
		v13 = v15
		goto _14
	_14:
		c.Y = v13
		v16 = c
		goto _17
	_17:
		a6.LowerBound = v16
		v18 = a6.UpperBound
		v19 = *(*Vec2)(unsafe.Pointer(v1 + uintptr(i)*8))
		v20 = v18.X
		v21 = v19.X
		if v20 > v21 {
			v24 = v20
		} else {
			v24 = v21
		}
		v22 = v24
		goto _23
	_23:
		c1.X = v22
		v25 = v18.Y
		v26 = v19.Y
		if v25 > v26 {
			v29 = v25
		} else {
			v29 = v26
		}
		v27 = v29
		goto _28
	_28:
		c1.Y = v27
		v30 = c1
		goto _31
	_31:
		a6.UpperBound = v30
		goto _3
	_3:
		;
		i++
	}
	r = Vec2{
		X: v2,
		Y: v2,
	}
	v32 = a6.LowerBound
	v33 = r
	v34 = Vec2{
		X: v32.X - v33.X,
		Y: v32.Y - v33.Y,
	}
	goto _35
_35:
	a6.LowerBound = v34
	v36 = a6.UpperBound
	v37 = r
	v38 = Vec2{
		X: v36.X + v37.X,
		Y: v36.Y + v37.Y,
	}
	goto _39
_39:
	a6.UpperBound = v38
	v40 = a6
	goto _41
_41:
	aabb = v40
	*(*WorldOverlapContext)(unsafe.Pointer(bp)) = WorldOverlapContext{
		World:       world,
		Fcn:         fcn,
		Filter:      filter,
		Proxy:       proxy,
		UserContext: context,
	}
	i1 = 0
	for {
		if !(i1 < int32(b2_bodyTypeCount)) {
			break
		}
		treeResult = b2DynamicTree_Query(tls, world+40+uintptr(i1)*72, aabb, filter.MaskBits, __ccgo_fp(TreeOverlapCallback), bp)
		treeStats.NodeVisits += treeResult.NodeVisits
		treeStats.LeafVisits += treeResult.LeafVisits
		goto _42
	_42:
		;
		i1++
	}
	return treeStats
}

func b2World_CastRay(tls *_Stack, worldId WorldId, origin Vec2, translation Vec2, filter QueryFilter, fcn uintptr, context uintptr) (r TreeStats) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var i int32
	var treeResult, treeStats TreeStats
	var world uintptr
	var _ /* input at bp+0 */ RayCastInput
	var _ /* worldContext at bp+24 */ WorldRayCastContext
	_, _, _, _ = i, treeResult, treeStats, world
	treeStats = TreeStats{}
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return treeStats
	}
	*(*RayCastInput)(unsafe.Pointer(bp)) = RayCastInput{
		Origin:      origin,
		Translation: translation,
		MaxFraction: float32FromFloat32(1),
	}
	*(*WorldRayCastContext)(unsafe.Pointer(bp + 24)) = WorldRayCastContext{
		World:       world,
		Fcn:         fcn,
		Filter:      filter,
		Fraction:    float32FromFloat32(1),
		UserContext: context,
	}
	i = 0
	for {
		if !(i < int32(b2_bodyTypeCount)) {
			break
		}
		treeResult = b2DynamicTree_RayCast(tls, world+40+uintptr(i)*72, bp, filter.MaskBits, __ccgo_fp(RayCastCallback), bp+24)
		treeStats.NodeVisits += treeResult.NodeVisits
		treeStats.LeafVisits += treeResult.LeafVisits
		if (*(*WorldRayCastContext)(unsafe.Pointer(bp + 24))).Fraction == float32FromFloat32(0) {
			return treeStats
		}
		(*(*RayCastInput)(unsafe.Pointer(bp))).MaxFraction = (*(*WorldRayCastContext)(unsafe.Pointer(bp + 24))).Fraction
		goto _1
	_1:
		;
		i++
	}
	return treeStats
}

func b2World_CastRayClosest(tls *_Stack, worldId WorldId, origin Vec2, translation Vec2, filter QueryFilter) (r RayResult) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var i int32
	var treeResult TreeStats
	var world uintptr
	var _ /* input at bp+40 */ RayCastInput
	var _ /* result at bp+0 */ RayResult
	var _ /* worldContext at bp+64 */ WorldRayCastContext
	_, _, _ = i, treeResult, world
	*(*RayResult)(unsafe.Pointer(bp)) = RayResult{}
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return *(*RayResult)(unsafe.Pointer(bp))
	}
	*(*RayCastInput)(unsafe.Pointer(bp + 40)) = RayCastInput{
		Origin:      origin,
		Translation: translation,
		MaxFraction: float32FromFloat32(1),
	}
	*(*WorldRayCastContext)(unsafe.Pointer(bp + 64)) = WorldRayCastContext{
		World:       world,
		Fcn:         __ccgo_fp(b2RayCastClosestFcn),
		Filter:      filter,
		Fraction:    float32FromFloat32(1),
		UserContext: bp,
	}
	i = 0
	for {
		if !(i < int32(b2_bodyTypeCount)) {
			break
		}
		treeResult = b2DynamicTree_RayCast(tls, world+40+uintptr(i)*72, bp+40, filter.MaskBits, __ccgo_fp(RayCastCallback), bp+64)
		(*(*RayResult)(unsafe.Pointer(bp))).NodeVisits += treeResult.NodeVisits
		(*(*RayResult)(unsafe.Pointer(bp))).LeafVisits += treeResult.LeafVisits
		if (*(*WorldRayCastContext)(unsafe.Pointer(bp + 64))).Fraction == float32FromFloat32(0) {
			return *(*RayResult)(unsafe.Pointer(bp))
		}
		(*(*RayCastInput)(unsafe.Pointer(bp + 40))).MaxFraction = (*(*WorldRayCastContext)(unsafe.Pointer(bp + 64))).Fraction
		goto _1
	_1:
		;
		i++
	}
	return *(*RayResult)(unsafe.Pointer(bp))
}

func b2World_CastShape(tls *_Stack, worldId WorldId, proxy uintptr, translation Vec2, filter QueryFilter, fcn uintptr, context uintptr) (r TreeStats) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var i int32
	var treeResult, treeStats TreeStats
	var world uintptr
	var _ /* input at bp+0 */ ShapeCastInput
	var _ /* worldContext at bp+88 */ WorldRayCastContext
	_, _, _, _ = i, treeResult, treeStats, world
	treeStats = TreeStats{}
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return treeStats
	}
	*(*ShapeCastInput)(unsafe.Pointer(bp)) = ShapeCastInput{}
	(*(*ShapeCastInput)(unsafe.Pointer(bp))).Proxy = *(*ShapeProxy)(unsafe.Pointer(proxy))
	(*(*ShapeCastInput)(unsafe.Pointer(bp))).Translation = translation
	(*(*ShapeCastInput)(unsafe.Pointer(bp))).MaxFraction = float32FromFloat32(1)
	*(*WorldRayCastContext)(unsafe.Pointer(bp + 88)) = WorldRayCastContext{
		World:       world,
		Fcn:         fcn,
		Filter:      filter,
		Fraction:    float32FromFloat32(1),
		UserContext: context,
	}
	i = 0
	for {
		if !(i < int32(b2_bodyTypeCount)) {
			break
		}
		treeResult = b2DynamicTree_ShapeCast(tls, world+40+uintptr(i)*72, bp, filter.MaskBits, __ccgo_fp(ShapeCastCallback), bp+88)
		treeStats.NodeVisits += treeResult.NodeVisits
		treeStats.LeafVisits += treeResult.LeafVisits
		if (*(*WorldRayCastContext)(unsafe.Pointer(bp + 88))).Fraction == float32FromFloat32(0) {
			return treeStats
		}
		(*(*ShapeCastInput)(unsafe.Pointer(bp))).MaxFraction = (*(*WorldRayCastContext)(unsafe.Pointer(bp + 88))).Fraction
		goto _1
	_1:
		;
		i++
	}
	return treeStats
}

func b2World_CastMover(tls *_Stack, worldId WorldId, mover uintptr, translation Vec2, filter QueryFilter) (r float32) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var i int32
	var world uintptr
	var _ /* input at bp+0 */ ShapeCastInput
	var _ /* worldContext at bp+88 */ WorldMoverCastContext
	_, _ = i, world
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return float32FromFloat32(1)
	}
	*(*ShapeCastInput)(unsafe.Pointer(bp)) = ShapeCastInput{}
	*(*Vec2)(unsafe.Pointer(bp)) = (*Capsule)(unsafe.Pointer(mover)).Center1
	*(*Vec2)(unsafe.Pointer(bp + 1*8)) = (*Capsule)(unsafe.Pointer(mover)).Center2
	(*(*ShapeCastInput)(unsafe.Pointer(bp))).Proxy.Count = int32(2)
	(*(*ShapeCastInput)(unsafe.Pointer(bp))).Proxy.Radius = (*Capsule)(unsafe.Pointer(mover)).Radius
	(*(*ShapeCastInput)(unsafe.Pointer(bp))).Translation = translation
	(*(*ShapeCastInput)(unsafe.Pointer(bp))).MaxFraction = float32FromFloat32(1)
	(*(*ShapeCastInput)(unsafe.Pointer(bp))).CanEncroach = uint8(true1)
	*(*WorldMoverCastContext)(unsafe.Pointer(bp + 88)) = WorldMoverCastContext{
		World:    world,
		Filter:   filter,
		Fraction: float32FromFloat32(1),
	}
	i = 0
	for {
		if !(i < int32(b2_bodyTypeCount)) {
			break
		}
		b2DynamicTree_ShapeCast(tls, world+40+uintptr(i)*72, bp, filter.MaskBits, __ccgo_fp(MoverCastCallback), bp+88)
		if (*(*WorldMoverCastContext)(unsafe.Pointer(bp + 88))).Fraction == float32FromFloat32(0) {
			return float32FromFloat32(0)
		}
		(*(*ShapeCastInput)(unsafe.Pointer(bp))).MaxFraction = (*(*WorldMoverCastContext)(unsafe.Pointer(bp + 88))).Fraction
		goto _1
	_1:
		;
		i++
	}
	return (*(*WorldMoverCastContext)(unsafe.Pointer(bp + 88))).Fraction
}

// C documentation
//
//	// It is tempting to use a shape proxy for the mover, but this makes handling deep overlap difficult and the generality may
//	// not be worth it.
func b2World_CollideMover(tls *_Stack, worldId WorldId, mover uintptr, filter QueryFilter, fcn uintptr, context uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var aabb AABB
	var c, c1, r, v1, v13, v15, v16, v17, v19, v2, v20, v31, v33, v34, v35 Vec2
	var i int32
	var world uintptr
	var v10, v12, v21, v22, v23, v25, v26, v27, v28, v3, v30, v4, v5, v7, v8, v9 float32
	var _ /* worldContext at bp+0 */ WorldMoverContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, c, c1, i, r, world, v1, v10, v12, v13, v15, v16, v17, v19, v2, v20, v21, v22, v23, v25, v26, v27, v28, v3, v30, v31, v33, v34, v35, v4, v5, v7, v8, v9
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	r = Vec2{
		X: (*Capsule)(unsafe.Pointer(mover)).Radius,
		Y: (*Capsule)(unsafe.Pointer(mover)).Radius,
	}
	v1 = (*Capsule)(unsafe.Pointer(mover)).Center1
	v2 = (*Capsule)(unsafe.Pointer(mover)).Center2
	v3 = v1.X
	v4 = v2.X
	if v3 < v4 {
		v7 = v3
	} else {
		v7 = v4
	}
	v5 = v7
	goto _6
_6:
	c.X = v5
	v8 = v1.Y
	v9 = v2.Y
	if v8 < v9 {
		v12 = v8
	} else {
		v12 = v9
	}
	v10 = v12
	goto _11
_11:
	c.Y = v10
	v13 = c
	goto _14
_14:
	v15 = v13
	v16 = r
	v17 = Vec2{
		X: v15.X - v16.X,
		Y: v15.Y - v16.Y,
	}
	goto _18
_18:
	aabb.LowerBound = v17
	v19 = (*Capsule)(unsafe.Pointer(mover)).Center1
	v20 = (*Capsule)(unsafe.Pointer(mover)).Center2
	v21 = v19.X
	v22 = v20.X
	if v21 > v22 {
		v25 = v21
	} else {
		v25 = v22
	}
	v23 = v25
	goto _24
_24:
	c1.X = v23
	v26 = v19.Y
	v27 = v20.Y
	if v26 > v27 {
		v30 = v26
	} else {
		v30 = v27
	}
	v28 = v30
	goto _29
_29:
	c1.Y = v28
	v31 = c1
	goto _32
_32:
	v33 = v31
	v34 = r
	v35 = Vec2{
		X: v33.X + v34.X,
		Y: v33.Y + v34.Y,
	}
	goto _36
_36:
	aabb.UpperBound = v35
	*(*WorldMoverContext)(unsafe.Pointer(bp)) = WorldMoverContext{
		World:       world,
		Fcn:         fcn,
		Filter:      filter,
		Mover:       *(*Capsule)(unsafe.Pointer(mover)),
		UserContext: context,
	}
	i = 0
	for {
		if !(i < int32(b2_bodyTypeCount)) {
			break
		}
		b2DynamicTree_Query(tls, world+40+uintptr(i)*72, aabb, filter.MaskBits, __ccgo_fp(TreeCollideCallback), bp)
		goto _37
	_37:
		;
		i++
	}
}

func b2World_SetCustomFilterCallback(tls *_Stack, worldId WorldId, fcn uintptr, context uintptr) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	(*b2World)(unsafe.Pointer(world)).CustomFilterFcn = fcn
	(*b2World)(unsafe.Pointer(world)).CustomFilterContext = context
}

func b2World_SetPreSolveCallback(tls *_Stack, worldId WorldId, fcn uintptr, context uintptr) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	(*b2World)(unsafe.Pointer(world)).PreSolveFcn = fcn
	(*b2World)(unsafe.Pointer(world)).PreSolveContext = context
}

func b2World_SetGravity(tls *_Stack, worldId WorldId, gravity Vec2) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	(*b2World)(unsafe.Pointer(world)).Gravity = gravity
}

func b2World_GetGravity(tls *_Stack, worldId WorldId) (r Vec2) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).Gravity
}

func b2World_Explode(tls *_Stack, worldId WorldId, explosionDef uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var aabb AABB
	var falloff, impulsePerLength, radius float32
	var maskBits uint64_t
	var position Vec2
	var world uintptr
	var _ /* explosionContext at bp+0 */ ExplosionContext
	_, _, _, _, _, _, _ = aabb, falloff, impulsePerLength, maskBits, position, radius, world
	maskBits = (*ExplosionDef)(unsafe.Pointer(explosionDef)).MaskBits
	position = (*ExplosionDef)(unsafe.Pointer(explosionDef)).Position
	radius = (*ExplosionDef)(unsafe.Pointer(explosionDef)).Radius
	falloff = (*ExplosionDef)(unsafe.Pointer(explosionDef)).Falloff
	impulsePerLength = (*ExplosionDef)(unsafe.Pointer(explosionDef)).ImpulsePerLength
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	*(*ExplosionContext)(unsafe.Pointer(bp)) = ExplosionContext{
		World:            world,
		Position:         position,
		Radius:           radius,
		Falloff:          falloff,
		ImpulsePerLength: impulsePerLength,
	}
	aabb.LowerBound.X = position.X - (radius + falloff)
	aabb.LowerBound.Y = position.Y - (radius + falloff)
	aabb.UpperBound.X = position.X + (radius + falloff)
	aabb.UpperBound.Y = position.Y + (radius + falloff)
	b2DynamicTree_Query(tls, world+40+uintptr(b2_dynamicBody)*72, aabb, maskBits, __ccgo_fp(ExplosionCallback), bp)
}

func b2World_RebuildStaticTree(tls *_Stack, worldId WorldId) {
	var staticTree, world uintptr
	_, _ = staticTree, world
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	staticTree = world + 40 + uintptr(b2_staticBody)*72
	b2DynamicTree_Rebuild(tls, staticTree, uint8(true1))
}

func b2World_EnableSpeculative(tls *_Stack, worldId WorldId, flag uint8) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	(*b2World)(unsafe.Pointer(world)).EnableSpeculative = flag
}
