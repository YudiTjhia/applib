package svcfactory

import (
	"applib/constant"
	"applib/isvc"
	"applib/svc"
)

func CreateLogBodySvc(svcLayerType string) isvc.ILogBodySvc {
	var iSvc isvc.ILogBodySvc
	if svcLayerType == constant.SERVICE_LAYER_DEFAULT {
		iSvc = svc.LogBodySvc{}
	} else {
		panic("Unknown svcLayerType=" + svcLayerType)
	}
	return iSvc
}
