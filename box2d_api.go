package box2d

import (
	"unsafe"
	"runtime"
)

type Body struct {
	Id  BodyId
	tls *TLS
}
type Shape struct {
	Id  ShapeId
	tls *TLS
}
type World struct {
	Id  WorldId
	tls *TLS
}
type Joint struct {
	Id  JointId
	tls *TLS
}
type Chain struct {
	Id  ChainId
	tls *TLS
}
type PrismaticJoint struct {
	Joint
}
type MotorJoint struct {
	Joint
}
type RevoluteJoint struct {
	Joint
}
type DistanceJoint struct {
	Joint
}
type WheelJoint struct {
	Joint
}
type WeldJoint struct {
	Joint
}
type MouseJoint struct {
	Joint
}
type FilterJoint struct {
	Joint
}
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

func (b Box2D) IsValidAABB(a1 AABB) (r uint8) {
	r = b2IsValidAABB(b.tls, a1)
	return
}
func (b World) CreateBody(def BodyDef) (r Body) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r.tls = b.tls
	r.Id = b2CreateBody(b.tls, b.Id, defStack.Addr)
	return
}
func (b Body) DestroyBody() {
	b2DestroyBody(b.tls, b.Id)
}
func (b Body) GetContactCapacity() (r int32) {
	r = b2Body_GetContactCapacity(b.tls, b.Id)
	return
}
func (b Body) GetContactData(contactData *ContactData, capacity int32) (r int32) {
	defer runtime.KeepAlive(contactData)
	escapes(contactData)
	r = b2Body_GetContactData(b.tls, b.Id, uintptr(unsafe.Pointer(contactData)), capacity)
	return
}
func (b Body) ComputeAABB() (r AABB) {
	r = b2Body_ComputeAABB(b.tls, b.Id)
	return
}
func (b Body) GetPosition() (r Vec2) {
	r = b2Body_GetPosition(b.tls, b.Id)
	return
}
func (b Body) GetRotation() (r Rot) {
	r = b2Body_GetRotation(b.tls, b.Id)
	return
}
func (b Body) GetTransform() (r Transform) {
	r = b2Body_GetTransform(b.tls, b.Id)
	return
}
func (b Body) GetLocalPoint(worldPoint Vec2) (r Vec2) {
	r = b2Body_GetLocalPoint(b.tls, b.Id, worldPoint)
	return
}
func (b Body) GetWorldPoint(localPoint Vec2) (r Vec2) {
	r = b2Body_GetWorldPoint(b.tls, b.Id, localPoint)
	return
}
func (b Body) GetLocalVector(worldVector Vec2) (r Vec2) {
	r = b2Body_GetLocalVector(b.tls, b.Id, worldVector)
	return
}
func (b Body) GetWorldVector(localVector Vec2) (r Vec2) {
	r = b2Body_GetWorldVector(b.tls, b.Id, localVector)
	return
}
func (b Body) SetTransform(position Vec2, rotation Rot) {
	b2Body_SetTransform(b.tls, b.Id, position, rotation)
}
func (b Body) GetLinearVelocity() (r Vec2) {
	r = b2Body_GetLinearVelocity(b.tls, b.Id)
	return
}
func (b Body) GetAngularVelocity() (r float32) {
	r = b2Body_GetAngularVelocity(b.tls, b.Id)
	return
}
func (b Body) SetLinearVelocity(linearVelocity Vec2) {
	b2Body_SetLinearVelocity(b.tls, b.Id, linearVelocity)
}
func (b Body) SetAngularVelocity(angularVelocity float32) {
	b2Body_SetAngularVelocity(b.tls, b.Id, angularVelocity)
}
func (b Body) SetTargetTransform(target Transform, timeStep float32) {
	b2Body_SetTargetTransform(b.tls, b.Id, target, timeStep)
}
func (b Body) GetLocalPointVelocity(localPoint Vec2) (r1 Vec2) {
	r1 = b2Body_GetLocalPointVelocity(b.tls, b.Id, localPoint)
	return
}
func (b Body) GetWorldPointVelocity(worldPoint Vec2) (r1 Vec2) {
	r1 = b2Body_GetWorldPointVelocity(b.tls, b.Id, worldPoint)
	return
}
func (b Body) ApplyForce(force Vec2, point Vec2, wake uint8) {
	b2Body_ApplyForce(b.tls, b.Id, force, point, wake)
}
func (b Body) ApplyForceToCenter(force Vec2, wake uint8) {
	b2Body_ApplyForceToCenter(b.tls, b.Id, force, wake)
}
func (b Body) ApplyTorque(torque float32, wake uint8) {
	b2Body_ApplyTorque(b.tls, b.Id, torque, wake)
}
func (b Body) ApplyLinearImpulse(impulse Vec2, point Vec2, wake uint8) {
	b2Body_ApplyLinearImpulse(b.tls, b.Id, impulse, point, wake)
}
func (b Body) ApplyLinearImpulseToCenter(impulse Vec2, wake uint8) {
	b2Body_ApplyLinearImpulseToCenter(b.tls, b.Id, impulse, wake)
}
func (b Body) ApplyAngularImpulse(impulse float32, wake uint8) {
	b2Body_ApplyAngularImpulse(b.tls, b.Id, impulse, wake)
}
func (b Body) GetType() (r BodyType) {
	r = b2Body_GetType(b.tls, b.Id)
	return
}
func (b Body) SetType(type1 BodyType) {
	b2Body_SetType(b.tls, b.Id, type1)
}
func (b Body) SetUserData(userData uintptr) {
	b2Body_SetUserData(b.tls, b.Id, userData)
}
func (b Body) GetUserData() (r uintptr) {
	r = b2Body_GetUserData(b.tls, b.Id)
	return
}
func (b Body) GetMass() (r float32) {
	r = b2Body_GetMass(b.tls, b.Id)
	return
}
func (b Body) GetRotationalInertia() (r float32) {
	r = b2Body_GetRotationalInertia(b.tls, b.Id)
	return
}
func (b Body) GetLocalCenterOfMass() (r Vec2) {
	r = b2Body_GetLocalCenterOfMass(b.tls, b.Id)
	return
}
func (b Body) GetWorldCenterOfMass() (r Vec2) {
	r = b2Body_GetWorldCenterOfMass(b.tls, b.Id)
	return
}
func (b Body) SetMassData(massData MassData) {
	b2Body_SetMassData(b.tls, b.Id, massData)
}
func (b Body) GetMassData() (r MassData) {
	r = b2Body_GetMassData(b.tls, b.Id)
	return
}
func (b Body) ApplyMassFromShapes() {
	b2Body_ApplyMassFromShapes(b.tls, b.Id)
}
func (b Body) SetLinearDamping(linearDamping float32) {
	b2Body_SetLinearDamping(b.tls, b.Id, linearDamping)
}
func (b Body) GetLinearDamping() (r float32) {
	r = b2Body_GetLinearDamping(b.tls, b.Id)
	return
}
func (b Body) SetAngularDamping(angularDamping float32) {
	b2Body_SetAngularDamping(b.tls, b.Id, angularDamping)
}
func (b Body) GetAngularDamping() (r float32) {
	r = b2Body_GetAngularDamping(b.tls, b.Id)
	return
}
func (b Body) SetGravityScale(gravityScale float32) {
	b2Body_SetGravityScale(b.tls, b.Id, gravityScale)
}
func (b Body) GetGravityScale() (r float32) {
	r = b2Body_GetGravityScale(b.tls, b.Id)
	return
}
func (b Body) IsAwake() (r uint8) {
	r = b2Body_IsAwake(b.tls, b.Id)
	return
}
func (b Body) SetAwake(awake uint8) {
	b2Body_SetAwake(b.tls, b.Id, awake)
}
func (b Body) IsEnabled() (r uint8) {
	r = b2Body_IsEnabled(b.tls, b.Id)
	return
}
func (b Body) IsSleepEnabled() (r uint8) {
	r = b2Body_IsSleepEnabled(b.tls, b.Id)
	return
}
func (b Body) SetSleepThreshold(sleepThreshold float32) {
	b2Body_SetSleepThreshold(b.tls, b.Id, sleepThreshold)
}
func (b Body) GetSleepThreshold() (r float32) {
	r = b2Body_GetSleepThreshold(b.tls, b.Id)
	return
}
func (b Body) EnableSleep(enableSleep uint8) {
	b2Body_EnableSleep(b.tls, b.Id, enableSleep)
}
func (b Body) Disable() {
	b2Body_Disable(b.tls, b.Id)
}
func (b Body) Enable() {
	b2Body_Enable(b.tls, b.Id)
}
func (b Body) SetFixedRotation(flag uint8) {
	b2Body_SetFixedRotation(b.tls, b.Id, flag)
}
func (b Body) IsFixedRotation() (r uint8) {
	r = b2Body_IsFixedRotation(b.tls, b.Id)
	return
}
func (b Body) SetBullet(flag uint8) {
	b2Body_SetBullet(b.tls, b.Id, flag)
}
func (b Body) IsBullet() (r uint8) {
	r = b2Body_IsBullet(b.tls, b.Id)
	return
}
func (b Body) EnableContactEvents(flag uint8) {
	b2Body_EnableContactEvents(b.tls, b.Id, flag)
}
func (b Body) EnableHitEvents(flag uint8) {
	b2Body_EnableHitEvents(b.tls, b.Id, flag)
}
func (b Body) GetWorld() (r World) {
	r.tls = b.tls
	r.Id = b2Body_GetWorld(b.tls, b.Id)
	return
}
func (b Body) GetShapeCount() (r int32) {
	r = b2Body_GetShapeCount(b.tls, b.Id)
	return
}
func (b Body) GetJointCount() (r int32) {
	r = b2Body_GetJointCount(b.tls, b.Id)
	return
}
func (b Box2D) SetLengthUnitsPerMeter(lengthUnits float32) {
	b2SetLengthUnitsPerMeter(b.tls, lengthUnits)
}
func (b Box2D) GetLengthUnitsPerMeter() (r float32) {
	r = b2GetLengthUnitsPerMeter(b.tls)
	return
}
func (b Box2D) GetVersion() (r Version) {
	r = b2GetVersion(b.tls)
	return
}
func (b Box2D) GetByteCount() (r int32) {
	r = b2GetByteCount(b.tls)
	return
}
func (b Box2D) GetSweepTransform(sweep Sweep, time float32) (r Transform) {
	sweepStack := copyToStack(b.tls, sweep)
	defer sweepStack.Free()
	r = b2GetSweepTransform(b.tls, sweepStack.Addr, time)
	return
}
func (b Box2D) SegmentDistance(p1 Vec2, q1 Vec2, p2 Vec2, q2 Vec2) (r1 SegmentDistanceResult) {
	r1 = b2SegmentDistance(b.tls, p1, q1, p2, q2)
	return
}
func (b Box2D) MakeProxy(points Vec2, count int32, radius float32) (r ShapeProxy) {
	pointsStack := copyToStack(b.tls, points)
	defer pointsStack.Free()
	r = b2MakeProxy(b.tls, pointsStack.Addr, count, radius)
	return
}
func (b Box2D) MakeOffsetProxy(points Vec2, count int32, radius float32, position Vec2, rotation Rot) (r ShapeProxy) {
	pointsStack := copyToStack(b.tls, points)
	defer pointsStack.Free()
	r = b2MakeOffsetProxy(b.tls, pointsStack.Addr, count, radius, position, rotation)
	return
}
func (b Box2D) ShapeDistance(input DistanceInput, cache *SimplexCache, simplexes *Simplex, simplexCapacity int32) (r DistanceOutput) {
	defer runtime.KeepAlive(simplexes)
	defer runtime.KeepAlive(cache)
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	escapes(cache)
	escapes(simplexes)
	r = b2ShapeDistance(b.tls, inputStack.Addr, uintptr(unsafe.Pointer(cache)), uintptr(unsafe.Pointer(simplexes)), simplexCapacity)
	return
}
func (b Box2D) ShapeCast(input ShapeCastPairInput) (r CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	r = b2ShapeCast(b.tls, inputStack.Addr)
	return
}
func (b Box2D) TimeOfImpact(input TOIInput) (r TOIOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	r = b2TimeOfImpact(b.tls, inputStack.Addr)
	return
}
func (b Joint) DistanceJoint_SetLength(length float32) {
	b2DistanceJoint_SetLength(b.tls, b.Id, length)
}
func (b Joint) DistanceJoint_GetLength() (r float32) {
	r = b2DistanceJoint_GetLength(b.tls, b.Id)
	return
}
func (b Joint) DistanceJoint_EnableLimit(enableLimit uint8) {
	b2DistanceJoint_EnableLimit(b.tls, b.Id, enableLimit)
}
func (b Joint) DistanceJoint_IsLimitEnabled() (r uint8) {
	r = b2DistanceJoint_IsLimitEnabled(b.tls, b.Id)
	return
}
func (b Joint) DistanceJoint_SetLengthRange(minLength float32, maxLength float32) {
	b2DistanceJoint_SetLengthRange(b.tls, b.Id, minLength, maxLength)
}
func (b Joint) DistanceJoint_GetMinLength() (r float32) {
	r = b2DistanceJoint_GetMinLength(b.tls, b.Id)
	return
}
func (b Joint) DistanceJoint_GetMaxLength() (r float32) {
	r = b2DistanceJoint_GetMaxLength(b.tls, b.Id)
	return
}
func (b Joint) DistanceJoint_GetCurrentLength() (r float32) {
	r = b2DistanceJoint_GetCurrentLength(b.tls, b.Id)
	return
}
func (b Joint) DistanceJoint_EnableSpring(enableSpring uint8) {
	b2DistanceJoint_EnableSpring(b.tls, b.Id, enableSpring)
}
func (b Joint) DistanceJoint_IsSpringEnabled() (r uint8) {
	r = b2DistanceJoint_IsSpringEnabled(b.tls, b.Id)
	return
}
func (b Joint) DistanceJoint_SetSpringHertz(hertz float32) {
	b2DistanceJoint_SetSpringHertz(b.tls, b.Id, hertz)
}
func (b Joint) DistanceJoint_SetSpringDampingRatio(dampingRatio float32) {
	b2DistanceJoint_SetSpringDampingRatio(b.tls, b.Id, dampingRatio)
}
func (b Joint) DistanceJoint_GetSpringHertz() (r float32) {
	r = b2DistanceJoint_GetSpringHertz(b.tls, b.Id)
	return
}
func (b Joint) DistanceJoint_GetSpringDampingRatio() (r float32) {
	r = b2DistanceJoint_GetSpringDampingRatio(b.tls, b.Id)
	return
}
func (b Joint) DistanceJoint_EnableMotor(enableMotor uint8) {
	b2DistanceJoint_EnableMotor(b.tls, b.Id, enableMotor)
}
func (b Joint) DistanceJoint_IsMotorEnabled() (r uint8) {
	r = b2DistanceJoint_IsMotorEnabled(b.tls, b.Id)
	return
}
func (b Joint) DistanceJoint_SetMotorSpeed(motorSpeed float32) {
	b2DistanceJoint_SetMotorSpeed(b.tls, b.Id, motorSpeed)
}
func (b Joint) DistanceJoint_GetMotorSpeed() (r float32) {
	r = b2DistanceJoint_GetMotorSpeed(b.tls, b.Id)
	return
}
func (b Joint) DistanceJoint_GetMotorForce() (r float32) {
	r = b2DistanceJoint_GetMotorForce(b.tls, b.Id)
	return
}
func (b Joint) DistanceJoint_SetMaxMotorForce(force float32) {
	b2DistanceJoint_SetMaxMotorForce(b.tls, b.Id, force)
}
func (b Joint) DistanceJoint_GetMaxMotorForce() (r float32) {
	r = b2DistanceJoint_GetMaxMotorForce(b.tls, b.Id)
	return
}
func (b Box2D) IsValidRay(input RayCastInput) (r uint8) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	r = b2IsValidRay(b.tls, inputStack.Addr)
	return
}
func (b Box2D) MakePolygon(hull Hull, radius float32) (r Polygon) {
	hullStack := copyToStack(b.tls, hull)
	defer hullStack.Free()
	r = b2MakePolygon(b.tls, hullStack.Addr, radius)
	return
}
func (b Box2D) MakeOffsetPolygon(hull Hull, position Vec2, rotation Rot) (r Polygon) {
	hullStack := copyToStack(b.tls, hull)
	defer hullStack.Free()
	r = b2MakeOffsetPolygon(b.tls, hullStack.Addr, position, rotation)
	return
}
func (b Box2D) MakeOffsetRoundedPolygon(hull Hull, position Vec2, rotation Rot, radius float32) (r Polygon) {
	hullStack := copyToStack(b.tls, hull)
	defer hullStack.Free()
	r = b2MakeOffsetRoundedPolygon(b.tls, hullStack.Addr, position, rotation, radius)
	return
}
func (b Box2D) MakeSquare(halfWidth float32) (r Polygon) {
	r = b2MakeSquare(b.tls, halfWidth)
	return
}
func (b Box2D) MakeBox(halfWidth float32, halfHeight float32) (r Polygon) {
	r = b2MakeBox(b.tls, halfWidth, halfHeight)
	return
}
func (b Box2D) MakeRoundedBox(halfWidth float32, halfHeight float32, radius float32) (r Polygon) {
	r = b2MakeRoundedBox(b.tls, halfWidth, halfHeight, radius)
	return
}
func (b Box2D) MakeOffsetBox(halfWidth float32, halfHeight float32, center Vec2, rotation Rot) (r Polygon) {
	r = b2MakeOffsetBox(b.tls, halfWidth, halfHeight, center, rotation)
	return
}
func (b Box2D) MakeOffsetRoundedBox(halfWidth float32, halfHeight float32, center Vec2, rotation Rot, radius float32) (r Polygon) {
	r = b2MakeOffsetRoundedBox(b.tls, halfWidth, halfHeight, center, rotation, radius)
	return
}
func (b Box2D) TransformPolygon(transform Transform, polygon Polygon) (r Polygon) {
	polygonStack := copyToStack(b.tls, polygon)
	defer polygonStack.Free()
	r = b2TransformPolygon(b.tls, transform, polygonStack.Addr)
	return
}
func (b Box2D) ComputeCircleMass(shape Circle, density float32) (r MassData) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2ComputeCircleMass(b.tls, shapeStack.Addr, density)
	return
}
func (b Box2D) ComputeCapsuleMass(shape Capsule, density float32) (r MassData) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2ComputeCapsuleMass(b.tls, shapeStack.Addr, density)
	return
}
func (b Box2D) ComputePolygonMass(shape Polygon, density float32) (r1 MassData) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r1 = b2ComputePolygonMass(b.tls, shapeStack.Addr, density)
	return
}
func (b Box2D) ComputeCircleAABB(shape Circle, xf Transform) (r1 AABB) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r1 = b2ComputeCircleAABB(b.tls, shapeStack.Addr, xf)
	return
}
func (b Box2D) ComputeCapsuleAABB(shape Capsule, xf Transform) (r1 AABB) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r1 = b2ComputeCapsuleAABB(b.tls, shapeStack.Addr, xf)
	return
}
func (b Box2D) ComputePolygonAABB(shape Polygon, xf Transform) (r1 AABB) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r1 = b2ComputePolygonAABB(b.tls, shapeStack.Addr, xf)
	return
}
func (b Box2D) ComputeSegmentAABB(shape Segment, xf Transform) (r AABB) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2ComputeSegmentAABB(b.tls, shapeStack.Addr, xf)
	return
}
func (b Box2D) PointInCircle(point Vec2, shape Circle) (r uint8) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2PointInCircle(b.tls, point, shapeStack.Addr)
	return
}
func (b Box2D) PointInCapsule(point Vec2, shape Capsule) (r uint8) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2PointInCapsule(b.tls, point, shapeStack.Addr)
	return
}
func (b Box2D) PointInPolygon(_point Vec2, shape Polygon) (r uint8) {
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2PointInPolygon(b.tls, _point, shapeStack.Addr)
	return
}
func (b Box2D) RayCastCircle(input RayCastInput, shape Circle) (r1 CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r1 = b2RayCastCircle(b.tls, inputStack.Addr, shapeStack.Addr)
	return
}
func (b Box2D) RayCastCapsule(input RayCastInput, shape Capsule) (r CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2RayCastCapsule(b.tls, inputStack.Addr, shapeStack.Addr)
	return
}
func (b Box2D) RayCastSegment(input RayCastInput, shape Segment, oneSided uint8) (r CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2RayCastSegment(b.tls, inputStack.Addr, shapeStack.Addr, oneSided)
	return
}
func (b Box2D) RayCastPolygon(input RayCastInput, shape Polygon) (r CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2RayCastPolygon(b.tls, inputStack.Addr, shapeStack.Addr)
	return
}
func (b Box2D) ShapeCastCircle(input ShapeCastInput, shape Circle) (r CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2ShapeCastCircle(b.tls, inputStack.Addr, shapeStack.Addr)
	return
}
func (b Box2D) ShapeCastCapsule(input ShapeCastInput, shape Capsule) (r CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2ShapeCastCapsule(b.tls, inputStack.Addr, shapeStack.Addr)
	return
}
func (b Box2D) ShapeCastSegment(input ShapeCastInput, shape Segment) (r CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2ShapeCastSegment(b.tls, inputStack.Addr, shapeStack.Addr)
	return
}
func (b Box2D) ShapeCastPolygon(input ShapeCastInput, shape Polygon) (r CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	shapeStack := copyToStack(b.tls, shape)
	defer shapeStack.Free()
	r = b2ShapeCastPolygon(b.tls, inputStack.Addr, shapeStack.Addr)
	return
}
func (b Box2D) ComputeHull(points Vec2, count int32) (r1 Hull) {
	pointsStack := copyToStack(b.tls, points)
	defer pointsStack.Free()
	r1 = b2ComputeHull(b.tls, pointsStack.Addr, count)
	return
}
func (b Box2D) ValidateHull(hull Hull) (r uint8) {
	hullStack := copyToStack(b.tls, hull)
	defer hullStack.Free()
	r = b2ValidateHull(b.tls, hullStack.Addr)
	return
}
func (b Box2D) DefaultDistanceJointDef() (r DistanceJointDef) {
	r = b2DefaultDistanceJointDef(b.tls)
	return
}
func (b Box2D) DefaultMotorJointDef() (r MotorJointDef) {
	r = b2DefaultMotorJointDef(b.tls)
	return
}
func (b Box2D) DefaultMouseJointDef() (r MouseJointDef) {
	r = b2DefaultMouseJointDef(b.tls)
	return
}
func (b Box2D) DefaultFilterJointDef() (r FilterJointDef) {
	r = b2DefaultFilterJointDef(b.tls)
	return
}
func (b Box2D) DefaultPrismaticJointDef() (r PrismaticJointDef) {
	r = b2DefaultPrismaticJointDef(b.tls)
	return
}
func (b Box2D) DefaultRevoluteJointDef() (r RevoluteJointDef) {
	r = b2DefaultRevoluteJointDef(b.tls)
	return
}
func (b Box2D) DefaultWeldJointDef() (r WeldJointDef) {
	r = b2DefaultWeldJointDef(b.tls)
	return
}
func (b Box2D) DefaultWheelJointDef() (r WheelJointDef) {
	r = b2DefaultWheelJointDef(b.tls)
	return
}
func (b Box2D) DefaultExplosionDef() (r ExplosionDef) {
	r = b2DefaultExplosionDef(b.tls)
	return
}
func (b World) CreateDistanceJoint(def DistanceJointDef) (r DistanceJoint) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r.tls = b.tls
	r.Id = b2CreateDistanceJoint(b.tls, b.Id, defStack.Addr)
	return
}
func (b World) CreateMotorJoint(def MotorJointDef) (r MotorJoint) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r.tls = b.tls
	r.Id = b2CreateMotorJoint(b.tls, b.Id, defStack.Addr)
	return
}
func (b World) CreateMouseJoint(def MouseJointDef) (r MouseJoint) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r.tls = b.tls
	r.Id = b2CreateMouseJoint(b.tls, b.Id, defStack.Addr)
	return
}
func (b World) CreateFilterJoint(def FilterJointDef) (r FilterJoint) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r.tls = b.tls
	r.Id = b2CreateFilterJoint(b.tls, b.Id, defStack.Addr)
	return
}
func (b World) CreateRevoluteJoint(def RevoluteJointDef) (r RevoluteJoint) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r.tls = b.tls
	r.Id = b2CreateRevoluteJoint(b.tls, b.Id, defStack.Addr)
	return
}
func (b World) CreatePrismaticJoint(def PrismaticJointDef) (r PrismaticJoint) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r.tls = b.tls
	r.Id = b2CreatePrismaticJoint(b.tls, b.Id, defStack.Addr)
	return
}
func (b World) CreateWeldJoint(def WeldJointDef) (r WeldJoint) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r.tls = b.tls
	r.Id = b2CreateWeldJoint(b.tls, b.Id, defStack.Addr)
	return
}
func (b World) CreateWheelJoint(def WheelJointDef) (r WheelJoint) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r.tls = b.tls
	r.Id = b2CreateWheelJoint(b.tls, b.Id, defStack.Addr)
	return
}
func (b Joint) DestroyJoint() {
	b2DestroyJoint(b.tls, b.Id)
}
func (b Joint) GetType() (r JointType) {
	r = b2Joint_GetType(b.tls, b.Id)
	return
}
func (b Joint) GetBodyA() (r Body) {
	r.tls = b.tls
	r.Id = b2Joint_GetBodyA(b.tls, b.Id)
	return
}
func (b Joint) GetBodyB() (r Body) {
	r.tls = b.tls
	r.Id = b2Joint_GetBodyB(b.tls, b.Id)
	return
}
func (b Joint) GetWorld() (r World) {
	r.tls = b.tls
	r.Id = b2Joint_GetWorld(b.tls, b.Id)
	return
}
func (b Joint) SetLocalAnchorA(localAnchor Vec2) {
	b2Joint_SetLocalAnchorA(b.tls, b.Id, localAnchor)
}
func (b Joint) GetLocalAnchorA() (r Vec2) {
	r = b2Joint_GetLocalAnchorA(b.tls, b.Id)
	return
}
func (b Joint) SetLocalAnchorB(localAnchor Vec2) {
	b2Joint_SetLocalAnchorB(b.tls, b.Id, localAnchor)
}
func (b Joint) GetLocalAnchorB() (r Vec2) {
	r = b2Joint_GetLocalAnchorB(b.tls, b.Id)
	return
}
func (b Joint) SetReferenceAngle(angleInRadians float32) {
	b2Joint_SetReferenceAngle(b.tls, b.Id, angleInRadians)
}
func (b Joint) GetReferenceAngle() (r float32) {
	r = b2Joint_GetReferenceAngle(b.tls, b.Id)
	return
}
func (b Joint) SetLocalAxisA(localAxis Vec2) {
	b2Joint_SetLocalAxisA(b.tls, b.Id, localAxis)
}
func (b Joint) GetLocalAxisA() (r Vec2) {
	r = b2Joint_GetLocalAxisA(b.tls, b.Id)
	return
}
func (b Joint) SetCollideConnected(shouldCollide uint8) {
	b2Joint_SetCollideConnected(b.tls, b.Id, shouldCollide)
}
func (b Joint) GetCollideConnected() (r uint8) {
	r = b2Joint_GetCollideConnected(b.tls, b.Id)
	return
}
func (b Joint) SetUserData(userData uintptr) {
	b2Joint_SetUserData(b.tls, b.Id, userData)
}
func (b Joint) GetUserData() (r uintptr) {
	r = b2Joint_GetUserData(b.tls, b.Id)
	return
}
func (b Joint) WakeBodies() {
	b2Joint_WakeBodies(b.tls, b.Id)
}
func (b Joint) GetConstraintForce() (r Vec2) {
	r = b2Joint_GetConstraintForce(b.tls, b.Id)
	return
}
func (b Joint) GetConstraintTorque() (r float32) {
	r = b2Joint_GetConstraintTorque(b.tls, b.Id)
	return
}
func (b Joint) GetLinearSeparation() (r float32) {
	r = b2Joint_GetLinearSeparation(b.tls, b.Id)
	return
}
func (b Joint) GetAngularSeparation() (r float32) {
	r = b2Joint_GetAngularSeparation(b.tls, b.Id)
	return
}
func (b Joint) GetConstraintTuning(hertz uintptr, dampingRatio uintptr) {
	b2Joint_GetConstraintTuning(b.tls, b.Id, hertz, dampingRatio)
}
func (b Joint) SetConstraintTuning(hertz float32, dampingRatio float32) {
	b2Joint_SetConstraintTuning(b.tls, b.Id, hertz, dampingRatio)
}
func (b Box2D) CollideCircles(circleA Circle, xfA Transform, circleB Circle, xfB Transform) (r Manifold) {
	circleAStack := copyToStack(b.tls, circleA)
	defer circleAStack.Free()
	circleBStack := copyToStack(b.tls, circleB)
	defer circleBStack.Free()
	r = b2CollideCircles(b.tls, circleAStack.Addr, xfA, circleBStack.Addr, xfB)
	return
}
func (b Box2D) CollideCapsuleAndCircle(capsuleA Capsule, xfA Transform, circleB Circle, xfB Transform) (r Manifold) {
	capsuleAStack := copyToStack(b.tls, capsuleA)
	defer capsuleAStack.Free()
	circleBStack := copyToStack(b.tls, circleB)
	defer circleBStack.Free()
	r = b2CollideCapsuleAndCircle(b.tls, capsuleAStack.Addr, xfA, circleBStack.Addr, xfB)
	return
}
func (b Box2D) CollidePolygonAndCircle(polygonA Polygon, xfA Transform, circleB Circle, xfB Transform) (r Manifold) {
	polygonAStack := copyToStack(b.tls, polygonA)
	defer polygonAStack.Free()
	circleBStack := copyToStack(b.tls, circleB)
	defer circleBStack.Free()
	r = b2CollidePolygonAndCircle(b.tls, polygonAStack.Addr, xfA, circleBStack.Addr, xfB)
	return
}
func (b Box2D) CollideCapsules(capsuleA Capsule, xfA Transform, capsuleB Capsule, xfB Transform) (r Manifold) {
	capsuleAStack := copyToStack(b.tls, capsuleA)
	defer capsuleAStack.Free()
	capsuleBStack := copyToStack(b.tls, capsuleB)
	defer capsuleBStack.Free()
	r = b2CollideCapsules(b.tls, capsuleAStack.Addr, xfA, capsuleBStack.Addr, xfB)
	return
}
func (b Box2D) CollideSegmentAndCapsule(segmentA Segment, xfA Transform, capsuleB Capsule, xfB Transform) (r Manifold) {
	segmentAStack := copyToStack(b.tls, segmentA)
	defer segmentAStack.Free()
	capsuleBStack := copyToStack(b.tls, capsuleB)
	defer capsuleBStack.Free()
	r = b2CollideSegmentAndCapsule(b.tls, segmentAStack.Addr, xfA, capsuleBStack.Addr, xfB)
	return
}
func (b Box2D) CollidePolygonAndCapsule(polygonA Polygon, xfA Transform, capsuleB Capsule, xfB Transform) (r Manifold) {
	polygonAStack := copyToStack(b.tls, polygonA)
	defer polygonAStack.Free()
	capsuleBStack := copyToStack(b.tls, capsuleB)
	defer capsuleBStack.Free()
	r = b2CollidePolygonAndCapsule(b.tls, polygonAStack.Addr, xfA, capsuleBStack.Addr, xfB)
	return
}
func (b Box2D) CollidePolygons(polygonA Polygon, xfA Transform, polygonB Polygon, xfB Transform) (r Manifold) {
	polygonAStack := copyToStack(b.tls, polygonA)
	defer polygonAStack.Free()
	polygonBStack := copyToStack(b.tls, polygonB)
	defer polygonBStack.Free()
	r = b2CollidePolygons(b.tls, polygonAStack.Addr, xfA, polygonBStack.Addr, xfB)
	return
}
func (b Box2D) CollideSegmentAndCircle(segmentA Segment, xfA Transform, circleB Circle, xfB Transform) (r Manifold) {
	segmentAStack := copyToStack(b.tls, segmentA)
	defer segmentAStack.Free()
	circleBStack := copyToStack(b.tls, circleB)
	defer circleBStack.Free()
	r = b2CollideSegmentAndCircle(b.tls, segmentAStack.Addr, xfA, circleBStack.Addr, xfB)
	return
}
func (b Box2D) CollideSegmentAndPolygon(segmentA Segment, xfA Transform, polygonB Polygon, xfB Transform) (r Manifold) {
	segmentAStack := copyToStack(b.tls, segmentA)
	defer segmentAStack.Free()
	polygonBStack := copyToStack(b.tls, polygonB)
	defer polygonBStack.Free()
	r = b2CollideSegmentAndPolygon(b.tls, segmentAStack.Addr, xfA, polygonBStack.Addr, xfB)
	return
}
func (b Box2D) CollideChainSegmentAndCircle(segmentA ChainSegment, xfA Transform, circleB Circle, xfB Transform) (r Manifold) {
	segmentAStack := copyToStack(b.tls, segmentA)
	defer segmentAStack.Free()
	circleBStack := copyToStack(b.tls, circleB)
	defer circleBStack.Free()
	r = b2CollideChainSegmentAndCircle(b.tls, segmentAStack.Addr, xfA, circleBStack.Addr, xfB)
	return
}
func (b Box2D) CollideChainSegmentAndCapsule(segmentA ChainSegment, xfA Transform, capsuleB Capsule, xfB Transform, cache uintptr) (r Manifold) {
	segmentAStack := copyToStack(b.tls, segmentA)
	defer segmentAStack.Free()
	capsuleBStack := copyToStack(b.tls, capsuleB)
	defer capsuleBStack.Free()
	r = b2CollideChainSegmentAndCapsule(b.tls, segmentAStack.Addr, xfA, capsuleBStack.Addr, xfB, cache)
	return
}
func (b Box2D) CollideChainSegmentAndPolygon(segmentA ChainSegment, xfA Transform, polygonB Polygon, xfB Transform, cache uintptr) (r Manifold) {
	segmentAStack := copyToStack(b.tls, segmentA)
	defer segmentAStack.Free()
	polygonBStack := copyToStack(b.tls, polygonB)
	defer polygonBStack.Free()
	r = b2CollideChainSegmentAndPolygon(b.tls, segmentAStack.Addr, xfA, polygonBStack.Addr, xfB, cache)
	return
}
func (b Box2D) IsValidFloat(a float32) (r uint8) {
	r = b2IsValidFloat(b.tls, a)
	return
}
func (b Box2D) IsValidVec2(v Vec2) (r uint8) {
	r = b2IsValidVec2(b.tls, v)
	return
}
func (b Box2D) IsValidRotation(q1 Rot) (r uint8) {
	r = b2IsValidRotation(b.tls, q1)
	return
}
func (b Box2D) IsValidPlane(a3 Plane) (r uint8) {
	r = b2IsValidPlane(b.tls, a3)
	return
}
func (b Box2D) Atan2(y float32, x float32) (r1 float32) {
	r1 = b2Atan2(b.tls, y, x)
	return
}
func (b Box2D) ComputeCosSin(radians1 float32) (r CosSin) {
	r = b2ComputeCosSin(b.tls, radians1)
	return
}
func (b Box2D) ComputeRotationBetweenUnitVectors(v1 Vec2, v2 Vec2) (r Rot) {
	r = b2ComputeRotationBetweenUnitVectors(b.tls, v1, v2)
	return
}
func (b Joint) MotorJoint_SetLinearOffset(linearOffset Vec2) {
	b2MotorJoint_SetLinearOffset(b.tls, b.Id, linearOffset)
}
func (b Joint) MotorJoint_GetLinearOffset() (r Vec2) {
	r = b2MotorJoint_GetLinearOffset(b.tls, b.Id)
	return
}
func (b Joint) MotorJoint_SetAngularOffset(angularOffset float32) {
	b2MotorJoint_SetAngularOffset(b.tls, b.Id, angularOffset)
}
func (b Joint) MotorJoint_GetAngularOffset() (r float32) {
	r = b2MotorJoint_GetAngularOffset(b.tls, b.Id)
	return
}
func (b Joint) MotorJoint_SetMaxForce(maxForce float32) {
	b2MotorJoint_SetMaxForce(b.tls, b.Id, maxForce)
}
func (b Joint) MotorJoint_GetMaxForce() (r float32) {
	r = b2MotorJoint_GetMaxForce(b.tls, b.Id)
	return
}
func (b Joint) MotorJoint_SetMaxTorque(maxTorque float32) {
	b2MotorJoint_SetMaxTorque(b.tls, b.Id, maxTorque)
}
func (b Joint) MotorJoint_GetMaxTorque() (r float32) {
	r = b2MotorJoint_GetMaxTorque(b.tls, b.Id)
	return
}
func (b Joint) MotorJoint_SetCorrectionFactor(correctionFactor float32) {
	b2MotorJoint_SetCorrectionFactor(b.tls, b.Id, correctionFactor)
}
func (b Joint) MotorJoint_GetCorrectionFactor() (r float32) {
	r = b2MotorJoint_GetCorrectionFactor(b.tls, b.Id)
	return
}
func (b Joint) MouseJoint_SetTarget(target Vec2) {
	b2MouseJoint_SetTarget(b.tls, b.Id, target)
}
func (b Joint) MouseJoint_GetTarget() (r Vec2) {
	r = b2MouseJoint_GetTarget(b.tls, b.Id)
	return
}
func (b Joint) MouseJoint_SetSpringHertz(hertz float32) {
	b2MouseJoint_SetSpringHertz(b.tls, b.Id, hertz)
}
func (b Joint) MouseJoint_GetSpringHertz() (r float32) {
	r = b2MouseJoint_GetSpringHertz(b.tls, b.Id)
	return
}
func (b Joint) MouseJoint_SetSpringDampingRatio(dampingRatio float32) {
	b2MouseJoint_SetSpringDampingRatio(b.tls, b.Id, dampingRatio)
}
func (b Joint) MouseJoint_GetSpringDampingRatio() (r float32) {
	r = b2MouseJoint_GetSpringDampingRatio(b.tls, b.Id)
	return
}
func (b Joint) MouseJoint_SetMaxForce(maxForce float32) {
	b2MouseJoint_SetMaxForce(b.tls, b.Id, maxForce)
}
func (b Joint) MouseJoint_GetMaxForce() (r float32) {
	r = b2MouseJoint_GetMaxForce(b.tls, b.Id)
	return
}
func (b Box2D) SolvePlanes(targetDelta Vec2, planes *CollisionPlane, count int32) (r PlaneSolverResult) {
	defer runtime.KeepAlive(planes)
	escapes(planes)
	r = b2SolvePlanes(b.tls, targetDelta, uintptr(unsafe.Pointer(planes)), count)
	return
}
func (b Box2D) ClipVector(vector Vec2, planes CollisionPlane, count int32) (r Vec2) {
	planesStack := copyToStack(b.tls, planes)
	defer planesStack.Free()
	r = b2ClipVector(b.tls, vector, planesStack.Addr, count)
	return
}
func (b Joint) PrismaticJoint_EnableSpring(enableSpring uint8) {
	b2PrismaticJoint_EnableSpring(b.tls, b.Id, enableSpring)
}
func (b Joint) PrismaticJoint_IsSpringEnabled() (r uint8) {
	r = b2PrismaticJoint_IsSpringEnabled(b.tls, b.Id)
	return
}
func (b Joint) PrismaticJoint_SetSpringHertz(hertz float32) {
	b2PrismaticJoint_SetSpringHertz(b.tls, b.Id, hertz)
}
func (b Joint) PrismaticJoint_GetSpringHertz() (r float32) {
	r = b2PrismaticJoint_GetSpringHertz(b.tls, b.Id)
	return
}
func (b Joint) PrismaticJoint_SetSpringDampingRatio(dampingRatio float32) {
	b2PrismaticJoint_SetSpringDampingRatio(b.tls, b.Id, dampingRatio)
}
func (b Joint) PrismaticJoint_GetSpringDampingRatio() (r float32) {
	r = b2PrismaticJoint_GetSpringDampingRatio(b.tls, b.Id)
	return
}
func (b Joint) PrismaticJoint_SetTargetTranslation(translation float32) {
	b2PrismaticJoint_SetTargetTranslation(b.tls, b.Id, translation)
}
func (b Joint) PrismaticJoint_GetTargetTranslation() (r float32) {
	r = b2PrismaticJoint_GetTargetTranslation(b.tls, b.Id)
	return
}
func (b Joint) PrismaticJoint_EnableLimit(enableLimit uint8) {
	b2PrismaticJoint_EnableLimit(b.tls, b.Id, enableLimit)
}
func (b Joint) PrismaticJoint_IsLimitEnabled() (r uint8) {
	r = b2PrismaticJoint_IsLimitEnabled(b.tls, b.Id)
	return
}
func (b Joint) PrismaticJoint_GetLowerLimit() (r float32) {
	r = b2PrismaticJoint_GetLowerLimit(b.tls, b.Id)
	return
}
func (b Joint) PrismaticJoint_GetUpperLimit() (r float32) {
	r = b2PrismaticJoint_GetUpperLimit(b.tls, b.Id)
	return
}
func (b Joint) PrismaticJoint_SetLimits(lower float32, upper float32) {
	b2PrismaticJoint_SetLimits(b.tls, b.Id, lower, upper)
}
func (b Joint) PrismaticJoint_EnableMotor(enableMotor uint8) {
	b2PrismaticJoint_EnableMotor(b.tls, b.Id, enableMotor)
}
func (b Joint) PrismaticJoint_IsMotorEnabled() (r uint8) {
	r = b2PrismaticJoint_IsMotorEnabled(b.tls, b.Id)
	return
}
func (b Joint) PrismaticJoint_SetMotorSpeed(motorSpeed float32) {
	b2PrismaticJoint_SetMotorSpeed(b.tls, b.Id, motorSpeed)
}
func (b Joint) PrismaticJoint_GetMotorSpeed() (r float32) {
	r = b2PrismaticJoint_GetMotorSpeed(b.tls, b.Id)
	return
}
func (b Joint) PrismaticJoint_GetMotorForce() (r float32) {
	r = b2PrismaticJoint_GetMotorForce(b.tls, b.Id)
	return
}
func (b Joint) PrismaticJoint_SetMaxMotorForce(force float32) {
	b2PrismaticJoint_SetMaxMotorForce(b.tls, b.Id, force)
}
func (b Joint) PrismaticJoint_GetMaxMotorForce() (r float32) {
	r = b2PrismaticJoint_GetMaxMotorForce(b.tls, b.Id)
	return
}
func (b Joint) PrismaticJoint_GetTranslation() (r float32) {
	r = b2PrismaticJoint_GetTranslation(b.tls, b.Id)
	return
}
func (b Joint) PrismaticJoint_GetSpeed() (r float32) {
	r = b2PrismaticJoint_GetSpeed(b.tls, b.Id)
	return
}
func (b Joint) RevoluteJoint_EnableSpring(enableSpring uint8) {
	b2RevoluteJoint_EnableSpring(b.tls, b.Id, enableSpring)
}
func (b Joint) RevoluteJoint_IsSpringEnabled() (r uint8) {
	r = b2RevoluteJoint_IsSpringEnabled(b.tls, b.Id)
	return
}
func (b Joint) RevoluteJoint_SetSpringHertz(hertz float32) {
	b2RevoluteJoint_SetSpringHertz(b.tls, b.Id, hertz)
}
func (b Joint) RevoluteJoint_GetSpringHertz() (r float32) {
	r = b2RevoluteJoint_GetSpringHertz(b.tls, b.Id)
	return
}
func (b Joint) RevoluteJoint_SetSpringDampingRatio(dampingRatio float32) {
	b2RevoluteJoint_SetSpringDampingRatio(b.tls, b.Id, dampingRatio)
}
func (b Joint) RevoluteJoint_GetSpringDampingRatio() (r float32) {
	r = b2RevoluteJoint_GetSpringDampingRatio(b.tls, b.Id)
	return
}
func (b Joint) RevoluteJoint_SetTargetAngle(angle float32) {
	b2RevoluteJoint_SetTargetAngle(b.tls, b.Id, angle)
}
func (b Joint) RevoluteJoint_GetTargetAngle() (r float32) {
	r = b2RevoluteJoint_GetTargetAngle(b.tls, b.Id)
	return
}
func (b Joint) RevoluteJoint_GetAngle() (r float32) {
	r = b2RevoluteJoint_GetAngle(b.tls, b.Id)
	return
}
func (b Joint) RevoluteJoint_EnableLimit(enableLimit uint8) {
	b2RevoluteJoint_EnableLimit(b.tls, b.Id, enableLimit)
}
func (b Joint) RevoluteJoint_IsLimitEnabled() (r uint8) {
	r = b2RevoluteJoint_IsLimitEnabled(b.tls, b.Id)
	return
}
func (b Joint) RevoluteJoint_GetLowerLimit() (r float32) {
	r = b2RevoluteJoint_GetLowerLimit(b.tls, b.Id)
	return
}
func (b Joint) RevoluteJoint_GetUpperLimit() (r float32) {
	r = b2RevoluteJoint_GetUpperLimit(b.tls, b.Id)
	return
}
func (b Joint) RevoluteJoint_SetLimits(lower float32, upper float32) {
	b2RevoluteJoint_SetLimits(b.tls, b.Id, lower, upper)
}
func (b Joint) RevoluteJoint_EnableMotor(enableMotor uint8) {
	b2RevoluteJoint_EnableMotor(b.tls, b.Id, enableMotor)
}
func (b Joint) RevoluteJoint_IsMotorEnabled() (r uint8) {
	r = b2RevoluteJoint_IsMotorEnabled(b.tls, b.Id)
	return
}
func (b Joint) RevoluteJoint_SetMotorSpeed(motorSpeed float32) {
	b2RevoluteJoint_SetMotorSpeed(b.tls, b.Id, motorSpeed)
}
func (b Joint) RevoluteJoint_GetMotorSpeed() (r float32) {
	r = b2RevoluteJoint_GetMotorSpeed(b.tls, b.Id)
	return
}
func (b Joint) RevoluteJoint_GetMotorTorque() (r float32) {
	r = b2RevoluteJoint_GetMotorTorque(b.tls, b.Id)
	return
}
func (b Joint) RevoluteJoint_SetMaxMotorTorque(torque float32) {
	b2RevoluteJoint_SetMaxMotorTorque(b.tls, b.Id, torque)
}
func (b Joint) RevoluteJoint_GetMaxMotorTorque() (r float32) {
	r = b2RevoluteJoint_GetMaxMotorTorque(b.tls, b.Id)
	return
}
func (b Body) CreateCircleShape(def ShapeDef, circle Circle) (r Shape) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	circleStack := copyToStack(b.tls, circle)
	defer circleStack.Free()
	r.tls = b.tls
	r.Id = b2CreateCircleShape(b.tls, b.Id, defStack.Addr, circleStack.Addr)
	return
}
func (b Body) CreateCapsuleShape(def ShapeDef, capsule Capsule) (r Shape) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	capsuleStack := copyToStack(b.tls, capsule)
	defer capsuleStack.Free()
	r.tls = b.tls
	r.Id = b2CreateCapsuleShape(b.tls, b.Id, defStack.Addr, capsuleStack.Addr)
	return
}
func (b Body) CreatePolygonShape(def ShapeDef, polygon Polygon) (r Shape) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	polygonStack := copyToStack(b.tls, polygon)
	defer polygonStack.Free()
	r.tls = b.tls
	r.Id = b2CreatePolygonShape(b.tls, b.Id, defStack.Addr, polygonStack.Addr)
	return
}
func (b Body) CreateSegmentShape(def ShapeDef, segment Segment) (r Shape) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	segmentStack := copyToStack(b.tls, segment)
	defer segmentStack.Free()
	r.tls = b.tls
	r.Id = b2CreateSegmentShape(b.tls, b.Id, defStack.Addr, segmentStack.Addr)
	return
}
func (b Shape) DestroyShape(updateBodyMass uint8) {
	b2DestroyShape(b.tls, b.Id, updateBodyMass)
}
func (b Body) CreateChain(def ChainDef) (r Chain) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r.tls = b.tls
	r.Id = b2CreateChain(b.tls, b.Id, defStack.Addr)
	return
}
func (b Chain) DestroyChain() {
	b2DestroyChain(b.tls, b.Id)
}
func (b Chain) GetWorld() (r World) {
	r.tls = b.tls
	r.Id = b2Chain_GetWorld(b.tls, b.Id)
	return
}
func (b Chain) GetSegmentCount() (r int32) {
	r = b2Chain_GetSegmentCount(b.tls, b.Id)
	return
}
func (b Shape) GetBody() (r Body) {
	r.tls = b.tls
	r.Id = b2Shape_GetBody(b.tls, b.Id)
	return
}
func (b Shape) GetWorld() (r World) {
	r.tls = b.tls
	r.Id = b2Shape_GetWorld(b.tls, b.Id)
	return
}
func (b Shape) SetUserData(userData uintptr) {
	b2Shape_SetUserData(b.tls, b.Id, userData)
}
func (b Shape) GetUserData() (r uintptr) {
	r = b2Shape_GetUserData(b.tls, b.Id)
	return
}
func (b Shape) IsSensor() (r uint8) {
	r = b2Shape_IsSensor(b.tls, b.Id)
	return
}
func (b Shape) TestPoint(point Vec2) (r uint8) {
	r = b2Shape_TestPoint(b.tls, b.Id, point)
	return
}
func (b Shape) RayCast(input RayCastInput) (r CastOutput) {
	inputStack := copyToStack(b.tls, input)
	defer inputStack.Free()
	r = b2Shape_RayCast(b.tls, b.Id, inputStack.Addr)
	return
}
func (b Shape) SetDensity(density float32, updateBodyMass uint8) {
	b2Shape_SetDensity(b.tls, b.Id, density, updateBodyMass)
}
func (b Shape) GetDensity() (r float32) {
	r = b2Shape_GetDensity(b.tls, b.Id)
	return
}
func (b Shape) SetFriction(friction float32) {
	b2Shape_SetFriction(b.tls, b.Id, friction)
}
func (b Shape) GetFriction() (r float32) {
	r = b2Shape_GetFriction(b.tls, b.Id)
	return
}
func (b Shape) SetRestitution(restitution float32) {
	b2Shape_SetRestitution(b.tls, b.Id, restitution)
}
func (b Shape) GetRestitution() (r float32) {
	r = b2Shape_GetRestitution(b.tls, b.Id)
	return
}
func (b Shape) SetMaterial(material int32) {
	b2Shape_SetMaterial(b.tls, b.Id, material)
}
func (b Shape) GetMaterial() (r int32) {
	r = b2Shape_GetMaterial(b.tls, b.Id)
	return
}
func (b Shape) GetSurfaceMaterial() (r SurfaceMaterial) {
	r = b2Shape_GetSurfaceMaterial(b.tls, b.Id)
	return
}
func (b Shape) SetSurfaceMaterial(surfaceMaterial SurfaceMaterial) {
	b2Shape_SetSurfaceMaterial(b.tls, b.Id, surfaceMaterial)
}
func (b Shape) GetFilter() (r Filter) {
	r = b2Shape_GetFilter(b.tls, b.Id)
	return
}
func (b Shape) SetFilter(filter Filter) {
	b2Shape_SetFilter(b.tls, b.Id, filter)
}
func (b Shape) EnableSensorEvents(flag uint8) {
	b2Shape_EnableSensorEvents(b.tls, b.Id, flag)
}
func (b Shape) AreSensorEventsEnabled() (r uint8) {
	r = b2Shape_AreSensorEventsEnabled(b.tls, b.Id)
	return
}
func (b Shape) EnableContactEvents(flag uint8) {
	b2Shape_EnableContactEvents(b.tls, b.Id, flag)
}
func (b Shape) AreContactEventsEnabled() (r uint8) {
	r = b2Shape_AreContactEventsEnabled(b.tls, b.Id)
	return
}
func (b Shape) EnablePreSolveEvents(flag uint8) {
	b2Shape_EnablePreSolveEvents(b.tls, b.Id, flag)
}
func (b Shape) ArePreSolveEventsEnabled() (r uint8) {
	r = b2Shape_ArePreSolveEventsEnabled(b.tls, b.Id)
	return
}
func (b Shape) EnableHitEvents(flag uint8) {
	b2Shape_EnableHitEvents(b.tls, b.Id, flag)
}
func (b Shape) AreHitEventsEnabled() (r uint8) {
	r = b2Shape_AreHitEventsEnabled(b.tls, b.Id)
	return
}
func (b Shape) GetType() (r ShapeType) {
	r = b2Shape_GetType(b.tls, b.Id)
	return
}
func (b Shape) GetCircle() (r Circle) {
	r = b2Shape_GetCircle(b.tls, b.Id)
	return
}
func (b Shape) GetSegment() (r Segment) {
	r = b2Shape_GetSegment(b.tls, b.Id)
	return
}
func (b Shape) GetChainSegment() (r ChainSegment) {
	r = b2Shape_GetChainSegment(b.tls, b.Id)
	return
}
func (b Shape) GetCapsule() (r Capsule) {
	r = b2Shape_GetCapsule(b.tls, b.Id)
	return
}
func (b Shape) GetPolygon() (r Polygon) {
	r = b2Shape_GetPolygon(b.tls, b.Id)
	return
}
func (b Shape) SetCircle(circle Circle) {
	circleStack := copyToStack(b.tls, circle)
	defer circleStack.Free()
	b2Shape_SetCircle(b.tls, b.Id, circleStack.Addr)
}
func (b Shape) SetCapsule(capsule Capsule) {
	capsuleStack := copyToStack(b.tls, capsule)
	defer capsuleStack.Free()
	b2Shape_SetCapsule(b.tls, b.Id, capsuleStack.Addr)
}
func (b Shape) SetSegment(segment Segment) {
	segmentStack := copyToStack(b.tls, segment)
	defer segmentStack.Free()
	b2Shape_SetSegment(b.tls, b.Id, segmentStack.Addr)
}
func (b Shape) SetPolygon(polygon Polygon) {
	polygonStack := copyToStack(b.tls, polygon)
	defer polygonStack.Free()
	b2Shape_SetPolygon(b.tls, b.Id, polygonStack.Addr)
}
func (b Shape) GetParentChain() (r Chain) {
	r.tls = b.tls
	r.Id = b2Shape_GetParentChain(b.tls, b.Id)
	return
}
func (b Chain) SetFriction(friction float32) {
	b2Chain_SetFriction(b.tls, b.Id, friction)
}
func (b Chain) GetFriction() (r float32) {
	r = b2Chain_GetFriction(b.tls, b.Id)
	return
}
func (b Chain) SetRestitution(restitution float32) {
	b2Chain_SetRestitution(b.tls, b.Id, restitution)
}
func (b Chain) GetRestitution() (r float32) {
	r = b2Chain_GetRestitution(b.tls, b.Id)
	return
}
func (b Chain) SetMaterial(material int32) {
	b2Chain_SetMaterial(b.tls, b.Id, material)
}
func (b Chain) GetMaterial() (r int32) {
	r = b2Chain_GetMaterial(b.tls, b.Id)
	return
}
func (b Shape) GetContactCapacity() (r int32) {
	r = b2Shape_GetContactCapacity(b.tls, b.Id)
	return
}
func (b Shape) GetContactData(contactData *ContactData, capacity int32) (r int32) {
	defer runtime.KeepAlive(contactData)
	escapes(contactData)
	r = b2Shape_GetContactData(b.tls, b.Id, uintptr(unsafe.Pointer(contactData)), capacity)
	return
}
func (b Shape) GetSensorCapacity() (r int32) {
	r = b2Shape_GetSensorCapacity(b.tls, b.Id)
	return
}
func (b Shape) GetAABB() (r AABB) {
	r = b2Shape_GetAABB(b.tls, b.Id)
	return
}
func (b Shape) GetMassData() (r MassData) {
	r = b2Shape_GetMassData(b.tls, b.Id)
	return
}
func (b Shape) GetClosestPoint(_target Vec2) (r Vec2) {
	r = b2Shape_GetClosestPoint(b.tls, b.Id, _target)
	return
}
func (b Box2D) GetTicks() (r uint64) {
	r = b2GetTicks(b.tls)
	return
}
func (b Box2D) GetMilliseconds(ticks uint64) (r float32) {
	r = b2GetMilliseconds(b.tls, ticks)
	return
}
func (b Box2D) GetMillisecondsAndReset(ticks uintptr) (r float32) {
	r = b2GetMillisecondsAndReset(b.tls, ticks)
	return
}
func (b Box2D) Yield() {
	b2Yield(b.tls)
}
func (b Box2D) Hash(hash uint32, data uintptr, count int32) (r uint32) {
	r = b2Hash(b.tls, hash, data, count)
	return
}
func (b Box2D) DefaultWorldDef() (r WorldDef) {
	r = b2DefaultWorldDef(b.tls)
	return
}
func (b Box2D) DefaultBodyDef() (r BodyDef) {
	r = b2DefaultBodyDef(b.tls)
	return
}
func (b Box2D) DefaultFilter() (r Filter) {
	r = b2DefaultFilter(b.tls)
	return
}
func (b Box2D) DefaultQueryFilter() (r QueryFilter) {
	r = b2DefaultQueryFilter(b.tls)
	return
}
func (b Box2D) DefaultShapeDef() (r ShapeDef) {
	r = b2DefaultShapeDef(b.tls)
	return
}
func (b Box2D) DefaultSurfaceMaterial() (r SurfaceMaterial) {
	r = b2DefaultSurfaceMaterial(b.tls)
	return
}
func (b Box2D) DefaultChainDef() (r ChainDef) {
	r = b2DefaultChainDef(b.tls)
	return
}
func (b Box2D) DefaultDebugDraw() (r DebugDraw) {
	r = b2DefaultDebugDraw(b.tls)
	return
}
func (b Joint) WeldJoint_SetLinearHertz(hertz float32) {
	b2WeldJoint_SetLinearHertz(b.tls, b.Id, hertz)
}
func (b Joint) WeldJoint_GetLinearHertz() (r float32) {
	r = b2WeldJoint_GetLinearHertz(b.tls, b.Id)
	return
}
func (b Joint) WeldJoint_SetLinearDampingRatio(dampingRatio float32) {
	b2WeldJoint_SetLinearDampingRatio(b.tls, b.Id, dampingRatio)
}
func (b Joint) WeldJoint_GetLinearDampingRatio() (r float32) {
	r = b2WeldJoint_GetLinearDampingRatio(b.tls, b.Id)
	return
}
func (b Joint) WeldJoint_SetAngularHertz(hertz float32) {
	b2WeldJoint_SetAngularHertz(b.tls, b.Id, hertz)
}
func (b Joint) WeldJoint_GetAngularHertz() (r float32) {
	r = b2WeldJoint_GetAngularHertz(b.tls, b.Id)
	return
}
func (b Joint) WeldJoint_SetAngularDampingRatio(dampingRatio float32) {
	b2WeldJoint_SetAngularDampingRatio(b.tls, b.Id, dampingRatio)
}
func (b Joint) WeldJoint_GetAngularDampingRatio() (r float32) {
	r = b2WeldJoint_GetAngularDampingRatio(b.tls, b.Id)
	return
}
func (b Joint) WheelJoint_EnableSpring(enableSpring uint8) {
	b2WheelJoint_EnableSpring(b.tls, b.Id, enableSpring)
}
func (b Joint) WheelJoint_IsSpringEnabled() (r uint8) {
	r = b2WheelJoint_IsSpringEnabled(b.tls, b.Id)
	return
}
func (b Joint) WheelJoint_SetSpringHertz(hertz float32) {
	b2WheelJoint_SetSpringHertz(b.tls, b.Id, hertz)
}
func (b Joint) WheelJoint_GetSpringHertz() (r float32) {
	r = b2WheelJoint_GetSpringHertz(b.tls, b.Id)
	return
}
func (b Joint) WheelJoint_SetSpringDampingRatio(dampingRatio float32) {
	b2WheelJoint_SetSpringDampingRatio(b.tls, b.Id, dampingRatio)
}
func (b Joint) WheelJoint_GetSpringDampingRatio() (r float32) {
	r = b2WheelJoint_GetSpringDampingRatio(b.tls, b.Id)
	return
}
func (b Joint) WheelJoint_EnableLimit(enableLimit uint8) {
	b2WheelJoint_EnableLimit(b.tls, b.Id, enableLimit)
}
func (b Joint) WheelJoint_IsLimitEnabled() (r uint8) {
	r = b2WheelJoint_IsLimitEnabled(b.tls, b.Id)
	return
}
func (b Joint) WheelJoint_GetLowerLimit() (r float32) {
	r = b2WheelJoint_GetLowerLimit(b.tls, b.Id)
	return
}
func (b Joint) WheelJoint_GetUpperLimit() (r float32) {
	r = b2WheelJoint_GetUpperLimit(b.tls, b.Id)
	return
}
func (b Joint) WheelJoint_SetLimits(lower float32, upper float32) {
	b2WheelJoint_SetLimits(b.tls, b.Id, lower, upper)
}
func (b Joint) WheelJoint_EnableMotor(enableMotor uint8) {
	b2WheelJoint_EnableMotor(b.tls, b.Id, enableMotor)
}
func (b Joint) WheelJoint_IsMotorEnabled() (r uint8) {
	r = b2WheelJoint_IsMotorEnabled(b.tls, b.Id)
	return
}
func (b Joint) WheelJoint_SetMotorSpeed(motorSpeed float32) {
	b2WheelJoint_SetMotorSpeed(b.tls, b.Id, motorSpeed)
}
func (b Joint) WheelJoint_GetMotorSpeed() (r float32) {
	r = b2WheelJoint_GetMotorSpeed(b.tls, b.Id)
	return
}
func (b Joint) WheelJoint_GetMotorTorque() (r float32) {
	r = b2WheelJoint_GetMotorTorque(b.tls, b.Id)
	return
}
func (b Joint) WheelJoint_SetMaxMotorTorque(torque float32) {
	b2WheelJoint_SetMaxMotorTorque(b.tls, b.Id, torque)
}
func (b Joint) WheelJoint_GetMaxMotorTorque() (r float32) {
	r = b2WheelJoint_GetMaxMotorTorque(b.tls, b.Id)
	return
}
func (b Box2D) CreateWorld(def WorldDef) (r World) {
	defStack := copyToStack(b.tls, def)
	defer defStack.Free()
	r.tls = b.tls
	r.Id = b2CreateWorld(b.tls, defStack.Addr)
	return
}
func (b World) DestroyWorld() {
	b2DestroyWorld(b.tls, b.Id)
}
func (b World) Step(timeStep float32, subStepCount int32) {
	b2World_Step(b.tls, b.Id, timeStep, subStepCount)
}
func (b World) Draw(draw *DebugDraw) {
	defer runtime.KeepAlive(draw)
	escapes(draw)
	b2World_Draw(b.tls, b.Id, uintptr(unsafe.Pointer(draw)))
}
func (b World) IsValid() (r uint8) {
	r = b2World_IsValid(b.tls, b.Id)
	return
}
func (b Body) IsValid() (r uint8) {
	r = b2Body_IsValid(b.tls, b.Id)
	return
}
func (b Shape) IsValid() (r uint8) {
	r = b2Shape_IsValid(b.tls, b.Id)
	return
}
func (b Chain) IsValid() (r uint8) {
	r = b2Chain_IsValid(b.tls, b.Id)
	return
}
func (b Joint) IsValid() (r uint8) {
	r = b2Joint_IsValid(b.tls, b.Id)
	return
}
func (b World) EnableSleeping(flag uint8) {
	b2World_EnableSleeping(b.tls, b.Id, flag)
}
func (b World) IsSleepingEnabled() (r uint8) {
	r = b2World_IsSleepingEnabled(b.tls, b.Id)
	return
}
func (b World) EnableWarmStarting(flag uint8) {
	b2World_EnableWarmStarting(b.tls, b.Id, flag)
}
func (b World) IsWarmStartingEnabled() (r uint8) {
	r = b2World_IsWarmStartingEnabled(b.tls, b.Id)
	return
}
func (b World) GetAwakeBodyCount() (r int32) {
	r = b2World_GetAwakeBodyCount(b.tls, b.Id)
	return
}
func (b World) EnableContinuous(flag uint8) {
	b2World_EnableContinuous(b.tls, b.Id, flag)
}
func (b World) IsContinuousEnabled() (r uint8) {
	r = b2World_IsContinuousEnabled(b.tls, b.Id)
	return
}
func (b World) SetRestitutionThreshold(value float32) {
	b2World_SetRestitutionThreshold(b.tls, b.Id, value)
}
func (b World) GetRestitutionThreshold() (r float32) {
	r = b2World_GetRestitutionThreshold(b.tls, b.Id)
	return
}
func (b World) SetHitEventThreshold(value float32) {
	b2World_SetHitEventThreshold(b.tls, b.Id, value)
}
func (b World) GetHitEventThreshold() (r float32) {
	r = b2World_GetHitEventThreshold(b.tls, b.Id)
	return
}
func (b World) SetContactTuning(hertz float32, dampingRatio float32, pushSpeed float32) {
	b2World_SetContactTuning(b.tls, b.Id, hertz, dampingRatio, pushSpeed)
}
func (b World) SetMaximumLinearSpeed(maximumLinearSpeed float32) {
	b2World_SetMaximumLinearSpeed(b.tls, b.Id, maximumLinearSpeed)
}
func (b World) GetMaximumLinearSpeed() (r float32) {
	r = b2World_GetMaximumLinearSpeed(b.tls, b.Id)
	return
}
func (b World) GetProfile() (r Profile) {
	r = b2World_GetProfile(b.tls, b.Id)
	return
}
func (b World) GetCounters() (r Counters) {
	r = b2World_GetCounters(b.tls, b.Id)
	return
}
func (b World) SetUserData(userData uintptr) {
	b2World_SetUserData(b.tls, b.Id, userData)
}
func (b World) GetUserData() (r uintptr) {
	r = b2World_GetUserData(b.tls, b.Id)
	return
}
func (b World) DumpMemoryStats() {
	b2World_DumpMemoryStats(b.tls, b.Id)
}
func (b World) OverlapShape(proxy ShapeProxy, filter QueryFilter, fcn uintptr, context uintptr) (r1 TreeStats) {
	proxyStack := copyToStack(b.tls, proxy)
	defer proxyStack.Free()
	r1 = b2World_OverlapShape(b.tls, b.Id, proxyStack.Addr, filter, fcn, context)
	return
}
func (b World) CastRay(origin Vec2, translation Vec2, filter QueryFilter, fcn uintptr, context uintptr) (r TreeStats) {
	r = b2World_CastRay(b.tls, b.Id, origin, translation, filter, fcn, context)
	return
}
func (b World) CastRayClosest(origin Vec2, translation Vec2, filter QueryFilter) (r RayResult) {
	r = b2World_CastRayClosest(b.tls, b.Id, origin, translation, filter)
	return
}
func (b World) CastShape(proxy ShapeProxy, translation Vec2, filter QueryFilter, fcn uintptr, context uintptr) (r TreeStats) {
	proxyStack := copyToStack(b.tls, proxy)
	defer proxyStack.Free()
	r = b2World_CastShape(b.tls, b.Id, proxyStack.Addr, translation, filter, fcn, context)
	return
}
func (b World) CastMover(mover Capsule, translation Vec2, filter QueryFilter) (r float32) {
	moverStack := copyToStack(b.tls, mover)
	defer moverStack.Free()
	r = b2World_CastMover(b.tls, b.Id, moverStack.Addr, translation, filter)
	return
}
func (b World) SetGravity(gravity Vec2) {
	b2World_SetGravity(b.tls, b.Id, gravity)
}
func (b World) GetGravity() (r Vec2) {
	r = b2World_GetGravity(b.tls, b.Id)
	return
}
func (b World) Explode(explosionDef ExplosionDef) {
	explosionDefStack := copyToStack(b.tls, explosionDef)
	defer explosionDefStack.Free()
	b2World_Explode(b.tls, b.Id, explosionDefStack.Addr)
}
func (b World) RebuildStaticTree() {
	b2World_RebuildStaticTree(b.tls, b.Id)
}
func (b World) EnableSpeculative(flag uint8) {
	b2World_EnableSpeculative(b.tls, b.Id, flag)
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
const JointTypeDistance JointType = b2_distanceJoint
const JointTypeFilter JointType = b2_filterJoint
const JointTypeMotor JointType = b2_motorJoint
const JointTypeMouse JointType = b2_mouseJoint
const JointTypePrismatic JointType = b2_prismaticJoint
const JointTypeRevolute JointType = b2_revoluteJoint
const JointTypeWeld JointType = b2_weldJoint
const JointTypeWheel JointType = b2_wheelJoint
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
