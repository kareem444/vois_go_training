package modify

import "reflect"

func Map[T any](data T, modifyObjects map[string]any) T {
	dataValue := reflect.ValueOf(&data).Elem()

	for key, value := range modifyObjects {
		field := dataValue.FieldByName(key)

		if field.IsValid() && field.CanSet() {
			field.Set(reflect.ValueOf(value))
		}
	}

	return data
}

func SliceMap[T any](data []T, modifyObjects map[string]any) []T {
	for i := range data {
		data[i] = Map(data[i], modifyObjects)
	}

	return data
}