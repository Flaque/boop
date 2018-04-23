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
		T__0=1, FLAG=2, COMMENT=3, NEWLINE=4, WHITESPACE=5, EQUALS=6, ASSIGN=7, 
		PIPE=8, LTHAN=9, GTHAN=10, LTHAN_EQUALS=11, GTHAN_EQUALS=12, LPAREN=13, 
		RPAREN=14, LSQUIG=15, RSQUIG=16, LBLOCK=17, RBLOCK=18, PLUS=19, SUB=20, 
		DIV=21, MULT=22, IMPORT=23, EXPORT=24, AS=25, IF=26, DO=27, END=28, FUNC=29, 
		RETURN=30, FOR=31, TRUE=32, FALSE=33, OR=34, AND=35, INT=36, FLOAT=37, 
		QUOTED=38, LETTERS=39, STRING=40;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "FLAG", "COMMENT", "NEWLINE", "WHITESPACE", "EQUALS", "ASSIGN", 
		"PIPE", "LTHAN", "GTHAN", "LTHAN_EQUALS", "GTHAN_EQUALS", "LPAREN", "RPAREN", 
		"LSQUIG", "RSQUIG", "LBLOCK", "RBLOCK", "PLUS", "SUB", "DIV", "MULT", 
		"IMPORT", "EXPORT", "AS", "IF", "DO", "END", "FUNC", "RETURN", "FOR", 
		"TRUE", "FALSE", "OR", "AND", "INT", "FLOAT", "QUOTED", "STRINGORNEW", 
		"LETTERS", "STRING", "ESC"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2*\u0105\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\3"+
		"\2\3\2\3\3\6\3[\n\3\r\3\16\3\\\3\3\6\3`\n\3\r\3\16\3a\3\4\3\4\7\4f\n\4"+
		"\f\4\16\4i\13\4\3\4\3\4\3\5\3\5\6\5o\n\5\r\5\16\5p\5\5s\n\5\3\6\6\6v\n"+
		"\6\r\6\16\6w\3\6\3\6\3\7\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f"+
		"\3\f\3\f\3\r\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22"+
		"\3\23\3\23\3\24\3\24\3\25\3\25\3\26\3\26\3\27\3\27\3\30\3\30\3\30\3\30"+
		"\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\33"+
		"\3\33\3\33\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36"+
		"\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3!\3!\3!\3!\3!\3\"\3\""+
		"\3\"\3\"\3\"\3\"\3#\3#\3#\3$\3$\3$\3$\3%\6%\u00df\n%\r%\16%\u00e0\3&\3"+
		"&\3&\3&\3\'\3\'\6\'\u00e9\n\'\r\'\16\'\u00ea\7\'\u00ed\n\'\f\'\16\'\u00f0"+
		"\13\'\3\'\3\'\3(\3(\5(\u00f6\n(\3)\6)\u00f9\n)\r)\16)\u00fa\3*\3*\5*\u00ff"+
		"\n*\3+\3+\3+\5+\u0104\n+\2\2,\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13"+
		"\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61"+
		"\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O\2Q)S*U\2\3\2\b\4\2"+
		"\f\f\17\17\4\2\13\13\"\"\3\2\62;\4\2C\\c|\4\2$$^^\n\2$$\61\61^^ddhhpp"+
		"ttvv\2\u010f\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2"+
		"\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2"+
		"\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2"+
		"\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2"+
		"\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3"+
		"\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2"+
		"\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\3"+
		"W\3\2\2\2\5Z\3\2\2\2\7c\3\2\2\2\tr\3\2\2\2\13u\3\2\2\2\r{\3\2\2\2\17~"+
		"\3\2\2\2\21\u0080\3\2\2\2\23\u0082\3\2\2\2\25\u0084\3\2\2\2\27\u0086\3"+
		"\2\2\2\31\u0089\3\2\2\2\33\u008c\3\2\2\2\35\u008e\3\2\2\2\37\u0090\3\2"+
		"\2\2!\u0092\3\2\2\2#\u0094\3\2\2\2%\u0096\3\2\2\2\'\u0098\3\2\2\2)\u009a"+
		"\3\2\2\2+\u009c\3\2\2\2-\u009e\3\2\2\2/\u00a0\3\2\2\2\61\u00a7\3\2\2\2"+
		"\63\u00ae\3\2\2\2\65\u00b1\3\2\2\2\67\u00b4\3\2\2\29\u00b7\3\2\2\2;\u00bb"+
		"\3\2\2\2=\u00c0\3\2\2\2?\u00c7\3\2\2\2A\u00cb\3\2\2\2C\u00d0\3\2\2\2E"+
		"\u00d6\3\2\2\2G\u00d9\3\2\2\2I\u00de\3\2\2\2K\u00e2\3\2\2\2M\u00e6\3\2"+
		"\2\2O\u00f5\3\2\2\2Q\u00f8\3\2\2\2S\u00fe\3\2\2\2U\u0100\3\2\2\2WX\7<"+
		"\2\2X\4\3\2\2\2Y[\7/\2\2ZY\3\2\2\2[\\\3\2\2\2\\Z\3\2\2\2\\]\3\2\2\2]_"+
		"\3\2\2\2^`\5Q)\2_^\3\2\2\2`a\3\2\2\2a_\3\2\2\2ab\3\2\2\2b\6\3\2\2\2cg"+
		"\7%\2\2df\n\2\2\2ed\3\2\2\2fi\3\2\2\2ge\3\2\2\2gh\3\2\2\2hj\3\2\2\2ig"+
		"\3\2\2\2jk\b\4\2\2k\b\3\2\2\2ls\t\2\2\2mo\t\2\2\2nm\3\2\2\2op\3\2\2\2"+
		"pn\3\2\2\2pq\3\2\2\2qs\3\2\2\2rl\3\2\2\2rn\3\2\2\2s\n\3\2\2\2tv\t\3\2"+
		"\2ut\3\2\2\2vw\3\2\2\2wu\3\2\2\2wx\3\2\2\2xy\3\2\2\2yz\b\6\3\2z\f\3\2"+
		"\2\2{|\7?\2\2|}\7?\2\2}\16\3\2\2\2~\177\7?\2\2\177\20\3\2\2\2\u0080\u0081"+
		"\7~\2\2\u0081\22\3\2\2\2\u0082\u0083\7>\2\2\u0083\24\3\2\2\2\u0084\u0085"+
		"\7@\2\2\u0085\26\3\2\2\2\u0086\u0087\7>\2\2\u0087\u0088\7?\2\2\u0088\30"+
		"\3\2\2\2\u0089\u008a\7@\2\2\u008a\u008b\7?\2\2\u008b\32\3\2\2\2\u008c"+
		"\u008d\7*\2\2\u008d\34\3\2\2\2\u008e\u008f\7+\2\2\u008f\36\3\2\2\2\u0090"+
		"\u0091\7}\2\2\u0091 \3\2\2\2\u0092\u0093\7\177\2\2\u0093\"\3\2\2\2\u0094"+
		"\u0095\7]\2\2\u0095$\3\2\2\2\u0096\u0097\7_\2\2\u0097&\3\2\2\2\u0098\u0099"+
		"\7-\2\2\u0099(\3\2\2\2\u009a\u009b\7/\2\2\u009b*\3\2\2\2\u009c\u009d\7"+
		"\61\2\2\u009d,\3\2\2\2\u009e\u009f\7,\2\2\u009f.\3\2\2\2\u00a0\u00a1\7"+
		"k\2\2\u00a1\u00a2\7o\2\2\u00a2\u00a3\7r\2\2\u00a3\u00a4\7q\2\2\u00a4\u00a5"+
		"\7t\2\2\u00a5\u00a6\7v\2\2\u00a6\60\3\2\2\2\u00a7\u00a8\7g\2\2\u00a8\u00a9"+
		"\7z\2\2\u00a9\u00aa\7r\2\2\u00aa\u00ab\7q\2\2\u00ab\u00ac\7t\2\2\u00ac"+
		"\u00ad\7v\2\2\u00ad\62\3\2\2\2\u00ae\u00af\7c\2\2\u00af\u00b0\7u\2\2\u00b0"+
		"\64\3\2\2\2\u00b1\u00b2\7k\2\2\u00b2\u00b3\7h\2\2\u00b3\66\3\2\2\2\u00b4"+
		"\u00b5\7f\2\2\u00b5\u00b6\7q\2\2\u00b68\3\2\2\2\u00b7\u00b8\7g\2\2\u00b8"+
		"\u00b9\7p\2\2\u00b9\u00ba\7f\2\2\u00ba:\3\2\2\2\u00bb\u00bc\7h\2\2\u00bc"+
		"\u00bd\7w\2\2\u00bd\u00be\7p\2\2\u00be\u00bf\7e\2\2\u00bf<\3\2\2\2\u00c0"+
		"\u00c1\7t\2\2\u00c1\u00c2\7g\2\2\u00c2\u00c3\7v\2\2\u00c3\u00c4\7w\2\2"+
		"\u00c4\u00c5\7t\2\2\u00c5\u00c6\7p\2\2\u00c6>\3\2\2\2\u00c7\u00c8\7h\2"+
		"\2\u00c8\u00c9\7q\2\2\u00c9\u00ca\7t\2\2\u00ca@\3\2\2\2\u00cb\u00cc\7"+
		"v\2\2\u00cc\u00cd\7t\2\2\u00cd\u00ce\7w\2\2\u00ce\u00cf\7g\2\2\u00cfB"+
		"\3\2\2\2\u00d0\u00d1\7h\2\2\u00d1\u00d2\7c\2\2\u00d2\u00d3\7n\2\2\u00d3"+
		"\u00d4\7u\2\2\u00d4\u00d5\7g\2\2\u00d5D\3\2\2\2\u00d6\u00d7\7q\2\2\u00d7"+
		"\u00d8\7t\2\2\u00d8F\3\2\2\2\u00d9\u00da\7c\2\2\u00da\u00db\7p\2\2\u00db"+
		"\u00dc\7f\2\2\u00dcH\3\2\2\2\u00dd\u00df\t\4\2\2\u00de\u00dd\3\2\2\2\u00df"+
		"\u00e0\3\2\2\2\u00e0\u00de\3\2\2\2\u00e0\u00e1\3\2\2\2\u00e1J\3\2\2\2"+
		"\u00e2\u00e3\5I%\2\u00e3\u00e4\7\60\2\2\u00e4\u00e5\5I%\2\u00e5L\3\2\2"+
		"\2\u00e6\u00ee\7$\2\2\u00e7\u00e9\5O(\2\u00e8\u00e7\3\2\2\2\u00e9\u00ea"+
		"\3\2\2\2\u00ea\u00e8\3\2\2\2\u00ea\u00eb\3\2\2\2\u00eb\u00ed\3\2\2\2\u00ec"+
		"\u00e8\3\2\2\2\u00ed\u00f0\3\2\2\2\u00ee\u00ec\3\2\2\2\u00ee\u00ef\3\2"+
		"\2\2\u00ef\u00f1\3\2\2\2\u00f0\u00ee\3\2\2\2\u00f1\u00f2\7$\2\2\u00f2"+
		"N\3\2\2\2\u00f3\u00f6\5S*\2\u00f4\u00f6\5\t\5\2\u00f5\u00f3\3\2\2\2\u00f5"+
		"\u00f4\3\2\2\2\u00f6P\3\2\2\2\u00f7\u00f9\t\5\2\2\u00f8\u00f7\3\2\2\2"+
		"\u00f9\u00fa\3\2\2\2\u00fa\u00f8\3\2\2\2\u00fa\u00fb\3\2\2\2\u00fbR\3"+
		"\2\2\2\u00fc\u00ff\5U+\2\u00fd\u00ff\n\6\2\2\u00fe\u00fc\3\2\2\2\u00fe"+
		"\u00fd\3\2\2\2\u00ffT\3\2\2\2\u0100\u0103\7^\2\2\u0101\u0104\t\7\2\2\u0102"+
		"\u0104\5Q)\2\u0103\u0101\3\2\2\2\u0103\u0102\3\2\2\2\u0104V\3\2\2\2\20"+
		"\2\\agprw\u00e0\u00ea\u00ee\u00f5\u00fa\u00fe\u0103\4\b\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}