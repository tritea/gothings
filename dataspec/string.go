package dataspec

import "fmt"

// StringDataSpec 字符串数据类型
//
// 使用方式:
//
//	{
//		"name": "model",
//		"description": "型号",
//		"required": true,
//		"data": {
//			"type": "string",
//			"specs": {
//				"length": 15
//			}
//		}
//	}
type StringDataSpec struct {
	// Length 字符串最大长度
	Length int32 `json:"length"`
}

func (n *StringDataSpec) Validate(v interface{}) (bool, error) {
	str, ok := v.(string)
	if !ok {
		return false, fmt.Errorf("StringDataSpecs: value type is not supported")
	}

	return n.ValidateString(str)
}

func (n *StringDataSpec) ValidateString(v string) (bool, error) {
	if n.Length == 0 {
		return true, nil
	}

	if len(v) > int(n.Length) {
		return false, fmt.Errorf("StringDataSpecs: string length must be range [%d, %d]", 0, n.Length)
	}
	return true, nil
}
