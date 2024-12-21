package utils

import (
	"reflect"
	"strings"
)

func DtoToModel(src interface{}, dest interface{}) error {
	srcVal := reflect.ValueOf(src).Elem()
	destVal := reflect.ValueOf(dest).Elem()

	if srcVal.Kind() != reflect.Struct || destVal.Kind() != reflect.Struct {
		return nil // или верните ошибку, если нужно
	}

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Type().Field(i)
		destField := destVal.FieldByName(srcField.Name)

		destFieldName := strings.ToLower(srcField.Name)
		found := false

		for j := 0; j < destVal.NumField(); j++ {
			destStructField := destVal.Type().Field(j)
			if strings.ToLower(destStructField.Name) == destFieldName {
				destField = destVal.Field(j)
				found = true
				break
			}
		}

		// Если поле найдено и можно установить значение
		if found && destField.IsValid() && destField.CanSet() {
			destField.Set(srcVal.Field(i))
		}
	}

	return nil
}
