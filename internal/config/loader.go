package config

type Config struct {
	SRV Server
	DB  Database
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}
