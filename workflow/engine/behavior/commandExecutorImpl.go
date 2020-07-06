package behavior

type CommandExecutorImpl struct {
	First CommandInterceptor
}

func (comm CommandExecutorImpl) Exe(conf Command) interface{} {
	return comm.First.Execute(conf)
}
