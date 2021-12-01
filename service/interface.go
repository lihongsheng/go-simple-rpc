package service

type Ser interface {
	Listen(port int, scheme string) error
	Close()
	AddReadHandler(func(param interface{}))
	AddWriteHandler(func(param interface{}))
}
