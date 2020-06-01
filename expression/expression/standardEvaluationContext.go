package expression

type StandardEvaluationContext struct {
	variables map[interface{}]interface{}
}

func (s *StandardEvaluationContext) setVariable(var1 string, var2 map[interface{}]interface{}) {
	s.variables[var1] = var2
}
