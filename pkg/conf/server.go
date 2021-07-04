package conf

import "time"

type Server struct {
	SizeX     int           `env:"SIZE_X" envDefault:"75"`
	SizeY     int           `env:"SIZE_Y" envDefault:"40"`
	TickTime  time.Duration `env:"TICK_TIME" envDefault:"80ms"`
	FoodCount int           `env:"FOOD_COUNT" envDefault:"3"`
}
