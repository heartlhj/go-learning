package behavior

type CommandInterceptor interface {
	Execute(command Command) interface{}

	SetNext(next CommandInterceptor)
}
