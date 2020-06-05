package main

import (
	"fmt"
	. "go-learning/expression"
	"go-learning/expression/spel"
)

//测试
func main() {
	context := spel.StandardEvaluationContext{}
	m := make(map[string]interface{})
	m["name"] = 200
	context.SetVariables(m)
	parser := SpelExpressionParser{}
	valueContext := parser.ParseExpression("#name <= 700").GetValueContext(&context)
	fmt.Println("结果为：", valueContext)
}

func add(start int, end int) int {
	i := start + end
	fmt.Println(i)
	return i
}
