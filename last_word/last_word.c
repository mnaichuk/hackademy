#include <unistd.h>

void print_last_word(char *str) 
{
  int index;
  int i = 0;
  if (str[0] != '\0') 
  {
    while (str[i+1] != '\0')
    {
      if (((str[i] == ' ') || (str[i] == '\n')) && (str[i+1] != ' ') && (str[i+1] != '\n')) 
      {
        index = i + 1;
      }
      i++;  
    }
    while (str[index] !='\0') 
    {
      write(1,str + index,1);
      index++;
    }
  }
  write(1,"\n",1);
}

int main() 
{
  print_last_word("1234   45677   ");
}
