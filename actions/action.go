package actions

import (
	"fmt"

	"github.com/AtomPod/thingmodel/thingmodel/dataspec"
)

// ActionDescription 动作(方法)描述，用于功能调用等
type ActionDescription struct {
	// Name 动作名称，用于调用使用
	Name string `json:"name"`

	// Description 事件描述
	Description string `json:"description"`

	// InputData 外部输入的数据描述
	InputData *dataspec.DataDescription `json:"input_data"`

	// OutputData 设备输出的数据描述
	OutputData *dataspec.DataDescription `json:"output_data"`
}

func (a *ActionDescription) UpdateData() error {
	if len(a.Name) == 0 {
		return fmt.Errorf("ActionDescription: name could not be empty")
	}

	if a.InputData == nil || a.OutputData == nil {
		return fmt.Errorf("ActionDescription: data field could not be empty")
	}

	return nil
}

// ValidateInput 验证输入的数据是否正确
func (a *ActionDescription) ValidateInput(v interface{}) (bool, error) {
	return a.InputData.Validate(v)
}

// ValidateOutput 验证输出的数据是否正确
func (a *ActionDescription) ValidateOutput(v interface{}) (bool, error) {
	return a.OutputData.Validate(v)
}
