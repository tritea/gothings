package property

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/AtomPod/thingmodel/thingmodel/dataspec"
)

// PropertyDescription 属性描述，用于描述某个属性使用，作为物模型中设备或者传感器所拥有的属性
type PropertyDescription struct {
	// Name 属性名称(作为key使用)
	Name string `json:"name"`

	// Description 属性描述，作为解释说明
	Description string `json:"description"`

	// Required 是否必须存在
	Required bool `json:"required"`

	// AccessMode 属性的访问支持，其中w代表可写、r代表可读，该属性仅支持这两个字符
	AccessMode string `json:"access_mode"`

	// Data 属性对应的数据描述
	Data *dataspec.DataDescription `json:"data"`
}

func (p *PropertyDescription) isAccessMode(c byte) bool {
	return c == 'w' || c == 'r'
}

// Readable 是否可读
func (p *PropertyDescription) Readable() bool {
	return strings.Contains(p.AccessMode, "r")
}

// Writable 是否可写
func (p *PropertyDescription) Writable() bool {
	return strings.Contains(p.AccessMode, "w")
}

func (p *PropertyDescription) UpdateData() error {
	if len(p.Name) == 0 {
		return fmt.Errorf("PropertyDescription: name could not be empty")
	}

	l := len(p.AccessMode)
	if l > 2 || (l > 0 && !p.isAccessMode(p.AccessMode[0])) || (l > 1 && !p.isAccessMode(p.AccessMode[1])) {
		return fmt.Errorf("PropertyDescription: access mode is invalid")
	}

	if p.Data == nil {
		return fmt.Errorf("PropertyDescription: data field could not be empty")
	}

	if err := p.Data.Parse(); err != nil {
		return err
	}
	return nil
}

func (p *PropertyDescription) Parse(b []byte) error {
	if err := json.Unmarshal([]byte(b), p); err != nil {
		return err
	}
	return p.UpdateData()
}

// Validate 验证数据是否正确
func (p *PropertyDescription) Validate(v interface{}) (bool, error) {
	return p.Data.Validate(v)
}
