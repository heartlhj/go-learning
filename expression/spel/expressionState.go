package spel

type ExpressionState struct {
	RelatedContext EvaluationContext

	RootObject TypedValue
}

func (e *ExpressionState) LookupVariable(name string) TypedValue {
	variable := e.RelatedContext.LookupVariable(name)
	return TypedValue{variable}
}
