package config

type Config struct {
	Server   Server
	Database Database
	Mail     Email
	Redis    Redis
}

type Server struct {
	Host string
	Port string
}

type Redis struct {
	Addr string
	Pass string
}

type Database struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

type Email struct {
	Host string
	User string
	Pass string
	Port string
}
