grammar Monkey;

// -- Parser ---------------------------------------

prog : stmt* expr? EOF ;

stmt : let_stmt        # LetStatement
     | return_stmt     # ReturnStatement
     | expression_stmt # ExpressionStatement
     ;

expr : INT                                                                   # IntegerLiteral
     | (TRUE | FALSE)                                                        # BooleanLiteral
     | STRING                                                                # StringLiteral
     | IDENT                                                                 # IdentifierExpression
     | op=(MINUS | BANG) expr                                                # UnaryOperatorExpression
     | expr LBRACKET expr RBRACKET                                           # IndexOperatorExpression
     | expr LPAREN expr_list RPAREN                                          # CallExpression
     | expr (ASTERISK | SLASH) expr                                          # MulDivBinaryExpression
     | expr (PLUS | MINUS) expr                                              # AddSubBinaryExpression
     | expr (LT | GT) expr                                                   # LesGreBinaryExpression
     | expr (EQ | NOT_EQ) expr                                               # EqualityBinaryExpression
     | LBRACKET expr_list RBRACKET                                           # ArrayExpression
     | IF LPAREN expr RPAREN LBRACE stmt* RBRACE (ELSE LBRACE stmt* RBRACE)? # IfExpression
     | FUNCTION LPAREN args_list RPAREN LBRACE stmt* RBRACE                  # FunctionExpression
     ;

// Statements
let_stmt : LET IDENT ASSIGN expr SEMICOLON ;

return_stmt : RETURN expr SEMICOLON ;

expression_stmt : expr SEMICOLON ;

// Expressions
expr_list :
          | expr (COMMA expr)*
          ;

args_list :
          | IDENT (COMMA IDENT)*
          ;

// -- Lexer ----------------------------------------

// Keywords
FUNCTION : 'fn' ;
LET      : 'let' ;
TRUE     : 'true' ;
FALSE    : 'false' ;
IF       : 'if' ;
ELSE     : 'else' ;
RETURN   : 'return' ;

// Identifiers
STRING : '"' (~'"')+ '"' ;
IDENT  : [a-zA-Z_]+ ;
INT    : [0-9]+ ;

// Operators
ASSIGN   : '=' ;
PLUS     : '+' ;
MINUS    : '-' ;
BANG     : '!' ;
ASTERISK : '*' ;
SLASH    : '/' ;

LT : '<' ;
GT : '>' ;

EQ     : '==' ;
NOT_EQ : '!=' ;

// Delimeters
COMMA     : ',' ;
COLON     : ':' ;
SEMICOLON : ';' ;

LPAREN   : '(' ;
RPAREN   : ')' ;
LBRACE   : '{' ;
RBRACE   : '}' ;
LBRACKET : '[' ;
RBRACKET : ']' ;

WS : [ \t\r\n] -> skip ;