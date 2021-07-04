package conf

type App struct {
	PrometheusBind string `env:"PROMETHEUS_BIND" envDefault:":2112"`
	ServerBind     string `env:"SERVER_BIND" envDefault:":8080"`

	DefaultServer Server
}
