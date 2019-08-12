#include <unistd.h>

char* print_word(char *str) {
  while ((*str == ' ') || (*str == '\n'))
  {
    str++;
  }
  while ((*str != '\0') && (*str != ' ') && (*str != '\n')) 
  {
    write(1,str,1);
    str++;
  }
  while ((*str != '\0') && ((*str == ' ') || (*str == '\n'))) 
  {
    str++;
  }
  return str;
}

void expand_str(char *str) 
{
  str = print_word(str);  
  while (*str != '\0') 
  {
    write(1,"   ",3);
    str = print_word(str);
  }
  write(1,"\n",1);
}

int main() 
{
   expand_str("     12 34 56 78     910     ");
}
