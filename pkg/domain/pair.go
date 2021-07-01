package domain

type Pair struct {
	X int
	Y int
}

func (a Pair) Add(b Pair) Pair {
	return Pair{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (a Pair) Negate() Pair {
	return Pair{
		X: -a.X,
		Y: -a.Y,
	}
}

func (a Pair) Sub(b Pair) Pair {
	return a.Add(b.Negate())
}
