package types

import "reflect"

func IsString(s interface{}) bool {
	return reflect.TypeOf(s).Kind() == reflect.String
}

func IsInt(s interface{}) bool {
	return reflect.TypeOf(s).Kind() == reflect.Int
}

func IsFloat(s interface{}) bool {
	return reflect.TypeOf(s).Kind() == reflect.Float64
}

func IsBool(s interface{}) bool {
	return reflect.TypeOf(s).Kind() == reflect.Bool
}

func IsSlice(s interface{}) bool {
	return reflect.TypeOf(s).Kind() == reflect.Slice
}

func IsMap(s interface{}) bool {
	return reflect.TypeOf(s).Kind() == reflect.Map
}

func IsStruct(s interface{}) bool {
	return reflect.TypeOf(s).Kind() == reflect.Struct
}

func IsPointer(s interface{}) bool {
	return reflect.TypeOf(s).Kind() == reflect.Ptr
}

func IsTime(s interface{}) bool {
	return reflect.TypeOf(s).Kind() == reflect.Struct && reflect.TypeOf(s).String() == "time.Time"
}