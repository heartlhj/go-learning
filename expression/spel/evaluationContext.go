package spel

type EvaluationContext interface {
	SetVariable(var1 string, var2 map[string]interface{})

	SetVariables(var2 map[string]interface{})

	LookupVariable(name string) interface{}
}
