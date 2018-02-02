package SlcStrtToCSV

import (
	"bytes"
	"encoding/csv"
	"github.com/spf13/cast"
	"reflect"
	"time"
	"errors"
)


func GetCSV(sliceOfTruct interface{}) (*bytes.Buffer, error) {
	data := reflect.ValueOf(sliceOfTruct)
	if data.Kind() != reflect.Slice {
		return &bytes.Buffer{}, errors.New("sliceofStructToCSV error: given a non-slice type")
	}

	sliceOfInterface := make([]interface{}, data.Len())

	for i := 0; i < data.Len(); i++ {
		sliceOfInterface[i] = data.Index(i).Interface()
	}
	if len(sliceOfInterface) > 0 {
		b := &bytes.Buffer{}
		writer := csv.NewWriter(b)
		titleSlice := []string{}
		fieldNameSlice:= []string{}
		val := reflect.Indirect(reflect.ValueOf(sliceOfInterface[0]))
		skip := map[int]bool{}
		for i := 0; i < val.NumField(); i++ {
			fieldNameSlice= append(fieldNameSlice, val.Type().Field(i).Name)
			if fieldNameSlice[i] != "-" {
				if fieldNameSlice[i] == "" {
					titleSlice = append(titleSlice, val.Type().Field(i).Name)
				} else {
				titleSlice = append(titleSlice, fieldNameSlice[i])
				}
			} else{
				skip[i] = true
			}
		}
		if err := writer.Write(titleSlice); err != nil {
			return &bytes.Buffer{}, errors.New("SlcStrtToCSV error: " + err.Error())
		}
		for _, value := range sliceOfInterface {
			val := reflect.Indirect(reflect.ValueOf(value))
			record := []string{}
			for i := 0; i < val.NumField(); i++ {
				if _, ok := skip[i]; ok {
					continue
				}
				if val.Field(i).Type().String() != "time.Time" {
					record = append(record, cast.ToString(val.Field(i).Interface()))
				} else {
					record = append(record, val.Field(i).Interface().(time.Time).Format("2006-01-02 15:04:05"))
				}
			}
			if err := writer.Write(record); err != nil {
				return &bytes.Buffer{}, errors.New("SlcStrtToCSV error: " + err.Error())
			}
		}
		writer.Flush()
		return b, nil
	}
	return &bytes.Buffer{}, errors.New("SlcStrtToCSV error: slice is empty.")
}
