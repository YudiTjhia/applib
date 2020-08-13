package app

type BaseSvc struct {
	App
	TableName     string
	DataLayerType string
	ConnectionID  string
	DbConfigFile  string
	AppConfigFile string
}
