package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Support    Support    `json:"support"`
	Postgresql Postgresql `json:"postgresql"`
	Worker     Worker     `json:"worker"`
	Mail       Mail       `json:"mail"`
	Mongo      Mongo      `json:"mongo"`
	Redis      Redis      `json:"redis"`
	Key        Key        `json:"key"`
}

type Support struct {
	Security int      `json:"security"`
	Addr     string   `json:"addr"`     // start port
	Home     string   `json:"home"`     // website homepage
	Query    []string `json:"query"`    // Table fields allowed to be queried
	Ip       []string `json:"ip"`       // allowed ip
	Suffixes []string `json:"suffixes"` // email suffixes allowed to sign up
	Admin    []string `json:"admin"`    // admin email address
}

type Postgresql struct {
	User     []string `json:"user"`
	Password []string `json:"password"`
	Host     []string `json:"host"`
	Port     []string `json:"port"`
	Name     []string `json:"name"`
}

type Mongo struct {
	User     []string `json:"user"`
	Password []string `json:"password"`
	Host     []string `json:"host"`
	Port     []string `json:"port"`
}

type Redis struct {
	Host     []string `json:"host"`
	Port     []string `json:"port"`
	Password []string `json:"password"`
}

type Worker struct {
	WorkerId int64 `json:"worker_id"`
	CenterId int64 `json:"center_id"`
	Sequence int64 `json:"sequence"`
	Epoch    int64 `json:"epoch"`
}

type Mail struct {
	From     string `json:"from"`
	User     string `json:"user"`
	Password string `json:"password"`
	Smtp     string `json:"smtp"`
	Port     string `json:"port"`
}

type Key struct {
	Token      string `json:"token"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

func (config *Config) Get() *Config {
	file, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, config)
	if err != nil {
		panic(err)
	}
	return config
}

func (config *Config) Set(s Support) (bool, error) {
	file, err := os.ReadFile("./config.json")
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(file, &s)
	if err != nil {
		return false, err
	}
	value, err := json.Marshal(s)
	if err != nil {
		return false, err
	}
	_, err = os.Stdout.Write(value)
	if err != nil {
		return false, err
	}
	return true, nil
}
