package evaluator

import (
	"monkey-antlr/object"
	"monkey-antlr/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Evaluator struct {
	input  string
	env    *object.Environment
	parser *parser.MonkeyParser
}

func New(input string, env *object.Environment) *Evaluator {
	lexer := parser.NewMonkeyLexer(antlr.NewInputStream(input))
	p := parser.NewMonkeyParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))
	return &Evaluator{input: input, env: env, parser: p}
}

func (e *Evaluator) Eval() object.Object {
	l := &ProgramListener{}
	antlr.ParseTreeWalkerDefault.Walk(l, e.parser.Prog())

	return l.Pop()
}
