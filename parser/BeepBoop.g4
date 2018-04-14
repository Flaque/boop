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
    : term                         #termExpr
    | MINUS expr                   #unaryMinusExpr
    | expr op=(PLUS | MINUS) expr  #additiveExpr
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