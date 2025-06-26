package ingest

import (
	"fmt"
	"net"
	"time"
)

type LogLevel int

const (
	LevelUnknown LogLevel = iota // Default
	LevelDebug
	LevelInfo
	LevelNotice
	LevelWarning
	LevelError
	LevelCritical
	LevelAlert
	LevelEmergency
)

func (l LogLevel) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelNotice:
		return "NOTICE"
	case LevelWarning:
		return "WARNING"
	case LevelError:
		return "ERROR"
	case LevelCritical:
		return "CRITICAL"
	case LevelAlert:
		return "ALERT"
	case LevelEmergency:
		return "EMERGENCY"
	default:
		return "UNKNOWN"
	}
}

type LogEvent struct {
	EventTimestamp  time.Time
	EngineTimestamp time.Time
	Level           LogLevel
	LoggerPriority  uint8
	Host            string
	Application     string
	ProcessID       string
	MessageID       string
	SourceAddr      [16]byte
	DestAddr        [16]byte
	SourcePort      uint16
	DestinationPort uint16
	Protocol        string
	Message         string
	RawJSON         string
	CustomFields    map[string]interface{}
}

func NewLogEvent() *LogEvent {
	return &LogEvent{
		EngineTimestamp: time.Now().UTC(),
		Level:           LevelUnknown,
		LoggerPriority:  0,
		SourceAddr:      [16]byte{}, // Initialize to all zeros
		DestAddr:        [16]byte{},
	}
}

// type parser interface {
// 	Parse(line []byte) (*LogEvent, error)
// 	MatchFormat(line []byte) bool
// }

func SetSourceAddr(ipstring string) (*[16]byte, error) {
	ip := net.ParseIP(ipstring)
	if ip == nil {
		return nil, fmt.Errorf("Couldn't Parse IP addresses from Logs")
	}
	if len(ip) == net.IPv4len {
		var addr [16]byte
		copy(addr[12:], ip.To4())
		return &addr, nil
	}
	var addr [16]byte
	copy(addr[:], ip.To16())
	return &addr, nil
}
