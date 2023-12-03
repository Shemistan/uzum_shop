package models

type Config struct {
	Host string `envconfig:"HOST"`

	GrpcPort        string `envconfig:"GRPC_PORT"`
	HttpPort        string `envconfig:"HTTP_PORT"`
	LoginClientPort string `envconfig:"LOGIN_CLIENT_PORT"`

	DbConnStr string `envconfig:"POSTGRES_DSN"`
}
