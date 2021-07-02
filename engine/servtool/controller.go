package servtool

import (
	"github.com/arthur-snake/snakego/pkg/domain"
)

type Controller struct {
	queue    []domain.BaseDirection
	dir      domain.BaseDirection
	lastMove domain.BaseDirection
}

func NewController() *Controller {
	return &Controller{
		dir: domain.Right,
	}
}

func (c *Controller) Turn(direction domain.BaseDirection) {
	tail := c.dir
	if len(c.queue) > 0 {
		tail = c.queue[len(c.queue)-1]
	}

	if tail == direction {
		// no seq repeats
		return
	}

	c.queue = append(c.queue, direction)
}

func (c *Controller) PreMove() domain.BaseDirection {
	if len(c.queue) == 0 {
		return c.dir
	}

	nxt := c.queue[0]
	c.queue = c.queue[1:]

	if nxt.Dir == c.lastMove.Dir.Negate() {
		return c.PreMove()
	}

	c.dir = nxt
	return nxt
}

func (c *Controller) PostMove() {
	c.lastMove = c.dir
}
