package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Postgresql Postgresql `yaml:"postgresql"`
	Worker     Worker     `yaml:"worker"`
	Mail       Mail       `yaml:"mail"`
	Mongo      Mongo      `yaml:"mongo"`
	Redis      Redis      `yaml:"redis"`
	Key        Key        `yaml:"key"`
	MobTech    MobTech    `yaml:"mob_tech"`
}

type Postgresql struct {
	User     []string `yaml:"user"`
	Password []string `yaml:"password"`
	Host     []string `yaml:"host"`
	Port     []string `yaml:"port"`
	Name     []string `yaml:"name"`
}

type Worker struct {
	WorkerId int64 `yaml:"workerId"`
	CenterId int64 `yaml:"centerId"`
	Sequence int64 `yaml:"sequence"`
	Epoch    int64 `yaml:"epoch"`
}

type Mail struct {
	From     string   `yaml:"from"`
	User     string   `yaml:"user"`
	Password string   `yaml:"password"`
	Admin    []string `yaml:"admin"`
	Smtp     string   `yaml:"smtp"`
	Port     int      `yaml:"port"`
}

type Mongo struct {
	User     []string `yaml:"user"`
	Password []string `yaml:"password"`
	Host     []string `yaml:"host"`
	Port     []string `yaml:"port"`
}

type Redis struct {
	Host     []string `yaml:"host"`
	Port     []string `yaml:"port"`
	Password []string `yaml:"password"`
}

type Key struct {
	Token      []byte `yaml:"token"`
	PrivateKey string `yaml:"privateKey"`
	PublicKey  string `yaml:"publicKey"`
}

type MobTech struct {
	Key    string `yaml:"key"`
	Secret string `yaml:"secret"`
}

func (config *Config) Yaml() *Config {
	filename, _ := os.ReadFile("./config.yaml")

	_ = yaml.Unmarshal(filename, config)

	return config
}
