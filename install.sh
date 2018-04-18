cd /usr/local/lib
sudo curl -O http://www.antlr.org/download/antlr-4.7.1-complete.jar
echo "export CLASSPATH=\".:/usr/local/lib/antlr-4.7.1-complete.jar:$$CLASSPATH\"" >> ~/.bashrc