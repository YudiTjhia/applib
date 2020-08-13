package funcs

import (
	"applib/constant"
	"errors"
	"fmt"
)

func CannotAddNew(entityName string) error {
	return errors.New(fmt.Sprintf(constant.CANNOT_ADD, entityName))
}
func CannotEdit(entityName string) error {
	return errors.New(fmt.Sprintf(constant.CANNOT_EDIT, entityName))
}
func CannotDelete(entityName string) error {
	return errors.New(fmt.Sprintf(constant.CANNOT_DELETE, entityName))
}
func CannotFind(entityName string) error {
	return errors.New(fmt.Sprintf(constant.CANNOT_FIND, entityName))
}
func CannotFindStr(entityName string) string {
	return fmt.Sprintf(constant.CANNOT_FIND, entityName)
}

func CannotFindProperty(entityName string, propertyName string) error {
	return errors.New(fmt.Sprintf(constant.CANNOT_FIND_PROPERTY, entityName, propertyName))
}

func MustBeUnique(entityName string, propertyName string) error {
	return errors.New(fmt.Sprintf(constant.MUST_BE_UNIQUE, entityName, propertyName))
}

func IsRequired(entityName string, propertyName string) error {
	return errors.New(fmt.Sprintf(constant.IS_REQUIRED, entityName, propertyName))
}
func MinLength(entityName string, propertyName string, minlength int) error {
	return errors.New(fmt.Sprintf(constant.MIN_LENGTH, entityName, propertyName, minlength))
}
func MaxLength(entityName string, propertyName string, maxlength int) error {
	return errors.New(fmt.Sprintf(constant.MAX_LENGTH, entityName, propertyName, maxlength))
}

func MustBeGtz(entityName string, propertyName string) error {
	return errors.New(fmt.Sprintf(constant.MUST_BE_GTZ, entityName, propertyName))
}
func MustBeGtez(entityName string, propertyName string) error {
	return errors.New(fmt.Sprintf(constant.MUST_BE_GTEZ, entityName, propertyName))
}
func MustBeIntRange(entityName string, propertyName string, min int, max int) error {
	return errors.New(fmt.Sprintf(constant.MUST_BE_INT_RANGE, entityName, propertyName, min, max))
}
func MustBeFlRange(entityName string, propertyName string, min float64, max float64) error {
	return errors.New(fmt.Sprintf(constant.MUST_BE_FL_RANGE, entityName, propertyName, min, max))
}
