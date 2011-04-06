lex.yy.c: start.l
	lex start.l
start: lex.yy.c start.l
	cc -o start lex.yy.c
clean:
	rm -f lex.yy.c start
run: start
	./start main.decaf
