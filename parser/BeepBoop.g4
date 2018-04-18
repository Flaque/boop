grammar BeepBoop;

beepboop: NEWLINE+ block EOF | block EOF;

block: statement+;

statement:
	assignment NEWLINE		# assignStatement
	| fncall NEWLINE		# fncallStatement
	| returnStat NEWLINE	# returnStatement
	| pipe NEWLINE			# pipeStatement;

ifstat:
	IF expr DO block END		# exprIfStatement
	| IF fncall DO block END	# fncallIfStatement;

funcdef: FUNC STRING label+ DO block END;

returnStat:
	RETURN expr		# exprReturn
	| RETURN fncall	# fncallReturn;

assignment: label ASSIGN expr | label ASSIGN fncall;

expr:
	expr op = (PLUS | MINUS) expr	# additiveExpr
	| MINUS expr					# unaryMinusExpr
	| term							# termExpr;

pipe: fncall PIPE fncall | fncall NEWLINE;

fncall: STRING | STRING term+;

term: label # labelTerm | STRING # stringTerm | INT # intTerm;

label: '$' STRING;

COMMENT: '#' ~[\r\n]* -> skip;
NEWLINE: [\r\n]+;
WHITESPACE: ' ' -> channel(HIDDEN);
TABSPACE: [\t]+ -> channel(HIDDEN);
INT: [0-9]+;
PLUS: '+';
MINUS: '-';
MULT: '*';
DIVIDE: '/';
ASSIGN: '=';
PIPE: '|';
LPAREN: '(';
RPAREN: ')';
STRING: [a-zA-Z]+;
IF: 'if';
DO: 'do';
END: 'end';
FUNC: 'func';
RETURN: 'return';
FOR: 'for';
IS: 'is';

UNKNOWN: . -> skip;