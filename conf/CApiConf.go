package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type CApiConf struct {
	ID       string `json:"id"`
	User     string `json:"user"`
	Passwd   string `json:"passwd"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
	BasePath string `json:"basePath"`
	Protocol string `json:"protocol"`
}

var _cApiConfs map[string]CApiConf

func GetCApi(configFile string) map[string]CApiConf {
	if _cApiConfs == nil {
		readCApiConfig(configFile)
	}
	return _cApiConfs
}

func GetCApiConfig(configFile string, connectionID string) CApiConf {
	if _cApiConfs == nil {
		readCApiConfig(configFile)
	}

	conf, e := _cApiConfs[connectionID]
	if !e {
		fmt.Println("Cannot find CApiConf. ConnectionID=" + connectionID)
	}
	return conf
}

func SerializeCApiConf(conf DbConf) string {
	b, _ := json.Marshal(conf)
	return string(b)
}

func readCApiConfig(fileName string) {
	fopen, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	bytesVal, err := ioutil.ReadAll(fopen)
	if err != nil {
		panic(err)
	}

	confs := []CApiConf{}
	err = json.Unmarshal(bytesVal, &confs)
	if err != nil {
		panic(err)
	}

	if len(confs) == 0 {
		panic("No CApiConfs defined")
	}

	_cApiConfs = map[string]CApiConf{}
	for _, conf := range confs {
		_, e := _cApiConfs[conf.ID]
		if !e {
			_cApiConfs[conf.ID] = conf
		} else {
			panic("Duplicate CApiConf.ID =" + conf.ID)
		}
	}
}
