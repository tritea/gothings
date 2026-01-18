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
	if err := json.Unmarshal([]byte(b), t); err != nil {
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
	for  i := 0; i < len(actions); i++ {
		if err := actions[i].UpdateData(); err != nil {
			return err
		}
	}

	return nil
}

// GetProperty 获取属性,若不存在，返回nil
func (t *ThingModel) GetProperty(name string) *property.PropertyDescription {
	for _, p := range t.Properties {
		if p.Name == name {
			return &p
		}
	}
	return nil
}

// GetEvent 获取事件，若不存在，返回nil
func (t *ThingModel) GetEvent(name string) *events.EventDescription {
	for _, e := range t.Events {
		if e.Name == name {
			return &e
		}
	}
	return nil
}

// GetAction 获取活动，若不存在，返回nil
func (t *ThingModel) GetAction(name string) *actions.ActionDescription {
	for _, p := range t.Actions {
		if p.Name == name {
			return &p
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
	return false, fmt.Errorf("action not found")
}
