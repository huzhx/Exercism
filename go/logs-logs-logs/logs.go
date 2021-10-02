package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Message extracts the message from the provided log line.
func Message(line string) string {
	return strings.TrimSpace(strings.Split(line, "]: ")[1])
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	message := Message(line)
	return utf8.RuneCountInString(message)
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	temp := strings.Split(line, "[")[1]
	level := strings.Split(temp, "]: ")[0]
	return strings.ToLower(level)
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	message := Message(line)
	level := LogLevel(line)
	return fmt.Sprintf("%v (%v)", message, level)
}
