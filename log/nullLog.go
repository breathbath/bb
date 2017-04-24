package log

type NullLog struct {

}

func (nl NullLog) Log(message string) {
}

func (nl NullLog) Logf(template string, values ... interface{}) {

}
