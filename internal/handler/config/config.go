package config

// Server config
type Server struct {
	Host string `yaml:"host" env:"PORT"`
	Port string `yaml:"port" env:"HOST" env-default:"0.0.0.0"`
}
