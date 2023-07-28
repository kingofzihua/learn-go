package log

import "strings"

// A Level is the importance or severity of a log event.
// The higher the level, the more important or severe the event.
type Level int

// LevelKey is logger level key.
const LevelKey = "level"

const (
	// LevelDebug is logger debug level.
	LevelDebug Level = -4
	// LevelInfo is logger info level.
	LevelInfo = 0
	// LevelWarn is logger warn level.
	LevelWarn = 4
	// LevelError is logger error level.
	LevelError = 8
)

func (l Level) Key() string {
	return LevelKey
}

// Level returns the receiver.
// It implements Leveler.
func (l Level) Level() Level { return l }

// String returns a name for the level.
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return ""
	}
}

// ParseLevel parses a level string into a logger Level value.
func ParseLevel(s string) Level {
	switch strings.ToUpper(s) {
	case "DEBUG":
		return LevelDebug
	case "INFO":
		return LevelInfo
	case "WARN":
		return LevelWarn
	case "ERROR":
		return LevelError
	}

	return LevelInfo
}
