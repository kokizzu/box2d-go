package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/oliverbestmann/box2d-go"
)

func b2New(gravity float64) Physics {
	def := b2.DefaultWorldDef()
	def.Gravity = b2.Vec2{Y: float32(gravity)}

	b2.EnableConcurrency(&def)

	return &Box2D{
		World: b2.CreateWorld(def),
	}
}

type Box2D struct {
	World b2.World
}

func (ph Box2D) Step(dt float64, subSteps int) {
	ph.World.Step(float32(dt), int32(subSteps))
}

func (ph Box2D) CreateSquare(halfSize, centerX, centerY, density float64) Body {
	p := pathSquare(halfSize)

	var tr b2.Transform
	tr.P.X = float32(centerX)
	tr.P.Y = float32(centerY)
	tr.Q.C = 1

	def := b2.DefaultBodyDef()
	def.Type1 = b2.DynamicBody
	body := ph.World.CreateBody(def)
	body.SetTransform(tr.P, tr.Q)

	shape := b2.DefaultShapeDef()
	shape.Density = float32(density)
	shape.Material.Restitution = 0
	shape.Material.Friction = 1
	body.CreatePolygonShape(shape, b2.MakeSquare(float32(halfSize)))

	return &b2Body{shape: p, body: body}
}

func (ph Box2D) CreateStaticLine(x0, y0, x1, y1 float64) Body {
	def := b2.DefaultBodyDef()
	def.Type1 = b2.StaticBody

	shape := b2.DefaultShapeDef()
	shape.Density = 1
	shape.Material.Restitution = 0
	shape.Material.Friction = 1

	body := ph.World.CreateBody(def)
	body.CreateSegmentShape(shape, b2.Segment{
		Point1: b2.Vec2{X: float32(x0), Y: float32(y0)},
		Point2: b2.Vec2{X: float32(x1), Y: float32(y1)},
	})

	return &b2Body{
		shape: pathLine(x0, y0, x1, y1),
		body:  body,
	}
}

type b2Body struct {
	shape vector.Path
	body  b2.Body
}

func (b *b2Body) Shape() vector.Path {
	return b.shape
}

func (b *b2Body) Position() (float64, float64) {
	pos := b.body.GetPosition()
	return float64(pos.X), float64(pos.Y)
}

func (b *b2Body) Rotation() float64 {
	rot := b.body.GetRotation()
	return math.Atan2(float64(rot.S), float64(rot.C))
}

func (b *b2Body) SetVelocity(x, y float64) {
	b.body.SetLinearVelocity(b2.Vec2{X: float32(x), Y: float32(y)})
}
