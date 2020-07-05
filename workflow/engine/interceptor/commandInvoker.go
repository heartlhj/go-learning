package interceptor

type CommandInvoker struct {
	Next CommandInterceptor
}

func (a CommandInvoker) Execute(command Command) interface{} {
	context, err := GetCommandContext()
	if err != nil {

	}
	execute := command.Execute(context)
	executeOperations(context)
	return execute
}

func executeOperations(context CommandContext) {
	context.Agenda.GetNextOperation().Run()
}

func (a *CommandInvoker) SetNext(next CommandInterceptor) {
	a.Next = next
}