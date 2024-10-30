package test

import (
	"fmt"
	"testing"

	"github.com/szlizard/gengine/builder"
	"github.com/szlizard/gengine/context"
	"github.com/szlizard/gengine/engine"
)

const in_operator_rule = `
rule "test in operator"
	begin
		if "hello" in "hello world" {
			println("hello Substring found!")
		}
		if "5" in InputAndResult.GetSlice() {
			println("Element found '5' in slice!")
		}
		if 4 in InputAndResult.GetSlice() {
			println("Element found 4 in slice!")
		}
		if 2.00000000000000 in InputAndResult.GetSlice() {
			println("Element found 2.0 in slice!")
		}
		if 2.00000000000001 in InputAndResult.GetSlice() {
			println("Element found 2.00000000000001 in slice!")
		}
		if -3.0000 in InputAndResult.GetSlice() {
			println("Element found -3 in slice!")
		}
		if "key" in InputAndResult.GetMap() {
			println("Key found in map!")
		}
		if "key1" in InputAndResult.GetMap() {
			println("Key1 found in map!")
		}
	end
`

func Test_in_operator(t *testing.T) {
	dataContext := context.NewDataContext()
	dataContext.Add("println", fmt.Println)

	InputAndResult := &InOperatorInputAndResult{}
	dataContext.Add("InputAndResult", InputAndResult)

	ruleBuilder := builder.NewRuleBuilder(dataContext)
	e1 := ruleBuilder.BuildRuleFromString(in_operator_rule)
	if e1 != nil {
		t.Fatalf("BuildRuleFromString error: %v", e1)
	}

	gengine := engine.NewGengine()
	e := gengine.Execute(ruleBuilder, true)
	if e != nil {
		t.Fatalf("Execute error: %v", e)
	}
}

type InOperatorInputAndResult struct{}

func (i *InOperatorInputAndResult) GetSlice() []interface{} {
	var aa float32 = 2.0
	var bb float64 = 2.0
	return []interface{}{1, aa, bb, -3, 4, "5"}
}

func (i *InOperatorInputAndResult) GetMap() map[interface{}]interface{} {
	return map[interface{}]interface{}{"key": "value"}
}
