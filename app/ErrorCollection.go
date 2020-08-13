package app

import "encoding/json"
import (
	"errors"
)

type ErrorCollection struct {
	Errors []string `json:"errors"`
}

func (col *ErrorCollection) Add(err error) {
	if err != nil {
		col.Errors = append(col.Errors, err.Error())
	}
}

func (col *ErrorCollection) AddString(errString string) {
	col.Errors = append(col.Errors, errString)
}

func (col *ErrorCollection) AddCollection(errCol ErrorCollection) {
	if errCol.HasErrors() {
		for _, er := range errCol.Errors {
			col.Errors = append(col.Errors, er)
		}
	}
}

func MarshallError(err error) ErrorCollection {
	errs := ErrorCollection{}
	json.Unmarshal([]byte(err.Error()), &errs)
	return errs
}

func (col ErrorCollection) HasErrorOf(msg string) bool {
	found := false
	for _, err := range col.Errors {
		if err == msg {
			found = true
			break
		}
	}
	return found
}

func (col ErrorCollection) HasErrors() bool {
	return (len(col.Errors) > 0)
}

func (col ErrorCollection) HasNoErrors() bool {
	return (len(col.Errors) == 0)
}

func (col ErrorCollection) Error() error {
	return errors.New(col.Errors[0])
}

func (col *ErrorCollection) Reset() {
	col.Errors = []string{}
}
