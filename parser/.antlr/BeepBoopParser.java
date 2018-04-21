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
		T__0=1, COMMENT=2, NEWLINE=3, WHITESPACE=4, EQUALS=5, ASSIGN=6, PIPE=7, 
		LTHAN=8, GTHAN=9, LTHAN_EQUALS=10, GTHAN_EQUALS=11, LPAREN=12, RPAREN=13, 
		LSQUIG=14, RSQUIG=15, LBLOCK=16, RBLOCK=17, IF=18, DO=19, END=20, FUNC=21, 
		RETURN=22, FOR=23, PLUS=24, SUB=25, DIV=26, MULT=27, TRUE=28, FALSE=29, 
		OR=30, AND=31, INT=32, QUOTED=33, LETTERS=34, STRING=35;
	public static final int
		RULE_boop = 0, RULE_code = 1, RULE_funcdef = 2, RULE_funcguts = 3, RULE_statement = 4, 
		RULE_ifstat = 5, RULE_returnstat = 6, RULE_structexpr = 7, RULE_assignstat = 8, 
		RULE_assign = 9, RULE_paren_fncall = 10, RULE_fncall = 11, RULE_term = 12, 
		RULE_list = 13, RULE_listterm = 14, RULE_conditional = 15, RULE_math = 16, 
		RULE_literal = 17, RULE_num = 18, RULE_boolexpr = 19, RULE_pipe = 20, 
		RULE_label = 21;
	public static final String[] ruleNames = {
		"boop", "code", "funcdef", "funcguts", "statement", "ifstat", "returnstat", 
		"structexpr", "assignstat", "assign", "paren_fncall", "fncall", "term", 
		"list", "listterm", "conditional", "math", "literal", "num", "boolexpr", 
		"pipe", "label"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "':'", null, null, null, "'=='", "'='", "'|'", "'<'", "'>'", "'<='", 
		"'>='", "'('", "')'", "'{'", "'}'", "'['", "']'", "'if'", "'do'", "'end'", 
		"'func'", "'return'", "'for'", "'add'", "'sub'", "'div'", "'mult'", "'true'", 
		"'false'", "'or'", "'and'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, "COMMENT", "NEWLINE", "WHITESPACE", "EQUALS", "ASSIGN", "PIPE", 
		"LTHAN", "GTHAN", "LTHAN_EQUALS", "GTHAN_EQUALS", "LPAREN", "RPAREN", 
		"LSQUIG", "RSQUIG", "LBLOCK", "RBLOCK", "IF", "DO", "END", "FUNC", "RETURN", 
		"FOR", "PLUS", "SUB", "DIV", "MULT", "TRUE", "FALSE", "OR", "AND", "INT", 
		"QUOTED", "LETTERS", "STRING"
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
			setState(72);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(50);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
				case 1:
					{
					setState(44);
					match(NEWLINE);
					}
					break;
				case 2:
					{
					setState(46); 
					_errHandler.sync(this);
					_alt = 1;
					do {
						switch (_alt) {
						case 1:
							{
							{
							setState(45);
							match(NEWLINE);
							}
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						setState(48); 
						_errHandler.sync(this);
						_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
					} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
					}
					break;
				}
				setState(58);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
				case 1:
					{
					setState(52);
					code();
					}
					break;
				case 2:
					{
					setState(54); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(53);
						code();
						}
						}
						setState(56); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << NEWLINE) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << IF) | (1L << FUNC) | (1L << RETURN) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << QUOTED) | (1L << LETTERS))) != 0) );
					}
					break;
				}
				setState(60);
				match(EOF);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(68);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
				case 1:
					{
					setState(62);
					code();
					}
					break;
				case 2:
					{
					setState(64); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(63);
						code();
						}
						}
						setState(66); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << NEWLINE) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << IF) | (1L << FUNC) | (1L << RETURN) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << QUOTED) | (1L << LETTERS))) != 0) );
					}
					break;
				}
				setState(70);
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
			setState(76);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case FUNC:
				_localctx = new FuncdefCodeContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(74);
				funcdef();
				}
				break;
			case T__0:
			case NEWLINE:
			case LPAREN:
			case LSQUIG:
			case LBLOCK:
			case IF:
			case RETURN:
			case SUB:
			case TRUE:
			case FALSE:
			case INT:
			case QUOTED:
			case LETTERS:
				_localctx = new StatementCodeContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(75);
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
			setState(108);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(78);
				match(FUNC);
				setState(79);
				label();
				setState(80);
				match(DO);
				setState(81);
				match(END);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(83);
				match(FUNC);
				setState(85); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(84);
					label();
					}
					}
					setState(87); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__0 );
				setState(89);
				match(DO);
				setState(90);
				match(END);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(92);
				match(FUNC);
				setState(93);
				label();
				setState(94);
				match(DO);
				setState(95);
				funcguts();
				setState(96);
				match(END);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(98);
				match(FUNC);
				setState(100); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(99);
					label();
					}
					}
					setState(102); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__0 );
				setState(104);
				match(DO);
				setState(105);
				funcguts();
				setState(106);
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
			setState(111); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(110);
				statement();
				}
				}
				setState(113); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << NEWLINE) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << IF) | (1L << RETURN) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << QUOTED) | (1L << LETTERS))) != 0) );
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
			setState(127);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				_localctx = new AssignStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(115);
				assignstat();
				}
				break;
			case 2:
				_localctx = new ReturnStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(116);
				returnstat();
				}
				break;
			case 3:
				_localctx = new FncallStatementContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(117);
				fncall();
				setState(118);
				match(NEWLINE);
				}
				break;
			case 4:
				_localctx = new IfStatementContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(120);
				ifstat();
				setState(121);
				match(NEWLINE);
				}
				break;
			case 5:
				_localctx = new PipeStatementContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(123);
				pipe();
				setState(124);
				match(NEWLINE);
				}
				break;
			case 6:
				_localctx = new NewlineStatementContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(126);
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
			setState(144);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(129);
				match(IF);
				setState(130);
				conditional(0);
				setState(131);
				match(DO);
				setState(133); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(132);
					statement();
					}
					}
					setState(135); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << NEWLINE) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << IF) | (1L << RETURN) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << QUOTED) | (1L << LETTERS))) != 0) );
				setState(137);
				match(END);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(139);
				match(IF);
				setState(140);
				conditional(0);
				setState(141);
				match(DO);
				setState(142);
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
			setState(167);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(146);
				match(RETURN);
				setState(147);
				term();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(148);
				match(RETURN);
				setState(149);
				term();
				setState(156);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
				case 1:
					{
					setState(150);
					match(NEWLINE);
					}
					break;
				case 2:
					{
					setState(152); 
					_errHandler.sync(this);
					_alt = 1;
					do {
						switch (_alt) {
						case 1:
							{
							{
							setState(151);
							match(NEWLINE);
							}
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						setState(154); 
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
				setState(158);
				match(RETURN);
				setState(165);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
				case 1:
					{
					setState(159);
					match(NEWLINE);
					}
					break;
				case 2:
					{
					setState(161); 
					_errHandler.sync(this);
					_alt = 1;
					do {
						switch (_alt) {
						case 1:
							{
							{
							setState(160);
							match(NEWLINE);
							}
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						setState(163); 
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
			setState(203);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(169);
				match(LSQUIG);
				setState(170);
				match(RSQUIG);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(171);
				match(LSQUIG);
				setState(172);
				match(NEWLINE);
				setState(179);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
				case 1:
					{
					setState(173);
					assignstat();
					}
					break;
				case 2:
					{
					setState(175); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(174);
						assignstat();
						}
						}
						setState(177); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( _la==T__0 );
					}
					break;
				}
				setState(181);
				match(RSQUIG);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(183);
				match(LSQUIG);
				setState(190);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
				case 1:
					{
					setState(184);
					assignstat();
					}
					break;
				case 2:
					{
					setState(186); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(185);
						assignstat();
						}
						}
						setState(188); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( _la==T__0 );
					}
					break;
				}
				setState(192);
				match(RSQUIG);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(194);
				match(LSQUIG);
				setState(195);
				assign();
				setState(196);
				match(RSQUIG);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(198);
				match(LSQUIG);
				setState(199);
				match(NEWLINE);
				setState(200);
				assign();
				setState(201);
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
			setState(205);
			assign();
			setState(206);
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
		enterRule(_localctx, 18, RULE_assign);
		try {
			setState(216);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				_localctx = new ExprAssignContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(208);
				label();
				setState(209);
				match(ASSIGN);
				setState(210);
				term();
				}
				break;
			case 2:
				_localctx = new FncallAssignContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(212);
				label();
				setState(213);
				match(ASSIGN);
				setState(214);
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
			setState(218);
			match(LPAREN);
			setState(219);
			fncall();
			setState(220);
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
			setState(229);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(222);
				label();
				setState(224); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(223);
					term();
					}
					}
					setState(226); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << QUOTED) | (1L << LETTERS))) != 0) );
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(228);
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
		enterRule(_localctx, 24, RULE_term);
		try {
			setState(237);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
			case 1:
				_localctx = new LabelTermContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(231);
				label();
				}
				break;
			case 2:
				_localctx = new LiteralTermContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(232);
				literal();
				}
				break;
			case 3:
				_localctx = new MathTermContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(233);
				math(0);
				}
				break;
			case 4:
				_localctx = new ParenfncallTermContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(234);
				paren_fncall();
				}
				break;
			case 5:
				_localctx = new StructTermContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(235);
				structexpr();
				}
				break;
			case 6:
				_localctx = new ListTermContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(236);
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
			setState(258);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(239);
				match(LBLOCK);
				setState(240);
				match(RBLOCK);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(241);
				match(LBLOCK);
				setState(243); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(242);
					listterm();
					}
					}
					setState(245); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << QUOTED) | (1L << LETTERS))) != 0) );
				setState(247);
				match(RBLOCK);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(249);
				match(LBLOCK);
				setState(250);
				match(NEWLINE);
				setState(252); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(251);
					listterm();
					}
					}
					setState(254); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << INT) | (1L << QUOTED) | (1L << LETTERS))) != 0) );
				setState(256);
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
			setState(264);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(260);
				term();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(261);
				term();
				setState(262);
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
			setState(288);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				{
				_localctx = new EqualsCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(267);
				term();
				setState(268);
				match(EQUALS);
				setState(269);
				term();
				}
				break;
			case 2:
				{
				_localctx = new LthanCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(271);
				term();
				setState(272);
				match(LTHAN);
				setState(273);
				term();
				}
				break;
			case 3:
				{
				_localctx = new GthanCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(275);
				term();
				setState(276);
				match(GTHAN);
				setState(277);
				term();
				}
				break;
			case 4:
				{
				_localctx = new LthanequalsCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(279);
				term();
				setState(280);
				match(LTHAN_EQUALS);
				setState(281);
				term();
				}
				break;
			case 5:
				{
				_localctx = new GthanequalsCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(283);
				term();
				setState(284);
				match(GTHAN_EQUALS);
				setState(285);
				term();
				}
				break;
			case 6:
				{
				_localctx = new BoolCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(287);
				boolexpr();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(298);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,35,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(296);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
					case 1:
						{
						_localctx = new OrCondContext(new ConditionalContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_conditional);
						setState(290);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(291);
						match(OR);
						setState(292);
						conditional(4);
						}
						break;
					case 2:
						{
						_localctx = new AndCondContext(new ConditionalContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_conditional);
						setState(293);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(294);
						match(AND);
						setState(295);
						conditional(3);
						}
						break;
					}
					} 
				}
				setState(300);
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
			setState(305);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SUB:
				{
				_localctx = new UnarySubMathContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(302);
				match(SUB);
				setState(303);
				math(2);
				}
				break;
			case INT:
				{
				_localctx = new SoloMathContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(304);
				num();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(312);
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
					setState(307);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(308);
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
					setState(309);
					math(4);
					}
					} 
				}
				setState(314);
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
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	 
		public LiteralContext() { }
		public void copyFrom(LiteralContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class QuotedLiteralContext extends LiteralContext {
		public TerminalNode QUOTED() { return getToken(BeepBoopParser.QUOTED, 0); }
		public QuotedLiteralContext(LiteralContext ctx) { copyFrom(ctx); }
	}
	public static class NumLiteralContext extends LiteralContext {
		public NumContext num() {
			return getRuleContext(NumContext.class,0);
		}
		public NumLiteralContext(LiteralContext ctx) { copyFrom(ctx); }
	}
	public static class LettersLiteralContext extends LiteralContext {
		public TerminalNode LETTERS() { return getToken(BeepBoopParser.LETTERS, 0); }
		public LettersLiteralContext(LiteralContext ctx) { copyFrom(ctx); }
	}
	public static class BoolLiteralContext extends LiteralContext {
		public BoolexprContext boolexpr() {
			return getRuleContext(BoolexprContext.class,0);
		}
		public BoolLiteralContext(LiteralContext ctx) { copyFrom(ctx); }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_literal);
		try {
			setState(319);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				_localctx = new NumLiteralContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(315);
				num();
				}
				break;
			case LETTERS:
				_localctx = new LettersLiteralContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(316);
				match(LETTERS);
				}
				break;
			case TRUE:
			case FALSE:
				_localctx = new BoolLiteralContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(317);
				boolexpr();
				}
				break;
			case QUOTED:
				_localctx = new QuotedLiteralContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(318);
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
			setState(321);
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
			setState(323);
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
			setState(337);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,39,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(325);
				term();
				setState(326);
				match(PIPE);
				setState(327);
				fncall();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(329);
				term();
				setState(330);
				match(PIPE);
				setState(331);
				pipe();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(333);
				fncall();
				setState(334);
				match(PIPE);
				setState(335);
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

	public static class LabelContext extends ParserRuleContext {
		public TerminalNode LETTERS() { return getToken(BeepBoopParser.LETTERS, 0); }
		public LabelContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_label; }
	}

	public final LabelContext label() throws RecognitionException {
		LabelContext _localctx = new LabelContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_label);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(339);
			match(T__0);
			setState(340);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3%\u0159\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\3\2\3\2\6\2\61\n\2"+
		"\r\2\16\2\62\5\2\65\n\2\3\2\3\2\6\29\n\2\r\2\16\2:\5\2=\n\2\3\2\3\2\3"+
		"\2\3\2\6\2C\n\2\r\2\16\2D\5\2G\n\2\3\2\3\2\5\2K\n\2\3\3\3\3\5\3O\n\3\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\6\4X\n\4\r\4\16\4Y\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\6\4g\n\4\r\4\16\4h\3\4\3\4\3\4\3\4\5\4o\n\4\3\5\6"+
		"\5r\n\5\r\5\16\5s\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6"+
		"\u0082\n\6\3\7\3\7\3\7\3\7\6\7\u0088\n\7\r\7\16\7\u0089\3\7\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\5\7\u0093\n\7\3\b\3\b\3\b\3\b\3\b\3\b\6\b\u009b\n\b\r\b"+
		"\16\b\u009c\5\b\u009f\n\b\3\b\3\b\3\b\6\b\u00a4\n\b\r\b\16\b\u00a5\5\b"+
		"\u00a8\n\b\5\b\u00aa\n\b\3\t\3\t\3\t\3\t\3\t\3\t\6\t\u00b2\n\t\r\t\16"+
		"\t\u00b3\5\t\u00b6\n\t\3\t\3\t\3\t\3\t\3\t\6\t\u00bd\n\t\r\t\16\t\u00be"+
		"\5\t\u00c1\n\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u00ce\n"+
		"\t\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13\u00db\n\13"+
		"\3\f\3\f\3\f\3\f\3\r\3\r\6\r\u00e3\n\r\r\r\16\r\u00e4\3\r\5\r\u00e8\n"+
		"\r\3\16\3\16\3\16\3\16\3\16\3\16\5\16\u00f0\n\16\3\17\3\17\3\17\3\17\6"+
		"\17\u00f6\n\17\r\17\16\17\u00f7\3\17\3\17\3\17\3\17\3\17\6\17\u00ff\n"+
		"\17\r\17\16\17\u0100\3\17\3\17\5\17\u0105\n\17\3\20\3\20\3\20\3\20\5\20"+
		"\u010b\n\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\5\21\u0123\n\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\7\21\u012b\n\21\f\21\16\21\u012e\13\21\3\22"+
		"\3\22\3\22\3\22\5\22\u0134\n\22\3\22\3\22\3\22\7\22\u0139\n\22\f\22\16"+
		"\22\u013c\13\22\3\23\3\23\3\23\3\23\5\23\u0142\n\23\3\24\3\24\3\25\3\25"+
		"\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\5\26\u0154"+
		"\n\26\3\27\3\27\3\27\3\27\2\4 \"\30\2\4\6\b\n\f\16\20\22\24\26\30\32\34"+
		"\36 \"$&(*,\2\4\3\2\32\33\3\2\36\37\2\u0180\2J\3\2\2\2\4N\3\2\2\2\6n\3"+
		"\2\2\2\bq\3\2\2\2\n\u0081\3\2\2\2\f\u0092\3\2\2\2\16\u00a9\3\2\2\2\20"+
		"\u00cd\3\2\2\2\22\u00cf\3\2\2\2\24\u00da\3\2\2\2\26\u00dc\3\2\2\2\30\u00e7"+
		"\3\2\2\2\32\u00ef\3\2\2\2\34\u0104\3\2\2\2\36\u010a\3\2\2\2 \u0122\3\2"+
		"\2\2\"\u0133\3\2\2\2$\u0141\3\2\2\2&\u0143\3\2\2\2(\u0145\3\2\2\2*\u0153"+
		"\3\2\2\2,\u0155\3\2\2\2.\65\7\5\2\2/\61\7\5\2\2\60/\3\2\2\2\61\62\3\2"+
		"\2\2\62\60\3\2\2\2\62\63\3\2\2\2\63\65\3\2\2\2\64.\3\2\2\2\64\60\3\2\2"+
		"\2\65<\3\2\2\2\66=\5\4\3\2\679\5\4\3\28\67\3\2\2\29:\3\2\2\2:8\3\2\2\2"+
		":;\3\2\2\2;=\3\2\2\2<\66\3\2\2\2<8\3\2\2\2=>\3\2\2\2>?\7\2\2\3?K\3\2\2"+
		"\2@G\5\4\3\2AC\5\4\3\2BA\3\2\2\2CD\3\2\2\2DB\3\2\2\2DE\3\2\2\2EG\3\2\2"+
		"\2F@\3\2\2\2FB\3\2\2\2GH\3\2\2\2HI\7\2\2\3IK\3\2\2\2J\64\3\2\2\2JF\3\2"+
		"\2\2K\3\3\2\2\2LO\5\6\4\2MO\5\n\6\2NL\3\2\2\2NM\3\2\2\2O\5\3\2\2\2PQ\7"+
		"\27\2\2QR\5,\27\2RS\7\25\2\2ST\7\26\2\2To\3\2\2\2UW\7\27\2\2VX\5,\27\2"+
		"WV\3\2\2\2XY\3\2\2\2YW\3\2\2\2YZ\3\2\2\2Z[\3\2\2\2[\\\7\25\2\2\\]\7\26"+
		"\2\2]o\3\2\2\2^_\7\27\2\2_`\5,\27\2`a\7\25\2\2ab\5\b\5\2bc\7\26\2\2co"+
		"\3\2\2\2df\7\27\2\2eg\5,\27\2fe\3\2\2\2gh\3\2\2\2hf\3\2\2\2hi\3\2\2\2"+
		"ij\3\2\2\2jk\7\25\2\2kl\5\b\5\2lm\7\26\2\2mo\3\2\2\2nP\3\2\2\2nU\3\2\2"+
		"\2n^\3\2\2\2nd\3\2\2\2o\7\3\2\2\2pr\5\n\6\2qp\3\2\2\2rs\3\2\2\2sq\3\2"+
		"\2\2st\3\2\2\2t\t\3\2\2\2u\u0082\5\22\n\2v\u0082\5\16\b\2wx\5\30\r\2x"+
		"y\7\5\2\2y\u0082\3\2\2\2z{\5\f\7\2{|\7\5\2\2|\u0082\3\2\2\2}~\5*\26\2"+
		"~\177\7\5\2\2\177\u0082\3\2\2\2\u0080\u0082\7\5\2\2\u0081u\3\2\2\2\u0081"+
		"v\3\2\2\2\u0081w\3\2\2\2\u0081z\3\2\2\2\u0081}\3\2\2\2\u0081\u0080\3\2"+
		"\2\2\u0082\13\3\2\2\2\u0083\u0084\7\24\2\2\u0084\u0085\5 \21\2\u0085\u0087"+
		"\7\25\2\2\u0086\u0088\5\n\6\2\u0087\u0086\3\2\2\2\u0088\u0089\3\2\2\2"+
		"\u0089\u0087\3\2\2\2\u0089\u008a\3\2\2\2\u008a\u008b\3\2\2\2\u008b\u008c"+
		"\7\26\2\2\u008c\u0093\3\2\2\2\u008d\u008e\7\24\2\2\u008e\u008f\5 \21\2"+
		"\u008f\u0090\7\25\2\2\u0090\u0091\7\26\2\2\u0091\u0093\3\2\2\2\u0092\u0083"+
		"\3\2\2\2\u0092\u008d\3\2\2\2\u0093\r\3\2\2\2\u0094\u0095\7\30\2\2\u0095"+
		"\u00aa\5\32\16\2\u0096\u0097\7\30\2\2\u0097\u009e\5\32\16\2\u0098\u009f"+
		"\7\5\2\2\u0099\u009b\7\5\2\2\u009a\u0099\3\2\2\2\u009b\u009c\3\2\2\2\u009c"+
		"\u009a\3\2\2\2\u009c\u009d\3\2\2\2\u009d\u009f\3\2\2\2\u009e\u0098\3\2"+
		"\2\2\u009e\u009a\3\2\2\2\u009f\u00aa\3\2\2\2\u00a0\u00a7\7\30\2\2\u00a1"+
		"\u00a8\7\5\2\2\u00a2\u00a4\7\5\2\2\u00a3\u00a2\3\2\2\2\u00a4\u00a5\3\2"+
		"\2\2\u00a5\u00a3\3\2\2\2\u00a5\u00a6\3\2\2\2\u00a6\u00a8\3\2\2\2\u00a7"+
		"\u00a1\3\2\2\2\u00a7\u00a3\3\2\2\2\u00a8\u00aa\3\2\2\2\u00a9\u0094\3\2"+
		"\2\2\u00a9\u0096\3\2\2\2\u00a9\u00a0\3\2\2\2\u00aa\17\3\2\2\2\u00ab\u00ac"+
		"\7\20\2\2\u00ac\u00ce\7\21\2\2\u00ad\u00ae\7\20\2\2\u00ae\u00b5\7\5\2"+
		"\2\u00af\u00b6\5\22\n\2\u00b0\u00b2\5\22\n\2\u00b1\u00b0\3\2\2\2\u00b2"+
		"\u00b3\3\2\2\2\u00b3\u00b1\3\2\2\2\u00b3\u00b4\3\2\2\2\u00b4\u00b6\3\2"+
		"\2\2\u00b5\u00af\3\2\2\2\u00b5\u00b1\3\2\2\2\u00b6\u00b7\3\2\2\2\u00b7"+
		"\u00b8\7\21\2\2\u00b8\u00ce\3\2\2\2\u00b9\u00c0\7\20\2\2\u00ba\u00c1\5"+
		"\22\n\2\u00bb\u00bd\5\22\n\2\u00bc\u00bb\3\2\2\2\u00bd\u00be\3\2\2\2\u00be"+
		"\u00bc\3\2\2\2\u00be\u00bf\3\2\2\2\u00bf\u00c1\3\2\2\2\u00c0\u00ba\3\2"+
		"\2\2\u00c0\u00bc\3\2\2\2\u00c1\u00c2\3\2\2\2\u00c2\u00c3\7\21\2\2\u00c3"+
		"\u00ce\3\2\2\2\u00c4\u00c5\7\20\2\2\u00c5\u00c6\5\24\13\2\u00c6\u00c7"+
		"\7\21\2\2\u00c7\u00ce\3\2\2\2\u00c8\u00c9\7\20\2\2\u00c9\u00ca\7\5\2\2"+
		"\u00ca\u00cb\5\24\13\2\u00cb\u00cc\7\21\2\2\u00cc\u00ce\3\2\2\2\u00cd"+
		"\u00ab\3\2\2\2\u00cd\u00ad\3\2\2\2\u00cd\u00b9\3\2\2\2\u00cd\u00c4\3\2"+
		"\2\2\u00cd\u00c8\3\2\2\2\u00ce\21\3\2\2\2\u00cf\u00d0\5\24\13\2\u00d0"+
		"\u00d1\7\5\2\2\u00d1\23\3\2\2\2\u00d2\u00d3\5,\27\2\u00d3\u00d4\7\b\2"+
		"\2\u00d4\u00d5\5\32\16\2\u00d5\u00db\3\2\2\2\u00d6\u00d7\5,\27\2\u00d7"+
		"\u00d8\7\b\2\2\u00d8\u00d9\5\30\r\2\u00d9\u00db\3\2\2\2\u00da\u00d2\3"+
		"\2\2\2\u00da\u00d6\3\2\2\2\u00db\25\3\2\2\2\u00dc\u00dd\7\16\2\2\u00dd"+
		"\u00de\5\30\r\2\u00de\u00df\7\17\2\2\u00df\27\3\2\2\2\u00e0\u00e2\5,\27"+
		"\2\u00e1\u00e3\5\32\16\2\u00e2\u00e1\3\2\2\2\u00e3\u00e4\3\2\2\2\u00e4"+
		"\u00e2\3\2\2\2\u00e4\u00e5\3\2\2\2\u00e5\u00e8\3\2\2\2\u00e6\u00e8\5,"+
		"\27\2\u00e7\u00e0\3\2\2\2\u00e7\u00e6\3\2\2\2\u00e8\31\3\2\2\2\u00e9\u00f0"+
		"\5,\27\2\u00ea\u00f0\5$\23\2\u00eb\u00f0\5\"\22\2\u00ec\u00f0\5\26\f\2"+
		"\u00ed\u00f0\5\20\t\2\u00ee\u00f0\5\34\17\2\u00ef\u00e9\3\2\2\2\u00ef"+
		"\u00ea\3\2\2\2\u00ef\u00eb\3\2\2\2\u00ef\u00ec\3\2\2\2\u00ef\u00ed\3\2"+
		"\2\2\u00ef\u00ee\3\2\2\2\u00f0\33\3\2\2\2\u00f1\u00f2\7\22\2\2\u00f2\u0105"+
		"\7\23\2\2\u00f3\u00f5\7\22\2\2\u00f4\u00f6\5\36\20\2\u00f5\u00f4\3\2\2"+
		"\2\u00f6\u00f7\3\2\2\2\u00f7\u00f5\3\2\2\2\u00f7\u00f8\3\2\2\2\u00f8\u00f9"+
		"\3\2\2\2\u00f9\u00fa\7\23\2\2\u00fa\u0105\3\2\2\2\u00fb\u00fc\7\22\2\2"+
		"\u00fc\u00fe\7\5\2\2\u00fd\u00ff\5\36\20\2\u00fe\u00fd\3\2\2\2\u00ff\u0100"+
		"\3\2\2\2\u0100\u00fe\3\2\2\2\u0100\u0101\3\2\2\2\u0101\u0102\3\2\2\2\u0102"+
		"\u0103\7\23\2\2\u0103\u0105\3\2\2\2\u0104\u00f1\3\2\2\2\u0104\u00f3\3"+
		"\2\2\2\u0104\u00fb\3\2\2\2\u0105\35\3\2\2\2\u0106\u010b\5\32\16\2\u0107"+
		"\u0108\5\32\16\2\u0108\u0109\7\5\2\2\u0109\u010b\3\2\2\2\u010a\u0106\3"+
		"\2\2\2\u010a\u0107\3\2\2\2\u010b\37\3\2\2\2\u010c\u010d\b\21\1\2\u010d"+
		"\u010e\5\32\16\2\u010e\u010f\7\7\2\2\u010f\u0110\5\32\16\2\u0110\u0123"+
		"\3\2\2\2\u0111\u0112\5\32\16\2\u0112\u0113\7\n\2\2\u0113\u0114\5\32\16"+
		"\2\u0114\u0123\3\2\2\2\u0115\u0116\5\32\16\2\u0116\u0117\7\13\2\2\u0117"+
		"\u0118\5\32\16\2\u0118\u0123\3\2\2\2\u0119\u011a\5\32\16\2\u011a\u011b"+
		"\7\f\2\2\u011b\u011c\5\32\16\2\u011c\u0123\3\2\2\2\u011d\u011e\5\32\16"+
		"\2\u011e\u011f\7\r\2\2\u011f\u0120\5\32\16\2\u0120\u0123\3\2\2\2\u0121"+
		"\u0123\5(\25\2\u0122\u010c\3\2\2\2\u0122\u0111\3\2\2\2\u0122\u0115\3\2"+
		"\2\2\u0122\u0119\3\2\2\2\u0122\u011d\3\2\2\2\u0122\u0121\3\2\2\2\u0123"+
		"\u012c\3\2\2\2\u0124\u0125\f\5\2\2\u0125\u0126\7 \2\2\u0126\u012b\5 \21"+
		"\6\u0127\u0128\f\4\2\2\u0128\u0129\7!\2\2\u0129\u012b\5 \21\5\u012a\u0124"+
		"\3\2\2\2\u012a\u0127\3\2\2\2\u012b\u012e\3\2\2\2\u012c\u012a\3\2\2\2\u012c"+
		"\u012d\3\2\2\2\u012d!\3\2\2\2\u012e\u012c\3\2\2\2\u012f\u0130\b\22\1\2"+
		"\u0130\u0131\7\33\2\2\u0131\u0134\5\"\22\4\u0132\u0134\5&\24\2\u0133\u012f"+
		"\3\2\2\2\u0133\u0132\3\2\2\2\u0134\u013a\3\2\2\2\u0135\u0136\f\5\2\2\u0136"+
		"\u0137\t\2\2\2\u0137\u0139\5\"\22\6\u0138\u0135\3\2\2\2\u0139\u013c\3"+
		"\2\2\2\u013a\u0138\3\2\2\2\u013a\u013b\3\2\2\2\u013b#\3\2\2\2\u013c\u013a"+
		"\3\2\2\2\u013d\u0142\5&\24\2\u013e\u0142\7$\2\2\u013f\u0142\5(\25\2\u0140"+
		"\u0142\7#\2\2\u0141\u013d\3\2\2\2\u0141\u013e\3\2\2\2\u0141\u013f\3\2"+
		"\2\2\u0141\u0140\3\2\2\2\u0142%\3\2\2\2\u0143\u0144\7\"\2\2\u0144\'\3"+
		"\2\2\2\u0145\u0146\t\3\2\2\u0146)\3\2\2\2\u0147\u0148\5\32\16\2\u0148"+
		"\u0149\7\t\2\2\u0149\u014a\5\30\r\2\u014a\u0154\3\2\2\2\u014b\u014c\5"+
		"\32\16\2\u014c\u014d\7\t\2\2\u014d\u014e\5*\26\2\u014e\u0154\3\2\2\2\u014f"+
		"\u0150\5\30\r\2\u0150\u0151\7\t\2\2\u0151\u0152\5*\26\2\u0152\u0154\3"+
		"\2\2\2\u0153\u0147\3\2\2\2\u0153\u014b\3\2\2\2\u0153\u014f\3\2\2\2\u0154"+
		"+\3\2\2\2\u0155\u0156\7\3\2\2\u0156\u0157\7$\2\2\u0157-\3\2\2\2*\62\64"+
		":<DFJNYhns\u0081\u0089\u0092\u009c\u009e\u00a5\u00a7\u00a9\u00b3\u00b5"+
		"\u00be\u00c0\u00cd\u00da\u00e4\u00e7\u00ef\u00f7\u0100\u0104\u010a\u0122"+
		"\u012a\u012c\u0133\u013a\u0141\u0153";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}