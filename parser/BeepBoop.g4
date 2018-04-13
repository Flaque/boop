grammar BeepBoop ;

beepboop : block EOF ;

block : statement+ ;

statement : expr NEWLINE;

expr 
    : term                         #termExpr
    | MINUS expr                   #unaryMinusExpr
    | expr op=(PLUS | MINUS) expr  #additiveExpr
    ;


term : INT ;

NEWLINE : [\r\n]+ ;
WHITESPACE : ' ' -> skip ;
INT : [0-9]+ ;
PLUS : '+';
MINUS : '-';