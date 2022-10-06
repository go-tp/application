package configs

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// 文件内容结构体映射
type Config struct {
	RunMod string        `yaml:"runMod"`
	Mysql  *MysqlConfig  `yaml:"mysql"`
	Redis  *RedisConfig  `yaml:"redis"`
	Logs   *LogsConfig   `yaml:"logs"`
	Device *DeviceConfig `yaml:"device"`
	JWT    *JwtConfig    `yaml:"jwt"`
	
}

type MysqlConfig struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Database string `yaml:"Database"`
}

type RedisConfig struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Db       string `yaml:"Db"`
}

type LogsConfig struct {
	Path string `yaml:"Path"`
}

type DeviceConfig struct {
	Cpu int `yaml:"Cpu"`
}

type JwtConfig struct {
	Secret string `yaml:"Secret"`
}

func ReadYaml() *Config {
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println("read yaml err: ", err)
		return nil
	}

	var _config *Config

	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println("Unmarshal yaml err: ", err)
		return nil
	}
	// fmt.Println("mysql host:", _config.Mysql.Host)
	return _config
}