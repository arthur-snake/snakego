package domain

type CellType string

var (
	FreeCell   CellType = "free"
	FoodCell   CellType = "food"
	PlayerCell CellType = "player"
	BlockCell  CellType = "block"
)

type Cell struct {
	ID ObjectID
}
