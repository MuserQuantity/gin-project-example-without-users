package utils

import (
	"encoding/json"
)

//func Struct2Map(obj interface{}) map[string]interface{} {
//	t := reflect.TypeOf(obj)
//	v := reflect.ValueOf(obj)
//
//	var data = make(map[string]interface{})
//	for i := 0; i < t.NumField(); i++ {
//		data[t.Field(i).Name] = v.Field(i).Interface()
//	}
//	return data
//}

func Struct2JsonStr(obj interface{}) string {
	byteData, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(byteData)
}

func JsonStr2Map(jsonStr string) map[string]interface{} {
	var mapResult map[string]interface{}
	_ = json.Unmarshal([]byte(jsonStr), &mapResult)
	return mapResult
}
