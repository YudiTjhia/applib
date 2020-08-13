package constant

import "applib/app"

const (
	DB_CONFIG_FILE        = "conf/db.conf"
	APP_CONFIG_FILE       = "conf/app.conf"
	CONN_ID_CONFIG_FILE   = "conf/connection-ids.conf"
	SERVICE_LAYER_DEFAULT = "DEFAULT"
)

func AppendConf(header *app.RequestHeader) {
	appConfigFile := app.AppConfigFile{}
	appConfigFile.ConnIDConfigFile = "conf/connection-ids.conf"
	appConfigFile.DBConfigFile = "conf/db.conf"
	appConfigFile.AppConfigFile = "conf/app.conf"
	header.AppConfigFile = appConfigFile
}

func GetAppConfigFile() app.AppConfigFile {
	appConfigFile := app.AppConfigFile{}
	appConfigFile.ConnIDConfigFile = "conf/connection-ids.conf"
	appConfigFile.DBConfigFile = "conf/db.conf"
	appConfigFile.AppConfigFile = "conf/app.conf"
	return appConfigFile
}
