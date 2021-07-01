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

type Cell struct {
	ID   ObjectID
	Food int
}

type CellWithLocation struct {
	Location Pair
	Cell     *Cell
}
