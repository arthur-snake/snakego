package maptool

import "github.com/arthur-snake/snakego/pkg/domain"

func CreateMap(size domain.FieldSize, zero domain.ObjectID) [][]domain.ObjectID {
	field := make([][]domain.ObjectID, size.SizeX)
	for x := 0; x < size.SizeX; x++ {
		field[x] = make([]domain.ObjectID, size.SizeY)
		for y := 0; y < size.SizeY; y++ {
			field[x][y] = zero
		}
	}

	return field
}
