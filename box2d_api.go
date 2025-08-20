package main

import (
	"unsafe"
	"runtime"
)

type Version = b2Version
type Vec2 = b2Vec2
type CosSin = b2CosSin
type Rot = b2Rot
type Transform = b2Transform
type Mat22 = b2Mat22
type AABB = b2AABB
type Plane = b2Plane
type SimplexCache = b2SimplexCache
type Hull = b2Hull
type RayCastInput = b2RayCastInput
type ShapeProxy = b2ShapeProxy
type ShapeCastInput = b2ShapeCastInput
type CastOutput = b2CastOutput
type MassData = b2MassData
type Circle = b2Circle
type Capsule = b2Capsule
type Polygon = b2Polygon
type Segment = b2Segment
type ChainSegment = b2ChainSegment
type SegmentDistanceResult = b2SegmentDistanceResult
type DistanceInput = b2DistanceInput
type DistanceOutput = b2DistanceOutput
type SimplexVertex = b2SimplexVertex
type Simplex = b2Simplex
type ShapeCastPairInput = b2ShapeCastPairInput
type Sweep = b2Sweep
type TOIInput = b2TOIInput
type TOIState = b2TOIState
type TOIOutput = b2TOIOutput
type ManifoldPoint = b2ManifoldPoint
type Manifold = b2Manifold
type TreeStats = b2TreeStats
type PlaneResult = b2PlaneResult
type CollisionPlane = b2CollisionPlane
type PlaneSolverResult = b2PlaneSolverResult
type WorldId = b2WorldId
type BodyId = b2BodyId
type ShapeId = b2ShapeId
type ChainId = b2ChainId
type JointId = b2JointId
type RayResult = b2RayResult
type WorldDef = b2WorldDef
type BodyType = b2BodyType
type BodyDef = b2BodyDef
type Filter = b2Filter
type QueryFilter = b2QueryFilter
type ShapeType = b2ShapeType
type SurfaceMaterial = b2SurfaceMaterial
type ShapeDef = b2ShapeDef
type ChainDef = b2ChainDef
type Profile = b2Profile
type Counters = b2Counters
type JointType = b2JointType
type DistanceJointDef = b2DistanceJointDef
type MotorJointDef = b2MotorJointDef
type MouseJointDef = b2MouseJointDef
type FilterJointDef = b2FilterJointDef
type PrismaticJointDef = b2PrismaticJointDef
type RevoluteJointDef = b2RevoluteJointDef
type WeldJointDef = b2WeldJointDef
type WheelJointDef = b2WheelJointDef
type ExplosionDef = b2ExplosionDef
type SensorBeginTouchEvent = b2SensorBeginTouchEvent
type SensorEndTouchEvent = b2SensorEndTouchEvent
type ContactBeginTouchEvent = b2ContactBeginTouchEvent
type ContactEndTouchEvent = b2ContactEndTouchEvent
type ContactHitEvent = b2ContactHitEvent
type BodyMoveEvent = b2BodyMoveEvent
type ContactData = b2ContactData
type HexColor = b2HexColor
type DebugDraw = b2DebugDraw

func (b *Box2D) IsValidAABB(a1 AABB) (r uint8) {
	r = b2IsValidAABB(b.tls, a1)
	return
}
func (b *Box2D) CreateBody(worldId WorldId, def BodyDef) (r BodyId) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r = b2CreateBody(b.tls, worldId, defStack.Addr)
	return
}
func (b *Box2D) DestroyBody(bodyId BodyId) {
	b2DestroyBody(b.tls, bodyId)
}
func (b *Box2D) BodyGetContactCapacity(bodyId BodyId) (r int32) {
	r = b2Body_GetContactCapacity(b.tls, bodyId)
	return
}
func (b *Box2D) BodyGetContactData(bodyId BodyId, contactData *ContactData, capacity int32) (r int32) {
	defer runtime.KeepAlive(contactData)
	escapes(contactData)
	r = b2Body_GetContactData(b.tls, bodyId, uintptr(unsafe.Pointer(contactData)), capacity)
	return
}
func (b *Box2D) BodyComputeAABB(bodyId BodyId) (r AABB) {
	r = b2Body_ComputeAABB(b.tls, bodyId)
	return
}
func (b *Box2D) BodyGetPosition(bodyId BodyId) (r Vec2) {
	r = b2Body_GetPosition(b.tls, bodyId)
	return
}
func (b *Box2D) BodyGetRotation(bodyId BodyId) (r Rot) {
	r = b2Body_GetRotation(b.tls, bodyId)
	return
}
func (b *Box2D) BodyGetTransform(bodyId BodyId) (r Transform) {
	r = b2Body_GetTransform(b.tls, bodyId)
	return
}
func (b *Box2D) BodyGetLocalPoint(bodyId BodyId, worldPoint Vec2) (r Vec2) {
	r = b2Body_GetLocalPoint(b.tls, bodyId, worldPoint)
	return
}
func (b *Box2D) BodyGetWorldPoint(bodyId BodyId, localPoint Vec2) (r Vec2) {
	r = b2Body_GetWorldPoint(b.tls, bodyId, localPoint)
	return
}
func (b *Box2D) BodyGetLocalVector(bodyId BodyId, worldVector Vec2) (r Vec2) {
	r = b2Body_GetLocalVector(b.tls, bodyId, worldVector)
	return
}
func (b *Box2D) BodyGetWorldVector(bodyId BodyId, localVector Vec2) (r Vec2) {
	r = b2Body_GetWorldVector(b.tls, bodyId, localVector)
	return
}
func (b *Box2D) BodySetTransform(bodyId BodyId, position Vec2, rotation Rot) {
	b2Body_SetTransform(b.tls, bodyId, position, rotation)
}
func (b *Box2D) BodyGetLinearVelocity(bodyId BodyId) (r Vec2) {
	r = b2Body_GetLinearVelocity(b.tls, bodyId)
	return
}
func (b *Box2D) BodyGetAngularVelocity(bodyId BodyId) (r float32) {
	r = b2Body_GetAngularVelocity(b.tls, bodyId)
	return
}
func (b *Box2D) BodySetLinearVelocity(bodyId BodyId, linearVelocity Vec2) {
	b2Body_SetLinearVelocity(b.tls, bodyId, linearVelocity)
}
func (b *Box2D) BodySetAngularVelocity(bodyId BodyId, angularVelocity float32) {
	b2Body_SetAngularVelocity(b.tls, bodyId, angularVelocity)
}
func (b *Box2D) BodySetTargetTransform(bodyId BodyId, target Transform, timeStep float32) {
	b2Body_SetTargetTransform(b.tls, bodyId, target, timeStep)
}
func (b *Box2D) BodyGetLocalPointVelocity(bodyId BodyId, localPoint Vec2) (r1 Vec2) {
	r1 = b2Body_GetLocalPointVelocity(b.tls, bodyId, localPoint)
	return
}
func (b *Box2D) BodyGetWorldPointVelocity(bodyId BodyId, worldPoint Vec2) (r1 Vec2) {
	r1 = b2Body_GetWorldPointVelocity(b.tls, bodyId, worldPoint)
	return
}
func (b *Box2D) BodyApplyForce(bodyId BodyId, force Vec2, point Vec2, wake uint8) {
	b2Body_ApplyForce(b.tls, bodyId, force, point, wake)
}
func (b *Box2D) BodyApplyForceToCenter(bodyId BodyId, force Vec2, wake uint8) {
	b2Body_ApplyForceToCenter(b.tls, bodyId, force, wake)
}
func (b *Box2D) BodyApplyTorque(bodyId BodyId, torque float32, wake uint8) {
	b2Body_ApplyTorque(b.tls, bodyId, torque, wake)
}
func (b *Box2D) BodyApplyLinearImpulse(bodyId BodyId, impulse Vec2, point Vec2, wake uint8) {
	b2Body_ApplyLinearImpulse(b.tls, bodyId, impulse, point, wake)
}
func (b *Box2D) BodyApplyLinearImpulseToCenter(bodyId BodyId, impulse Vec2, wake uint8) {
	b2Body_ApplyLinearImpulseToCenter(b.tls, bodyId, impulse, wake)
}
func (b *Box2D) BodyApplyAngularImpulse(bodyId BodyId, impulse float32, wake uint8) {
	b2Body_ApplyAngularImpulse(b.tls, bodyId, impulse, wake)
}
func (b *Box2D) BodyGetType(bodyId BodyId) (r BodyType) {
	r = b2Body_GetType(b.tls, bodyId)
	return
}
func (b *Box2D) BodySetType(bodyId BodyId, type1 BodyType) {
	b2Body_SetType(b.tls, bodyId, type1)
}
func (b *Box2D) BodySetUserData(bodyId BodyId, userData uintptr) {
	b2Body_SetUserData(b.tls, bodyId, userData)
}
func (b *Box2D) BodyGetUserData(bodyId BodyId) (r uintptr) {
	r = b2Body_GetUserData(b.tls, bodyId)
	return
}
func (b *Box2D) BodyGetMass(bodyId BodyId) (r float32) {
	r = b2Body_GetMass(b.tls, bodyId)
	return
}
func (b *Box2D) BodyGetRotationalInertia(bodyId BodyId) (r float32) {
	r = b2Body_GetRotationalInertia(b.tls, bodyId)
	return
}
func (b *Box2D) BodyGetLocalCenterOfMass(bodyId BodyId) (r Vec2) {
	r = b2Body_GetLocalCenterOfMass(b.tls, bodyId)
	return
}
func (b *Box2D) BodyGetWorldCenterOfMass(bodyId BodyId) (r Vec2) {
	r = b2Body_GetWorldCenterOfMass(b.tls, bodyId)
	return
}
func (b *Box2D) BodySetMassData(bodyId BodyId, massData MassData) {
	b2Body_SetMassData(b.tls, bodyId, massData)
}
func (b *Box2D) BodyGetMassData(bodyId BodyId) (r MassData) {
	r = b2Body_GetMassData(b.tls, bodyId)
	return
}
func (b *Box2D) BodyApplyMassFromShapes(bodyId BodyId) {
	b2Body_ApplyMassFromShapes(b.tls, bodyId)
}
func (b *Box2D) BodySetLinearDamping(bodyId BodyId, linearDamping float32) {
	b2Body_SetLinearDamping(b.tls, bodyId, linearDamping)
}
func (b *Box2D) BodyGetLinearDamping(bodyId BodyId) (r float32) {
	r = b2Body_GetLinearDamping(b.tls, bodyId)
	return
}
func (b *Box2D) BodySetAngularDamping(bodyId BodyId, angularDamping float32) {
	b2Body_SetAngularDamping(b.tls, bodyId, angularDamping)
}
func (b *Box2D) BodyGetAngularDamping(bodyId BodyId) (r float32) {
	r = b2Body_GetAngularDamping(b.tls, bodyId)
	return
}
func (b *Box2D) BodySetGravityScale(bodyId BodyId, gravityScale float32) {
	b2Body_SetGravityScale(b.tls, bodyId, gravityScale)
}
func (b *Box2D) BodyGetGravityScale(bodyId BodyId) (r float32) {
	r = b2Body_GetGravityScale(b.tls, bodyId)
	return
}
func (b *Box2D) BodyIsAwake(bodyId BodyId) (r uint8) {
	r = b2Body_IsAwake(b.tls, bodyId)
	return
}
func (b *Box2D) BodySetAwake(bodyId BodyId, awake uint8) {
	b2Body_SetAwake(b.tls, bodyId, awake)
}
func (b *Box2D) BodyIsEnabled(bodyId BodyId) (r uint8) {
	r = b2Body_IsEnabled(b.tls, bodyId)
	return
}
func (b *Box2D) BodyIsSleepEnabled(bodyId BodyId) (r uint8) {
	r = b2Body_IsSleepEnabled(b.tls, bodyId)
	return
}
func (b *Box2D) BodySetSleepThreshold(bodyId BodyId, sleepThreshold float32) {
	b2Body_SetSleepThreshold(b.tls, bodyId, sleepThreshold)
}
func (b *Box2D) BodyGetSleepThreshold(bodyId BodyId) (r float32) {
	r = b2Body_GetSleepThreshold(b.tls, bodyId)
	return
}
func (b *Box2D) BodyEnableSleep(bodyId BodyId, enableSleep uint8) {
	b2Body_EnableSleep(b.tls, bodyId, enableSleep)
}
func (b *Box2D) BodyDisable(bodyId BodyId) {
	b2Body_Disable(b.tls, bodyId)
}
func (b *Box2D) BodyEnable(bodyId BodyId) {
	b2Body_Enable(b.tls, bodyId)
}
func (b *Box2D) BodySetFixedRotation(bodyId BodyId, flag uint8) {
	b2Body_SetFixedRotation(b.tls, bodyId, flag)
}
func (b *Box2D) BodyIsFixedRotation(bodyId BodyId) (r uint8) {
	r = b2Body_IsFixedRotation(b.tls, bodyId)
	return
}
func (b *Box2D) BodySetBullet(bodyId BodyId, flag uint8) {
	b2Body_SetBullet(b.tls, bodyId, flag)
}
func (b *Box2D) BodyIsBullet(bodyId BodyId) (r uint8) {
	r = b2Body_IsBullet(b.tls, bodyId)
	return
}
func (b *Box2D) BodyEnableContactEvents(bodyId BodyId, flag uint8) {
	b2Body_EnableContactEvents(b.tls, bodyId, flag)
}
func (b *Box2D) BodyEnableHitEvents(bodyId BodyId, flag uint8) {
	b2Body_EnableHitEvents(b.tls, bodyId, flag)
}
func (b *Box2D) BodyGetWorld(bodyId BodyId) (r WorldId) {
	r = b2Body_GetWorld(b.tls, bodyId)
	return
}
func (b *Box2D) BodyGetShapeCount(bodyId BodyId) (r int32) {
	r = b2Body_GetShapeCount(b.tls, bodyId)
	return
}
func (b *Box2D) BodyGetJointCount(bodyId BodyId) (r int32) {
	r = b2Body_GetJointCount(b.tls, bodyId)
	return
}
func (b *Box2D) SetLengthUnitsPerMeter(lengthUnits float32) {
	b2SetLengthUnitsPerMeter(b.tls, lengthUnits)
}
func (b *Box2D) GetLengthUnitsPerMeter() (r float32) {
	r = b2GetLengthUnitsPerMeter(b.tls)
	return
}
func (b *Box2D) GetVersion() (r Version) {
	r = b2GetVersion(b.tls)
	return
}
func (b *Box2D) GetByteCount() (r int32) {
	r = b2GetByteCount(b.tls)
	return
}
func (b *Box2D) GetSweepTransform(sweep Sweep, time float32) (r Transform) {
	sweepStack := copyToStack(b.tls, sweep)
	defer sweepStack.Free()
	r = b2GetSweepTransform(b.tls, sweepStack.Addr, time)
	return
}
func (b *Box2D) SegmentDistance(p1 Vec2, q1 Vec2, p2 Vec2, q2 Vec2) (r1 SegmentDistanceResult) {
	r1 = b2SegmentDistance(b.tls, p1, q1, p2, q2)
	return
}
func (b *Box2D) MakeProxy(points Vec2, count int32, radius float32) (r ShapeProxy) {
	pointsStack := copyToStack(b.tls, points)
	defer pointsStack.Free()
	r = b2MakeProxy(b.tls, pointsStack.Addr, count, radius)
	return
}
func (b *Box2D) MakeOffsetProxy(points Vec2, count int32, radius float32, position Vec2, rotation Rot) (r ShapeProxy) {
	pointsStack := copyToStack(b.tls, points)
	defer pointsStack.Free()
	r = b2MakeOffsetProxy(b.tls, pointsStack.Addr, count, radius, position, rotation)
	return
}
func (b *Box2D) ShapeDistance(input DistanceInput, cache *SimplexCache, simplexes *Simplex, simplexCapacity int32) (r DistanceOutput) {
	defer runtime.KeepAlive(simplexes)
	defer runtime.KeepAlive(cache)
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	escapes(cache)
	escapes(simplexes)
	r = b2ShapeDistance(b.tls, inputStack.Addr, uintptr(unsafe.Pointer(cache)), uintptr(unsafe.Pointer(simplexes)), simplexCapacity)
	return
}
func (b *Box2D) ShapeCast(input ShapeCastPairInput) (r CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	r = b2ShapeCast(b.tls, inputStack.Addr)
	return
}
func (b *Box2D) TimeOfImpact(input TOIInput) (r TOIOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	r = b2TimeOfImpact(b.tls, inputStack.Addr)
	return
}
func (b *Box2D) DistanceJointSetLength(jointId JointId, length float32) {
	b2DistanceJoint_SetLength(b.tls, jointId, length)
}
func (b *Box2D) DistanceJointGetLength(jointId JointId) (r float32) {
	r = b2DistanceJoint_GetLength(b.tls, jointId)
	return
}
func (b *Box2D) DistanceJointEnableLimit(jointId JointId, enableLimit uint8) {
	b2DistanceJoint_EnableLimit(b.tls, jointId, enableLimit)
}
func (b *Box2D) DistanceJointIsLimitEnabled(jointId JointId) (r uint8) {
	r = b2DistanceJoint_IsLimitEnabled(b.tls, jointId)
	return
}
func (b *Box2D) DistanceJointSetLengthRange(jointId JointId, minLength float32, maxLength float32) {
	b2DistanceJoint_SetLengthRange(b.tls, jointId, minLength, maxLength)
}
func (b *Box2D) DistanceJointGetMinLength(jointId JointId) (r float32) {
	r = b2DistanceJoint_GetMinLength(b.tls, jointId)
	return
}
func (b *Box2D) DistanceJointGetMaxLength(jointId JointId) (r float32) {
	r = b2DistanceJoint_GetMaxLength(b.tls, jointId)
	return
}
func (b *Box2D) DistanceJointGetCurrentLength(jointId JointId) (r float32) {
	r = b2DistanceJoint_GetCurrentLength(b.tls, jointId)
	return
}
func (b *Box2D) DistanceJointEnableSpring(jointId JointId, enableSpring uint8) {
	b2DistanceJoint_EnableSpring(b.tls, jointId, enableSpring)
}
func (b *Box2D) DistanceJointIsSpringEnabled(jointId JointId) (r uint8) {
	r = b2DistanceJoint_IsSpringEnabled(b.tls, jointId)
	return
}
func (b *Box2D) DistanceJointSetSpringHertz(jointId JointId, hertz float32) {
	b2DistanceJoint_SetSpringHertz(b.tls, jointId, hertz)
}
func (b *Box2D) DistanceJointSetSpringDampingRatio(jointId JointId, dampingRatio float32) {
	b2DistanceJoint_SetSpringDampingRatio(b.tls, jointId, dampingRatio)
}
func (b *Box2D) DistanceJointGetSpringHertz(jointId JointId) (r float32) {
	r = b2DistanceJoint_GetSpringHertz(b.tls, jointId)
	return
}
func (b *Box2D) DistanceJointGetSpringDampingRatio(jointId JointId) (r float32) {
	r = b2DistanceJoint_GetSpringDampingRatio(b.tls, jointId)
	return
}
func (b *Box2D) DistanceJointEnableMotor(jointId JointId, enableMotor uint8) {
	b2DistanceJoint_EnableMotor(b.tls, jointId, enableMotor)
}
func (b *Box2D) DistanceJointIsMotorEnabled(jointId JointId) (r uint8) {
	r = b2DistanceJoint_IsMotorEnabled(b.tls, jointId)
	return
}
func (b *Box2D) DistanceJointSetMotorSpeed(jointId JointId, motorSpeed float32) {
	b2DistanceJoint_SetMotorSpeed(b.tls, jointId, motorSpeed)
}
func (b *Box2D) DistanceJointGetMotorSpeed(jointId JointId) (r float32) {
	r = b2DistanceJoint_GetMotorSpeed(b.tls, jointId)
	return
}
func (b *Box2D) DistanceJointGetMotorForce(jointId JointId) (r float32) {
	r = b2DistanceJoint_GetMotorForce(b.tls, jointId)
	return
}
func (b *Box2D) DistanceJointSetMaxMotorForce(jointId JointId, force float32) {
	b2DistanceJoint_SetMaxMotorForce(b.tls, jointId, force)
}
func (b *Box2D) DistanceJointGetMaxMotorForce(jointId JointId) (r float32) {
	r = b2DistanceJoint_GetMaxMotorForce(b.tls, jointId)
	return
}
func (b *Box2D) IsValidRay(input RayCastInput) (r uint8) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	r = b2IsValidRay(b.tls, inputStack.Addr)
	return
}
func (b *Box2D) MakePolygon(hull Hull, radius float32) (r Polygon) {
	hullStack := copyToStack(b.tls, hull)
	defer hullStack.Free()
	r = b2MakePolygon(b.tls, hullStack.Addr, radius)
	return
}
func (b *Box2D) MakeOffsetPolygon(hull Hull, position Vec2, rotation Rot) (r Polygon) {
	hullStack := copyToStack(b.tls, hull)
	defer hullStack.Free()
	r = b2MakeOffsetPolygon(b.tls, hullStack.Addr, position, rotation)
	return
}
func (b *Box2D) MakeOffsetRoundedPolygon(hull Hull, position Vec2, rotation Rot, radius float32) (r Polygon) {
	hullStack := copyToStack(b.tls, hull)
	defer hullStack.Free()
	r = b2MakeOffsetRoundedPolygon(b.tls, hullStack.Addr, position, rotation, radius)
	return
}
func (b *Box2D) MakeSquare(halfWidth float32) (r Polygon) {
	r = b2MakeSquare(b.tls, halfWidth)
	return
}
func (b *Box2D) MakeBox(halfWidth float32, halfHeight float32) (r Polygon) {
	r = b2MakeBox(b.tls, halfWidth, halfHeight)
	return
}
func (b *Box2D) MakeRoundedBox(halfWidth float32, halfHeight float32, radius float32) (r Polygon) {
	r = b2MakeRoundedBox(b.tls, halfWidth, halfHeight, radius)
	return
}
func (b *Box2D) MakeOffsetBox(halfWidth float32, halfHeight float32, center Vec2, rotation Rot) (r Polygon) {
	r = b2MakeOffsetBox(b.tls, halfWidth, halfHeight, center, rotation)
	return
}
func (b *Box2D) MakeOffsetRoundedBox(halfWidth float32, halfHeight float32, center Vec2, rotation Rot, radius float32) (r Polygon) {
	r = b2MakeOffsetRoundedBox(b.tls, halfWidth, halfHeight, center, rotation, radius)
	return
}
func (b *Box2D) TransformPolygon(transform Transform, polygon Polygon) (r Polygon) {
	polygonStack := copyToStack(b.tls, polygon)
	defer polygonStack.Free()
	r = b2TransformPolygon(b.tls, transform, polygonStack.Addr)
	return
}
func (b *Box2D) ComputeCircleMass(shape Circle, density float32) (r MassData) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2ComputeCircleMass(b.tls, shapeStack.Addr, density)
	return
}
func (b *Box2D) ComputeCapsuleMass(shape Capsule, density float32) (r MassData) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2ComputeCapsuleMass(b.tls, shapeStack.Addr, density)
	return
}
func (b *Box2D) ComputePolygonMass(shape Polygon, density float32) (r1 MassData) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r1 = b2ComputePolygonMass(b.tls, shapeStack.Addr, density)
	return
}
func (b *Box2D) ComputeCircleAABB(shape Circle, xf Transform) (r1 AABB) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r1 = b2ComputeCircleAABB(b.tls, shapeStack.Addr, xf)
	return
}
func (b *Box2D) ComputeCapsuleAABB(shape Capsule, xf Transform) (r1 AABB) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r1 = b2ComputeCapsuleAABB(b.tls, shapeStack.Addr, xf)
	return
}
func (b *Box2D) ComputePolygonAABB(shape Polygon, xf Transform) (r1 AABB) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r1 = b2ComputePolygonAABB(b.tls, shapeStack.Addr, xf)
	return
}
func (b *Box2D) ComputeSegmentAABB(shape Segment, xf Transform) (r AABB) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2ComputeSegmentAABB(b.tls, shapeStack.Addr, xf)
	return
}
func (b *Box2D) PointInCircle(point Vec2, shape Circle) (r uint8) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2PointInCircle(b.tls, point, shapeStack.Addr)
	return
}
func (b *Box2D) PointInCapsule(point Vec2, shape Capsule) (r uint8) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2PointInCapsule(b.tls, point, shapeStack.Addr)
	return
}
func (b *Box2D) PointInPolygon(_point Vec2, shape Polygon) (r uint8) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2PointInPolygon(b.tls, _point, shapeStack.Addr)
	return
}
func (b *Box2D) RayCastCircle(input RayCastInput, shape Circle) (r1 CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r1 = b2RayCastCircle(b.tls, inputStack.Addr, shapeStack.Addr)
	return
}
func (b *Box2D) RayCastCapsule(input RayCastInput, shape Capsule) (r CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2RayCastCapsule(b.tls, inputStack.Addr, shapeStack.Addr)
	return
}
func (b *Box2D) RayCastSegment(input RayCastInput, shape Segment, oneSided uint8) (r CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2RayCastSegment(b.tls, inputStack.Addr, shapeStack.Addr, oneSided)
	return
}
func (b *Box2D) RayCastPolygon(input RayCastInput, shape Polygon) (r CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2RayCastPolygon(b.tls, inputStack.Addr, shapeStack.Addr)
	return
}
func (b *Box2D) ShapeCastCircle(input ShapeCastInput, shape Circle) (r CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2ShapeCastCircle(b.tls, inputStack.Addr, shapeStack.Addr)
	return
}
func (b *Box2D) ShapeCastCapsule(input ShapeCastInput, shape Capsule) (r CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2ShapeCastCapsule(b.tls, inputStack.Addr, shapeStack.Addr)
	return
}
func (b *Box2D) ShapeCastSegment(input ShapeCastInput, shape Segment) (r CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2ShapeCastSegment(b.tls, inputStack.Addr, shapeStack.Addr)
	return
}
func (b *Box2D) ShapeCastPolygon(input ShapeCastInput, shape Polygon) (r CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2ShapeCastPolygon(b.tls, inputStack.Addr, shapeStack.Addr)
	return
}
func (b *Box2D) ComputeHull(points Vec2, count int32) (r1 Hull) {
	pointsStack := copyToStack(b.tls, points)
	defer pointsStack.Free()
	r1 = b2ComputeHull(b.tls, pointsStack.Addr, count)
	return
}
func (b *Box2D) ValidateHull(hull Hull) (r uint8) {
	hullStack := copyToStack(b.tls, hull)
	defer hullStack.Free()
	r = b2ValidateHull(b.tls, hullStack.Addr)
	return
}
func (b *Box2D) DefaultDistanceJointDef() (r DistanceJointDef) {
	r = b2DefaultDistanceJointDef(b.tls)
	return
}
func (b *Box2D) DefaultMotorJointDef() (r MotorJointDef) {
	r = b2DefaultMotorJointDef(b.tls)
	return
}
func (b *Box2D) DefaultMouseJointDef() (r MouseJointDef) {
	r = b2DefaultMouseJointDef(b.tls)
	return
}
func (b *Box2D) DefaultFilterJointDef() (r FilterJointDef) {
	r = b2DefaultFilterJointDef(b.tls)
	return
}
func (b *Box2D) DefaultPrismaticJointDef() (r PrismaticJointDef) {
	r = b2DefaultPrismaticJointDef(b.tls)
	return
}
func (b *Box2D) DefaultRevoluteJointDef() (r RevoluteJointDef) {
	r = b2DefaultRevoluteJointDef(b.tls)
	return
}
func (b *Box2D) DefaultWeldJointDef() (r WeldJointDef) {
	r = b2DefaultWeldJointDef(b.tls)
	return
}
func (b *Box2D) DefaultWheelJointDef() (r WheelJointDef) {
	r = b2DefaultWheelJointDef(b.tls)
	return
}
func (b *Box2D) DefaultExplosionDef() (r ExplosionDef) {
	r = b2DefaultExplosionDef(b.tls)
	return
}
func (b *Box2D) CreateDistanceJoint(worldId WorldId, def DistanceJointDef) (r JointId) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r = b2CreateDistanceJoint(b.tls, worldId, defStack.Addr)
	return
}
func (b *Box2D) CreateMotorJoint(worldId WorldId, def MotorJointDef) (r JointId) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r = b2CreateMotorJoint(b.tls, worldId, defStack.Addr)
	return
}
func (b *Box2D) CreateMouseJoint(worldId WorldId, def MouseJointDef) (r JointId) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r = b2CreateMouseJoint(b.tls, worldId, defStack.Addr)
	return
}
func (b *Box2D) CreateFilterJoint(worldId WorldId, def FilterJointDef) (r JointId) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r = b2CreateFilterJoint(b.tls, worldId, defStack.Addr)
	return
}
func (b *Box2D) CreateRevoluteJoint(worldId WorldId, def RevoluteJointDef) (r JointId) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r = b2CreateRevoluteJoint(b.tls, worldId, defStack.Addr)
	return
}
func (b *Box2D) CreatePrismaticJoint(worldId WorldId, def PrismaticJointDef) (r JointId) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r = b2CreatePrismaticJoint(b.tls, worldId, defStack.Addr)
	return
}
func (b *Box2D) CreateWeldJoint(worldId WorldId, def WeldJointDef) (r JointId) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r = b2CreateWeldJoint(b.tls, worldId, defStack.Addr)
	return
}
func (b *Box2D) CreateWheelJoint(worldId WorldId, def WheelJointDef) (r JointId) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r = b2CreateWheelJoint(b.tls, worldId, defStack.Addr)
	return
}
func (b *Box2D) DestroyJoint(jointId JointId) {
	b2DestroyJoint(b.tls, jointId)
}
func (b *Box2D) JointGetType(jointId JointId) (r JointType) {
	r = b2Joint_GetType(b.tls, jointId)
	return
}
func (b *Box2D) JointGetBodyA(jointId JointId) (r BodyId) {
	r = b2Joint_GetBodyA(b.tls, jointId)
	return
}
func (b *Box2D) JointGetBodyB(jointId JointId) (r BodyId) {
	r = b2Joint_GetBodyB(b.tls, jointId)
	return
}
func (b *Box2D) JointGetWorld(jointId JointId) (r WorldId) {
	r = b2Joint_GetWorld(b.tls, jointId)
	return
}
func (b *Box2D) JointSetLocalAnchorA(jointId JointId, localAnchor Vec2) {
	b2Joint_SetLocalAnchorA(b.tls, jointId, localAnchor)
}
func (b *Box2D) JointGetLocalAnchorA(jointId JointId) (r Vec2) {
	r = b2Joint_GetLocalAnchorA(b.tls, jointId)
	return
}
func (b *Box2D) JointSetLocalAnchorB(jointId JointId, localAnchor Vec2) {
	b2Joint_SetLocalAnchorB(b.tls, jointId, localAnchor)
}
func (b *Box2D) JointGetLocalAnchorB(jointId JointId) (r Vec2) {
	r = b2Joint_GetLocalAnchorB(b.tls, jointId)
	return
}
func (b *Box2D) JointSetReferenceAngle(jointId JointId, angleInRadians float32) {
	b2Joint_SetReferenceAngle(b.tls, jointId, angleInRadians)
}
func (b *Box2D) JointGetReferenceAngle(jointId JointId) (r float32) {
	r = b2Joint_GetReferenceAngle(b.tls, jointId)
	return
}
func (b *Box2D) JointSetLocalAxisA(jointId JointId, localAxis Vec2) {
	b2Joint_SetLocalAxisA(b.tls, jointId, localAxis)
}
func (b *Box2D) JointGetLocalAxisA(jointId JointId) (r Vec2) {
	r = b2Joint_GetLocalAxisA(b.tls, jointId)
	return
}
func (b *Box2D) JointSetCollideConnected(jointId JointId, shouldCollide uint8) {
	b2Joint_SetCollideConnected(b.tls, jointId, shouldCollide)
}
func (b *Box2D) JointGetCollideConnected(jointId JointId) (r uint8) {
	r = b2Joint_GetCollideConnected(b.tls, jointId)
	return
}
func (b *Box2D) JointSetUserData(jointId JointId, userData uintptr) {
	b2Joint_SetUserData(b.tls, jointId, userData)
}
func (b *Box2D) JointGetUserData(jointId JointId) (r uintptr) {
	r = b2Joint_GetUserData(b.tls, jointId)
	return
}
func (b *Box2D) JointWakeBodies(jointId JointId) {
	b2Joint_WakeBodies(b.tls, jointId)
}
func (b *Box2D) JointGetConstraintForce(jointId JointId) (r Vec2) {
	r = b2Joint_GetConstraintForce(b.tls, jointId)
	return
}
func (b *Box2D) JointGetConstraintTorque(jointId JointId) (r float32) {
	r = b2Joint_GetConstraintTorque(b.tls, jointId)
	return
}
func (b *Box2D) JointGetLinearSeparation(jointId JointId) (r float32) {
	r = b2Joint_GetLinearSeparation(b.tls, jointId)
	return
}
func (b *Box2D) JointGetAngularSeparation(jointId JointId) (r float32) {
	r = b2Joint_GetAngularSeparation(b.tls, jointId)
	return
}
func (b *Box2D) JointGetConstraintTuning(jointId JointId, hertz uintptr, dampingRatio uintptr) {
	b2Joint_GetConstraintTuning(b.tls, jointId, hertz, dampingRatio)
}
func (b *Box2D) JointSetConstraintTuning(jointId JointId, hertz float32, dampingRatio float32) {
	b2Joint_SetConstraintTuning(b.tls, jointId, hertz, dampingRatio)
}
func (b *Box2D) CollideCircles(circleA Circle, xfA Transform, circleB Circle, xfB Transform) (r Manifold) {
	circleAStack := copyToStack(b.tls, circleA)
	defer circleAStack.Free()
	circleBStack := copyToStack(b.tls, circleB)
	defer circleBStack.Free()
	r = b2CollideCircles(b.tls, circleAStack.Addr, xfA, circleBStack.Addr, xfB)
	return
}
func (b *Box2D) CollideCapsuleAndCircle(capsuleA Capsule, xfA Transform, circleB Circle, xfB Transform) (r Manifold) {
	capsuleAStack := copyToStack(b.tls, capsuleA)
	defer capsuleAStack.Free()
	circleBStack := copyToStack(b.tls, circleB)
	defer circleBStack.Free()
	r = b2CollideCapsuleAndCircle(b.tls, capsuleAStack.Addr, xfA, circleBStack.Addr, xfB)
	return
}
func (b *Box2D) CollidePolygonAndCircle(polygonA Polygon, xfA Transform, circleB Circle, xfB Transform) (r Manifold) {
	polygonAStack := copyToStack(b.tls, polygonA)
	defer polygonAStack.Free()
	circleBStack := copyToStack(b.tls, circleB)
	defer circleBStack.Free()
	r = b2CollidePolygonAndCircle(b.tls, polygonAStack.Addr, xfA, circleBStack.Addr, xfB)
	return
}
func (b *Box2D) CollideCapsules(capsuleA Capsule, xfA Transform, capsuleB Capsule, xfB Transform) (r Manifold) {
	capsuleAStack := copyToStack(b.tls, capsuleA)
	defer capsuleAStack.Free()
	capsuleBStack := copyToStack(b.tls, capsuleB)
	defer capsuleBStack.Free()
	r = b2CollideCapsules(b.tls, capsuleAStack.Addr, xfA, capsuleBStack.Addr, xfB)
	return
}
func (b *Box2D) CollideSegmentAndCapsule(segmentA Segment, xfA Transform, capsuleB Capsule, xfB Transform) (r Manifold) {
	segmentAStack := copyToStack(b.tls, segmentA)
	defer segmentAStack.Free()
	capsuleBStack := copyToStack(b.tls, capsuleB)
	defer capsuleBStack.Free()
	r = b2CollideSegmentAndCapsule(b.tls, segmentAStack.Addr, xfA, capsuleBStack.Addr, xfB)
	return
}
func (b *Box2D) CollidePolygonAndCapsule(polygonA Polygon, xfA Transform, capsuleB Capsule, xfB Transform) (r Manifold) {
	polygonAStack := copyToStack(b.tls, polygonA)
	defer polygonAStack.Free()
	capsuleBStack := copyToStack(b.tls, capsuleB)
	defer capsuleBStack.Free()
	r = b2CollidePolygonAndCapsule(b.tls, polygonAStack.Addr, xfA, capsuleBStack.Addr, xfB)
	return
}
func (b *Box2D) CollidePolygons(polygonA Polygon, xfA Transform, polygonB Polygon, xfB Transform) (r Manifold) {
	polygonAStack := copyToStack(b.tls, polygonA)
	defer polygonAStack.Free()
	polygonBStack := copyToStack(b.tls, polygonB)
	defer polygonBStack.Free()
	r = b2CollidePolygons(b.tls, polygonAStack.Addr, xfA, polygonBStack.Addr, xfB)
	return
}
func (b *Box2D) CollideSegmentAndCircle(segmentA Segment, xfA Transform, circleB Circle, xfB Transform) (r Manifold) {
	segmentAStack := copyToStack(b.tls, segmentA)
	defer segmentAStack.Free()
	circleBStack := copyToStack(b.tls, circleB)
	defer circleBStack.Free()
	r = b2CollideSegmentAndCircle(b.tls, segmentAStack.Addr, xfA, circleBStack.Addr, xfB)
	return
}
func (b *Box2D) CollideSegmentAndPolygon(segmentA Segment, xfA Transform, polygonB Polygon, xfB Transform) (r Manifold) {
	segmentAStack := copyToStack(b.tls, segmentA)
	defer segmentAStack.Free()
	polygonBStack := copyToStack(b.tls, polygonB)
	defer polygonBStack.Free()
	r = b2CollideSegmentAndPolygon(b.tls, segmentAStack.Addr, xfA, polygonBStack.Addr, xfB)
	return
}
func (b *Box2D) CollideChainSegmentAndCircle(segmentA ChainSegment, xfA Transform, circleB Circle, xfB Transform) (r Manifold) {
	segmentAStack := copyToStack(b.tls, segmentA)
	defer segmentAStack.Free()
	circleBStack := copyToStack(b.tls, circleB)
	defer circleBStack.Free()
	r = b2CollideChainSegmentAndCircle(b.tls, segmentAStack.Addr, xfA, circleBStack.Addr, xfB)
	return
}
func (b *Box2D) CollideChainSegmentAndCapsule(segmentA ChainSegment, xfA Transform, capsuleB Capsule, xfB Transform, cache uintptr) (r Manifold) {
	segmentAStack := copyToStack(b.tls, segmentA)
	defer segmentAStack.Free()
	capsuleBStack := copyToStack(b.tls, capsuleB)
	defer capsuleBStack.Free()
	r = b2CollideChainSegmentAndCapsule(b.tls, segmentAStack.Addr, xfA, capsuleBStack.Addr, xfB, cache)
	return
}
func (b *Box2D) CollideChainSegmentAndPolygon(segmentA ChainSegment, xfA Transform, polygonB Polygon, xfB Transform, cache uintptr) (r Manifold) {
	segmentAStack := copyToStack(b.tls, segmentA)
	defer segmentAStack.Free()
	polygonBStack := copyToStack(b.tls, polygonB)
	defer polygonBStack.Free()
	r = b2CollideChainSegmentAndPolygon(b.tls, segmentAStack.Addr, xfA, polygonBStack.Addr, xfB, cache)
	return
}
func (b *Box2D) IsValidFloat(a float32) (r uint8) {
	r = b2IsValidFloat(b.tls, a)
	return
}
func (b *Box2D) IsValidVec2(v Vec2) (r uint8) {
	r = b2IsValidVec2(b.tls, v)
	return
}
func (b *Box2D) IsValidRotation(q1 Rot) (r uint8) {
	r = b2IsValidRotation(b.tls, q1)
	return
}
func (b *Box2D) IsValidPlane(a3 Plane) (r uint8) {
	r = b2IsValidPlane(b.tls, a3)
	return
}
func (b *Box2D) Atan2(y float32, x float32) (r1 float32) {
	r1 = b2Atan2(b.tls, y, x)
	return
}
func (b *Box2D) ComputeCosSin(radians1 float32) (r CosSin) {
	r = b2ComputeCosSin(b.tls, radians1)
	return
}
func (b *Box2D) ComputeRotationBetweenUnitVectors(v1 Vec2, v2 Vec2) (r Rot) {
	r = b2ComputeRotationBetweenUnitVectors(b.tls, v1, v2)
	return
}
func (b *Box2D) MotorJointSetLinearOffset(jointId JointId, linearOffset Vec2) {
	b2MotorJoint_SetLinearOffset(b.tls, jointId, linearOffset)
}
func (b *Box2D) MotorJointGetLinearOffset(jointId JointId) (r Vec2) {
	r = b2MotorJoint_GetLinearOffset(b.tls, jointId)
	return
}
func (b *Box2D) MotorJointSetAngularOffset(jointId JointId, angularOffset float32) {
	b2MotorJoint_SetAngularOffset(b.tls, jointId, angularOffset)
}
func (b *Box2D) MotorJointGetAngularOffset(jointId JointId) (r float32) {
	r = b2MotorJoint_GetAngularOffset(b.tls, jointId)
	return
}
func (b *Box2D) MotorJointSetMaxForce(jointId JointId, maxForce float32) {
	b2MotorJoint_SetMaxForce(b.tls, jointId, maxForce)
}
func (b *Box2D) MotorJointGetMaxForce(jointId JointId) (r float32) {
	r = b2MotorJoint_GetMaxForce(b.tls, jointId)
	return
}
func (b *Box2D) MotorJointSetMaxTorque(jointId JointId, maxTorque float32) {
	b2MotorJoint_SetMaxTorque(b.tls, jointId, maxTorque)
}
func (b *Box2D) MotorJointGetMaxTorque(jointId JointId) (r float32) {
	r = b2MotorJoint_GetMaxTorque(b.tls, jointId)
	return
}
func (b *Box2D) MotorJointSetCorrectionFactor(jointId JointId, correctionFactor float32) {
	b2MotorJoint_SetCorrectionFactor(b.tls, jointId, correctionFactor)
}
func (b *Box2D) MotorJointGetCorrectionFactor(jointId JointId) (r float32) {
	r = b2MotorJoint_GetCorrectionFactor(b.tls, jointId)
	return
}
func (b *Box2D) MouseJointSetTarget(jointId JointId, target Vec2) {
	b2MouseJoint_SetTarget(b.tls, jointId, target)
}
func (b *Box2D) MouseJointGetTarget(jointId JointId) (r Vec2) {
	r = b2MouseJoint_GetTarget(b.tls, jointId)
	return
}
func (b *Box2D) MouseJointSetSpringHertz(jointId JointId, hertz float32) {
	b2MouseJoint_SetSpringHertz(b.tls, jointId, hertz)
}
func (b *Box2D) MouseJointGetSpringHertz(jointId JointId) (r float32) {
	r = b2MouseJoint_GetSpringHertz(b.tls, jointId)
	return
}
func (b *Box2D) MouseJointSetSpringDampingRatio(jointId JointId, dampingRatio float32) {
	b2MouseJoint_SetSpringDampingRatio(b.tls, jointId, dampingRatio)
}
func (b *Box2D) MouseJointGetSpringDampingRatio(jointId JointId) (r float32) {
	r = b2MouseJoint_GetSpringDampingRatio(b.tls, jointId)
	return
}
func (b *Box2D) MouseJointSetMaxForce(jointId JointId, maxForce float32) {
	b2MouseJoint_SetMaxForce(b.tls, jointId, maxForce)
}
func (b *Box2D) MouseJointGetMaxForce(jointId JointId) (r float32) {
	r = b2MouseJoint_GetMaxForce(b.tls, jointId)
	return
}
func (b *Box2D) SolvePlanes(targetDelta Vec2, planes *CollisionPlane, count int32) (r PlaneSolverResult) {
	defer runtime.KeepAlive(planes)
	escapes(planes)
	r = b2SolvePlanes(b.tls, targetDelta, uintptr(unsafe.Pointer(planes)), count)
	return
}
func (b *Box2D) ClipVector(vector Vec2, planes CollisionPlane, count int32) (r Vec2) {
	planesStack := copyToStack(b.tls, planes)
	defer planesStack.Free()
	r = b2ClipVector(b.tls, vector, planesStack.Addr, count)
	return
}
func (b *Box2D) PrismaticJointEnableSpring(jointId JointId, enableSpring uint8) {
	b2PrismaticJoint_EnableSpring(b.tls, jointId, enableSpring)
}
func (b *Box2D) PrismaticJointIsSpringEnabled(jointId JointId) (r uint8) {
	r = b2PrismaticJoint_IsSpringEnabled(b.tls, jointId)
	return
}
func (b *Box2D) PrismaticJointSetSpringHertz(jointId JointId, hertz float32) {
	b2PrismaticJoint_SetSpringHertz(b.tls, jointId, hertz)
}
func (b *Box2D) PrismaticJointGetSpringHertz(jointId JointId) (r float32) {
	r = b2PrismaticJoint_GetSpringHertz(b.tls, jointId)
	return
}
func (b *Box2D) PrismaticJointSetSpringDampingRatio(jointId JointId, dampingRatio float32) {
	b2PrismaticJoint_SetSpringDampingRatio(b.tls, jointId, dampingRatio)
}
func (b *Box2D) PrismaticJointGetSpringDampingRatio(jointId JointId) (r float32) {
	r = b2PrismaticJoint_GetSpringDampingRatio(b.tls, jointId)
	return
}
func (b *Box2D) PrismaticJointSetTargetTranslation(jointId JointId, translation float32) {
	b2PrismaticJoint_SetTargetTranslation(b.tls, jointId, translation)
}
func (b *Box2D) PrismaticJointGetTargetTranslation(jointId JointId) (r float32) {
	r = b2PrismaticJoint_GetTargetTranslation(b.tls, jointId)
	return
}
func (b *Box2D) PrismaticJointEnableLimit(jointId JointId, enableLimit uint8) {
	b2PrismaticJoint_EnableLimit(b.tls, jointId, enableLimit)
}
func (b *Box2D) PrismaticJointIsLimitEnabled(jointId JointId) (r uint8) {
	r = b2PrismaticJoint_IsLimitEnabled(b.tls, jointId)
	return
}
func (b *Box2D) PrismaticJointGetLowerLimit(jointId JointId) (r float32) {
	r = b2PrismaticJoint_GetLowerLimit(b.tls, jointId)
	return
}
func (b *Box2D) PrismaticJointGetUpperLimit(jointId JointId) (r float32) {
	r = b2PrismaticJoint_GetUpperLimit(b.tls, jointId)
	return
}
func (b *Box2D) PrismaticJointSetLimits(jointId JointId, lower float32, upper float32) {
	b2PrismaticJoint_SetLimits(b.tls, jointId, lower, upper)
}
func (b *Box2D) PrismaticJointEnableMotor(jointId JointId, enableMotor uint8) {
	b2PrismaticJoint_EnableMotor(b.tls, jointId, enableMotor)
}
func (b *Box2D) PrismaticJointIsMotorEnabled(jointId JointId) (r uint8) {
	r = b2PrismaticJoint_IsMotorEnabled(b.tls, jointId)
	return
}
func (b *Box2D) PrismaticJointSetMotorSpeed(jointId JointId, motorSpeed float32) {
	b2PrismaticJoint_SetMotorSpeed(b.tls, jointId, motorSpeed)
}
func (b *Box2D) PrismaticJointGetMotorSpeed(jointId JointId) (r float32) {
	r = b2PrismaticJoint_GetMotorSpeed(b.tls, jointId)
	return
}
func (b *Box2D) PrismaticJointGetMotorForce(jointId JointId) (r float32) {
	r = b2PrismaticJoint_GetMotorForce(b.tls, jointId)
	return
}
func (b *Box2D) PrismaticJointSetMaxMotorForce(jointId JointId, force float32) {
	b2PrismaticJoint_SetMaxMotorForce(b.tls, jointId, force)
}
func (b *Box2D) PrismaticJointGetMaxMotorForce(jointId JointId) (r float32) {
	r = b2PrismaticJoint_GetMaxMotorForce(b.tls, jointId)
	return
}
func (b *Box2D) PrismaticJointGetTranslation(jointId JointId) (r float32) {
	r = b2PrismaticJoint_GetTranslation(b.tls, jointId)
	return
}
func (b *Box2D) PrismaticJointGetSpeed(jointId JointId) (r float32) {
	r = b2PrismaticJoint_GetSpeed(b.tls, jointId)
	return
}
func (b *Box2D) RevoluteJointEnableSpring(jointId JointId, enableSpring uint8) {
	b2RevoluteJoint_EnableSpring(b.tls, jointId, enableSpring)
}
func (b *Box2D) RevoluteJointIsSpringEnabled(jointId JointId) (r uint8) {
	r = b2RevoluteJoint_IsSpringEnabled(b.tls, jointId)
	return
}
func (b *Box2D) RevoluteJointSetSpringHertz(jointId JointId, hertz float32) {
	b2RevoluteJoint_SetSpringHertz(b.tls, jointId, hertz)
}
func (b *Box2D) RevoluteJointGetSpringHertz(jointId JointId) (r float32) {
	r = b2RevoluteJoint_GetSpringHertz(b.tls, jointId)
	return
}
func (b *Box2D) RevoluteJointSetSpringDampingRatio(jointId JointId, dampingRatio float32) {
	b2RevoluteJoint_SetSpringDampingRatio(b.tls, jointId, dampingRatio)
}
func (b *Box2D) RevoluteJointGetSpringDampingRatio(jointId JointId) (r float32) {
	r = b2RevoluteJoint_GetSpringDampingRatio(b.tls, jointId)
	return
}
func (b *Box2D) RevoluteJointSetTargetAngle(jointId JointId, angle float32) {
	b2RevoluteJoint_SetTargetAngle(b.tls, jointId, angle)
}
func (b *Box2D) RevoluteJointGetTargetAngle(jointId JointId) (r float32) {
	r = b2RevoluteJoint_GetTargetAngle(b.tls, jointId)
	return
}
func (b *Box2D) RevoluteJointGetAngle(jointId JointId) (r float32) {
	r = b2RevoluteJoint_GetAngle(b.tls, jointId)
	return
}
func (b *Box2D) RevoluteJointEnableLimit(jointId JointId, enableLimit uint8) {
	b2RevoluteJoint_EnableLimit(b.tls, jointId, enableLimit)
}
func (b *Box2D) RevoluteJointIsLimitEnabled(jointId JointId) (r uint8) {
	r = b2RevoluteJoint_IsLimitEnabled(b.tls, jointId)
	return
}
func (b *Box2D) RevoluteJointGetLowerLimit(jointId JointId) (r float32) {
	r = b2RevoluteJoint_GetLowerLimit(b.tls, jointId)
	return
}
func (b *Box2D) RevoluteJointGetUpperLimit(jointId JointId) (r float32) {
	r = b2RevoluteJoint_GetUpperLimit(b.tls, jointId)
	return
}
func (b *Box2D) RevoluteJointSetLimits(jointId JointId, lower float32, upper float32) {
	b2RevoluteJoint_SetLimits(b.tls, jointId, lower, upper)
}
func (b *Box2D) RevoluteJointEnableMotor(jointId JointId, enableMotor uint8) {
	b2RevoluteJoint_EnableMotor(b.tls, jointId, enableMotor)
}
func (b *Box2D) RevoluteJointIsMotorEnabled(jointId JointId) (r uint8) {
	r = b2RevoluteJoint_IsMotorEnabled(b.tls, jointId)
	return
}
func (b *Box2D) RevoluteJointSetMotorSpeed(jointId JointId, motorSpeed float32) {
	b2RevoluteJoint_SetMotorSpeed(b.tls, jointId, motorSpeed)
}
func (b *Box2D) RevoluteJointGetMotorSpeed(jointId JointId) (r float32) {
	r = b2RevoluteJoint_GetMotorSpeed(b.tls, jointId)
	return
}
func (b *Box2D) RevoluteJointGetMotorTorque(jointId JointId) (r float32) {
	r = b2RevoluteJoint_GetMotorTorque(b.tls, jointId)
	return
}
func (b *Box2D) RevoluteJointSetMaxMotorTorque(jointId JointId, torque float32) {
	b2RevoluteJoint_SetMaxMotorTorque(b.tls, jointId, torque)
}
func (b *Box2D) RevoluteJointGetMaxMotorTorque(jointId JointId) (r float32) {
	r = b2RevoluteJoint_GetMaxMotorTorque(b.tls, jointId)
	return
}
func (b *Box2D) CreateCircleShape(bodyId BodyId, def ShapeDef, circle Circle) (r ShapeId) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	circleStack := copyToStack(b.tls, circle)
	defer circleStack.Free()
	r = b2CreateCircleShape(b.tls, bodyId, defStack.Addr, circleStack.Addr)
	return
}
func (b *Box2D) CreateCapsuleShape(bodyId BodyId, def ShapeDef, capsule Capsule) (r ShapeId) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	capsuleStack := copyToStack(b.tls, capsule)
	defer capsuleStack.Free()
	r = b2CreateCapsuleShape(b.tls, bodyId, defStack.Addr, capsuleStack.Addr)
	return
}
func (b *Box2D) CreatePolygonShape(bodyId BodyId, def ShapeDef, polygon Polygon) (r ShapeId) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	polygonStack := copyToStack(b.tls, polygon)
	defer polygonStack.Free()
	r = b2CreatePolygonShape(b.tls, bodyId, defStack.Addr, polygonStack.Addr)
	return
}
func (b *Box2D) CreateSegmentShape(bodyId BodyId, def ShapeDef, segment Segment) (r ShapeId) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	segmentStack := copyToStack(b.tls, segment)
	defer segmentStack.Free()
	r = b2CreateSegmentShape(b.tls, bodyId, defStack.Addr, segmentStack.Addr)
	return
}
func (b *Box2D) DestroyShape(shapeId ShapeId, updateBodyMass uint8) {
	b2DestroyShape(b.tls, shapeId, updateBodyMass)
}
func (b *Box2D) CreateChain(bodyId BodyId, def ChainDef) (r ChainId) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r = b2CreateChain(b.tls, bodyId, defStack.Addr)
	return
}
func (b *Box2D) DestroyChain(chainId ChainId) {
	b2DestroyChain(b.tls, chainId)
}
func (b *Box2D) ChainGetWorld(chainId ChainId) (r WorldId) {
	r = b2Chain_GetWorld(b.tls, chainId)
	return
}
func (b *Box2D) ChainGetSegmentCount(chainId ChainId) (r int32) {
	r = b2Chain_GetSegmentCount(b.tls, chainId)
	return
}
func (b *Box2D) ShapeGetBody(shapeId ShapeId) (r BodyId) {
	r = b2Shape_GetBody(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeGetWorld(shapeId ShapeId) (r WorldId) {
	r = b2Shape_GetWorld(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeSetUserData(shapeId ShapeId, userData uintptr) {
	b2Shape_SetUserData(b.tls, shapeId, userData)
}
func (b *Box2D) ShapeGetUserData(shapeId ShapeId) (r uintptr) {
	r = b2Shape_GetUserData(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeIsSensor(shapeId ShapeId) (r uint8) {
	r = b2Shape_IsSensor(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeTestPoint(shapeId ShapeId, point Vec2) (r uint8) {
	r = b2Shape_TestPoint(b.tls, shapeId, point)
	return
}
func (b *Box2D) ShapeRayCast(shapeId ShapeId, input RayCastInput) (r CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	r = b2Shape_RayCast(b.tls, shapeId, inputStack.Addr)
	return
}
func (b *Box2D) ShapeSetDensity(shapeId ShapeId, density float32, updateBodyMass uint8) {
	b2Shape_SetDensity(b.tls, shapeId, density, updateBodyMass)
}
func (b *Box2D) ShapeGetDensity(shapeId ShapeId) (r float32) {
	r = b2Shape_GetDensity(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeSetFriction(shapeId ShapeId, friction float32) {
	b2Shape_SetFriction(b.tls, shapeId, friction)
}
func (b *Box2D) ShapeGetFriction(shapeId ShapeId) (r float32) {
	r = b2Shape_GetFriction(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeSetRestitution(shapeId ShapeId, restitution float32) {
	b2Shape_SetRestitution(b.tls, shapeId, restitution)
}
func (b *Box2D) ShapeGetRestitution(shapeId ShapeId) (r float32) {
	r = b2Shape_GetRestitution(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeSetMaterial(shapeId ShapeId, material int32) {
	b2Shape_SetMaterial(b.tls, shapeId, material)
}
func (b *Box2D) ShapeGetMaterial(shapeId ShapeId) (r int32) {
	r = b2Shape_GetMaterial(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeGetSurfaceMaterial(shapeId ShapeId) (r SurfaceMaterial) {
	r = b2Shape_GetSurfaceMaterial(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeSetSurfaceMaterial(shapeId ShapeId, surfaceMaterial SurfaceMaterial) {
	b2Shape_SetSurfaceMaterial(b.tls, shapeId, surfaceMaterial)
}
func (b *Box2D) ShapeGetFilter(shapeId ShapeId) (r Filter) {
	r = b2Shape_GetFilter(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeSetFilter(shapeId ShapeId, filter Filter) {
	b2Shape_SetFilter(b.tls, shapeId, filter)
}
func (b *Box2D) ShapeEnableSensorEvents(shapeId ShapeId, flag uint8) {
	b2Shape_EnableSensorEvents(b.tls, shapeId, flag)
}
func (b *Box2D) ShapeAreSensorEventsEnabled(shapeId ShapeId) (r uint8) {
	r = b2Shape_AreSensorEventsEnabled(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeEnableContactEvents(shapeId ShapeId, flag uint8) {
	b2Shape_EnableContactEvents(b.tls, shapeId, flag)
}
func (b *Box2D) ShapeAreContactEventsEnabled(shapeId ShapeId) (r uint8) {
	r = b2Shape_AreContactEventsEnabled(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeEnablePreSolveEvents(shapeId ShapeId, flag uint8) {
	b2Shape_EnablePreSolveEvents(b.tls, shapeId, flag)
}
func (b *Box2D) ShapeArePreSolveEventsEnabled(shapeId ShapeId) (r uint8) {
	r = b2Shape_ArePreSolveEventsEnabled(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeEnableHitEvents(shapeId ShapeId, flag uint8) {
	b2Shape_EnableHitEvents(b.tls, shapeId, flag)
}
func (b *Box2D) ShapeAreHitEventsEnabled(shapeId ShapeId) (r uint8) {
	r = b2Shape_AreHitEventsEnabled(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeGetType(shapeId ShapeId) (r ShapeType) {
	r = b2Shape_GetType(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeGetCircle(shapeId ShapeId) (r Circle) {
	r = b2Shape_GetCircle(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeGetSegment(shapeId ShapeId) (r Segment) {
	r = b2Shape_GetSegment(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeGetChainSegment(shapeId ShapeId) (r ChainSegment) {
	r = b2Shape_GetChainSegment(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeGetCapsule(shapeId ShapeId) (r Capsule) {
	r = b2Shape_GetCapsule(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeGetPolygon(shapeId ShapeId) (r Polygon) {
	r = b2Shape_GetPolygon(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeSetCircle(shapeId ShapeId, circle Circle) {
	circleStack := copyToStack(b.tls, circle)
	defer circleStack.Free()
	b2Shape_SetCircle(b.tls, shapeId, circleStack.Addr)
}
func (b *Box2D) ShapeSetCapsule(shapeId ShapeId, capsule Capsule) {
	capsuleStack := copyToStack(b.tls, capsule)
	defer capsuleStack.Free()
	b2Shape_SetCapsule(b.tls, shapeId, capsuleStack.Addr)
}
func (b *Box2D) ShapeSetSegment(shapeId ShapeId, segment Segment) {
	segmentStack := copyToStack(b.tls, segment)
	defer segmentStack.Free()
	b2Shape_SetSegment(b.tls, shapeId, segmentStack.Addr)
}
func (b *Box2D) ShapeSetPolygon(shapeId ShapeId, polygon Polygon) {
	polygonStack := copyToStack(b.tls, polygon)
	defer polygonStack.Free()
	b2Shape_SetPolygon(b.tls, shapeId, polygonStack.Addr)
}
func (b *Box2D) ShapeGetParentChain(shapeId ShapeId) (r ChainId) {
	r = b2Shape_GetParentChain(b.tls, shapeId)
	return
}
func (b *Box2D) ChainSetFriction(chainId ChainId, friction float32) {
	b2Chain_SetFriction(b.tls, chainId, friction)
}
func (b *Box2D) ChainGetFriction(chainId ChainId) (r float32) {
	r = b2Chain_GetFriction(b.tls, chainId)
	return
}
func (b *Box2D) ChainSetRestitution(chainId ChainId, restitution float32) {
	b2Chain_SetRestitution(b.tls, chainId, restitution)
}
func (b *Box2D) ChainGetRestitution(chainId ChainId) (r float32) {
	r = b2Chain_GetRestitution(b.tls, chainId)
	return
}
func (b *Box2D) ChainSetMaterial(chainId ChainId, material int32) {
	b2Chain_SetMaterial(b.tls, chainId, material)
}
func (b *Box2D) ChainGetMaterial(chainId ChainId) (r int32) {
	r = b2Chain_GetMaterial(b.tls, chainId)
	return
}
func (b *Box2D) ShapeGetContactCapacity(shapeId ShapeId) (r int32) {
	r = b2Shape_GetContactCapacity(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeGetContactData(shapeId ShapeId, contactData *ContactData, capacity int32) (r int32) {
	defer runtime.KeepAlive(contactData)
	escapes(contactData)
	r = b2Shape_GetContactData(b.tls, shapeId, uintptr(unsafe.Pointer(contactData)), capacity)
	return
}
func (b *Box2D) ShapeGetSensorCapacity(shapeId ShapeId) (r int32) {
	r = b2Shape_GetSensorCapacity(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeGetSensorOverlaps(shapeId ShapeId, overlaps *ShapeId, capacity int32) (r int32) {
	defer runtime.KeepAlive(overlaps)
	escapes(overlaps)
	r = b2Shape_GetSensorOverlaps(b.tls, shapeId, uintptr(unsafe.Pointer(overlaps)), capacity)
	return
}
func (b *Box2D) ShapeGetAABB(shapeId ShapeId) (r AABB) {
	r = b2Shape_GetAABB(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeGetMassData(shapeId ShapeId) (r MassData) {
	r = b2Shape_GetMassData(b.tls, shapeId)
	return
}
func (b *Box2D) ShapeGetClosestPoint(shapeId ShapeId, _target Vec2) (r Vec2) {
	r = b2Shape_GetClosestPoint(b.tls, shapeId, _target)
	return
}
func (b *Box2D) GetTicks() (r uint64) {
	r = b2GetTicks(b.tls)
	return
}
func (b *Box2D) GetMilliseconds(ticks uint64) (r float32) {
	r = b2GetMilliseconds(b.tls, ticks)
	return
}
func (b *Box2D) GetMillisecondsAndReset(ticks uintptr) (r float32) {
	r = b2GetMillisecondsAndReset(b.tls, ticks)
	return
}
func (b *Box2D) Yield() {
	b2Yield(b.tls)
}
func (b *Box2D) Hash(hash uint32, data uintptr, count int32) (r uint32) {
	r = b2Hash(b.tls, hash, data, count)
	return
}
func (b *Box2D) DefaultWorldDef() (r WorldDef) {
	r = b2DefaultWorldDef(b.tls)
	return
}
func (b *Box2D) DefaultBodyDef() (r BodyDef) {
	r = b2DefaultBodyDef(b.tls)
	return
}
func (b *Box2D) DefaultFilter() (r Filter) {
	r = b2DefaultFilter(b.tls)
	return
}
func (b *Box2D) DefaultQueryFilter() (r QueryFilter) {
	r = b2DefaultQueryFilter(b.tls)
	return
}
func (b *Box2D) DefaultShapeDef() (r ShapeDef) {
	r = b2DefaultShapeDef(b.tls)
	return
}
func (b *Box2D) DefaultSurfaceMaterial() (r SurfaceMaterial) {
	r = b2DefaultSurfaceMaterial(b.tls)
	return
}
func (b *Box2D) DefaultChainDef() (r ChainDef) {
	r = b2DefaultChainDef(b.tls)
	return
}
func (b *Box2D) DefaultDebugDraw() (r DebugDraw) {
	r = b2DefaultDebugDraw(b.tls)
	return
}
func (b *Box2D) WeldJointSetLinearHertz(jointId JointId, hertz float32) {
	b2WeldJoint_SetLinearHertz(b.tls, jointId, hertz)
}
func (b *Box2D) WeldJointGetLinearHertz(jointId JointId) (r float32) {
	r = b2WeldJoint_GetLinearHertz(b.tls, jointId)
	return
}
func (b *Box2D) WeldJointSetLinearDampingRatio(jointId JointId, dampingRatio float32) {
	b2WeldJoint_SetLinearDampingRatio(b.tls, jointId, dampingRatio)
}
func (b *Box2D) WeldJointGetLinearDampingRatio(jointId JointId) (r float32) {
	r = b2WeldJoint_GetLinearDampingRatio(b.tls, jointId)
	return
}
func (b *Box2D) WeldJointSetAngularHertz(jointId JointId, hertz float32) {
	b2WeldJoint_SetAngularHertz(b.tls, jointId, hertz)
}
func (b *Box2D) WeldJointGetAngularHertz(jointId JointId) (r float32) {
	r = b2WeldJoint_GetAngularHertz(b.tls, jointId)
	return
}
func (b *Box2D) WeldJointSetAngularDampingRatio(jointId JointId, dampingRatio float32) {
	b2WeldJoint_SetAngularDampingRatio(b.tls, jointId, dampingRatio)
}
func (b *Box2D) WeldJointGetAngularDampingRatio(jointId JointId) (r float32) {
	r = b2WeldJoint_GetAngularDampingRatio(b.tls, jointId)
	return
}
func (b *Box2D) WheelJointEnableSpring(jointId JointId, enableSpring uint8) {
	b2WheelJoint_EnableSpring(b.tls, jointId, enableSpring)
}
func (b *Box2D) WheelJointIsSpringEnabled(jointId JointId) (r uint8) {
	r = b2WheelJoint_IsSpringEnabled(b.tls, jointId)
	return
}
func (b *Box2D) WheelJointSetSpringHertz(jointId JointId, hertz float32) {
	b2WheelJoint_SetSpringHertz(b.tls, jointId, hertz)
}
func (b *Box2D) WheelJointGetSpringHertz(jointId JointId) (r float32) {
	r = b2WheelJoint_GetSpringHertz(b.tls, jointId)
	return
}
func (b *Box2D) WheelJointSetSpringDampingRatio(jointId JointId, dampingRatio float32) {
	b2WheelJoint_SetSpringDampingRatio(b.tls, jointId, dampingRatio)
}
func (b *Box2D) WheelJointGetSpringDampingRatio(jointId JointId) (r float32) {
	r = b2WheelJoint_GetSpringDampingRatio(b.tls, jointId)
	return
}
func (b *Box2D) WheelJointEnableLimit(jointId JointId, enableLimit uint8) {
	b2WheelJoint_EnableLimit(b.tls, jointId, enableLimit)
}
func (b *Box2D) WheelJointIsLimitEnabled(jointId JointId) (r uint8) {
	r = b2WheelJoint_IsLimitEnabled(b.tls, jointId)
	return
}
func (b *Box2D) WheelJointGetLowerLimit(jointId JointId) (r float32) {
	r = b2WheelJoint_GetLowerLimit(b.tls, jointId)
	return
}
func (b *Box2D) WheelJointGetUpperLimit(jointId JointId) (r float32) {
	r = b2WheelJoint_GetUpperLimit(b.tls, jointId)
	return
}
func (b *Box2D) WheelJointSetLimits(jointId JointId, lower float32, upper float32) {
	b2WheelJoint_SetLimits(b.tls, jointId, lower, upper)
}
func (b *Box2D) WheelJointEnableMotor(jointId JointId, enableMotor uint8) {
	b2WheelJoint_EnableMotor(b.tls, jointId, enableMotor)
}
func (b *Box2D) WheelJointIsMotorEnabled(jointId JointId) (r uint8) {
	r = b2WheelJoint_IsMotorEnabled(b.tls, jointId)
	return
}
func (b *Box2D) WheelJointSetMotorSpeed(jointId JointId, motorSpeed float32) {
	b2WheelJoint_SetMotorSpeed(b.tls, jointId, motorSpeed)
}
func (b *Box2D) WheelJointGetMotorSpeed(jointId JointId) (r float32) {
	r = b2WheelJoint_GetMotorSpeed(b.tls, jointId)
	return
}
func (b *Box2D) WheelJointGetMotorTorque(jointId JointId) (r float32) {
	r = b2WheelJoint_GetMotorTorque(b.tls, jointId)
	return
}
func (b *Box2D) WheelJointSetMaxMotorTorque(jointId JointId, torque float32) {
	b2WheelJoint_SetMaxMotorTorque(b.tls, jointId, torque)
}
func (b *Box2D) WheelJointGetMaxMotorTorque(jointId JointId) (r float32) {
	r = b2WheelJoint_GetMaxMotorTorque(b.tls, jointId)
	return
}
func (b *Box2D) CreateWorld(def WorldDef) (r WorldId) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r = b2CreateWorld(b.tls, defStack.Addr)
	return
}
func (b *Box2D) DestroyWorld(worldId WorldId) {
	b2DestroyWorld(b.tls, worldId)
}
func (b *Box2D) WorldStep(worldId WorldId, timeStep float32, subStepCount int32) {
	b2World_Step(b.tls, worldId, timeStep, subStepCount)
}
func (b *Box2D) WorldDraw(worldId WorldId, draw *DebugDraw) {
	defer runtime.KeepAlive(draw)
	escapes(draw)
	b2World_Draw(b.tls, worldId, uintptr(unsafe.Pointer(draw)))
}
func (b *Box2D) WorldIsValid(id WorldId) (r uint8) {
	r = b2World_IsValid(b.tls, id)
	return
}
func (b *Box2D) BodyIsValid(id BodyId) (r uint8) {
	r = b2Body_IsValid(b.tls, id)
	return
}
func (b *Box2D) ShapeIsValid(id ShapeId) (r uint8) {
	r = b2Shape_IsValid(b.tls, id)
	return
}
func (b *Box2D) ChainIsValid(id ChainId) (r uint8) {
	r = b2Chain_IsValid(b.tls, id)
	return
}
func (b *Box2D) JointIsValid(id JointId) (r uint8) {
	r = b2Joint_IsValid(b.tls, id)
	return
}
func (b *Box2D) WorldEnableSleeping(worldId WorldId, flag uint8) {
	b2World_EnableSleeping(b.tls, worldId, flag)
}
func (b *Box2D) WorldIsSleepingEnabled(worldId WorldId) (r uint8) {
	r = b2World_IsSleepingEnabled(b.tls, worldId)
	return
}
func (b *Box2D) WorldEnableWarmStarting(worldId WorldId, flag uint8) {
	b2World_EnableWarmStarting(b.tls, worldId, flag)
}
func (b *Box2D) WorldIsWarmStartingEnabled(worldId WorldId) (r uint8) {
	r = b2World_IsWarmStartingEnabled(b.tls, worldId)
	return
}
func (b *Box2D) WorldGetAwakeBodyCount(worldId WorldId) (r int32) {
	r = b2World_GetAwakeBodyCount(b.tls, worldId)
	return
}
func (b *Box2D) WorldEnableContinuous(worldId WorldId, flag uint8) {
	b2World_EnableContinuous(b.tls, worldId, flag)
}
func (b *Box2D) WorldIsContinuousEnabled(worldId WorldId) (r uint8) {
	r = b2World_IsContinuousEnabled(b.tls, worldId)
	return
}
func (b *Box2D) WorldSetRestitutionThreshold(worldId WorldId, value float32) {
	b2World_SetRestitutionThreshold(b.tls, worldId, value)
}
func (b *Box2D) WorldGetRestitutionThreshold(worldId WorldId) (r float32) {
	r = b2World_GetRestitutionThreshold(b.tls, worldId)
	return
}
func (b *Box2D) WorldSetHitEventThreshold(worldId WorldId, value float32) {
	b2World_SetHitEventThreshold(b.tls, worldId, value)
}
func (b *Box2D) WorldGetHitEventThreshold(worldId WorldId) (r float32) {
	r = b2World_GetHitEventThreshold(b.tls, worldId)
	return
}
func (b *Box2D) WorldSetContactTuning(worldId WorldId, hertz float32, dampingRatio float32, pushSpeed float32) {
	b2World_SetContactTuning(b.tls, worldId, hertz, dampingRatio, pushSpeed)
}
func (b *Box2D) WorldSetMaximumLinearSpeed(worldId WorldId, maximumLinearSpeed float32) {
	b2World_SetMaximumLinearSpeed(b.tls, worldId, maximumLinearSpeed)
}
func (b *Box2D) WorldGetMaximumLinearSpeed(worldId WorldId) (r float32) {
	r = b2World_GetMaximumLinearSpeed(b.tls, worldId)
	return
}
func (b *Box2D) WorldGetProfile(worldId WorldId) (r Profile) {
	r = b2World_GetProfile(b.tls, worldId)
	return
}
func (b *Box2D) WorldGetCounters(worldId WorldId) (r Counters) {
	r = b2World_GetCounters(b.tls, worldId)
	return
}
func (b *Box2D) WorldSetUserData(worldId WorldId, userData uintptr) {
	b2World_SetUserData(b.tls, worldId, userData)
}
func (b *Box2D) WorldGetUserData(worldId WorldId) (r uintptr) {
	r = b2World_GetUserData(b.tls, worldId)
	return
}
func (b *Box2D) WorldDumpMemoryStats(worldId WorldId) {
	b2World_DumpMemoryStats(b.tls, worldId)
}
func (b *Box2D) WorldOverlapShape(worldId WorldId, proxy ShapeProxy, filter QueryFilter, fcn uintptr, context uintptr) (r1 TreeStats) {
	proxyStack := copyToStack(b.tls, proxy)
	defer proxyStack.Free()
	r1 = b2World_OverlapShape(b.tls, worldId, proxyStack.Addr, filter, fcn, context)
	return
}
func (b *Box2D) WorldCastRay(worldId WorldId, origin Vec2, translation Vec2, filter QueryFilter, fcn uintptr, context uintptr) (r TreeStats) {
	r = b2World_CastRay(b.tls, worldId, origin, translation, filter, fcn, context)
	return
}
func (b *Box2D) WorldCastRayClosest(worldId WorldId, origin Vec2, translation Vec2, filter QueryFilter) (r RayResult) {
	r = b2World_CastRayClosest(b.tls, worldId, origin, translation, filter)
	return
}
func (b *Box2D) WorldCastShape(worldId WorldId, proxy ShapeProxy, translation Vec2, filter QueryFilter, fcn uintptr, context uintptr) (r TreeStats) {
	proxyStack := copyToStack(b.tls, proxy)
	defer proxyStack.Free()
	r = b2World_CastShape(b.tls, worldId, proxyStack.Addr, translation, filter, fcn, context)
	return
}
func (b *Box2D) WorldCastMover(worldId WorldId, mover Capsule, translation Vec2, filter QueryFilter) (r float32) {
	moverStack := copyToStack(b.tls, mover)
	defer moverStack.Free()
	r = b2World_CastMover(b.tls, worldId, moverStack.Addr, translation, filter)
	return
}
func (b *Box2D) WorldSetGravity(worldId WorldId, gravity Vec2) {
	b2World_SetGravity(b.tls, worldId, gravity)
}
func (b *Box2D) WorldGetGravity(worldId WorldId) (r Vec2) {
	r = b2World_GetGravity(b.tls, worldId)
	return
}
func (b *Box2D) WorldExplode(worldId WorldId, explosionDef ExplosionDef) {
	explosionDefStack := copyToStack(b.tls, explosionDef)
	defer explosionDefStack.Free()
	b2World_Explode(b.tls, worldId, explosionDefStack.Addr)
}
func (b *Box2D) WorldRebuildStaticTree(worldId WorldId) {
	b2World_RebuildStaticTree(b.tls, worldId)
}
func (b *Box2D) WorldEnableSpeculative(worldId WorldId, flag uint8) {
	b2World_EnableSpeculative(b.tls, worldId, flag)
}

const ToiStateUnknown TOIState = b2_toiStateUnknown
const ToiStateFailed TOIState = b2_toiStateFailed
const ToiStateOverlapped TOIState = b2_toiStateOverlapped
const ToiStateHit TOIState = b2_toiStateHit
const ToiStateSeparated TOIState = b2_toiStateSeparated
const StaticBody BodyType = b2_staticBody
const KinematicBody BodyType = b2_kinematicBody
const DynamicBody BodyType = b2_dynamicBody
const CircleShape ShapeType = b2_circleShape
const CapsuleShape ShapeType = b2_capsuleShape
const SegmentShape ShapeType = b2_segmentShape
const PolygonShape ShapeType = b2_polygonShape
const ChainSegmentShape ShapeType = b2_chainSegmentShape
const DistanceJoint JointType = b2_distanceJoint
const FilterJoint JointType = b2_filterJoint
const MotorJoint JointType = b2_motorJoint
const MouseJoint JointType = b2_mouseJoint
const PrismaticJoint JointType = b2_prismaticJoint
const RevoluteJoint JointType = b2_revoluteJoint
const WeldJoint JointType = b2_weldJoint
const WheelJoint JointType = b2_wheelJoint
const ColorAliceBlue HexColor = b2_colorAliceBlue
const ColorAntiqueWhite HexColor = b2_colorAntiqueWhite
const ColorAqua HexColor = b2_colorAqua
const ColorAquamarine HexColor = b2_colorAquamarine
const ColorAzure HexColor = b2_colorAzure
const ColorBeige HexColor = b2_colorBeige
const ColorBisque HexColor = b2_colorBisque
const ColorBlack HexColor = b2_colorBlack
const ColorBlanchedAlmond HexColor = b2_colorBlanchedAlmond
const ColorBlue HexColor = b2_colorBlue
const ColorBlueViolet HexColor = b2_colorBlueViolet
const ColorBrown HexColor = b2_colorBrown
const ColorBurlywood HexColor = b2_colorBurlywood
const ColorCadetBlue HexColor = b2_colorCadetBlue
const ColorChartreuse HexColor = b2_colorChartreuse
const ColorChocolate HexColor = b2_colorChocolate
const ColorCoral HexColor = b2_colorCoral
const ColorCornflowerBlue HexColor = b2_colorCornflowerBlue
const ColorCornsilk HexColor = b2_colorCornsilk
const ColorCrimson HexColor = b2_colorCrimson
const ColorCyan HexColor = b2_colorCyan
const ColorDarkBlue HexColor = b2_colorDarkBlue
const ColorDarkCyan HexColor = b2_colorDarkCyan
const ColorDarkGoldenRod HexColor = b2_colorDarkGoldenRod
const ColorDarkGray HexColor = b2_colorDarkGray
const ColorDarkGreen HexColor = b2_colorDarkGreen
const ColorDarkKhaki HexColor = b2_colorDarkKhaki
const ColorDarkMagenta HexColor = b2_colorDarkMagenta
const ColorDarkOliveGreen HexColor = b2_colorDarkOliveGreen
const ColorDarkOrange HexColor = b2_colorDarkOrange
const ColorDarkOrchid HexColor = b2_colorDarkOrchid
const ColorDarkRed HexColor = b2_colorDarkRed
const ColorDarkSalmon HexColor = b2_colorDarkSalmon
const ColorDarkSeaGreen HexColor = b2_colorDarkSeaGreen
const ColorDarkSlateBlue HexColor = b2_colorDarkSlateBlue
const ColorDarkSlateGray HexColor = b2_colorDarkSlateGray
const ColorDarkTurquoise HexColor = b2_colorDarkTurquoise
const ColorDarkViolet HexColor = b2_colorDarkViolet
const ColorDeepPink HexColor = b2_colorDeepPink
const ColorDeepSkyBlue HexColor = b2_colorDeepSkyBlue
const ColorDimGray HexColor = b2_colorDimGray
const ColorDodgerBlue HexColor = b2_colorDodgerBlue
const ColorFireBrick HexColor = b2_colorFireBrick
const ColorFloralWhite HexColor = b2_colorFloralWhite
const ColorForestGreen HexColor = b2_colorForestGreen
const ColorFuchsia HexColor = b2_colorFuchsia
const ColorGainsboro HexColor = b2_colorGainsboro
const ColorGhostWhite HexColor = b2_colorGhostWhite
const ColorGold HexColor = b2_colorGold
const ColorGoldenRod HexColor = b2_colorGoldenRod
const ColorGray HexColor = b2_colorGray
const ColorGreen HexColor = b2_colorGreen
const ColorGreenYellow HexColor = b2_colorGreenYellow
const ColorHoneyDew HexColor = b2_colorHoneyDew
const ColorHotPink HexColor = b2_colorHotPink
const ColorIndianRed HexColor = b2_colorIndianRed
const ColorIndigo HexColor = b2_colorIndigo
const ColorIvory HexColor = b2_colorIvory
const ColorKhaki HexColor = b2_colorKhaki
const ColorLavender HexColor = b2_colorLavender
const ColorLavenderBlush HexColor = b2_colorLavenderBlush
const ColorLawnGreen HexColor = b2_colorLawnGreen
const ColorLemonChiffon HexColor = b2_colorLemonChiffon
const ColorLightBlue HexColor = b2_colorLightBlue
const ColorLightCoral HexColor = b2_colorLightCoral
const ColorLightCyan HexColor = b2_colorLightCyan
const ColorLightGoldenRodYellow HexColor = b2_colorLightGoldenRodYellow
const ColorLightGray HexColor = b2_colorLightGray
const ColorLightGreen HexColor = b2_colorLightGreen
const ColorLightPink HexColor = b2_colorLightPink
const ColorLightSalmon HexColor = b2_colorLightSalmon
const ColorLightSeaGreen HexColor = b2_colorLightSeaGreen
const ColorLightSkyBlue HexColor = b2_colorLightSkyBlue
const ColorLightSlateGray HexColor = b2_colorLightSlateGray
const ColorLightSteelBlue HexColor = b2_colorLightSteelBlue
const ColorLightYellow HexColor = b2_colorLightYellow
const ColorLime HexColor = b2_colorLime
const ColorLimeGreen HexColor = b2_colorLimeGreen
const ColorLinen HexColor = b2_colorLinen
const ColorMagenta HexColor = b2_colorMagenta
const ColorMaroon HexColor = b2_colorMaroon
const ColorMediumAquaMarine HexColor = b2_colorMediumAquaMarine
const ColorMediumBlue HexColor = b2_colorMediumBlue
const ColorMediumOrchid HexColor = b2_colorMediumOrchid
const ColorMediumPurple HexColor = b2_colorMediumPurple
const ColorMediumSeaGreen HexColor = b2_colorMediumSeaGreen
const ColorMediumSlateBlue HexColor = b2_colorMediumSlateBlue
const ColorMediumSpringGreen HexColor = b2_colorMediumSpringGreen
const ColorMediumTurquoise HexColor = b2_colorMediumTurquoise
const ColorMediumVioletRed HexColor = b2_colorMediumVioletRed
const ColorMidnightBlue HexColor = b2_colorMidnightBlue
const ColorMintCream HexColor = b2_colorMintCream
const ColorMistyRose HexColor = b2_colorMistyRose
const ColorMoccasin HexColor = b2_colorMoccasin
const ColorNavajoWhite HexColor = b2_colorNavajoWhite
const ColorNavy HexColor = b2_colorNavy
const ColorOldLace HexColor = b2_colorOldLace
const ColorOlive HexColor = b2_colorOlive
const ColorOliveDrab HexColor = b2_colorOliveDrab
const ColorOrange HexColor = b2_colorOrange
const ColorOrangeRed HexColor = b2_colorOrangeRed
const ColorOrchid HexColor = b2_colorOrchid
const ColorPaleGoldenRod HexColor = b2_colorPaleGoldenRod
const ColorPaleGreen HexColor = b2_colorPaleGreen
const ColorPaleTurquoise HexColor = b2_colorPaleTurquoise
const ColorPaleVioletRed HexColor = b2_colorPaleVioletRed
const ColorPapayaWhip HexColor = b2_colorPapayaWhip
const ColorPeachPuff HexColor = b2_colorPeachPuff
const ColorPeru HexColor = b2_colorPeru
const ColorPink HexColor = b2_colorPink
const ColorPlum HexColor = b2_colorPlum
const ColorPowderBlue HexColor = b2_colorPowderBlue
const ColorPurple HexColor = b2_colorPurple
const ColorRebeccaPurple HexColor = b2_colorRebeccaPurple
const ColorRed HexColor = b2_colorRed
const ColorRosyBrown HexColor = b2_colorRosyBrown
const ColorRoyalBlue HexColor = b2_colorRoyalBlue
const ColorSaddleBrown HexColor = b2_colorSaddleBrown
const ColorSalmon HexColor = b2_colorSalmon
const ColorSandyBrown HexColor = b2_colorSandyBrown
const ColorSeaGreen HexColor = b2_colorSeaGreen
const ColorSeaShell HexColor = b2_colorSeaShell
const ColorSienna HexColor = b2_colorSienna
const ColorSilver HexColor = b2_colorSilver
const ColorSkyBlue HexColor = b2_colorSkyBlue
const ColorSlateBlue HexColor = b2_colorSlateBlue
const ColorSlateGray HexColor = b2_colorSlateGray
const ColorSnow HexColor = b2_colorSnow
const ColorSpringGreen HexColor = b2_colorSpringGreen
const ColorSteelBlue HexColor = b2_colorSteelBlue
const ColorTan HexColor = b2_colorTan
const ColorTeal HexColor = b2_colorTeal
const ColorThistle HexColor = b2_colorThistle
const ColorTomato HexColor = b2_colorTomato
const ColorTurquoise HexColor = b2_colorTurquoise
const ColorViolet HexColor = b2_colorViolet
const ColorWheat HexColor = b2_colorWheat
const ColorWhite HexColor = b2_colorWhite
const ColorWhiteSmoke HexColor = b2_colorWhiteSmoke
const ColorYellow HexColor = b2_colorYellow
const ColorYellowGreen HexColor = b2_colorYellowGreen
const ColorBox2DRed HexColor = b2_colorBox2DRed
const ColorBox2DBlue HexColor = b2_colorBox2DBlue
const ColorBox2DGreen HexColor = b2_colorBox2DGreen
const ColorBox2DYellow HexColor = b2_colorBox2DYellow
