grammar BeepBoop ;

beepboop : expr ;

expr 
    : term             #termExpr
    | expr PLUS expr   #addExpr
    | expr MINUS expr  #minusExpr
    ;


term : INT ;

NEWLINE : [\r\n]+ ;
WHITESPACE : ' ' -> skip ;
INT : [0-9]+ ;
PLUS : '+';
MINUS : '-';