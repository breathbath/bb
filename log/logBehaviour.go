package log

type LogBehaviour interface {
	Log(message string)
	Logf(template string, values ... interface{})
}
