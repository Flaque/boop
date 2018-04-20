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
		T__0=1, T__1=2, COMMENT=3, NEWLINE=4, WHITESPACE=5, IF=6, DO=7, END=8, 
		FUNC=9, RETURN=10, FOR=11, PLUS=12, SUB=13, DIV=14, MULT=15, TRUE=16, 
		FALSE=17, OR=18, AND=19, EQUALS=20, ASSIGN=21, PIPE=22, LTHAN=23, GTHAN=24, 
		LTHAN_EQUALS=25, GTHAN_EQUALS=26, LPAREN=27, RPAREN=28, LSQUIG=29, RSQUIG=30, 
		LBLOCK=31, RBLOCK=32, INT=33, STRING=34;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "T__1", "COMMENT", "NEWLINE", "WHITESPACE", "IF", "DO", "END", 
		"FUNC", "RETURN", "FOR", "PLUS", "SUB", "DIV", "MULT", "TRUE", "FALSE", 
		"OR", "AND", "EQUALS", "ASSIGN", "PIPE", "LTHAN", "GTHAN", "LTHAN_EQUALS", 
		"GTHAN_EQUALS", "LPAREN", "RPAREN", "LSQUIG", "RSQUIG", "LBLOCK", "RBLOCK", 
		"INT", "STRING"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2$\u00c7\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\3\2\3\2\3\3\3\3\3\4\3\4\7\4N\n\4\f\4\16\4Q\13\4\3\4"+
		"\3\4\3\5\3\5\6\5W\n\5\r\5\16\5X\5\5[\n\5\3\6\6\6^\n\6\r\6\16\6_\3\6\3"+
		"\6\3\7\3\7\3\7\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\13\3"+
		"\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\16\3\16"+
		"\3\16\3\16\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21"+
		"\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\24\3\24\3\24"+
		"\3\24\3\25\3\25\3\25\3\26\3\26\3\27\3\27\3\30\3\30\3\31\3\31\3\32\3\32"+
		"\3\32\3\33\3\33\3\33\3\34\3\34\3\35\3\35\3\36\3\36\3\37\3\37\3 \3 \3!"+
		"\3!\3\"\6\"\u00bf\n\"\r\"\16\"\u00c0\3#\6#\u00c4\n#\r#\16#\u00c5\2\2$"+
		"\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20"+
		"\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37"+
		"= ?!A\"C#E$\3\2\6\4\2\f\f\17\17\4\2\13\13\"\"\3\2\62;\4\2C\\c|\2\u00cc"+
		"\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2"+
		"\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2"+
		"\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2"+
		"\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2"+
		"\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3"+
		"\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\3G\3\2\2"+
		"\2\5I\3\2\2\2\7K\3\2\2\2\tZ\3\2\2\2\13]\3\2\2\2\rc\3\2\2\2\17f\3\2\2\2"+
		"\21i\3\2\2\2\23m\3\2\2\2\25r\3\2\2\2\27y\3\2\2\2\31}\3\2\2\2\33\u0081"+
		"\3\2\2\2\35\u0085\3\2\2\2\37\u0089\3\2\2\2!\u008e\3\2\2\2#\u0093\3\2\2"+
		"\2%\u0099\3\2\2\2\'\u009c\3\2\2\2)\u00a0\3\2\2\2+\u00a3\3\2\2\2-\u00a5"+
		"\3\2\2\2/\u00a7\3\2\2\2\61\u00a9\3\2\2\2\63\u00ab\3\2\2\2\65\u00ae\3\2"+
		"\2\2\67\u00b1\3\2\2\29\u00b3\3\2\2\2;\u00b5\3\2\2\2=\u00b7\3\2\2\2?\u00b9"+
		"\3\2\2\2A\u00bb\3\2\2\2C\u00be\3\2\2\2E\u00c3\3\2\2\2GH\7<\2\2H\4\3\2"+
		"\2\2IJ\7$\2\2J\6\3\2\2\2KO\7%\2\2LN\n\2\2\2ML\3\2\2\2NQ\3\2\2\2OM\3\2"+
		"\2\2OP\3\2\2\2PR\3\2\2\2QO\3\2\2\2RS\b\4\2\2S\b\3\2\2\2T[\t\2\2\2UW\t"+
		"\2\2\2VU\3\2\2\2WX\3\2\2\2XV\3\2\2\2XY\3\2\2\2Y[\3\2\2\2ZT\3\2\2\2ZV\3"+
		"\2\2\2[\n\3\2\2\2\\^\t\3\2\2]\\\3\2\2\2^_\3\2\2\2_]\3\2\2\2_`\3\2\2\2"+
		"`a\3\2\2\2ab\b\6\3\2b\f\3\2\2\2cd\7k\2\2de\7h\2\2e\16\3\2\2\2fg\7f\2\2"+
		"gh\7q\2\2h\20\3\2\2\2ij\7g\2\2jk\7p\2\2kl\7f\2\2l\22\3\2\2\2mn\7h\2\2"+
		"no\7w\2\2op\7p\2\2pq\7e\2\2q\24\3\2\2\2rs\7t\2\2st\7g\2\2tu\7v\2\2uv\7"+
		"w\2\2vw\7t\2\2wx\7p\2\2x\26\3\2\2\2yz\7h\2\2z{\7q\2\2{|\7t\2\2|\30\3\2"+
		"\2\2}~\7c\2\2~\177\7f\2\2\177\u0080\7f\2\2\u0080\32\3\2\2\2\u0081\u0082"+
		"\7u\2\2\u0082\u0083\7w\2\2\u0083\u0084\7d\2\2\u0084\34\3\2\2\2\u0085\u0086"+
		"\7f\2\2\u0086\u0087\7k\2\2\u0087\u0088\7x\2\2\u0088\36\3\2\2\2\u0089\u008a"+
		"\7o\2\2\u008a\u008b\7w\2\2\u008b\u008c\7n\2\2\u008c\u008d\7v\2\2\u008d"+
		" \3\2\2\2\u008e\u008f\7v\2\2\u008f\u0090\7t\2\2\u0090\u0091\7w\2\2\u0091"+
		"\u0092\7g\2\2\u0092\"\3\2\2\2\u0093\u0094\7h\2\2\u0094\u0095\7c\2\2\u0095"+
		"\u0096\7n\2\2\u0096\u0097\7u\2\2\u0097\u0098\7g\2\2\u0098$\3\2\2\2\u0099"+
		"\u009a\7q\2\2\u009a\u009b\7t\2\2\u009b&\3\2\2\2\u009c\u009d\7c\2\2\u009d"+
		"\u009e\7p\2\2\u009e\u009f\7f\2\2\u009f(\3\2\2\2\u00a0\u00a1\7?\2\2\u00a1"+
		"\u00a2\7?\2\2\u00a2*\3\2\2\2\u00a3\u00a4\7?\2\2\u00a4,\3\2\2\2\u00a5\u00a6"+
		"\7~\2\2\u00a6.\3\2\2\2\u00a7\u00a8\7>\2\2\u00a8\60\3\2\2\2\u00a9\u00aa"+
		"\7@\2\2\u00aa\62\3\2\2\2\u00ab\u00ac\7>\2\2\u00ac\u00ad\7?\2\2\u00ad\64"+
		"\3\2\2\2\u00ae\u00af\7@\2\2\u00af\u00b0\7?\2\2\u00b0\66\3\2\2\2\u00b1"+
		"\u00b2\7*\2\2\u00b28\3\2\2\2\u00b3\u00b4\7+\2\2\u00b4:\3\2\2\2\u00b5\u00b6"+
		"\7}\2\2\u00b6<\3\2\2\2\u00b7\u00b8\7\177\2\2\u00b8>\3\2\2\2\u00b9\u00ba"+
		"\7]\2\2\u00ba@\3\2\2\2\u00bb\u00bc\7_\2\2\u00bcB\3\2\2\2\u00bd\u00bf\t"+
		"\4\2\2\u00be\u00bd\3\2\2\2\u00bf\u00c0\3\2\2\2\u00c0\u00be\3\2\2\2\u00c0"+
		"\u00c1\3\2\2\2\u00c1D\3\2\2\2\u00c2\u00c4\t\5\2\2\u00c3\u00c2\3\2\2\2"+
		"\u00c4\u00c5\3\2\2\2\u00c5\u00c3\3\2\2\2\u00c5\u00c6\3\2\2\2\u00c6F\3"+
		"\2\2\2\t\2OXZ_\u00c0\u00c5\4\b\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}