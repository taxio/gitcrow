package di

import (
	"github.com/google/wire"
	"github.com/spf13/cobra"
)

type App struct {
	Cmd *cobra.Command
}

var AppSet = wire.NewSet(wire.Struct(new(App), "*"))
