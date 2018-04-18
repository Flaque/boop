// Generated from /home/flaque/projects/go/src/github.com/Flaque/boop/parser/BeepBoop.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class BeepBoopLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, COMMENT=2, NEWLINE=3, WHITESPACE=4, TABSPACE=5, INT=6, PLUS=7, 
		MINUS=8, MULT=9, DIVIDE=10, ASSIGN=11, PIPE=12, LPAREN=13, RPAREN=14, 
		STRING=15, IF=16, DO=17, END=18, FUNC=19, RETURN=20, FOR=21, IS=22, UNKNOWN=23;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "COMMENT", "NEWLINE", "WHITESPACE", "TABSPACE", "INT", "PLUS", 
		"MINUS", "MULT", "DIVIDE", "ASSIGN", "PIPE", "LPAREN", "RPAREN", "STRING", 
		"IF", "DO", "END", "FUNC", "RETURN", "FOR", "IS", "UNKNOWN"
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


	public BeepBoopLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "BeepBoop.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\31\u0087\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\3\2"+
		"\3\2\3\3\3\3\7\3\66\n\3\f\3\16\39\13\3\3\3\3\3\3\4\6\4>\n\4\r\4\16\4?"+
		"\3\5\3\5\3\5\3\5\3\6\6\6G\n\6\r\6\16\6H\3\6\3\6\3\7\6\7N\n\7\r\7\16\7"+
		"O\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17"+
		"\3\20\6\20c\n\20\r\20\16\20d\3\21\3\21\3\21\3\22\3\22\3\22\3\23\3\23\3"+
		"\23\3\23\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3"+
		"\26\3\26\3\26\3\26\3\27\3\27\3\27\3\30\3\30\3\30\3\30\2\2\31\3\3\5\4\7"+
		"\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22"+
		"#\23%\24\'\25)\26+\27-\30/\31\3\2\6\4\2\f\f\17\17\3\2\13\13\3\2\62;\4"+
		"\2C\\c|\2\u008b\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3"+
		"\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2"+
		"\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3"+
		"\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2"+
		"\2\2\2/\3\2\2\2\3\61\3\2\2\2\5\63\3\2\2\2\7=\3\2\2\2\tA\3\2\2\2\13F\3"+
		"\2\2\2\rM\3\2\2\2\17Q\3\2\2\2\21S\3\2\2\2\23U\3\2\2\2\25W\3\2\2\2\27Y"+
		"\3\2\2\2\31[\3\2\2\2\33]\3\2\2\2\35_\3\2\2\2\37b\3\2\2\2!f\3\2\2\2#i\3"+
		"\2\2\2%l\3\2\2\2\'p\3\2\2\2)u\3\2\2\2+|\3\2\2\2-\u0080\3\2\2\2/\u0083"+
		"\3\2\2\2\61\62\7&\2\2\62\4\3\2\2\2\63\67\7%\2\2\64\66\n\2\2\2\65\64\3"+
		"\2\2\2\669\3\2\2\2\67\65\3\2\2\2\678\3\2\2\28:\3\2\2\29\67\3\2\2\2:;\b"+
		"\3\2\2;\6\3\2\2\2<>\t\2\2\2=<\3\2\2\2>?\3\2\2\2?=\3\2\2\2?@\3\2\2\2@\b"+
		"\3\2\2\2AB\7\"\2\2BC\3\2\2\2CD\b\5\3\2D\n\3\2\2\2EG\t\3\2\2FE\3\2\2\2"+
		"GH\3\2\2\2HF\3\2\2\2HI\3\2\2\2IJ\3\2\2\2JK\b\6\3\2K\f\3\2\2\2LN\t\4\2"+
		"\2ML\3\2\2\2NO\3\2\2\2OM\3\2\2\2OP\3\2\2\2P\16\3\2\2\2QR\7-\2\2R\20\3"+
		"\2\2\2ST\7/\2\2T\22\3\2\2\2UV\7,\2\2V\24\3\2\2\2WX\7\61\2\2X\26\3\2\2"+
		"\2YZ\7?\2\2Z\30\3\2\2\2[\\\7~\2\2\\\32\3\2\2\2]^\7*\2\2^\34\3\2\2\2_`"+
		"\7+\2\2`\36\3\2\2\2ac\t\5\2\2ba\3\2\2\2cd\3\2\2\2db\3\2\2\2de\3\2\2\2"+
		"e \3\2\2\2fg\7k\2\2gh\7h\2\2h\"\3\2\2\2ij\7f\2\2jk\7q\2\2k$\3\2\2\2lm"+
		"\7g\2\2mn\7p\2\2no\7f\2\2o&\3\2\2\2pq\7h\2\2qr\7w\2\2rs\7p\2\2st\7e\2"+
		"\2t(\3\2\2\2uv\7t\2\2vw\7g\2\2wx\7v\2\2xy\7w\2\2yz\7t\2\2z{\7p\2\2{*\3"+
		"\2\2\2|}\7h\2\2}~\7q\2\2~\177\7t\2\2\177,\3\2\2\2\u0080\u0081\7k\2\2\u0081"+
		"\u0082\7u\2\2\u0082.\3\2\2\2\u0083\u0084\13\2\2\2\u0084\u0085\3\2\2\2"+
		"\u0085\u0086\b\30\2\2\u0086\60\3\2\2\2\b\2\67?HOd\4\b\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}