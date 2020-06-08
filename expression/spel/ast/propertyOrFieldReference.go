package ast

import (
	. "github.com/heartlhj/go-learning/expression/spel"
)

type PropertyOrFieldReference struct {
	*SpelNodeImpl
	NullSafe                            bool
	Name                                string
	OriginalPrimitiveExitTypeDescriptor string
	CachedReadAccessor                  PropertyAccessor
}

type ValueRef interface {
	GetValue() TypedValue
}

type AccessorLValue struct {
	ref                    PropertyOrFieldReference
	contextObject          TypedValue
	evalContext            EvaluationContext
	autoGrowNullReferences bool
}

func (this AccessorLValue) GetValue() TypedValue {
	return this.ref.getValueInternal(this.contextObject, this.evalContext, this.autoGrowNullReferences)
}

func (this PropertyOrFieldReference) GetValueRef(state ExpressionState) ValueRef {
	return AccessorLValue{ref: this, contextObject: state.GetActiveContextObject(), evalContext: state.RelatedContext,
		autoGrowNullReferences: false}
}

func (this PropertyOrFieldReference) GetValueInternal(state ExpressionState) TypedValue {
	return this.getValueInternal(state.GetActiveContextObject(), state.GetEvaluationContext(), false)
}

func (this PropertyOrFieldReference) getValueInternal(contextObject TypedValue, evalContext EvaluationContext, isAutoGrowNullReferences bool) TypedValue {
	return this.readProperty(contextObject, evalContext, this.Name)
}

func (this PropertyOrFieldReference) readProperty(contextObject TypedValue, evalContext EvaluationContext, name string) TypedValue {
	accessors := evalContext.GetPropertyAccessors()
	for _, accessor := range accessors {
		_, ok := accessor.(ReflectivePropertyAccessor)
		if !ok {
			this.CachedReadAccessor = accessor
			return accessor.Read(evalContext, contextObject.Value, name)
		}
	}
	return TypedValue{}
}
