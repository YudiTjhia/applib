package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type ConnectionIDConf struct {
	ID            string `json:"id"`
	DataLayerType string `json:"dataLayerType"`
	ConnectionID  string `json:"connectionID"`
}

var _connectionIDs map[string]ConnectionIDConf
var connIDMutex = sync.RWMutex{}


func GetConnectionID(connectionIDFile string, connID string) ConnectionIDConf {

	if _connectionIDs == nil {

		fopen, err := os.Open(connectionIDFile)
		if err != nil {
			panic(err)
		}

		bytesVal, err := ioutil.ReadAll(fopen)
		if err != nil {
			panic(err)
		}

		connIDConfs := []ConnectionIDConf{}
		err = json.Unmarshal(bytesVal, &connIDConfs)
		if err != nil {
			panic(err)
		}

		if len(connIDConfs) == 0 {
			fmt.Print("No connectionIDs defined")
			panic("No connectionIDs defined")
		}

		connIDMutex.Lock()
		_connectionIDs = map[string]ConnectionIDConf{}
		for _, connIDConf := range connIDConfs {
			_, e := _connectionIDs[connIDConf.ID]
			if !e {
				_connectionIDs[connIDConf.ID] = connIDConf
			} else {
				fmt.Print("Duplicate DbConf.ID =" + connIDConf.ID)
				panic("Duplicate DbConf.ID =" + connIDConf.ID)
			}
		}
		connIDMutex.Unlock()
	}

	connIDConf, e := _connectionIDs[connID]
	if !e {
		fmt.Print("Cannot find connIDConf=" + connID)
		panic("Cannot find connIDConf=" + connID)
	}

	return connIDConf

}



func GetConnectionID2(connectionIDFile string, connID string) (ConnectionIDConf, error) {

	if _connectionIDs == nil {
		fopen, err := os.Open(connectionIDFile)
		if err != nil {
			//panic(err)
			return ConnectionIDConf{}, err

		}

		bytesVal, err := ioutil.ReadAll(fopen)
		if err != nil {
			//panic(err)
			return ConnectionIDConf{}, err
		}

		connIDConfs := []ConnectionIDConf{}
		err = json.Unmarshal(bytesVal, &connIDConfs)
		if err != nil {
			//panic(err)
			return ConnectionIDConf{}, err
		}

		if len(connIDConfs) == 0 {
			fmt.Print("No connectionIDs defined")
			//panic("No connectionIDs defined")
			return ConnectionIDConf{}, errors.New("No connectionIDs defined")
		}


		connIDMutex.Lock()
		_connectionIDs = map[string]ConnectionIDConf{}
		for _, connIDConf := range connIDConfs {
			_, e := _connectionIDs[connIDConf.ID]
			if !e {
				_connectionIDs[connIDConf.ID] = connIDConf
			} else {
				fmt.Print("Duplicate DbConf.ID =" + connIDConf.ID)
				//panic("Duplicate DbConf.ID =" + connIDConf.ID)
				return ConnectionIDConf{}, errors.New("Duplicate DbConf.ID =" + connIDConf.ID)
			}
		}
		connIDMutex.Unlock()
	}

	connIDConf, e := _connectionIDs[connID]
	if !e {
		fmt.Print("Cannot find connIDConf=" + connID)
		return ConnectionIDConf{}, errors.New("Cannot find connIDConf=" + connID)
	}

	return connIDConf, nil

}
