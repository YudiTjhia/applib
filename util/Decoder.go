package util

import (
	"net/http"
	"strconv"
)

func  DecodeQueryKey(r *http.Request, key string) string {
	val := r.URL.Query()[key]
	if val != nil {
		return val[0]
	}
	return ""
}

func  DecodeQueryInt(r *http.Request, field string) int{
	val := DecodeQueryKey(r, field)
	if val!="" {
		intVal, _ := strconv.Atoi(val)
		return intVal
	}
	return 0
}

func  DecodeQueryFloat32(r *http.Request, field string) float32 {
	val := DecodeQueryKey(r, field)
	if val!="" {
		intVal, _ := strconv.ParseFloat(val,32)
		return float32(intVal)
	}
	return float32(0)
}

func  DecodeQueryFloat64(r *http.Request, field string) float64 {
	val := DecodeQueryKey(r, field)
	if val!="" {
		intVal, _ := strconv.ParseFloat(val,64)
		return float64(intVal)
	}
	return float64(0)
}


