grammar BeepBoop ;

beepboop : block EOF ;

block : statement+ ;

statement 
    : expr NEWLINE       #exprStatement
    | assignment NEWLINE #assignStatement
    ;

assignment
    : label ASSIGN expr 
    ;

expr 
    : term                         #termExpr
    | MINUS expr                   #unaryMinusExpr
    | expr op=(PLUS | MINUS) expr  #additiveExpr
    ;


term : INT ;
label : '$'STRING ; 

NEWLINE : [\r\n]+ ;
WHITESPACE : ' ' -> skip ;
INT : [0-9]+ ;
PLUS : '+' ;
MINUS : '-' ;
ASSIGN : '=' ;
STRING : [a-zA-Z] ;