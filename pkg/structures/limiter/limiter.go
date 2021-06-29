package limiter

import "time"

// TODO: give meaningful names

type Limiter struct {
	A float64
	B float64

	Mid  float64
	Last time.Time
}

func NewLimiter(a, b float64) *Limiter {
	return &Limiter{
		A:    a,
		B:    b,
		Mid:  a,
		Last: time.Now(),
	}
}

func (l *Limiter) Calc() float64 {
	now := time.Now()
	cur := now.Sub(l.Last)
	l.Mid += l.B * (float64(cur.Microseconds()) - l.Mid)
	l.Last = now
	return l.Mid
}
