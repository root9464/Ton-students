package utils

import (
	"errors"
	"reflect"
)

func ConvertDtoToEntity[T any](dto interface{}) (*T, error) {
	result := new(T)

	dtoVal := reflect.ValueOf(dto)
	if dtoVal.Kind() != reflect.Ptr {
		return nil, errors.New("dto must be a pointer to a struct")
	}

	dtoElem := dtoVal.Elem()
	if dtoElem.Kind() != reflect.Struct {
		return nil, errors.New("dto must be a struct")
	}

	resultVal := reflect.ValueOf(result).Elem()

	for i := 0; i < dtoElem.NumField(); i++ {
		dtoField := dtoElem.Field(i)
		resultField := resultVal.Field(i)

		if dtoField.Type().AssignableTo(resultField.Type()) {
			resultField.Set(dtoField)
		}
	}

	return result, nil
}
