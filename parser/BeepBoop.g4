grammar BeepBoop;

beepboop: NEWLINE+ code+ EOF | code+ EOF;

code: statement # statementCode | funcdef # funcdefCode;

statement:
	fncall NEWLINE			# fncallStatement
	| assignment NEWLINE	# assignStatement
	| returnStat NEWLINE	# returnStatement
	| ifstat NEWLINE		# ifStatement
	| pipe NEWLINE			# pipeStatement
	| NEWLINE				# noopStatement;

funcguts: statement+;

funcdef:
	FUNC STRING label+ DO funcguts END NEWLINE
	| FUNC STRING DO funcguts END NEWLINE
	| FUNC STRING DO END NEWLINE;

fncall: STRING | STRING term+;

ifstat:
	IF expr DO code END		# exprIfStatement
	| IF fncall DO code END	# fncallIfStatement;

returnStat:
	RETURN expr		# exprReturn
	| RETURN fncall	# fncallReturn;

assignment: label ASSIGN expr | label ASSIGN fncall;

expr:
	expr op = (PLUS | MINUS) expr	# additiveExpr
	| MINUS expr					# unaryMinusExpr
	| term							# termExpr;

pipe: fncall PIPE fncall | fncall NEWLINE+;

term: label # labelTerm | STRING # stringTerm | INT # intTerm;

// Comments and whitespace
COMMENT: '#' ~[\r\n]* -> skip;
NEWLINE: [\r\n] | [\r\n]+;
WHITESPACE: (' ' | '\t')+ -> channel(HIDDEN);

// Keywords
IF: 'if';
DO: 'do';
END: 'end';
FUNC: 'func';
RETURN: 'return';
FOR: 'for';
IS: 'is';

label: '$' STRING;
PLUS: '+';
MINUS: '-';
MULT: '*';
DIVIDE: '/';
ASSIGN: '=';
PIPE: '|';
LPAREN: '(';
RPAREN: ')';
STRING: [a-zA-Z]+;
INT: [0-9]+;
