package main


type LogMessage struct {
	Message string `json:"name"`
	Timestamp int64 `json:"timestamp"`
	Severity string `json:"severity"`
}

type LogMessages []LogMessage
