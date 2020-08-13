package validator

import (
	"applib/app"
	"applib/funcs"
	"strings"
)

type IValidator interface {
	Execute() (bool, error)
	SetValue(value interface{})
}

type IValidatorObserve interface {
	GetValue(name string) (interface{}, error)
}

type IObjectValidator interface {
	Validate(fields []string) app.ErrorCollection
}

type ObjectValidator struct {
	validators       map[string][]IValidator
	appErr           app.ErrorCollection
	validatorObserve IValidatorObserve
}

const (
	DATA_TYPE_STRING  = "string"
	DATA_TYPE_INT     = "int"
	DATA_TYPE_FLOAT64 = "float64"

	VALIDATOR_TYPE_REQUIRED   = "required"
	VALIDATOR_TYPE_MIN_LENGTH = "min_length"
	VALIDATOR_TYPE_MAX_LENGTH = "max_length"

	VALIDATOR_TYPE_FL64_GTZ   = "fl64_gtz"
	VALIDATOR_TYPE_FL64_GTEZ  = "fl64_gtez"
	VALIDATOR_TYPE_FL64_RANGE = "fl64_range"

	VALIDATOR_TYPE_INT_GTZ   = "int_gtz"
	VALIDATOR_TYPE_INT_GTEZ  = "int_gtez"
	VALIDATOR_TYPE_INT_RANGE = "int_range"
)

func (objectValidator *ObjectValidator) AddRequiredValidator(entityName string, field string) {
	requiredValidator := RequiredValidator{}
	requiredValidator.SetName(entityName, field)
	objectValidator.AddValidator(field, &requiredValidator)
}

func (objectValidator *ObjectValidator) AddMinLengthValidator(entityName string, field string, minlength int) {
	minLengthValidator := MinLengthValidator{}
	minLengthValidator.SetName(entityName, field)
	minLengthValidator.SetMinLength(minlength)
	objectValidator.AddValidator(field, &minLengthValidator)
}

func (objectValidator *ObjectValidator) AddMaxLengthValidator(entityName string, field string, maxlength int) {
	maxLengthValidator := MaxLengthValidator{}
	maxLengthValidator.SetName(entityName, field)
	maxLengthValidator.SetMaxLength(maxlength)
	objectValidator.AddValidator(field, &maxLengthValidator)
}

func (objectValidator *ObjectValidator) AddFl64GtzValidator(entityName string, field string) {
	validator := Fl64GtzValidator{}
	validator.SetName(entityName, field)
	objectValidator.AddValidator(field, &validator)
}



func (objectValidator *ObjectValidator) AddFl64GtezValidator(entityName string, field string) {
	validator := Fl64GtezValidator{}
	validator.SetName(entityName, field)
	objectValidator.AddValidator(field, &validator)
}

func (objectValidator *ObjectValidator) AddFlRangeValidator(entityName string, field string, min float64, max float64) {
	validator := FlRangeValidator{}
	validator.SetName(entityName, field)
	validator.SetRange(min, max)
	objectValidator.AddValidator(field, &validator)
}

func (objectValidator *ObjectValidator) AddIntGtzValidator(entityName string, field string) {
	validator := IntGtzValidator{}
	validator.SetName(entityName, field)
	objectValidator.AddValidator(field, &validator)
}

func (objectValidator *ObjectValidator) AddIntGtezValidator(entityName string, field string) {
	validator := IntGtezValidator{}
	validator.SetName(entityName, field)
	objectValidator.AddValidator(field, &validator)
}

func (objectValidator *ObjectValidator) AddIntRangeValidator(entityName string, field string, min int, max int) {
	validator := IntRangeValidator{}
	validator.SetName(entityName, field)
	validator.SetRange(min, max)
	objectValidator.AddValidator(field, &validator)
}

func (objectValidator *ObjectValidator) AddValidator(name string, validator IValidator) {
	if objectValidator.validators == nil {
		objectValidator.validators = map[string][]IValidator{}
	}

	_, e := objectValidator.validators[name]
	if !e {
		objectValidator.validators[name] = []IValidator{}
	}
	objectValidator.validators[name] = append(objectValidator.validators[name], validator)
}

func (objectValidator *ObjectValidator) SetValidatorObserve(observe IValidatorObserve) {
	objectValidator.validatorObserve = observe
}

func (objectValidator *ObjectValidator) Validate(fields []string) app.ErrorCollection {
	objectValidator.appErr.Reset()
	for i := 0; i < len(fields); i++ {
		field := fields[i]
		data, err := objectValidator.validatorObserve.GetValue(field)
		if err == nil {
			validators, e := objectValidator.validators[field]
			if e {
				for j := 0; j < len(validators); j++ {
					validator := validators[j]
					validator.SetValue(data)
					b, err := validator.Execute()
					if !b {
						objectValidator.appErr.Add(err)
					}
				}
			}
		}
	}

	return objectValidator.appErr
}

func (objectValidator *ObjectValidator) Validate2(observe IValidatorObserve, fields []string) app.ErrorCollection {
	objectValidator.appErr.Reset()
	for i := 0; i < len(fields); i++ {
		field := fields[i]
		data, err := observe.GetValue(field)
		if err == nil {
			validators, e := objectValidator.validators[field]
			if e {
				for j := 0; j < len(validators); j++ {
					validator := validators[j]
					validator.SetValue(data)
					b, err := validator.Execute()
					if !b {
						objectValidator.appErr.Add(err)
					}
				}
			}
		}
	}

	return objectValidator.appErr
}

type MaxLengthValidator struct {
	entityName   string
	propertyName string
	value        string
	maxLength    int
}

func (validator *MaxLengthValidator) SetMaxLength(maxLength int) {
	validator.maxLength = maxLength
}
func (validator *MaxLengthValidator) SetValue(value interface{}) {
	validator.value = value.(string)
}
func (validator *MaxLengthValidator) SetName(entityName string, propertyName string) {
	validator.entityName = entityName
	validator.propertyName = propertyName
}
func (validator *MaxLengthValidator) GetType() string {
	return VALIDATOR_TYPE_MAX_LENGTH
}
func (validator MaxLengthValidator) Execute() (bool, error) {
	if len(validator.value) > validator.maxLength {
		return false, funcs.MaxLength(validator.entityName, validator.propertyName, validator.maxLength)
	}
	return true, nil
}

type MinLengthValidator struct {
	entityName   string
	propertyName string
	value        string
	minLength    int
}

func (validator *MinLengthValidator) SetMinLength(minLength int) {
	validator.minLength = minLength
}
func (validator *MinLengthValidator) SetValue(value interface{}) {
	validator.value = value.(string)
}
func (validator *MinLengthValidator) SetName(entityName string, propertyName string) {
	validator.entityName = entityName
	validator.propertyName = propertyName
}
func (validator *MinLengthValidator) GetType() string {
	return VALIDATOR_TYPE_MIN_LENGTH
}
func (validator MinLengthValidator) Execute() (bool, error) {
	if len(validator.value) < validator.minLength {
		return false, funcs.MinLength(validator.entityName, validator.propertyName, validator.minLength)
	}
	return true, nil
}

type RequiredValidator struct {
	entityName   string
	propertyName string
	value        string
}

func (validator *RequiredValidator) SetValue(value interface{}) {
	validator.value = value.(string)
}
func (validator *RequiredValidator) SetName(entityName string, propertyName string) {
	validator.entityName = entityName
	validator.propertyName = propertyName
}
func (validator RequiredValidator) GetType() string {
	return VALIDATOR_TYPE_REQUIRED
}
func (validator RequiredValidator) Execute() (bool, error) {
	if strings.TrimSpace(validator.value) == "" {
		return false, funcs.IsRequired(validator.entityName, validator.propertyName)
	}
	return true, nil
}

type Fl64GtzValidator struct {
	entityName   string
	propertyName string
	value        float64
}

func (validator *Fl64GtzValidator) SetValue(value interface{}) {
	validator.value = value.(float64)
}
func (validator *Fl64GtzValidator) SetName(entityName string, propertyName string) {
	validator.entityName = entityName
	validator.propertyName = propertyName
}
func (validator Fl64GtzValidator) GetType() string {
	return VALIDATOR_TYPE_FL64_GTZ
}
func (validator Fl64GtzValidator) Execute() (bool, error) {
	if validator.value <= float64(0) {
		return false, funcs.MustBeGtz(validator.entityName, validator.propertyName)
	}
	return true, nil
}

type Fl64GtezValidator struct {
	entityName   string
	propertyName string
	value        float64
}

func (validator *Fl64GtezValidator) SetValue(value interface{}) {
	validator.value = value.(float64)
}
func (validator *Fl64GtezValidator) SetName(entityName string, propertyName string) {
	validator.entityName = entityName
	validator.propertyName = propertyName
}
func (validator Fl64GtezValidator) GetType() string {
	return VALIDATOR_TYPE_FL64_GTEZ
}
func (validator Fl64GtezValidator) Execute() (bool, error) {
	if validator.value < float64(0) {
		return false, funcs.MustBeGtez(validator.entityName, validator.propertyName)
	}
	return true, nil
}

type IntGtzValidator struct {
	entityName   string
	propertyName string
	value        int
}

func (validator *IntGtzValidator) SetValue(value interface{}) {
	validator.value = value.(int)
}
func (validator *IntGtzValidator) SetName(entityName string, propertyName string) {
	validator.entityName = entityName
	validator.propertyName = propertyName
}
func (validator IntGtzValidator) GetType() string {
	return VALIDATOR_TYPE_INT_GTZ
}
func (validator IntGtzValidator) Execute() (bool, error) {
	if validator.value <= int(0) {
		return false, funcs.MustBeGtz(validator.entityName, validator.propertyName)
	}
	return true, nil
}

type IntGtezValidator struct {
	entityName   string
	propertyName string
	value        int
}

func (validator *IntGtezValidator) SetValue(value interface{}) {
	validator.value = value.(int)
}
func (validator *IntGtezValidator) SetName(entityName string, propertyName string) {
	validator.entityName = entityName
	validator.propertyName = propertyName
}
func (validator IntGtezValidator) GetType() string {
	return VALIDATOR_TYPE_INT_GTEZ
}
func (validator IntGtezValidator) Execute() (bool, error) {
	if validator.value < int(0) {
		return false, funcs.MustBeGtez(validator.entityName, validator.propertyName)
	}
	return true, nil
}

type IntRangeValidator struct {
	entityName   string
	propertyName string
	lower        int
	upper        int
	value        int
}

func (validator *IntRangeValidator) SetValue(value interface{}) {
	validator.value = value.(int)
}
func (validator *IntRangeValidator) SetRange(min int, max int) {
	validator.lower = min
	validator.upper = max
}
func (validator *IntRangeValidator) SetName(entityName string, propertyName string) {
	validator.entityName = entityName
	validator.propertyName = propertyName
}
func (validator IntRangeValidator) GetType() string {
	return VALIDATOR_TYPE_INT_RANGE
}
func (validator IntRangeValidator) Execute() (bool, error) {
	if validator.value < validator.lower || validator.value > validator.upper {
		return false, funcs.MustBeIntRange(validator.entityName, validator.propertyName,
			validator.lower, validator.upper)
	}
	return true, nil
}

type FlRangeValidator struct {
	entityName   string
	propertyName string
	lower        float64
	upper        float64
	value        float64
}

func (validator *FlRangeValidator) SetValue(value interface{}) {
	validator.value = value.(float64)
}
func (validator *FlRangeValidator) SetRange(min float64, max float64) {
	validator.lower = min
	validator.upper = max
}
func (validator *FlRangeValidator) SetName(entityName string, propertyName string) {
	validator.entityName = entityName
	validator.propertyName = propertyName
}
func (validator FlRangeValidator) GetType() string {
	return VALIDATOR_TYPE_FL64_RANGE
}
func (validator FlRangeValidator) Execute() (bool, error) {
	if validator.value < validator.lower || validator.value > validator.upper {
		return false, funcs.MustBeFlRange(validator.entityName, validator.propertyName,
			validator.lower, validator.upper)
	}
	return true, nil
}
