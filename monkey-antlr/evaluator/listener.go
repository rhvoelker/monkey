package evaluator

import (
	"fmt"
	"monkey-antlr/object"
	"monkey-antlr/parser"
	"strconv"
)

type ProgramListener struct {
	*parser.BaseMonkeyListener

	stack  []object.Object
	errors []string
}

func (l *ProgramListener) Push(item object.Object) {
	l.stack = append(l.stack, item)
}

func (l *ProgramListener) Pop() object.Object {
	if len(l.stack) == 0 {
		return nil
	}

	result := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *ProgramListener) ExitIntegerLiteral(c *parser.IntegerLiteralContext) {
	value, err := strconv.ParseInt(c.GetText(), 0, 64)
	if err != nil {
		panic(fmt.Sprintf("Could not parse '%q' as integer. Check grammar.", c.GetText()))
	}
	l.Push(&object.Integer{Value: value})
}
