package svcfactory

import (
	"applib/constant"
	"applib/isvc"
	"applib/svc"
)

func CreateLogResponseSvc(svcLayerType string) isvc.ILogResponseSvc {

	var iSvc isvc.ILogResponseSvc
	if svcLayerType == constant.SERVICE_LAYER_DEFAULT {
		iSvc = svc.LogResponseSvc{}

	} else {
		panic("Unknown svcLayerType=" + svcLayerType)
	}

	return iSvc

}
