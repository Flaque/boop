grammar BeepBoop;

boop: NEWLINE+ funcdef+ EOF | NEWLINE+ statement+ EOF;

funcdef:
	FUNC label DO END
	| FUNC label+ DO END
	| FUNC label DO statement+ END
	| FUNC label+ DO statement+ END;

statement:
	assignstat
	| fncall NEWLINE
	| ifstat NEWLINE
	| pipe NEWLINE
	| NEWLINE;

ifstat:
	IF conditional DO statement+ END
	| IF conditional DO END;

assignstat: assign NEWLINE;

assign: label ASSIGN term | label ASSIGN fncall;

paren_fncall: LPAREN fncall RPAREN;

fncall: label term+ | label;

term:
	label			# labelTerm // Also a function call
	| literal		# literalTerm
	| math			# mathTerm
	| paren_fncall	# parenfncallTerm
	| struct		# structTerm
	| list			# listTerm;

struct: LSQUIG RSQUIG | LSQUIG assignstat+ RSQUIG;

list: LBLOCK RBLOCK | LBLOCK term+ RBLOCK;

conditional:
	term EQUALS term				# equalsCond
	| term LTHAN term				# lthanCond
	| term GTHAN term				# gthanCond
	| term LTHAN_EQUALS term		# lthanequalsCond
	| term GTHAN_EQUALS term		# gthanequalsCond
	| conditional OR conditional	# orCond
	| conditional AND conditional	# andCond
	| bool							# boolCond;

math:
	num op = (PLUS | SUB) num	# additiveMath
	| SUB num					# unarySubMath
	| num						# soloMath;

literal: num | STRING | bool;
num: INT;
bool: TRUE | FALSE;

pipe: term PIPE fncall | term PIPE pipe | fncall PIPE pipe;

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
PLUS: 'plus';
SUB: 'sub';
DIV: 'div';
MULT: 'mult';
TRUE: 'true';
FALSE: 'false';
OR: 'or';
AND: 'and';

label: ':' STRING;
EQUALS: '==';
ASSIGN: '=';
PIPE: '|';
LTHAN: '<';
GTHAN: '>';
LTHAN_EQUALS: '<=';
GTHAN_EQUALS: '>=';
LPAREN: '(';
RPAREN: ')';
LSQUIG: '{';
RSQUIG: '}';
LBLOCK: '[';
RBLOCK: ']';
INT: [0-9]+;
STRING: [a-zA-Z]+;
