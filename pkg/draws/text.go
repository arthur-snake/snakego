package draws

import (
	"image"
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

func Text(h int, text string) []Slab {
	img := image.NewRGBA(image.Rect(0, 0, len(text)*8, h-1))
	addLabel(img, 0, h-1, text)

	var slabs []Slab
	for x := 0; x < img.Bounds().Dx(); x++ {
		slab := Slab{}
		for y := 0; y < img.Bounds().Dy(); y++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a > 0 {
				slab.Filled = append(slab.Filled, y)
			}
		}

		slabs = append(slabs, slab)
	}

	return slabs
}
