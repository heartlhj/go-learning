package spel

//map存储
type StandardEvaluationContext struct {
	Variables map[string]interface{}
}

func (s *StandardEvaluationContext) SetVariable(var1 string, var2 map[string]interface{}) {
	s.Variables[var1] = var2
}

func (s *StandardEvaluationContext) SetVariables(var2 map[string]interface{}) {
	s.Variables = var2
}

func (s *StandardEvaluationContext) LookupVariable(name string) interface{} {
	return s.Variables[name]
}
