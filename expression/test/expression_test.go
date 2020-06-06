package test

import (
	"fmt"
	. "go-learning/expression"
	"go-learning/expression/spel"
	"testing"
)

//测试
func TestEQ(t *testing.T) {
	context := spel.StandardEvaluationContext{}
	m := make(map[string]interface{})
	m["name"] = "lisi"
	m["age"] = 18
	context.SetVariables(m)
	parser := SpelExpressionParser{}
	valueContext := parser.ParseExpression("#name=='lisi'").GetValueContext(&context)
	fmt.Println("结果为：", valueContext)
}

func TestAnd(t *testing.T) {
	context := spel.StandardEvaluationContext{}
	m := make(map[string]interface{})
	m["name"] = "lisi"
	m["age"] = 18
	context.SetVariables(m)
	parser := SpelExpressionParser{}
	valueContext := parser.ParseExpression("#name=='lisi' && #age>=3").GetValueContext(&context)
	fmt.Println("结果为：", valueContext)
}

func TestGT(t *testing.T) {
	context := spel.StandardEvaluationContext{}
	m := make(map[string]interface{})
	m["name"] = "lisi"
	m["age"] = 18
	context.SetVariables(m)
	parser := SpelExpressionParser{}
	valueContext := parser.ParseExpression("#age>=10").GetValueContext(&context)
	fmt.Println("结果为：", valueContext)
}

func TestFloat(t *testing.T) {
	context := spel.StandardEvaluationContext{}
	m := make(map[string]interface{})
	var ageFloat float64
	ageFloat = 10
	m["num"] = ageFloat
	context.SetVariables(m)
	parser := SpelExpressionParser{}
	valueContext := parser.ParseExpression("#num>=9.8").GetValueContext(&context)
	fmt.Println("结果为：", valueContext)
}
