package interceptor

type CommandInterceptor interface {
	Execute(command Command) interface{}

	SetNext(next CommandInterceptor)
}
