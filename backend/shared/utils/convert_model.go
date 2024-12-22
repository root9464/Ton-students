package utils

import (
	"fmt"
	"reflect"
	"strings"
)

// MapStruct is a generic function to map values from one struct to another.
func MapStruct[T any](src interface{}) (T, error) {
	var dest T
	// Ensure src is a pointer before using reflect.ValueOf
	srcVal := reflect.ValueOf(src)
	if srcVal.Kind() != reflect.Ptr {
		return dest, fmt.Errorf("source must be a pointer")
	}
	srcVal = srcVal.Elem() // Dereference the pointer to get the struct

	destVal := reflect.ValueOf(&dest).Elem()

	if srcVal.Kind() != reflect.Struct || destVal.Kind() != reflect.Struct {
		return dest, fmt.Errorf("source and destination must be structs")
	}

	// Iterate over the fields of the source struct
	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Type().Field(i)
		destField := destVal.FieldByName(srcField.Name)

		// Check if field name matches case-insensitively
		destFieldName := strings.ToLower(srcField.Name)
		found := false

		fmt.Printf("Src field: %s (value: %v)\n", srcField.Name, srcVal.Field(i)) // Debug: print source field name and value

		for j := 0; j < destVal.NumField(); j++ {
			destStructField := destVal.Type().Field(j)
			fmt.Printf("Dest field: %s\n", destStructField.Name) // Debug: print destination field name

			if strings.ToLower(destStructField.Name) == destFieldName {
				destField = destVal.Field(j)
				found = true
				break
			}
		}

		// If the field was found and is valid, copy the value
		if found && destField.IsValid() && destField.CanSet() {
			if destField.Kind() == reflect.Struct {
				// Recursive call for nested structs
				nestedValue, err := MapStruct[T](srcVal.Field(i).Addr().Interface())
				if err != nil {
					return dest, err
				}
				destField.Set(reflect.ValueOf(nestedValue))
			} else {
				// Set the field value directly
				destField.Set(srcVal.Field(i))
			}
		}
	}

	return dest, nil
}
