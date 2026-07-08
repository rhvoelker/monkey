// Code generated from C:/Users/Sheep/Documents/Projects/monkey-interpreter/monkey-antlr/parser/Monkey.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Monkey
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MonkeyParser struct {
	*antlr.BaseParser
}

var MonkeyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func monkeyParserInit() {
	staticData := &MonkeyParserStaticData
	staticData.LiteralNames = []string{
		"", "'fn'", "'let'", "'true'", "'false'", "'if'", "'else'", "'return'",
		"", "", "", "'='", "'+'", "'-'", "'!'", "'*'", "'/'", "'<'", "'>'",
		"'=='", "'!='", "','", "':'", "';'", "'('", "')'", "'{'", "'}'", "'['",
		"']'",
	}
	staticData.SymbolicNames = []string{
		"", "FUNCTION", "LET", "TRUE", "FALSE", "IF", "ELSE", "RETURN", "STRING",
		"IDENT", "INT", "ASSIGN", "PLUS", "MINUS", "BANG", "ASTERISK", "SLASH",
		"LT", "GT", "EQ", "NOT_EQ", "COMMA", "COLON", "SEMICOLON", "LPAREN",
		"RPAREN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "stmt", "expr", "let_stmt", "return_stmt", "expression_stmt",
		"expr_list", "args_list",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 30, 141, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 4, 0, 18, 8, 0, 11, 0, 12,
		0, 19, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 27, 8, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 5, 2, 47, 8, 2, 10, 2, 12, 2, 50, 9, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 5, 2, 56, 8, 2, 10, 2, 12, 2, 59, 9, 2, 1, 2, 3, 2, 62, 8, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 70, 8, 2, 10, 2, 12, 2, 73, 9,
		2, 1, 2, 1, 2, 3, 2, 77, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 5, 2, 101, 8, 2, 10, 2, 12, 2, 104, 9, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 5, 6, 123, 8, 6, 10, 6, 12, 6, 126, 9, 6, 3, 6, 128,
		8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 134, 8, 7, 10, 7, 12, 7, 137, 9, 7,
		3, 7, 139, 8, 7, 1, 7, 0, 1, 4, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 5, 1,
		0, 13, 14, 1, 0, 15, 16, 1, 0, 12, 13, 1, 0, 17, 18, 1, 0, 19, 20, 157,
		0, 17, 1, 0, 0, 0, 2, 26, 1, 0, 0, 0, 4, 76, 1, 0, 0, 0, 6, 105, 1, 0,
		0, 0, 8, 111, 1, 0, 0, 0, 10, 115, 1, 0, 0, 0, 12, 127, 1, 0, 0, 0, 14,
		138, 1, 0, 0, 0, 16, 18, 3, 2, 1, 0, 17, 16, 1, 0, 0, 0, 18, 19, 1, 0,
		0, 0, 19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 22,
		5, 0, 0, 1, 22, 1, 1, 0, 0, 0, 23, 27, 3, 6, 3, 0, 24, 27, 3, 8, 4, 0,
		25, 27, 3, 10, 5, 0, 26, 23, 1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 26, 25, 1,
		0, 0, 0, 27, 3, 1, 0, 0, 0, 28, 29, 6, 2, -1, 0, 29, 77, 5, 10, 0, 0, 30,
		77, 5, 3, 0, 0, 31, 77, 5, 4, 0, 0, 32, 77, 5, 8, 0, 0, 33, 77, 5, 9, 0,
		0, 34, 35, 7, 0, 0, 0, 35, 77, 3, 4, 2, 10, 36, 37, 5, 28, 0, 0, 37, 38,
		3, 12, 6, 0, 38, 39, 5, 29, 0, 0, 39, 77, 1, 0, 0, 0, 40, 41, 5, 5, 0,
		0, 41, 42, 5, 24, 0, 0, 42, 43, 3, 4, 2, 0, 43, 44, 5, 25, 0, 0, 44, 48,
		5, 26, 0, 0, 45, 47, 3, 2, 1, 0, 46, 45, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0,
		48, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 51, 1, 0, 0, 0, 50, 48, 1,
		0, 0, 0, 51, 61, 5, 27, 0, 0, 52, 53, 5, 6, 0, 0, 53, 57, 5, 26, 0, 0,
		54, 56, 3, 2, 1, 0, 55, 54, 1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1,
		0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 60,
		62, 5, 27, 0, 0, 61, 52, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 77, 1, 0,
		0, 0, 63, 64, 5, 1, 0, 0, 64, 65, 5, 24, 0, 0, 65, 66, 3, 14, 7, 0, 66,
		67, 5, 25, 0, 0, 67, 71, 5, 26, 0, 0, 68, 70, 3, 2, 1, 0, 69, 68, 1, 0,
		0, 0, 70, 73, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 74,
		1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 75, 5, 27, 0, 0, 75, 77, 1, 0, 0, 0,
		76, 28, 1, 0, 0, 0, 76, 30, 1, 0, 0, 0, 76, 31, 1, 0, 0, 0, 76, 32, 1,
		0, 0, 0, 76, 33, 1, 0, 0, 0, 76, 34, 1, 0, 0, 0, 76, 36, 1, 0, 0, 0, 76,
		40, 1, 0, 0, 0, 76, 63, 1, 0, 0, 0, 77, 102, 1, 0, 0, 0, 78, 79, 10, 7,
		0, 0, 79, 80, 7, 1, 0, 0, 80, 101, 3, 4, 2, 8, 81, 82, 10, 6, 0, 0, 82,
		83, 7, 2, 0, 0, 83, 101, 3, 4, 2, 7, 84, 85, 10, 5, 0, 0, 85, 86, 7, 3,
		0, 0, 86, 101, 3, 4, 2, 6, 87, 88, 10, 4, 0, 0, 88, 89, 7, 4, 0, 0, 89,
		101, 3, 4, 2, 5, 90, 91, 10, 9, 0, 0, 91, 92, 5, 28, 0, 0, 92, 93, 3, 4,
		2, 0, 93, 94, 5, 29, 0, 0, 94, 101, 1, 0, 0, 0, 95, 96, 10, 8, 0, 0, 96,
		97, 5, 24, 0, 0, 97, 98, 3, 12, 6, 0, 98, 99, 5, 25, 0, 0, 99, 101, 1,
		0, 0, 0, 100, 78, 1, 0, 0, 0, 100, 81, 1, 0, 0, 0, 100, 84, 1, 0, 0, 0,
		100, 87, 1, 0, 0, 0, 100, 90, 1, 0, 0, 0, 100, 95, 1, 0, 0, 0, 101, 104,
		1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 5, 1, 0, 0,
		0, 104, 102, 1, 0, 0, 0, 105, 106, 5, 2, 0, 0, 106, 107, 5, 9, 0, 0, 107,
		108, 5, 11, 0, 0, 108, 109, 3, 4, 2, 0, 109, 110, 5, 23, 0, 0, 110, 7,
		1, 0, 0, 0, 111, 112, 5, 7, 0, 0, 112, 113, 3, 4, 2, 0, 113, 114, 5, 23,
		0, 0, 114, 9, 1, 0, 0, 0, 115, 116, 3, 4, 2, 0, 116, 117, 5, 23, 0, 0,
		117, 11, 1, 0, 0, 0, 118, 128, 1, 0, 0, 0, 119, 124, 3, 4, 2, 0, 120, 121,
		5, 21, 0, 0, 121, 123, 3, 4, 2, 0, 122, 120, 1, 0, 0, 0, 123, 126, 1, 0,
		0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 128, 1, 0, 0, 0,
		126, 124, 1, 0, 0, 0, 127, 118, 1, 0, 0, 0, 127, 119, 1, 0, 0, 0, 128,
		13, 1, 0, 0, 0, 129, 139, 1, 0, 0, 0, 130, 135, 5, 9, 0, 0, 131, 132, 5,
		21, 0, 0, 132, 134, 5, 9, 0, 0, 133, 131, 1, 0, 0, 0, 134, 137, 1, 0, 0,
		0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 139, 1, 0, 0, 0, 137,
		135, 1, 0, 0, 0, 138, 129, 1, 0, 0, 0, 138, 130, 1, 0, 0, 0, 139, 15, 1,
		0, 0, 0, 13, 19, 26, 48, 57, 61, 71, 76, 100, 102, 124, 127, 135, 138,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MonkeyParserInit initializes any static state used to implement MonkeyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMonkeyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MonkeyParserInit() {
	staticData := &MonkeyParserStaticData
	staticData.once.Do(monkeyParserInit)
}

// NewMonkeyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMonkeyParser(input antlr.TokenStream) *MonkeyParser {
	MonkeyParserInit()
	this := new(MonkeyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MonkeyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Monkey.g4"

	return this
}

// MonkeyParser tokens.
const (
	MonkeyParserEOF       = antlr.TokenEOF
	MonkeyParserFUNCTION  = 1
	MonkeyParserLET       = 2
	MonkeyParserTRUE      = 3
	MonkeyParserFALSE     = 4
	MonkeyParserIF        = 5
	MonkeyParserELSE      = 6
	MonkeyParserRETURN    = 7
	MonkeyParserSTRING    = 8
	MonkeyParserIDENT     = 9
	MonkeyParserINT       = 10
	MonkeyParserASSIGN    = 11
	MonkeyParserPLUS      = 12
	MonkeyParserMINUS     = 13
	MonkeyParserBANG      = 14
	MonkeyParserASTERISK  = 15
	MonkeyParserSLASH     = 16
	MonkeyParserLT        = 17
	MonkeyParserGT        = 18
	MonkeyParserEQ        = 19
	MonkeyParserNOT_EQ    = 20
	MonkeyParserCOMMA     = 21
	MonkeyParserCOLON     = 22
	MonkeyParserSEMICOLON = 23
	MonkeyParserLPAREN    = 24
	MonkeyParserRPAREN    = 25
	MonkeyParserLBRACE    = 26
	MonkeyParserRBRACE    = 27
	MonkeyParserLBRACKET  = 28
	MonkeyParserRBRACKET  = 29
	MonkeyParserWS        = 30
)

// MonkeyParser rules.
const (
	MonkeyParserRULE_prog            = 0
	MonkeyParserRULE_stmt            = 1
	MonkeyParserRULE_expr            = 2
	MonkeyParserRULE_let_stmt        = 3
	MonkeyParserRULE_return_stmt     = 4
	MonkeyParserRULE_expression_stmt = 5
	MonkeyParserRULE_expr_list       = 6
	MonkeyParserRULE_args_list       = 7
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(MonkeyParserEOF, 0)
}

func (s *ProgContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MonkeyParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268462014) != 0) {
		{
			p.SetState(16)
			p.Stmt()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(21)
		p.Match(MonkeyParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Let_stmt() ILet_stmtContext
	Return_stmt() IReturn_stmtContext
	Expression_stmt() IExpression_stmtContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Let_stmt() ILet_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILet_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILet_stmtContext)
}

func (s *StmtContext) Return_stmt() IReturn_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturn_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturn_stmtContext)
}

func (s *StmtContext) Expression_stmt() IExpression_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_stmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MonkeyParserRULE_stmt)
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(23)
			p.Let_stmt()
		}

	case MonkeyParserRETURN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
			p.Return_stmt()
		}

	case MonkeyParserFUNCTION, MonkeyParserTRUE, MonkeyParserFALSE, MonkeyParserIF, MonkeyParserSTRING, MonkeyParserIDENT, MonkeyParserINT, MonkeyParserMINUS, MonkeyParserBANG, MonkeyParserLBRACKET:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(25)
			p.Expression_stmt()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	STRING() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	MINUS() antlr.TerminalNode
	BANG() antlr.TerminalNode
	LBRACKET() antlr.TerminalNode
	Expr_list() IExpr_listContext
	RBRACKET() antlr.TerminalNode
	IF() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllLBRACE() []antlr.TerminalNode
	LBRACE(i int) antlr.TerminalNode
	AllRBRACE() []antlr.TerminalNode
	RBRACE(i int) antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	ELSE() antlr.TerminalNode
	FUNCTION() antlr.TerminalNode
	Args_list() IArgs_listContext
	ASTERISK() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	EQ() antlr.TerminalNode
	NOT_EQ() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) INT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserINT, 0)
}

func (s *ExprContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserTRUE, 0)
}

func (s *ExprContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserFALSE, 0)
}

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSTRING, 0)
}

func (s *ExprContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIDENT, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MonkeyParserMINUS, 0)
}

func (s *ExprContext) BANG() antlr.TerminalNode {
	return s.GetToken(MonkeyParserBANG, 0)
}

func (s *ExprContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLBRACKET, 0)
}

func (s *ExprContext) Expr_list() IExpr_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_listContext)
}

func (s *ExprContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRBRACKET, 0)
}

func (s *ExprContext) IF() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIF, 0)
}

func (s *ExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLPAREN, 0)
}

func (s *ExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRPAREN, 0)
}

func (s *ExprContext) AllLBRACE() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserLBRACE)
}

func (s *ExprContext) LBRACE(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserLBRACE, i)
}

func (s *ExprContext) AllRBRACE() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserRBRACE)
}

func (s *ExprContext) RBRACE(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserRBRACE, i)
}

func (s *ExprContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ExprContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserELSE, 0)
}

func (s *ExprContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(MonkeyParserFUNCTION, 0)
}

func (s *ExprContext) Args_list() IArgs_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgs_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgs_listContext)
}

func (s *ExprContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(MonkeyParserASTERISK, 0)
}

func (s *ExprContext) SLASH() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSLASH, 0)
}

func (s *ExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPLUS, 0)
}

func (s *ExprContext) LT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLT, 0)
}

func (s *ExprContext) GT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserGT, 0)
}

func (s *ExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(MonkeyParserEQ, 0)
}

func (s *ExprContext) NOT_EQ() antlr.TerminalNode {
	return s.GetToken(MonkeyParserNOT_EQ, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *MonkeyParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, MonkeyParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserINT:
		{
			p.SetState(29)
			p.Match(MonkeyParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserTRUE:
		{
			p.SetState(30)
			p.Match(MonkeyParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserFALSE:
		{
			p.SetState(31)
			p.Match(MonkeyParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserSTRING:
		{
			p.SetState(32)
			p.Match(MonkeyParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserIDENT:
		{
			p.SetState(33)
			p.Match(MonkeyParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserMINUS, MonkeyParserBANG:
		{
			p.SetState(34)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MonkeyParserMINUS || _la == MonkeyParserBANG) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(35)
			p.expr(10)
		}

	case MonkeyParserLBRACKET:
		{
			p.SetState(36)
			p.Match(MonkeyParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.Expr_list()
		}
		{
			p.SetState(38)
			p.Match(MonkeyParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserIF:
		{
			p.SetState(40)
			p.Match(MonkeyParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.Match(MonkeyParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.expr(0)
		}
		{
			p.SetState(43)
			p.Match(MonkeyParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)
			p.Match(MonkeyParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268462014) != 0 {
			{
				p.SetState(45)
				p.Stmt()
			}

			p.SetState(50)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(51)
			p.Match(MonkeyParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(52)
				p.Match(MonkeyParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(53)
				p.Match(MonkeyParserLBRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268462014) != 0 {
				{
					p.SetState(54)
					p.Stmt()
				}

				p.SetState(59)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(60)
				p.Match(MonkeyParserRBRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case MonkeyParserFUNCTION:
		{
			p.SetState(63)
			p.Match(MonkeyParserFUNCTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Match(MonkeyParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Args_list()
		}
		{
			p.SetState(66)
			p.Match(MonkeyParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Match(MonkeyParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268462014) != 0 {
			{
				p.SetState(68)
				p.Stmt()
			}

			p.SetState(73)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(74)
			p.Match(MonkeyParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(78)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(79)
					_la = p.GetTokenStream().LA(1)

					if !(_la == MonkeyParserASTERISK || _la == MonkeyParserSLASH) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(80)
					p.expr(8)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(82)
					_la = p.GetTokenStream().LA(1)

					if !(_la == MonkeyParserPLUS || _la == MonkeyParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(83)
					p.expr(7)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(85)
					_la = p.GetTokenStream().LA(1)

					if !(_la == MonkeyParserLT || _la == MonkeyParserGT) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(86)
					p.expr(6)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(88)
					_la = p.GetTokenStream().LA(1)

					if !(_la == MonkeyParserEQ || _la == MonkeyParserNOT_EQ) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(89)
					p.expr(5)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(91)
					p.Match(MonkeyParserLBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(92)
					p.expr(0)
				}
				{
					p.SetState(93)
					p.Match(MonkeyParserRBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(95)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(96)
					p.Match(MonkeyParserLPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(97)
					p.Expr_list()
				}
				{
					p.SetState(98)
					p.Match(MonkeyParserRPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILet_stmtContext is an interface to support dynamic dispatch.
type ILet_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LET() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext
	SEMICOLON() antlr.TerminalNode

	// IsLet_stmtContext differentiates from other interfaces.
	IsLet_stmtContext()
}

type Let_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLet_stmtContext() *Let_stmtContext {
	var p = new(Let_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_let_stmt
	return p
}

func InitEmptyLet_stmtContext(p *Let_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_let_stmt
}

func (*Let_stmtContext) IsLet_stmtContext() {}

func NewLet_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Let_stmtContext {
	var p = new(Let_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_let_stmt

	return p
}

func (s *Let_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Let_stmtContext) LET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLET, 0)
}

func (s *Let_stmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIDENT, 0)
}

func (s *Let_stmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserASSIGN, 0)
}

func (s *Let_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Let_stmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSEMICOLON, 0)
}

func (s *Let_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Let_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Let_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterLet_stmt(s)
	}
}

func (s *Let_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitLet_stmt(s)
	}
}

func (s *Let_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitLet_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Let_stmt() (localctx ILet_stmtContext) {
	localctx = NewLet_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MonkeyParserRULE_let_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(MonkeyParserLET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(MonkeyParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(MonkeyParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
		p.expr(0)
	}
	{
		p.SetState(109)
		p.Match(MonkeyParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturn_stmtContext is an interface to support dynamic dispatch.
type IReturn_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expr() IExprContext
	SEMICOLON() antlr.TerminalNode

	// IsReturn_stmtContext differentiates from other interfaces.
	IsReturn_stmtContext()
}

type Return_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_stmtContext() *Return_stmtContext {
	var p = new(Return_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_return_stmt
	return p
}

func InitEmptyReturn_stmtContext(p *Return_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_return_stmt
}

func (*Return_stmtContext) IsReturn_stmtContext() {}

func NewReturn_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_stmtContext {
	var p = new(Return_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_return_stmt

	return p
}

func (s *Return_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_stmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRETURN, 0)
}

func (s *Return_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Return_stmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSEMICOLON, 0)
}

func (s *Return_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterReturn_stmt(s)
	}
}

func (s *Return_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitReturn_stmt(s)
	}
}

func (s *Return_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitReturn_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Return_stmt() (localctx IReturn_stmtContext) {
	localctx = NewReturn_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MonkeyParserRULE_return_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(MonkeyParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.expr(0)
	}
	{
		p.SetState(113)
		p.Match(MonkeyParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpression_stmtContext is an interface to support dynamic dispatch.
type IExpression_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	SEMICOLON() antlr.TerminalNode

	// IsExpression_stmtContext differentiates from other interfaces.
	IsExpression_stmtContext()
}

type Expression_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_stmtContext() *Expression_stmtContext {
	var p = new(Expression_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_expression_stmt
	return p
}

func InitEmptyExpression_stmtContext(p *Expression_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_expression_stmt
}

func (*Expression_stmtContext) IsExpression_stmtContext() {}

func NewExpression_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_stmtContext {
	var p = new(Expression_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_expression_stmt

	return p
}

func (s *Expression_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Expression_stmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSEMICOLON, 0)
}

func (s *Expression_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterExpression_stmt(s)
	}
}

func (s *Expression_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitExpression_stmt(s)
	}
}

func (s *Expression_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitExpression_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Expression_stmt() (localctx IExpression_stmtContext) {
	localctx = NewExpression_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MonkeyParserRULE_expression_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.expr(0)
	}
	{
		p.SetState(116)
		p.Match(MonkeyParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpr_listContext is an interface to support dynamic dispatch.
type IExpr_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExpr_listContext differentiates from other interfaces.
	IsExpr_listContext()
}

type Expr_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_listContext() *Expr_listContext {
	var p = new(Expr_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_expr_list
	return p
}

func InitEmptyExpr_listContext(p *Expr_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_expr_list
}

func (*Expr_listContext) IsExpr_listContext() {}

func NewExpr_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_listContext {
	var p = new(Expr_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_expr_list

	return p
}

func (s *Expr_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_listContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Expr_listContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Expr_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserCOMMA)
}

func (s *Expr_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserCOMMA, i)
}

func (s *Expr_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterExpr_list(s)
	}
}

func (s *Expr_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitExpr_list(s)
	}
}

func (s *Expr_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitExpr_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Expr_list() (localctx IExpr_listContext) {
	localctx = NewExpr_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MonkeyParserRULE_expr_list)
	var _la int

	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserRPAREN, MonkeyParserRBRACKET:
		p.EnterOuterAlt(localctx, 1)

	case MonkeyParserFUNCTION, MonkeyParserTRUE, MonkeyParserFALSE, MonkeyParserIF, MonkeyParserSTRING, MonkeyParserIDENT, MonkeyParserINT, MonkeyParserMINUS, MonkeyParserBANG, MonkeyParserLBRACKET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(119)
			p.expr(0)
		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MonkeyParserCOMMA {
			{
				p.SetState(120)
				p.Match(MonkeyParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(121)
				p.expr(0)
			}

			p.SetState(126)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgs_listContext is an interface to support dynamic dispatch.
type IArgs_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgs_listContext differentiates from other interfaces.
	IsArgs_listContext()
}

type Args_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgs_listContext() *Args_listContext {
	var p = new(Args_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_args_list
	return p
}

func InitEmptyArgs_listContext(p *Args_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_args_list
}

func (*Args_listContext) IsArgs_listContext() {}

func NewArgs_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Args_listContext {
	var p = new(Args_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_args_list

	return p
}

func (s *Args_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Args_listContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserIDENT)
}

func (s *Args_listContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserIDENT, i)
}

func (s *Args_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserCOMMA)
}

func (s *Args_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserCOMMA, i)
}

func (s *Args_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Args_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Args_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterArgs_list(s)
	}
}

func (s *Args_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitArgs_list(s)
	}
}

func (s *Args_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitArgs_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Args_list() (localctx IArgs_listContext) {
	localctx = NewArgs_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MonkeyParserRULE_args_list)
	var _la int

	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserRPAREN:
		p.EnterOuterAlt(localctx, 1)

	case MonkeyParserIDENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			p.Match(MonkeyParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MonkeyParserCOMMA {
			{
				p.SetState(131)
				p.Match(MonkeyParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(132)
				p.Match(MonkeyParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(137)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *MonkeyParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MonkeyParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
