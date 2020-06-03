package spel

import . "go-learning/expression"

type ExpressionState struct {
	relatedContext EvaluationContext

	rootObject TypedValue
}
