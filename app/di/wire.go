//+build wireinject

package di

import (
	"github.com/google/wire"
)

func NewApp() (*App, error) {
	wire.Build(AppSet, ProvideSet)
	return &App{}, nil
}
