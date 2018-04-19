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
		T__0=1, COMMENT=2, NEWLINE=3, WHITESPACE=4, IF=5, DO=6, END=7, FUNC=8, 
		RETURN=9, FOR=10, IS=11, PLUS=12, MINUS=13, MULT=14, DIVIDE=15, ASSIGN=16, 
		PIPE=17, LPAREN=18, RPAREN=19, INT=20, STRING=21;
	public static final int
		RULE_beepboop = 0, RULE_code = 1, RULE_statement = 2, RULE_funcguts = 3, 
		RULE_funcdef = 4, RULE_fncall = 5, RULE_ifstat = 6, RULE_returnStat = 7, 
		RULE_assignment = 8, RULE_expr = 9, RULE_pipe = 10, RULE_term = 11, RULE_label = 12;
	public static final String[] ruleNames = {
		"beepboop", "code", "statement", "funcguts", "funcdef", "fncall", "ifstat", 
		"returnStat", "assignment", "expr", "pipe", "term", "label"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'$'", null, null, null, "'if'", "'do'", "'end'", "'func'", "'return'", 
		"'for'", "'is'", "'+'", "'-'", "'*'", "'/'", "'='", "'|'", "'('", "')'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, "COMMENT", "NEWLINE", "WHITESPACE", "IF", "DO", "END", "FUNC", 
		"RETURN", "FOR", "IS", "PLUS", "MINUS", "MULT", "DIVIDE", "ASSIGN", "PIPE", 
		"LPAREN", "RPAREN", "INT", "STRING"
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
	public static class BeepboopContext extends ParserRuleContext {
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
		public BeepboopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_beepboop; }
	}

	public final BeepboopContext beepboop() throws RecognitionException {
		BeepboopContext _localctx = new BeepboopContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_beepboop);
		int _la;
		try {
			int _alt;
			setState(45);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(27); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(26);
						match(NEWLINE);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(29); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				setState(32); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(31);
					code();
					}
					}
					setState(34); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << NEWLINE) | (1L << FUNC) | (1L << RETURN) | (1L << STRING))) != 0) );
				setState(36);
				match(EOF);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(39); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(38);
					code();
					}
					}
					setState(41); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << NEWLINE) | (1L << FUNC) | (1L << RETURN) | (1L << STRING))) != 0) );
				setState(43);
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
			setState(49);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
			case NEWLINE:
			case RETURN:
			case STRING:
				_localctx = new StatementCodeContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(47);
				statement();
				}
				break;
			case FUNC:
				_localctx = new FuncdefCodeContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(48);
				funcdef();
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
		public AssignmentContext assignment() {
			return getRuleContext(AssignmentContext.class,0);
		}
		public TerminalNode NEWLINE() { return getToken(BeepBoopParser.NEWLINE, 0); }
		public AssignStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class NoopStatementContext extends StatementContext {
		public TerminalNode NEWLINE() { return getToken(BeepBoopParser.NEWLINE, 0); }
		public NoopStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class ReturnStatementContext extends StatementContext {
		public ReturnStatContext returnStat() {
			return getRuleContext(ReturnStatContext.class,0);
		}
		public TerminalNode NEWLINE() { return getToken(BeepBoopParser.NEWLINE, 0); }
		public ReturnStatementContext(StatementContext ctx) { copyFrom(ctx); }
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
		enterRule(_localctx, 4, RULE_statement);
		try {
			setState(64);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				_localctx = new FncallStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(51);
				fncall();
				setState(52);
				match(NEWLINE);
				}
				break;
			case 2:
				_localctx = new AssignStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(54);
				assignment();
				setState(55);
				match(NEWLINE);
				}
				break;
			case 3:
				_localctx = new ReturnStatementContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(57);
				returnStat();
				setState(58);
				match(NEWLINE);
				}
				break;
			case 4:
				_localctx = new PipeStatementContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(60);
				pipe();
				setState(61);
				match(NEWLINE);
				}
				break;
			case 5:
				_localctx = new NoopStatementContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(63);
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
			setState(67); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(66);
				statement();
				}
				}
				setState(69); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << NEWLINE) | (1L << RETURN) | (1L << STRING))) != 0) );
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
		public TerminalNode STRING() { return getToken(BeepBoopParser.STRING, 0); }
		public TerminalNode DO() { return getToken(BeepBoopParser.DO, 0); }
		public FuncgutsContext funcguts() {
			return getRuleContext(FuncgutsContext.class,0);
		}
		public TerminalNode END() { return getToken(BeepBoopParser.END, 0); }
		public TerminalNode NEWLINE() { return getToken(BeepBoopParser.NEWLINE, 0); }
		public List<LabelContext> label() {
			return getRuleContexts(LabelContext.class);
		}
		public LabelContext label(int i) {
			return getRuleContext(LabelContext.class,i);
		}
		public FuncdefContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcdef; }
	}

	public final FuncdefContext funcdef() throws RecognitionException {
		FuncdefContext _localctx = new FuncdefContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_funcdef);
		int _la;
		try {
			setState(95);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(71);
				match(FUNC);
				setState(72);
				match(STRING);
				setState(74); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(73);
					label();
					}
					}
					setState(76); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__0 );
				setState(78);
				match(DO);
				setState(79);
				funcguts();
				setState(80);
				match(END);
				setState(81);
				match(NEWLINE);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(83);
				match(FUNC);
				setState(84);
				match(STRING);
				setState(85);
				match(DO);
				setState(86);
				funcguts();
				setState(87);
				match(END);
				setState(88);
				match(NEWLINE);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(90);
				match(FUNC);
				setState(91);
				match(STRING);
				setState(92);
				match(DO);
				setState(93);
				match(END);
				setState(94);
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

	public static class FncallContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(BeepBoopParser.STRING, 0); }
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
		enterRule(_localctx, 10, RULE_fncall);
		int _la;
		try {
			setState(104);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(97);
				match(STRING);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(98);
				match(STRING);
				setState(100); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(99);
					term();
					}
					}
					setState(102); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << INT) | (1L << STRING))) != 0) );
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
		public IfstatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifstat; }
	 
		public IfstatContext() { }
		public void copyFrom(IfstatContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class ExprIfStatementContext extends IfstatContext {
		public TerminalNode IF() { return getToken(BeepBoopParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode DO() { return getToken(BeepBoopParser.DO, 0); }
		public CodeContext code() {
			return getRuleContext(CodeContext.class,0);
		}
		public TerminalNode END() { return getToken(BeepBoopParser.END, 0); }
		public ExprIfStatementContext(IfstatContext ctx) { copyFrom(ctx); }
	}
	public static class FncallIfStatementContext extends IfstatContext {
		public TerminalNode IF() { return getToken(BeepBoopParser.IF, 0); }
		public FncallContext fncall() {
			return getRuleContext(FncallContext.class,0);
		}
		public TerminalNode DO() { return getToken(BeepBoopParser.DO, 0); }
		public CodeContext code() {
			return getRuleContext(CodeContext.class,0);
		}
		public TerminalNode END() { return getToken(BeepBoopParser.END, 0); }
		public FncallIfStatementContext(IfstatContext ctx) { copyFrom(ctx); }
	}

	public final IfstatContext ifstat() throws RecognitionException {
		IfstatContext _localctx = new IfstatContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_ifstat);
		try {
			setState(118);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				_localctx = new ExprIfStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(106);
				match(IF);
				setState(107);
				expr(0);
				setState(108);
				match(DO);
				setState(109);
				code();
				setState(110);
				match(END);
				}
				break;
			case 2:
				_localctx = new FncallIfStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(112);
				match(IF);
				setState(113);
				fncall();
				setState(114);
				match(DO);
				setState(115);
				code();
				setState(116);
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

	public static class ReturnStatContext extends ParserRuleContext {
		public ReturnStatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnStat; }
	 
		public ReturnStatContext() { }
		public void copyFrom(ReturnStatContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class FncallReturnContext extends ReturnStatContext {
		public TerminalNode RETURN() { return getToken(BeepBoopParser.RETURN, 0); }
		public FncallContext fncall() {
			return getRuleContext(FncallContext.class,0);
		}
		public FncallReturnContext(ReturnStatContext ctx) { copyFrom(ctx); }
	}
	public static class ExprReturnContext extends ReturnStatContext {
		public TerminalNode RETURN() { return getToken(BeepBoopParser.RETURN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ExprReturnContext(ReturnStatContext ctx) { copyFrom(ctx); }
	}

	public final ReturnStatContext returnStat() throws RecognitionException {
		ReturnStatContext _localctx = new ReturnStatContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_returnStat);
		try {
			setState(124);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				_localctx = new ExprReturnContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(120);
				match(RETURN);
				setState(121);
				expr(0);
				}
				break;
			case 2:
				_localctx = new FncallReturnContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(122);
				match(RETURN);
				setState(123);
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

	public static class AssignmentContext extends ParserRuleContext {
		public LabelContext label() {
			return getRuleContext(LabelContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(BeepBoopParser.ASSIGN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public FncallContext fncall() {
			return getRuleContext(FncallContext.class,0);
		}
		public AssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment; }
	}

	public final AssignmentContext assignment() throws RecognitionException {
		AssignmentContext _localctx = new AssignmentContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_assignment);
		try {
			setState(134);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(126);
				label();
				setState(127);
				match(ASSIGN);
				setState(128);
				expr(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(130);
				label();
				setState(131);
				match(ASSIGN);
				setState(132);
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

	public static class ExprContext extends ParserRuleContext {
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	 
		public ExprContext() { }
		public void copyFrom(ExprContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class UnaryMinusExprContext extends ExprContext {
		public TerminalNode MINUS() { return getToken(BeepBoopParser.MINUS, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public UnaryMinusExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class TermExprContext extends ExprContext {
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public TermExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class AdditiveExprContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode PLUS() { return getToken(BeepBoopParser.PLUS, 0); }
		public TerminalNode MINUS() { return getToken(BeepBoopParser.MINUS, 0); }
		public AdditiveExprContext(ExprContext ctx) { copyFrom(ctx); }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 18;
		enterRecursionRule(_localctx, 18, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(140);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MINUS:
				{
				_localctx = new UnaryMinusExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(137);
				match(MINUS);
				setState(138);
				expr(2);
				}
				break;
			case T__0:
			case INT:
			case STRING:
				{
				_localctx = new TermExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(139);
				term();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(147);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new AdditiveExprContext(new ExprContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_expr);
					setState(142);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(143);
					((AdditiveExprContext)_localctx).op = _input.LT(1);
					_la = _input.LA(1);
					if ( !(_la==PLUS || _la==MINUS) ) {
						((AdditiveExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					setState(144);
					expr(4);
					}
					} 
				}
				setState(149);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
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

	public static class PipeContext extends ParserRuleContext {
		public List<FncallContext> fncall() {
			return getRuleContexts(FncallContext.class);
		}
		public FncallContext fncall(int i) {
			return getRuleContext(FncallContext.class,i);
		}
		public TerminalNode PIPE() { return getToken(BeepBoopParser.PIPE, 0); }
		public List<TerminalNode> NEWLINE() { return getTokens(BeepBoopParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(BeepBoopParser.NEWLINE, i);
		}
		public PipeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pipe; }
	}

	public final PipeContext pipe() throws RecognitionException {
		PipeContext _localctx = new PipeContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_pipe);
		try {
			int _alt;
			setState(160);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(150);
				fncall();
				setState(151);
				match(PIPE);
				setState(152);
				fncall();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(154);
				fncall();
				setState(156); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(155);
						match(NEWLINE);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(158); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
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
	public static class StringTermContext extends TermContext {
		public TerminalNode STRING() { return getToken(BeepBoopParser.STRING, 0); }
		public StringTermContext(TermContext ctx) { copyFrom(ctx); }
	}
	public static class LabelTermContext extends TermContext {
		public LabelContext label() {
			return getRuleContext(LabelContext.class,0);
		}
		public LabelTermContext(TermContext ctx) { copyFrom(ctx); }
	}
	public static class IntTermContext extends TermContext {
		public TerminalNode INT() { return getToken(BeepBoopParser.INT, 0); }
		public IntTermContext(TermContext ctx) { copyFrom(ctx); }
	}

	public final TermContext term() throws RecognitionException {
		TermContext _localctx = new TermContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_term);
		try {
			setState(165);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				_localctx = new LabelTermContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(162);
				label();
				}
				break;
			case STRING:
				_localctx = new StringTermContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(163);
				match(STRING);
				}
				break;
			case INT:
				_localctx = new IntTermContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(164);
				match(INT);
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

	public static class LabelContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(BeepBoopParser.STRING, 0); }
		public LabelContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_label; }
	}

	public final LabelContext label() throws RecognitionException {
		LabelContext _localctx = new LabelContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_label);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(167);
			match(T__0);
			setState(168);
			match(STRING);
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
		case 9:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 3);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\27\u00ad\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\3\2\6\2\36\n\2\r\2\16\2\37\3\2\6\2#\n"+
		"\2\r\2\16\2$\3\2\3\2\3\2\6\2*\n\2\r\2\16\2+\3\2\3\2\5\2\60\n\2\3\3\3\3"+
		"\5\3\64\n\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4C\n"+
		"\4\3\5\6\5F\n\5\r\5\16\5G\3\6\3\6\3\6\6\6M\n\6\r\6\16\6N\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6b\n\6\3\7\3"+
		"\7\3\7\6\7g\n\7\r\7\16\7h\5\7k\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\5\by\n\b\3\t\3\t\3\t\3\t\5\t\177\n\t\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\5\n\u0089\n\n\3\13\3\13\3\13\3\13\5\13\u008f\n\13\3\13\3"+
		"\13\3\13\7\13\u0094\n\13\f\13\16\13\u0097\13\13\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\6\f\u009f\n\f\r\f\16\f\u00a0\5\f\u00a3\n\f\3\r\3\r\3\r\5\r\u00a8\n"+
		"\r\3\16\3\16\3\16\3\16\2\3\24\17\2\4\6\b\n\f\16\20\22\24\26\30\32\2\3"+
		"\3\2\16\17\2\u00b7\2/\3\2\2\2\4\63\3\2\2\2\6B\3\2\2\2\bE\3\2\2\2\na\3"+
		"\2\2\2\fj\3\2\2\2\16x\3\2\2\2\20~\3\2\2\2\22\u0088\3\2\2\2\24\u008e\3"+
		"\2\2\2\26\u00a2\3\2\2\2\30\u00a7\3\2\2\2\32\u00a9\3\2\2\2\34\36\7\5\2"+
		"\2\35\34\3\2\2\2\36\37\3\2\2\2\37\35\3\2\2\2\37 \3\2\2\2 \"\3\2\2\2!#"+
		"\5\4\3\2\"!\3\2\2\2#$\3\2\2\2$\"\3\2\2\2$%\3\2\2\2%&\3\2\2\2&\'\7\2\2"+
		"\3\'\60\3\2\2\2(*\5\4\3\2)(\3\2\2\2*+\3\2\2\2+)\3\2\2\2+,\3\2\2\2,-\3"+
		"\2\2\2-.\7\2\2\3.\60\3\2\2\2/\35\3\2\2\2/)\3\2\2\2\60\3\3\2\2\2\61\64"+
		"\5\6\4\2\62\64\5\n\6\2\63\61\3\2\2\2\63\62\3\2\2\2\64\5\3\2\2\2\65\66"+
		"\5\f\7\2\66\67\7\5\2\2\67C\3\2\2\289\5\22\n\29:\7\5\2\2:C\3\2\2\2;<\5"+
		"\20\t\2<=\7\5\2\2=C\3\2\2\2>?\5\26\f\2?@\7\5\2\2@C\3\2\2\2AC\7\5\2\2B"+
		"\65\3\2\2\2B8\3\2\2\2B;\3\2\2\2B>\3\2\2\2BA\3\2\2\2C\7\3\2\2\2DF\5\6\4"+
		"\2ED\3\2\2\2FG\3\2\2\2GE\3\2\2\2GH\3\2\2\2H\t\3\2\2\2IJ\7\n\2\2JL\7\27"+
		"\2\2KM\5\32\16\2LK\3\2\2\2MN\3\2\2\2NL\3\2\2\2NO\3\2\2\2OP\3\2\2\2PQ\7"+
		"\b\2\2QR\5\b\5\2RS\7\t\2\2ST\7\5\2\2Tb\3\2\2\2UV\7\n\2\2VW\7\27\2\2WX"+
		"\7\b\2\2XY\5\b\5\2YZ\7\t\2\2Z[\7\5\2\2[b\3\2\2\2\\]\7\n\2\2]^\7\27\2\2"+
		"^_\7\b\2\2_`\7\t\2\2`b\7\5\2\2aI\3\2\2\2aU\3\2\2\2a\\\3\2\2\2b\13\3\2"+
		"\2\2ck\7\27\2\2df\7\27\2\2eg\5\30\r\2fe\3\2\2\2gh\3\2\2\2hf\3\2\2\2hi"+
		"\3\2\2\2ik\3\2\2\2jc\3\2\2\2jd\3\2\2\2k\r\3\2\2\2lm\7\7\2\2mn\5\24\13"+
		"\2no\7\b\2\2op\5\4\3\2pq\7\t\2\2qy\3\2\2\2rs\7\7\2\2st\5\f\7\2tu\7\b\2"+
		"\2uv\5\4\3\2vw\7\t\2\2wy\3\2\2\2xl\3\2\2\2xr\3\2\2\2y\17\3\2\2\2z{\7\13"+
		"\2\2{\177\5\24\13\2|}\7\13\2\2}\177\5\f\7\2~z\3\2\2\2~|\3\2\2\2\177\21"+
		"\3\2\2\2\u0080\u0081\5\32\16\2\u0081\u0082\7\22\2\2\u0082\u0083\5\24\13"+
		"\2\u0083\u0089\3\2\2\2\u0084\u0085\5\32\16\2\u0085\u0086\7\22\2\2\u0086"+
		"\u0087\5\f\7\2\u0087\u0089\3\2\2\2\u0088\u0080\3\2\2\2\u0088\u0084\3\2"+
		"\2\2\u0089\23\3\2\2\2\u008a\u008b\b\13\1\2\u008b\u008c\7\17\2\2\u008c"+
		"\u008f\5\24\13\4\u008d\u008f\5\30\r\2\u008e\u008a\3\2\2\2\u008e\u008d"+
		"\3\2\2\2\u008f\u0095\3\2\2\2\u0090\u0091\f\5\2\2\u0091\u0092\t\2\2\2\u0092"+
		"\u0094\5\24\13\6\u0093\u0090\3\2\2\2\u0094\u0097\3\2\2\2\u0095\u0093\3"+
		"\2\2\2\u0095\u0096\3\2\2\2\u0096\25\3\2\2\2\u0097\u0095\3\2\2\2\u0098"+
		"\u0099\5\f\7\2\u0099\u009a\7\23\2\2\u009a\u009b\5\f\7\2\u009b\u00a3\3"+
		"\2\2\2\u009c\u009e\5\f\7\2\u009d\u009f\7\5\2\2\u009e\u009d\3\2\2\2\u009f"+
		"\u00a0\3\2\2\2\u00a0\u009e\3\2\2\2\u00a0\u00a1\3\2\2\2\u00a1\u00a3\3\2"+
		"\2\2\u00a2\u0098\3\2\2\2\u00a2\u009c\3\2\2\2\u00a3\27\3\2\2\2\u00a4\u00a8"+
		"\5\32\16\2\u00a5\u00a8\7\27\2\2\u00a6\u00a8\7\26\2\2\u00a7\u00a4\3\2\2"+
		"\2\u00a7\u00a5\3\2\2\2\u00a7\u00a6\3\2\2\2\u00a8\31\3\2\2\2\u00a9\u00aa"+
		"\7\3\2\2\u00aa\u00ab\7\27\2\2\u00ab\33\3\2\2\2\25\37$+/\63BGNahjx~\u0088"+
		"\u008e\u0095\u00a0\u00a2\u00a7";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}