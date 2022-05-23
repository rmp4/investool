package config

import (
	"fmt"
	"investool/pkg/log"
	"investool/pkg/utils"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Server   Server   `toml:"server"`
	Database Database `toml:"database"`
	Redis    Redis    `toml:"redis"`
}

type Server struct {
	RunMode      string        `toml:"RunMode"`
	HTTPPort     int           `toml:"HttpPort"`
	ReadTimeout  time.Duration `toml:"ReadTimeout"`
	WriteTimeout time.Duration `toml:"WriteTimeout"`
}
type Database struct {
	Type        string `toml:"Type"`
	User        string `toml:"User"`
	Password    string `toml:"Password"`
	Host        string `toml:"Host"`
	Name        string `toml:"Name"`
	TablePrefix string `toml:"TablePrefix"`
}
type Redis struct {
	Host        string `toml:"Host"`
	Password    string `toml:"Password"`
	MaxIdle     int    `toml:"MaxIdle"`
	MaxActive   int    `toml:"MaxActive"`
	IdleTimeout int    `toml:"IdleTimeout"`
}

var config *Config

func Init(fileName string) error {
	viper.SetConfigName(fileName)
	viper.SetConfigType("toml")
	dir, err := utils.GetCurrentPath()
	if err == nil {
		viper.AddConfigPath(dir + "/configs/")
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		// fmt.Println("Config file changed:", e.Name)
		log.ServiceLog().Debug(fmt.Sprintln("Config file changed:", e.Name))
		reloadConfig()
	})
	viper.WatchConfig()

	return readConfig()
}

func readConfig() error {
	if err := viper.ReadInConfig(); err == nil { // Find and read the config file
		log.ServiceLog().Debug(fmt.Sprintln("Using config file:", viper.ConfigFileUsed()))
		config = &Config{}
		viper.Unmarshal(config)
		log.ServiceLog().Info(fmt.Sprintf("Config:%+v\n", config))
	} else { // Handle errors reading the config file
		// load default config
		log.ServiceLog().Error(fmt.Sprintln("config not find", viper.ConfigFileUsed(), err.Error()))
		return err
	}
	return nil
}

func reloadConfig() {

}

func Get() Config {
	if config == nil {
		viper.Unmarshal(config)
	}
	return *config
}
