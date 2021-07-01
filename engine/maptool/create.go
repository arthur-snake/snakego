package maptool

import "github.com/arthur-snake/snakego/pkg/domain"

func CreateMap(size domain.FieldSize, zero domain.Cell) [][]domain.Cell {
	field := make([][]domain.Cell, size.SizeX)
	for x := 0; x < size.SizeX; x++ {
		field[x] = make([]domain.Cell, size.SizeY)
		for y := 0; y < size.SizeY; y++ {
			field[x][y] = zero
		}
	}

	return field
}
