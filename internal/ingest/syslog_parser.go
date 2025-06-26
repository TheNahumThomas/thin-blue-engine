package ingest

import (
	"fmt"
	"time"

	"github.com/influxdata/go-syslog/v3"
	"github.com/influxdata/go-syslog/v3/rfc5424"
	rfc "github.com/influxdata/go-syslog/v3/rfc5424"
)

type SyslogParser struct {
	parser syslog.Machine
}

func NewSyslogParser() *SyslogParser {
	return &SyslogParser{
		parser: rfc.NewMachine(rfc.WithBestEffort()),
	}

}

func (p *SyslogParser) Parse(line []byte) (*LogEvent, error) {
	msg := NewLogEvent()
	parsed, err := p.parser.Parse(line)
	base, ok := parsed.(*rfc5424.SyslogMessage)
	if !ok {
		return msg, fmt.Errorf("Couldn't retrieve RFC5424.SysLogMessage Object: %w", err)
	}
	if err != nil {
		return msg, fmt.Errorf("failed to parse syslog message: %w", err)
	}
	if parsed == nil {
		return &LogEvent{}, fmt.Errorf("parsed syslog message is nil")
	}

	msg = &LogEvent{
		EventTimestamp:  time.Now(),
		Host:            "",
		Application:     "",
		ProcessID:       "",
		MessageID:       "",
		SourcePort:      0,
		DestinationPort: 0,
		Protocol:        "",
		Message:         "",
		RawJSON:         "",
		CustomFields:    make(map[string]interface{}),
	}

	if base.Timestamp != nil {
		msg.EventTimestamp = *base.Timestamp
	}
	if base.Priority != nil {
		msg.LoggerPriority = uint8(*base.Priority)
	}
	if base.Hostname != nil {
		msg.Host = *base.Hostname
	}
	if base.Appname != nil {
		msg.Application = *base.Appname
	}
	if base.ProcID != nil {
		msg.ProcessID = *base.ProcID
	}
	if base.MsgID != nil {
		msg.MessageID = *base.MsgID
	}
	if base.StructuredData != nil {
		for sdID, params := range *base.StructuredData {
			fieldMap := make(map[string]string)
			for paramName, paramValue := range params {
				fieldMap[paramName] = paramValue
			}
			msg.CustomFields[string(sdID)] = fieldMap
		}
	}

	return msg, nil

}
