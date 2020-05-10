package config

// Server config
type Server struct {
	Host string `yaml:"host" env:"PORT"`
	Port string `yaml:"port" env:"host" env-default:"0.0.0.0"`
}
