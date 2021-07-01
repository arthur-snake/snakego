package maptool

import "github.com/arthur-snake/snakego/pkg/domain"

type Distribution map[domain.ObjectID][]domain.CellWithLocation

func DistributeMap(size domain.FieldSize, f [][]domain.Cell) Distribution {
	distr := make(Distribution)
	IterateCustom(size, domain.Pair{}, func(loc domain.Pair) {
		ptr := &f[loc.X][loc.Y]
		distr[ptr.ID] = append(distr[ptr.ID], domain.CellWithLocation{
			Location: loc,
			Cell:     ptr,
		})
	})

	return distr
}
