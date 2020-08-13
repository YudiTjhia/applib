package lang

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"sync"
)

var _lang map[string]string
var langMutex = sync.RWMutex{}

func firstUpper(key string, delimiter string, newSeparator string) string {
	splKeys := strings.Split(key, delimiter)
	resKey := ""
	for i:=0;i<len(splKeys);i++ {
		runKey := []rune(splKeys[i])
		word := strings.ToUpper(string(runKey[0])) + strings.ToLower(string(runKey[1:]))
		resKey += word + newSeparator
	}
	return resKey
}

func LoadLocal(filename string) {
	langMutex.Lock()
	bytes, err := ioutil.ReadFile(filename)
	if err==nil {
		json.Unmarshal(bytes, &_lang)
	}
	langMutex.Unlock()
}

func Local(key string) string {
	val, e:= _lang[key]
	if e {
		return val
	} else {
		return firstUpper(key, "_"," ")
	}
}
