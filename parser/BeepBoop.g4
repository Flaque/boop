grammar BeepBoop;
boop: (NEWLINE | NEWLINE+) (code | code+) EOF
	| (code | code+) EOF;

code: funcdef # funcdefCode | statement # statementCode;

funcdef:
	FUNC label DO END
	| FUNC label+ DO END
	| FUNC label DO funcguts END
	| FUNC label+ DO funcguts END;

funcguts: statement+;

statement:
	assignstat				    # assignStatement
	| returnstat			    # returnStatement
	| exportstat		    	# exportStatement
	| importstat NEWLINE	# importStatement
	| fncall NEWLINE		  # fncallStatement
	| ifstat NEWLINE		  # ifStatement
	| pipe NEWLINE			  # pipeStatement
	| NEWLINE			      	# newlineStatement;

importstat: IMPORT words | IMPORT words AS label;
exportstat: EXPORT (term | label);

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
	label ASSIGN term	  	# exprAssign
	| label ASSIGN fncall	# fncallAssign;

paren_fncall: LPAREN fncall RPAREN;

fncall: label fnargs+ | label;

fnargs: FLAG #flagFnargs
	| MULT     #multFnargs
	| DIV      #divFnargs
	| term     #termFnargs
	; 

term:
	label			      # labelTerm // Also a function call
	| literal		    # literalTerm
	| math			    # mathTerm
	| paren_fncall	# parenfncallTerm
	| structexpr	  # structTerm
	| list		    	# listTerm;

list:
	LBLOCK RBLOCK
	| LBLOCK listterm+ RBLOCK
	| LBLOCK NEWLINE listterm+ RBLOCK;

listterm: term | term NEWLINE;

conditional:
	term EQUALS term				# equalsCond
	| term LTHAN term				# lthanCond
	| term GTHAN term				      # gthanCond
	| term LTHAN_EQUALS term		  # lthanequalsCond
	| term GTHAN_EQUALS term	  	# gthanequalsCond
	| conditional OR conditional	# orCond
	| conditional AND conditional	# andCond
	| boolexpr					        	# boolCond;

math:
	math op = (PLUS | SUB) math		# additiveMath
	| math op = (MULT | DIV) math	# multiplicativeMath
	| SUB math			        			# unarySubMath
	| num						            	# soloMath;

literal:
	num			    # numLiteral
	| words	   	# wordsLiteral
	| boolexpr	# boolLiteral;

num: INT  #intNum 
	| FLOAT #floatNum;

words: LETTERS #letterWords
	| STRING     #stringWords
	| QUOTED     #quotedWords;

boolexpr: TRUE | FALSE;

pipe:
	pipe PIPE fncall          #pipeToFncallPipe
	| pipe PIPE pipe          #pipeToPipe
	| pipe NEWLINE PIPE pipe  #newlinePipe
	| fncall					 				#fncallPipe
	| term  						 			#termPipe;

label: ':' LETTERS;
FLAG: '-'+ LETTERS+;

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
PLUS: '+';
SUB: '-';
DIV: '/';
MULT: '*';

// Keywords
IMPORT: 'import';
EXPORT: 'export';
AS: 'as';
IF: 'if';
DO: 'do';
END: 'end';
FUNC: 'func';
RETURN: 'return';
FOR: 'for';
TRUE: 'true';
FALSE: 'false';
OR: 'or';
AND: 'and';

// Int is above so we recognize it as as a number and not a string
INT: [0-9]+;
FLOAT: INT '.' INT;

// Multiline strings
QUOTED: '"' (STRINGORNEW+)* '"';
fragment STRINGORNEW: (STRING | NEWLINE);

// Strings matching all unicode 
LETTERS: [a-zA-Z]+;
STRING: (ESC | ~["\\]);
fragment ESC: '\\' (["\\/bfnrt] | LETTERS);