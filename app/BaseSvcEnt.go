package app

import "net/http"
import "strconv"

type BaseSvcEnt struct {
	WfAction string `json:"wfAction"`
	WfNotes string `json:"wfNotes"`
}

func (BaseSvcEnt) DecodeQueryKey(r *http.Request, key string) string {
	val := r.URL.Query()[key]
	if val != nil {
		return val[0]
	}
	return ""
}

func (baseSvcEnt BaseSvcEnt) DecodeQueryInt(r *http.Request, field string) int{
	val := baseSvcEnt.DecodeQueryKey(r, field)
	if val!="" {
		intVal, _ := strconv.Atoi(val)
		return intVal
	}
	return 0
}

func (baseSvcEnt BaseSvcEnt) DecodeQueryBool(r *http.Request, field string) bool {
	val := baseSvcEnt.DecodeQueryKey(r, field)
	if val!="" {
		boolVal, _ := strconv.ParseBool(val)
		return boolVal
	}
	return false
}

func (baseSvcEnt BaseSvcEnt) DecodeQueryFloat32(r *http.Request, field string) float32 {
	val := baseSvcEnt.DecodeQueryKey(r, field)
	if val!="" {
		intVal, _ := strconv.ParseFloat(val,32)
		return float32(intVal)
	}
	return float32(0)
}

func (baseSvcEnt BaseSvcEnt) DecodeQueryFloat64(r *http.Request, field string) float64 {
	val := baseSvcEnt.DecodeQueryKey(r, field)
	if val!="" {
		intVal, _ := strconv.ParseFloat(val,64)
		return float64(intVal)
	}
	return float64(0)
}

