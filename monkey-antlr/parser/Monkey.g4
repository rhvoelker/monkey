grammar Monkey;

// -- Parser ---------------------------------------

prog : stmt+ EOF ;
stmt : let_stmt
     | return_stmt
     | expression_stmt
     ;
expr : INT
     | TRUE
     | FALSE
     | STRING
     | IDENT
     | (MINUS | BANG) expr
     | expr LBRACKET expr RBRACKET
     | expr LPAREN expr_list RPAREN
     | expr (ASTERISK | SLASH) expr
     | expr (PLUS | MINUS) expr
     | expr (LT | GT) expr
     | expr (EQ | NOT_EQ) expr
     | LBRACKET expr_list RBRACKET
     | IF LPAREN expr RPAREN LBRACE stmt* RBRACE (ELSE LBRACE stmt* RBRACE)?
     | FUNCTION LPAREN args_list RPAREN LBRACE stmt* RBRACE
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