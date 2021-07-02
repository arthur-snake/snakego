package conf

import (
	"github.com/caarlos0/env/v6"
)

type App struct {
	PrometheusBind string `env:"PROMETHEUS_BIND" envDefault:":2112"`
	ServerBind     string `env:"SERVER_BIND" envDefault:":8080"`

	SizeX     int `env:"SIZE_X" envDefault:"75"`
	SizeY     int `env:"SIZE_Y" envDefault:"40"`
	TickTime  int `env:"TICK_TIME" envDefault:"80"`
	FoodCount int `env:"FOOD_COUNT" envDefault:"3"`
}

func ParseEnv() (*App, error) {
	cfg := App{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
