package util

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func NowIn(locationName string) (time.Time, error) {
	loc, err := time.LoadLocation(locationName)
	if err == nil {
		return time.Now().In(loc), nil
	}
	return time.Time{}, err
}

func NowInJakarta() (time.Time, error) {
	return NowIn("Asia/Jakarta")
}

func ParseDate(dateStr string) (time.Time, error) {
	tDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return tDate, nil
}

func ParseDateNoError(dateStr string) time.Time {
	tDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}
	}
	return tDate
}

func ToUTC(dateStr string) (time.Time, error) {
	tDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return tDate.UTC(), nil
}

func ParseDateIn(dateStr string, locationName string) (time.Time, error) {
	loc, err := time.LoadLocation(locationName)
	if err != nil {
		return time.Time{}, err
	}

	tDate, err := time.ParseInLocation("2006-01-02", dateStr, loc)
	if err != nil {
		return time.Time{}, err
	}

	return tDate, nil

}

func ParseTime(dateStr string) (time.Time, error) {
	tDate, err := time.Parse("15:04:05", dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return tDate, nil
}

func ParseTimeIn(timeStr string, locationName string) (time.Time, error) {
	loc, err := time.LoadLocation(locationName)
	if err != nil {
		return time.Time{}, err
	}

	tDate, err := time.ParseInLocation("15:04:05", timeStr, loc)
	if err != nil {
		return time.Time{}, err
	}

	return tDate, nil
}

func ParseDateTimeIn(dateTimeStr string, locationName string) (time.Time, error) {
	loc, err := time.LoadLocation(locationName)
	if err != nil {
		return time.Time{}, err
	}

	tDate, err := time.ParseInLocation("2006-01-02 15:04:05", dateTimeStr, loc)
	if err != nil {
		return time.Time{}, err
	}

	return tDate, nil
}


func ParseDateTimeIn2(dateTimeStr string, locationName string) (time.Time, error) {
	loc, err := time.LoadLocation(locationName)
	if err != nil {
		return time.Time{}, err
	}

	tDate, err := time.ParseInLocation("2006-01-02 15:04", dateTimeStr, loc)
	if err != nil {
		return time.Time{}, err
	}

	return tDate, nil
}

func ParseDateInJakarta(dateStr string) (time.Time, error) {
	return ParseDateIn(dateStr, "Asia/Jakarta")
}

func ParseTimeInJakarta(dateStr string) (time.Time, error) {
	return ParseTimeIn(dateStr, "Asia/Jakarta")
}


func StartDateGtEndDate(startDate time.Time, endDate time.Time) error {
	if endDate.Sub(startDate) < 0 {
		return errors.New("End Date must be greater than Start Date")
	}
	return nil

}

func DateStr(dt time.Time) string {
	return dt.Format("2006-01-02")
}

func DateTimeStr(dt time.Time) string {
	return dt.Format("2006-01-02 15:04:05")
}

func TimeStr(dt time.Time) string {
	return dt.Format("15:04:05")
}

func Required(val string, errorMsg string) error {
	if val == "" {
		return errors.New(errorMsg)
	}

	return nil
}

func IsBool(val int8, errorMsg string) error {
	if !(val == 0 || val == 1) {
		return errors.New(errorMsg)
	}
	return nil
}

func IsLteFloat64(val float64, compare float64, errorMsg string) error {
	if val <= compare {
		return errors.New(errorMsg)
	}
	return nil
}

func IsLtFloat64(val float64, compare float64, errorMsg string) error {
	if val < compare {
		return errors.New(errorMsg)
	}
	return nil
}

func IsLteInt(val int, compare int, errorMsg string) error {
	if val <= compare {
		return errors.New(errorMsg)
	}
	return nil
}

func IsLteInt64(val int64, compare int64, errorMsg string) error {
	if val <= compare {
		return errors.New(errorMsg)
	}
	return nil
}

func ParseCode(code string) string {

	cols := strings.Split(code, " ")
	str := ""
	for _, col := range cols {
		col = strings.TrimSpace(col)
		str += strings.ToLower(col) + "-"
	}
	str = str[0 : len(str)-1]
	return str
}

func StrToInt(str string) int {
	x, err := strconv.Atoi(str)
	if err == nil {
		return x
	}
	return 0
}

func StrToFloat64(str string) float64 {
	x, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return x
	}
	return 0
}

func ToJson(data interface{}) string {
	b, e:=json.Marshal(data)
	if e==nil {
		return string(b)
	} else {
		return string(e.Error())
	}
}

func DecodeQuery(r *http.Request, key string) string {
	val := r.URL.Query()[key]
	if val != nil {
		return val[0]
	}
	return ""
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func DirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return false, err
}