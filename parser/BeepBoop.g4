grammar BeepBoop ;

beepboop : block EOF ;

block : statement+ ;

statement 
    : assignment NEWLINE #assignStatement
    | fncall NEWLINE     #fncallStatement
    ;

assignment
    : label ASSIGN expr   
    | label ASSIGN fncall
    ;

expr 
    : expr op=(PLUS | MINUS) expr  #additiveExpr
    | MINUS expr                   #unaryMinusExpr
    | term                         #termExpr
    ;

fncall
    : STRING       
    | STRING term+ 
    ;

term 
   : label  #labelTerm
   | STRING #stringTerm
   | INT    #intTerm
   ;

label : '$'STRING ; 

NEWLINE : [\r\n]+ ;
WHITESPACE : ' ' -> skip ;
INT : [0-9]+ ;
PLUS : '+' ;
MINUS : '-' ;
ASSIGN : '=' ;
STRING : [a-zA-Z]+ ;