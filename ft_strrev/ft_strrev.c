#include <stdio.h>

char * ft_strrev(char * str)
{
  int n = 0;
  while (str[n] != '\0')
  {
    n++;
  }
  int i;
  char temp;
  for (i = 0; i < n/2 ; i++) 
  {
    temp = str[i];
    str[i] = str[n - 1 - i];
    str[n - 1 - i] = temp;    
  }
  return str;  
}

void print_str(char *s) 
{  
   int i = 0;
   while (s[i] != '\0') 
   {
     printf("%c",s[i]);
     i++;
   }
}

int main() 
{
   char *s = ft_strrev("1234567");
   print_str(s);
}
