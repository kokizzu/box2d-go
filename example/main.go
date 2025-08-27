package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/oliverbestmann/box2d-go"
)

func main() {
	b := box2d.NewBox2D()

	def := b.DefaultWorldDef()
	def.Gravity = box2d.Vec2{Y: -10}
	w := b.CreateWorld(def)

	// floor
	floor := b.DefaultBodyDef()
	floor.Type1 = box2d.StaticBody
	{
		floor := w.CreateBody(floor)
		floor.CreateSegmentShape(b.DefaultShapeDef(), box2d.Segment{
			Point1: box2d.Vec2{X: -10},
			Point2: box2d.Vec2{X: 10},
		})
	}

	square := b.MakeSquare(0.5)

	var bodies []box2d.Body

	const Layers = 20

	// create layers of boxes
	for l := range Layers {
		for i := range l {
			centerX := float32(i) + 0.5 - float32(l)/2
			centerY := Layers - float32(l)

			var tr box2d.Transform
			tr.P.X = centerX
			tr.P.Y = centerY
			tr.Q.C = 1

			box := b.DefaultBodyDef()
			box.Type1 = box2d.DynamicBody
			body := w.CreateBody(box)
			body.SetTransform(tr.P, tr.Q)
			body.CreatePolygonShape(b.DefaultShapeDef(), square)

			bodies = append(bodies, body)
		}
	}

	ebiten.SetWindowSize(1200, 800)
	_ = ebiten.RunGame(&viz{w: w, bodies: bodies})
}

type viz struct {
	w           box2d.World
	bodies      []box2d.Body
	physicsTime float64
}

func (v *viz) Update() error {
	startTime := time.Now()
	v.w.Step(1/60.0, 4)

	if v.physicsTime == 0 {
		v.physicsTime = time.Since(startTime).Seconds()
	} else {
		v.physicsTime = 0.95*v.physicsTime + 0.05*time.Since(startTime).Seconds()
	}

	return nil
}

func (v *viz) Draw(screen *ebiten.Image) {
	w, h := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())

	var toScreen ebiten.GeoM

	toScreen.Scale(30.0, -30.0)
	toScreen.Translate(w/2, h-50)

	var floor vector.Path
	floor.MoveTo(-10, 0)
	floor.LineTo(10, 0)

	floor.MoveTo(0, 0)
	floor.LineTo(0, 0.1)

	var p vector.Path

	p.Reset()
	p.AddPath(&floor, &vector.AddPathOptions{GeoM: toScreen})
	vector.StrokePath(screen, &p, &vector.StrokeOptions{Width: 1}, &vector.DrawPathOptions{})

	var box vector.Path
	box.MoveTo(-0.5, -0.5)
	box.LineTo(-0.5, 0.5)
	box.LineTo(0.5, 0.5)
	box.LineTo(0.5, -0.5)
	box.Close()

	for _, body := range v.bodies {
		pos := body.GetPosition()

		var toWorld ebiten.GeoM
		toWorld.Translate(float64(pos.X), float64(pos.Y))
		toWorld.Concat(toScreen)

		p.Reset()
		p.AddPath(&box, &vector.AddPathOptions{GeoM: toWorld})
		vector.StrokePath(screen, &p, &vector.StrokeOptions{Width: 1}, &vector.DrawPathOptions{})
	}

	text := fmt.Sprintf("physics: %1.2fms", v.physicsTime*1000.0)
	ebitenutil.DebugPrintAt(screen, text, 16, 16)
}

func (v *viz) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
