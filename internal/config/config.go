package config

type Server struct {
	Port    string `yaml:"port"`
	Host    string `yaml:"host"`
	JokeUrl string `yaml:"jokeUrl"`
}
