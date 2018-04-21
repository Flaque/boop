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
public class BoopLexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"COMMENT", "NEWLINE", "WHITESPACE", "EQUALS", "ASSIGN", "PIPE", "LTHAN", 
		"GTHAN", "LTHAN_EQUALS", "GTHAN_EQUALS", "LPAREN", "RPAREN", "LSQUIG", 
		"RSQUIG", "LBLOCK", "RBLOCK", "LABEL", "IF", "DO", "END", "FUNC", "RETURN", 
		"FOR", "PLUS", "SUB", "DIV", "MULT", "TRUE", "FALSE", "OR", "AND", "INT", 
		"QUOTED", "STRINGORNEW", "STRING", "ESC", "UNICODE", "HEX"
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


	public BoopLexer(CharStream input) {
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2$\u00ee\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\3\2\3\2\7\2R\n\2\f\2\16\2"+
		"U\13\2\3\2\3\2\3\3\3\3\6\3[\n\3\r\3\16\3\\\5\3_\n\3\3\4\6\4b\n\4\r\4\16"+
		"\4c\3\4\3\4\3\5\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\n\3"+
		"\13\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21"+
		"\3\22\3\22\6\22\u0087\n\22\r\22\16\22\u0088\3\23\3\23\3\23\3\24\3\24\3"+
		"\24\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3"+
		"\27\3\27\3\27\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3"+
		"\32\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3"+
		"\35\3\36\3\36\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3 \3 \3 \3 \3!\6!\u00c9"+
		"\n!\r!\16!\u00ca\3\"\3\"\6\"\u00cf\n\"\r\"\16\"\u00d0\7\"\u00d3\n\"\f"+
		"\"\16\"\u00d6\13\"\3\"\3\"\3#\3#\5#\u00dc\n#\3$\3$\5$\u00e0\n$\3%\3%\3"+
		"%\5%\u00e5\n%\3&\3&\3&\3&\3&\3&\3\'\3\'\2\2(\3\3\5\4\7\5\t\6\13\7\r\b"+
		"\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26"+
		"+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E\2G$I\2K\2M\2\3"+
		"\2\t\4\2\f\f\17\17\4\2\13\13\"\"\4\2C\\c|\3\2\62;\4\2$$^^\n\2$$\61\61"+
		"^^ddhhppttvv\5\2\62;CHch\2\u00f4\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2"+
		"\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2"+
		"\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2"+
		"\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2"+
		"\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2"+
		"\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2"+
		"\2C\3\2\2\2\2G\3\2\2\2\3O\3\2\2\2\5^\3\2\2\2\7a\3\2\2\2\tg\3\2\2\2\13"+
		"j\3\2\2\2\rl\3\2\2\2\17n\3\2\2\2\21p\3\2\2\2\23r\3\2\2\2\25u\3\2\2\2\27"+
		"x\3\2\2\2\31z\3\2\2\2\33|\3\2\2\2\35~\3\2\2\2\37\u0080\3\2\2\2!\u0082"+
		"\3\2\2\2#\u0084\3\2\2\2%\u008a\3\2\2\2\'\u008d\3\2\2\2)\u0090\3\2\2\2"+
		"+\u0094\3\2\2\2-\u0099\3\2\2\2/\u00a0\3\2\2\2\61\u00a4\3\2\2\2\63\u00a8"+
		"\3\2\2\2\65\u00ac\3\2\2\2\67\u00b0\3\2\2\29\u00b5\3\2\2\2;\u00ba\3\2\2"+
		"\2=\u00c0\3\2\2\2?\u00c3\3\2\2\2A\u00c8\3\2\2\2C\u00cc\3\2\2\2E\u00db"+
		"\3\2\2\2G\u00df\3\2\2\2I\u00e1\3\2\2\2K\u00e6\3\2\2\2M\u00ec\3\2\2\2O"+
		"S\7%\2\2PR\n\2\2\2QP\3\2\2\2RU\3\2\2\2SQ\3\2\2\2ST\3\2\2\2TV\3\2\2\2U"+
		"S\3\2\2\2VW\b\2\2\2W\4\3\2\2\2X_\t\2\2\2Y[\t\2\2\2ZY\3\2\2\2[\\\3\2\2"+
		"\2\\Z\3\2\2\2\\]\3\2\2\2]_\3\2\2\2^X\3\2\2\2^Z\3\2\2\2_\6\3\2\2\2`b\t"+
		"\3\2\2a`\3\2\2\2bc\3\2\2\2ca\3\2\2\2cd\3\2\2\2de\3\2\2\2ef\b\4\3\2f\b"+
		"\3\2\2\2gh\7?\2\2hi\7?\2\2i\n\3\2\2\2jk\7?\2\2k\f\3\2\2\2lm\7~\2\2m\16"+
		"\3\2\2\2no\7>\2\2o\20\3\2\2\2pq\7@\2\2q\22\3\2\2\2rs\7>\2\2st\7?\2\2t"+
		"\24\3\2\2\2uv\7@\2\2vw\7?\2\2w\26\3\2\2\2xy\7*\2\2y\30\3\2\2\2z{\7+\2"+
		"\2{\32\3\2\2\2|}\7}\2\2}\34\3\2\2\2~\177\7\177\2\2\177\36\3\2\2\2\u0080"+
		"\u0081\7]\2\2\u0081 \3\2\2\2\u0082\u0083\7_\2\2\u0083\"\3\2\2\2\u0084"+
		"\u0086\7<\2\2\u0085\u0087\t\4\2\2\u0086\u0085\3\2\2\2\u0087\u0088\3\2"+
		"\2\2\u0088\u0086\3\2\2\2\u0088\u0089\3\2\2\2\u0089$\3\2\2\2\u008a\u008b"+
		"\7k\2\2\u008b\u008c\7h\2\2\u008c&\3\2\2\2\u008d\u008e\7f\2\2\u008e\u008f"+
		"\7q\2\2\u008f(\3\2\2\2\u0090\u0091\7g\2\2\u0091\u0092\7p\2\2\u0092\u0093"+
		"\7f\2\2\u0093*\3\2\2\2\u0094\u0095\7h\2\2\u0095\u0096\7w\2\2\u0096\u0097"+
		"\7p\2\2\u0097\u0098\7e\2\2\u0098,\3\2\2\2\u0099\u009a\7t\2\2\u009a\u009b"+
		"\7g\2\2\u009b\u009c\7v\2\2\u009c\u009d\7w\2\2\u009d\u009e\7t\2\2\u009e"+
		"\u009f\7p\2\2\u009f.\3\2\2\2\u00a0\u00a1\7h\2\2\u00a1\u00a2\7q\2\2\u00a2"+
		"\u00a3\7t\2\2\u00a3\60\3\2\2\2\u00a4\u00a5\7c\2\2\u00a5\u00a6\7f\2\2\u00a6"+
		"\u00a7\7f\2\2\u00a7\62\3\2\2\2\u00a8\u00a9\7u\2\2\u00a9\u00aa\7w\2\2\u00aa"+
		"\u00ab\7d\2\2\u00ab\64\3\2\2\2\u00ac\u00ad\7f\2\2\u00ad\u00ae\7k\2\2\u00ae"+
		"\u00af\7x\2\2\u00af\66\3\2\2\2\u00b0\u00b1\7o\2\2\u00b1\u00b2\7w\2\2\u00b2"+
		"\u00b3\7n\2\2\u00b3\u00b4\7v\2\2\u00b48\3\2\2\2\u00b5\u00b6\7v\2\2\u00b6"+
		"\u00b7\7t\2\2\u00b7\u00b8\7w\2\2\u00b8\u00b9\7g\2\2\u00b9:\3\2\2\2\u00ba"+
		"\u00bb\7h\2\2\u00bb\u00bc\7c\2\2\u00bc\u00bd\7n\2\2\u00bd\u00be\7u\2\2"+
		"\u00be\u00bf\7g\2\2\u00bf<\3\2\2\2\u00c0\u00c1\7q\2\2\u00c1\u00c2\7t\2"+
		"\2\u00c2>\3\2\2\2\u00c3\u00c4\7c\2\2\u00c4\u00c5\7p\2\2\u00c5\u00c6\7"+
		"f\2\2\u00c6@\3\2\2\2\u00c7\u00c9\t\5\2\2\u00c8\u00c7\3\2\2\2\u00c9\u00ca"+
		"\3\2\2\2\u00ca\u00c8\3\2\2\2\u00ca\u00cb\3\2\2\2\u00cbB\3\2\2\2\u00cc"+
		"\u00d4\7$\2\2\u00cd\u00cf\5E#\2\u00ce\u00cd\3\2\2\2\u00cf\u00d0\3\2\2"+
		"\2\u00d0\u00ce\3\2\2\2\u00d0\u00d1\3\2\2\2\u00d1\u00d3\3\2\2\2\u00d2\u00ce"+
		"\3\2\2\2\u00d3\u00d6\3\2\2\2\u00d4\u00d2\3\2\2\2\u00d4\u00d5\3\2\2\2\u00d5"+
		"\u00d7\3\2\2\2\u00d6\u00d4\3\2\2\2\u00d7\u00d8\7$\2\2\u00d8D\3\2\2\2\u00d9"+
		"\u00dc\5G$\2\u00da\u00dc\5\5\3\2\u00db\u00d9\3\2\2\2\u00db\u00da\3\2\2"+
		"\2\u00dcF\3\2\2\2\u00dd\u00e0\5I%\2\u00de\u00e0\n\6\2\2\u00df\u00dd\3"+
		"\2\2\2\u00df\u00de\3\2\2\2\u00e0H\3\2\2\2\u00e1\u00e4\7^\2\2\u00e2\u00e5"+
		"\t\7\2\2\u00e3\u00e5\5K&\2\u00e4\u00e2\3\2\2\2\u00e4\u00e3\3\2\2\2\u00e5"+
		"J\3\2\2\2\u00e6\u00e7\7w\2\2\u00e7\u00e8\5M\'\2\u00e8\u00e9\5M\'\2\u00e9"+
		"\u00ea\5M\'\2\u00ea\u00eb\5M\'\2\u00ebL\3\2\2\2\u00ec\u00ed\t\b\2\2\u00ed"+
		"N\3\2\2\2\16\2S\\^c\u0088\u00ca\u00d0\u00d4\u00db\u00df\u00e4\4\b\2\2"+
		"\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}