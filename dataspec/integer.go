package dataspec

import (
	"fmt"
	"reflect"
)

// IntegerDataSpec 整数数据类型，包括有符号和无符号，正常应该使用有符合，因为数据范围更通用
//
// 使用方式:
//
//	{
//		"name": "age",
//		"description": "年龄",
//		"required": true,
//		"data": {
//			"type": "integer",
//			"specs": {
//				"min": 1,
//				"max": 100,
//				"step": 5,
//				"unit": "y"
//			}
//		}
//	}
type IntegerDataSpec struct {
	// Min 最小值，若不设置则会取Int64最小值
	Min int64 `json:"min"`

	// Max 最大值，若不设置则会取Int64最大值
	Max int64 `json:"max"`

	// Step 步进，单步进为零时，不使用
	Step int64 `json:"step"`

	// Unit 单位
	Unit string `json:"unit"`
}

func (n *IntegerDataSpec) Validate(v interface{}) (bool, error) {
	var result int64
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result = value.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result = int64(value.Uint())
	default:
		return false, fmt.Errorf("value type is not supported")
	}
	return n.ValidateInteger(result)
}

func (n *IntegerDataSpec) ValidateInteger(v int64) (bool, error) {
	if v < n.Min || v > n.Max {
		return false, fmt.Errorf("IntegerDataSpecs: value must be range [%d, %d]", n.Min, n.Max)
	}

	step := n.Step
	if step != 0 {
		dv := v - n.Min
		if dv%step != 0 {
			return false, fmt.Errorf("IntegerDataSpecs: value must be step by [%d]", step)
		}
	}
	return true, nil
}
