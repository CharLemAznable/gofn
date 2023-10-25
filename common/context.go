package common

type Context interface {
	Get() interface{}
	Set(interface{})
	GetErr() error
	SetErr(error)
	Interrupted() bool
	SetInterrupt(bool)
}
