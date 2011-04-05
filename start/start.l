%{
	#define ID 1
	#define NULLL 2
	#define TRUE 3
	#define FALSE 4
	#define EQ 5
	#define NE 6
	#define OR 7
	#define AND 8
	#define IF 9
	#define WHILE 10
	#define ELSE 11
	#define RELOP 14 
	#define PRINT 15
	#define PRINTLN 16
	#define PROGRAM 22
	#define FUNCTION 25
	#define RETURN 26
	#define BEGINN 30
	#define END 31
	#define INT 32
	#define BOOLEAN 33
	#define VOID 34
	#define PERIOD 37
	#define COMMA 38
	#define SEMICOLON 39
	#define LPAREN 40
	#define RPAREN 41
	#define ASSIGN 42
	#define COMMENT 43
	#define QUOTE 44
	#define NOT 45
	#define DIV 46
	#define MULT 47
	#define ADD 48
	#define SUB 49
	#define MAX_NUM_IDENTS 50
	#define MAX_ID_LENGTH 51
	#define GTE 52
	#define LTE 53
	#define GT 54
	#define LT 55
	#define RSHIFT 56
	#define LSHIFT 57
	
	int line;
	int yylval;
	char idents [MAX_NUM_IDENTS][MAX_ID_LENGTH];

%}

delim     [ \t]

ws        {delim}*
letter    [A-Za-z]
digit     [0-9]
id        {letter}({letter}|{digit})*


%%

{ws}      
[iI][nN][tT]		{return(INT);}
[vV][oO][iI][dD]	{return(VOID);}
[pP][rR][iI][nN][tT]	{return(PRINT);}
[pP][rR][iI][nN][tT][lL][nN]	 {return(PRINTLN);}
[fF][uU][nN][cC][tT][iI][oO][nN] {return(FUNCTION);}
[pP][rR][oO][gG][rR][aA][mM]	 {return(PROGRAM);}
[tT][rR][uU][eE]	{return(TRUE);}
[fF][aA][lL][sS][eE]	{return(FALSE);}
[iI][fF]	 	{return(IF);}
[eE][lL][sS][eE]	{return(ELSE);}
[bB][eE][gG][iI][nN]	{return(BEGINN);}
[eE][nN][dD]		{return(END);}
[wW][hH][iI][lL][eE]	{return(WHILE);}
[rR][eE][tT][uU][rR][nN]	{return(RETURN);}
[bB][oO][oO][lL][eE][aA][nN]	{return(BOOLEAN);}
[nN][uU][lL][lL]		{return(NULLL);}
{id}     {yylval = (int) installID(); return(ID);}

"("      {return (LPAREN);}
")"      {return (RPAREN);}
"\""     {return (QUOTE);}
"//"     {return (COMMENT);}
"+"      {return (ADD);}
"-"      {return (SUB);}
"*"      {return (MULT);}
"/"      {return (DIV);}
"!"      {return (NOT);}
"="      {return (ASSIGN);}
"."      {return (PERIOD);}
","      {return (COMMA);}
";"      {return (SEMICOLON);}
";"      {return (SEMICOLON);}

"&&"     {yylval = AND; return(RELOP);}
"||"     {yylval = OR; return(RELOP);}
"=="     {yylval = EQ; return(RELOP);}
"!="     {yylval = NE; return(RELOP);}
">"      {yylval = GT; return(RELOP);}
"<"      {yylval = LT; return(RELOP);}
">="     {yylval = GTE; return(RELOP);}
"<="     {yylval = LTE; return(RELOP);}
.	 {printf("Unrecognized token: %s\n",yytext);}

%%

int installID() {
        /* function to install the lexeme, whose first
           character is pointed to by yytext, and whose
           length is yyleng, into the symbol table and
           return a pointer thereto */

           return (int) yytext;
}

int installNum() {
        /* similar to installID, but puts numerical
           constants into a separate table */
}


int yywrap() {
     /* dump contents of id and num tables */
        printf("There were %d lines\n",line);
	printf("Identifiers:\n");
       // etc.
     return 1;
}

int main() {
 	int tkn = 0;
   	while (tkn = yylex()) {
      	  switch (tkn) {
           	case INT:
             		printf("Got INT\n");
              		break;
           	case VOID:
             		printf("Got VOID\n");
              		break;
           	case PRINT:
             		printf("Got PRINT\n");
              		break;
           	case PRINTLN:
             		printf("Got PRINTLN\n");
              		break;
           	case FUNCTION:
             		printf("Got FUNCTION\n");
              		break;
		case PROGRAM:
			printf("Got PROGRAM\n");
			break;
		case TRUE:
			printf("Got TRUE\n");
			break;
		case FALSE:
			printf("Got FALSE\n");
			break;
           	case IF:
             		printf("Got IF\n");
              		break;
           	case ELSE:
              		printf("Got ELSE\n");
              		break;
           	case BEGINN:
              		printf("Got BEGIN\n");
              		break;
           	case END:
              		printf("Got END\n");
              		break;
           	case WHILE:
              		printf("Got WHILE\n");
              		break;
           	case RETURN:
              		printf("Got RETURN\n");
              		break;
           	case BOOLEAN:
              		printf("Got BOOLEAN\n");
              		break;
           	case NULLL:
              		printf("Got NULL\n");
              		break;
		case RELOP:
			printf("Got RELOP... ");
			switch (yylval) {
				case EQ:
					printf("EQ\n");
					break;
				case NE:
					printf("NE\n");
					break;
				case GT:
					printf("GT\n");
					break;
				case LT:
					printf("LT\n");
					break;
				case GTE:
					printf("GTE\n");
					break;
				case LTE:
					printf("LTE\n");
					break;
				case AND:
					printf("AND\n");
					break;
				case OR:
					printf("OR\n");
					break;
			}
			break;
		case SEMICOLON:
			printf("Got SEMICOLON\n");
			break;
		case LPAREN:
			printf("Got LPAREN\n");
			break;
		case RPAREN:
			printf("Got RPAREN\n");
			break;
		case QUOTE:
			printf("Got QUOTE\n");
			break;
		case PERIOD:
			printf("Got PERIOD\n");
			break;
		case COMMA:
			printf("Got COMMA\n");
			break;
		case NOT:
			printf("Got NOT\n");
			break;
		case MULT:
			printf("Got MULT\n");
			break;
		case DIV:
			printf("Got DIV\n");
			break;
		case ADD:
			printf("Got ADD\n");
			break;
		case SUB:
			printf("Got SUB\n");
			break;
		case ASSIGN:
			printf("Got ASSIGN\n");
			break;
		case COMMENT:
			printf("Got COMMENT\n");
			break;
		case ID:
			printf("Got ID... %s\n", yylval);
			break;
		DEFAULT:
			printf("Missing some token in case\n");
			break;
		   }
	  }
}
