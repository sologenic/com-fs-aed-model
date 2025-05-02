package domain

import (
	aedgrpc "github.com/sologenic/com-fs-aed-model"
)

func ParseFieldValue[T any](aed *aedgrpc.AED, field aedgrpc.Field) T {
	var t T
	for _, v := range aed.Value {
		if v.Field == field {
			switch any(t).(type) {
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
	return t
}

// No other type check required as `CreateFieldValue` is to be used in line with the `ParseFieldValue`.
func CreateFieldValue[T any](field aedgrpc.Field, value T) *aedgrpc.Value {
	vObj := &aedgrpc.Value{Field: field}
	switch v := any(value).(type) {
	case float64:
		floatVal := v
		vObj.Float64Val = &floatVal
	case int64:
		intVal := v
		vObj.Int64Val = &intVal
	case string:
		strVal := v
		vObj.StringVal = &strVal
	}
	return vObj
}
