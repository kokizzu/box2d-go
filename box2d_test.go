package b2

import (
	"fmt"
	"testing"
)

func TestBounce(t *testing.T) {
	fmt.Println(GetVersion())

	def := DefaultWorldDef()
	var w World = CreateWorld(def)

	w.SetGravity(Vec2{Y: -10})

	ballDef := DefaultBodyDef()
	ballDef.Type1 = DynamicBody
	ballDef.Position.Y = 5

	var ball Body = w.CreateBody(ballDef)

	ballShape := DefaultShapeDef()
	ballShape.Material.Restitution = 0.2
	ballShape.EnableContactEvents = 1
	ballShape.EnableHitEvents = 1
	ball.CreateCircleShape(ballShape, Circle{Radius: 1})

	groundDef := DefaultBodyDef()
	groundDef.Type1 = StaticBody

	var ground Body = w.CreateBody(groundDef)

	groundSegment := Segment{
		Point1: Vec2{-20.0, 0.0},
		Point2: Vec2{20.0, 0.0},
	}

	groundShape := DefaultShapeDef()
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

func TestWorld_Draw(t *testing.T) {
	w := CreateWorld(DefaultWorldDef())

	b := w.CreateBody(DefaultBodyDef())
	b.CreateSegmentShape(DefaultShapeDef(), Segment{
		Point1: Vec2{1, 1},
		Point2: Vec2{2, 5},
	})

	w.Draw(DebugDraw{
		DrawSegment: func(p1 Vec2, p2 Vec2, color HexColor) {
			fmt.Println("draw segment", p1, p2)
		},

		DrawShapes: true,
	})
}
