package test

import (
	"fmt"
	. "github.com/heartlhj/go-expression/expression"
	"github.com/heartlhj/go-expression/expression/spel"
	"reflect"
	"testing"
)

type Order struct {
	name string
	age  int
}

//测试数组
func TestIndex(t *testing.T) {
	context := spel.StandardEvaluationContext{}
	context.AddPropertyAccessor(spel.MapAccessor{})
	m := make(map[string]interface{})
	m["name"] = "lisi"
	m["age"] = 18
	m1 := make(map[string]interface{})
	//切片
	//orders := make([]Order, 2)
	//数组
	orders := [2]Order{}
	orders[0] = Order{name: "lisi", age: 12}
	orders[1] = Order{name: "wang", age: 24}
	m1["code"] = orders
	m["order"] = m1
	context.SetVariables(m)
	parser := SpelExpressionParser{}
	valueContext := parser.ParseExpression("#order.code[0].name=='lisi'").GetValueContext(&context)
	fmt.Println("结果为：", valueContext)
}

//测试复合属性，对象里的字段
func TestCompound(t *testing.T) {
	context := spel.StandardEvaluationContext{}
	context.AddPropertyAccessor(spel.MapAccessor{})
	m := make(map[string]interface{})
	m["name"] = "lisi"
	m["age"] = 18
	m1 := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m2["num"] = 12
	m1["code"] = m2
	m["order"] = m1
	context.SetVariables(m)
	parser := SpelExpressionParser{}
	valueContext := parser.ParseExpression("#order.code.num==12").GetValueContext(&context)
	fmt.Println("结果为：", valueContext)
}

//测试字符相等
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

//测试与操作
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

//测试大于等于
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

//测试float类型 #num>=9f 会转为float64 #num>=9转为int
func TestFloat(t *testing.T) {
	context := spel.StandardEvaluationContext{}
	m := make(map[string]interface{})
	var ageFloat float64
	ageFloat = 10
	m["num"] = ageFloat
	context.SetVariables(m)
	parser := SpelExpressionParser{}
	valueContext := parser.ParseExpression("#num>=9f").GetValueContext(&context)
	fmt.Println("结果为：", valueContext)
}

//从interface获取数组/切片指定下标结构体的某一字段
func TestReflect(t *testing.T) {
	orders := make([]Order, 2)
	orders[0] = Order{name: "lisi", age: 18}
	orders[1] = Order{name: "wang", age: 24}
	var test interface{}
	test = orders
	//获取orders切片下标为1的数据
	index := reflect.ValueOf(test).Index(1)
	//取得Order对象类型
	orderType := index.Type()
	//取得name属性
	nameFile, _ := orderType.FieldByName("name")
	//取得name的类型
	nameType := nameFile.Type.Kind()
	//取得name字段
	name := index.FieldByName("name")
	var nameValue interface{}
	switch nameType {
	case reflect.String:
		nameValue = name.String()
		break
	case reflect.Int:
		nameValue = name.Int()
		break
	}
	fmt.Println("name字段类型为：", nameType)
	fmt.Println("name字段值为：", nameValue)
}
