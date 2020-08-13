package validator

type LogBodyValidator struct {
	ObjectValidator
}

func (validator *LogBodyValidator) Init(entityName string) {
	validator.AddRequiredValidator(entityName, "ID")
	validator.AddMinLengthValidator(entityName, "ID", 4)
	validator.AddMaxLengthValidator(entityName, "ID", 40)
	validator.AddRequiredValidator(entityName, "SystemID")
	validator.AddMinLengthValidator(entityName, "SystemID", 4)
	validator.AddMaxLengthValidator(entityName, "SystemID", 40)
	validator.AddRequiredValidator(entityName, "LogID")
	validator.AddMinLengthValidator(entityName, "LogID", 4)
	validator.AddMaxLengthValidator(entityName, "LogID", 40)
	validator.AddRequiredValidator(entityName, "LogType")
	validator.AddMinLengthValidator(entityName, "LogType", 4)
	validator.AddMaxLengthValidator(entityName, "LogType", 10)
	validator.AddRequiredValidator(entityName, "LogTag")
	validator.AddMinLengthValidator(entityName, "LogTag", 4)
	validator.AddMaxLengthValidator(entityName, "LogTag", 20)
	validator.AddRequiredValidator(entityName, "Message")
	validator.AddRequiredValidator(entityName, "LogDate")

}
