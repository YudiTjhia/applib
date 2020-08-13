package conf

import (
	"encoding/json"
	"os"
	"sync"
)

type AppConf struct {
	BasePath  string `json:"basePath"`
	ClientUrl  string `json:"clientUrl"`
	EnableLog bool   `json:"enableLog"`
	EnableNotification bool `json:"enableNotification"`
	NotificationDelay int `json:"notificationDelay"`
	EnableReminder bool `json:"enableReminder"`
	ReminderDelay int `json:"reminderDelay"`
	
}

var (
	_appConf     AppConf
	_appConfRead = false
)

var appConfMutex = sync.RWMutex{}

func GetAppConf(configFile string) AppConf {

	if _appConfRead == false {

		appConfMutex.Lock()
		f, err := os.Open(configFile)
		defer f.Close()
		if err != nil {
			panic(err)
		}

		decoder := json.NewDecoder(f)
		err = decoder.Decode(&_appConf)
		if err != nil {
			panic(err)
		}

		_appConfRead = true
		appConfMutex.Unlock()

	}

	return _appConf

}

func IsEnableLog(appConfigFile string) bool {

	if !_appConfRead {
		GetAppConf(appConfigFile)
	}

	return _appConf.EnableLog

}
