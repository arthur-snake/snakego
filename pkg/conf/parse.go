package conf

import "github.com/caarlos0/env/v6"

func ParseEnv() (*App, error) {
	cfg := App{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
