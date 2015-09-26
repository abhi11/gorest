package main


type LogMessage struct {
	Message string `json:"message"`
	Timestamp int64 `json:"timestamp"`
	Severity string `json:"severity"`
}

type LogMessages []LogMessage
