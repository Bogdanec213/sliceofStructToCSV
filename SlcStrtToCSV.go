package SlcStrtToCSV

import (
	"bytes"
	"encoding/csv"
	"github.com/spf13/cast"
	"reflect"
	"time"
)

func GetCSV(sliceOfTruct interface{}) (*bytes.Buffer, string) {
	data := reflect.ValueOf(sliceOfTruct)
	if data.Kind() != reflect.Slice {
		return &bytes.Buffer{}, "sliceofStructToCSV error: given a non-slice type"
	}

	sliceOfInterface := make([]interface{}, data.Len())

	for i := 0; i < data.Len(); i++ {
		sliceOfInterface[i] = data.Index(i).Interface()
	}
	if len(sliceOfInterface) > 0 {
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
			return &bytes.Buffer{}, "SlcStrtToCSV error: " + err.Error()
		}
		for _, value := range sliceOfInterface {
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
				return &bytes.Buffer{}, "SlcStrtToCSV error: " + err.Error()
			}
		}
		writer.Flush()
		return b, ""
	}
	return &bytes.Buffer{}, "SlcStrtToCSV error: slice is empty."
}
