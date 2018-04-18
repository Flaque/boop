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
		T__0=1, COMMENT=2, NEWLINE=3, WHITESPACE=4, TABSPACE=5, INT=6, PLUS=7, 
		MINUS=8, MULT=9, DIVIDE=10, ASSIGN=11, PIPE=12, LPAREN=13, RPAREN=14, 
		STRING=15, IF=16, DO=17, END=18, FUNC=19, RETURN=20, FOR=21, IS=22, UNKNOWN=23;
	public static final int
		RULE_beepboop = 0, RULE_block = 1, RULE_statement = 2, RULE_ifstat = 3, 
		RULE_funcdef = 4, RULE_returnStat = 5, RULE_assignment = 6, RULE_expr = 7, 
		RULE_pipe = 8, RULE_fncall = 9, RULE_term = 10, RULE_label = 11;
	public static final String[] ruleNames = {
		"beepboop", "block", "statement", "ifstat", "funcdef", "returnStat", "assignment", 
		"expr", "pipe", "fncall", "term", "label"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'$'", null, null, "' '", null, null, "'+'", "'-'", "'*'", "'/'", 
		"'='", "'|'", "'('", "')'", null, "'if'", "'do'", "'end'", "'func'", "'return'", 
		"'for'", "'is'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, "COMMENT", "NEWLINE", "WHITESPACE", "TABSPACE", "INT", "PLUS", 
		"MINUS", "MULT", "DIVIDE", "ASSIGN", "PIPE", "LPAREN", "RPAREN", "STRING", 
		"IF", "DO", "END", "FUNC", "RETURN", "FOR", "IS", "UNKNOWN"
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
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode EOF() { return getToken(BeepBoopParser.EOF, 0); }
		public List<TerminalNode> NEWLINE() { return getTokens(BeepBoopParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(BeepBoopParser.NEWLINE, i);
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
			setState(35);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NEWLINE:
				enterOuterAlt(_localctx, 1);
				{
				setState(25); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(24);
					match(NEWLINE);
					}
					}
					setState(27); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==NEWLINE );
				setState(29);
				block();
				setState(30);
				match(EOF);
				}
				break;
			case T__0:
			case STRING:
			case RETURN:
				enterOuterAlt(_localctx, 2);
				{
				setState(32);
				block();
				setState(33);
				match(EOF);
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

	public static class BlockContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(38); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(37);
				statement();
				}
				}
				setState(40); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << STRING) | (1L << RETURN))) != 0) );
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
			setState(54);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				_localctx = new AssignStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(42);
				assignment();
				setState(43);
				match(NEWLINE);
				}
				break;
			case 2:
				_localctx = new FncallStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(45);
				fncall();
				setState(46);
				match(NEWLINE);
				}
				break;
			case 3:
				_localctx = new ReturnStatementContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(48);
				returnStat();
				setState(49);
				match(NEWLINE);
				}
				break;
			case 4:
				_localctx = new PipeStatementContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(51);
				pipe();
				setState(52);
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
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
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
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode END() { return getToken(BeepBoopParser.END, 0); }
		public FncallIfStatementContext(IfstatContext ctx) { copyFrom(ctx); }
	}

	public final IfstatContext ifstat() throws RecognitionException {
		IfstatContext _localctx = new IfstatContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_ifstat);
		try {
			setState(68);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				_localctx = new ExprIfStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(56);
				match(IF);
				setState(57);
				expr(0);
				setState(58);
				match(DO);
				setState(59);
				block();
				setState(60);
				match(END);
				}
				break;
			case 2:
				_localctx = new FncallIfStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(62);
				match(IF);
				setState(63);
				fncall();
				setState(64);
				match(DO);
				setState(65);
				block();
				setState(66);
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

	public static class FuncdefContext extends ParserRuleContext {
		public TerminalNode FUNC() { return getToken(BeepBoopParser.FUNC, 0); }
		public TerminalNode STRING() { return getToken(BeepBoopParser.STRING, 0); }
		public TerminalNode DO() { return getToken(BeepBoopParser.DO, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode END() { return getToken(BeepBoopParser.END, 0); }
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
			enterOuterAlt(_localctx, 1);
			{
			setState(70);
			match(FUNC);
			setState(71);
			match(STRING);
			setState(73); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(72);
				label();
				}
				}
				setState(75); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==T__0 );
			setState(77);
			match(DO);
			setState(78);
			block();
			setState(79);
			match(END);
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
		enterRule(_localctx, 10, RULE_returnStat);
		try {
			setState(85);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				_localctx = new ExprReturnContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(81);
				match(RETURN);
				setState(82);
				expr(0);
				}
				break;
			case 2:
				_localctx = new FncallReturnContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(83);
				match(RETURN);
				setState(84);
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
		enterRule(_localctx, 12, RULE_assignment);
		try {
			setState(95);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(87);
				label();
				setState(88);
				match(ASSIGN);
				setState(89);
				expr(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(91);
				label();
				setState(92);
				match(ASSIGN);
				setState(93);
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
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(101);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MINUS:
				{
				_localctx = new UnaryMinusExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(98);
				match(MINUS);
				setState(99);
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
				setState(100);
				term();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(108);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new AdditiveExprContext(new ExprContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_expr);
					setState(103);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(104);
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
					setState(105);
					expr(4);
					}
					} 
				}
				setState(110);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
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
		public TerminalNode NEWLINE() { return getToken(BeepBoopParser.NEWLINE, 0); }
		public PipeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pipe; }
	}

	public final PipeContext pipe() throws RecognitionException {
		PipeContext _localctx = new PipeContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_pipe);
		try {
			setState(118);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(111);
				fncall();
				setState(112);
				match(PIPE);
				setState(113);
				fncall();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(115);
				fncall();
				setState(116);
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
		enterRule(_localctx, 18, RULE_fncall);
		int _la;
		try {
			setState(127);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(120);
				match(STRING);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(121);
				match(STRING);
				setState(123); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(122);
					term();
					}
					}
					setState(125); 
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
		enterRule(_localctx, 20, RULE_term);
		try {
			setState(132);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				_localctx = new LabelTermContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(129);
				label();
				}
				break;
			case STRING:
				_localctx = new StringTermContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(130);
				match(STRING);
				}
				break;
			case INT:
				_localctx = new IntTermContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(131);
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
		enterRule(_localctx, 22, RULE_label);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(134);
			match(T__0);
			setState(135);
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
		case 7:
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\31\u008c\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\3\2\6\2\34\n\2\r\2\16\2\35\3\2\3\2\3\2\3\2\3\2\3"+
		"\2\5\2&\n\2\3\3\6\3)\n\3\r\3\16\3*\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\5\49\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5"+
		"\5\5G\n\5\3\6\3\6\3\6\6\6L\n\6\r\6\16\6M\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3"+
		"\7\5\7X\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\5\bb\n\b\3\t\3\t\3\t\3\t\5"+
		"\th\n\t\3\t\3\t\3\t\7\tm\n\t\f\t\16\tp\13\t\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\5\ny\n\n\3\13\3\13\3\13\6\13~\n\13\r\13\16\13\177\5\13\u0082\n\13\3"+
		"\f\3\f\3\f\5\f\u0087\n\f\3\r\3\r\3\r\3\r\2\3\20\16\2\4\6\b\n\f\16\20\22"+
		"\24\26\30\2\3\3\2\t\n\2\u0090\2%\3\2\2\2\4(\3\2\2\2\68\3\2\2\2\bF\3\2"+
		"\2\2\nH\3\2\2\2\fW\3\2\2\2\16a\3\2\2\2\20g\3\2\2\2\22x\3\2\2\2\24\u0081"+
		"\3\2\2\2\26\u0086\3\2\2\2\30\u0088\3\2\2\2\32\34\7\5\2\2\33\32\3\2\2\2"+
		"\34\35\3\2\2\2\35\33\3\2\2\2\35\36\3\2\2\2\36\37\3\2\2\2\37 \5\4\3\2 "+
		"!\7\2\2\3!&\3\2\2\2\"#\5\4\3\2#$\7\2\2\3$&\3\2\2\2%\33\3\2\2\2%\"\3\2"+
		"\2\2&\3\3\2\2\2\')\5\6\4\2(\'\3\2\2\2)*\3\2\2\2*(\3\2\2\2*+\3\2\2\2+\5"+
		"\3\2\2\2,-\5\16\b\2-.\7\5\2\2.9\3\2\2\2/\60\5\24\13\2\60\61\7\5\2\2\61"+
		"9\3\2\2\2\62\63\5\f\7\2\63\64\7\5\2\2\649\3\2\2\2\65\66\5\22\n\2\66\67"+
		"\7\5\2\2\679\3\2\2\28,\3\2\2\28/\3\2\2\28\62\3\2\2\28\65\3\2\2\29\7\3"+
		"\2\2\2:;\7\22\2\2;<\5\20\t\2<=\7\23\2\2=>\5\4\3\2>?\7\24\2\2?G\3\2\2\2"+
		"@A\7\22\2\2AB\5\24\13\2BC\7\23\2\2CD\5\4\3\2DE\7\24\2\2EG\3\2\2\2F:\3"+
		"\2\2\2F@\3\2\2\2G\t\3\2\2\2HI\7\25\2\2IK\7\21\2\2JL\5\30\r\2KJ\3\2\2\2"+
		"LM\3\2\2\2MK\3\2\2\2MN\3\2\2\2NO\3\2\2\2OP\7\23\2\2PQ\5\4\3\2QR\7\24\2"+
		"\2R\13\3\2\2\2ST\7\26\2\2TX\5\20\t\2UV\7\26\2\2VX\5\24\13\2WS\3\2\2\2"+
		"WU\3\2\2\2X\r\3\2\2\2YZ\5\30\r\2Z[\7\r\2\2[\\\5\20\t\2\\b\3\2\2\2]^\5"+
		"\30\r\2^_\7\r\2\2_`\5\24\13\2`b\3\2\2\2aY\3\2\2\2a]\3\2\2\2b\17\3\2\2"+
		"\2cd\b\t\1\2de\7\n\2\2eh\5\20\t\4fh\5\26\f\2gc\3\2\2\2gf\3\2\2\2hn\3\2"+
		"\2\2ij\f\5\2\2jk\t\2\2\2km\5\20\t\6li\3\2\2\2mp\3\2\2\2nl\3\2\2\2no\3"+
		"\2\2\2o\21\3\2\2\2pn\3\2\2\2qr\5\24\13\2rs\7\16\2\2st\5\24\13\2ty\3\2"+
		"\2\2uv\5\24\13\2vw\7\5\2\2wy\3\2\2\2xq\3\2\2\2xu\3\2\2\2y\23\3\2\2\2z"+
		"\u0082\7\21\2\2{}\7\21\2\2|~\5\26\f\2}|\3\2\2\2~\177\3\2\2\2\177}\3\2"+
		"\2\2\177\u0080\3\2\2\2\u0080\u0082\3\2\2\2\u0081z\3\2\2\2\u0081{\3\2\2"+
		"\2\u0082\25\3\2\2\2\u0083\u0087\5\30\r\2\u0084\u0087\7\21\2\2\u0085\u0087"+
		"\7\b\2\2\u0086\u0083\3\2\2\2\u0086\u0084\3\2\2\2\u0086\u0085\3\2\2\2\u0087"+
		"\27\3\2\2\2\u0088\u0089\7\3\2\2\u0089\u008a\7\21\2\2\u008a\31\3\2\2\2"+
		"\20\35%*8FMWagnx\177\u0081\u0086";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}