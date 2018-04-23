// Generated from /home/flaque/projects/go/src/github.com/Flaque/boop/parser/BeepBoop.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class BeepBoopParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, FLAG=2, COMMENT=3, NEWLINE=4, WHITESPACE=5, EQUALS=6, ASSIGN=7, 
		PIPE=8, LTHAN=9, GTHAN=10, LTHAN_EQUALS=11, GTHAN_EQUALS=12, LPAREN=13, 
		RPAREN=14, LSQUIG=15, RSQUIG=16, LBLOCK=17, RBLOCK=18, PLUS=19, SUB=20, 
		DIV=21, MULT=22, IMPORT=23, EXPORT=24, AS=25, IF=26, DO=27, END=28, FUNC=29, 
		RETURN=30, FOR=31, TRUE=32, FALSE=33, OR=34, AND=35, INT=36, FLOAT=37, 
		QUOTED=38, LETTERS=39, STRING=40;
	public static final int
		RULE_boop = 0, RULE_code = 1, RULE_funcdef = 2, RULE_funcguts = 3, RULE_statement = 4, 
		RULE_importstat = 5, RULE_exportstat = 6, RULE_ifstat = 7, RULE_returnstat = 8, 
		RULE_structexpr = 9, RULE_assignstat = 10, RULE_assign = 11, RULE_paren_fncall = 12, 
		RULE_fncall = 13, RULE_fnargs = 14, RULE_term = 15, RULE_list = 16, RULE_listterm = 17, 
		RULE_conditional = 18, RULE_math = 19, RULE_literal = 20, RULE_num = 21, 
		RULE_words = 22, RULE_boolexpr = 23, RULE_pipe = 24, RULE_label = 25;
	public static final String[] ruleNames = {
		"boop", "code", "funcdef", "funcguts", "statement", "importstat", "exportstat", 
		"ifstat", "returnstat", "structexpr", "assignstat", "assign", "paren_fncall", 
		"fncall", "fnargs", "term", "list", "listterm", "conditional", "math", 
		"literal", "num", "words", "boolexpr", "pipe", "label"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "':'", null, null, null, null, "'=='", "'='", "'|'", "'<'", "'>'", 
		"'<='", "'>='", "'('", "')'", "'{'", "'}'", "'['", "']'", "'+'", "'-'", 
		"'/'", "'*'", "'import'", "'export'", "'as'", "'if'", "'do'", "'end'", 
		"'func'", "'return'", "'for'", "'true'", "'false'", "'or'", "'and'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, "FLAG", "COMMENT", "NEWLINE", "WHITESPACE", "EQUALS", "ASSIGN", 
		"PIPE", "LTHAN", "GTHAN", "LTHAN_EQUALS", "GTHAN_EQUALS", "LPAREN", "RPAREN", 
		"LSQUIG", "RSQUIG", "LBLOCK", "RBLOCK", "PLUS", "SUB", "DIV", "MULT", 
		"IMPORT", "EXPORT", "AS", "IF", "DO", "END", "FUNC", "RETURN", "FOR", 
		"TRUE", "FALSE", "OR", "AND", "INT", "FLOAT", "QUOTED", "LETTERS", "STRING"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "BeepBoop.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public BeepBoopParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class BoopContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(BeepBoopParser.EOF, 0); }
		public List<TerminalNode> NEWLINE() { return getTokens(BeepBoopParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(BeepBoopParser.NEWLINE, i);
		}
		public List<CodeContext> code() {
			return getRuleContexts(CodeContext.class);
		}
		public CodeContext code(int i) {
			return getRuleContext(CodeContext.class,i);
		}
		public BoopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_boop; }
	}

	public final BoopContext boop() throws RecognitionException {
		BoopContext _localctx = new BoopContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_boop);
		int _la;
		try {
			int _alt;
			setState(80);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(58);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
				case 1:
					{
					setState(52);
					match(NEWLINE);
					}
					break;
				case 2:
					{
					setState(54); 
					_errHandler.sync(this);
					_alt = 1;
					do {
						switch (_alt) {
						case 1:
							{
							{
							setState(53);
							match(NEWLINE);
							}
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						setState(56); 
						_errHandler.sync(this);
						_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
					} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
					}
					break;
				}
				setState(66);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
				case 1:
					{
					setState(60);
					code();
					}
					break;
				case 2:
					{
					setState(62); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(61);
						code();
						}
						}
						setState(64); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << NEWLINE) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << SUB) | (1L << IMPORT) | (1L << EXPORT) | (1L << IF) | (1L << FUNC) | (1L << RETURN) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << FLOAT) | (1L << QUOTED) | (1L << LETTERS) | (1L << STRING))) != 0) );
					}
					break;
				}
				setState(68);
				match(EOF);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(76);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
				case 1:
					{
					setState(70);
					code();
					}
					break;
				case 2:
					{
					setState(72); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(71);
						code();
						}
						}
						setState(74); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << NEWLINE) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << SUB) | (1L << IMPORT) | (1L << EXPORT) | (1L << IF) | (1L << FUNC) | (1L << RETURN) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << FLOAT) | (1L << QUOTED) | (1L << LETTERS) | (1L << STRING))) != 0) );
					}
					break;
				}
				setState(78);
				match(EOF);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CodeContext extends ParserRuleContext {
		public CodeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_code; }
	 
		public CodeContext() { }
		public void copyFrom(CodeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class StatementCodeContext extends CodeContext {
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public StatementCodeContext(CodeContext ctx) { copyFrom(ctx); }
	}
	public static class FuncdefCodeContext extends CodeContext {
		public FuncdefContext funcdef() {
			return getRuleContext(FuncdefContext.class,0);
		}
		public FuncdefCodeContext(CodeContext ctx) { copyFrom(ctx); }
	}

	public final CodeContext code() throws RecognitionException {
		CodeContext _localctx = new CodeContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_code);
		try {
			setState(84);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case FUNC:
				_localctx = new FuncdefCodeContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(82);
				funcdef();
				}
				break;
			case T__0:
			case NEWLINE:
			case LPAREN:
			case LSQUIG:
			case LBLOCK:
			case SUB:
			case IMPORT:
			case EXPORT:
			case IF:
			case RETURN:
			case TRUE:
			case FALSE:
			case INT:
			case FLOAT:
			case QUOTED:
			case LETTERS:
			case STRING:
				_localctx = new StatementCodeContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(83);
				statement();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FuncdefContext extends ParserRuleContext {
		public TerminalNode FUNC() { return getToken(BeepBoopParser.FUNC, 0); }
		public List<LabelContext> label() {
			return getRuleContexts(LabelContext.class);
		}
		public LabelContext label(int i) {
			return getRuleContext(LabelContext.class,i);
		}
		public TerminalNode DO() { return getToken(BeepBoopParser.DO, 0); }
		public TerminalNode END() { return getToken(BeepBoopParser.END, 0); }
		public FuncgutsContext funcguts() {
			return getRuleContext(FuncgutsContext.class,0);
		}
		public FuncdefContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcdef; }
	}

	public final FuncdefContext funcdef() throws RecognitionException {
		FuncdefContext _localctx = new FuncdefContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_funcdef);
		int _la;
		try {
			setState(116);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(86);
				match(FUNC);
				setState(87);
				label();
				setState(88);
				match(DO);
				setState(89);
				match(END);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(91);
				match(FUNC);
				setState(93); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(92);
					label();
					}
					}
					setState(95); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__0 );
				setState(97);
				match(DO);
				setState(98);
				match(END);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(100);
				match(FUNC);
				setState(101);
				label();
				setState(102);
				match(DO);
				setState(103);
				funcguts();
				setState(104);
				match(END);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(106);
				match(FUNC);
				setState(108); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(107);
					label();
					}
					}
					setState(110); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__0 );
				setState(112);
				match(DO);
				setState(113);
				funcguts();
				setState(114);
				match(END);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FuncgutsContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public FuncgutsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcguts; }
	}

	public final FuncgutsContext funcguts() throws RecognitionException {
		FuncgutsContext _localctx = new FuncgutsContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_funcguts);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(119); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(118);
				statement();
				}
				}
				setState(121); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << NEWLINE) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << SUB) | (1L << IMPORT) | (1L << EXPORT) | (1L << IF) | (1L << RETURN) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << FLOAT) | (1L << QUOTED) | (1L << LETTERS) | (1L << STRING))) != 0) );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementContext extends ParserRuleContext {
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	 
		public StatementContext() { }
		public void copyFrom(StatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class PipeStatementContext extends StatementContext {
		public PipeContext pipe() {
			return getRuleContext(PipeContext.class,0);
		}
		public TerminalNode NEWLINE() { return getToken(BeepBoopParser.NEWLINE, 0); }
		public PipeStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class ExportStatementContext extends StatementContext {
		public ExportstatContext exportstat() {
			return getRuleContext(ExportstatContext.class,0);
		}
		public ExportStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class ImportStatementContext extends StatementContext {
		public ImportstatContext importstat() {
			return getRuleContext(ImportstatContext.class,0);
		}
		public TerminalNode NEWLINE() { return getToken(BeepBoopParser.NEWLINE, 0); }
		public ImportStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class AssignStatementContext extends StatementContext {
		public AssignstatContext assignstat() {
			return getRuleContext(AssignstatContext.class,0);
		}
		public AssignStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class NewlineStatementContext extends StatementContext {
		public TerminalNode NEWLINE() { return getToken(BeepBoopParser.NEWLINE, 0); }
		public NewlineStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class ReturnStatementContext extends StatementContext {
		public ReturnstatContext returnstat() {
			return getRuleContext(ReturnstatContext.class,0);
		}
		public ReturnStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class IfStatementContext extends StatementContext {
		public IfstatContext ifstat() {
			return getRuleContext(IfstatContext.class,0);
		}
		public TerminalNode NEWLINE() { return getToken(BeepBoopParser.NEWLINE, 0); }
		public IfStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class FncallStatementContext extends StatementContext {
		public FncallContext fncall() {
			return getRuleContext(FncallContext.class,0);
		}
		public TerminalNode NEWLINE() { return getToken(BeepBoopParser.NEWLINE, 0); }
		public FncallStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_statement);
		try {
			setState(139);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				_localctx = new AssignStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(123);
				assignstat();
				}
				break;
			case 2:
				_localctx = new ReturnStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(124);
				returnstat();
				}
				break;
			case 3:
				_localctx = new ExportStatementContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(125);
				exportstat();
				}
				break;
			case 4:
				_localctx = new ImportStatementContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(126);
				importstat();
				setState(127);
				match(NEWLINE);
				}
				break;
			case 5:
				_localctx = new FncallStatementContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(129);
				fncall();
				setState(130);
				match(NEWLINE);
				}
				break;
			case 6:
				_localctx = new IfStatementContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(132);
				ifstat();
				setState(133);
				match(NEWLINE);
				}
				break;
			case 7:
				_localctx = new PipeStatementContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(135);
				pipe(0);
				setState(136);
				match(NEWLINE);
				}
				break;
			case 8:
				_localctx = new NewlineStatementContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(138);
				match(NEWLINE);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ImportstatContext extends ParserRuleContext {
		public TerminalNode IMPORT() { return getToken(BeepBoopParser.IMPORT, 0); }
		public WordsContext words() {
			return getRuleContext(WordsContext.class,0);
		}
		public TerminalNode AS() { return getToken(BeepBoopParser.AS, 0); }
		public LabelContext label() {
			return getRuleContext(LabelContext.class,0);
		}
		public ImportstatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importstat; }
	}

	public final ImportstatContext importstat() throws RecognitionException {
		ImportstatContext _localctx = new ImportstatContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_importstat);
		try {
			setState(148);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(141);
				match(IMPORT);
				setState(142);
				words();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(143);
				match(IMPORT);
				setState(144);
				words();
				setState(145);
				match(AS);
				setState(146);
				label();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExportstatContext extends ParserRuleContext {
		public TerminalNode EXPORT() { return getToken(BeepBoopParser.EXPORT, 0); }
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public LabelContext label() {
			return getRuleContext(LabelContext.class,0);
		}
		public ExportstatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exportstat; }
	}

	public final ExportstatContext exportstat() throws RecognitionException {
		ExportstatContext _localctx = new ExportstatContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_exportstat);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(150);
			match(EXPORT);
			setState(153);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				{
				setState(151);
				term();
				}
				break;
			case 2:
				{
				setState(152);
				label();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IfstatContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(BeepBoopParser.IF, 0); }
		public ConditionalContext conditional() {
			return getRuleContext(ConditionalContext.class,0);
		}
		public TerminalNode DO() { return getToken(BeepBoopParser.DO, 0); }
		public TerminalNode END() { return getToken(BeepBoopParser.END, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public IfstatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifstat; }
	}

	public final IfstatContext ifstat() throws RecognitionException {
		IfstatContext _localctx = new IfstatContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_ifstat);
		int _la;
		try {
			setState(170);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(155);
				match(IF);
				setState(156);
				conditional(0);
				setState(157);
				match(DO);
				setState(159); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(158);
					statement();
					}
					}
					setState(161); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << NEWLINE) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << SUB) | (1L << IMPORT) | (1L << EXPORT) | (1L << IF) | (1L << RETURN) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << FLOAT) | (1L << QUOTED) | (1L << LETTERS) | (1L << STRING))) != 0) );
				setState(163);
				match(END);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(165);
				match(IF);
				setState(166);
				conditional(0);
				setState(167);
				match(DO);
				setState(168);
				match(END);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ReturnstatContext extends ParserRuleContext {
		public TerminalNode RETURN() { return getToken(BeepBoopParser.RETURN, 0); }
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public List<TerminalNode> NEWLINE() { return getTokens(BeepBoopParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(BeepBoopParser.NEWLINE, i);
		}
		public ReturnstatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnstat; }
	}

	public final ReturnstatContext returnstat() throws RecognitionException {
		ReturnstatContext _localctx = new ReturnstatContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_returnstat);
		try {
			int _alt;
			setState(193);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(172);
				match(RETURN);
				setState(173);
				term();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(174);
				match(RETURN);
				setState(175);
				term();
				setState(182);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
				case 1:
					{
					setState(176);
					match(NEWLINE);
					}
					break;
				case 2:
					{
					setState(178); 
					_errHandler.sync(this);
					_alt = 1;
					do {
						switch (_alt) {
						case 1:
							{
							{
							setState(177);
							match(NEWLINE);
							}
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						setState(180); 
						_errHandler.sync(this);
						_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
					} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
					}
					break;
				}
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(184);
				match(RETURN);
				setState(191);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
				case 1:
					{
					setState(185);
					match(NEWLINE);
					}
					break;
				case 2:
					{
					setState(187); 
					_errHandler.sync(this);
					_alt = 1;
					do {
						switch (_alt) {
						case 1:
							{
							{
							setState(186);
							match(NEWLINE);
							}
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						setState(189); 
						_errHandler.sync(this);
						_alt = getInterpreter().adaptivePredict(_input,19,_ctx);
					} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
					}
					break;
				}
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StructexprContext extends ParserRuleContext {
		public TerminalNode LSQUIG() { return getToken(BeepBoopParser.LSQUIG, 0); }
		public TerminalNode RSQUIG() { return getToken(BeepBoopParser.RSQUIG, 0); }
		public TerminalNode NEWLINE() { return getToken(BeepBoopParser.NEWLINE, 0); }
		public List<AssignstatContext> assignstat() {
			return getRuleContexts(AssignstatContext.class);
		}
		public AssignstatContext assignstat(int i) {
			return getRuleContext(AssignstatContext.class,i);
		}
		public AssignContext assign() {
			return getRuleContext(AssignContext.class,0);
		}
		public StructexprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structexpr; }
	}

	public final StructexprContext structexpr() throws RecognitionException {
		StructexprContext _localctx = new StructexprContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_structexpr);
		int _la;
		try {
			setState(229);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(195);
				match(LSQUIG);
				setState(196);
				match(RSQUIG);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(197);
				match(LSQUIG);
				setState(198);
				match(NEWLINE);
				setState(205);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
				case 1:
					{
					setState(199);
					assignstat();
					}
					break;
				case 2:
					{
					setState(201); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(200);
						assignstat();
						}
						}
						setState(203); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( _la==T__0 );
					}
					break;
				}
				setState(207);
				match(RSQUIG);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(209);
				match(LSQUIG);
				setState(216);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
				case 1:
					{
					setState(210);
					assignstat();
					}
					break;
				case 2:
					{
					setState(212); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(211);
						assignstat();
						}
						}
						setState(214); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( _la==T__0 );
					}
					break;
				}
				setState(218);
				match(RSQUIG);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(220);
				match(LSQUIG);
				setState(221);
				assign();
				setState(222);
				match(RSQUIG);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(224);
				match(LSQUIG);
				setState(225);
				match(NEWLINE);
				setState(226);
				assign();
				setState(227);
				match(RSQUIG);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AssignstatContext extends ParserRuleContext {
		public AssignContext assign() {
			return getRuleContext(AssignContext.class,0);
		}
		public TerminalNode NEWLINE() { return getToken(BeepBoopParser.NEWLINE, 0); }
		public AssignstatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignstat; }
	}

	public final AssignstatContext assignstat() throws RecognitionException {
		AssignstatContext _localctx = new AssignstatContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_assignstat);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(231);
			assign();
			setState(232);
			match(NEWLINE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AssignContext extends ParserRuleContext {
		public AssignContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign; }
	 
		public AssignContext() { }
		public void copyFrom(AssignContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class FncallAssignContext extends AssignContext {
		public LabelContext label() {
			return getRuleContext(LabelContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(BeepBoopParser.ASSIGN, 0); }
		public FncallContext fncall() {
			return getRuleContext(FncallContext.class,0);
		}
		public FncallAssignContext(AssignContext ctx) { copyFrom(ctx); }
	}
	public static class ExprAssignContext extends AssignContext {
		public LabelContext label() {
			return getRuleContext(LabelContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(BeepBoopParser.ASSIGN, 0); }
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public ExprAssignContext(AssignContext ctx) { copyFrom(ctx); }
	}

	public final AssignContext assign() throws RecognitionException {
		AssignContext _localctx = new AssignContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_assign);
		try {
			setState(242);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				_localctx = new ExprAssignContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(234);
				label();
				setState(235);
				match(ASSIGN);
				setState(236);
				term();
				}
				break;
			case 2:
				_localctx = new FncallAssignContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(238);
				label();
				setState(239);
				match(ASSIGN);
				setState(240);
				fncall();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Paren_fncallContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(BeepBoopParser.LPAREN, 0); }
		public FncallContext fncall() {
			return getRuleContext(FncallContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(BeepBoopParser.RPAREN, 0); }
		public Paren_fncallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paren_fncall; }
	}

	public final Paren_fncallContext paren_fncall() throws RecognitionException {
		Paren_fncallContext _localctx = new Paren_fncallContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_paren_fncall);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(244);
			match(LPAREN);
			setState(245);
			fncall();
			setState(246);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FncallContext extends ParserRuleContext {
		public LabelContext label() {
			return getRuleContext(LabelContext.class,0);
		}
		public List<FnargsContext> fnargs() {
			return getRuleContexts(FnargsContext.class);
		}
		public FnargsContext fnargs(int i) {
			return getRuleContext(FnargsContext.class,i);
		}
		public FncallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fncall; }
	}

	public final FncallContext fncall() throws RecognitionException {
		FncallContext _localctx = new FncallContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_fncall);
		try {
			int _alt;
			setState(255);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(248);
				label();
				setState(250); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(249);
						fnargs();
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(252); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(254);
				label();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FnargsContext extends ParserRuleContext {
		public FnargsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fnargs; }
	 
		public FnargsContext() { }
		public void copyFrom(FnargsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class TermFnargsContext extends FnargsContext {
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public TermFnargsContext(FnargsContext ctx) { copyFrom(ctx); }
	}
	public static class DivFnargsContext extends FnargsContext {
		public TerminalNode DIV() { return getToken(BeepBoopParser.DIV, 0); }
		public DivFnargsContext(FnargsContext ctx) { copyFrom(ctx); }
	}
	public static class FlagFnargsContext extends FnargsContext {
		public TerminalNode FLAG() { return getToken(BeepBoopParser.FLAG, 0); }
		public FlagFnargsContext(FnargsContext ctx) { copyFrom(ctx); }
	}
	public static class MultFnargsContext extends FnargsContext {
		public TerminalNode MULT() { return getToken(BeepBoopParser.MULT, 0); }
		public MultFnargsContext(FnargsContext ctx) { copyFrom(ctx); }
	}

	public final FnargsContext fnargs() throws RecognitionException {
		FnargsContext _localctx = new FnargsContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_fnargs);
		try {
			setState(261);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case FLAG:
				_localctx = new FlagFnargsContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(257);
				match(FLAG);
				}
				break;
			case MULT:
				_localctx = new MultFnargsContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(258);
				match(MULT);
				}
				break;
			case DIV:
				_localctx = new DivFnargsContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(259);
				match(DIV);
				}
				break;
			case T__0:
			case LPAREN:
			case LSQUIG:
			case LBLOCK:
			case SUB:
			case TRUE:
			case FALSE:
			case INT:
			case FLOAT:
			case QUOTED:
			case LETTERS:
			case STRING:
				_localctx = new TermFnargsContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(260);
				term();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TermContext extends ParserRuleContext {
		public TermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_term; }
	 
		public TermContext() { }
		public void copyFrom(TermContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class ParenfncallTermContext extends TermContext {
		public Paren_fncallContext paren_fncall() {
			return getRuleContext(Paren_fncallContext.class,0);
		}
		public ParenfncallTermContext(TermContext ctx) { copyFrom(ctx); }
	}
	public static class LiteralTermContext extends TermContext {
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public LiteralTermContext(TermContext ctx) { copyFrom(ctx); }
	}
	public static class LabelTermContext extends TermContext {
		public LabelContext label() {
			return getRuleContext(LabelContext.class,0);
		}
		public LabelTermContext(TermContext ctx) { copyFrom(ctx); }
	}
	public static class StructTermContext extends TermContext {
		public StructexprContext structexpr() {
			return getRuleContext(StructexprContext.class,0);
		}
		public StructTermContext(TermContext ctx) { copyFrom(ctx); }
	}
	public static class MathTermContext extends TermContext {
		public MathContext math() {
			return getRuleContext(MathContext.class,0);
		}
		public MathTermContext(TermContext ctx) { copyFrom(ctx); }
	}
	public static class ListTermContext extends TermContext {
		public ListContext list() {
			return getRuleContext(ListContext.class,0);
		}
		public ListTermContext(TermContext ctx) { copyFrom(ctx); }
	}

	public final TermContext term() throws RecognitionException {
		TermContext _localctx = new TermContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_term);
		try {
			setState(269);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				_localctx = new LabelTermContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(263);
				label();
				}
				break;
			case 2:
				_localctx = new LiteralTermContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(264);
				literal();
				}
				break;
			case 3:
				_localctx = new MathTermContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(265);
				math(0);
				}
				break;
			case 4:
				_localctx = new ParenfncallTermContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(266);
				paren_fncall();
				}
				break;
			case 5:
				_localctx = new StructTermContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(267);
				structexpr();
				}
				break;
			case 6:
				_localctx = new ListTermContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(268);
				list();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ListContext extends ParserRuleContext {
		public TerminalNode LBLOCK() { return getToken(BeepBoopParser.LBLOCK, 0); }
		public TerminalNode RBLOCK() { return getToken(BeepBoopParser.RBLOCK, 0); }
		public List<ListtermContext> listterm() {
			return getRuleContexts(ListtermContext.class);
		}
		public ListtermContext listterm(int i) {
			return getRuleContext(ListtermContext.class,i);
		}
		public TerminalNode NEWLINE() { return getToken(BeepBoopParser.NEWLINE, 0); }
		public ListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list; }
	}

	public final ListContext list() throws RecognitionException {
		ListContext _localctx = new ListContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_list);
		int _la;
		try {
			setState(290);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(271);
				match(LBLOCK);
				setState(272);
				match(RBLOCK);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(273);
				match(LBLOCK);
				setState(275); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(274);
					listterm();
					}
					}
					setState(277); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << FLOAT) | (1L << QUOTED) | (1L << LETTERS) | (1L << STRING))) != 0) );
				setState(279);
				match(RBLOCK);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(281);
				match(LBLOCK);
				setState(282);
				match(NEWLINE);
				setState(284); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(283);
					listterm();
					}
					}
					setState(286); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << FLOAT) | (1L << QUOTED) | (1L << LETTERS) | (1L << STRING))) != 0) );
				setState(288);
				match(RBLOCK);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ListtermContext extends ParserRuleContext {
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public TerminalNode NEWLINE() { return getToken(BeepBoopParser.NEWLINE, 0); }
		public ListtermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listterm; }
	}

	public final ListtermContext listterm() throws RecognitionException {
		ListtermContext _localctx = new ListtermContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_listterm);
		try {
			setState(296);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,35,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(292);
				term();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(293);
				term();
				setState(294);
				match(NEWLINE);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConditionalContext extends ParserRuleContext {
		public ConditionalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditional; }
	 
		public ConditionalContext() { }
		public void copyFrom(ConditionalContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class GthanequalsCondContext extends ConditionalContext {
		public List<TermContext> term() {
			return getRuleContexts(TermContext.class);
		}
		public TermContext term(int i) {
			return getRuleContext(TermContext.class,i);
		}
		public TerminalNode GTHAN_EQUALS() { return getToken(BeepBoopParser.GTHAN_EQUALS, 0); }
		public GthanequalsCondContext(ConditionalContext ctx) { copyFrom(ctx); }
	}
	public static class EqualsCondContext extends ConditionalContext {
		public List<TermContext> term() {
			return getRuleContexts(TermContext.class);
		}
		public TermContext term(int i) {
			return getRuleContext(TermContext.class,i);
		}
		public TerminalNode EQUALS() { return getToken(BeepBoopParser.EQUALS, 0); }
		public EqualsCondContext(ConditionalContext ctx) { copyFrom(ctx); }
	}
	public static class OrCondContext extends ConditionalContext {
		public List<ConditionalContext> conditional() {
			return getRuleContexts(ConditionalContext.class);
		}
		public ConditionalContext conditional(int i) {
			return getRuleContext(ConditionalContext.class,i);
		}
		public TerminalNode OR() { return getToken(BeepBoopParser.OR, 0); }
		public OrCondContext(ConditionalContext ctx) { copyFrom(ctx); }
	}
	public static class GthanCondContext extends ConditionalContext {
		public List<TermContext> term() {
			return getRuleContexts(TermContext.class);
		}
		public TermContext term(int i) {
			return getRuleContext(TermContext.class,i);
		}
		public TerminalNode GTHAN() { return getToken(BeepBoopParser.GTHAN, 0); }
		public GthanCondContext(ConditionalContext ctx) { copyFrom(ctx); }
	}
	public static class AndCondContext extends ConditionalContext {
		public List<ConditionalContext> conditional() {
			return getRuleContexts(ConditionalContext.class);
		}
		public ConditionalContext conditional(int i) {
			return getRuleContext(ConditionalContext.class,i);
		}
		public TerminalNode AND() { return getToken(BeepBoopParser.AND, 0); }
		public AndCondContext(ConditionalContext ctx) { copyFrom(ctx); }
	}
	public static class LthanCondContext extends ConditionalContext {
		public List<TermContext> term() {
			return getRuleContexts(TermContext.class);
		}
		public TermContext term(int i) {
			return getRuleContext(TermContext.class,i);
		}
		public TerminalNode LTHAN() { return getToken(BeepBoopParser.LTHAN, 0); }
		public LthanCondContext(ConditionalContext ctx) { copyFrom(ctx); }
	}
	public static class LthanequalsCondContext extends ConditionalContext {
		public List<TermContext> term() {
			return getRuleContexts(TermContext.class);
		}
		public TermContext term(int i) {
			return getRuleContext(TermContext.class,i);
		}
		public TerminalNode LTHAN_EQUALS() { return getToken(BeepBoopParser.LTHAN_EQUALS, 0); }
		public LthanequalsCondContext(ConditionalContext ctx) { copyFrom(ctx); }
	}
	public static class BoolCondContext extends ConditionalContext {
		public BoolexprContext boolexpr() {
			return getRuleContext(BoolexprContext.class,0);
		}
		public BoolCondContext(ConditionalContext ctx) { copyFrom(ctx); }
	}

	public final ConditionalContext conditional() throws RecognitionException {
		return conditional(0);
	}

	private ConditionalContext conditional(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ConditionalContext _localctx = new ConditionalContext(_ctx, _parentState);
		ConditionalContext _prevctx = _localctx;
		int _startState = 36;
		enterRecursionRule(_localctx, 36, RULE_conditional, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(320);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,36,_ctx) ) {
			case 1:
				{
				_localctx = new EqualsCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(299);
				term();
				setState(300);
				match(EQUALS);
				setState(301);
				term();
				}
				break;
			case 2:
				{
				_localctx = new LthanCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(303);
				term();
				setState(304);
				match(LTHAN);
				setState(305);
				term();
				}
				break;
			case 3:
				{
				_localctx = new GthanCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(307);
				term();
				setState(308);
				match(GTHAN);
				setState(309);
				term();
				}
				break;
			case 4:
				{
				_localctx = new LthanequalsCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(311);
				term();
				setState(312);
				match(LTHAN_EQUALS);
				setState(313);
				term();
				}
				break;
			case 5:
				{
				_localctx = new GthanequalsCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(315);
				term();
				setState(316);
				match(GTHAN_EQUALS);
				setState(317);
				term();
				}
				break;
			case 6:
				{
				_localctx = new BoolCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(319);
				boolexpr();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(330);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,38,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(328);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,37,_ctx) ) {
					case 1:
						{
						_localctx = new OrCondContext(new ConditionalContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_conditional);
						setState(322);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(323);
						match(OR);
						setState(324);
						conditional(4);
						}
						break;
					case 2:
						{
						_localctx = new AndCondContext(new ConditionalContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_conditional);
						setState(325);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(326);
						match(AND);
						setState(327);
						conditional(3);
						}
						break;
					}
					} 
				}
				setState(332);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,38,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class MathContext extends ParserRuleContext {
		public MathContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_math; }
	 
		public MathContext() { }
		public void copyFrom(MathContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class SoloMathContext extends MathContext {
		public NumContext num() {
			return getRuleContext(NumContext.class,0);
		}
		public SoloMathContext(MathContext ctx) { copyFrom(ctx); }
	}
	public static class MultiplicativeMathContext extends MathContext {
		public Token op;
		public List<MathContext> math() {
			return getRuleContexts(MathContext.class);
		}
		public MathContext math(int i) {
			return getRuleContext(MathContext.class,i);
		}
		public TerminalNode MULT() { return getToken(BeepBoopParser.MULT, 0); }
		public TerminalNode DIV() { return getToken(BeepBoopParser.DIV, 0); }
		public MultiplicativeMathContext(MathContext ctx) { copyFrom(ctx); }
	}
	public static class AdditiveMathContext extends MathContext {
		public Token op;
		public List<MathContext> math() {
			return getRuleContexts(MathContext.class);
		}
		public MathContext math(int i) {
			return getRuleContext(MathContext.class,i);
		}
		public TerminalNode PLUS() { return getToken(BeepBoopParser.PLUS, 0); }
		public TerminalNode SUB() { return getToken(BeepBoopParser.SUB, 0); }
		public AdditiveMathContext(MathContext ctx) { copyFrom(ctx); }
	}
	public static class UnarySubMathContext extends MathContext {
		public TerminalNode SUB() { return getToken(BeepBoopParser.SUB, 0); }
		public MathContext math() {
			return getRuleContext(MathContext.class,0);
		}
		public UnarySubMathContext(MathContext ctx) { copyFrom(ctx); }
	}

	public final MathContext math() throws RecognitionException {
		return math(0);
	}

	private MathContext math(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		MathContext _localctx = new MathContext(_ctx, _parentState);
		MathContext _prevctx = _localctx;
		int _startState = 38;
		enterRecursionRule(_localctx, 38, RULE_math, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(337);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SUB:
				{
				_localctx = new UnarySubMathContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(334);
				match(SUB);
				setState(335);
				math(4);
				}
				break;
			case INT:
			case FLOAT:
				{
				_localctx = new SoloMathContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(336);
				num();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(347);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,41,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(345);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,40,_ctx) ) {
					case 1:
						{
						_localctx = new AdditiveMathContext(new MathContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_math);
						setState(339);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(340);
						((AdditiveMathContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==PLUS || _la==SUB) ) {
							((AdditiveMathContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(341);
						math(4);
						}
						break;
					case 2:
						{
						_localctx = new MultiplicativeMathContext(new MathContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_math);
						setState(342);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(343);
						((MultiplicativeMathContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==DIV || _la==MULT) ) {
							((MultiplicativeMathContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(344);
						math(3);
						}
						break;
					}
					} 
				}
				setState(349);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,41,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class LiteralContext extends ParserRuleContext {
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	 
		public LiteralContext() { }
		public void copyFrom(LiteralContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class NumLiteralContext extends LiteralContext {
		public NumContext num() {
			return getRuleContext(NumContext.class,0);
		}
		public NumLiteralContext(LiteralContext ctx) { copyFrom(ctx); }
	}
	public static class WordsLiteralContext extends LiteralContext {
		public WordsContext words() {
			return getRuleContext(WordsContext.class,0);
		}
		public WordsLiteralContext(LiteralContext ctx) { copyFrom(ctx); }
	}
	public static class BoolLiteralContext extends LiteralContext {
		public BoolexprContext boolexpr() {
			return getRuleContext(BoolexprContext.class,0);
		}
		public BoolLiteralContext(LiteralContext ctx) { copyFrom(ctx); }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_literal);
		try {
			setState(353);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
			case FLOAT:
				_localctx = new NumLiteralContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(350);
				num();
				}
				break;
			case QUOTED:
			case LETTERS:
			case STRING:
				_localctx = new WordsLiteralContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(351);
				words();
				}
				break;
			case TRUE:
			case FALSE:
				_localctx = new BoolLiteralContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(352);
				boolexpr();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NumContext extends ParserRuleContext {
		public NumContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_num; }
	 
		public NumContext() { }
		public void copyFrom(NumContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class FloatNumContext extends NumContext {
		public TerminalNode FLOAT() { return getToken(BeepBoopParser.FLOAT, 0); }
		public FloatNumContext(NumContext ctx) { copyFrom(ctx); }
	}
	public static class IntNumContext extends NumContext {
		public TerminalNode INT() { return getToken(BeepBoopParser.INT, 0); }
		public IntNumContext(NumContext ctx) { copyFrom(ctx); }
	}

	public final NumContext num() throws RecognitionException {
		NumContext _localctx = new NumContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_num);
		try {
			setState(357);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				_localctx = new IntNumContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(355);
				match(INT);
				}
				break;
			case FLOAT:
				_localctx = new FloatNumContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(356);
				match(FLOAT);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class WordsContext extends ParserRuleContext {
		public WordsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_words; }
	 
		public WordsContext() { }
		public void copyFrom(WordsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class QuotedWordsContext extends WordsContext {
		public TerminalNode QUOTED() { return getToken(BeepBoopParser.QUOTED, 0); }
		public QuotedWordsContext(WordsContext ctx) { copyFrom(ctx); }
	}
	public static class LetterWordsContext extends WordsContext {
		public TerminalNode LETTERS() { return getToken(BeepBoopParser.LETTERS, 0); }
		public LetterWordsContext(WordsContext ctx) { copyFrom(ctx); }
	}
	public static class StringWordsContext extends WordsContext {
		public TerminalNode STRING() { return getToken(BeepBoopParser.STRING, 0); }
		public StringWordsContext(WordsContext ctx) { copyFrom(ctx); }
	}

	public final WordsContext words() throws RecognitionException {
		WordsContext _localctx = new WordsContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_words);
		try {
			setState(362);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LETTERS:
				_localctx = new LetterWordsContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(359);
				match(LETTERS);
				}
				break;
			case STRING:
				_localctx = new StringWordsContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(360);
				match(STRING);
				}
				break;
			case QUOTED:
				_localctx = new QuotedWordsContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(361);
				match(QUOTED);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BoolexprContext extends ParserRuleContext {
		public TerminalNode TRUE() { return getToken(BeepBoopParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(BeepBoopParser.FALSE, 0); }
		public BoolexprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_boolexpr; }
	}

	public final BoolexprContext boolexpr() throws RecognitionException {
		BoolexprContext _localctx = new BoolexprContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_boolexpr);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(364);
			_la = _input.LA(1);
			if ( !(_la==TRUE || _la==FALSE) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PipeContext extends ParserRuleContext {
		public PipeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pipe; }
	 
		public PipeContext() { }
		public void copyFrom(PipeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class PipeToFncallPipeContext extends PipeContext {
		public PipeContext pipe() {
			return getRuleContext(PipeContext.class,0);
		}
		public TerminalNode PIPE() { return getToken(BeepBoopParser.PIPE, 0); }
		public FncallContext fncall() {
			return getRuleContext(FncallContext.class,0);
		}
		public PipeToFncallPipeContext(PipeContext ctx) { copyFrom(ctx); }
	}
	public static class PipeToPipeContext extends PipeContext {
		public List<PipeContext> pipe() {
			return getRuleContexts(PipeContext.class);
		}
		public PipeContext pipe(int i) {
			return getRuleContext(PipeContext.class,i);
		}
		public TerminalNode PIPE() { return getToken(BeepBoopParser.PIPE, 0); }
		public PipeToPipeContext(PipeContext ctx) { copyFrom(ctx); }
	}
	public static class TermPipeContext extends PipeContext {
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public TermPipeContext(PipeContext ctx) { copyFrom(ctx); }
	}
	public static class NewlinePipeContext extends PipeContext {
		public List<PipeContext> pipe() {
			return getRuleContexts(PipeContext.class);
		}
		public PipeContext pipe(int i) {
			return getRuleContext(PipeContext.class,i);
		}
		public TerminalNode NEWLINE() { return getToken(BeepBoopParser.NEWLINE, 0); }
		public TerminalNode PIPE() { return getToken(BeepBoopParser.PIPE, 0); }
		public NewlinePipeContext(PipeContext ctx) { copyFrom(ctx); }
	}
	public static class FncallPipeContext extends PipeContext {
		public FncallContext fncall() {
			return getRuleContext(FncallContext.class,0);
		}
		public FncallPipeContext(PipeContext ctx) { copyFrom(ctx); }
	}

	public final PipeContext pipe() throws RecognitionException {
		return pipe(0);
	}

	private PipeContext pipe(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		PipeContext _localctx = new PipeContext(_ctx, _parentState);
		PipeContext _prevctx = _localctx;
		int _startState = 48;
		enterRecursionRule(_localctx, 48, RULE_pipe, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(369);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,45,_ctx) ) {
			case 1:
				{
				_localctx = new FncallPipeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(367);
				fncall();
				}
				break;
			case 2:
				{
				_localctx = new TermPipeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(368);
				term();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(383);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,47,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(381);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,46,_ctx) ) {
					case 1:
						{
						_localctx = new PipeToPipeContext(new PipeContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_pipe);
						setState(371);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(372);
						match(PIPE);
						setState(373);
						pipe(5);
						}
						break;
					case 2:
						{
						_localctx = new NewlinePipeContext(new PipeContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_pipe);
						setState(374);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(375);
						match(NEWLINE);
						setState(376);
						match(PIPE);
						setState(377);
						pipe(4);
						}
						break;
					case 3:
						{
						_localctx = new PipeToFncallPipeContext(new PipeContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_pipe);
						setState(378);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(379);
						match(PIPE);
						setState(380);
						fncall();
						}
						break;
					}
					} 
				}
				setState(385);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,47,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class LabelContext extends ParserRuleContext {
		public TerminalNode LETTERS() { return getToken(BeepBoopParser.LETTERS, 0); }
		public LabelContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_label; }
	}

	public final LabelContext label() throws RecognitionException {
		LabelContext _localctx = new LabelContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_label);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(386);
			match(T__0);
			setState(387);
			match(LETTERS);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 18:
			return conditional_sempred((ConditionalContext)_localctx, predIndex);
		case 19:
			return math_sempred((MathContext)_localctx, predIndex);
		case 24:
			return pipe_sempred((PipeContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean conditional_sempred(ConditionalContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 3);
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean math_sempred(MathContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 3);
		case 3:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean pipe_sempred(PipeContext _localctx, int predIndex) {
		switch (predIndex) {
		case 4:
			return precpred(_ctx, 4);
		case 5:
			return precpred(_ctx, 3);
		case 6:
			return precpred(_ctx, 5);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3*\u0188\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\3\2\3\2\6\29\n\2\r\2\16\2:\5\2=\n\2\3\2\3\2\6\2A"+
		"\n\2\r\2\16\2B\5\2E\n\2\3\2\3\2\3\2\3\2\6\2K\n\2\r\2\16\2L\5\2O\n\2\3"+
		"\2\3\2\5\2S\n\2\3\3\3\3\5\3W\n\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\6\4`\n\4"+
		"\r\4\16\4a\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\6\4o\n\4\r\4\16"+
		"\4p\3\4\3\4\3\4\3\4\5\4w\n\4\3\5\6\5z\n\5\r\5\16\5{\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6\u008e\n\6\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\3\7\5\7\u0097\n\7\3\b\3\b\3\b\5\b\u009c\n\b\3\t\3\t\3\t"+
		"\3\t\6\t\u00a2\n\t\r\t\16\t\u00a3\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u00ad"+
		"\n\t\3\n\3\n\3\n\3\n\3\n\3\n\6\n\u00b5\n\n\r\n\16\n\u00b6\5\n\u00b9\n"+
		"\n\3\n\3\n\3\n\6\n\u00be\n\n\r\n\16\n\u00bf\5\n\u00c2\n\n\5\n\u00c4\n"+
		"\n\3\13\3\13\3\13\3\13\3\13\3\13\6\13\u00cc\n\13\r\13\16\13\u00cd\5\13"+
		"\u00d0\n\13\3\13\3\13\3\13\3\13\3\13\6\13\u00d7\n\13\r\13\16\13\u00d8"+
		"\5\13\u00db\n\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\5\13\u00e8\n\13\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u00f5"+
		"\n\r\3\16\3\16\3\16\3\16\3\17\3\17\6\17\u00fd\n\17\r\17\16\17\u00fe\3"+
		"\17\5\17\u0102\n\17\3\20\3\20\3\20\3\20\5\20\u0108\n\20\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\5\21\u0110\n\21\3\22\3\22\3\22\3\22\6\22\u0116\n\22\r"+
		"\22\16\22\u0117\3\22\3\22\3\22\3\22\3\22\6\22\u011f\n\22\r\22\16\22\u0120"+
		"\3\22\3\22\5\22\u0125\n\22\3\23\3\23\3\23\3\23\5\23\u012b\n\23\3\24\3"+
		"\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3"+
		"\24\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u0143\n\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\7\24\u014b\n\24\f\24\16\24\u014e\13\24\3\25\3\25\3\25\3\25"+
		"\5\25\u0154\n\25\3\25\3\25\3\25\3\25\3\25\3\25\7\25\u015c\n\25\f\25\16"+
		"\25\u015f\13\25\3\26\3\26\3\26\5\26\u0164\n\26\3\27\3\27\5\27\u0168\n"+
		"\27\3\30\3\30\3\30\5\30\u016d\n\30\3\31\3\31\3\32\3\32\3\32\5\32\u0174"+
		"\n\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\7\32\u0180\n\32"+
		"\f\32\16\32\u0183\13\32\3\33\3\33\3\33\3\33\2\5&(\62\34\2\4\6\b\n\f\16"+
		"\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\2\5\3\2\25\26\3\2\27\30\3\2"+
		"\"#\2\u01b7\2R\3\2\2\2\4V\3\2\2\2\6v\3\2\2\2\by\3\2\2\2\n\u008d\3\2\2"+
		"\2\f\u0096\3\2\2\2\16\u0098\3\2\2\2\20\u00ac\3\2\2\2\22\u00c3\3\2\2\2"+
		"\24\u00e7\3\2\2\2\26\u00e9\3\2\2\2\30\u00f4\3\2\2\2\32\u00f6\3\2\2\2\34"+
		"\u0101\3\2\2\2\36\u0107\3\2\2\2 \u010f\3\2\2\2\"\u0124\3\2\2\2$\u012a"+
		"\3\2\2\2&\u0142\3\2\2\2(\u0153\3\2\2\2*\u0163\3\2\2\2,\u0167\3\2\2\2."+
		"\u016c\3\2\2\2\60\u016e\3\2\2\2\62\u0173\3\2\2\2\64\u0184\3\2\2\2\66="+
		"\7\6\2\2\679\7\6\2\28\67\3\2\2\29:\3\2\2\2:8\3\2\2\2:;\3\2\2\2;=\3\2\2"+
		"\2<\66\3\2\2\2<8\3\2\2\2=D\3\2\2\2>E\5\4\3\2?A\5\4\3\2@?\3\2\2\2AB\3\2"+
		"\2\2B@\3\2\2\2BC\3\2\2\2CE\3\2\2\2D>\3\2\2\2D@\3\2\2\2EF\3\2\2\2FG\7\2"+
		"\2\3GS\3\2\2\2HO\5\4\3\2IK\5\4\3\2JI\3\2\2\2KL\3\2\2\2LJ\3\2\2\2LM\3\2"+
		"\2\2MO\3\2\2\2NH\3\2\2\2NJ\3\2\2\2OP\3\2\2\2PQ\7\2\2\3QS\3\2\2\2R<\3\2"+
		"\2\2RN\3\2\2\2S\3\3\2\2\2TW\5\6\4\2UW\5\n\6\2VT\3\2\2\2VU\3\2\2\2W\5\3"+
		"\2\2\2XY\7\37\2\2YZ\5\64\33\2Z[\7\35\2\2[\\\7\36\2\2\\w\3\2\2\2]_\7\37"+
		"\2\2^`\5\64\33\2_^\3\2\2\2`a\3\2\2\2a_\3\2\2\2ab\3\2\2\2bc\3\2\2\2cd\7"+
		"\35\2\2de\7\36\2\2ew\3\2\2\2fg\7\37\2\2gh\5\64\33\2hi\7\35\2\2ij\5\b\5"+
		"\2jk\7\36\2\2kw\3\2\2\2ln\7\37\2\2mo\5\64\33\2nm\3\2\2\2op\3\2\2\2pn\3"+
		"\2\2\2pq\3\2\2\2qr\3\2\2\2rs\7\35\2\2st\5\b\5\2tu\7\36\2\2uw\3\2\2\2v"+
		"X\3\2\2\2v]\3\2\2\2vf\3\2\2\2vl\3\2\2\2w\7\3\2\2\2xz\5\n\6\2yx\3\2\2\2"+
		"z{\3\2\2\2{y\3\2\2\2{|\3\2\2\2|\t\3\2\2\2}\u008e\5\26\f\2~\u008e\5\22"+
		"\n\2\177\u008e\5\16\b\2\u0080\u0081\5\f\7\2\u0081\u0082\7\6\2\2\u0082"+
		"\u008e\3\2\2\2\u0083\u0084\5\34\17\2\u0084\u0085\7\6\2\2\u0085\u008e\3"+
		"\2\2\2\u0086\u0087\5\20\t\2\u0087\u0088\7\6\2\2\u0088\u008e\3\2\2\2\u0089"+
		"\u008a\5\62\32\2\u008a\u008b\7\6\2\2\u008b\u008e\3\2\2\2\u008c\u008e\7"+
		"\6\2\2\u008d}\3\2\2\2\u008d~\3\2\2\2\u008d\177\3\2\2\2\u008d\u0080\3\2"+
		"\2\2\u008d\u0083\3\2\2\2\u008d\u0086\3\2\2\2\u008d\u0089\3\2\2\2\u008d"+
		"\u008c\3\2\2\2\u008e\13\3\2\2\2\u008f\u0090\7\31\2\2\u0090\u0097\5.\30"+
		"\2\u0091\u0092\7\31\2\2\u0092\u0093\5.\30\2\u0093\u0094\7\33\2\2\u0094"+
		"\u0095\5\64\33\2\u0095\u0097\3\2\2\2\u0096\u008f\3\2\2\2\u0096\u0091\3"+
		"\2\2\2\u0097\r\3\2\2\2\u0098\u009b\7\32\2\2\u0099\u009c\5 \21\2\u009a"+
		"\u009c\5\64\33\2\u009b\u0099\3\2\2\2\u009b\u009a\3\2\2\2\u009c\17\3\2"+
		"\2\2\u009d\u009e\7\34\2\2\u009e\u009f\5&\24\2\u009f\u00a1\7\35\2\2\u00a0"+
		"\u00a2\5\n\6\2\u00a1\u00a0\3\2\2\2\u00a2\u00a3\3\2\2\2\u00a3\u00a1\3\2"+
		"\2\2\u00a3\u00a4\3\2\2\2\u00a4\u00a5\3\2\2\2\u00a5\u00a6\7\36\2\2\u00a6"+
		"\u00ad\3\2\2\2\u00a7\u00a8\7\34\2\2\u00a8\u00a9\5&\24\2\u00a9\u00aa\7"+
		"\35\2\2\u00aa\u00ab\7\36\2\2\u00ab\u00ad\3\2\2\2\u00ac\u009d\3\2\2\2\u00ac"+
		"\u00a7\3\2\2\2\u00ad\21\3\2\2\2\u00ae\u00af\7 \2\2\u00af\u00c4\5 \21\2"+
		"\u00b0\u00b1\7 \2\2\u00b1\u00b8\5 \21\2\u00b2\u00b9\7\6\2\2\u00b3\u00b5"+
		"\7\6\2\2\u00b4\u00b3\3\2\2\2\u00b5\u00b6\3\2\2\2\u00b6\u00b4\3\2\2\2\u00b6"+
		"\u00b7\3\2\2\2\u00b7\u00b9\3\2\2\2\u00b8\u00b2\3\2\2\2\u00b8\u00b4\3\2"+
		"\2\2\u00b9\u00c4\3\2\2\2\u00ba\u00c1\7 \2\2\u00bb\u00c2\7\6\2\2\u00bc"+
		"\u00be\7\6\2\2\u00bd\u00bc\3\2\2\2\u00be\u00bf\3\2\2\2\u00bf\u00bd\3\2"+
		"\2\2\u00bf\u00c0\3\2\2\2\u00c0\u00c2\3\2\2\2\u00c1\u00bb\3\2\2\2\u00c1"+
		"\u00bd\3\2\2\2\u00c2\u00c4\3\2\2\2\u00c3\u00ae\3\2\2\2\u00c3\u00b0\3\2"+
		"\2\2\u00c3\u00ba\3\2\2\2\u00c4\23\3\2\2\2\u00c5\u00c6\7\21\2\2\u00c6\u00e8"+
		"\7\22\2\2\u00c7\u00c8\7\21\2\2\u00c8\u00cf\7\6\2\2\u00c9\u00d0\5\26\f"+
		"\2\u00ca\u00cc\5\26\f\2\u00cb\u00ca\3\2\2\2\u00cc\u00cd\3\2\2\2\u00cd"+
		"\u00cb\3\2\2\2\u00cd\u00ce\3\2\2\2\u00ce\u00d0\3\2\2\2\u00cf\u00c9\3\2"+
		"\2\2\u00cf\u00cb\3\2\2\2\u00d0\u00d1\3\2\2\2\u00d1\u00d2\7\22\2\2\u00d2"+
		"\u00e8\3\2\2\2\u00d3\u00da\7\21\2\2\u00d4\u00db\5\26\f\2\u00d5\u00d7\5"+
		"\26\f\2\u00d6\u00d5\3\2\2\2\u00d7\u00d8\3\2\2\2\u00d8\u00d6\3\2\2\2\u00d8"+
		"\u00d9\3\2\2\2\u00d9\u00db\3\2\2\2\u00da\u00d4\3\2\2\2\u00da\u00d6\3\2"+
		"\2\2\u00db\u00dc\3\2\2\2\u00dc\u00dd\7\22\2\2\u00dd\u00e8\3\2\2\2\u00de"+
		"\u00df\7\21\2\2\u00df\u00e0\5\30\r\2\u00e0\u00e1\7\22\2\2\u00e1\u00e8"+
		"\3\2\2\2\u00e2\u00e3\7\21\2\2\u00e3\u00e4\7\6\2\2\u00e4\u00e5\5\30\r\2"+
		"\u00e5\u00e6\7\22\2\2\u00e6\u00e8\3\2\2\2\u00e7\u00c5\3\2\2\2\u00e7\u00c7"+
		"\3\2\2\2\u00e7\u00d3\3\2\2\2\u00e7\u00de\3\2\2\2\u00e7\u00e2\3\2\2\2\u00e8"+
		"\25\3\2\2\2\u00e9\u00ea\5\30\r\2\u00ea\u00eb\7\6\2\2\u00eb\27\3\2\2\2"+
		"\u00ec\u00ed\5\64\33\2\u00ed\u00ee\7\t\2\2\u00ee\u00ef\5 \21\2\u00ef\u00f5"+
		"\3\2\2\2\u00f0\u00f1\5\64\33\2\u00f1\u00f2\7\t\2\2\u00f2\u00f3\5\34\17"+
		"\2\u00f3\u00f5\3\2\2\2\u00f4\u00ec\3\2\2\2\u00f4\u00f0\3\2\2\2\u00f5\31"+
		"\3\2\2\2\u00f6\u00f7\7\17\2\2\u00f7\u00f8\5\34\17\2\u00f8\u00f9\7\20\2"+
		"\2\u00f9\33\3\2\2\2\u00fa\u00fc\5\64\33\2\u00fb\u00fd\5\36\20\2\u00fc"+
		"\u00fb\3\2\2\2\u00fd\u00fe\3\2\2\2\u00fe\u00fc\3\2\2\2\u00fe\u00ff\3\2"+
		"\2\2\u00ff\u0102\3\2\2\2\u0100\u0102\5\64\33\2\u0101\u00fa\3\2\2\2\u0101"+
		"\u0100\3\2\2\2\u0102\35\3\2\2\2\u0103\u0108\7\4\2\2\u0104\u0108\7\30\2"+
		"\2\u0105\u0108\7\27\2\2\u0106\u0108\5 \21\2\u0107\u0103\3\2\2\2\u0107"+
		"\u0104\3\2\2\2\u0107\u0105\3\2\2\2\u0107\u0106\3\2\2\2\u0108\37\3\2\2"+
		"\2\u0109\u0110\5\64\33\2\u010a\u0110\5*\26\2\u010b\u0110\5(\25\2\u010c"+
		"\u0110\5\32\16\2\u010d\u0110\5\24\13\2\u010e\u0110\5\"\22\2\u010f\u0109"+
		"\3\2\2\2\u010f\u010a\3\2\2\2\u010f\u010b\3\2\2\2\u010f\u010c\3\2\2\2\u010f"+
		"\u010d\3\2\2\2\u010f\u010e\3\2\2\2\u0110!\3\2\2\2\u0111\u0112\7\23\2\2"+
		"\u0112\u0125\7\24\2\2\u0113\u0115\7\23\2\2\u0114\u0116\5$\23\2\u0115\u0114"+
		"\3\2\2\2\u0116\u0117\3\2\2\2\u0117\u0115\3\2\2\2\u0117\u0118\3\2\2\2\u0118"+
		"\u0119\3\2\2\2\u0119\u011a\7\24\2\2\u011a\u0125\3\2\2\2\u011b\u011c\7"+
		"\23\2\2\u011c\u011e\7\6\2\2\u011d\u011f\5$\23\2\u011e\u011d\3\2\2\2\u011f"+
		"\u0120\3\2\2\2\u0120\u011e\3\2\2\2\u0120\u0121\3\2\2\2\u0121\u0122\3\2"+
		"\2\2\u0122\u0123\7\24\2\2\u0123\u0125\3\2\2\2\u0124\u0111\3\2\2\2\u0124"+
		"\u0113\3\2\2\2\u0124\u011b\3\2\2\2\u0125#\3\2\2\2\u0126\u012b\5 \21\2"+
		"\u0127\u0128\5 \21\2\u0128\u0129\7\6\2\2\u0129\u012b\3\2\2\2\u012a\u0126"+
		"\3\2\2\2\u012a\u0127\3\2\2\2\u012b%\3\2\2\2\u012c\u012d\b\24\1\2\u012d"+
		"\u012e\5 \21\2\u012e\u012f\7\b\2\2\u012f\u0130\5 \21\2\u0130\u0143\3\2"+
		"\2\2\u0131\u0132\5 \21\2\u0132\u0133\7\13\2\2\u0133\u0134\5 \21\2\u0134"+
		"\u0143\3\2\2\2\u0135\u0136\5 \21\2\u0136\u0137\7\f\2\2\u0137\u0138\5 "+
		"\21\2\u0138\u0143\3\2\2\2\u0139\u013a\5 \21\2\u013a\u013b\7\r\2\2\u013b"+
		"\u013c\5 \21\2\u013c\u0143\3\2\2\2\u013d\u013e\5 \21\2\u013e\u013f\7\16"+
		"\2\2\u013f\u0140\5 \21\2\u0140\u0143\3\2\2\2\u0141\u0143\5\60\31\2\u0142"+
		"\u012c\3\2\2\2\u0142\u0131\3\2\2\2\u0142\u0135\3\2\2\2\u0142\u0139\3\2"+
		"\2\2\u0142\u013d\3\2\2\2\u0142\u0141\3\2\2\2\u0143\u014c\3\2\2\2\u0144"+
		"\u0145\f\5\2\2\u0145\u0146\7$\2\2\u0146\u014b\5&\24\6\u0147\u0148\f\4"+
		"\2\2\u0148\u0149\7%\2\2\u0149\u014b\5&\24\5\u014a\u0144\3\2\2\2\u014a"+
		"\u0147\3\2\2\2\u014b\u014e\3\2\2\2\u014c\u014a\3\2\2\2\u014c\u014d\3\2"+
		"\2\2\u014d\'\3\2\2\2\u014e\u014c\3\2\2\2\u014f\u0150\b\25\1\2\u0150\u0151"+
		"\7\26\2\2\u0151\u0154\5(\25\6\u0152\u0154\5,\27\2\u0153\u014f\3\2\2\2"+
		"\u0153\u0152\3\2\2\2\u0154\u015d\3\2\2\2\u0155\u0156\f\5\2\2\u0156\u0157"+
		"\t\2\2\2\u0157\u015c\5(\25\6\u0158\u0159\f\4\2\2\u0159\u015a\t\3\2\2\u015a"+
		"\u015c\5(\25\5\u015b\u0155\3\2\2\2\u015b\u0158\3\2\2\2\u015c\u015f\3\2"+
		"\2\2\u015d\u015b\3\2\2\2\u015d\u015e\3\2\2\2\u015e)\3\2\2\2\u015f\u015d"+
		"\3\2\2\2\u0160\u0164\5,\27\2\u0161\u0164\5.\30\2\u0162\u0164\5\60\31\2"+
		"\u0163\u0160\3\2\2\2\u0163\u0161\3\2\2\2\u0163\u0162\3\2\2\2\u0164+\3"+
		"\2\2\2\u0165\u0168\7&\2\2\u0166\u0168\7\'\2\2\u0167\u0165\3\2\2\2\u0167"+
		"\u0166\3\2\2\2\u0168-\3\2\2\2\u0169\u016d\7)\2\2\u016a\u016d\7*\2\2\u016b"+
		"\u016d\7(\2\2\u016c\u0169\3\2\2\2\u016c\u016a\3\2\2\2\u016c\u016b\3\2"+
		"\2\2\u016d/\3\2\2\2\u016e\u016f\t\4\2\2\u016f\61\3\2\2\2\u0170\u0171\b"+
		"\32\1\2\u0171\u0174\5\34\17\2\u0172\u0174\5 \21\2\u0173\u0170\3\2\2\2"+
		"\u0173\u0172\3\2\2\2\u0174\u0181\3\2\2\2\u0175\u0176\f\6\2\2\u0176\u0177"+
		"\7\n\2\2\u0177\u0180\5\62\32\7\u0178\u0179\f\5\2\2\u0179\u017a\7\6\2\2"+
		"\u017a\u017b\7\n\2\2\u017b\u0180\5\62\32\6\u017c\u017d\f\7\2\2\u017d\u017e"+
		"\7\n\2\2\u017e\u0180\5\34\17\2\u017f\u0175\3\2\2\2\u017f\u0178\3\2\2\2"+
		"\u017f\u017c\3\2\2\2\u0180\u0183\3\2\2\2\u0181\u017f\3\2\2\2\u0181\u0182"+
		"\3\2\2\2\u0182\63\3\2\2\2\u0183\u0181\3\2\2\2\u0184\u0185\7\3\2\2\u0185"+
		"\u0186\7)\2\2\u0186\65\3\2\2\2\62:<BDLNRVapv{\u008d\u0096\u009b\u00a3"+
		"\u00ac\u00b6\u00b8\u00bf\u00c1\u00c3\u00cd\u00cf\u00d8\u00da\u00e7\u00f4"+
		"\u00fe\u0101\u0107\u010f\u0117\u0120\u0124\u012a\u0142\u014a\u014c\u0153"+
		"\u015b\u015d\u0163\u0167\u016c\u0173\u017f\u0181";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}