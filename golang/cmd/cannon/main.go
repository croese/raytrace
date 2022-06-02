package main

import (
	"fmt"
	"log"
	"os"

	"github.com/croese/raytrace/canvas"
	"github.com/croese/raytrace/tuple"
)

type projectile struct {
	pos tuple.Tuple4
	v   tuple.Tuple4
}

type environment struct {
	gravity tuple.Tuple4
	wind    tuple.Tuple4
}

func tick(env environment, proj projectile) projectile {
	pos := proj.pos.Plus(proj.v)
	v := proj.v.Plus(env.gravity).Plus(env.wind)
	return projectile{
		pos: pos,
		v:   v,
	}
}

func main() {
	p := projectile{
		pos: tuple.Point(0, 1, 0),
		v:   tuple.Vector(1, 1.8, 0).Norm().ScalarMult(11.25),
	}

	e := environment{
		gravity: tuple.Vector(0, -0.1, 0),
		wind:    tuple.Vector(-0.01, 0, 0),
	}

	c := canvas.New(900, 550)
	plotColor := tuple.Color(1, 0, 0)

	ticks := 0
	for p.pos.Y() > 0 {
		p = tick(e, p)
		ticks += 1
		fmt.Printf("tick %d: %+v\n", ticks, p.pos)

		convertedX := int(p.pos.X())
		convertedY := c.Height() - int(p.pos.Y())

		c.WritePixel(convertedX, convertedY, plotColor)
	}
	fmt.Printf("%d ticks to reach the ground\n", ticks)

	file, err := os.Create("cannon.ppm") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(c.ToPPMString())
}
