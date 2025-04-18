package reflection

import (
	"fmt"
	"reflect"
	"strings"
)

func ReflectionLowerCaseFieldNames(v interface{}) (map[string]interface{}, error) {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("非结构体类型")
	}

	result := make(map[string]interface{})
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		key := strings.ToLower(field.Name) // 字段名转为小写
		result[key] = val.Field(i).Interface()
	}
	return result, nil
}
