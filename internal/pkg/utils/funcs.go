package utils

import "reflect"

const fieldTagName = "stm"

// StrToMap Convert struct to map[string]int64 based on stm tag
func StrToMap(s interface{}) map[string]int64 {
	rv := make(map[string]int64)
	t, v := reflect.TypeOf(s).Elem(), reflect.ValueOf(s).Elem()
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)
		if tag := ft.Tag.Get(fieldTagName); tag != "" {
			rv[tag] = v.Field(i).Int()
		}
	}
	return rv
}
