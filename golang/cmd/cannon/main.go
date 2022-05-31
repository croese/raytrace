package main

import (
	"fmt"

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
		v:   tuple.Vector(1, 1, 0).Norm(),
	}

	e := environment{
		gravity: tuple.Vector(0, -0.1, 0),
		wind:    tuple.Vector(-0.01, 0, 0),
	}

	ticks := 0
	for p.pos.Y() > 0 {
		p = tick(e, p)
		ticks += 1
		fmt.Printf("tick %d: %+v\n", ticks, p.pos)
	}
	fmt.Printf("%d ticks to reach the ground\n", ticks)
}
