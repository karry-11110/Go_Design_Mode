package build

import "fmt"

type Player struct {
	name   string
	number int
}
type Config struct {
	number int
}
type ConfigOptFunc func(option *Config)

func NewPlayer(name string, opts ...ConfigOptFunc) (*Player, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	option := &Config{
		number: 30,
	}
	for _, opt := range opts {
		opt(option)
	}
	if option.number < 0 {
		return nil, fmt.Errorf("error")
	}
	return &Player{
		name:   name,
		number: option.number,
	}, nil
}
