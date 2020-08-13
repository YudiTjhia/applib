package conf

import (
	"encoding/json"
	"os"
	"strconv"
)

type ServerConf struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	WriteTimeout int    `json:"writeTimeout"`
	ReadTimeout  int    `json:"readTimeout"`
	AccountID    string `json:"accountID"`
	WindowsService bool `json:"windowsService"`
}

func (serverConf ServerConf) GetUrl() string {
	return serverConf.Host + ":" + strconv.Itoa(serverConf.Port)
}

func (serverConf *ServerConf) Load(configFile string) {
	f, err := os.Open(configFile)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&serverConf)
	if err != nil {
		panic(err)
	}
}
