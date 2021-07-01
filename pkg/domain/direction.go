package domain

type BaseDirection struct {
	Dir  Pair
	Name string
}

var baseDirections = []BaseDirection{
	{Pair{0, -1}, "UP"},
	{Pair{-1, 0}, "LEFT"},
	{Pair{0, 1}, "DOWN"},
	{Pair{1, 0}, "RIGHT"},
}

var (
	Up    = baseDirections[0]
	Left  = baseDirections[1]
	Down  = baseDirections[2]
	Right = baseDirections[3]
)

func ParseBaseDirection(str string) (BaseDirection, bool) {
	for _, dir := range baseDirections {
		if dir.Name == str {
			return dir, true
		}
	}

	return BaseDirection{}, false
}
