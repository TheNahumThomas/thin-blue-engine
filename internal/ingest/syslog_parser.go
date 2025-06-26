package ingest

import (
	"fmt"

	syslog "github.com/leodido/go-syslog/v4"
	"github.com/leodido/go-syslog/v4/rfc5424"
)

type SyslogParser struct {
	parser syslog.Machine
}

func NewSyslogParser(options ...syslog.MachineOption) *SyslogParser {
	return &SyslogParser{
		parser: rfc5424.NewParser(rfc5424.WithBestEffort()),
	}
}

func (p *SyslogParser) BuildLogObject(syslog []byte) (*LogEvent, error) {
	l := NewLogEvent()

	parsedSyslog, err := p.Parse(syslog)
	if err != nil {
		return nil, fmt.Errorf("Failed to build log object from syslog message: %w", err)
	}

	l.CustomFields = make(map[string]interface{})

	if _, err := p.CopyTo(l, parsedSyslog); err != nil {
		return nil, fmt.Errorf("Failed to copy syslog message to log object: %w", err)
	}
	// WRITE CODE FOR REGEXP PARSING OF SDID AND SD-PARAMS

	return l, nil
}

func (p *SyslogParser) Parse(line []byte) (syslog.Message, error) {
	parsedMsg, err := p.parser.Parse(line)
	if err != nil {
		return nil, fmt.Errorf("Failed to pass syslog message to parser: %w", err)
	}

	return parsedMsg, nil
}

func (p *SyslogParser) CopyTo(l *LogEvent, m syslog.Message) (*LogEvent, error) {

	base, ok := m.(*rfc5424.SyslogMessage)
	if !ok {
		return nil, fmt.Errorf("Failed to cast syslog message to rfc5424.SyslogMessage (syslog_parser.CopyTo)")
	}

	if base.Timestamp != nil {
		l.EventTimestamp = *base.Timestamp
	} else {
		l.EventTimestamp = l.EngineTimestamp
	}
	if base.Priority != nil {
		l.LoggerPriority = *base.Priority
	} else {
		l.LoggerPriority = 0
	}
	if base.Hostname != nil {
		l.Host = *base.Hostname
	} else {
		l.Host = ""
	}
	if base.Appname != nil {
		l.Application = *base.Appname
	} else {
		l.Application = ""
	}
	if base.ProcID != nil {
		l.ProcessID = *base.ProcID
	} else {
		l.ProcessID = ""
	}
	if base.MsgID != nil {
		l.MessageID = *base.MsgID
	} else {
		l.MessageID = ""
	}
	if base.StructuredData != nil {
		for sdID, params := range *base.StructuredData {
			fieldMap := make(map[string]interface{})
			for paramName, paramValue := range params {
				fieldMap[paramName] = paramValue
			}
			l.CustomFields[string(sdID)] = fieldMap
		}
	} else {
		l.CustomFields = make(map[string]interface{})
	}

	return &LogEvent{}, nil
}

// func getStructuredData(customFields map[string]interface{}) map[string]interface{} {
// 	re :=

// }
