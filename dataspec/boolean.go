package dataspec

import "fmt"

// BooleanDataSpec 布尔数据类型
//
// 使用方式:
//
// {
// 	"name": "lighting",
// 	"description": "开关灯",
// 	"required": true,
// 	"data": {
// 		"type": "boolean",
// 		"specs": {
//			"true_desc": "开",
//			"false_desc": "关"
// 		}
// 	}
// }
type BooleanDataSpec struct {
	// TrueDesc 为真时的描述
	TrueDesc string `json:"true_desc"`

	// FalseDesc 为假时的描述
	FalseDesc string `json:"false_desc"`
}

func (n *BooleanDataSpec) Validate(v interface{}) (bool, error) {
	_, ok := v.(bool)
	if !ok {
		return false, fmt.Errorf("BooleanDataSpecs: value type is not supported")
	}

	return true, nil
}
