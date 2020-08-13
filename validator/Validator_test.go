package validator

import (
	"applib/app"
	"applib/funcs"
	"fmt"
	"testing"
)

type Entity struct {
	objectValidator IObjectValidator

	ID     string
	Name   string
	Amount float64
	Count  int
}

const ENTITY = "Entity"

func (ent *Entity) InitValidator() {
	entValidator := EntityValidator{}
	entValidator.Init(ENTITY)
	entValidator.SetValidatorObserve(ent)
	ent.objectValidator = &entValidator
}

func (ent Entity) GetValue(name string) (interface{}, error) {
	switch name {
	case "ID":
		return ent.ID, nil
	case "Name":
		return ent.Name, nil
	case "Amount":
		return ent.Amount, nil
	case "Count":
		return ent.Count, nil
	default:
		return nil, funcs.CannotFindProperty(name, "ObjectValidator")
	}
}

func (ent Entity) ValidateID() app.ErrorCollection {
	fields := []string{"ID"}
	return ent.objectValidator.Validate(fields)
}

func (ent Entity) ValidateBase() app.ErrorCollection {
	fields := []string{"Name", "Amount", "Count"}
	return ent.objectValidator.Validate(fields)
}

type EntityValidator struct {
	ObjectValidator
}

func (validator *EntityValidator) Init(entityName string) {

	validator.AddRequiredValidator(entityName, "ID")
	validator.AddMinLengthValidator(entityName, "ID", 4)
	validator.AddMaxLengthValidator(entityName, "ID", 5)

	validator.AddRequiredValidator(entityName, "Name")
	validator.AddMinLengthValidator(entityName, "Name", 4)
	validator.AddMaxLengthValidator(entityName, "Name", 5)

	validator.AddFl64GtzValidator(entityName, "Amount")
	validator.AddIntGtezValidator(entityName, "Count")

}

func TestRequiredValidator_Execute(t *testing.T) {

	appErr := app.ErrorCollection{}
	ent := Entity{}
	ent.Count = -1
	ent.InitValidator()

	appErr.AddCollection(ent.ValidateBase())
	appErr.AddCollection(ent.ValidateID())

	for _, err := range appErr.Errors {
		fmt.Println(err)
	}
}
