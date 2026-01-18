package dataspec

import (
	"fmt"
	"reflect"
)

// StructDataSpec 结构体数据类型，该类型将spec当成一个子data进行处理，类似下列使用，specs每一个字段都为data
//
// 使用方式:
//
//	{
//		"name": "hello",
//		"description": "hello world",
//		"required": true,
//		"data": {
//			"type": "struct",
//			"specs": {
//				"name": {
//					"type": "string",
//					"specs": {
//						"length": 15
//					}
//				},
//				"age": {
//					"type": "integer",
//					"specs": {
//						"min": 0,
//						"max": 15,
//						"step": 1,
//						"unit": "y"
//					}
//				}
//			}
//		}
//	}
type StructDataSpec map[string]*DataDescription

func (a StructDataSpec) Validate(v interface{}) (bool, error) {
	value := reflect.ValueOf(v)
	typ := value.Type()
	kind := value.Kind()

	if kind == reflect.Pointer {
		value = value.Elem()
		typ = value.Type()
		kind = value.Kind()
	}

	if kind != reflect.Map && kind != reflect.Struct {
		return false, fmt.Errorf("StructDataSpecs: value type is not supported")
	}

	if kind == reflect.Map {
		iter := value.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()

			if k.Kind() != reflect.String {
				return false, fmt.Errorf("StructDataSpecs: type of key must be string")
			}

			kind := v.Kind()
			if kind == reflect.Interface || kind == reflect.Pointer {
				v = v.Elem()
				if !v.IsValid() {
					return false, fmt.Errorf("StructDataSpecs: field is invalid or nil")
				}
			}

			key := k.String()
			dd, ok := a[key]
			if !ok {
				return false, fmt.Errorf("StructDataSpecs: field [%s] is not allowed", key)
			}

			if ok, err := validateReflectData(dd, v); !ok {
				return ok, err
			}
		}
		return true, nil
	}

	if kind == reflect.Struct {
		numOfFields := value.NumField()
		for i := 0; i < numOfFields; i++ {
			value := value.Field(i)
			field := typ.Field(i)

			key := ""
			if k, ok := field.Tag.Lookup("json"); ok {
				key = k
			} else {
				key = field.Name
			}

			kind := value.Kind()
			if kind == reflect.Interface || kind == reflect.Pointer {
				value = value.Elem()

				if !value.IsValid() {
					return false, fmt.Errorf("StructDataSpecs: field is invalid or nil")
				}
			}

			dd, ok := a[key]
			if !ok {
				return false, fmt.Errorf("StructDataSpecs: field [%s] is not allowed", key)
			}

			if ok, err := validateReflectData(dd, value); !ok {
				return ok, err
			}
		}
		return true, nil
	}
	return false, fmt.Errorf("StructDataSpecs: argument type is not supported")
}
