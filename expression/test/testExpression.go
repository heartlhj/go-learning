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
	parser.ParseExpression("#name=='li'").GetValueContext(&context)
}

func add(start int, end int) int {
	i := start + end
	fmt.Println(i)
	return i
}
