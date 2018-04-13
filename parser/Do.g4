grammar Do ;

do : math ;

math : term | math '+' term | math '-' term ;
term : INT ;

NEWLINE : [\r\n]+ ;
WHITESPACE : ' ' -> skip ;
INT : [0-9]+ ;