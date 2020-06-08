package ast

import (
	. "github.com/heartlhj/go-learning/expression/spel"
)

//多维数组
type Indexer struct {
	*SpelNodeImpl
	cachedReadName     string
	cachedReadAccessor PropertyAccessor
}

func (this Indexer) GetValueRef(state ExpressionState) ValueRef {
	//context := state.GetActiveContextObject()
	//target := context.Value
	//targetDescriptor := context.GetTypeDescriptor()

	return nil
}

func (this Indexer) GetValueInternal(state ExpressionState) TypedValue {
	return this.GetValueRef(state).GetValue()
}

func (this *Indexer) setValue(state ExpressionState) TypedValue {
	return this.GetValueRef(state).GetValue()
}
