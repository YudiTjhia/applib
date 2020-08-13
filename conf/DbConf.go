package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type DbConf struct {
	ID       string `json:"id"`
	User     string `json:"user"`
	Passwd   string `json:"passwd"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
	DbName   string `json:"dbName"`
	DbType   string `json:"dbType"`
	BasePath string `json:"basePath"`
	Protocol string `json:"protocol"`
	AppName  string `json:"appName"`
	Singular bool `json:"singular"`
	Debug bool `json:"debug"`
}

var _dbConfs map[string]DbConf
var dbConfMutex = sync.RWMutex{}

func GetDbConfigs(configFile string) map[string]DbConf {
	if _dbConfs == nil {
		readDbConfig(configFile)
	}
	return _dbConfs
}


func GetDbConfigs2(configFile string) (map[string]DbConf, error) {
	if _dbConfs == nil {
		err := readDbConfig2(configFile)
		if err!=nil {
			return nil, err
		}
	}

	return _dbConfs, nil
}


func GetDbConfig(configFile string, connectionID string) DbConf {
	if _dbConfs == nil {
		readDbConfig(configFile)
	}

	dbConf, e := _dbConfs[connectionID]
	if !e {
		fmt.Println("Cannot find DbConf. ConnectionIDConf=" + connectionID)
	}
	return dbConf
}

func SerializeDbConf(conf DbConf) string {
	b, _ := json.Marshal(conf)
	return string(b)
}

func readDbConfig(fileName string) {

	fopen, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	bytesVal, err := ioutil.ReadAll(fopen)
	if err != nil {
		panic(err)
	}

	dbConfs := []DbConf{}
	err = json.Unmarshal(bytesVal, &dbConfs)
	if err != nil {
		panic(err)
	}

	if len(dbConfs) == 0 {
		panic("No DbConfs defined")
	}

	dbConfMutex.Lock()
	_dbConfs = map[string]DbConf{}
	for _, dbConf := range dbConfs {
		_, e := _dbConfs[dbConf.ID]
		if !e {

			_dbConfs[dbConf.ID] = dbConf
		} else {
			panic("Duplicate DbConf.ID =" + dbConf.ID)
		}
	}

	dbConfMutex.Unlock()

}


func readDbConfig2(fileName string) error {

	fopen, err := os.Open(fileName)
	if err != nil {
		//panic(err)
		return err
	}

	bytesVal, err := ioutil.ReadAll(fopen)
	if err != nil {
		//panic(err)
		return err
	}

	dbConfs := []DbConf{}
	err = json.Unmarshal(bytesVal, &dbConfs)
	if err != nil {
		//panic(err)
		return err
	}

	if len(dbConfs) == 0 {
		//panic("No DbConfs defined")
		return errors.New("No DbConfs defined")
	}

	dbConfMutex.Lock()
	_dbConfs = map[string]DbConf{}
	for _, dbConf := range dbConfs {
		_, e := _dbConfs[dbConf.ID]
		if !e {

			_dbConfs[dbConf.ID] = dbConf
		} else {
			//panic("Duplicate DbConf.ID =" + dbConf.ID)
			return errors.New("Duplicate DbConf.ID =" + dbConf.ID)
		}
	}

	dbConfMutex.Unlock()

	return nil

}
