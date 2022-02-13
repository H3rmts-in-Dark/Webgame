package model

import (
	"fmt"
	"reflect"
)

// Decode adds custom conversions to mapstructure decoder
// Array t o String, as timeUUID's are used as ids and wont get converted
// to string by default
func Decode(type1 reflect.Kind, type2 reflect.Kind, int interface{}) (interface{}, error) {
	if type1 == reflect.Array && type2 == reflect.String {
		return fmt.Sprintf("%s", int), nil
	}

	return int, nil
}
