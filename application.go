package application

import (
	"context"
	"go.uber.org/dig"
)

type Application struct {
	container *dig.Container
	modules   []Module
	context   context.Context
}

func (a *Application) Modules() []Module {
	return a.modules
}

func (a *Application) Context() context.Context {
	return a.context
}

func (a *Application) Container() *dig.Container {
	return a.container
}

func New(ctx context.Context, modules []Module) (*Application, error) {
	container := dig.New()

	app := &Application{
		container: container,
		context:   ctx,
		modules:   modules,
	}

	// If not exist config set default
	cfg := app.provideConfig()

	// If not exist logger set default
	logger := app.provideLogger()

	app.modules = append([]Module{cfg, logger}, app.modules...)

	for _, module := range app.modules {
		if appListener, ok := module.(ModulusInitListener); ok {
			err := appListener.Init(app)
			if err != nil {
				return app, err
			}
		}
	}

	for _, module := range app.modules {
		if serviceProvider, ok := module.(ModulusServiceProvider); ok {
			if services := serviceProvider.ProvidedServices(); services != nil {
				for _, service := range services {
					err := app.container.Provide(service)
					if err != nil {
						return app, err
					}
				}
			}
		}
	}

	return app, nil
}

func (a *Application) Start() error {
	ch := make(chan error)

	for _, module := range a.modules {
		appListener, ok := module.(ModulusStartListener)
		if ok {
			go func() {
				err := appListener.Start(a)
				if err != nil {
					ch <- err
				}
			}()
		}
	}

	err := <-ch
	if err != nil {
		return err
	}

	return nil
}

func (a *Application) Stop() error {
	ch := make(chan error)

	for _, module := range a.modules {
		appListener, ok := module.(ModulusStopListener)
		if ok {
			go func() {
				err := appListener.Stop(a)
				if err != nil {
					ch <- err
				}
			}()
		}
	}

	err := <-ch
	if err != nil {
		return err
	}

	return nil
}

func (a *Application) Logger() Logger {
	var logger Logger
	err := a.container.Invoke(
		func(dep Logger) error {
			logger = dep
			return nil
		},
	)
	if err != nil || logger == nil {
		return NewDefaultLogger()
	}

	return logger
}

func (a *Application) Config() Config {
	var config Config
	err := a.container.Invoke(
		func(dep Config) error {
			config = dep
			return nil
		},
	)
	if err != nil || config == nil {
		return NewDefaultConfig()
	}

	return config
}

func (a *Application) provideLogger() Logger {
	var logger Logger
	for _, module := range a.modules {
		var ok bool
		if logger, ok = module.(Logger); ok {
			break
		}
	}
	if logger == nil {
		logger = NewDefaultLogger()
	}

	err := a.container.Provide(func() Logger { return logger })
	if err != nil {
		panic("Logger cannot be provided")
	}

	return logger
}

func (a *Application) provideConfig() Config {
	var cfg Config
	for _, module := range a.modules {
		var ok bool
		if cfg, ok = module.(Config); ok {
			break
		}
	}
	if cfg == nil {
		cfg = NewDefaultConfig()
	}

	err := a.container.Provide(func() Config { return cfg })
	if err != nil {
		panic("Config cannot be provided")
	}

	return cfg
}
