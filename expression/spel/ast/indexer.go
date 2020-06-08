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
	context := state.GetActiveContextObject()
	target := context.Value
	targetDescriptor := context.GetTypeDescriptor()
	var indexValue TypedValue
	var index interface{}
	_, ok := target.(map[interface{}]interface{})
	reference, isOK := this.Children[0].(PropertyOrFieldReference)
	if ok && isOK {
		index = reference.Name
		indexValue = TypedValue{Value: index}
	} else {
		defer state.PopActiveContextObject()
		state.PushActiveContextObject(state.GetRootContextObject())
		indexValue = this.Children[0].GetValueInternal(state)
		index = indexValue.Value
		if index == nil {
			panic("No index")
		}
	}

	if target == nil {
		panic("Cannot index into a null value")
	}
	_, okArr := target.([]interface{})
	if okArr {
		var key interface{}
		key = index

	}
	return nil
}

func (this Indexer) GetValueInternal(state ExpressionState) TypedValue {
	return this.GetValueRef(state).GetValue()
}

func (this *Indexer) setValue(state ExpressionState) TypedValue {
	return this.GetValueRef(state).GetValue()
}
