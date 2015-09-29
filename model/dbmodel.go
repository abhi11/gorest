package main


type LogMessage struct {
	Timestamp int64 `json:"timestamp"`
	LogLevel string `json:"log_level"`
	MobileDeviceId string `json:"mobile_device_id"`
	TwistDeviceId string `json:"twist_device_id"`
	Log string `json:"log"`
}

type LogMessages []LogMessage
