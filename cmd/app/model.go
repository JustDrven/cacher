package app

type Application struct {
	Debug  bool
	Config ApplicationConfig
}

type ApplicationConfig struct {
	Addr string
}
