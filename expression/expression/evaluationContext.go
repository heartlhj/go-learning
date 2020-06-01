package expression

type EvaluationContext interface {
	setVariable(var1 string, var2 map[interface{}]interface{})
}
