package maptool

import (
	"math/rand"

	"github.com/arthur-snake/snakego/pkg/domain"
)

type Distribution map[domain.ObjectID][]domain.CellWithLocation

func DistributeMap(size domain.FieldSize, f [][]domain.Cell) Distribution {
	distr := make(Distribution)
	IterateCustom(size, domain.Pair{}, func(loc domain.Pair) {
		ptr := &f[loc.X][loc.Y]
		distr[ptr.ID] = append(distr[ptr.ID], domain.CellWithLocation{
			Location: loc,
			Cell:     ptr,
		})
		shuffleTail(distr[ptr.ID])
	})

	return distr
}

func shuffleTail(arr []domain.CellWithLocation) {
	n := len(arr)
	i := rand.Intn(n) //nolint:gosec

	arr[i], arr[n-1] = arr[n-1], arr[i]
}
