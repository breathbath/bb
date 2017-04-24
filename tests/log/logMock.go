package log

type LogMock struct {
	Messages []string
	FormattedMessages []FormattedCall
}

type FormattedCall struct {
	Message string
	Arguments [] interface{}
}

func NewLogMock() *LogMock{
	return &LogMock{[]string{}, []FormattedCall{}}
}

func (l *LogMock) Log(message string) {
	l.Messages = append(l.Messages, message)
}
func (l *LogMock) Logf(template string, values ... interface{}) {
	call := FormattedCall{template, values}
	l.FormattedMessages = append(l.FormattedMessages, call)
}

func (l *LogMock) HasMessages() bool{
	return len(l.FormattedMessages) > 0 || len(l.Messages) > 0
}


