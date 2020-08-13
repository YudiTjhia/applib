package validator

type LogResponseValidator struct {
	ObjectValidator
}

func (validator *LogResponseValidator) Init(entityName string) {
	validator.AddRequiredValidator(entityName, "ID")
	validator.AddMinLengthValidator(entityName, "ID", 4)
	validator.AddMaxLengthValidator(entityName, "ID", 40)
	validator.AddRequiredValidator(entityName, "SystemID")
	validator.AddMinLengthValidator(entityName, "SystemID", 4)
	validator.AddMaxLengthValidator(entityName, "SystemID", 40)
	validator.AddRequiredValidator(entityName, "AppID")
	validator.AddMinLengthValidator(entityName, "AppID", 4)
	validator.AddMaxLengthValidator(entityName, "AppID", 40)
	validator.AddRequiredValidator(entityName, "ServiceName")
	validator.AddMinLengthValidator(entityName, "ServiceName", 4)
	validator.AddMaxLengthValidator(entityName, "ServiceName", 255)

}
