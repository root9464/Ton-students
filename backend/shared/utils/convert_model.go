package utils

import (
	"errors"
	"reflect"
)

// ConvertDtoToEntity преобразует DTO в Entity с учетом совпадающих полей.
func ConvertDtoToEntity[T any](dto interface{}) (*T, error) {
	result := new(T)

	dtoVal := reflect.ValueOf(dto)
	if dtoVal.Kind() != reflect.Ptr || dtoVal.Elem().Kind() != reflect.Struct {
		return nil, errors.New("dto must be a pointer to a struct")
	}

	resultVal := reflect.ValueOf(result).Elem()
	dtoElem := dtoVal.Elem()

	for i := 0; i < resultVal.NumField(); i++ {
		resultField := resultVal.Type().Field(i)
		resultFieldValue := resultVal.Field(i)

		if !resultFieldValue.CanSet() {
			continue
		}

		dtoFieldValue := dtoElem.FieldByName(resultField.Name)
		if dtoFieldValue.IsValid() && dtoFieldValue.Type().AssignableTo(resultFieldValue.Type()) {
			resultFieldValue.Set(dtoFieldValue)
		}
	}

	return result, nil
}
