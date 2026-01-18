package events

// EventType 事件类型
type EventType string

const (
	// Info 普通消息，包括各种推送
	Info EventType = "info"

	// Warning 警告消息
	Warning EventType = "warning"

	// Error 错误消息，比如某些功能无法启用等
	Error EventType = "error"

	// Alert 警报消息，非常严重的信息，比如传感器感应到危险信息
	Alert EventType = "alert"
)
