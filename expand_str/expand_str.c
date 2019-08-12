#include <stdarg.h>
#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
void expand_str(char *s)
{
  while(*s)
  {
    if(*s==' ')
    {
      write(1,"   ",3);
      while(*s==' ')
    }

  }



}

int main()
{
expand_str("Hello" );
}
