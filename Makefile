ANTLR4= java -Xmx500M -cp "/usr/local/lib/antlr-4.7.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool

build:
	$(ANTLR4) -Dlanguage=Go -visitor parser/BeepBoop.g4