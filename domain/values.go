package domain

import (
	aedgrpc "github.com/sologenic/com-fs-aed-model"
)

type FieldValue interface {
	~string | ~float64 | ~int64
}

func ParseFieldValue[T FieldValue](aed *aedgrpc.AED, field aedgrpc.Field) T {
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
func CreateFieldValue[T FieldValue](field aedgrpc.Field, value T) *aedgrpc.Value {
	vObj := &aedgrpc.Value{Field: field}
	switch v := any(value).(type) {
	case float64:
		vObj.Float64Val = &v
	case int64:
		vObj.Int64Val = &v
	case string:
		vObj.StringVal = &v
	}
	return vObj
}
