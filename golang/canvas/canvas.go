package canvas

import (
	"fmt"
	"math"

	"github.com/croese/raytrace/tuple"
)

var blackPixel = tuple.Color(0, 0, 0)

const maxPPMLineLength = 70

type Canvas struct {
	width  int
	height int
	pixels []tuple.Tuple4
}

func New(width, height int) *Canvas {
	pixels := make([]tuple.Tuple4, width*height)
	for i := 0; i < len(pixels); i++ {
		pixels[i] = blackPixel
	}

	return &Canvas{
		width:  width,
		height: height,
		pixels: pixels,
	}
}

func (c *Canvas) Width() int {
	return c.width
}

func (c *Canvas) Height() int {
	return c.height
}

func (c *Canvas) PixelAt(x, y int) tuple.Tuple4 {
	if x < 0 || x >= c.width {
		panic(fmt.Sprintf("x out of range: %d", x))
	}
	if y < 0 || y >= c.height {
		panic(fmt.Sprintf("y out of range: %d", y))
	}

	index := y*c.width + x
	return c.pixels[index]
}

func (c *Canvas) WritePixel(x, y int, color tuple.Tuple4) {
	if x < 0 || x >= c.width || y < 0 || y >= c.height {
		return
	}

	index := y*c.width + x
	c.pixels[index] = color
}

func (c *Canvas) ToPPMString() string {
	header := fmt.Sprintf("P3\n%d %d\n255\n", c.width, c.height)
	pixels := ""

	line := ""
	startOfLine := true
	prefix := ""
	for i, p := range c.pixels {
		startOfLine = i%c.width == 0
		if i > 0 && startOfLine {
			pixels += line + "\n"
			line = ""
		}

		if startOfLine {
			prefix = ""
		} else {
			prefix = " "
		}
		r := clampComponent(p.Red())
		g := clampComponent(p.Green())
		b := clampComponent(p.Blue())
		rStr := fmt.Sprint(r)
		gStr := fmt.Sprint(g)
		bStr := fmt.Sprint(b)

		if len(line)+len(rStr)+1 > maxPPMLineLength {
			pixels += line + "\n"
			line = rStr
		} else {
			line += prefix + rStr
		}

		if len(line)+len(gStr)+1 > maxPPMLineLength {
			pixels += line + "\n"
			line = gStr
		} else {
			line += " " + gStr
		}

		if len(line)+len(bStr)+1 > maxPPMLineLength {
			pixels += line + "\n"
			line = bStr
		} else {
			line += " " + bStr
		}
	}

	pixels += line + "\n"

	return header + pixels
}

func clampComponent(value float64) int {
	v := int(math.Round(value * 255))
	if v < 0 {
		return 0
	} else if v > 255 {
		return 255
	}
	return v
}
