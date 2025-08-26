package box2d

import (
	"fmt"
	"testing"
)

func TestBounce(t *testing.T) {
	b := NewBox2D()

	fmt.Println(b.GetVersion())

	def := b.DefaultWorldDef()
	var w World = b.CreateWorld(def)

	w.SetGravity(Vec2{Y: -10})

	ballDef := b.DefaultBodyDef()
	ballDef.Type1 = DynamicBody
	ballDef.Position.Y = 5

	var ball Body = w.CreateBody(ballDef)

	ballShape := b.DefaultShapeDef()
	ballShape.Material.Restitution = 0.2
	ballShape.EnableContactEvents = 1
	ballShape.EnableHitEvents = 1
	ball.CreateCircleShape(ballShape, Circle{Radius: 1})

	groundDef := b.DefaultBodyDef()
	groundDef.Type1 = StaticBody

	var ground Body = w.CreateBody(groundDef)

	groundSegment := Segment{
		Point1: Vec2{-20.0, 0.0},
		Point2: Vec2{20.0, 0.0},
	}

	groundShape := b.DefaultShapeDef()
	ground.CreateSegmentShape(groundShape, groundSegment)

	for idx := range 100 {
		w.Step(1.0/60.0, 1)
		fmt.Println(idx, "pos", ball.GetPosition(), "vec", ball.GetLinearVelocity().Y)

		// read movement events
		for _, ev := range w.GetBodyEvents().MoveEvents {
			fmt.Printf("Movement: %#v\n", ev)
		}

		// and check for collisions
		contacts := w.GetContactEvents()
		for _, ev := range contacts.BeginEvents {
			fmt.Printf("ContactBeginTouchEvent: %#v\n", ev)
		}

		for _, ev := range contacts.EndEvents {
			fmt.Printf("ContactEndTouchEvent: %#v\n", ev)
		}

		for _, ev := range contacts.HitEvents {
			fmt.Printf("HitEvent: %#v\n", ev)
		}
	}
}
