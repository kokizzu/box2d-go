package box2d

import (
	"fmt"
	"testing"
)

func TestBounce(t *testing.T) {
	b := NewBox2D()

	fmt.Println(b.GetVersion())

	def := b.DefaultWorldDef()
	w := b.CreateWorld(def)

	b.WorldSetGravity(w, Vec2{Y: -10})

	ballDef := b.DefaultBodyDef()
	ballDef.Type1 = DynamicBody
	ballDef.Position.Y = 5

	ball := b.CreateBody(w, ballDef)

	ballShape := b.DefaultShapeDef()
	ballShape.Material.Restitution = 0.2
	ballShape.EnableContactEvents = 1
	ballShape.EnableHitEvents = 1
	b.CreateCircleShape(ball, ballShape, Circle{Radius: 1})

	groundDef := b.DefaultBodyDef()
	groundDef.Type1 = StaticBody

	ground := b.CreateBody(w, groundDef)

	groundSegment := b2Segment{
		Point1: b2Vec2{-20.0, 0.0},
		Point2: b2Vec2{20.0, 0.0},
	}

	groundShape := b.DefaultShapeDef()
	b.CreateSegmentShape(ground, groundShape, groundSegment)

	for idx := range 100 {
		b.WorldStep(w, 1.0/60.0, 1)
		fmt.Println(idx, "pos", b.BodyGetPosition(ball), "vec", b.BodyGetLinearVelocity(ball).Y)

		// read movement events
		for _, ev := range b.WorldGetBodyEvents(w).MoveEvents {
			fmt.Printf("Movement: %#v\n", ev)
		}

		// and check for collisions
		contacts := b.WorldGetContactEvents(w)
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
