package evaluator

import (
	"monkey-antlr/object"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Eval_Integers(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		//{"-5", -5},
		//{"-10", -10},
		//{"5 + 5 + 5 + 5 - 10", 10},
		//{"2 * 2 * 2 * 2 * 2", 32},
		//{"-50 + 100 - 50", 0},
		//{"5 * 2 + 10", 20},
		//{"5 + 2 * 10", 25},
		//{"20 + 2 * -10", 0},
		//{"50 / 2 * 2 + 10", 60},
		//{"2 * (5 + 10)", 30},
		//{"3 * 3 * 3 + 10", 37},
		//{"3 * (3 * 3) + 10", 37},
		//{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	eval := New(input, object.NewEnvironment())
	return eval.Eval()
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) {
	assert.IsType(t, &object.Integer{}, obj)
	result := obj.(*object.Integer).Value
	assert.Equal(t, expected, result)
}
