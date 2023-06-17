package max

import (
	"reflect"
)

func Max(values interface{}) interface{} {
	var _max int
	if reflect.TypeOf(values).Kind() == reflect.Slice && reflect.TypeOf(values).Elem().Kind() == reflect.Int {
		slice := reflect.ValueOf(values)
		_max = slice.Index(0).Interface().(int)
		for i := 1; i < slice.Len(); i++ {
			value := slice.Index(i).Interface().(int)
			if _max < value {
				_max = value
			}
		}
	} else if reflect.TypeOf(values).Kind() == reflect.Slice && reflect.TypeOf(values).Elem().Kind() == reflect.String {
		slice := reflect.ValueOf(values)
		sMax := slice.Index(0).Interface().(string)
		for i := 1; i < slice.Len(); i++ {
			value := slice.Index(i).Interface().(string)
			if sMax < value {
				sMax = value
			}
		}
		return sMax
	}
	return _max
}
