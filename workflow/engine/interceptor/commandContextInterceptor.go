package interceptor

type CommandContextInterceptor struct {
	Next                       CommandInterceptor
	ProcessEngineConfiguration *ProcessEngineConfiguration
	CommandContextFactory      CommandContextFactory
}

func (a CommandContextInterceptor) Execute(command Command) interface{} {
	context, err := GetCommandContext()
	if err != nil {
		context = a.CommandContextFactory.CreateCommandContext(command)
	}
	SetCommandContext(context)
	return a.Next.Execute(command)
}

func (a *CommandContextInterceptor) SetNext(next CommandInterceptor) {
	a.Next = next
}
