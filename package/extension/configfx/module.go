package configfx

import (
	"go.uber.org/fx"
)

type Option[T any] struct {
	Variable T
}

func ForRoot[T any](option Option[T]) fx.Option {
	return fx.Module("configfx",

		fx.Provide(
			func() IConfig[T] {
				return NewConfig[T](option.Variable)
			},
		),
	)
}

type IConfig[T any] interface {
	Get() T
}

type config[T any] struct {
	conf T
}

// Get implements IConfig.
func (c *config[T]) Get() T {

	return c.conf
}

func NewConfig[T any](conf T) IConfig[T] {
	return &config[T]{
		conf: conf,
	}

}
