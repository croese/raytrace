package canvas_test

import (
	"testing"

	"github.com/croese/raytrace/canvas"
	"github.com/croese/raytrace/tuple"
)

func checkTupleEquality(t *testing.T, a, b tuple.Tuple4) {
	t.Helper()
	if !a.Equals(b) {
		t.Errorf("expected %+v and %+v to be equal", a, b)
	}
}

func TestCreateCanvas(t *testing.T) {
	c := canvas.New(10, 20)

	t.Run("Dimensions", func(t *testing.T) {
		if c.Width() != 10 {
			t.Errorf("expected width to be 10, got %d", c.Width())
		}

		if c.Height() != 20 {
			t.Errorf("expected height to be 20, got %d", c.Height())
		}
	})

	t.Run("DefaultCanvasColor", func(t *testing.T) {
		black := tuple.Color(0, 0, 0)
		for y := 0; y < c.Height(); y++ {
			for x := 0; x < c.Width(); x++ {
				checkTupleEquality(t, black, c.PixelAt(x, y))
			}
		}
	})

}

func TestWritePixel(t *testing.T) {
	c := canvas.New(10, 20)
	red := tuple.Color(1, 0, 0)

	c.WritePixel(2, 3, red)

	checkTupleEquality(t, red, c.PixelAt(2, 3))
}

func TestPPM(t *testing.T) {
	t.Run("PixelData", func(t *testing.T) {
		c := canvas.New(5, 3)
		c1 := tuple.Color(1.5, 0, 0)
		c2 := tuple.Color(0, 0.5, 0)
		c3 := tuple.Color(-0.5, 0, 1)

		c.WritePixel(0, 0, c1)
		c.WritePixel(2, 1, c2)
		c.WritePixel(4, 2, c3)

		ppm := c.ToPPMString()

		expected := `P3
5 3
255
255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`

		if ppm != expected {
			t.Errorf("invalid PPM data. expected %q, got %q", expected, ppm)
		}
	})

	t.Run("SplitLines", func(t *testing.T) {
		c := canvas.New(10, 2)
		color := tuple.Color(1, 0.8, 0.6)
		for x := 0; x < c.Width(); x++ {
			for y := 0; y < c.Height(); y++ {
				c.WritePixel(x, y, color)
			}
		}

		expected := `P3
10 2
255
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
`

		ppm := c.ToPPMString()

		if ppm != expected {
			t.Errorf("invalid PPM data. expected %q, got %q", expected, ppm)
		}
	})
}
