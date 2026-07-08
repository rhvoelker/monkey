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
		4, 1, 30, 144, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 5, 0, 18, 8, 0, 10, 0, 12,
		0, 21, 9, 0, 1, 0, 3, 0, 24, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1,
		31, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 50, 8, 2, 10, 2, 12, 2, 53,
		9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 59, 8, 2, 10, 2, 12, 2, 62, 9, 2, 1,
		2, 3, 2, 65, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 73, 8, 2,
		10, 2, 12, 2, 76, 9, 2, 1, 2, 1, 2, 3, 2, 80, 8, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 104, 8, 2, 10, 2, 12, 2, 107,
		9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 126, 8, 6, 10, 6, 12, 6, 129,
		9, 6, 3, 6, 131, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 137, 8, 7, 10, 7,
		12, 7, 140, 9, 7, 3, 7, 142, 8, 7, 1, 7, 0, 1, 4, 8, 0, 2, 4, 6, 8, 10,
		12, 14, 0, 6, 1, 0, 3, 4, 1, 0, 13, 14, 1, 0, 15, 16, 1, 0, 12, 13, 1,
		0, 17, 18, 1, 0, 19, 20, 160, 0, 19, 1, 0, 0, 0, 2, 30, 1, 0, 0, 0, 4,
		79, 1, 0, 0, 0, 6, 108, 1, 0, 0, 0, 8, 114, 1, 0, 0, 0, 10, 118, 1, 0,
		0, 0, 12, 130, 1, 0, 0, 0, 14, 141, 1, 0, 0, 0, 16, 18, 3, 2, 1, 0, 17,
		16, 1, 0, 0, 0, 18, 21, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0,
		0, 20, 23, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 22, 24, 3, 4, 2, 0, 23, 22,
		1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 26, 5, 0, 0, 1,
		26, 1, 1, 0, 0, 0, 27, 31, 3, 6, 3, 0, 28, 31, 3, 8, 4, 0, 29, 31, 3, 10,
		5, 0, 30, 27, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30, 29, 1, 0, 0, 0, 31, 3,
		1, 0, 0, 0, 32, 33, 6, 2, -1, 0, 33, 80, 5, 10, 0, 0, 34, 80, 7, 0, 0,
		0, 35, 80, 5, 8, 0, 0, 36, 80, 5, 9, 0, 0, 37, 38, 7, 1, 0, 0, 38, 80,
		3, 4, 2, 10, 39, 40, 5, 28, 0, 0, 40, 41, 3, 12, 6, 0, 41, 42, 5, 29, 0,
		0, 42, 80, 1, 0, 0, 0, 43, 44, 5, 5, 0, 0, 44, 45, 5, 24, 0, 0, 45, 46,
		3, 4, 2, 0, 46, 47, 5, 25, 0, 0, 47, 51, 5, 26, 0, 0, 48, 50, 3, 2, 1,
		0, 49, 48, 1, 0, 0, 0, 50, 53, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 51, 52,
		1, 0, 0, 0, 52, 54, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 54, 64, 5, 27, 0, 0,
		55, 56, 5, 6, 0, 0, 56, 60, 5, 26, 0, 0, 57, 59, 3, 2, 1, 0, 58, 57, 1,
		0, 0, 0, 59, 62, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61,
		63, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 63, 65, 5, 27, 0, 0, 64, 55, 1, 0,
		0, 0, 64, 65, 1, 0, 0, 0, 65, 80, 1, 0, 0, 0, 66, 67, 5, 1, 0, 0, 67, 68,
		5, 24, 0, 0, 68, 69, 3, 14, 7, 0, 69, 70, 5, 25, 0, 0, 70, 74, 5, 26, 0,
		0, 71, 73, 3, 2, 1, 0, 72, 71, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74, 72,
		1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 77, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0,
		77, 78, 5, 27, 0, 0, 78, 80, 1, 0, 0, 0, 79, 32, 1, 0, 0, 0, 79, 34, 1,
		0, 0, 0, 79, 35, 1, 0, 0, 0, 79, 36, 1, 0, 0, 0, 79, 37, 1, 0, 0, 0, 79,
		39, 1, 0, 0, 0, 79, 43, 1, 0, 0, 0, 79, 66, 1, 0, 0, 0, 80, 105, 1, 0,
		0, 0, 81, 82, 10, 7, 0, 0, 82, 83, 7, 2, 0, 0, 83, 104, 3, 4, 2, 8, 84,
		85, 10, 6, 0, 0, 85, 86, 7, 3, 0, 0, 86, 104, 3, 4, 2, 7, 87, 88, 10, 5,
		0, 0, 88, 89, 7, 4, 0, 0, 89, 104, 3, 4, 2, 6, 90, 91, 10, 4, 0, 0, 91,
		92, 7, 5, 0, 0, 92, 104, 3, 4, 2, 5, 93, 94, 10, 9, 0, 0, 94, 95, 5, 28,
		0, 0, 95, 96, 3, 4, 2, 0, 96, 97, 5, 29, 0, 0, 97, 104, 1, 0, 0, 0, 98,
		99, 10, 8, 0, 0, 99, 100, 5, 24, 0, 0, 100, 101, 3, 12, 6, 0, 101, 102,
		5, 25, 0, 0, 102, 104, 1, 0, 0, 0, 103, 81, 1, 0, 0, 0, 103, 84, 1, 0,
		0, 0, 103, 87, 1, 0, 0, 0, 103, 90, 1, 0, 0, 0, 103, 93, 1, 0, 0, 0, 103,
		98, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105, 106, 1,
		0, 0, 0, 106, 5, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108, 109, 5, 2, 0, 0,
		109, 110, 5, 9, 0, 0, 110, 111, 5, 11, 0, 0, 111, 112, 3, 4, 2, 0, 112,
		113, 5, 23, 0, 0, 113, 7, 1, 0, 0, 0, 114, 115, 5, 7, 0, 0, 115, 116, 3,
		4, 2, 0, 116, 117, 5, 23, 0, 0, 117, 9, 1, 0, 0, 0, 118, 119, 3, 4, 2,
		0, 119, 120, 5, 23, 0, 0, 120, 11, 1, 0, 0, 0, 121, 131, 1, 0, 0, 0, 122,
		127, 3, 4, 2, 0, 123, 124, 5, 21, 0, 0, 124, 126, 3, 4, 2, 0, 125, 123,
		1, 0, 0, 0, 126, 129, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 128, 1, 0,
		0, 0, 128, 131, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 130, 121, 1, 0, 0, 0,
		130, 122, 1, 0, 0, 0, 131, 13, 1, 0, 0, 0, 132, 142, 1, 0, 0, 0, 133, 138,
		5, 9, 0, 0, 134, 135, 5, 21, 0, 0, 135, 137, 5, 9, 0, 0, 136, 134, 1, 0,
		0, 0, 137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0,
		139, 142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 132, 1, 0, 0, 0, 141,
		133, 1, 0, 0, 0, 142, 15, 1, 0, 0, 0, 14, 19, 23, 30, 51, 60, 64, 74, 79,
		103, 105, 127, 130, 138, 141,
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
	Expr() IExprContext

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

func (s *ProgContext) Expr() IExprContext {
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(16)
				p.Stmt()
			}

		}
		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268461882) != 0 {
		{
			p.SetState(22)
			p.expr(0)
		}

	}
	{
		p.SetState(25)
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

func (s *StmtContext) CopyAll(ctx *StmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionStatementContext struct {
	StmtContext
}

func NewExpressionStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) Expression_stmt() IExpression_stmtContext {
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

func (s *ExpressionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitExpressionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStatementContext struct {
	StmtContext
}

func NewReturnStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) Return_stmt() IReturn_stmtContext {
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

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type LetStatementContext struct {
	StmtContext
}

func NewLetStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetStatementContext {
	var p = new(LetStatementContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *LetStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetStatementContext) Let_stmt() ILet_stmtContext {
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

func (s *LetStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterLetStatement(s)
	}
}

func (s *LetStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitLetStatement(s)
	}
}

func (s *LetStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitLetStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MonkeyParserRULE_stmt)
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserLET:
		localctx = NewLetStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)
			p.Let_stmt()
		}

	case MonkeyParserRETURN:
		localctx = NewReturnStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(28)
			p.Return_stmt()
		}

	case MonkeyParserFUNCTION, MonkeyParserTRUE, MonkeyParserFALSE, MonkeyParserIF, MonkeyParserSTRING, MonkeyParserIDENT, MonkeyParserINT, MonkeyParserMINUS, MonkeyParserBANG, MonkeyParserLBRACKET:
		localctx = NewExpressionStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(29)
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

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfExpressionContext struct {
	ExprContext
}

func NewIfExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfExpressionContext {
	var p = new(IfExpressionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IfExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExpressionContext) IF() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIF, 0)
}

func (s *IfExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLPAREN, 0)
}

func (s *IfExpressionContext) Expr() IExprContext {
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

func (s *IfExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRPAREN, 0)
}

func (s *IfExpressionContext) AllLBRACE() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserLBRACE)
}

func (s *IfExpressionContext) LBRACE(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserLBRACE, i)
}

func (s *IfExpressionContext) AllRBRACE() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserRBRACE)
}

func (s *IfExpressionContext) RBRACE(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserRBRACE, i)
}

func (s *IfExpressionContext) AllStmt() []IStmtContext {
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

func (s *IfExpressionContext) Stmt(i int) IStmtContext {
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

func (s *IfExpressionContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserELSE, 0)
}

func (s *IfExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterIfExpression(s)
	}
}

func (s *IfExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitIfExpression(s)
	}
}

func (s *IfExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitIfExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryOperatorExpressionContext struct {
	ExprContext
}

func NewUnaryOperatorExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryOperatorExpressionContext {
	var p = new(UnaryOperatorExpressionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *UnaryOperatorExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperatorExpressionContext) Expr() IExprContext {
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

func (s *UnaryOperatorExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MonkeyParserMINUS, 0)
}

func (s *UnaryOperatorExpressionContext) BANG() antlr.TerminalNode {
	return s.GetToken(MonkeyParserBANG, 0)
}

func (s *UnaryOperatorExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterUnaryOperatorExpression(s)
	}
}

func (s *UnaryOperatorExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitUnaryOperatorExpression(s)
	}
}

func (s *UnaryOperatorExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitUnaryOperatorExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanLiteralContext struct {
	ExprContext
}

func NewBooleanLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserFALSE, 0)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitBooleanLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualityBinaryExpressionContext struct {
	ExprContext
}

func NewEqualityBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityBinaryExpressionContext {
	var p = new(EqualityBinaryExpressionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EqualityBinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityBinaryExpressionContext) AllExpr() []IExprContext {
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

func (s *EqualityBinaryExpressionContext) Expr(i int) IExprContext {
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

func (s *EqualityBinaryExpressionContext) EQ() antlr.TerminalNode {
	return s.GetToken(MonkeyParserEQ, 0)
}

func (s *EqualityBinaryExpressionContext) NOT_EQ() antlr.TerminalNode {
	return s.GetToken(MonkeyParserNOT_EQ, 0)
}

func (s *EqualityBinaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterEqualityBinaryExpression(s)
	}
}

func (s *EqualityBinaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitEqualityBinaryExpression(s)
	}
}

func (s *EqualityBinaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitEqualityBinaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierExpressionContext struct {
	ExprContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIDENT, 0)
}

func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitIdentifierExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type LesGreBinaryExpressionContext struct {
	ExprContext
}

func NewLesGreBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LesGreBinaryExpressionContext {
	var p = new(LesGreBinaryExpressionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LesGreBinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LesGreBinaryExpressionContext) AllExpr() []IExprContext {
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

func (s *LesGreBinaryExpressionContext) Expr(i int) IExprContext {
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

func (s *LesGreBinaryExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLT, 0)
}

func (s *LesGreBinaryExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserGT, 0)
}

func (s *LesGreBinaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterLesGreBinaryExpression(s)
	}
}

func (s *LesGreBinaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitLesGreBinaryExpression(s)
	}
}

func (s *LesGreBinaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitLesGreBinaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexOperatorExpressionContext struct {
	ExprContext
}

func NewIndexOperatorExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexOperatorExpressionContext {
	var p = new(IndexOperatorExpressionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IndexOperatorExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexOperatorExpressionContext) AllExpr() []IExprContext {
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

func (s *IndexOperatorExpressionContext) Expr(i int) IExprContext {
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

func (s *IndexOperatorExpressionContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLBRACKET, 0)
}

func (s *IndexOperatorExpressionContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRBRACKET, 0)
}

func (s *IndexOperatorExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterIndexOperatorExpression(s)
	}
}

func (s *IndexOperatorExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitIndexOperatorExpression(s)
	}
}

func (s *IndexOperatorExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitIndexOperatorExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubBinaryExpressionContext struct {
	ExprContext
}

func NewAddSubBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubBinaryExpressionContext {
	var p = new(AddSubBinaryExpressionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AddSubBinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubBinaryExpressionContext) AllExpr() []IExprContext {
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

func (s *AddSubBinaryExpressionContext) Expr(i int) IExprContext {
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

func (s *AddSubBinaryExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPLUS, 0)
}

func (s *AddSubBinaryExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MonkeyParserMINUS, 0)
}

func (s *AddSubBinaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterAddSubBinaryExpression(s)
	}
}

func (s *AddSubBinaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitAddSubBinaryExpression(s)
	}
}

func (s *AddSubBinaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitAddSubBinaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionExpressionContext struct {
	ExprContext
}

func NewFunctionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExpressionContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(MonkeyParserFUNCTION, 0)
}

func (s *FunctionExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLPAREN, 0)
}

func (s *FunctionExpressionContext) Args_list() IArgs_listContext {
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

func (s *FunctionExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRPAREN, 0)
}

func (s *FunctionExpressionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLBRACE, 0)
}

func (s *FunctionExpressionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRBRACE, 0)
}

func (s *FunctionExpressionContext) AllStmt() []IStmtContext {
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

func (s *FunctionExpressionContext) Stmt(i int) IStmtContext {
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

func (s *FunctionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitFunctionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringLiteralContext struct {
	ExprContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSTRING, 0)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayExpressionContext struct {
	ExprContext
}

func NewArrayExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayExpressionContext {
	var p = new(ArrayExpressionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArrayExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayExpressionContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLBRACKET, 0)
}

func (s *ArrayExpressionContext) Expr_list() IExpr_listContext {
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

func (s *ArrayExpressionContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRBRACKET, 0)
}

func (s *ArrayExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterArrayExpression(s)
	}
}

func (s *ArrayExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitArrayExpression(s)
	}
}

func (s *ArrayExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitArrayExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallExpressionContext struct {
	ExprContext
}

func NewCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExpressionContext {
	var p = new(CallExpressionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExpressionContext) Expr() IExprContext {
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

func (s *CallExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLPAREN, 0)
}

func (s *CallExpressionContext) Expr_list() IExpr_listContext {
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

func (s *CallExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRPAREN, 0)
}

func (s *CallExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterCallExpression(s)
	}
}

func (s *CallExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitCallExpression(s)
	}
}

func (s *CallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntegerLiteralContext struct {
	ExprContext
}

func NewIntegerLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserINT, 0)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitIntegerLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivBinaryExpressionContext struct {
	ExprContext
}

func NewMulDivBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivBinaryExpressionContext {
	var p = new(MulDivBinaryExpressionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MulDivBinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivBinaryExpressionContext) AllExpr() []IExprContext {
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

func (s *MulDivBinaryExpressionContext) Expr(i int) IExprContext {
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

func (s *MulDivBinaryExpressionContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(MonkeyParserASTERISK, 0)
}

func (s *MulDivBinaryExpressionContext) SLASH() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSLASH, 0)
}

func (s *MulDivBinaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterMulDivBinaryExpression(s)
	}
}

func (s *MulDivBinaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitMulDivBinaryExpression(s)
	}
}

func (s *MulDivBinaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyVisitor:
		return t.VisitMulDivBinaryExpression(s)

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
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserINT:
		localctx = NewIntegerLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(33)
			p.Match(MonkeyParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserTRUE, MonkeyParserFALSE:
		localctx = NewBooleanLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(34)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MonkeyParserTRUE || _la == MonkeyParserFALSE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case MonkeyParserSTRING:
		localctx = NewStringLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(35)
			p.Match(MonkeyParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserIDENT:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(36)
			p.Match(MonkeyParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserMINUS, MonkeyParserBANG:
		localctx = NewUnaryOperatorExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(37)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MonkeyParserMINUS || _la == MonkeyParserBANG) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(38)
			p.expr(10)
		}

	case MonkeyParserLBRACKET:
		localctx = NewArrayExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(39)
			p.Match(MonkeyParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.Expr_list()
		}
		{
			p.SetState(41)
			p.Match(MonkeyParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserIF:
		localctx = NewIfExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(43)
			p.Match(MonkeyParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)
			p.Match(MonkeyParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.expr(0)
		}
		{
			p.SetState(46)
			p.Match(MonkeyParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Match(MonkeyParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268462014) != 0 {
			{
				p.SetState(48)
				p.Stmt()
			}

			p.SetState(53)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(54)
			p.Match(MonkeyParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(55)
				p.Match(MonkeyParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(56)
				p.Match(MonkeyParserLBRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(60)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268462014) != 0 {
				{
					p.SetState(57)
					p.Stmt()
				}

				p.SetState(62)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(63)
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
		localctx = NewFunctionExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(66)
			p.Match(MonkeyParserFUNCTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Match(MonkeyParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Args_list()
		}
		{
			p.SetState(69)
			p.Match(MonkeyParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(MonkeyParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268462014) != 0 {
			{
				p.SetState(71)
				p.Stmt()
			}

			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(77)
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
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(103)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivBinaryExpressionContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(82)
					_la = p.GetTokenStream().LA(1)

					if !(_la == MonkeyParserASTERISK || _la == MonkeyParserSLASH) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(83)
					p.expr(8)
				}

			case 2:
				localctx = NewAddSubBinaryExpressionContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(85)
					_la = p.GetTokenStream().LA(1)

					if !(_la == MonkeyParserPLUS || _la == MonkeyParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(86)
					p.expr(7)
				}

			case 3:
				localctx = NewLesGreBinaryExpressionContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(88)
					_la = p.GetTokenStream().LA(1)

					if !(_la == MonkeyParserLT || _la == MonkeyParserGT) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(89)
					p.expr(6)
				}

			case 4:
				localctx = NewEqualityBinaryExpressionContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(91)
					_la = p.GetTokenStream().LA(1)

					if !(_la == MonkeyParserEQ || _la == MonkeyParserNOT_EQ) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(92)
					p.expr(5)
				}

			case 5:
				localctx = NewIndexOperatorExpressionContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(94)
					p.Match(MonkeyParserLBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(95)
					p.expr(0)
				}
				{
					p.SetState(96)
					p.Match(MonkeyParserRBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 6:
				localctx = NewCallExpressionContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(98)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(99)
					p.Match(MonkeyParserLPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(100)
					p.Expr_list()
				}
				{
					p.SetState(101)
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
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
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
		p.SetState(108)
		p.Match(MonkeyParserLET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(MonkeyParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Match(MonkeyParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.expr(0)
	}
	{
		p.SetState(112)
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
		p.SetState(114)
		p.Match(MonkeyParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
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
		p.SetState(118)
		p.expr(0)
	}
	{
		p.SetState(119)
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

	p.SetState(130)
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
			p.SetState(122)
			p.expr(0)
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MonkeyParserCOMMA {
			{
				p.SetState(123)
				p.Match(MonkeyParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(124)
				p.expr(0)
			}

			p.SetState(129)
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

	p.SetState(141)
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
			p.SetState(133)
			p.Match(MonkeyParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MonkeyParserCOMMA {
			{
				p.SetState(134)
				p.Match(MonkeyParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(135)
				p.Match(MonkeyParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(140)
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
