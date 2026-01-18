package events

import (
	"fmt"

	"github.com/AtomPod/thingmodel/thingmodel/dataspec"
)

// EventDescription 事件描述，用于消息上报等，包含几种情况，例如info、alert、error、warning等
type EventDescription struct {
	// Name 事件名称
	Name string `json:"name"`

	// Description 事件描述
	Description string `json:"description"`

	// Type 事件类型
	Type EventType `json:"type"`

	// Data 上报的数据描述
	Data *dataspec.DataDescription `json:"data"`
}

func (e *EventDescription) UpdateData() error {
	if len(e.Name) == 0 {
		return fmt.Errorf("EventDescription: name could not be empty")
	}

	if e.Data == nil {
		return fmt.Errorf("EventDescription: data field could not be empty")
	}

	if err := e.Data.Parse(); err != nil {
		return err
	}
	return nil
}

// Validate 验证数据是否正确
func (e *EventDescription) Validate(v interface{}) (bool, error) {
	return e.Data.Validate(v)
}
