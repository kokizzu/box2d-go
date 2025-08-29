package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const Layers = 20

func main() {
	// defer profile.Start(profile.CPUProfile).Stop()

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(1200, 800)
	ops := ebiten.RunGameOptions{}
	ops.SingleThread = true
	_ = ebiten.RunGameWithOptions(&viz{}, &ops)
}

type viz struct {
	physics     Physics
	bodies      []Body
	physicsTime float64
}

func (v *viz) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyB) {
		v.initialize(b2New)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		v.initialize(cpNew)
	}

	if v.physics != nil {
		startTime := time.Now()
		v.physics.Step(1/60.0, 4)

		if v.physicsTime == 0 {
			v.physicsTime = time.Since(startTime).Seconds()
		} else {
			v.physicsTime = 0.95*v.physicsTime + 0.05*time.Since(startTime).Seconds()
		}
	}

	return nil
}

func (v *viz) Draw(screen *ebiten.Image) {
	if v.physics == nil {
		text := "press b for box2d, c for jakecoffman/cp"
		ebitenutil.DebugPrintAt(screen, text, 16, 16)
		return
	}

	w, h := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())

	var toScreen ebiten.GeoM

	var minX, minY, maxX, maxY float64
	for _, body := range v.bodies {
		x, y := body.Position()
		minX = min(minX, x)
		maxX = max(maxX, x)
		minY = min(minY, y)
		maxY = max(minY, y)
	}

	_ = maxY

	// scale := (h - 50) / (Layers + 5)
	scale := w / (5 * Layers)

	toScreen.Scale(scale, -scale)
	toScreen.Translate(w/2, h-50)

	var p vector.Path

	for _, body := range v.bodies {
		var toWorld ebiten.GeoM
		toWorld.Rotate(body.Rotation())
		toWorld.Translate(body.Position())
		toWorld.Concat(toScreen)

		shape := body.Shape()

		p.Reset()
		p.AddPath(&shape, &vector.AddPathOptions{GeoM: toWorld})
		vector.StrokePath(screen, &p, &vector.StrokeOptions{Width: 1}, &vector.DrawPathOptions{})
	}

	text := fmt.Sprintf("physics %T: %1.2fms", v.physics, v.physicsTime*1000.0)
	ebitenutil.DebugPrintAt(screen, text, 16, 16)
}

func (v *viz) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (v *viz) initialize(makePhysics func(gravity float64) Physics) {
	ph := makePhysics(-10)

	var bodies []Body

	bodies = append(bodies, ph.CreateStaticLine(-2*Layers, 0, 2*Layers, 0))
	bodies = append(bodies, ph.CreateStaticLine(-2*Layers, 0, -2*Layers, 100))
	bodies = append(bodies, ph.CreateStaticLine(2*Layers, 0, 2*Layers, 100))

	// create layers of boxes
	for l := range Layers {
		for i := range l {
			centerX := float64(i) + 0.5 - float64(l)/2
			centerY := (Layers - float64(l) - 0.5) * 1.0

			bodies = append(bodies, ph.CreateSquare(0.5, centerX, centerY, 1))
		}
	}

	// create a fast moving box
	b := ph.CreateSquare(2, -0.5*Layers, Layers, 10)
	b.SetVelocity(50, -100)
	bodies = append(bodies, b)

	v.physics = ph
	v.bodies = bodies
}

func pathSquare(halfSize float64) vector.Path {
	var p vector.Path
	p.MoveTo(float32(-halfSize), -float32(halfSize))
	p.LineTo(-float32(halfSize), float32(halfSize))
	p.LineTo(float32(halfSize), float32(halfSize))
	p.LineTo(float32(halfSize), -float32(halfSize))
	p.Close()
	return p
}

func pathLine(x0 float64, y0 float64, x1 float64, y1 float64) vector.Path {
	var p vector.Path
	p.MoveTo(float32(x0), float32(y0))
	p.LineTo(float32(x1), float32(y1))
	return p
}

type Body interface {
	Shape() vector.Path
	Position() (float64, float64)
	Rotation() float64
	SetVelocity(x, y float64)
}

type Physics interface {
	CreateSquare(halfSize, centerX, centerY, density float64) Body
	CreateStaticLine(x0, y0, x1, y1 float64) Body
	Step(dt float64, subSteps int)
}
