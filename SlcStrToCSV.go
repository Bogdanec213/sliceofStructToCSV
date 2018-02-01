package SlcStrToCSV

import (
	"bytes"
	"encoding/csv"
	"github.com/spf13/cast"
	"reflect"
	"time"
)

func GetCSV(data []interface{}) (*bytes.Buffer, string) {
	if len(data) > 0 {
		b := &bytes.Buffer{}
		writer := csv.NewWriter(b)
		titleSlice := []string{}
		val := reflect.Indirect(reflect.ValueOf(data[0]))
		for i := 0; i < val.NumField(); i++ {
			if val.Type().Field(i).Tag.Get("title") == "" {
				titleSlice = append(titleSlice, val.Type().Field(i).Name)
			} else {
				titleSlice = append(titleSlice, val.Type().Field(i).Tag.Get("title"))
			}
		}
		if err := writer.Write(titleSlice); err != nil {
			return &bytes.Buffer{}, "sliceofStructToCSV error: " + err.Error()
		}
		for _, value := range data {
			val := reflect.Indirect(reflect.ValueOf(value))
			record := []string{}
			for i := 0; i < val.NumField(); i++ {
				if val.Field(i).Type().String() != "time.Time" {
					record = append(record, cast.ToString(val.Field(i).Interface()))
				} else {
					record = append(record, val.Field(i).Interface().(time.Time).Format("2006-01-02 15:04:05"))
				}
			}
			if err := writer.Write(record); err != nil {
				return &bytes.Buffer{}, "sliceofStructToCSV error: " + err.Error()
			}
		}
		writer.Flush()
		return b, ""
	}
	return &bytes.Buffer{}, "sliceofStructToCSV error: slice is empty."
}
