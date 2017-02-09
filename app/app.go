package app

func Setup() (*App, error) {
	// create some configurations

	return NewApp(nil), nil
}

func NewApp(config *Config) *App {
	return &App {
		config,
	}
}

type App struct {
	*Config
}

func (app *App) Run() error {
	// run stuff forever in here
	return nil
}