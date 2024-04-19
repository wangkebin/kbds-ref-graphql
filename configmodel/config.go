package configmodel

type Config struct {
	ConnString string `mapstructure:"connection_string"`
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	Debug      bool   `mapstructure:"debug"`
}

var GlobalConfig Config
