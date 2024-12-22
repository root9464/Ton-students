package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func MapStruct[T any](src interface{}) (*T, error) {
	dest := new(T)

	srcVal := reflect.ValueOf(src)
	if srcVal.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("source must be a pointer")
	}
	srcVal = srcVal.Elem()
	destVal := reflect.ValueOf(dest).Elem()

	if srcVal.Kind() != reflect.Struct || destVal.Kind() != reflect.Struct {
		return nil, fmt.Errorf("source and destination must be structs")
	}

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Type().Field(i)
		destField := destVal.FieldByName(srcField.Name)
		destFieldName := strings.ToLower(srcField.Name)
		found := false

		fmt.Printf("Src field: %s (value: %v)\n", srcField.Name, srcVal.Field(i))

		for j := 0; j < destVal.NumField(); j++ {
			destStructField := destVal.Type().Field(j)
			fmt.Printf("Dest field: %s\n", destStructField.Name)
			if strings.ToLower(destStructField.Name) == destFieldName {
				destField = destVal.Field(j)
				found = true
				break
			}
		}

		if found && destField.IsValid() && destField.CanSet() {
			if destField.Kind() == reflect.Struct {
				nestedValue, err := MapStruct[T](srcVal.Field(i).Addr().Interface())
				if err != nil {
					return nil, err
				}
				destField.Set(reflect.ValueOf(nestedValue).Elem())
			} else {
				destField.Set(srcVal.Field(i))
			}
		}
	}

	return dest, nil
}

// DtoToModel преобразует структуру src в структуру типа *T
func DtoToModel[T any](src interface{}, dest T) (*T, error) {
	var empty *T // Инициализация переменной типа *T с "пустым" значением (nil)

	srcVal := reflect.ValueOf(src).Elem()

	if srcVal.Kind() != reflect.Struct {
		return empty, fmt.Errorf("source must be a struct")
	}

	// Создаем новый указатель на структуру типа T
	destVal := reflect.New(reflect.TypeOf(dest)).Elem()

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Type().Field(i)
		destField := destVal.FieldByName(srcField.Name)

		destFieldName := strings.ToLower(srcField.Name)
		found := false

		// Найти соответствующее поле в destination структуре
		for j := 0; j < destVal.NumField(); j++ {
			destStructField := destVal.Type().Field(j)
			if strings.ToLower(destStructField.Name) == destFieldName {
				destField = destVal.Field(j)
				found = true
				break
			}
		}

		if found && destField.IsValid() && destField.CanSet() {
			if destField.Kind() == reflect.Struct {
				// Рекурсивно вызываем для вложенных структур
				newDestField, err := DtoToModel(srcVal.Field(i).Addr().Interface(), destField.Addr().Interface())
				if err != nil {
					return empty, err
				}
				destField.Set(reflect.ValueOf(newDestField).Elem())
			} else {
				destField.Set(srcVal.Field(i))
			}
		}
	}

	return destVal.Addr().Interface().(*T), nil
}
