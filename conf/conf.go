package conf

import (
	"errors"
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

var (
	Conf              config
	defaultConfigFile = "./conf/conf.toml"
)

type config struct {
	//App config
	App app `toml:"app"`
}

type app struct {
	Port    string `toml:"port"`
	Timeout int64  `toml:"timeout"`
}

func InitConfig(configFile string) (err error) {
	var configBytes []byte
	if configFile == "" {
		configBytes, err = ioutil.ReadFile(defaultConfigFile)
	} else {
		configBytes, err = ioutil.ReadFile(configFile)
	}

	if err != nil {
		return errors.New("config load err:" + err.Error())
	}
	_, err = toml.Decode(string(configBytes), &Conf)
	if err != nil {
		return errors.New("config decode err:" + err.Error())
	}
	return nil
}
