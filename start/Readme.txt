
Contents:
  - A start to your lex program is in start.l
  - A Makefile for it if you are using a unix shell 
       (either native or with cygwin, etc. on windows)
  - 3 sample inputs and the output it producess on those inputs

Using a unix shell with lex and yacc installed:

To create the lex.yy.c file from the lex:
> lex start.l

To compile lex.yy.c into an executable called 'start':
> cc - o start lex.yy.c

To do both of the above with the Makefile:
> Make start

To run reading from standard input and printing to standard output:
> ./start

To run reading from a redirected file but printing to standard output:
> ./start < t1.txt

To run reading from a redirected file and printing to a file:
> ./start < t1.txt > t1Out.txt


If you are using Windows, I would highly suggest you install cygwin, 
making sure to install lex and yacc.  Note: they may be called flex or bison on the install.  Flex and bison are open-sourced versions of lex and yacc.
