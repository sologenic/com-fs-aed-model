package domain

import (
	aedgrpc "github.com/sologenic/com-fs-aed-model"
)

func ParseFieldValue[T any](aed *aedgrpc.AED, field aedgrpc.Field) T {
	var zero T
	for _, v := range aed.Value {
		if v.Field == field {
			switch any(zero).(type) {
			case float64:
				if v.Float64Val != nil {
					return any(*v.Float64Val).(T)
				}
			case int64:
				if v.Int64Val != nil {
					return any(*v.Int64Val).(T)
				}
			case string:
				if v.StringVal != nil {
					return any(*v.StringVal).(T)
				}
			}
		}
	}
	return zero
}

func CreateFieldValue[T any](field aedgrpc.Field, value T) *aedgrpc.Value {
	valueObj := &aedgrpc.Value{Field: field}

	switch v := any(value).(type) {
	case float64:
		floatVal := v
		valueObj.Float64Val = &floatVal
	case int64:
		intVal := v
		valueObj.Int64Val = &intVal
	case string:
		strVal := v
		valueObj.StringVal = &strVal
	}

	return valueObj
}
