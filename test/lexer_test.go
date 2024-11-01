package test

import (
	"testing"

	"github.com/szlizard/gengine/builder"
	"github.com/szlizard/gengine/context"
	"github.com/szlizard/gengine/engine"
)

func Test_lexer(t *testing.T) {

	lexer_rule := `
rule "test" salience 1
begin
规则管理
end
`
	dataContext := context.NewDataContext()
	ruleBuilder := builder.NewRuleBuilder(dataContext)
	e1 := ruleBuilder.BuildRuleFromString(lexer_rule)
	if e1 != nil {
		t.Logf("build rule error: %v", e1)
	}

	gengine := engine.NewGengine()
	e := gengine.Execute(ruleBuilder, true)
	if e != nil {
		t.Logf("build rule error: %v", e)
	}

}
