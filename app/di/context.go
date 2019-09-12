package di

var (
	appName    = "gitcrow"
	appVersion = "v0.0.1"
)

type AppContext struct {
	Name    string
	Version string
}

func provideAppContext() *AppContext {
	return &AppContext{
		Name:    appName,
		Version: appVersion,
	}
}
