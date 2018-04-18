grammar BeepBoop;

beepboop: block EOF;

block: statement+;

statement:
	assignment NEWLINE		# assignStatement
	| fncall NEWLINE		# fncallStatement
	| returnStat NEWLINE	# returnStatement;

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

fncall: STRING | STRING term+;

term: label # labelTerm | STRING # stringTerm | INT # intTerm;

label: '$' STRING;

NEWLINE: [\r\n]+;
WHITESPACE: ' ' -> skip;
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