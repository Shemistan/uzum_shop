package models

type Config struct {
	App APP      `envconfig:"APP"`
	DB  Postgres `envconfig:"POSTGRES"`
}

type APP struct {
	PortGRPC   string `envconfig:"PORT_GRPC"`
	PortHTTP   string `envconfig:"PORT_HTTP"`
	AuthClient string `envconfig:"AUTH_CLIENT"`
}

type Postgres struct {
	User     string `envconfig:"USER" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	Host     string `envconfig:"HOST" required:"true"`
	Port     string `envconfig:"PORT" required:"true"`
	Database string `envconfig:"DATABASE" required:"true"`
	SSLMode  string `envconfig:"SSL_MODE" default:"disable"`
}
