package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/jakecoffman/cp/v2"
)

func cpNew(gravity float64) Physics {
	space := cp.NewSpace()
	space.SetGravity(cp.Vector{Y: gravity})
	return Chipmunk{Space: space}
}

type Chipmunk struct {
	Space *cp.Space
}

func (ph Chipmunk) Step(dt float64, subSteps int) {
	ph.Space.Iterations = uint(subSteps)
	ph.Space.Step(dt)
}

func (ph Chipmunk) CreateSquare(halfSize float64, centerX, centerY float64) Body {
	body := cp.NewBody(1, 1)

	shape := cp.NewBox(body, halfSize*2, halfSize*2, 0)
	shape.SetElasticity(0)
	shape.SetDensity(1)
	shape.SetFriction(1)
	body.AddShape(shape)
	body.SetPosition(cp.Vector{X: centerX, Y: centerY})

	ph.Space.AddShape(shape)
	ph.Space.AddBody(body)

	return &cpBody{
		shape: pathSquare(halfSize),
		body:  body,
	}
}

func (ph Chipmunk) CreateStaticLine(x0, y0, x1, y1 float64) Body {
	body := cp.NewStaticBody()

	a := cp.Vector{X: x0, Y: y0}
	b := cp.Vector{X: x1, Y: y1}
	shape := cp.NewSegment(body, a, b, 0)
	shape.SetElasticity(0)
	shape.SetDensity(1)
	shape.SetFriction(1)
	body.AddShape(shape)

	ph.Space.AddShape(shape)
	ph.Space.AddBody(body)

	return &cpBody{
		shape: pathLine(x0, y0, x1, y1),
		body:  body,
	}
}

type cpBody struct {
	shape vector.Path
	body  *cp.Body
}

func (b *cpBody) Shape() vector.Path {
	return b.shape
}

func (b *cpBody) Position() (float64, float64) {
	pos := b.body.Position()
	return pos.X, pos.Y
}

func (b *cpBody) Rotation() float64 {
	rot := b.body.Rotation()
	return math.Atan2(rot.Y, rot.X)
}
