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
		T__0=1, NEWLINE=2, WHITESPACE=3, INT=4, PLUS=5, MINUS=6, MULT=7, DIVIDE=8, 
		ASSIGN=9, PIPE=10, LPAREN=11, RPAREN=12, STRING=13, IF=14, DO=15, END=16, 
		FUNC=17, RETURN=18, FOR=19, IS=20;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "NEWLINE", "WHITESPACE", "INT", "PLUS", "MINUS", "MULT", "DIVIDE", 
		"ASSIGN", "PIPE", "LPAREN", "RPAREN", "STRING", "IF", "DO", "END", "FUNC", 
		"RETURN", "FOR", "IS"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'$'", null, "' '", null, "'+'", "'-'", "'*'", "'/'", "'='", "'|'", 
		"'('", "')'", null, "'if'", "'do'", "'end'", "'func'", "'return'", "'for'", 
		"'is'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, "NEWLINE", "WHITESPACE", "INT", "PLUS", "MINUS", "MULT", "DIVIDE", 
		"ASSIGN", "PIPE", "LPAREN", "RPAREN", "STRING", "IF", "DO", "END", "FUNC", 
		"RETURN", "FOR", "IS"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\26m\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\3\2\3\2\3\3\6\3/\n\3\r\3\16\3\60\3\4\3"+
		"\4\3\4\3\4\3\5\6\58\n\5\r\5\16\59\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\n"+
		"\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16\6\16M\n\16\r\16\16\16N\3\17\3\17\3"+
		"\17\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\23\3"+
		"\23\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\25\3\25\3\25\2\2\26"+
		"\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20"+
		"\37\21!\22#\23%\24\'\25)\26\3\2\5\4\2\f\f\17\17\3\2\62;\4\2C\\c|\2o\2"+
		"\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2"+
		"\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2"+
		"\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2"+
		"\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\3+\3\2\2\2\5.\3\2\2\2\7\62\3\2\2"+
		"\2\t\67\3\2\2\2\13;\3\2\2\2\r=\3\2\2\2\17?\3\2\2\2\21A\3\2\2\2\23C\3\2"+
		"\2\2\25E\3\2\2\2\27G\3\2\2\2\31I\3\2\2\2\33L\3\2\2\2\35P\3\2\2\2\37S\3"+
		"\2\2\2!V\3\2\2\2#Z\3\2\2\2%_\3\2\2\2\'f\3\2\2\2)j\3\2\2\2+,\7&\2\2,\4"+
		"\3\2\2\2-/\t\2\2\2.-\3\2\2\2/\60\3\2\2\2\60.\3\2\2\2\60\61\3\2\2\2\61"+
		"\6\3\2\2\2\62\63\7\"\2\2\63\64\3\2\2\2\64\65\b\4\2\2\65\b\3\2\2\2\668"+
		"\t\3\2\2\67\66\3\2\2\289\3\2\2\29\67\3\2\2\29:\3\2\2\2:\n\3\2\2\2;<\7"+
		"-\2\2<\f\3\2\2\2=>\7/\2\2>\16\3\2\2\2?@\7,\2\2@\20\3\2\2\2AB\7\61\2\2"+
		"B\22\3\2\2\2CD\7?\2\2D\24\3\2\2\2EF\7~\2\2F\26\3\2\2\2GH\7*\2\2H\30\3"+
		"\2\2\2IJ\7+\2\2J\32\3\2\2\2KM\t\4\2\2LK\3\2\2\2MN\3\2\2\2NL\3\2\2\2NO"+
		"\3\2\2\2O\34\3\2\2\2PQ\7k\2\2QR\7h\2\2R\36\3\2\2\2ST\7f\2\2TU\7q\2\2U"+
		" \3\2\2\2VW\7g\2\2WX\7p\2\2XY\7f\2\2Y\"\3\2\2\2Z[\7h\2\2[\\\7w\2\2\\]"+
		"\7p\2\2]^\7e\2\2^$\3\2\2\2_`\7t\2\2`a\7g\2\2ab\7v\2\2bc\7w\2\2cd\7t\2"+
		"\2de\7p\2\2e&\3\2\2\2fg\7h\2\2gh\7q\2\2hi\7t\2\2i(\3\2\2\2jk\7k\2\2kl"+
		"\7u\2\2l*\3\2\2\2\6\2\609N\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}