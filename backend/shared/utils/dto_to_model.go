package utils

import (
	"fmt"
	"reflect"
)

func ConvertDtoToEntity[T any](src interface{}, dest T) (*T, error) {
	empty := new(T)

	srcVal := reflect.ValueOf(src)
	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	if srcVal.Kind() != reflect.Struct {
		return empty, fmt.Errorf("source must be a struct")
	}

	destVal := reflect.New(reflect.TypeOf(dest)).Elem()

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Type().Field(i)
		srcFieldValue := srcVal.Field(i)
		destField := destVal.FieldByName(srcField.Name)

		if destField.IsValid() && destField.CanSet() {
			if srcFieldValue.Type().AssignableTo(destField.Type()) {
				destField.Set(srcFieldValue)
			} else if srcFieldValue.Type().ConvertibleTo(destField.Type()) {
				destField.Set(srcFieldValue.Convert(destField.Type()))
			} else {
				return empty, fmt.Errorf("cannot assign field %s: incompatible types", srcField.Name)
			}
		}
	}

	result := destVal.Addr().Interface().(*T)
	return result, nil
}
