package servtool

import (
	"github.com/arthur-snake/snakego/pkg/domain"
)

type Controller struct {
	Direction domain.BaseDirection
}

func NewController() *Controller {
	return &Controller{
		Direction: domain.Right,
	}
}

func (c *Controller) Turn(direction domain.BaseDirection) {
	c.Direction = direction
}

func (c *Controller) Move() domain.BaseDirection {
	return c.Direction
}
