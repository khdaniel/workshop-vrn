package config

// Server config
type Server struct {
	host string `yaml:"host" env:"PORT"`
	port string `yaml:"port" env:"HOST" env-default:"0.0.0.0"`
}
