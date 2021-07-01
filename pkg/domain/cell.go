package domain

// ObjectID represents id of an object inside a cell.
type ObjectID string

type CellType string

var (
	FreeCell   CellType = "free"
	FoodCell   CellType = "food"
	PlayerCell CellType = "player"
	BlockCell  CellType = "block"
)
