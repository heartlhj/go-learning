package ast

import (
	. "github.com/heartlhj/go-learning/expression/spel"
)

type CompoundExpression struct {
	*SpelNodeImpl
}

//复合表达式
func (this *CompoundExpression) GetValueInternal(state ExpressionState) TypedValue {
	if this.getChildCount() == 1 {
		return TypedValue{}
	}
	nextNode := this.Children[0]
	result := nextNode.GetValueInternal(state)
	count := this.getChildCount()
	for i := 1; i < count-1; i++ {
		defer state.PopActiveContextObjectNull()
		state.PushActiveContextObject(result)
		nextNode = this.Children[i]
		result = nextNode.GetValueInternal(state)
	}
	defer state.PushActiveContextObject(result)
	state.PushActiveContextObject(result)
	nextNode = this.Children[count-1]

	ref := nextNode.GetValueRef(state)
	return ref.GetValue()

}

func (o *CompoundExpression) getChildCount() int {
	return len(o.Children)
}
