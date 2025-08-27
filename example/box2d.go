package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/oliverbestmann/box2d-go"
)

func b2New(gravity float64) Physics {
	b := box2d.NewBox2D()

	def := b.DefaultWorldDef()
	def.Gravity = box2d.Vec2{Y: float32(gravity)}

	return &Box2D{
		Box:   b,
		World: b.CreateWorld(def),
	}
}

type Box2D struct {
	Box   *box2d.Box2D
	World box2d.World
}

func (ph Box2D) Step(dt float64, subSteps int) {
	ph.World.Step(float32(dt), int32(subSteps))
}

func (ph Box2D) CreateSquare(halfSize float64, centerX, centerY float64) Body {
	p := pathSquare(halfSize)

	var tr box2d.Transform
	tr.P.X = float32(centerX)
	tr.P.Y = float32(centerY)
	tr.Q.C = 1

	def := ph.Box.DefaultBodyDef()
	def.Type1 = box2d.DynamicBody
	body := ph.World.CreateBody(def)
	body.SetTransform(tr.P, tr.Q)

	shape := ph.Box.DefaultShapeDef()
	shape.Density = 1
	shape.Material.Restitution = 0
	shape.Material.Friction = 1
	body.CreatePolygonShape(shape, ph.Box.MakeSquare(float32(halfSize)))

	return &b2Body{shape: p, body: body}
}

func (ph Box2D) CreateStaticLine(x0, y0, x1, y1 float64) Body {
	def := ph.Box.DefaultBodyDef()
	def.Type1 = box2d.StaticBody

	shape := ph.Box.DefaultShapeDef()
	shape.Density = 1
	shape.Material.Restitution = 0
	shape.Material.Friction = 1

	body := ph.World.CreateBody(def)
	body.CreateSegmentShape(shape, box2d.Segment{
		Point1: box2d.Vec2{X: float32(x0), Y: float32(y0)},
		Point2: box2d.Vec2{X: float32(x1), Y: float32(y1)},
	})

	return &b2Body{
		shape: pathLine(x0, y0, x1, y1),
		body:  body,
	}
}

type b2Body struct {
	shape vector.Path
	body  box2d.Body
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
