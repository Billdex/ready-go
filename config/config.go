package config

import "github.com/BurntSushi/toml"

var cfg Configuration

type Configuration struct {
	App   AppConfig `toml:"app"`
	Mysql DBConfig  `toml:"mysql"`
	Log   LogConfig `toml:"log"`
}

type AppConfig struct {
	Address string `toml:"address"`
	RunMode string `toml:"run_mode"`
}

type DBConfig struct {
	DSN             string `toml:"dsn"`
	MaxOpenConns    int    `toml:"max_open_conns"`
	MaxIdleConns    int    `toml:"max_idle_conns"`
	ConnMaxLifetime int    `toml:"conn_max_lifetime"`
}

type LogConfig struct {
	Level string `toml:"level"`
	Style string `toml:"style"`
}

func InitConfig(path string) error {
	_, err := toml.DecodeFile(path, &cfg)
	return err
}

func GetConfig() Configuration {
	return cfg
}
