package b2

import (
	"slices"
	"unsafe"
)

type DrawPolygon func(vertices []Vec2, color HexColor)
type DrawSolidPolygon func(transform Transform, vertices []Vec2, radius float32, color HexColor)
type DrawCircle func(center Vec2, radius float32, color HexColor)
type DrawSolidCircle func(transform Transform, radius float32, color HexColor)
type DrawSolidCapsule func(p1 Vec2, p2 Vec2, radius float32, color HexColor)
type DrawSegment func(p1 Vec2, p2 Vec2, color HexColor)
type DrawTransform func(transform Transform)
type DrawPoint func(p Vec2, size float32, color HexColor)
type DrawString func(p Vec2, s string, color HexColor)

type DebugDraw struct {
	DrawPolygon      DrawPolygon
	DrawSolidPolygon DrawSolidPolygon
	DrawCircle       DrawCircle
	DrawSolidCircle  DrawSolidCircle
	DrawSolidCapsule DrawSolidCapsule
	DrawSegment      DrawSegment
	DrawTransform    DrawTransform
	DrawPoint        DrawPoint
	DrawString       DrawString

	DrawingBounds    AABB
	UseDrawingBounds bool

	DrawShapes           bool
	DrawJoints           bool
	DrawJointExtras      bool
	DrawBounds           bool
	DrawMass             bool
	DrawBodyNames        bool
	DrawContacts         bool
	DrawGraphColors      bool
	DrawContactNormals   bool
	DrawContactImpulses  bool
	DrawContactFeatures  bool
	DrawFrictionImpulses bool
	DrawIslands          bool
}

func (b World) Draw(draw DebugDraw) {
	context := theStack.RegisterObject(&draw)
	defer theStack.ForgetObject(context)

	var dd b2DebugDraw
	dd.DrawSegmentFcn = __ccgo_fp(cbDrawSegment)
	dd.DrawShapes = boolUint8(draw.DrawShapes)
	dd.Context = context

	ddStack := copyToStack(theStack, dd)
	defer ddStack.Free()

	b2World_Draw(theStack, b.Id, ddStack.Addr)
}

func cbDrawPolygon(tls *_Stack, vertices uintptr, vertexCount int32, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawPolygon != nil {
		vertices := toSlice[Vec2](vertices, vertexCount)
		ctx.DrawPolygon(vertices, color)
	}
}
func cbDrawSolidPolygon(tls *_Stack, transform Transform, vertices uintptr, vertexCount int32, radius float32, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawSolidPolygon != nil {
		vertices := toSlice[Vec2](vertices, vertexCount)
		ctx.DrawSolidPolygon(transform, vertices, radius, color)
	}
}
func cbDrawCircle(tls *_Stack, center Vec2, radius float32, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawCircle != nil {
		ctx.DrawCircle(center, radius, color)
	}
}
func cbDrawSolidCircle(tls *_Stack, transform Transform, radius float32, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawSolidCircle != nil {
		ctx.DrawSolidCircle(transform, radius, color)
	}
}
func cbDrawSolidCapsule(tls *_Stack, p1 Vec2, p2 Vec2, radius float32, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawSolidCapsule != nil {
		ctx.DrawSolidCapsule(p1, p2, radius, color)
	}
}
func cbDrawSegment(tls *_Stack, p1 Vec2, p2 Vec2, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawSegment != nil {
		ctx.DrawSegment(p1, p2, color)
	}
}
func cbDrawTransform(tls *_Stack, transform Transform, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawTransform != nil {
		ctx.DrawTransform(transform)
	}
}
func cbDrawPoint(tls *_Stack, p Vec2, size float32, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawPoint != nil {
		ctx.DrawPoint(p, size, color)
	}
}
func cbDrawString(tls *_Stack, p Vec2, s uintptr, color HexColor, context uintptr) {
	ctx := tls.GetObject(context).(*DebugDraw)
	if ctx.DrawString != nil {
		ctx.DrawString(p, toString(s), color)
	}
}

func toSlice[T any](ptr uintptr, count int32) []T {
	return slices.Clone(unsafe.Slice((*T)(unsafe.Pointer(ptr)), count))
}
