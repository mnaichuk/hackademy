#include<stdio.h>
#include<unistd.h>
#include<stdlib.h>
void last_word(char *s)
{
  char *str;
  int flagspace=1;
  int nowords=1;
  while(*s)
  {
    if(*s!=' ')
      nowords=0;
    if(*s!=' '&&flagspace==1)
    {
      str=s;
      flagspace=0;
    }
    if(*s==' ')
      flagspace=1;
    s++;
  }
  if (nowords)
    write(1,"\n",1);
  else
  {
    while(*str)
    {
      write(1,str,1);
      str++;
    }
    write(1,"\n",1);
  }
}
int main()
{
  last_word("");
  last_word("Ken Barbie");
}