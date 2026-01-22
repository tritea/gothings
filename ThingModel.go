package thingmodel

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/AtomPod/thingmodel/thingmodel/actions"
	"github.com/AtomPod/thingmodel/thingmodel/events"
	"github.com/AtomPod/thingmodel/thingmodel/property"
)

type ThingModel struct {
	// ID 物模型ID
	ID string `json:"id"`

	// Name 物模型名称
	Name string `json:"name"`

	// CreatedAt 创建时间
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt 更新时间
	UpdatedAt time.Time `json:"updated_at"`

	// Properties 属性列表
	Properties []property.PropertyDescription `json:"properties"`

	// Actions 动作列表
	Actions []actions.ActionDescription `json:"actions"`

	// Events 事件列表
	Events []events.EventDescription `json:"events"`
}

func (t *ThingModel) Parse(b []byte) error {
	if err := json.Unmarshal(b, t); err != nil {
		return err
	}

	if t.Properties == nil {
		t.Properties = make([]property.PropertyDescription, 0)
	}

	if t.Events == nil {
		t.Events = make([]events.EventDescription, 0)
	}

	if t.Actions == nil {
		t.Actions = make([]actions.ActionDescription, 0)
	}

	props := t.Properties
	for i := 0; i < len(props); i++ {
		if err := props[i].UpdateData(); err != nil {
			return err
		}

		if props[i].AccessMode == "" {
			props[i].AccessMode = "wr"
		}
	}

	events := t.Events
	for i := 0; i < len(events); i++ {
		if err := events[i].UpdateData(); err != nil {
			return err
		}
	}

	actions := t.Actions
	for i := 0; i < len(actions); i++ {
		if err := actions[i].UpdateData(); err != nil {
			return err
		}
	}

	return nil
}

// GetProperty 获取属性,若不存在，返回nil
func (t *ThingModel) GetProperty(name string) *property.PropertyDescription {
	for i := range t.Properties {
		if t.Properties[i].Name == name {
			return &t.Properties[i]
		}
	}
	return nil
}

// GetEvent 获取事件，若不存在，返回nil
func (t *ThingModel) GetEvent(name string) *events.EventDescription {
	for i := range t.Events {
		if t.Events[i].Name == name {
			return &t.Events[i]
		}
	}
	return nil
}

// GetAction 获取活动，若不存在，返回nil
func (t *ThingModel) GetAction(name string) *actions.ActionDescription {
	for i := range t.Actions {
		if t.Actions[i].Name == name {
			return &t.Actions[i]
		}
	}
	return nil
}

func (t *ThingModel) ValidateProperty(name string, v interface{}) (bool, error) {
	for _, p := range t.Properties {
		if p.Name == name {
			return p.Validate(v)
		}
	}
	return false, fmt.Errorf("property not found")
}

func (t *ThingModel) ValidateActionInput(name string, v interface{}) (bool, error) {
	for _, p := range t.Actions {
		if p.Name == name {
			return p.InputData.Validate(v)
		}
	}
	return false, fmt.Errorf("action not found")
}

func (t *ThingModel) ValidateActionOutput(name string, v interface{}) (bool, error) {
	for _, p := range t.Actions {
		if p.Name == name {
			return p.OutputData.Validate(v)
		}
	}
	return false, fmt.Errorf("action not found")
}

func (t *ThingModel) ValidateEvent(name string, v interface{}) (bool, error) {
	for _, e := range t.Events {
		if e.Name == name {
			return e.Data.Validate(v)
		}
	}
	return false, fmt.Errorf("event not found")
}

// ToJSON 将 ThingModel 序列化为 JSON 字节
func (t *ThingModel) ToJSON() ([]byte, error) {
	return json.Marshal(t)
}

// Validate 验证整个 ThingModel 的完整性
func (t *ThingModel) Validate() error {
	for _, p := range t.Properties {
		if p.Data == nil {
			return fmt.Errorf("property '%s': data field is nil", p.Name)
		}
	}

	for _, e := range t.Events {
		if e.Data == nil {
			return fmt.Errorf("event '%s': data field is nil", e.Name)
		}
	}

	for _, a := range t.Actions {
		if a.InputData == nil || a.OutputData == nil {
			return fmt.Errorf("action '%s': data field is nil", a.Name)
		}
	}

	return nil
}

// AddProperty 添加一个属性
func (t *ThingModel) AddProperty(prop property.PropertyDescription) error {
	if prop.Name == "" {
		return fmt.Errorf("property name cannot be empty")
	}
	if prop.Data == nil {
		return fmt.Errorf("property data field cannot be nil")
	}
	t.Properties = append(t.Properties, prop)
	return nil
}

// AddEvent 添加一个事件
func (t *ThingModel) AddEvent(evt events.EventDescription) error {
	if evt.Name == "" {
		return fmt.Errorf("event name cannot be empty")
	}
	if evt.Data == nil {
		return fmt.Errorf("event data field cannot be nil")
	}
	t.Events = append(t.Events, evt)
	return nil
}

// AddAction 添加一个动作
func (t *ThingModel) AddAction(act actions.ActionDescription) error {
	if act.Name == "" {
		return fmt.Errorf("action name cannot be empty")
	}
	if act.InputData == nil || act.OutputData == nil {
		return fmt.Errorf("action input/output data field cannot be nil")
	}
	t.Actions = append(t.Actions, act)
	return nil
}

// RemoveProperty 删除指定名称的属性，返回是否成功
func (t *ThingModel) RemoveProperty(name string) bool {
	for i, p := range t.Properties {
		if p.Name == name {
			t.Properties = append(t.Properties[:i], t.Properties[i+1:]...)
			return true
		}
	}
	return false
}

// RemoveEvent 删除指定名称的事件，返回是否成功
func (t *ThingModel) RemoveEvent(name string) bool {
	for i, e := range t.Events {
		if e.Name == name {
			t.Events = append(t.Events[:i], t.Events[i+1:]...)
			return true
		}
	}
	return false
}

// RemoveAction 删除指定名称的动作，返回是否成功
func (t *ThingModel) RemoveAction(name string) bool {
	for i, a := range t.Actions {
		if a.Name == name {
			t.Actions = append(t.Actions[:i], t.Actions[i+1:]...)
			return true
		}
	}
	return false
}

// GetReadableProperties 获取所有可读的属性
func (t *ThingModel) GetReadableProperties() []property.PropertyDescription {
	result := make([]property.PropertyDescription, 0)
	for _, p := range t.Properties {
		if p.Readable() {
			result = append(result, p)
		}
	}
	return result
}

// GetWritableProperties 获取所有可写的属性
func (t *ThingModel) GetWritableProperties() []property.PropertyDescription {
	result := make([]property.PropertyDescription, 0)
	for _, p := range t.Properties {
		if p.Writable() {
			result = append(result, p)
		}
	}
	return result
}
