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
		T__0=1, T__1=2, COMMENT=3, NEWLINE=4, WHITESPACE=5, IF=6, DO=7, END=8, 
		FUNC=9, RETURN=10, FOR=11, PLUS=12, SUB=13, DIV=14, MULT=15, TRUE=16, 
		FALSE=17, OR=18, AND=19, EQUALS=20, ASSIGN=21, PIPE=22, LTHAN=23, GTHAN=24, 
		LTHAN_EQUALS=25, GTHAN_EQUALS=26, LPAREN=27, RPAREN=28, LSQUIG=29, RSQUIG=30, 
		LBLOCK=31, RBLOCK=32, INT=33, STRING=34;
	public static final int
		RULE_boop = 0, RULE_code = 1, RULE_funcdef = 2, RULE_funcguts = 3, RULE_statement = 4, 
		RULE_ifstat = 5, RULE_returnstat = 6, RULE_structexpr = 7, RULE_assignstat = 8, 
		RULE_assign = 9, RULE_paren_fncall = 10, RULE_fncall = 11, RULE_term = 12, 
		RULE_list = 13, RULE_listterm = 14, RULE_conditional = 15, RULE_math = 16, 
		RULE_literal = 17, RULE_num = 18, RULE_boolexpr = 19, RULE_pipe = 20, 
		RULE_label = 21, RULE_quoted = 22, RULE_stringornew = 23;
	public static final String[] ruleNames = {
		"boop", "code", "funcdef", "funcguts", "statement", "ifstat", "returnstat", 
		"structexpr", "assignstat", "assign", "paren_fncall", "fncall", "term", 
		"list", "listterm", "conditional", "math", "literal", "num", "boolexpr", 
		"pipe", "label", "quoted", "stringornew"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "':'", "'\"'", null, null, null, "'if'", "'do'", "'end'", "'func'", 
		"'return'", "'for'", "'add'", "'sub'", "'div'", "'mult'", "'true'", "'false'", 
		"'or'", "'and'", "'=='", "'='", "'|'", "'<'", "'>'", "'<='", "'>='", "'('", 
		"')'", "'{'", "'}'", "'['", "']'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, "COMMENT", "NEWLINE", "WHITESPACE", "IF", "DO", "END", 
		"FUNC", "RETURN", "FOR", "PLUS", "SUB", "DIV", "MULT", "TRUE", "FALSE", 
		"OR", "AND", "EQUALS", "ASSIGN", "PIPE", "LTHAN", "GTHAN", "LTHAN_EQUALS", 
		"GTHAN_EQUALS", "LPAREN", "RPAREN", "LSQUIG", "RSQUIG", "LBLOCK", "RBLOCK", 
		"INT", "STRING"
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
			setState(76);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(54);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
				case 1:
					{
					setState(48);
					match(NEWLINE);
					}
					break;
				case 2:
					{
					setState(50); 
					_errHandler.sync(this);
					_alt = 1;
					do {
						switch (_alt) {
						case 1:
							{
							{
							setState(49);
							match(NEWLINE);
							}
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						setState(52); 
						_errHandler.sync(this);
						_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
					} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
					}
					break;
				}
				setState(62);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
				case 1:
					{
					setState(56);
					code();
					}
					break;
				case 2:
					{
					setState(58); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(57);
						code();
						}
						}
						setState(60); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << NEWLINE) | (1L << IF) | (1L << FUNC) | (1L << RETURN) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << INT) | (1L << STRING))) != 0) );
					}
					break;
				}
				setState(64);
				match(EOF);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(72);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
				case 1:
					{
					setState(66);
					code();
					}
					break;
				case 2:
					{
					setState(68); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(67);
						code();
						}
						}
						setState(70); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << NEWLINE) | (1L << IF) | (1L << FUNC) | (1L << RETURN) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << INT) | (1L << STRING))) != 0) );
					}
					break;
				}
				setState(74);
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
			setState(80);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case FUNC:
				_localctx = new FuncdefCodeContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(78);
				funcdef();
				}
				break;
			case T__0:
			case T__1:
			case NEWLINE:
			case IF:
			case RETURN:
			case SUB:
			case TRUE:
			case FALSE:
			case LPAREN:
			case LSQUIG:
			case LBLOCK:
			case INT:
			case STRING:
				_localctx = new StatementCodeContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(79);
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
			setState(112);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(82);
				match(FUNC);
				setState(83);
				label();
				setState(84);
				match(DO);
				setState(85);
				match(END);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(87);
				match(FUNC);
				setState(89); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(88);
					label();
					}
					}
					setState(91); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__0 );
				setState(93);
				match(DO);
				setState(94);
				match(END);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(96);
				match(FUNC);
				setState(97);
				label();
				setState(98);
				match(DO);
				setState(99);
				funcguts();
				setState(100);
				match(END);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(102);
				match(FUNC);
				setState(104); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(103);
					label();
					}
					}
					setState(106); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__0 );
				setState(108);
				match(DO);
				setState(109);
				funcguts();
				setState(110);
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
			setState(115); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(114);
				statement();
				}
				}
				setState(117); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << NEWLINE) | (1L << IF) | (1L << RETURN) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << INT) | (1L << STRING))) != 0) );
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
			setState(131);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				_localctx = new AssignStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(119);
				assignstat();
				}
				break;
			case 2:
				_localctx = new ReturnStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(120);
				returnstat();
				}
				break;
			case 3:
				_localctx = new FncallStatementContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(121);
				fncall();
				setState(122);
				match(NEWLINE);
				}
				break;
			case 4:
				_localctx = new IfStatementContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(124);
				ifstat();
				setState(125);
				match(NEWLINE);
				}
				break;
			case 5:
				_localctx = new PipeStatementContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(127);
				pipe();
				setState(128);
				match(NEWLINE);
				}
				break;
			case 6:
				_localctx = new NewlineStatementContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(130);
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
			setState(148);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(133);
				match(IF);
				setState(134);
				conditional(0);
				setState(135);
				match(DO);
				setState(137); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(136);
					statement();
					}
					}
					setState(139); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << NEWLINE) | (1L << IF) | (1L << RETURN) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << INT) | (1L << STRING))) != 0) );
				setState(141);
				match(END);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(143);
				match(IF);
				setState(144);
				conditional(0);
				setState(145);
				match(DO);
				setState(146);
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
			setState(171);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(150);
				match(RETURN);
				setState(151);
				term();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(152);
				match(RETURN);
				setState(153);
				term();
				setState(160);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
				case 1:
					{
					setState(154);
					match(NEWLINE);
					}
					break;
				case 2:
					{
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
				setState(162);
				match(RETURN);
				setState(169);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
				case 1:
					{
					setState(163);
					match(NEWLINE);
					}
					break;
				case 2:
					{
					setState(165); 
					_errHandler.sync(this);
					_alt = 1;
					do {
						switch (_alt) {
						case 1:
							{
							{
							setState(164);
							match(NEWLINE);
							}
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						setState(167); 
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
			setState(207);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(173);
				match(LSQUIG);
				setState(174);
				match(RSQUIG);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(175);
				match(LSQUIG);
				setState(176);
				match(NEWLINE);
				setState(183);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
				case 1:
					{
					setState(177);
					assignstat();
					}
					break;
				case 2:
					{
					setState(179); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(178);
						assignstat();
						}
						}
						setState(181); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( _la==T__0 );
					}
					break;
				}
				setState(185);
				match(RSQUIG);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(187);
				match(LSQUIG);
				setState(194);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
				case 1:
					{
					setState(188);
					assignstat();
					}
					break;
				case 2:
					{
					setState(190); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(189);
						assignstat();
						}
						}
						setState(192); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( _la==T__0 );
					}
					break;
				}
				setState(196);
				match(RSQUIG);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(198);
				match(LSQUIG);
				setState(199);
				assign();
				setState(200);
				match(RSQUIG);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(202);
				match(LSQUIG);
				setState(203);
				match(NEWLINE);
				setState(204);
				assign();
				setState(205);
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
			setState(209);
			assign();
			setState(210);
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
			setState(220);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				_localctx = new ExprAssignContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(212);
				label();
				setState(213);
				match(ASSIGN);
				setState(214);
				term();
				}
				break;
			case 2:
				_localctx = new FncallAssignContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(216);
				label();
				setState(217);
				match(ASSIGN);
				setState(218);
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
			setState(222);
			match(LPAREN);
			setState(223);
			fncall();
			setState(224);
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
			setState(233);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(226);
				label();
				setState(228); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(227);
					term();
					}
					}
					setState(230); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << INT) | (1L << STRING))) != 0) );
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(232);
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
			setState(241);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
			case 1:
				_localctx = new LabelTermContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(235);
				label();
				}
				break;
			case 2:
				_localctx = new LiteralTermContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(236);
				literal();
				}
				break;
			case 3:
				_localctx = new MathTermContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(237);
				math(0);
				}
				break;
			case 4:
				_localctx = new ParenfncallTermContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(238);
				paren_fncall();
				}
				break;
			case 5:
				_localctx = new StructTermContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(239);
				structexpr();
				}
				break;
			case 6:
				_localctx = new ListTermContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(240);
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
			setState(262);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(243);
				match(LBLOCK);
				setState(244);
				match(RBLOCK);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(245);
				match(LBLOCK);
				setState(247); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(246);
					listterm();
					}
					}
					setState(249); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << INT) | (1L << STRING))) != 0) );
				setState(251);
				match(RBLOCK);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(253);
				match(LBLOCK);
				setState(254);
				match(NEWLINE);
				setState(256); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(255);
					listterm();
					}
					}
					setState(258); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << SUB) | (1L << TRUE) | (1L << FALSE) | (1L << LPAREN) | (1L << LSQUIG) | (1L << LBLOCK) | (1L << INT) | (1L << STRING))) != 0) );
				setState(260);
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
			setState(268);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(264);
				term();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(265);
				term();
				setState(266);
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
			setState(292);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				{
				_localctx = new EqualsCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(271);
				term();
				setState(272);
				match(EQUALS);
				setState(273);
				term();
				}
				break;
			case 2:
				{
				_localctx = new LthanCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(275);
				term();
				setState(276);
				match(LTHAN);
				setState(277);
				term();
				}
				break;
			case 3:
				{
				_localctx = new GthanCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(279);
				term();
				setState(280);
				match(GTHAN);
				setState(281);
				term();
				}
				break;
			case 4:
				{
				_localctx = new LthanequalsCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(283);
				term();
				setState(284);
				match(LTHAN_EQUALS);
				setState(285);
				term();
				}
				break;
			case 5:
				{
				_localctx = new GthanequalsCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(287);
				term();
				setState(288);
				match(GTHAN_EQUALS);
				setState(289);
				term();
				}
				break;
			case 6:
				{
				_localctx = new BoolCondContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(291);
				boolexpr();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(302);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,35,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(300);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
					case 1:
						{
						_localctx = new OrCondContext(new ConditionalContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_conditional);
						setState(294);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(295);
						match(OR);
						setState(296);
						conditional(4);
						}
						break;
					case 2:
						{
						_localctx = new AndCondContext(new ConditionalContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_conditional);
						setState(297);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(298);
						match(AND);
						setState(299);
						conditional(3);
						}
						break;
					}
					} 
				}
				setState(304);
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
			setState(309);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SUB:
				{
				_localctx = new UnarySubMathContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(306);
				match(SUB);
				setState(307);
				math(2);
				}
				break;
			case INT:
				{
				_localctx = new SoloMathContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(308);
				num();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(316);
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
					setState(311);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(312);
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
					setState(313);
					math(4);
					}
					} 
				}
				setState(318);
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
		public QuotedContext quoted() {
			return getRuleContext(QuotedContext.class,0);
		}
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_literal);
		try {
			setState(323);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(319);
				num();
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 2);
				{
				setState(320);
				match(STRING);
				}
				break;
			case TRUE:
			case FALSE:
				enterOuterAlt(_localctx, 3);
				{
				setState(321);
				boolexpr();
				}
				break;
			case T__1:
				enterOuterAlt(_localctx, 4);
				{
				setState(322);
				quoted();
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
			setState(325);
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
			setState(327);
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
			setState(341);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,39,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(329);
				term();
				setState(330);
				match(PIPE);
				setState(331);
				fncall();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(333);
				term();
				setState(334);
				match(PIPE);
				setState(335);
				pipe();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(337);
				fncall();
				setState(338);
				match(PIPE);
				setState(339);
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
		public TerminalNode STRING() { return getToken(BeepBoopParser.STRING, 0); }
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
			setState(343);
			match(T__0);
			setState(344);
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

	public static class QuotedContext extends ParserRuleContext {
		public List<StringornewContext> stringornew() {
			return getRuleContexts(StringornewContext.class);
		}
		public StringornewContext stringornew(int i) {
			return getRuleContext(StringornewContext.class,i);
		}
		public QuotedContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_quoted; }
	}

	public final QuotedContext quoted() throws RecognitionException {
		QuotedContext _localctx = new QuotedContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_quoted);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(346);
			match(T__1);
			setState(353);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,41,_ctx) ) {
			case 1:
				{
				setState(347);
				stringornew();
				}
				break;
			case 2:
				{
				setState(349); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(348);
					stringornew();
					}
					}
					setState(351); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==NEWLINE || _la==STRING );
				}
				break;
			}
			setState(355);
			match(T__1);
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

	public static class StringornewContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(BeepBoopParser.STRING, 0); }
		public TerminalNode NEWLINE() { return getToken(BeepBoopParser.NEWLINE, 0); }
		public StringornewContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stringornew; }
	}

	public final StringornewContext stringornew() throws RecognitionException {
		StringornewContext _localctx = new StringornewContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_stringornew);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(357);
			_la = _input.LA(1);
			if ( !(_la==NEWLINE || _la==STRING) ) {
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3$\u016a\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\3\2\3\2\6\2\65\n\2\r\2\16\2\66\5\29\n\2\3\2\3\2\6\2=\n\2\r\2\16\2>\5"+
		"\2A\n\2\3\2\3\2\3\2\3\2\6\2G\n\2\r\2\16\2H\5\2K\n\2\3\2\3\2\5\2O\n\2\3"+
		"\3\3\3\5\3S\n\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\6\4\\\n\4\r\4\16\4]\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\6\4k\n\4\r\4\16\4l\3\4\3\4\3\4"+
		"\3\4\5\4s\n\4\3\5\6\5v\n\5\r\5\16\5w\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\5\6\u0086\n\6\3\7\3\7\3\7\3\7\6\7\u008c\n\7\r\7\16\7\u008d"+
		"\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7\u0097\n\7\3\b\3\b\3\b\3\b\3\b\3\b\6\b"+
		"\u009f\n\b\r\b\16\b\u00a0\5\b\u00a3\n\b\3\b\3\b\3\b\6\b\u00a8\n\b\r\b"+
		"\16\b\u00a9\5\b\u00ac\n\b\5\b\u00ae\n\b\3\t\3\t\3\t\3\t\3\t\3\t\6\t\u00b6"+
		"\n\t\r\t\16\t\u00b7\5\t\u00ba\n\t\3\t\3\t\3\t\3\t\3\t\6\t\u00c1\n\t\r"+
		"\t\16\t\u00c2\5\t\u00c5\n\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\5\t\u00d2\n\t\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5"+
		"\13\u00df\n\13\3\f\3\f\3\f\3\f\3\r\3\r\6\r\u00e7\n\r\r\r\16\r\u00e8\3"+
		"\r\5\r\u00ec\n\r\3\16\3\16\3\16\3\16\3\16\3\16\5\16\u00f4\n\16\3\17\3"+
		"\17\3\17\3\17\6\17\u00fa\n\17\r\17\16\17\u00fb\3\17\3\17\3\17\3\17\3\17"+
		"\6\17\u0103\n\17\r\17\16\17\u0104\3\17\3\17\5\17\u0109\n\17\3\20\3\20"+
		"\3\20\3\20\5\20\u010f\n\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\5\21"+
		"\u0127\n\21\3\21\3\21\3\21\3\21\3\21\3\21\7\21\u012f\n\21\f\21\16\21\u0132"+
		"\13\21\3\22\3\22\3\22\3\22\5\22\u0138\n\22\3\22\3\22\3\22\7\22\u013d\n"+
		"\22\f\22\16\22\u0140\13\22\3\23\3\23\3\23\3\23\5\23\u0146\n\23\3\24\3"+
		"\24\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3"+
		"\26\5\26\u0158\n\26\3\27\3\27\3\27\3\30\3\30\3\30\6\30\u0160\n\30\r\30"+
		"\16\30\u0161\5\30\u0164\n\30\3\30\3\30\3\31\3\31\3\31\2\4 \"\32\2\4\6"+
		"\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\2\5\3\2\16\17\3\2\22\23"+
		"\4\2\6\6$$\2\u0191\2N\3\2\2\2\4R\3\2\2\2\6r\3\2\2\2\bu\3\2\2\2\n\u0085"+
		"\3\2\2\2\f\u0096\3\2\2\2\16\u00ad\3\2\2\2\20\u00d1\3\2\2\2\22\u00d3\3"+
		"\2\2\2\24\u00de\3\2\2\2\26\u00e0\3\2\2\2\30\u00eb\3\2\2\2\32\u00f3\3\2"+
		"\2\2\34\u0108\3\2\2\2\36\u010e\3\2\2\2 \u0126\3\2\2\2\"\u0137\3\2\2\2"+
		"$\u0145\3\2\2\2&\u0147\3\2\2\2(\u0149\3\2\2\2*\u0157\3\2\2\2,\u0159\3"+
		"\2\2\2.\u015c\3\2\2\2\60\u0167\3\2\2\2\629\7\6\2\2\63\65\7\6\2\2\64\63"+
		"\3\2\2\2\65\66\3\2\2\2\66\64\3\2\2\2\66\67\3\2\2\2\679\3\2\2\28\62\3\2"+
		"\2\28\64\3\2\2\29@\3\2\2\2:A\5\4\3\2;=\5\4\3\2<;\3\2\2\2=>\3\2\2\2><\3"+
		"\2\2\2>?\3\2\2\2?A\3\2\2\2@:\3\2\2\2@<\3\2\2\2AB\3\2\2\2BC\7\2\2\3CO\3"+
		"\2\2\2DK\5\4\3\2EG\5\4\3\2FE\3\2\2\2GH\3\2\2\2HF\3\2\2\2HI\3\2\2\2IK\3"+
		"\2\2\2JD\3\2\2\2JF\3\2\2\2KL\3\2\2\2LM\7\2\2\3MO\3\2\2\2N8\3\2\2\2NJ\3"+
		"\2\2\2O\3\3\2\2\2PS\5\6\4\2QS\5\n\6\2RP\3\2\2\2RQ\3\2\2\2S\5\3\2\2\2T"+
		"U\7\13\2\2UV\5,\27\2VW\7\t\2\2WX\7\n\2\2Xs\3\2\2\2Y[\7\13\2\2Z\\\5,\27"+
		"\2[Z\3\2\2\2\\]\3\2\2\2][\3\2\2\2]^\3\2\2\2^_\3\2\2\2_`\7\t\2\2`a\7\n"+
		"\2\2as\3\2\2\2bc\7\13\2\2cd\5,\27\2de\7\t\2\2ef\5\b\5\2fg\7\n\2\2gs\3"+
		"\2\2\2hj\7\13\2\2ik\5,\27\2ji\3\2\2\2kl\3\2\2\2lj\3\2\2\2lm\3\2\2\2mn"+
		"\3\2\2\2no\7\t\2\2op\5\b\5\2pq\7\n\2\2qs\3\2\2\2rT\3\2\2\2rY\3\2\2\2r"+
		"b\3\2\2\2rh\3\2\2\2s\7\3\2\2\2tv\5\n\6\2ut\3\2\2\2vw\3\2\2\2wu\3\2\2\2"+
		"wx\3\2\2\2x\t\3\2\2\2y\u0086\5\22\n\2z\u0086\5\16\b\2{|\5\30\r\2|}\7\6"+
		"\2\2}\u0086\3\2\2\2~\177\5\f\7\2\177\u0080\7\6\2\2\u0080\u0086\3\2\2\2"+
		"\u0081\u0082\5*\26\2\u0082\u0083\7\6\2\2\u0083\u0086\3\2\2\2\u0084\u0086"+
		"\7\6\2\2\u0085y\3\2\2\2\u0085z\3\2\2\2\u0085{\3\2\2\2\u0085~\3\2\2\2\u0085"+
		"\u0081\3\2\2\2\u0085\u0084\3\2\2\2\u0086\13\3\2\2\2\u0087\u0088\7\b\2"+
		"\2\u0088\u0089\5 \21\2\u0089\u008b\7\t\2\2\u008a\u008c\5\n\6\2\u008b\u008a"+
		"\3\2\2\2\u008c\u008d\3\2\2\2\u008d\u008b\3\2\2\2\u008d\u008e\3\2\2\2\u008e"+
		"\u008f\3\2\2\2\u008f\u0090\7\n\2\2\u0090\u0097\3\2\2\2\u0091\u0092\7\b"+
		"\2\2\u0092\u0093\5 \21\2\u0093\u0094\7\t\2\2\u0094\u0095\7\n\2\2\u0095"+
		"\u0097\3\2\2\2\u0096\u0087\3\2\2\2\u0096\u0091\3\2\2\2\u0097\r\3\2\2\2"+
		"\u0098\u0099\7\f\2\2\u0099\u00ae\5\32\16\2\u009a\u009b\7\f\2\2\u009b\u00a2"+
		"\5\32\16\2\u009c\u00a3\7\6\2\2\u009d\u009f\7\6\2\2\u009e\u009d\3\2\2\2"+
		"\u009f\u00a0\3\2\2\2\u00a0\u009e\3\2\2\2\u00a0\u00a1\3\2\2\2\u00a1\u00a3"+
		"\3\2\2\2\u00a2\u009c\3\2\2\2\u00a2\u009e\3\2\2\2\u00a3\u00ae\3\2\2\2\u00a4"+
		"\u00ab\7\f\2\2\u00a5\u00ac\7\6\2\2\u00a6\u00a8\7\6\2\2\u00a7\u00a6\3\2"+
		"\2\2\u00a8\u00a9\3\2\2\2\u00a9\u00a7\3\2\2\2\u00a9\u00aa\3\2\2\2\u00aa"+
		"\u00ac\3\2\2\2\u00ab\u00a5\3\2\2\2\u00ab\u00a7\3\2\2\2\u00ac\u00ae\3\2"+
		"\2\2\u00ad\u0098\3\2\2\2\u00ad\u009a\3\2\2\2\u00ad\u00a4\3\2\2\2\u00ae"+
		"\17\3\2\2\2\u00af\u00b0\7\37\2\2\u00b0\u00d2\7 \2\2\u00b1\u00b2\7\37\2"+
		"\2\u00b2\u00b9\7\6\2\2\u00b3\u00ba\5\22\n\2\u00b4\u00b6\5\22\n\2\u00b5"+
		"\u00b4\3\2\2\2\u00b6\u00b7\3\2\2\2\u00b7\u00b5\3\2\2\2\u00b7\u00b8\3\2"+
		"\2\2\u00b8\u00ba\3\2\2\2\u00b9\u00b3\3\2\2\2\u00b9\u00b5\3\2\2\2\u00ba"+
		"\u00bb\3\2\2\2\u00bb\u00bc\7 \2\2\u00bc\u00d2\3\2\2\2\u00bd\u00c4\7\37"+
		"\2\2\u00be\u00c5\5\22\n\2\u00bf\u00c1\5\22\n\2\u00c0\u00bf\3\2\2\2\u00c1"+
		"\u00c2\3\2\2\2\u00c2\u00c0\3\2\2\2\u00c2\u00c3\3\2\2\2\u00c3\u00c5\3\2"+
		"\2\2\u00c4\u00be\3\2\2\2\u00c4\u00c0\3\2\2\2\u00c5\u00c6\3\2\2\2\u00c6"+
		"\u00c7\7 \2\2\u00c7\u00d2\3\2\2\2\u00c8\u00c9\7\37\2\2\u00c9\u00ca\5\24"+
		"\13\2\u00ca\u00cb\7 \2\2\u00cb\u00d2\3\2\2\2\u00cc\u00cd\7\37\2\2\u00cd"+
		"\u00ce\7\6\2\2\u00ce\u00cf\5\24\13\2\u00cf\u00d0\7 \2\2\u00d0\u00d2\3"+
		"\2\2\2\u00d1\u00af\3\2\2\2\u00d1\u00b1\3\2\2\2\u00d1\u00bd\3\2\2\2\u00d1"+
		"\u00c8\3\2\2\2\u00d1\u00cc\3\2\2\2\u00d2\21\3\2\2\2\u00d3\u00d4\5\24\13"+
		"\2\u00d4\u00d5\7\6\2\2\u00d5\23\3\2\2\2\u00d6\u00d7\5,\27\2\u00d7\u00d8"+
		"\7\27\2\2\u00d8\u00d9\5\32\16\2\u00d9\u00df\3\2\2\2\u00da\u00db\5,\27"+
		"\2\u00db\u00dc\7\27\2\2\u00dc\u00dd\5\30\r\2\u00dd\u00df\3\2\2\2\u00de"+
		"\u00d6\3\2\2\2\u00de\u00da\3\2\2\2\u00df\25\3\2\2\2\u00e0\u00e1\7\35\2"+
		"\2\u00e1\u00e2\5\30\r\2\u00e2\u00e3\7\36\2\2\u00e3\27\3\2\2\2\u00e4\u00e6"+
		"\5,\27\2\u00e5\u00e7\5\32\16\2\u00e6\u00e5\3\2\2\2\u00e7\u00e8\3\2\2\2"+
		"\u00e8\u00e6\3\2\2\2\u00e8\u00e9\3\2\2\2\u00e9\u00ec\3\2\2\2\u00ea\u00ec"+
		"\5,\27\2\u00eb\u00e4\3\2\2\2\u00eb\u00ea\3\2\2\2\u00ec\31\3\2\2\2\u00ed"+
		"\u00f4\5,\27\2\u00ee\u00f4\5$\23\2\u00ef\u00f4\5\"\22\2\u00f0\u00f4\5"+
		"\26\f\2\u00f1\u00f4\5\20\t\2\u00f2\u00f4\5\34\17\2\u00f3\u00ed\3\2\2\2"+
		"\u00f3\u00ee\3\2\2\2\u00f3\u00ef\3\2\2\2\u00f3\u00f0\3\2\2\2\u00f3\u00f1"+
		"\3\2\2\2\u00f3\u00f2\3\2\2\2\u00f4\33\3\2\2\2\u00f5\u00f6\7!\2\2\u00f6"+
		"\u0109\7\"\2\2\u00f7\u00f9\7!\2\2\u00f8\u00fa\5\36\20\2\u00f9\u00f8\3"+
		"\2\2\2\u00fa\u00fb\3\2\2\2\u00fb\u00f9\3\2\2\2\u00fb\u00fc\3\2\2\2\u00fc"+
		"\u00fd\3\2\2\2\u00fd\u00fe\7\"\2\2\u00fe\u0109\3\2\2\2\u00ff\u0100\7!"+
		"\2\2\u0100\u0102\7\6\2\2\u0101\u0103\5\36\20\2\u0102\u0101\3\2\2\2\u0103"+
		"\u0104\3\2\2\2\u0104\u0102\3\2\2\2\u0104\u0105\3\2\2\2\u0105\u0106\3\2"+
		"\2\2\u0106\u0107\7\"\2\2\u0107\u0109\3\2\2\2\u0108\u00f5\3\2\2\2\u0108"+
		"\u00f7\3\2\2\2\u0108\u00ff\3\2\2\2\u0109\35\3\2\2\2\u010a\u010f\5\32\16"+
		"\2\u010b\u010c\5\32\16\2\u010c\u010d\7\6\2\2\u010d\u010f\3\2\2\2\u010e"+
		"\u010a\3\2\2\2\u010e\u010b\3\2\2\2\u010f\37\3\2\2\2\u0110\u0111\b\21\1"+
		"\2\u0111\u0112\5\32\16\2\u0112\u0113\7\26\2\2\u0113\u0114\5\32\16\2\u0114"+
		"\u0127\3\2\2\2\u0115\u0116\5\32\16\2\u0116\u0117\7\31\2\2\u0117\u0118"+
		"\5\32\16\2\u0118\u0127\3\2\2\2\u0119\u011a\5\32\16\2\u011a\u011b\7\32"+
		"\2\2\u011b\u011c\5\32\16\2\u011c\u0127\3\2\2\2\u011d\u011e\5\32\16\2\u011e"+
		"\u011f\7\33\2\2\u011f\u0120\5\32\16\2\u0120\u0127\3\2\2\2\u0121\u0122"+
		"\5\32\16\2\u0122\u0123\7\34\2\2\u0123\u0124\5\32\16\2\u0124\u0127\3\2"+
		"\2\2\u0125\u0127\5(\25\2\u0126\u0110\3\2\2\2\u0126\u0115\3\2\2\2\u0126"+
		"\u0119\3\2\2\2\u0126\u011d\3\2\2\2\u0126\u0121\3\2\2\2\u0126\u0125\3\2"+
		"\2\2\u0127\u0130\3\2\2\2\u0128\u0129\f\5\2\2\u0129\u012a\7\24\2\2\u012a"+
		"\u012f\5 \21\6\u012b\u012c\f\4\2\2\u012c\u012d\7\25\2\2\u012d\u012f\5"+
		" \21\5\u012e\u0128\3\2\2\2\u012e\u012b\3\2\2\2\u012f\u0132\3\2\2\2\u0130"+
		"\u012e\3\2\2\2\u0130\u0131\3\2\2\2\u0131!\3\2\2\2\u0132\u0130\3\2\2\2"+
		"\u0133\u0134\b\22\1\2\u0134\u0135\7\17\2\2\u0135\u0138\5\"\22\4\u0136"+
		"\u0138\5&\24\2\u0137\u0133\3\2\2\2\u0137\u0136\3\2\2\2\u0138\u013e\3\2"+
		"\2\2\u0139\u013a\f\5\2\2\u013a\u013b\t\2\2\2\u013b\u013d\5\"\22\6\u013c"+
		"\u0139\3\2\2\2\u013d\u0140\3\2\2\2\u013e\u013c\3\2\2\2\u013e\u013f\3\2"+
		"\2\2\u013f#\3\2\2\2\u0140\u013e\3\2\2\2\u0141\u0146\5&\24\2\u0142\u0146"+
		"\7$\2\2\u0143\u0146\5(\25\2\u0144\u0146\5.\30\2\u0145\u0141\3\2\2\2\u0145"+
		"\u0142\3\2\2\2\u0145\u0143\3\2\2\2\u0145\u0144\3\2\2\2\u0146%\3\2\2\2"+
		"\u0147\u0148\7#\2\2\u0148\'\3\2\2\2\u0149\u014a\t\3\2\2\u014a)\3\2\2\2"+
		"\u014b\u014c\5\32\16\2\u014c\u014d\7\30\2\2\u014d\u014e\5\30\r\2\u014e"+
		"\u0158\3\2\2\2\u014f\u0150\5\32\16\2\u0150\u0151\7\30\2\2\u0151\u0152"+
		"\5*\26\2\u0152\u0158\3\2\2\2\u0153\u0154\5\30\r\2\u0154\u0155\7\30\2\2"+
		"\u0155\u0156\5*\26\2\u0156\u0158\3\2\2\2\u0157\u014b\3\2\2\2\u0157\u014f"+
		"\3\2\2\2\u0157\u0153\3\2\2\2\u0158+\3\2\2\2\u0159\u015a\7\3\2\2\u015a"+
		"\u015b\7$\2\2\u015b-\3\2\2\2\u015c\u0163\7\4\2\2\u015d\u0164\5\60\31\2"+
		"\u015e\u0160\5\60\31\2\u015f\u015e\3\2\2\2\u0160\u0161\3\2\2\2\u0161\u015f"+
		"\3\2\2\2\u0161\u0162\3\2\2\2\u0162\u0164\3\2\2\2\u0163\u015d\3\2\2\2\u0163"+
		"\u015f\3\2\2\2\u0164\u0165\3\2\2\2\u0165\u0166\7\4\2\2\u0166/\3\2\2\2"+
		"\u0167\u0168\t\4\2\2\u0168\61\3\2\2\2,\668>@HJNR]lrw\u0085\u008d\u0096"+
		"\u00a0\u00a2\u00a9\u00ab\u00ad\u00b7\u00b9\u00c2\u00c4\u00d1\u00de\u00e8"+
		"\u00eb\u00f3\u00fb\u0104\u0108\u010e\u0126\u012e\u0130\u0137\u013e\u0145"+
		"\u0157\u0161\u0163";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}