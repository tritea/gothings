package dataspec

import (
	"fmt"
	"reflect"
)

// ArrayDataSpec 数组数据，为了限制，同一数组数据类型应该相同
//
// 使用方式:
//
//	构造一个 字符串数组，每个字符串最长15
//	{
//		"name": "hello",
//		"description": "hello world",
//		"required": true,
//		"data": {
//			"type": "array",
//			"specs": {
//				"length": 5
//				"data": {
//					"type": "string",
//					"specs": {
//						"length": 15
//					}
//				}
//			}
//		}
//	}
type ArrayDataSpec struct {
	// Length 长度
	Length int32 `json:"length"`

	// Data 数组数据类型
	Data *DataDescription `json:"data"`
}

func (a *ArrayDataSpec) Validate(v interface{}) (bool, error) {
	value := reflect.ValueOf(v)
	typ := value.Type()
	kind := typ.Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return false, fmt.Errorf("ArrayDataSpecs: value type is not supported")
	}

	len := value.Len()
	if a.Length != int32(len) {
		return false, fmt.Errorf("ArrayDataSpecs: array size too large or too small")
	}

	for i := 0; i < len; i++ {
		elemVal := value.Index(i)
		if ok, err := validateReflectData(a.Data, elemVal); !ok {
			return ok, err
		}
	}
	return true, nil
}
