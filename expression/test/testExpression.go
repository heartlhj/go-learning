package main

import (
	"fmt"
	. "go-learning/expression"
	"go-learning/expression/spel"
)

func main() {
	context := spel.StandardEvaluationContext{}
	m := make(map[string]interface{})
	m["name"] = "lisi"
	context.SetVariables(m)
	parser := SpelExpressionParser{}
	valueContext := parser.ParseExpression("#name=='lisi'").GetValueContext(&context)
	fmt.Println("结果为：", valueContext)
}

func add(start int, end int) int {
	i := start + end
	fmt.Println(i)
	return i
}
