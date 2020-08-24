package conf

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

func NewConfig(environment string) Config {

	// configure logging
	logging()
	config := &Config{}

	log.Info().Msg("stuff goes here")

	filePath := "../../conf/" + environment + ".yml"
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal().Msg("Failed to open config")
	}

	d := yaml.NewDecoder(file)
	if err := d.Decode(config); err != nil {
		log.Fatal().Msg("Failed to load config")
	}

	return *config

}

func logging() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}

// Config struct for webapp config
type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
type Datasource struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Config struct {
	Server     `yaml:"server"`
	Datasource `yaml:"datasource"`
}
