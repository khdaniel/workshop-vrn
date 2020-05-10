package config

// Server config
type Server struct {
	Host string `yaml:"host" env:"HOST"`
	Port string `yaml:"port" env:"PORT" env-default:"0.0.0.0"`

	JokeURL string `yaml:"joke-url" env:"JOKE_URL"`
}
