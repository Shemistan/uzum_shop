package models

type Config struct {
	Host string

	GrpcPort string
	HttpPort string
	LoginClientPort string

	DbConnStr string
}
