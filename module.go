package application

type Module interface{}

// ModulusInitListener if service provider implements this method it will be called after
// providing dependencies of the module
type ModulusInitListener interface {
	// Init is called for each module to initialize module's variables
	Init(app *Application) error
}

// ModulusServiceProvider describes all services of a module in the dependency injection container
// it is the first step of the application running
type ModulusServiceProvider interface {
	// ProvidedServices returns a list of constructors that presents all services of the module.
	// All of them will be placed in the dependency injection container
	ProvidedServices() []interface{}
}

// ModulusStartListener if service provider implements this method it will be called after
// initializing the routes
type ModulusStartListener interface {
	// Start module's application such as a web-server
	Start(app *Application) error
}

// ModulusStopListener if service provider implements this method it will be called after
// stopping the application
type ModulusStopListener interface {
	// Stop may close some resources of a module, for example a db connection
	Stop(app *Application) error
}
