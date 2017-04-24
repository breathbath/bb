package log

import "fmt"

type TtlLog struct {

}

func (ttlLog TtlLog) Log(message string) {
	fmt.Println(message)
}

func (ttlLog TtlLog) Logf(template string, values ... interface{}) {
	fmt.Printf(template + "\n", values...)
}