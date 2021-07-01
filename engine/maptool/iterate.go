package maptool

import "github.com/arthur-snake/snakego/pkg/domain"

func IterateX(start, end int, shiftDir int, callback func(int)) {
	if shiftDir <= 0 {
		for x := start; x < end; x++ {
			callback(x)
		}
	} else {
		for x := 0; x < end-start; x++ {
			callback(end - 1 - x)
		}
	}
}

func IterateCustom(size domain.FieldSize, dir domain.Pair, callback func(pair domain.Pair)) {
	IterateX(0, size.SizeX, dir.X, func(x int) {
		IterateX(0, size.SizeY, dir.Y, func(y int) {
			callback(domain.Pair{X: x, Y: y})
		})
	})
}
