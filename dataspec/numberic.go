package dataspec

import (
	"fmt"
	"math"
	"reflect"
)

// NumericDataSpec 数组数据类型，用于浮点数使用
//
// 使用方式:
//
//	{
//		"name": "temperature",
//		"description": "温度",
//		"required": true,
//		"data": {
//			"type": "number",
//			"specs": {
//				"min": -275.0,
//				"max": 1000.0,
//				"step": 0.1,
//				"unit": "°"
//			}
//		}
//	}
type NumericDataSpec struct {
	// Min 最小值，若不设置，使用double最小值
	Min float64 `json:"min"`

	// Max 最大值，若不设置，使用double最大值
	Max float64 `json:"max"`

	// Step 步进, 若为零，则不使用
	Step float64 `json:"step"`

	// Unit 单位
	Unit string `json:"unit"`

	// Precision 精度
	Precision float64 `json:"precision"`
}

func (n *NumericDataSpec) Validate(v interface{}) (bool, error) {
	var result float64
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result = float64(value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result = float64(value.Uint())
	case reflect.Float32, reflect.Float64:
		result = value.Float()
	default:
		return false, fmt.Errorf("NumericDataSpecs: value type is not supported")
	}
	return n.ValidateNumber(result)
}

func (n *NumericDataSpec) ValidateNumber(v float64) (bool, error) {
	if v < n.Min || v > n.Max {
		return false, fmt.Errorf("NumericDataSpecs: value must be range [%f, %f]", n.Min, n.Max)
	}

	step := n.Step
	if math.Abs(step) > n.Precision {
		dv := v - n.Min
		s := math.Mod(dv, step)
		if math.Abs(s-step) > n.Precision && s > n.Precision {
			return false, fmt.Errorf("IntegerDataSpecs: value must be step by [%f]", step)
		}
	}
	return true, nil
}
