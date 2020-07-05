package interceptor

type CommandContextFactory struct {
}

func (factory CommandContextFactory) CreateCommandContext(command Command) CommandContext {
	context := CommandContext{Command: command}
	return context
}
