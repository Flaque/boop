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
		COMMENT=1, NEWLINE=2, WHITESPACE=3, EQUALS=4, ASSIGN=5, PIPE=6, LTHAN=7, 
		GTHAN=8, LTHAN_EQUALS=9, GTHAN_EQUALS=10, LPAREN=11, RPAREN=12, LSQUIG=13, 
		RSQUIG=14, LBLOCK=15, RBLOCK=16, LABEL=17, IF=18, DO=19, END=20, FUNC=21, 
		RETURN=22, FOR=23, PLUS=24, SUB=25, DIV=26, MULT=27, TRUE=28, FALSE=29, 
		OR=30, AND=31, INT=32, QUOTED=33, STRING=34;
	public static final int
		RULE_boop = 0, RULE_code = 1, RULE_funcdef = 2, RULE_funcguts = 3, RULE_statement = 4, 
		RULE_ifstat = 5, RULE_returnstat = 6, RULE_structexpr = 7, RULE_assignstat = 8, 
		RULE_assign = 9, RULE_paren_fncall = 10, RULE_fncall = 11, RULE_term = 12, 
		RULE_list = 13, RULE_listterm = 14, RULE_conditional = 15, RULE_math = 16, 
		RULE_literal = 17, RULE_num = 18, RULE_boolexpr = 19, RULE_pipe = 20;
	public static final String[] ruleNames = {
		"boop", "code", "funcdef", "funcguts", "statement", "ifstat", "returnstat", 
		"structexpr", "assignstat", "assign", "paren_fncall", "fncall", "term", 
		"list", "listterm", "conditional", "math", "literal", "num", "boolexpr", 
		"pipe"
	};

	private static final String[] _LITERAL_NAMES = {
		null, null, null, null, "'=='", "'='", "'|'", "'<'", "'>'", "'<='", "'>='", 
		"'('", "')'", "'{'", "'}'", "'['", "']'", null, "'if'", "'do'", "'end'", 
		"'func'", "'return'", "'for'", "'add'", "'sub'", "'div'", "'mult'", "'true'", 
		"'false'", "'or'", "'and'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, "COMMENT", "NEWLINE", "WHITESPACE", "EQUALS", "ASSIGN", "PIPE", 
		"LTHAN", "GTHAN", "LTHAN_EQUALS", "GTHAN_EQUALS", "LPAREN", "RPAREN", 
		"LSQUIG", "RSQUIG", "LBLOCK", "RBLOCK", "LABEL", "IF", "DO", "END", "FUNC", 
		"RETURN", "FOR", "PLUS", "SUB", "DIV", "MULT", "TRUE", "FALSE", "OR", 
		"AND", "INT", "QUOTED", "STRING"
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
			setState(70);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(48);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
				case 1:
					{
					setState(42);
					match(NEWLINE);
					}
					break;
				case 2:
					{
					setState(44); 
					_errHandler.sync(this);
					_alt = 1;
					do {
						switch (_alt) {
						case 1:
							{
							{
							setState(43);
							match(NEWLINE);
							}
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						setState(46); 
						_errHandler.sync(this);
						_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
					} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
					}
					break;
				}
				setState(56);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
				case 1:
					{
					setState(50);
					code();
					}
					break;
				case 2:
					{
					setState(52); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(51);
						code();
						}
						}
						setState(54); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << NEWLINE) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << LABEL) | (1L << IF) | (1L << FUNC) | (1L << RETURN) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << QUOTED) | (1L << STRING))) != 0) );
					}
					break;
				}
				setState(58);
				match(EOF);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(66);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
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
					} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << NEWLINE) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << LABEL) | (1L << IF) | (1L << FUNC) | (1L << RETURN) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << QUOTED) | (1L << STRING))) != 0) );
					}
					break;
				}
				setState(68);
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
			setState(74);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case FUNC:
				_localctx = new FuncdefCodeContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(72);
				funcdef();
				}
				break;
			case NEWLINE:
			case LPAREN:
			case LSQUIG:
			case LBLOCK:
			case LABEL:
			case IF:
			case RETURN:
			case SUB:
			case TRUE:
			case FALSE:
			case INT:
			case QUOTED:
			case STRING:
				_localctx = new StatementCodeContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(73);
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
		public List<TerminalNode> LABEL() { return getTokens(BeepBoopParser.LABEL); }
		public TerminalNode LABEL(int i) {
			return getToken(BeepBoopParser.LABEL, i);
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
			setState(104);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(76);
				match(FUNC);
				setState(77);
				match(LABEL);
				setState(78);
				match(DO);
				setState(79);
				match(END);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(80);
				match(FUNC);
				setState(82); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(81);
					match(LABEL);
					}
					}
					setState(84); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==LABEL );
				setState(86);
				match(DO);
				setState(87);
				match(END);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(88);
				match(FUNC);
				setState(89);
				match(LABEL);
				setState(90);
				match(DO);
				setState(91);
				funcguts();
				setState(92);
				match(END);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(94);
				match(FUNC);
				setState(96); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(95);
					match(LABEL);
					}
					}
					setState(98); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==LABEL );
				setState(100);
				match(DO);
				setState(101);
				funcguts();
				setState(102);
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
			setState(107); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(106);
				statement();
				}
				}
				setState(109); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << NEWLINE) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << LABEL) | (1L << IF) | (1L << RETURN) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << QUOTED) | (1L << STRING))) != 0) );
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
			setState(123);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				_localctx = new AssignStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(111);
				assignstat();
				}
				break;
			case 2:
				_localctx = new ReturnStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(112);
				returnstat();
				}
				break;
			case 3:
				_localctx = new FncallStatementContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(113);
				fncall();
				setState(114);
				match(NEWLINE);
				}
				break;
			case 4:
				_localctx = new IfStatementContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(116);
				ifstat();
				setState(117);
				match(NEWLINE);
				}
				break;
			case 5:
				_localctx = new PipeStatementContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(119);
				pipe();
				setState(120);
				match(NEWLINE);
				}
				break;
			case 6:
				_localctx = new NewlineStatementContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(122);
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
		enterRule(_localctx, 10, RULE_ifstat);
		int _la;
		try {
			setState(140);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(125);
				match(IF);
				setState(126);
				conditional(0);
				setState(127);
				match(DO);
				setState(129); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(128);
					statement();
					}
					}
					setState(131); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << NEWLINE) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << LABEL) | (1L << IF) | (1L << RETURN) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << QUOTED) | (1L << STRING))) != 0) );
				setState(133);
				match(END);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(135);
				match(IF);
				setState(136);
				conditional(0);
				setState(137);
				match(DO);
				setState(138);
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
		enterRule(_localctx, 12, RULE_returnstat);
		try {
			int _alt;
			setState(163);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(142);
				match(RETURN);
				setState(143);
				term();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(144);
				match(RETURN);
				setState(145);
				term();
				setState(152);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
				case 1:
					{
					setState(146);
					match(NEWLINE);
					}
					break;
				case 2:
					{
					setState(148); 
					_errHandler.sync(this);
					_alt = 1;
					do {
						switch (_alt) {
						case 1:
							{
							{
							setState(147);
							match(NEWLINE);
							}
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						setState(150); 
						_errHandler.sync(this);
						_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
					} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
					}
					break;
				}
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(154);
				match(RETURN);
				setState(161);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
				case 1:
					{
					setState(155);
					match(NEWLINE);
					}
					break;
				case 2:
					{
					setState(157); 
					_errHandler.sync(this);
					_alt = 1;
					do {
						switch (_alt) {
						case 1:
							{
							{
							setState(156);
							match(NEWLINE);
							}
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						setState(159); 
						_errHandler.sync(this);
						_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
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
		enterRule(_localctx, 14, RULE_structexpr);
		int _la;
		try {
			setState(199);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(165);
				match(LSQUIG);
				setState(166);
				match(RSQUIG);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(167);
				match(LSQUIG);
				setState(168);
				match(NEWLINE);
				setState(175);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
				case 1:
					{
					setState(169);
					assignstat();
					}
					break;
				case 2:
					{
					setState(171); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(170);
						assignstat();
						}
						}
						setState(173); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( _la==LABEL );
					}
					break;
				}
				setState(177);
				match(RSQUIG);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(179);
				match(LSQUIG);
				setState(186);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
				case 1:
					{
					setState(180);
					assignstat();
					}
					break;
				case 2:
					{
					setState(182); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(181);
						assignstat();
						}
						}
						setState(184); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( _la==LABEL );
					}
					break;
				}
				setState(188);
				match(RSQUIG);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(190);
				match(LSQUIG);
				setState(191);
				assign();
				setState(192);
				match(RSQUIG);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(194);
				match(LSQUIG);
				setState(195);
				match(NEWLINE);
				setState(196);
				assign();
				setState(197);
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
		enterRule(_localctx, 16, RULE_assignstat);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(201);
			assign();
			setState(202);
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
		public TerminalNode LABEL() { return getToken(BeepBoopParser.LABEL, 0); }
		public TerminalNode ASSIGN() { return getToken(BeepBoopParser.ASSIGN, 0); }
		public FncallContext fncall() {
			return getRuleContext(FncallContext.class,0);
		}
		public FncallAssignContext(AssignContext ctx) { copyFrom(ctx); }
	}
	public static class ExprAssignContext extends AssignContext {
		public TerminalNode LABEL() { return getToken(BeepBoopParser.LABEL, 0); }
		public TerminalNode ASSIGN() { return getToken(BeepBoopParser.ASSIGN, 0); }
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public ExprAssignContext(AssignContext ctx) { copyFrom(ctx); }
	}

	public final AssignContext assign() throws RecognitionException {
		AssignContext _localctx = new AssignContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_assign);
		try {
			setState(210);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				_localctx = new ExprAssignContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(204);
				match(LABEL);
				setState(205);
				match(ASSIGN);
				setState(206);
				term();
				}
				break;
			case 2:
				_localctx = new FncallAssignContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(207);
				match(LABEL);
				setState(208);
				match(ASSIGN);
				setState(209);
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
		enterRule(_localctx, 20, RULE_paren_fncall);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(212);
			match(LPAREN);
			setState(213);
			fncall();
			setState(214);
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
		public TerminalNode LABEL() { return getToken(BeepBoopParser.LABEL, 0); }
		public List<TermContext> term() {
			return getRuleContexts(TermContext.class);
		}
		public TermContext term(int i) {
			return getRuleContext(TermContext.class,i);
		}
		public FncallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fncall; }
	}

	public final FncallContext fncall() throws RecognitionException {
		FncallContext _localctx = new FncallContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_fncall);
		int _la;
		try {
			setState(223);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(216);
				match(LABEL);
				setState(218); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(217);
					term();
					}
					}
					setState(220); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << LABEL) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << QUOTED) | (1L << STRING))) != 0) );
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(222);
				match(LABEL);
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
		public TerminalNode LABEL() { return getToken(BeepBoopParser.LABEL, 0); }
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
		enterRule(_localctx, 24, RULE_term);
		try {
			setState(231);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
			case 1:
				_localctx = new LabelTermContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(225);
				match(LABEL);
				}
				break;
			case 2:
				_localctx = new LiteralTermContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(226);
				literal();
				}
				break;
			case 3:
				_localctx = new MathTermContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(227);
				math(0);
				}
				break;
			case 4:
				_localctx = new ParenfncallTermContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(228);
				paren_fncall();
				}
				break;
			case 5:
				_localctx = new StructTermContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(229);
				structexpr();
				}
				break;
			case 6:
				_localctx = new ListTermContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(230);
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
		enterRule(_localctx, 26, RULE_list);
		int _la;
		try {
			setState(252);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(233);
				match(LBLOCK);
				setState(234);
				match(RBLOCK);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(235);
				match(LBLOCK);
				setState(237); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(236);
					listterm();
					}
					}
					setState(239); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << LABEL) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << QUOTED) | (1L << STRING))) != 0) );
				setState(241);
				match(RBLOCK);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(243);
				match(LBLOCK);
				setState(244);
				match(NEWLINE);
				setState(246); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(245);
					listterm();
					}
					}
					setState(248); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << LABEL) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << QUOTED) | (1L << STRING))) != 0) );
				setState(250);
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
		enterRule(_localctx, 28, RULE_listterm);
		try {
			setState(258);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(254);
				term();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(255);
				term();
				setState(256);
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
		int _startState = 30;
		enterRecursionRule(_localctx, 30, RULE_conditional, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(282);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				{
				_localctx = new EqualsCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(261);
				term();
				setState(262);
				match(EQUALS);
				setState(263);
				term();
				}
				break;
			case 2:
				{
				_localctx = new LthanCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(265);
				term();
				setState(266);
				match(LTHAN);
				setState(267);
				term();
				}
				break;
			case 3:
				{
				_localctx = new GthanCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(269);
				term();
				setState(270);
				match(GTHAN);
				setState(271);
				term();
				}
				break;
			case 4:
				{
				_localctx = new LthanequalsCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(273);
				term();
				setState(274);
				match(LTHAN_EQUALS);
				setState(275);
				term();
				}
				break;
			case 5:
				{
				_localctx = new GthanequalsCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(277);
				term();
				setState(278);
				match(GTHAN_EQUALS);
				setState(279);
				term();
				}
				break;
			case 6:
				{
				_localctx = new BoolCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(281);
				boolexpr();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(292);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,35,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(290);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
					case 1:
						{
						_localctx = new OrCondContext(new ConditionalContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_conditional);
						setState(284);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(285);
						match(OR);
						setState(286);
						conditional(4);
						}
						break;
					case 2:
						{
						_localctx = new AndCondContext(new ConditionalContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_conditional);
						setState(287);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(288);
						match(AND);
						setState(289);
						conditional(3);
						}
						break;
					}
					} 
				}
				setState(294);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,35,_ctx);
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
		int _startState = 32;
		enterRecursionRule(_localctx, 32, RULE_math, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(299);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SUB:
				{
				_localctx = new UnarySubMathContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(296);
				match(SUB);
				setState(297);
				math(2);
				}
				break;
			case INT:
				{
				_localctx = new SoloMathContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(298);
				num();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(306);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,37,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new AdditiveMathContext(new MathContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_math);
					setState(301);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(302);
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
					setState(303);
					math(4);
					}
					} 
				}
				setState(308);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,37,_ctx);
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
		public NumContext num() {
			return getRuleContext(NumContext.class,0);
		}
		public TerminalNode STRING() { return getToken(BeepBoopParser.STRING, 0); }
		public BoolexprContext boolexpr() {
			return getRuleContext(BoolexprContext.class,0);
		}
		public TerminalNode QUOTED() { return getToken(BeepBoopParser.QUOTED, 0); }
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_literal);
		try {
			setState(313);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(309);
				num();
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 2);
				{
				setState(310);
				match(STRING);
				}
				break;
			case TRUE:
			case FALSE:
				enterOuterAlt(_localctx, 3);
				{
				setState(311);
				boolexpr();
				}
				break;
			case QUOTED:
				enterOuterAlt(_localctx, 4);
				{
				setState(312);
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

	public static class NumContext extends ParserRuleContext {
		public TerminalNode INT() { return getToken(BeepBoopParser.INT, 0); }
		public NumContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_num; }
	}

	public final NumContext num() throws RecognitionException {
		NumContext _localctx = new NumContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_num);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(315);
			match(INT);
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
		enterRule(_localctx, 38, RULE_boolexpr);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(317);
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
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public TerminalNode PIPE() { return getToken(BeepBoopParser.PIPE, 0); }
		public FncallContext fncall() {
			return getRuleContext(FncallContext.class,0);
		}
		public PipeContext pipe() {
			return getRuleContext(PipeContext.class,0);
		}
		public PipeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pipe; }
	}

	public final PipeContext pipe() throws RecognitionException {
		PipeContext _localctx = new PipeContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_pipe);
		try {
			setState(331);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,39,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(319);
				term();
				setState(320);
				match(PIPE);
				setState(321);
				fncall();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(323);
				term();
				setState(324);
				match(PIPE);
				setState(325);
				pipe();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(327);
				fncall();
				setState(328);
				match(PIPE);
				setState(329);
				pipe();
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 15:
			return conditional_sempred((ConditionalContext)_localctx, predIndex);
		case 16:
			return math_sempred((MathContext)_localctx, predIndex);
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
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3$\u0150\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\3\2\3\2\6\2/\n\2\r\2\16\2\60"+
		"\5\2\63\n\2\3\2\3\2\6\2\67\n\2\r\2\16\28\5\2;\n\2\3\2\3\2\3\2\3\2\6\2"+
		"A\n\2\r\2\16\2B\5\2E\n\2\3\2\3\2\5\2I\n\2\3\3\3\3\5\3M\n\3\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\6\4U\n\4\r\4\16\4V\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\6\4c\n\4\r\4\16\4d\3\4\3\4\3\4\3\4\5\4k\n\4\3\5\6\5n\n\5\r\5\16\5o"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6~\n\6\3\7\3\7\3\7"+
		"\3\7\6\7\u0084\n\7\r\7\16\7\u0085\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7\u008f"+
		"\n\7\3\b\3\b\3\b\3\b\3\b\3\b\6\b\u0097\n\b\r\b\16\b\u0098\5\b\u009b\n"+
		"\b\3\b\3\b\3\b\6\b\u00a0\n\b\r\b\16\b\u00a1\5\b\u00a4\n\b\5\b\u00a6\n"+
		"\b\3\t\3\t\3\t\3\t\3\t\3\t\6\t\u00ae\n\t\r\t\16\t\u00af\5\t\u00b2\n\t"+
		"\3\t\3\t\3\t\3\t\3\t\6\t\u00b9\n\t\r\t\16\t\u00ba\5\t\u00bd\n\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u00ca\n\t\3\n\3\n\3\n\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\5\13\u00d5\n\13\3\f\3\f\3\f\3\f\3\r\3\r\6\r"+
		"\u00dd\n\r\r\r\16\r\u00de\3\r\5\r\u00e2\n\r\3\16\3\16\3\16\3\16\3\16\3"+
		"\16\5\16\u00ea\n\16\3\17\3\17\3\17\3\17\6\17\u00f0\n\17\r\17\16\17\u00f1"+
		"\3\17\3\17\3\17\3\17\3\17\6\17\u00f9\n\17\r\17\16\17\u00fa\3\17\3\17\5"+
		"\17\u00ff\n\17\3\20\3\20\3\20\3\20\5\20\u0105\n\20\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\5\21\u011d\n\21\3\21\3\21\3\21\3\21\3\21\3\21\7\21"+
		"\u0125\n\21\f\21\16\21\u0128\13\21\3\22\3\22\3\22\3\22\5\22\u012e\n\22"+
		"\3\22\3\22\3\22\7\22\u0133\n\22\f\22\16\22\u0136\13\22\3\23\3\23\3\23"+
		"\3\23\5\23\u013c\n\23\3\24\3\24\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26"+
		"\3\26\3\26\3\26\3\26\3\26\3\26\5\26\u014e\n\26\3\26\2\4 \"\27\2\4\6\b"+
		"\n\f\16\20\22\24\26\30\32\34\36 \"$&(*\2\4\3\2\32\33\3\2\36\37\2\u0178"+
		"\2H\3\2\2\2\4L\3\2\2\2\6j\3\2\2\2\bm\3\2\2\2\n}\3\2\2\2\f\u008e\3\2\2"+
		"\2\16\u00a5\3\2\2\2\20\u00c9\3\2\2\2\22\u00cb\3\2\2\2\24\u00d4\3\2\2\2"+
		"\26\u00d6\3\2\2\2\30\u00e1\3\2\2\2\32\u00e9\3\2\2\2\34\u00fe\3\2\2\2\36"+
		"\u0104\3\2\2\2 \u011c\3\2\2\2\"\u012d\3\2\2\2$\u013b\3\2\2\2&\u013d\3"+
		"\2\2\2(\u013f\3\2\2\2*\u014d\3\2\2\2,\63\7\4\2\2-/\7\4\2\2.-\3\2\2\2/"+
		"\60\3\2\2\2\60.\3\2\2\2\60\61\3\2\2\2\61\63\3\2\2\2\62,\3\2\2\2\62.\3"+
		"\2\2\2\63:\3\2\2\2\64;\5\4\3\2\65\67\5\4\3\2\66\65\3\2\2\2\678\3\2\2\2"+
		"8\66\3\2\2\289\3\2\2\29;\3\2\2\2:\64\3\2\2\2:\66\3\2\2\2;<\3\2\2\2<=\7"+
		"\2\2\3=I\3\2\2\2>E\5\4\3\2?A\5\4\3\2@?\3\2\2\2AB\3\2\2\2B@\3\2\2\2BC\3"+
		"\2\2\2CE\3\2\2\2D>\3\2\2\2D@\3\2\2\2EF\3\2\2\2FG\7\2\2\3GI\3\2\2\2H\62"+
		"\3\2\2\2HD\3\2\2\2I\3\3\2\2\2JM\5\6\4\2KM\5\n\6\2LJ\3\2\2\2LK\3\2\2\2"+
		"M\5\3\2\2\2NO\7\27\2\2OP\7\23\2\2PQ\7\25\2\2Qk\7\26\2\2RT\7\27\2\2SU\7"+
		"\23\2\2TS\3\2\2\2UV\3\2\2\2VT\3\2\2\2VW\3\2\2\2WX\3\2\2\2XY\7\25\2\2Y"+
		"k\7\26\2\2Z[\7\27\2\2[\\\7\23\2\2\\]\7\25\2\2]^\5\b\5\2^_\7\26\2\2_k\3"+
		"\2\2\2`b\7\27\2\2ac\7\23\2\2ba\3\2\2\2cd\3\2\2\2db\3\2\2\2de\3\2\2\2e"+
		"f\3\2\2\2fg\7\25\2\2gh\5\b\5\2hi\7\26\2\2ik\3\2\2\2jN\3\2\2\2jR\3\2\2"+
		"\2jZ\3\2\2\2j`\3\2\2\2k\7\3\2\2\2ln\5\n\6\2ml\3\2\2\2no\3\2\2\2om\3\2"+
		"\2\2op\3\2\2\2p\t\3\2\2\2q~\5\22\n\2r~\5\16\b\2st\5\30\r\2tu\7\4\2\2u"+
		"~\3\2\2\2vw\5\f\7\2wx\7\4\2\2x~\3\2\2\2yz\5*\26\2z{\7\4\2\2{~\3\2\2\2"+
		"|~\7\4\2\2}q\3\2\2\2}r\3\2\2\2}s\3\2\2\2}v\3\2\2\2}y\3\2\2\2}|\3\2\2\2"+
		"~\13\3\2\2\2\177\u0080\7\24\2\2\u0080\u0081\5 \21\2\u0081\u0083\7\25\2"+
		"\2\u0082\u0084\5\n\6\2\u0083\u0082\3\2\2\2\u0084\u0085\3\2\2\2\u0085\u0083"+
		"\3\2\2\2\u0085\u0086\3\2\2\2\u0086\u0087\3\2\2\2\u0087\u0088\7\26\2\2"+
		"\u0088\u008f\3\2\2\2\u0089\u008a\7\24\2\2\u008a\u008b\5 \21\2\u008b\u008c"+
		"\7\25\2\2\u008c\u008d\7\26\2\2\u008d\u008f\3\2\2\2\u008e\177\3\2\2\2\u008e"+
		"\u0089\3\2\2\2\u008f\r\3\2\2\2\u0090\u0091\7\30\2\2\u0091\u00a6\5\32\16"+
		"\2\u0092\u0093\7\30\2\2\u0093\u009a\5\32\16\2\u0094\u009b\7\4\2\2\u0095"+
		"\u0097\7\4\2\2\u0096\u0095\3\2\2\2\u0097\u0098\3\2\2\2\u0098\u0096\3\2"+
		"\2\2\u0098\u0099\3\2\2\2\u0099\u009b\3\2\2\2\u009a\u0094\3\2\2\2\u009a"+
		"\u0096\3\2\2\2\u009b\u00a6\3\2\2\2\u009c\u00a3\7\30\2\2\u009d\u00a4\7"+
		"\4\2\2\u009e\u00a0\7\4\2\2\u009f\u009e\3\2\2\2\u00a0\u00a1\3\2\2\2\u00a1"+
		"\u009f\3\2\2\2\u00a1\u00a2\3\2\2\2\u00a2\u00a4\3\2\2\2\u00a3\u009d\3\2"+
		"\2\2\u00a3\u009f\3\2\2\2\u00a4\u00a6\3\2\2\2\u00a5\u0090\3\2\2\2\u00a5"+
		"\u0092\3\2\2\2\u00a5\u009c\3\2\2\2\u00a6\17\3\2\2\2\u00a7\u00a8\7\17\2"+
		"\2\u00a8\u00ca\7\20\2\2\u00a9\u00aa\7\17\2\2\u00aa\u00b1\7\4\2\2\u00ab"+
		"\u00b2\5\22\n\2\u00ac\u00ae\5\22\n\2\u00ad\u00ac\3\2\2\2\u00ae\u00af\3"+
		"\2\2\2\u00af\u00ad\3\2\2\2\u00af\u00b0\3\2\2\2\u00b0\u00b2\3\2\2\2\u00b1"+
		"\u00ab\3\2\2\2\u00b1\u00ad\3\2\2\2\u00b2\u00b3\3\2\2\2\u00b3\u00b4\7\20"+
		"\2\2\u00b4\u00ca\3\2\2\2\u00b5\u00bc\7\17\2\2\u00b6\u00bd\5\22\n\2\u00b7"+
		"\u00b9\5\22\n\2\u00b8\u00b7\3\2\2\2\u00b9\u00ba\3\2\2\2\u00ba\u00b8\3"+
		"\2\2\2\u00ba\u00bb\3\2\2\2\u00bb\u00bd\3\2\2\2\u00bc\u00b6\3\2\2\2\u00bc"+
		"\u00b8\3\2\2\2\u00bd\u00be\3\2\2\2\u00be\u00bf\7\20\2\2\u00bf\u00ca\3"+
		"\2\2\2\u00c0\u00c1\7\17\2\2\u00c1\u00c2\5\24\13\2\u00c2\u00c3\7\20\2\2"+
		"\u00c3\u00ca\3\2\2\2\u00c4\u00c5\7\17\2\2\u00c5\u00c6\7\4\2\2\u00c6\u00c7"+
		"\5\24\13\2\u00c7\u00c8\7\20\2\2\u00c8\u00ca\3\2\2\2\u00c9\u00a7\3\2\2"+
		"\2\u00c9\u00a9\3\2\2\2\u00c9\u00b5\3\2\2\2\u00c9\u00c0\3\2\2\2\u00c9\u00c4"+
		"\3\2\2\2\u00ca\21\3\2\2\2\u00cb\u00cc\5\24\13\2\u00cc\u00cd\7\4\2\2\u00cd"+
		"\23\3\2\2\2\u00ce\u00cf\7\23\2\2\u00cf\u00d0\7\7\2\2\u00d0\u00d5\5\32"+
		"\16\2\u00d1\u00d2\7\23\2\2\u00d2\u00d3\7\7\2\2\u00d3\u00d5\5\30\r\2\u00d4"+
		"\u00ce\3\2\2\2\u00d4\u00d1\3\2\2\2\u00d5\25\3\2\2\2\u00d6\u00d7\7\r\2"+
		"\2\u00d7\u00d8\5\30\r\2\u00d8\u00d9\7\16\2\2\u00d9\27\3\2\2\2\u00da\u00dc"+
		"\7\23\2\2\u00db\u00dd\5\32\16\2\u00dc\u00db\3\2\2\2\u00dd\u00de\3\2\2"+
		"\2\u00de\u00dc\3\2\2\2\u00de\u00df\3\2\2\2\u00df\u00e2\3\2\2\2\u00e0\u00e2"+
		"\7\23\2\2\u00e1\u00da\3\2\2\2\u00e1\u00e0\3\2\2\2\u00e2\31\3\2\2\2\u00e3"+
		"\u00ea\7\23\2\2\u00e4\u00ea\5$\23\2\u00e5\u00ea\5\"\22\2\u00e6\u00ea\5"+
		"\26\f\2\u00e7\u00ea\5\20\t\2\u00e8\u00ea\5\34\17\2\u00e9\u00e3\3\2\2\2"+
		"\u00e9\u00e4\3\2\2\2\u00e9\u00e5\3\2\2\2\u00e9\u00e6\3\2\2\2\u00e9\u00e7"+
		"\3\2\2\2\u00e9\u00e8\3\2\2\2\u00ea\33\3\2\2\2\u00eb\u00ec\7\21\2\2\u00ec"+
		"\u00ff\7\22\2\2\u00ed\u00ef\7\21\2\2\u00ee\u00f0\5\36\20\2\u00ef\u00ee"+
		"\3\2\2\2\u00f0\u00f1\3\2\2\2\u00f1\u00ef\3\2\2\2\u00f1\u00f2\3\2\2\2\u00f2"+
		"\u00f3\3\2\2\2\u00f3\u00f4\7\22\2\2\u00f4\u00ff\3\2\2\2\u00f5\u00f6\7"+
		"\21\2\2\u00f6\u00f8\7\4\2\2\u00f7\u00f9\5\36\20\2\u00f8\u00f7\3\2\2\2"+
		"\u00f9\u00fa\3\2\2\2\u00fa\u00f8\3\2\2\2\u00fa\u00fb\3\2\2\2\u00fb\u00fc"+
		"\3\2\2\2\u00fc\u00fd\7\22\2\2\u00fd\u00ff\3\2\2\2\u00fe\u00eb\3\2\2\2"+
		"\u00fe\u00ed\3\2\2\2\u00fe\u00f5\3\2\2\2\u00ff\35\3\2\2\2\u0100\u0105"+
		"\5\32\16\2\u0101\u0102\5\32\16\2\u0102\u0103\7\4\2\2\u0103\u0105\3\2\2"+
		"\2\u0104\u0100\3\2\2\2\u0104\u0101\3\2\2\2\u0105\37\3\2\2\2\u0106\u0107"+
		"\b\21\1\2\u0107\u0108\5\32\16\2\u0108\u0109\7\6\2\2\u0109\u010a\5\32\16"+
		"\2\u010a\u011d\3\2\2\2\u010b\u010c\5\32\16\2\u010c\u010d\7\t\2\2\u010d"+
		"\u010e\5\32\16\2\u010e\u011d\3\2\2\2\u010f\u0110\5\32\16\2\u0110\u0111"+
		"\7\n\2\2\u0111\u0112\5\32\16\2\u0112\u011d\3\2\2\2\u0113\u0114\5\32\16"+
		"\2\u0114\u0115\7\13\2\2\u0115\u0116\5\32\16\2\u0116\u011d\3\2\2\2\u0117"+
		"\u0118\5\32\16\2\u0118\u0119\7\f\2\2\u0119\u011a\5\32\16\2\u011a\u011d"+
		"\3\2\2\2\u011b\u011d\5(\25\2\u011c\u0106\3\2\2\2\u011c\u010b\3\2\2\2\u011c"+
		"\u010f\3\2\2\2\u011c\u0113\3\2\2\2\u011c\u0117\3\2\2\2\u011c\u011b\3\2"+
		"\2\2\u011d\u0126\3\2\2\2\u011e\u011f\f\5\2\2\u011f\u0120\7 \2\2\u0120"+
		"\u0125\5 \21\6\u0121\u0122\f\4\2\2\u0122\u0123\7!\2\2\u0123\u0125\5 \21"+
		"\5\u0124\u011e\3\2\2\2\u0124\u0121\3\2\2\2\u0125\u0128\3\2\2\2\u0126\u0124"+
		"\3\2\2\2\u0126\u0127\3\2\2\2\u0127!\3\2\2\2\u0128\u0126\3\2\2\2\u0129"+
		"\u012a\b\22\1\2\u012a\u012b\7\33\2\2\u012b\u012e\5\"\22\4\u012c\u012e"+
		"\5&\24\2\u012d\u0129\3\2\2\2\u012d\u012c\3\2\2\2\u012e\u0134\3\2\2\2\u012f"+
		"\u0130\f\5\2\2\u0130\u0131\t\2\2\2\u0131\u0133\5\"\22\6\u0132\u012f\3"+
		"\2\2\2\u0133\u0136\3\2\2\2\u0134\u0132\3\2\2\2\u0134\u0135\3\2\2\2\u0135"+
		"#\3\2\2\2\u0136\u0134\3\2\2\2\u0137\u013c\5&\24\2\u0138\u013c\7$\2\2\u0139"+
		"\u013c\5(\25\2\u013a\u013c\7#\2\2\u013b\u0137\3\2\2\2\u013b\u0138\3\2"+
		"\2\2\u013b\u0139\3\2\2\2\u013b\u013a\3\2\2\2\u013c%\3\2\2\2\u013d\u013e"+
		"\7\"\2\2\u013e\'\3\2\2\2\u013f\u0140\t\3\2\2\u0140)\3\2\2\2\u0141\u0142"+
		"\5\32\16\2\u0142\u0143\7\b\2\2\u0143\u0144\5\30\r\2\u0144\u014e\3\2\2"+
		"\2\u0145\u0146\5\32\16\2\u0146\u0147\7\b\2\2\u0147\u0148\5*\26\2\u0148"+
		"\u014e\3\2\2\2\u0149\u014a\5\30\r\2\u014a\u014b\7\b\2\2\u014b\u014c\5"+
		"*\26\2\u014c\u014e\3\2\2\2\u014d\u0141\3\2\2\2\u014d\u0145\3\2\2\2\u014d"+
		"\u0149\3\2\2\2\u014e+\3\2\2\2*\60\628:BDHLVdjo}\u0085\u008e\u0098\u009a"+
		"\u00a1\u00a3\u00a5\u00af\u00b1\u00ba\u00bc\u00c9\u00d4\u00de\u00e1\u00e9"+
		"\u00f1\u00fa\u00fe\u0104\u011c\u0124\u0126\u012d\u0134\u013b\u014d";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}