package maptool

import "github.com/arthur-snake/snakego/pkg/domain"

func ShiftDir(size domain.FieldSize, f [][]domain.Cell, dir domain.Pair, fill domain.Cell) {
	IterateCustom(size, dir, func(cur domain.Pair) {
		nxt := cur.Sub(dir)
		if !size.IsInside(nxt) {
			f[cur.X][cur.Y] = fill
		} else {
			f[cur.X][cur.Y] = f[nxt.X][nxt.Y]
		}
	})
}
