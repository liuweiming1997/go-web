package logs

import "testing"

func TestLogToFile(t *testing.T) {
	LogToFile("test")
	LogToFile("")
}
