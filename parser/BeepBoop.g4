grammar BeepBoop;

boop: (NEWLINE | NEWLINE+) (code | code+) EOF
	| (code | code+) EOF;

code: funcdef # funcdefCode | statement # statementCode;

funcdef:
	FUNC LABEL DO END
	| FUNC LABEL+ DO END
	| FUNC LABEL DO funcguts END
	| FUNC LABEL+ DO funcguts END;

funcguts: statement+;

statement:
	assignstat			# assignStatement
	| returnstat		# returnStatement
	| fncall NEWLINE	# fncallStatement
	| ifstat NEWLINE	# ifStatement
	| pipe NEWLINE		# pipeStatement
	| NEWLINE			# newlineStatement;

ifstat:
	IF conditional DO statement+ END
	| IF conditional DO END;

returnstat:
	RETURN term
	| RETURN term (NEWLINE | NEWLINE+)
	| RETURN (NEWLINE | NEWLINE+);

structexpr:
	LSQUIG RSQUIG
	| LSQUIG NEWLINE (assignstat | assignstat+) RSQUIG
	| LSQUIG (assignstat | assignstat+) RSQUIG
	| LSQUIG assign RSQUIG
	| LSQUIG NEWLINE assign RSQUIG;

assignstat: assign NEWLINE;

assign:
	LABEL ASSIGN term		# exprAssign
	| LABEL ASSIGN fncall	# fncallAssign;

paren_fncall: LPAREN fncall RPAREN;

fncall: LABEL term+ | LABEL;

term:
	LABEL			# labelTerm // Also a function call
	| literal		# literalTerm
	| math			# mathTerm
	| paren_fncall	# parenfncallTerm
	| structexpr	# structTerm
	| list			# listTerm;

list:
	LBLOCK RBLOCK
	| LBLOCK listterm+ RBLOCK
	| LBLOCK NEWLINE listterm+ RBLOCK;

listterm: term | term NEWLINE;

conditional:
	term EQUALS term				# equalsCond
	| term LTHAN term				# lthanCond
	| term GTHAN term				# gthanCond
	| term LTHAN_EQUALS term		# lthanequalsCond
	| term GTHAN_EQUALS term		# gthanequalsCond
	| conditional OR conditional	# orCond
	| conditional AND conditional	# andCond
	| boolexpr						# boolCond;

math:
	math op = (PLUS | SUB) math	# additiveMath
	| SUB math					# unarySubMath
	| num						# soloMath;

literal: num | STRING | boolexpr | QUOTED;
num: INT;
boolexpr: TRUE | FALSE;

pipe: term PIPE fncall | term PIPE pipe | fncall PIPE pipe;

// Comments and whitespace
COMMENT: '#' ~[\r\n]* -> skip;
NEWLINE: [\r\n] | [\r\n]+;
WHITESPACE: (' ' | '\t')+ -> channel(HIDDEN);

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

LABEL: ':' [a-zA-Z]+; // TODO Refactor unicode so this works

// Keywords
IF: 'if';
DO: 'do';
END: 'end';
FUNC: 'func';
RETURN: 'return';
FOR: 'for';
PLUS: 'add';
SUB: 'sub';
DIV: 'div';
MULT: 'mult';
TRUE: 'true';
FALSE: 'false';
OR: 'or';
AND: 'and';

// Int is above so we recognize it as as a number and not a string
INT: [0-9]+;

// Multiline strings
QUOTED: '"' (STRINGORNEW+)* '"';
fragment STRINGORNEW: (STRING | NEWLINE);

// Strings matching all unicode 
STRING: (ESC | ~["\\]);
fragment ESC: '\\' (["\\/bfnrt] | UNICODE);
fragment UNICODE: 'u' HEX HEX HEX HEX;
fragment HEX: [0-9a-fA-F];